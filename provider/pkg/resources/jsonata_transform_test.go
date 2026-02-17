// Copyright 2024, Pulumi Corporation.

package resources

import (
	"testing"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-go-provider/resourcex"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPathMatcher(t *testing.T) {
	t.Run("exact match without wildcard", func(t *testing.T) {
		pm := NewPathMatcher("engine")
		assert.True(t, pm.Match("engine"))
		assert.False(t, pm.Match("Engine"))
		assert.False(t, pm.Match("other"))
	})

	t.Run("nested path without wildcard", func(t *testing.T) {
		pm := NewPathMatcher("fileSystemProtection/replicationOverwriteProtection")
		assert.True(t, pm.Match("fileSystemProtection/replicationOverwriteProtection"))
		assert.False(t, pm.Match("fileSystemProtection/other"))
	})

	t.Run("array wildcard matches numeric index", func(t *testing.T) {
		pm := NewPathMatcher("securityGroupEgress/*/ipProtocol")
		assert.True(t, pm.Match("securityGroupEgress/0/ipProtocol"))
		assert.True(t, pm.Match("securityGroupEgress/1/ipProtocol"))
		assert.True(t, pm.Match("securityGroupEgress/99/ipProtocol"))
		assert.False(t, pm.Match("securityGroupEgress/ipProtocol"))
		assert.False(t, pm.Match("securityGroupIngress/0/ipProtocol"))
	})

	t.Run("multiple wildcards", func(t *testing.T) {
		pm := NewPathMatcher("a/*/b/*/c")
		assert.True(t, pm.Match("a/0/b/1/c"))
		assert.True(t, pm.Match("a/5/b/10/c"))
		assert.False(t, pm.Match("a/0/b/c"))
	})

	t.Run("hasWildcard", func(t *testing.T) {
		assert.True(t, NewPathMatcher("a/*/b").HasWildcard())
		assert.False(t, NewPathMatcher("a/b/c").HasWildcard())
	})

	t.Run("getPropertyName", func(t *testing.T) {
		assert.Equal(t, "ipProtocol", NewPathMatcher("securityGroupEgress/*/ipProtocol").GetPropertyName())
		assert.Equal(t, "engine", NewPathMatcher("engine").GetPropertyName())
	})
}

func TestTransformCache(t *testing.T) {
	cache := &TransformCache{
		transforms: make(map[string][]CompiledTransform),
	}

	t.Run("returns nil for nil spec", func(t *testing.T) {
		result := cache.GetOrCompile("test", nil)
		assert.Nil(t, result)
	})

	t.Run("returns nil for empty transforms", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{
			PropertyTransforms: nil,
		}
		result := cache.GetOrCompile("test", spec)
		assert.Nil(t, result)
	})

	t.Run("compiles and caches valid transforms", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{
			PropertyTransforms: map[string]string{
				"engine": "$lowercase(Engine)",
			},
		}

		// First call should compile
		result1 := cache.GetOrCompile("aws-native:rds:DBCluster", spec)
		require.NotNil(t, result1)
		assert.Len(t, result1, 1)
		assert.Equal(t, "engine", result1[0].Path)
		assert.Equal(t, "$lowercase(Engine)", result1[0].RawExpr)
		assert.NotNil(t, result1[0].Expression)
		assert.NotNil(t, result1[0].PathMatcher)

		// Second call should return cached
		result2 := cache.GetOrCompile("aws-native:rds:DBCluster", spec)
		assert.Equal(t, result1, result2)
	})

	t.Run("skips invalid transforms", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{
			PropertyTransforms: map[string]string{
				"valid":   "$lowercase(Valid)",
				"invalid": "this is not valid jsonata $$$%%%",
			},
		}

		result := cache.GetOrCompile("test-invalid", spec)
		// Should only have the valid one
		assert.Len(t, result, 1)
		assert.Equal(t, "valid", result[0].Path)
	})
}

func TestEvaluateTransform(t *testing.T) {
	t.Run("lowercase transform", func(t *testing.T) {
		transforms := compileTransforms(map[string]string{
			"engine": "$lowercase(Engine)",
		})
		require.Len(t, transforms, 1)

		// Test with uppercase input
		result, err := EvaluateTransform(transforms[0], "MYSQL", map[string]interface{}{
			"Engine": "MYSQL",
		})
		require.NoError(t, err)
		assert.Equal(t, "mysql", result)
	})

	t.Run("returns original on undefined", func(t *testing.T) {
		transforms := compileTransforms(map[string]string{
			"test": "$nonexistent.field",
		})
		require.Len(t, transforms, 1)

		result, err := EvaluateTransform(transforms[0], "original", map[string]interface{}{})
		require.NoError(t, err)
		assert.Equal(t, "original", result)
	})
}

