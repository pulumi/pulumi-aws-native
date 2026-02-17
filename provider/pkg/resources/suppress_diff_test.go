// Copyright 2024, Pulumi Corporation.

package resources

import (
	"testing"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterAWSPrefixedTags(t *testing.T) {
	t.Run("filters aws: prefixed tags in string map not in original", func(t *testing.T) {
		newTags := resource.NewObjectProperty(resource.PropertyMap{
			"Name":            resource.NewStringProperty("my-resource"),
			"aws:managed:tag": resource.NewStringProperty("value"),
			"Env":             resource.NewStringProperty("prod"),
		})
		originalTags := resource.NewObjectProperty(resource.PropertyMap{
			"Name": resource.NewStringProperty("my-resource"),
		})

		result := filterAWSPrefixedTags(newTags, originalTags)

		assert.True(t, result.IsObject())
		obj := result.ObjectValue()
		assert.Contains(t, obj, resource.PropertyKey("Name"))
		assert.Contains(t, obj, resource.PropertyKey("Env"))
		assert.NotContains(t, obj, resource.PropertyKey("aws:managed:tag"))
	})

	t.Run("preserves aws: prefixed tags in string map that were in original", func(t *testing.T) {
		newTags := resource.NewObjectProperty(resource.PropertyMap{
			"aws:custom:tag": resource.NewStringProperty("value"),
		})
		originalTags := resource.NewObjectProperty(resource.PropertyMap{
			"aws:custom:tag": resource.NewStringProperty("value"),
		})

		result := filterAWSPrefixedTags(newTags, originalTags)

		assert.True(t, result.IsObject())
		assert.Contains(t, result.ObjectValue(), resource.PropertyKey("aws:custom:tag"))
	})

	t.Run("filters aws: prefixed tags not in original", func(t *testing.T) {
		newTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("my-resource"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("aws:elasticfilesystem:default-backup"),
				"value": resource.NewStringProperty("enabled"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Environment"),
				"value": resource.NewStringProperty("prod"),
			}),
		})

		originalTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("my-resource"),
			}),
		})

		result := filterAWSPrefixedTags(newTags, originalTags)

		assert.True(t, result.IsArray())
		assert.Len(t, result.ArrayValue(), 2)
		// Should contain Name and Environment, but not aws:elasticfilesystem:default-backup
		keys := make([]string, 0)
		for _, tag := range result.ArrayValue() {
			keys = append(keys, tag.ObjectValue()["key"].StringValue())
		}
		assert.Contains(t, keys, "Name")
		assert.Contains(t, keys, "Environment")
		assert.NotContains(t, keys, "aws:elasticfilesystem:default-backup")
	})

	t.Run("preserves aws: tags that were in original", func(t *testing.T) {
		newTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("aws:custom:tag"),
				"value": resource.NewStringProperty("value"),
			}),
		})

		originalTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("aws:custom:tag"),
				"value": resource.NewStringProperty("value"),
			}),
		})

		result := filterAWSPrefixedTags(newTags, originalTags)

		assert.True(t, result.IsArray())
		assert.Len(t, result.ArrayValue(), 1)
		assert.Equal(t, "aws:custom:tag", result.ArrayValue()[0].ObjectValue()["key"].StringValue())
	})

	t.Run("handles uppercase Key property", func(t *testing.T) {
		newTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"Key":   resource.NewStringProperty("Name"),
				"Value": resource.NewStringProperty("my-resource"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"Key":   resource.NewStringProperty("aws:servicecatalog:applicationName"),
				"Value": resource.NewStringProperty("my-app"),
			}),
		})

		originalTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"Key":   resource.NewStringProperty("Name"),
				"Value": resource.NewStringProperty("my-resource"),
			}),
		})

		result := filterAWSPrefixedTags(newTags, originalTags)

		assert.True(t, result.IsArray())
		assert.Len(t, result.ArrayValue(), 1)
		assert.Equal(t, "Name", result.ArrayValue()[0].ObjectValue()["Key"].StringValue())
	})

	t.Run("returns null when all tags filtered", func(t *testing.T) {
		newTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("aws:managed:tag"),
				"value": resource.NewStringProperty("value"),
			}),
		})

		originalTags := resource.NewNullProperty()

		result := filterAWSPrefixedTags(newTags, originalTags)

		assert.True(t, result.IsNull())
	})

	t.Run("passes through non-array values unchanged", func(t *testing.T) {
		newTags := resource.NewStringProperty("not-an-array")
		originalTags := resource.NewNullProperty()

		result := filterAWSPrefixedTags(newTags, originalTags)

		assert.Equal(t, newTags, result)
	})
}

