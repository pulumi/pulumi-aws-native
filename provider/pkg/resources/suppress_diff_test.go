// Copyright 2024, Pulumi Corporation.

package resources

import (
	"testing"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
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

func TestSuppressEFSReplicationProtectionTransition(t *testing.T) {
	t.Run("suppresses REPLICATING transition", func(t *testing.T) {
		oldFsp := resource.NewObjectProperty(resource.PropertyMap{
			"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
		})
		newFsp := resource.NewObjectProperty(resource.PropertyMap{
			"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
		})

		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{},
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"fileSystemProtection": {Old: oldFsp, New: newFsp},
			},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		result := suppressEFSReplicationProtectionTransition(diff)

		_, hasUpdate := result.Updates["fileSystemProtection"]
		assert.False(t, hasUpdate, "REPLICATING transition should be suppressed")
	})

	t.Run("allows non-REPLICATING updates", func(t *testing.T) {
		oldFsp := resource.NewObjectProperty(resource.PropertyMap{
			"replicationOverwriteProtection": resource.NewStringProperty("ENABLED"),
		})
		newFsp := resource.NewObjectProperty(resource.PropertyMap{
			"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
		})

		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{},
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"fileSystemProtection": {Old: oldFsp, New: newFsp},
			},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		result := suppressEFSReplicationProtectionTransition(diff)

		// Should not be suppressed - it's not REPLICATING
		_, hasUpdate := result.Updates["fileSystemProtection"]
		assert.True(t, hasUpdate)
	})

	t.Run("ignores non-object updates", func(t *testing.T) {
		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{},
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"fileSystemProtection": {
					Old: resource.NewStringProperty("old"),
					New: resource.NewStringProperty("new"),
				},
			},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		result := suppressEFSReplicationProtectionTransition(diff)

		_, hasUpdate := result.Updates["fileSystemProtection"]
		assert.True(t, hasUpdate, "non-object updates should pass through")
	})
}

func TestSuppressAWSManagedDiffs(t *testing.T) {
	t.Run("handles nil diff", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{TagsProperty: "tags"}
		result := SuppressAWSManagedDiffs("aws-native:s3:Bucket", spec, nil, resource.PropertyMap{})
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

		result := SuppressAWSManagedDiffs("aws-native:s3:Bucket", spec, diff, resource.PropertyMap{})

		_, hasAdd := result.Adds["tags"]
		assert.False(t, hasAdd)
	})

	t.Run("applies EFS-specific suppression", func(t *testing.T) {
		spec := &metadata.CloudAPIResource{TagsProperty: "fileSystemTags"}
		oldFsp := resource.NewObjectProperty(resource.PropertyMap{
			"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
		})
		newFsp := resource.NewObjectProperty(resource.PropertyMap{
			"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
		})

		diff := &resource.ObjectDiff{
			Adds: resource.PropertyMap{},
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"fileSystemProtection": {Old: oldFsp, New: newFsp},
			},
			Deletes: resource.PropertyMap{},
			Sames:   resource.PropertyMap{},
		}

		originalInputs := resource.PropertyMap{
			"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
				"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
			}),
		}

		result := SuppressAWSManagedDiffs("aws-native:efs:FileSystem", spec, diff, originalInputs)

		_, hasUpdate := result.Updates["fileSystemProtection"]
		assert.False(t, hasUpdate)
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

		result := SuppressAWSManagedDiffs("aws-native:some:Resource", spec, diff, resource.PropertyMap{})

		// Should pass through unchanged
		_, hasAdd := result.Adds["someProperty"]
		assert.True(t, hasAdd)
	})
}