func TestValuesEquivalent(t *testing.T) {
	t.Run("exact string match", func(t *testing.T) {
		assert.True(t, ValuesEquivalent("tcp", "tcp"))
		assert.False(t, ValuesEquivalent("tcp", "udp"))
	})

	t.Run("deep-equals typed comparisons", func(t *testing.T) {
		assert.True(t, ValuesEquivalent(6, 6))
		assert.True(t, ValuesEquivalent("6", "6"))
		assert.True(t, ValuesEquivalent(int(6), float64(6.0)))
	})

	t.Run("regex pattern match", func(t *testing.T) {
		// AWS KMS key ID pattern
		pattern := "arn:.+?:kms:.+?:.+?:key/12345"
		value := "arn:aws:kms:us-east-1:123456789012:key/12345"
		assert.True(t, ValuesEquivalent(value, pattern))
	})

	t.Run("simple regex with .*", func(t *testing.T) {
		assert.True(t, ValuesEquivalent("hello world", "hello.*"))
		assert.False(t, ValuesEquivalent("goodbye", "hello.*"))
	})

	t.Run("quoted regex pattern", func(t *testing.T) {
		pattern := "\"^arn:aws[a-zA-Z-]*:iam::123456789012:[a-zA-Z-]*\""
		assert.True(t, ValuesEquivalent("arn:aws:iam::123456789012:root", pattern))
	})

	// Real CloudFormation propertyTransform patterns from AWS resource schemas.
	// These demonstrate the regex matching capability.
	t.Run("CloudFormation engineVersion transform pattern", func(t *testing.T) {
		// Transform: $join([$string(EngineVersion), ".*"])
		// This allows "5.8" to match "5.8.1", "5.8.2", etc.
		assert.True(t, ValuesEquivalent("5.8.1", "5.8.*"))
		assert.True(t, ValuesEquivalent("5.8.23", "5.8.*"))
		assert.False(t, ValuesEquivalent("5.9.1", "5.8.*"))
	})

	t.Run("CloudFormation principal ARN transform pattern", func(t *testing.T) {
		// Transform: $join(["^arn:aws[a-zA-Z-]*:iam::",Principal,":[a-zA-Z-]*"])
		// This matches ARNs with optional partitions (aws, aws-cn, aws-us-gov)
		pattern := "^arn:aws[a-zA-Z-]*:iam::123456789012:[a-zA-Z-]*"
		assert.True(t, ValuesEquivalent("arn:aws:iam::123456789012:root", pattern))
		assert.True(t, ValuesEquivalent("arn:aws-cn:iam::123456789012:user", pattern))
		assert.True(t, ValuesEquivalent("arn:aws-us-gov:iam::123456789012:role", pattern))
		assert.False(t, ValuesEquivalent("arn:aws:iam::999999999999:root", pattern))
	})

	t.Run("CloudFormation KMS ARN transform pattern", func(t *testing.T) {
		// Transform from metadata:
		// $join(["arn:(aws)...[:]{1}key\\/", KmsKeyId])
		pattern := "arn:(aws)[-]{0,1}[a-z]{0,2}[-]{0,1}[a-z]{0,3}:kms:[a-z]{2}[-]{1}[a-z]{3,10}[-]{0,1}[a-z]{0,10}[-]{1}[1-3]{1}:[0-9]{12}[:]{1}key\\/12345"
		assert.True(t, ValuesEquivalent("arn:aws:kms:us-east-1:123456789012:key/12345", pattern))
		assert.False(t, ValuesEquivalent("arn:aws:kms:us-east-1:123456789012:key/98765", pattern))
	})

	t.Run("CloudFormation Lambda function ARN alternation pattern", func(t *testing.T) {
		// Transform from metadata:
		// $join(["((arn:.*:lambda:.*:[0-9]{12}:function)|([0-9]{12}:function)):", Name])
		pattern := "((arn:.*:lambda:.*:[0-9]{12}:function)|([0-9]{12}:function)):my-func"
		assert.True(t, ValuesEquivalent("arn:aws:lambda:us-east-1:123456789012:function:my-func", pattern))
		assert.True(t, ValuesEquivalent("123456789012:function:my-func", pattern))
		assert.False(t, ValuesEquivalent("arn:aws:lambda:us-east-1:123456789012:function:other-func", pattern))
	})

	t.Run("does not coerce cross-type literals", func(t *testing.T) {
		assert.False(t, ValuesEquivalent("1", 1))
		assert.False(t, ValuesEquivalent("6", float64(6.0)))
		assert.False(t, ValuesEquivalent(true, "true"))
		assert.False(t, ValuesEquivalent(false, "false"))
	})

	// Numeric edge cases - important for AWS responses which may return
	// numbers as different Go types depending on JSON parsing context.
	t.Run("numeric edge cases", func(t *testing.T) {
		tests := []struct {
			name     string
			a, b     interface{}
			expected bool
		}{
			// Same type comparisons
			{"int vs int equal", int(6), int(6), true},
			{"int vs int different", int(6), int(7), false},
			{"float64 vs float64 equal", float64(6.0), float64(6.0), true},
			{"float64 vs float64 different", float64(6.0), float64(7.0), false},
			{"float64 decimal", float64(6.5), float64(6.5), true},

				// Cross-type numeric comparisons (common when comparing decoded AWS responses)
				// are handled by DeepEquals via Pulumi PropertyValue normalization.
				{"int vs float64 equal", int(6), float64(6.0), true},
				{"float64 vs int equal", float64(6.0), int(6), true},

			// Cross-type number/string comparisons are intentionally strict.
			{"string number vs float64", "6", float64(6.0), false},
			{"string number vs int", "6", int(6), false},
			{"float64 vs string number", float64(6.0), "6", false},

			// Different values should not match
			{"string vs different number", "6", float64(7.0), false},
			{"different string numbers", "6", "7", false},

			// Non-numeric strings
			{"non-numeric string equal", "tcp", "tcp", true},
			{"non-numeric string different", "tcp", "udp", false},

			// Boolean comparisons
			{"bool true", true, true, true},
			{"bool false", false, false, true},
			{"bool different", true, false, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := ValuesEquivalent(tt.a, tt.b)
				assert.Equal(t, tt.expected, result, "ValuesEquivalent(%v, %v)", tt.a, tt.b)
			})
		}
	})
}

