package resources

import (
	"testing"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPathClassifierProjectWritableOutputState(t *testing.T) {
	spec := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"code": {TypeSpec: pschema.TypeSpec{Ref: "#/types/aws-native:lambda:FunctionCode"}},
			"arn":  {TypeSpec: pschema.TypeSpec{Type: "string"}},
		},
		Outputs: map[string]pschema.PropertySpec{
			"code": {TypeSpec: pschema.TypeSpec{Ref: "#/types/aws-native:lambda:FunctionCode"}},
			"arn":  {TypeSpec: pschema.TypeSpec{Type: "string"}},
		},
		ReadOnly:  []string{"arn", "code/resolvedImageUri"},
		WriteOnly: []string{"code/imageUri"},
	}
	types := map[string]metadata.CloudAPIType{
		"aws-native:lambda:FunctionCode": {
			Type:     "object",
			Required: []string{"s3Bucket"},
			Properties: map[string]pschema.PropertySpec{
				"imageUri":         {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"resolvedImageUri": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"s3Bucket":         {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
	}
	classifier := NewPathClassifier(&spec, types)

	projected := classifier.ProjectWritableOutputState(resource.PropertyMap{
		"arn": resource.NewStringProperty("arn"),
		"code": resource.NewObjectProperty(resource.PropertyMap{
			"imageUri":         resource.NewStringProperty("secret"),
			"resolvedImageUri": resource.NewStringProperty("sha"),
			"s3Bucket":         resource.NewStringProperty("bucket"),
		}),
	})

	require.True(t, projected.HasValue("code"))
	code := projected["code"].ObjectValue()
	assert.Equal(t, resource.NewStringProperty("bucket"), code["s3Bucket"])
	assert.False(t, code.HasValue("imageUri"))
	assert.False(t, code.HasValue("resolvedImageUri"))
	assert.False(t, projected.HasValue("arn"))

	info, ok := classifier.Classify("code/s3Bucket")
	require.True(t, ok)
	assert.Equal(t, ConcreteField, info.Kind)
	assert.True(t, info.Required)
}

func TestPathClassifierActualInputBaselineOwnership(t *testing.T) {
	spec := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"backupTarget": {TypeSpec: pschema.TypeSpec{Type: "string"}},
			"settings":     {TypeSpec: pschema.TypeSpec{Ref: "#/types/aws-native:test:Settings"}},
			"tags": {
				TypeSpec: pschema.TypeSpec{
					Type:                 "object",
					AdditionalProperties: &pschema.TypeSpec{Type: "string"},
				},
			},
		},
		Outputs: map[string]pschema.PropertySpec{
			"backupTarget": {TypeSpec: pschema.TypeSpec{Type: "string"}},
			"settings":     {TypeSpec: pschema.TypeSpec{Ref: "#/types/aws-native:test:Settings"}},
			"tags": {
				TypeSpec: pschema.TypeSpec{
					Type:                 "object",
					AdditionalProperties: &pschema.TypeSpec{Type: "string"},
				},
			},
		},
	}
	types := map[string]metadata.CloudAPIType{
		"aws-native:test:Settings": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"defaulted": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"managed":   {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
	}
	classifier := NewPathClassifier(&spec, types)
	actual := resource.PropertyMap{
		"backupTarget": resource.NewStringProperty("region"),
		"settings": resource.NewObjectProperty(resource.PropertyMap{
			"defaulted": resource.NewStringProperty("aws"),
			"managed":   resource.NewStringProperty("drifted"),
		}),
		"tags": resource.NewObjectProperty(resource.PropertyMap{
			"owner": resource.NewStringProperty("team"),
			"extra": resource.NewStringProperty("external"),
		}),
	}
	oldDesired := resource.PropertyMap{
		"settings": resource.NewObjectProperty(resource.PropertyMap{
			"managed": resource.NewStringProperty("old"),
		}),
		"tags": resource.NewObjectProperty(resource.PropertyMap{
			"owner": resource.NewStringProperty("team"),
		}),
	}
	newDesired := resource.PropertyMap{
		"settings": resource.NewObjectProperty(resource.PropertyMap{
			"managed": resource.NewStringProperty("old"),
		}),
		"tags": resource.NewObjectProperty(resource.PropertyMap{
			"owner": resource.NewStringProperty("team"),
		}),
	}

	baseline := classifier.actualInputBaseline(oldDesired, actual, newDesired)

	assert.False(t, baseline.HasValue("backupTarget"))
	settings := baseline["settings"].ObjectValue()
	assert.False(t, settings.HasValue("defaulted"))
	assert.Equal(t, resource.NewStringProperty("drifted"), settings["managed"])
	tags := baseline["tags"].ObjectValue()
	assert.Equal(t, resource.NewStringProperty("team"), tags["owner"])
	assert.Equal(t, resource.NewStringProperty("external"), tags["extra"])
}

func TestPathClassifierActualInputBaselinePreservesSecretWrappers(t *testing.T) {
	spec := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"settings": {TypeSpec: pschema.TypeSpec{Ref: "#/types/aws-native:test:Settings"}},
			"tags": {
				TypeSpec: pschema.TypeSpec{
					Type:                 "object",
					AdditionalProperties: &pschema.TypeSpec{Type: "string"},
				},
			},
		},
		Outputs: map[string]pschema.PropertySpec{
			"settings": {TypeSpec: pschema.TypeSpec{Ref: "#/types/aws-native:test:Settings"}},
			"tags": {
				TypeSpec: pschema.TypeSpec{
					Type:                 "object",
					AdditionalProperties: &pschema.TypeSpec{Type: "string"},
				},
			},
		},
	}
	types := map[string]metadata.CloudAPIType{
		"aws-native:test:Settings": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"password": {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
	}
	classifier := NewPathClassifier(&spec, types)

	baseline := classifier.actualInputBaseline(
		resource.PropertyMap{
			"settings": resource.NewObjectProperty(resource.PropertyMap{
				"password": resource.MakeSecret(resource.NewStringProperty("old")),
			}),
			"tags": resource.NewObjectProperty(resource.PropertyMap{
				"secret": resource.MakeSecret(resource.NewStringProperty("old")),
			}),
		},
		resource.PropertyMap{
			"settings": resource.NewObjectProperty(resource.PropertyMap{
				"password": resource.NewStringProperty("drifted"),
			}),
			"tags": resource.NewObjectProperty(resource.PropertyMap{
				"secret": resource.NewStringProperty("drifted"),
				"plain":  resource.NewStringProperty("value"),
			}),
		},
		resource.PropertyMap{
			"settings": resource.NewObjectProperty(resource.PropertyMap{
				"password": resource.MakeSecret(resource.NewStringProperty("old")),
			}),
			"tags": resource.NewObjectProperty(resource.PropertyMap{
				"secret": resource.MakeSecret(resource.NewStringProperty("old")),
			}),
		})

	password := baseline["settings"].ObjectValue()["password"]
	require.True(t, password.IsSecret())
	assert.Equal(t, resource.NewStringProperty("drifted"), password.SecretValue().Element)
	secretTag := baseline["tags"].ObjectValue()["secret"]
	require.True(t, secretTag.IsSecret())
	assert.Equal(t, resource.NewStringProperty("drifted"), secretTag.SecretValue().Element)
	assert.False(t, baseline["tags"].ObjectValue()["plain"].IsSecret())
}

func TestPathClassifierArrayOwnership(t *testing.T) {
	spec := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"rules": {
				TypeSpec: pschema.TypeSpec{
					Type:  "array",
					Items: &pschema.TypeSpec{Ref: "#/types/aws-native:test:Rule"},
				},
			},
		},
		Outputs: map[string]pschema.PropertySpec{
			"rules": {
				TypeSpec: pschema.TypeSpec{
					Type:  "array",
					Items: &pschema.TypeSpec{Ref: "#/types/aws-native:test:Rule"},
				},
			},
		},
	}
	types := map[string]metadata.CloudAPIType{
		"aws-native:test:Rule": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"defaulted": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"managed":   {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
	}
	classifier := NewPathClassifier(&spec, types)

	info, ok := classifier.Classify("rules/0/managed")
	require.True(t, ok)
	assert.Equal(t, ConcreteField, info.Kind)

	baseline := classifier.actualInputBaseline(
		resource.PropertyMap{
			"rules": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"managed": resource.NewStringProperty("old"),
				}),
			}),
		},
		resource.PropertyMap{
			"rules": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"defaulted": resource.NewStringProperty("aws"),
					"managed":   resource.NewStringProperty("drifted"),
				}),
			}),
		},
		resource.PropertyMap{
			"rules": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"managed": resource.NewStringProperty("old"),
				}),
			}),
		})

	rules := baseline["rules"].ArrayValue()
	rule := rules[0].ObjectValue()
	assert.False(t, rule.HasValue("defaulted"))
	assert.Equal(t, resource.NewStringProperty("drifted"), rule["managed"])
}

