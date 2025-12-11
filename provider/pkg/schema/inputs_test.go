// Copyright 2016-2024, Pulumi Corporation.

package schema

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetPropertiesWithNestedReadOnly(t *testing.T) {
	tests := []struct {
		name     string
		readOnly []string
		expected map[string]bool
	}{
		{
			name:     "empty read-only list",
			readOnly: []string{},
			expected: map[string]bool{},
		},
		{
			name:     "top-level read-only properties are excluded",
			readOnly: []string{"arn", "fileSystemId"},
			expected: map[string]bool{},
		},
		{
			name:     "nested read-only paths return top-level property",
			readOnly: []string{"replicationConfiguration/destinations/*/status"},
			expected: map[string]bool{"replicationConfiguration": true},
		},
		{
			name: "mixed read-only paths",
			readOnly: []string{
				"arn",
				"fileSystemId",
				"replicationConfiguration/destinations/*/status",
				"replicationConfiguration/destinations/*/statusMessage",
			},
			expected: map[string]bool{"replicationConfiguration": true},
		},
		{
			name: "multiple properties with nested read-only",
			readOnly: []string{
				"propA/nested/value",
				"propB/other/nested",
			},
			expected: map[string]bool{"propA": true, "propB": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getPropertiesWithNestedReadOnly(tt.readOnly)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFilterEffectivelyReadOnlyProperties(t *testing.T) {
	tests := []struct {
		name            string
		refreshedInputs resource.PropertyMap
		originalInputs  resource.PropertyMap
		readOnly        []string
		expected        resource.PropertyMap
	}{
		{
			name: "keeps property that was in original inputs",
			refreshedInputs: resource.PropertyMap{
				"replicationConfiguration": resource.NewObjectProperty(resource.PropertyMap{
					"destinations": resource.NewArrayProperty([]resource.PropertyValue{}),
				}),
			},
			originalInputs: resource.PropertyMap{
				"replicationConfiguration": resource.NewObjectProperty(resource.PropertyMap{}),
			},
			readOnly: []string{"replicationConfiguration/destinations/*/status"},
			expected: resource.PropertyMap{
				"replicationConfiguration": resource.NewObjectProperty(resource.PropertyMap{
					"destinations": resource.NewArrayProperty([]resource.PropertyValue{}),
				}),
			},
		},
		{
			name: "removes property with nested read-only that was NOT in original inputs",
			refreshedInputs: resource.PropertyMap{
				"replicationConfiguration": resource.NewObjectProperty(resource.PropertyMap{
					"destinations": resource.NewArrayProperty([]resource.PropertyValue{}),
				}),
			},
			originalInputs: resource.PropertyMap{},
			readOnly:       []string{"replicationConfiguration/destinations/*/status"},
			expected:       resource.PropertyMap{},
		},
		{
			name: "keeps property without nested read-only paths",
			refreshedInputs: resource.PropertyMap{
				"encrypted": resource.NewBoolProperty(true),
			},
			originalInputs: resource.PropertyMap{},
			readOnly:       []string{"arn", "fileSystemId"},
			expected: resource.PropertyMap{
				"encrypted": resource.NewBoolProperty(true),
			},
		},
		{
			name: "EFS replication scenario - filters replicationConfiguration from destination",
			refreshedInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
				}),
				"replicationConfiguration": resource.NewObjectProperty(resource.PropertyMap{
					"destinations": resource.NewArrayProperty([]resource.PropertyValue{
						resource.NewObjectProperty(resource.PropertyMap{
							"fileSystemId": resource.NewStringProperty("fs-12345"),
							"region":       resource.NewStringProperty("us-west-2"),
							"status":       resource.NewStringProperty("ENABLED"),
						}),
					}),
				}),
			},
			originalInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
				}),
			},
			readOnly: []string{
				"arn",
				"fileSystemId",
				"replicationConfiguration/destinations/*/status",
				"replicationConfiguration/destinations/*/statusMessage",
			},
			expected: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
				}),
			},
		},
		{
			name: "keeps multiple properties that were in original inputs",
			refreshedInputs: resource.PropertyMap{
				"encrypted":                    resource.NewBoolProperty(true),
				"performanceMode":              resource.NewStringProperty("generalPurpose"),
				"throughputMode":               resource.NewStringProperty("elastic"),
				"kmsKeyId":                     resource.NewStringProperty("arn:aws:kms:us-east-1:123456789012:key/abc"),
				"provisionedThroughputInMibps": resource.NewNumberProperty(100),
			},
			originalInputs: resource.PropertyMap{
				"encrypted":       resource.NewBoolProperty(true),
				"performanceMode": resource.NewStringProperty("generalPurpose"),
			},
			readOnly: []string{"arn", "fileSystemId"},
			expected: resource.PropertyMap{
				"encrypted":                    resource.NewBoolProperty(true),
				"performanceMode":              resource.NewStringProperty("generalPurpose"),
				"throughputMode":               resource.NewStringProperty("elastic"),
				"kmsKeyId":                     resource.NewStringProperty("arn:aws:kms:us-east-1:123456789012:key/abc"),
				"provisionedThroughputInMibps": resource.NewNumberProperty(100),
			},
		},
		{
			name: "handles nil original inputs",
			refreshedInputs: resource.PropertyMap{
				"encrypted": resource.NewBoolProperty(true),
			},
			originalInputs: nil,
			readOnly:       []string{"arn"},
			expected: resource.PropertyMap{
				"encrypted": resource.NewBoolProperty(true),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterEffectivelyReadOnlyProperties(tt.refreshedInputs, tt.originalInputs, tt.readOnly)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestApplyReadOnlyNestedProperties(t *testing.T) {
	tests := []struct {
		name      string
		newInputs resource.PropertyMap
		oldInputs resource.PropertyMap
		readOnly  []string
		expected  resource.PropertyMap
	}{
		{
			name: "applies nested read-only value from old to new",
			newInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
				}),
			},
			oldInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
				}),
			},
			readOnly: []string{"fileSystemProtection/replicationOverwriteProtection"},
			expected: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
				}),
			},
		},
		{
			name: "no change when old and new have same value",
			newInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("ENABLED"),
				}),
			},
			oldInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("ENABLED"),
				}),
			},
			readOnly: []string{"fileSystemProtection/replicationOverwriteProtection"},
			expected: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("ENABLED"),
				}),
			},
		},
		{
			name: "no change when property not in read-only list",
			newInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
				}),
			},
			oldInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
				}),
			},
			readOnly: []string{"arn", "fileSystemId"},
			expected: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
				}),
			},
		},
		{
			name: "no change when property not in old inputs",
			newInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
				}),
			},
			oldInputs: resource.PropertyMap{},
			readOnly:  []string{"fileSystemProtection/replicationOverwriteProtection"},
			expected: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
				}),
			},
		},
		{
			name: "no change when property not in new inputs",
			newInputs: resource.PropertyMap{
				"encrypted": resource.NewBoolProperty(true),
			},
			oldInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
				}),
			},
			readOnly: []string{"fileSystemProtection/replicationOverwriteProtection"},
			expected: resource.PropertyMap{
				"encrypted": resource.NewBoolProperty(true),
			},
		},
		{
			name: "preserves other properties in the object",
			newInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
					"otherProperty":                  resource.NewStringProperty("value"),
				}),
			},
			oldInputs: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
				}),
			},
			readOnly: []string{"fileSystemProtection/replicationOverwriteProtection"},
			expected: resource.PropertyMap{
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
					"otherProperty":                  resource.NewStringProperty("value"),
				}),
			},
		},
		{
			name: "EFS FileSystem scenario - applies REPLICATING from refresh",
			newInputs: resource.PropertyMap{
				"encrypted":       resource.NewBoolProperty(true),
				"performanceMode": resource.NewStringProperty("generalPurpose"),
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
				}),
			},
			oldInputs: resource.PropertyMap{
				"encrypted":       resource.NewBoolProperty(true),
				"performanceMode": resource.NewStringProperty("generalPurpose"),
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
				}),
			},
			readOnly: []string{
				"arn",
				"fileSystemId",
				"replicationConfiguration/destinations/*/status",
				"fileSystemProtection/replicationOverwriteProtection",
			},
			expected: resource.PropertyMap{
				"encrypted":       resource.NewBoolProperty(true),
				"performanceMode": resource.NewStringProperty("generalPurpose"),
				"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
					"replicationOverwriteProtection": resource.NewStringProperty("REPLICATING"),
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ApplyReadOnlyNestedProperties(tt.newInputs, tt.oldInputs, tt.readOnly)
			assert.Equal(t, tt.expected, result)
		})
	}
}