func TestFindTransformForPath(t *testing.T) {
	transforms := compileTransforms(map[string]string{
		"engine":                           "$lowercase(Engine)",
		"securityGroupEgress/*/ipProtocol": "$lowercase(IpProtocol)",
		"fileSystemProtection/replicationOverwriteProtection": "$uppercase(ReplicationOverwriteProtection)",
	})

	t.Run("finds exact match", func(t *testing.T) {
		transform := FindTransformForPath(transforms, "engine")
		require.NotNil(t, transform)
		assert.Equal(t, "engine", transform.Path)
	})

	t.Run("finds wildcard match", func(t *testing.T) {
		transform := FindTransformForPath(transforms, "securityGroupEgress/0/ipProtocol")
		require.NotNil(t, transform)
		assert.Equal(t, "securityGroupEgress/*/ipProtocol", transform.Path)
	})

	t.Run("finds nested match", func(t *testing.T) {
		transform := FindTransformForPath(transforms, "fileSystemProtection/replicationOverwriteProtection")
		require.NotNil(t, transform)
	})

	t.Run("returns nil for no match", func(t *testing.T) {
		transform := FindTransformForPath(transforms, "unknownProperty")
		assert.Nil(t, transform)
	})
}

func TestExtractPropertyContext(t *testing.T) {
	props := map[string]interface{}{
		"name": "test",
		"securityGroupEgress": []interface{}{
			map[string]interface{}{
				"ipProtocol": "tcp",
				"fromPort":   float64(80),
				"toPort":     float64(80),
			},
			map[string]interface{}{
				"ipProtocol": "udp",
				"fromPort":   float64(53),
				"toPort":     float64(53),
			},
		},
		"nested": map[string]interface{}{
			"child": map[string]interface{}{
				"value": "deep",
			},
		},
	}

	t.Run("top-level property returns full props", func(t *testing.T) {
		ctx := ExtractPropertyContext(props, "name")
		assert.Equal(t, props, ctx)
	})

	t.Run("array element returns parent element", func(t *testing.T) {
		ctx := ExtractPropertyContext(props, "securityGroupEgress/0/ipProtocol")
		require.NotNil(t, ctx)
		assert.Equal(t, "tcp", ctx["ipProtocol"])
		assert.Equal(t, float64(80), ctx["fromPort"])
	})

	t.Run("nested object returns parent object", func(t *testing.T) {
		ctx := ExtractPropertyContext(props, "nested/child/value")
		require.NotNil(t, ctx)
		assert.Equal(t, "deep", ctx["value"])
	})

	t.Run("returns nil for invalid path", func(t *testing.T) {
		ctx := ExtractPropertyContext(props, "nonexistent/path/value")
		assert.Nil(t, ctx)
	})
}