func TestPathClassifierNestedRequiredHandlesRecursiveTypesAtArbitraryDepth(t *testing.T) {
	spec := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"children": {
				TypeSpec: pschema.TypeSpec{
					Type:  "array",
					Items: &pschema.TypeSpec{Ref: "#/types/aws-native:test:Child"},
				},
			},
		},
	}
	types := map[string]metadata.CloudAPIType{
		"aws-native:test:Child": {
			Type:     "object",
			Required: []string{"name"},
			Properties: map[string]pschema.PropertySpec{
				"name": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"children": {
					TypeSpec: pschema.TypeSpec{
						Type:  "array",
						Items: &pschema.TypeSpec{Ref: "#/types/aws-native:test:Child"},
					},
				},
			},
		},
	}

	classifier := NewPathClassifier(&spec, types)

	info, ok := classifier.Classify("children/0/name")
	require.True(t, ok)
	assert.True(t, info.Required)

	info, ok = classifier.Classify("children/0/children/0/children/0/name")
	require.True(t, ok)
	assert.True(t, info.Required)
}

func TestPathClassifierNestedRequiredHandlesRealRecursiveType(t *testing.T) {
	spec := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"children": {
				TypeSpec: pschema.TypeSpec{
					Type:  "array",
					Items: &pschema.TypeSpec{Ref: "#/types/aws-native:amplifyuibuilder:ComponentChild"},
				},
			},
		},
	}
	types := map[string]metadata.CloudAPIType{
		"aws-native:amplifyuibuilder:ComponentChild": {
			Type:     "object",
			Required: []string{"componentType", "name", "properties"},
			Properties: map[string]pschema.PropertySpec{
				"children": {
					TypeSpec: pschema.TypeSpec{
						Type:  "array",
						Items: &pschema.TypeSpec{Ref: "#/types/aws-native:amplifyuibuilder:ComponentChild"},
					},
				},
				"componentType": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"name":          {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"properties":    {TypeSpec: pschema.TypeSpec{Type: "object"}},
			},
		},
	}

	classifier := NewPathClassifier(&spec, types)

	info, ok := classifier.Classify("children/0/children/0/children/0/name")
	require.True(t, ok)
	assert.True(t, info.Required)
	assert.Equal(t, ConcreteField, info.Kind)
}