func TestSuppressAWSManagedTagAdditions(t *testing.T) {
	t.Run("removes aws: tag additions from diff", func(t *testing.T) {
		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{
				"fileSystemTags": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewObjectProperty(resource.PropertyMap{
						"key":   resource.NewStringProperty("aws:elasticfilesystem:default-backup"),
						"value": resource.NewStringProperty("enabled"),
					}),
				}),
			},
			Updates: map[resource.PropertyKey]resource.ValueDiff{},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		originalInputs := resource.PropertyMap{}

		result := suppressAWSManagedTagAdditions("fileSystemTags", diff, originalInputs)

		_, hasAdd := result.Adds["fileSystemTags"]
		assert.False(t, hasAdd, "aws: prefixed tag addition should be removed")
	})

	t.Run("preserves non-aws: tag additions", func(t *testing.T) {
		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{
				"tags": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewObjectProperty(resource.PropertyMap{
						"key":   resource.NewStringProperty("Name"),
						"value": resource.NewStringProperty("my-resource"),
					}),
				}),
			},
			Updates: map[resource.PropertyKey]resource.ValueDiff{},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		originalInputs := resource.PropertyMap{}

		result := suppressAWSManagedTagAdditions("tags", diff, originalInputs)

		addedTags, hasAdd := result.Adds["tags"]
		assert.True(t, hasAdd)
		assert.Len(t, addedTags.ArrayValue(), 1)
	})

	t.Run("filters aws: tags from updates", func(t *testing.T) {
		oldTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("my-resource"),
			}),
		})

		newTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("my-resource"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("aws:managed:tag"),
				"value": resource.NewStringProperty("value"),
			}),
		})

		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{},
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"tags": {Old: oldTags, New: newTags},
			},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		originalInputs := resource.PropertyMap{
			"tags": oldTags,
		}

		result := suppressAWSManagedTagAdditions("tags", diff, originalInputs)

		// After filtering aws: tag, old and new should be equal, so update should be removed
		_, hasUpdate := result.Updates["tags"]
		assert.False(t, hasUpdate, "update should be removed when filtered tags match old")
	})

	t.Run("keeps user tag updates when aws: tags are filtered out (array)", func(t *testing.T) {
		oldTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("old"),
			}),
		})

		newTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("new"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("aws:managed:tag"),
				"value": resource.NewStringProperty("value"),
			}),
		})

		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{},
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"tags": {Old: oldTags, New: newTags},
			},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		originalInputs := resource.PropertyMap{
			"tags": oldTags,
		}

		result := suppressAWSManagedTagAdditions("tags", diff, originalInputs)

		update, hasUpdate := result.Updates["tags"]
		assert.True(t, hasUpdate, "non-aws tag change should be preserved")
		assert.True(t, update.New.IsArray())
		assert.Len(t, update.New.ArrayValue(), 1)
		assert.Equal(t, "new", update.New.ArrayValue()[0].ObjectValue()["value"].StringValue())
	})

	t.Run("keeps user tag updates when aws: tags are filtered out (map)", func(t *testing.T) {
		oldTags := resource.NewObjectProperty(resource.PropertyMap{
			"Name": resource.NewStringProperty("old"),
		})
		newTags := resource.NewObjectProperty(resource.PropertyMap{
			"Name":            resource.NewStringProperty("new"),
			"aws:managed:tag": resource.NewStringProperty("value"),
		})

		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{},
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"tags": {Old: oldTags, New: newTags},
			},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		originalInputs := resource.PropertyMap{
			"tags": oldTags,
		}

		result := suppressAWSManagedTagAdditions("tags", diff, originalInputs)

		update, hasUpdate := result.Updates["tags"]
		assert.True(t, hasUpdate, "non-aws tag change should be preserved")
		assert.True(t, update.New.IsObject())
		assert.Contains(t, update.New.ObjectValue(), resource.PropertyKey("Name"))
		assert.Equal(t, "new", update.New.ObjectValue()[resource.PropertyKey("Name")].StringValue())
		assert.NotContains(t, update.New.ObjectValue(), resource.PropertyKey("aws:managed:tag"))
	})
}