func TestSuppressPropertyTransformDiffs(t *testing.T) {
	t.Run("suppresses lowercase equivalent diff", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{
			PropertyTransforms: map[string]string{
				"engine": "$lowercase(Engine)",
			},
		}

		diff := &resource.ObjectDiff{
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"engine": {
					Old: resource.NewStringProperty("MYSQL"),
					New: resource.NewStringProperty("mysql"),
				},
			},
		}

		originalInputs := resource.PropertyMap{
			"engine": resource.NewStringProperty("MYSQL"),
		}

		result := suppressPropertyTransformDiffs("aws-native:rds:DBCluster", spec, diff, originalInputs, NewTransformCache())
		assert.Empty(t, result.Updates, "Diff should be suppressed for case-insensitive match")
	})

	t.Run("preserves real diff", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{
			PropertyTransforms: map[string]string{
				"engine": "$lowercase(Engine)",
			},
		}

		diff := &resource.ObjectDiff{
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"engine": {
					Old: resource.NewStringProperty("mysql"),
					New: resource.NewStringProperty("postgresql"),
				},
			},
		}

		originalInputs := resource.PropertyMap{
			"engine": resource.NewStringProperty("mysql"),
		}

		result := suppressPropertyTransformDiffs("aws-native:rds:DBCluster", spec, diff, originalInputs, NewTransformCache())
		assert.Len(t, result.Updates, 1, "Real diff should be preserved")
	})

	t.Run("handles nested object diff", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{
			PropertyTransforms: map[string]string{
				"fileSystemProtection/replicationOverwriteProtection": "$uppercase(ReplicationOverwriteProtection)='DISABLED' ? 'REPLICATING' : $uppercase(ReplicationOverwriteProtection)",
			},
		}

		oldProps := resource.PropertyMap{
			"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
				"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
			}),
		}
		newProps := resource.PropertyMap{
			"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
				"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
			}),
		}
		diff := oldProps.Diff(newProps)
		require.NotNil(t, diff, "diff should not be nil")

		result := suppressPropertyTransformDiffs("aws-native:efs:FileSystem", spec, diff, oldProps, NewTransformCache())
		// This specific transform should suppress DISABLED -> REPLICATING
		// The transform evaluates to REPLICATING when input is DISABLED
		assert.Empty(t, result.Updates, "EFS replication transition should be suppressed")
	})
}

// TestDecodeValueForJSONata verifies that resourcex.DecodeValue correctly converts
// PropertyValues to Go interfaces for JSONata evaluation. This is critical because
// JSONata transforms need plain Go values, not Pulumi's PropertyValue types.
// The key behaviors we rely on are: secret unwrapping, nested structure conversion,
// and consistent type representation for comparison.
func TestDecodeValueForJSONata(t *testing.T) {
	t.Run("converts basic types", func(t *testing.T) {
		assert.Equal(t, true, resourcex.DecodeValue(resource.NewBoolProperty(true)))
		assert.Equal(t, float64(42), resourcex.DecodeValue(resource.NewNumberProperty(42)))
		assert.Equal(t, "hello", resourcex.DecodeValue(resource.NewStringProperty("hello")))
		assert.Nil(t, resourcex.DecodeValue(resource.NewNullProperty()))
	})

	t.Run("converts arrays", func(t *testing.T) {
		arr := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewStringProperty("a"),
			resource.NewStringProperty("b"),
		})
		result := resourcex.DecodeValue(arr)
		assert.Equal(t, []interface{}{"a", "b"}, result)
	})

	t.Run("converts objects", func(t *testing.T) {
		obj := resource.NewObjectProperty(resource.PropertyMap{
			"key": resource.NewStringProperty("value"),
		})
		result := resourcex.DecodeValue(obj)
		expected := map[string]interface{}{"key": "value"}
		assert.Equal(t, expected, result)
	})

	t.Run("converts nested structures", func(t *testing.T) {
		nested := resource.NewObjectProperty(resource.PropertyMap{
			"items": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"name": resource.NewStringProperty("item1"),
				}),
			}),
		})
		result := resourcex.DecodeValue(nested)
		expected := map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{"name": "item1"},
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("handles secret values", func(t *testing.T) {
		secret := resource.MakeSecret(resource.NewStringProperty("secret-value"))
		result := resourcex.DecodeValue(secret)
		assert.Equal(t, "secret-value", result)
	})

	t.Run("handles secret with nested object", func(t *testing.T) {
		secret := resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{
			"username": resource.NewStringProperty("admin"),
			"password": resource.NewStringProperty("hunter2"),
		}))
		result := resourcex.DecodeValue(secret)
		expected := map[string]interface{}{
			"username": "admin",
			"password": "hunter2",
		}
		assert.Equal(t, expected, result)
	})
}

func compileTransforms(specs map[string]string) []CompiledTransform {
	cache := &TransformCache{
		transforms: make(map[string][]CompiledTransform),
	}
	spec := &metadata.CloudAPIResource{
		PropertyTransforms: specs,
	}
	return cache.GetOrCompile("test", spec)
}
