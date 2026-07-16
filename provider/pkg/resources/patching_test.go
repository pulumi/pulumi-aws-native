package resources

import (
	"testing"

	"github.com/mattbaird/jsonpatch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/default_tags"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
)

func TestCalcPatch(t *testing.T) {
	type args struct {
		oldInputs resource.PropertyMap
		newInputs resource.PropertyMap
		spec      metadata.CloudAPIResource
		types     map[string]metadata.CloudAPIType
	}
	type implementation func(t *testing.T, args args) []jsonpatch.JsonPatchOperation
	typed := func(t *testing.T, args args) []jsonpatch.JsonPatchOperation {
		t.Helper()
		patch, err := CalcPatch(args.oldInputs, args.newInputs, args.spec, args.types)
		assert.NoError(t, err)
		return patch
	}
	untyped := func(t *testing.T, args args) []jsonpatch.JsonPatchOperation {
		t.Helper()
		// Convert all keys to upper camel case to match the behavior of the SdkToCfn function.
		keysToUpperCamel := func(s string) (string, bool) {
			return naming.ToUpperCamel(s), true
		}
		patch, err := CalculateUntypedPatch(ExtensionResourceInputs{
			Properties: args.oldInputs.MapRepl(keysToUpperCamel, nil),
		}, ExtensionResourceInputs{
			Properties: args.newInputs.MapRepl(keysToUpperCamel, nil),
			CreateOnly: Map(args.spec.CreateOnly, naming.ToUpperCamel),
			WriteOnly:  Map(args.spec.WriteOnly, naming.ToUpperCamel),
		})
		assert.NoError(t, err)
		return patch
	}

	for name, run := range map[string]implementation{
		"Typed":   typed,
		"Untyped": untyped,
	} {
		t.Run(name, func(t *testing.T) {
			t.Run("no diff", func(t *testing.T) {
				patch := run(t, args{
					oldInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("value1"),
					},
					newInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("value1"),
					},
					spec: metadata.CloudAPIResource{
						Inputs: map[string]schema.PropertySpec{
							"prop1": {TypeSpec: schema.TypeSpec{Type: "string"}},
						},
					}})
				assert.Empty(t, patch)
			})
			t.Run("no diff with must send props", func(t *testing.T) {
				expected := []jsonpatch.JsonPatchOperation{
					{Operation: "add", Path: "/Prop1", Value: "1"},
				}
				patch := run(t, args{
					oldInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2"),
					},
					newInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2"),
					},
					spec: metadata.CloudAPIResource{
						Inputs: map[string]schema.PropertySpec{
							"prop1": {TypeSpec: schema.TypeSpec{Type: "string"}},
							"prop2": {TypeSpec: schema.TypeSpec{Type: "string"}},
						},
						WriteOnly: []string{"prop1"},
					}})
				assert.Equal(t, expected, patch)
			})
			t.Run("always sends write-only properties", func(t *testing.T) {
				expected := []jsonpatch.JsonPatchOperation{
					{Operation: "add", Path: "/Prop1", Value: "1"},
					{Operation: "replace", Path: "/Prop2", Value: "2b"},
				}
				patch := run(t, args{
					oldInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2a"),
					},
					newInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2b"),
					},
					spec: metadata.CloudAPIResource{
						Inputs: map[string]schema.PropertySpec{
							"prop1": {TypeSpec: schema.TypeSpec{Type: "string"}},
							"prop2": {TypeSpec: schema.TypeSpec{Type: "string"}},
						},
						WriteOnly: []string{"prop1"},
					}})
				assert.Equal(t, expected, patch)
			})
			t.Run("don't send write-only, create-only properties", func(t *testing.T) {
				expected := []jsonpatch.JsonPatchOperation{
					{Operation: "replace", Path: "/Prop2", Value: "2b"},
				}
				patch := run(t, args{
					oldInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2a"),
					},
					newInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2b"),
					},
					spec: metadata.CloudAPIResource{
						Inputs: map[string]schema.PropertySpec{
							"prop1": {TypeSpec: schema.TypeSpec{Type: "string"}},
							"prop2": {TypeSpec: schema.TypeSpec{Type: "string"}},
						},
						WriteOnly:  []string{"prop1"},
						CreateOnly: []string{"prop1"},
					}})
				assert.Equal(t, expected, patch)
			})
		})
	}
}