func TestPathHelpersNestedReadWriteDelete(t *testing.T) {
	m := resource.PropertyMap{}
	SetPath(m, "code/imageUri", resource.NewStringProperty("image"))
	v, ok := GetPath(m, "code/imageUri")
	require.True(t, ok)
	assert.Equal(t, "image", v.StringValue())

	DeletePath(m, "code/imageUri")
	_, ok = GetPath(m, "code/imageUri")
	assert.False(t, ok)
}

func TestPathHelpersSetPathWithShapeUsesArrayGuide(t *testing.T) {
	shape := resource.PropertyMap{
		"defaultActions": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"authenticateOidcConfig": resource.NewObjectProperty(resource.PropertyMap{
					"clientSecret": resource.NewStringProperty("old-secret"),
				}),
			}),
		}),
	}
	m := resource.PropertyMap{}
	SetPathWithShape(m, shape, "defaultActions/0/authenticateOidcConfig/clientSecret", resource.NewStringProperty("secret"))

	require.True(t, m["defaultActions"].IsArray())
	secret, ok := GetPath(m, "defaultActions/0/authenticateOidcConfig/clientSecret")
	require.True(t, ok)
	assert.Equal(t, "secret", secret.StringValue())
}

func TestExpandMatchingPaths(t *testing.T) {
	inputs := resource.PropertyMap{
		"defaultActions": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"authenticateOidcConfig": resource.NewObjectProperty(resource.PropertyMap{
					"clientSecret": resource.NewStringProperty("secret-0"),
				}),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"authenticateOidcConfig": resource.NewObjectProperty(resource.PropertyMap{
					"clientSecret": resource.NewStringProperty("secret-1"),
				}),
			}),
		}),
	}

	assert.Equal(t, []string{
		"defaultActions/0/authenticateOidcConfig/clientSecret",
		"defaultActions/1/authenticateOidcConfig/clientSecret",
	}, ExpandMatchingPaths(inputs, "defaultActions/*/authenticateOidcConfig/clientSecret"))
}