func TestBuildCfnContext(t *testing.T) {
	t.Run("builds nested structure for multi-segment paths", func(t *testing.T) {
		// For path fileSystemProtection/replicationOverwriteProtection
		// the context should have:
		// - FileSystemProtection.ReplicationOverwriteProtection = value (nested)
		// - ReplicationOverwriteProtection = value (leaf)
		baseContext := map[string]interface{}{}
		path := "fileSystemProtection/replicationOverwriteProtection"
		value := "DISABLED"

		ctx := buildCfnContext(baseContext, path, value, nil)

		// Check leaf value is present
		assert.Equal(t, "DISABLED", ctx["replicationOverwriteProtection"])
		assert.Equal(t, "DISABLED", ctx["ReplicationOverwriteProtection"])

		// Check nested structure with various name variations
		fsp, ok := ctx["FileSystemProtection"].(map[string]interface{})
		assert.True(t, ok, "FileSystemProtection should be a map")
		assert.Equal(t, "DISABLED", fsp["ReplicationOverwriteProtection"])
	})

	t.Run("handles single-segment paths", func(t *testing.T) {
		baseContext := map[string]interface{}{}
		path := "ipProtocol"
		value := "tcp"

		ctx := buildCfnContext(baseContext, path, value, nil)

		assert.Equal(t, "tcp", ctx["ipProtocol"])
		assert.Equal(t, "tcp", ctx["IpProtocol"])
		// Note: Without irreversibleNames lookup, only PascalCase is added
	})

	t.Run("preserves base context", func(t *testing.T) {
		baseContext := map[string]interface{}{
			"existingProp": "existingValue",
		}
		path := "newProp"
		value := "newValue"

		ctx := buildCfnContext(baseContext, path, value, nil)

		assert.Equal(t, "existingValue", ctx["existingProp"])
		assert.Equal(t, "newValue", ctx["newProp"])
	})
}

func TestSuppressAWSManagedDiffs(t *testing.T) {
	t.Run("handles nil diff", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{TagsProperty: "tags"}
		result := SuppressAWSManagedDiffs("aws-native:s3:Bucket", spec, nil, resource.PropertyMap{}, NewTransformCache())
		assert.Nil(t, result)
	})

	t.Run("applies generic tag filtering", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{TagsProperty: "tags"}
		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{
				"tags": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewObjectProperty(resource.PropertyMap{
						"key":   resource.NewStringProperty("aws:managed:tag"),
						"value": resource.NewStringProperty("value"),
					}),
				}),
			},
			Updates: map[resource.PropertyKey]resource.ValueDiff{},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		result := SuppressAWSManagedDiffs("aws-native:s3:Bucket", spec, diff, resource.PropertyMap{}, NewTransformCache())

		_, hasAdd := result.Adds["tags"]
		assert.False(t, hasAdd)
	})

	t.Run("applies EFS-specific suppression via propertyTransform", func(t *testing.T) {
		// EFS replication protection is now handled via propertyTransform
		// The transform normalizes DISABLED -> REPLICATING for comparison
		// NOTE: The actual CloudFormation expression uses nested path: FileSystemProtection.ReplicationOverwriteProtection
		// This tests that buildCfnContext correctly builds the nested context structure
		spec := &metadata.CloudAPIResource{
			TagsProperty: "fileSystemTags",
			PropertyTransforms: map[string]string{
				"fileSystemProtection/replicationOverwriteProtection": "$uppercase(FileSystemProtection.ReplicationOverwriteProtection)='DISABLED' ? 'REPLICATING' : $uppercase(FileSystemProtection.ReplicationOverwriteProtection)",
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

		result := SuppressAWSManagedDiffs("aws-native:efs:FileSystem", spec, diff, oldProps, NewTransformCache())

		_, hasUpdate := result.Updates["fileSystemProtection"]
		assert.False(t, hasUpdate, "DISABLED -> REPLICATING should be suppressed via propertyTransform")
	})

	t.Run("skips tag filtering when no TagsProperty", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{TagsProperty: ""} // No tags property
		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{
				"someProperty": resource.NewStringProperty("value"),
			},
			Updates: map[resource.PropertyKey]resource.ValueDiff{},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		result := SuppressAWSManagedDiffs("aws-native:some:Resource", spec, diff, resource.PropertyMap{}, NewTransformCache())

		// Should pass through unchanged
		_, hasAdd := result.Adds["someProperty"]
		assert.True(t, hasAdd)
	})
}