func TestCalcPatchWithActualOutputs(t *testing.T) {
	t.Run("write-only falls back to old desired input", func(t *testing.T) {
		spec := metadata.CloudAPIResource{
			Inputs: map[string]schema.PropertySpec{
				"password": {TypeSpec: schema.TypeSpec{Type: "string"}},
				"name":     {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			WriteOnly: []string{"password"},
		}
		patch, err := CalcPatchWithActualOutputs(
			resource.PropertyMap{
				"password": resource.NewStringProperty("old-secret"),
				"name":     resource.NewStringProperty("old-name"),
			},
			resource.PropertyMap{
				"name": resource.NewStringProperty("old-name"),
			},
			resource.PropertyMap{
				"password": resource.NewStringProperty("old-secret"),
				"name":     resource.NewStringProperty("new-name"),
			},
			spec, nil, nil, "", nil)
		require.NoError(t, err)
		assert.Equal(t, []jsonpatch.JsonPatchOperation{
			{Operation: "replace", Path: "/Name", Value: "new-name"},
			{Operation: "add", Path: "/Password", Value: "old-secret"},
		}, patch)
	})

	t.Run("create-only write-only falls back without resend", func(t *testing.T) {
		spec := metadata.CloudAPIResource{
			Inputs: map[string]schema.PropertySpec{
				"password": {TypeSpec: schema.TypeSpec{Type: "string"}},
				"name":     {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			WriteOnly:  []string{"password"},
			CreateOnly: []string{"password"},
		}
		patch, err := CalcPatchWithActualOutputs(
			resource.PropertyMap{
				"password": resource.NewStringProperty("old-secret"),
				"name":     resource.NewStringProperty("old-name"),
			},
			resource.PropertyMap{
				"name": resource.NewStringProperty("old-name"),
			},
			resource.PropertyMap{
				"password": resource.NewStringProperty("old-secret"),
				"name":     resource.NewStringProperty("new-name"),
			},
			spec, nil, nil, "", nil)
		require.NoError(t, err)
		assert.Equal(t, []jsonpatch.JsonPatchOperation{
			{Operation: "replace", Path: "/Name", Value: "new-name"},
		}, patch)
	})

	t.Run("write-only descendant of create-only parent falls back without resend", func(t *testing.T) {
		spec := metadata.CloudAPIResource{
			Inputs: map[string]schema.PropertySpec{
				"configurationProperties": {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Ref: "#/types/aws-native:test:ConfigurationProperty"},
					},
				},
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Type: "string"},
					},
				},
			},
			Outputs: map[string]schema.PropertySpec{
				"configurationProperties": {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Ref: "#/types/aws-native:test:ConfigurationProperty"},
					},
				},
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Type: "string"},
					},
				},
			},
			WriteOnly:  []string{"configurationProperties/*/type"},
			CreateOnly: []string{"configurationProperties"},
		}
		types := map[string]metadata.CloudAPIType{
			"aws-native:test:ConfigurationProperty": {
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"name": {TypeSpec: schema.TypeSpec{Type: "string"}},
					"type": {TypeSpec: schema.TypeSpec{Type: "string"}},
				},
			},
		}
		oldInputs := resource.PropertyMap{
			"configurationProperties": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"name": resource.NewStringProperty("Owner"),
					"type": resource.NewStringProperty("String"),
				}),
			}),
			"tags": resource.NewObjectProperty(resource.PropertyMap{
				"env": resource.NewStringProperty("old"),
			}),
		}
		actualOutputs := resource.PropertyMap{
			"configurationProperties": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"name": resource.NewStringProperty("Owner"),
				}),
			}),
			"tags": resource.NewObjectProperty(resource.PropertyMap{
				"env": resource.NewStringProperty("old"),
			}),
		}
		newInputs := resource.PropertyMap{
			"configurationProperties": oldInputs["configurationProperties"],
			"tags": resource.NewObjectProperty(resource.PropertyMap{
				"env": resource.NewStringProperty("new"),
			}),
		}

		patch, err := CalcPatchWithActualOutputs(oldInputs, actualOutputs, newInputs, spec, types, nil, "", nil)
		require.NoError(t, err)
		assert.Equal(t, []jsonpatch.JsonPatchOperation{
			{Operation: "replace", Path: "/Tags", Value: map[string]interface{}{"env": "new"}},
		}, patch)
	})

	t.Run("owned map drift is reconciled", func(t *testing.T) {
		spec := metadata.CloudAPIResource{
			Inputs: map[string]schema.PropertySpec{
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Type: "string"},
					},
				},
			},
			Outputs: map[string]schema.PropertySpec{
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Type: "string"},
					},
				},
			},
		}
		patch, err := CalcPatchWithActualOutputs(
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			})},
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
				"extra": resource.NewStringProperty("external"),
			})},
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			})},
			spec, nil, nil, "", nil)
		require.NoError(t, err)
		assert.Equal(t, []jsonpatch.JsonPatchOperation{
			{
				Operation: "replace",
				Path:      "/Tags",
				Value: map[string]interface{}{
					"owner": "team",
				},
			},
		}, patch)
	})

	t.Run("aws managed tag drift is suppressed", func(t *testing.T) {
		spec := metadata.CloudAPIResource{
			TagsProperty: "tags",
			Inputs: map[string]schema.PropertySpec{
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Type: "string"},
					},
				},
			},
			Outputs: map[string]schema.PropertySpec{
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Type: "string"},
					},
				},
			},
		}
		patch, err := CalcPatchWithActualOutputs(
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			})},
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner":       resource.NewStringProperty("team"),
				"aws:managed": resource.NewStringProperty("external"),
			})},
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			})},
			spec, nil, nil, "aws-native:test:Resource", NewTransformCache())
		require.NoError(t, err)
		assert.Empty(t, patch)
	})

	t.Run("WAF WebAcl updates resend canonical rules", func(t *testing.T) {
		const tokenDomainsProperty = "tokenDomains"

		spec := metadata.CloudAPIResource{
			Inputs: map[string]schema.PropertySpec{
				wafv2RulesProperty: {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Type: "string"},
					},
				},
				tokenDomainsProperty: {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Type: "string"},
					},
				},
			},
			Outputs: map[string]schema.PropertySpec{
				wafv2RulesProperty: {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Type: "string"},
					},
				},
				tokenDomainsProperty: {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Type: "string"},
					},
				},
			},
		}
		rules := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewStringProperty("canonical-rule"),
		})
		oldInputs := resource.PropertyMap{
			wafv2RulesProperty: rules,
			tokenDomainsProperty: resource.NewArrayProperty(
				[]resource.PropertyValue{resource.NewStringProperty("old.example")},
			),
		}
		newInputs := resource.PropertyMap{
			wafv2RulesProperty: rules,
			tokenDomainsProperty: resource.NewArrayProperty(
				[]resource.PropertyValue{resource.NewStringProperty("new.example")},
			),
		}

		patch, err := CalcPatchWithActualOutputs(
			oldInputs,
			oldInputs,
			newInputs,
			spec,
			nil,
			nil,
			awsNativeWafv2WebAclToken,
			NewTransformCache(),
		)
		require.NoError(t, err)
		assert.Equal(t, []jsonpatch.JsonPatchOperation{
			{Operation: "replace", Path: "/Rules", Value: []interface{}{"canonical-rule"}},
			{Operation: "replace", Path: "/TokenDomains", Value: []interface{}{"new.example"}},
		}, patch)
	})

	t.Run("key value array tag reorder is suppressed", func(t *testing.T) {
		spec := metadata.CloudAPIResource{
			TagsProperty: "tags",
			TagsStyle:    default_tags.TagsStyleKeyValueArray,
			Inputs: map[string]schema.PropertySpec{
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Ref: "#/types/aws-native:test:Tag"},
					},
				},
			},
			Outputs: map[string]schema.PropertySpec{
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Ref: "#/types/aws-native:test:Tag"},
					},
				},
			},
		}
		types := map[string]metadata.CloudAPIType{
			"aws-native:test:Tag": {
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"key":   {TypeSpec: schema.TypeSpec{Type: "string"}},
					"value": {TypeSpec: schema.TypeSpec{Type: "string"}},
				},
			},
		}
		oldTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Environment"),
				"value": resource.NewStringProperty("prod"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("my-resource"),
			}),
		})
		actualTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("my-resource"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Environment"),
				"value": resource.NewStringProperty("prod"),
			}),
		})
		patch, err := CalcPatchWithActualOutputs(
			resource.PropertyMap{"tags": oldTags},
			resource.PropertyMap{"tags": actualTags},
			resource.PropertyMap{"tags": oldTags},
			spec, types, nil, "aws-native:test:Resource", NewTransformCache())
		require.NoError(t, err)
		assert.Empty(t, patch)
	})

	t.Run("normalized IAM policy document is suppressed", func(t *testing.T) {
		spec := metadata.CloudAPIResource{
			Inputs: map[string]schema.PropertySpec{
				"assumeRolePolicyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
				"policies": {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Ref: "#/types/aws-native:iam:RolePolicy"},
					},
				},
			},
			Outputs: map[string]schema.PropertySpec{
				"assumeRolePolicyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
				"policies": {
					TypeSpec: schema.TypeSpec{
						Type:  "array",
						Items: &schema.TypeSpec{Ref: "#/types/aws-native:iam:RolePolicy"},
					},
				},
			},
		}
		types := map[string]metadata.CloudAPIType{
			"aws-native:iam:RolePolicy": {
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"policyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
					"policyName":     {TypeSpec: schema.TypeSpec{Type: "string"}},
				},
			},
		}
		oldInputs := resource.PropertyMap{
			"assumeRolePolicyDocument": resource.NewStringProperty(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Effect": "Allow",
					"Principal": {"Service": "ec2.amazonaws.com"},
					"Action": "sts:AssumeRole"
				}]
			}`),
			"policies": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"policyName": resource.NewStringProperty("test-policy"),
					"policyDocument": resource.NewObjectProperty(resource.PropertyMap{
						"Version": resource.NewStringProperty("2012-10-17"),
						"Statement": resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewObjectProperty(resource.PropertyMap{
								"Effect":   resource.NewStringProperty("Allow"),
								"Action":   resource.NewStringProperty("*"),
								"Resource": resource.NewStringProperty("*"),
							}),
						}),
					}),
				}),
			}),
		}
		actualOutputs := resource.PropertyMap{
			"assumeRolePolicyDocument": resource.NewObjectProperty(resource.PropertyMap{
				"Version": resource.NewStringProperty("2012-10-17"),
				"Statement": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewObjectProperty(resource.PropertyMap{
						"Effect":    resource.NewStringProperty("Allow"),
						"Principal": resource.NewObjectProperty(resource.PropertyMap{"Service": resource.NewStringProperty("ec2.amazonaws.com")}),
						"Action":    resource.NewStringProperty("sts:AssumeRole"),
					}),
				}),
			}),
			"policies": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"policyName": resource.NewStringProperty("test-policy"),
					"policyDocument": resource.NewObjectProperty(resource.PropertyMap{
						"Version": resource.NewStringProperty("2012-10-17"),
						"Statement": resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewObjectProperty(resource.PropertyMap{
								"Effect":   resource.NewStringProperty("Allow"),
								"Action":   resource.NewArrayProperty([]resource.PropertyValue{resource.NewStringProperty("*")}),
								"Resource": resource.NewStringProperty("*"),
							}),
						}),
					}),
				}),
			}),
		}
		patch, err := CalcPatchWithActualOutputs(
			oldInputs, actualOutputs, oldInputs, spec, types, nil, "aws-native:iam:Role", NewTransformCache())
		require.NoError(t, err)
		assert.Empty(t, patch)
	})

}

// Map applies the given function to each element of the given slice and returns a new slice with the results.
func Map[T, U any](s []T, f func(T) U) []U {
	r := Prealloc[U](len(s))
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

// Preallocates a slice of type T of a given length. If the input length is 0 then the returned slice will be nil.
func Prealloc[T any](capacity int) []T {
	if capacity == 0 {
		return nil
	}
	return make([]T, 0, capacity)
}