func TestPathClassifierWriteOnlyWildcardFallback(t *testing.T) {
	spec := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"defaultActions": {
				TypeSpec: pschema.TypeSpec{
					Type:  "array",
					Items: &pschema.TypeSpec{Ref: "#/types/aws-native:test:Action"},
				},
			},
		},
		Outputs: map[string]pschema.PropertySpec{
			"defaultActions": {
				TypeSpec: pschema.TypeSpec{
					Type:  "array",
					Items: &pschema.TypeSpec{Ref: "#/types/aws-native:test:Action"},
				},
			},
		},
		WriteOnly: []string{"defaultActions/*/authenticateOidcConfig/clientSecret"},
	}
	types := map[string]metadata.CloudAPIType{
		"aws-native:test:Action": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"authenticateOidcConfig": {TypeSpec: pschema.TypeSpec{Ref: "#/types/aws-native:test:AuthenticateOidcConfig"}},
			},
		},
		"aws-native:test:AuthenticateOidcConfig": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"clientId":     {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"clientSecret": {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
	}
	classifier := NewPathClassifier(&spec, types)

	oldDesired := resource.PropertyMap{
		"defaultActions": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"authenticateOidcConfig": resource.NewObjectProperty(resource.PropertyMap{
					"clientId":     resource.NewStringProperty("client"),
					"clientSecret": resource.NewStringProperty("secret"),
				}),
			}),
		}),
	}
	actual := resource.PropertyMap{
		"defaultActions": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"authenticateOidcConfig": resource.NewObjectProperty(resource.PropertyMap{
					"clientId": resource.NewStringProperty("client"),
				}),
			}),
		}),
	}

	baseline := classifier.actualInputBaseline(oldDesired, actual, oldDesired)
	require.True(t, baseline["defaultActions"].IsArray())
	secret, ok := GetPath(baseline, "defaultActions/0/authenticateOidcConfig/clientSecret")
	require.True(t, ok)
	assert.Equal(t, "secret", secret.StringValue())
}

func TestPathClassifierWriteOnlyOutputFallbackPreservesCreateOnlyAndArrayShape(t *testing.T) {
	spec := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"defaultActions": {
				TypeSpec: pschema.TypeSpec{
					Type:  "array",
					Items: &pschema.TypeSpec{Ref: "#/types/aws-native:test:Action"},
				},
			},
		},
		Outputs: map[string]pschema.PropertySpec{
			"defaultActions": {
				TypeSpec: pschema.TypeSpec{
					Type:  "array",
					Items: &pschema.TypeSpec{Ref: "#/types/aws-native:test:Action"},
				},
			},
		},
		WriteOnly:  []string{"defaultActions/*/authenticateOidcConfig/clientSecret"},
		CreateOnly: []string{"defaultActions/*/authenticateOidcConfig/clientSecret"},
	}
	types := map[string]metadata.CloudAPIType{
		"aws-native:test:Action": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"authenticateOidcConfig": {TypeSpec: pschema.TypeSpec{Ref: "#/types/aws-native:test:AuthenticateOidcConfig"}},
			},
		},
		"aws-native:test:AuthenticateOidcConfig": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"clientSecret": {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
	}
	classifier := NewPathClassifier(&spec, types)

	oldDesired := resource.PropertyMap{
		"defaultActions": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"authenticateOidcConfig": resource.NewObjectProperty(resource.PropertyMap{
					"clientSecret": resource.NewStringProperty("secret"),
				}),
			}),
		}),
	}
	outputs := resource.PropertyMap{}
	classifier.AddWriteOnlyOutputFallbacks(outputs, oldDesired)

	require.True(t, outputs["defaultActions"].IsArray())
	secret, ok := GetPath(outputs, "defaultActions/0/authenticateOidcConfig/clientSecret")
	require.True(t, ok)
	assert.Equal(t, "secret", secret.StringValue())
}
