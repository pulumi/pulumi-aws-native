// Copyright 2024, Pulumi Corporation.

package resources

import (
	"testing"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Integration tests for propertyTransform support.
// These tests validate the full pipeline from real CloudFormation transform expressions
// through diff suppression with realistic property values.

// TestIntegration_RDS_DBCluster_LowercaseTransforms tests the RDS DBCluster transforms
// which normalize various identifiers to lowercase.
//
// CloudFormation schema excerpt:
//
//	"propertyTransform": {
//	  "/properties/DBClusterIdentifier": "$lowercase(DBClusterIdentifier)",
//	  "/properties/Engine": "$lowercase(Engine)",
//	  "/properties/DBSubnetGroupName": "$lowercase(DBSubnetGroupName)"
//	}
func TestIntegration_RDS_DBCluster_LowercaseTransforms(t *testing.T) {
	// Use unique token to avoid cache pollution from other tests
	resourceToken := "integration-test:rds:DBCluster-lowercase"

	spec := &metadata.CloudAPIResource{
		CfType: "AWS::RDS::DBCluster",
		PropertyTransforms: map[string]string{
			"dbClusterIdentifier":         "$lowercase(DBClusterIdentifier)",
			"engine":                      "$lowercase(Engine)",
			"dbClusterParameterGroupName": "$lowercase(DBClusterParameterGroupName)",
			"dbSubnetGroupName":           "$lowercase(DBSubnetGroupName)",
		},
		// IrreversibleNames maps SDK names to CloudFormation names where simple PascalCase
		// doesn't work (e.g., dbClusterIdentifier -> DBClusterIdentifier, not DbClusterIdentifier).
		// In production, this is populated during schema codegen from CloudFormation specs.
		IrreversibleNames: map[string]string{
			"dbClusterIdentifier":         "DBClusterIdentifier",
			"dbClusterParameterGroupName": "DBClusterParameterGroupName",
			"dbSubnetGroupName":           "DBSubnetGroupName",
		},
	}

	t.Run("suppresses case-only differences in identifiers", func(t *testing.T) {
		oldProps := resource.PropertyMap{
			"dbClusterIdentifier": resource.NewStringProperty("MY-CLUSTER"),
			"engine":              resource.NewStringProperty("AURORA-MYSQL"),
		}
		newProps := resource.PropertyMap{
			"dbClusterIdentifier": resource.NewStringProperty("my-cluster"),
			"engine":              resource.NewStringProperty("aurora-mysql"),
		}
		diff := oldProps.Diff(newProps)
		require.NotNil(t, diff, "diff should not be nil")

		result := suppressPropertyTransformDiffs(resourceToken, spec, diff, oldProps, NewTransformCache())

		assert.Empty(t, result.Updates, "Case-insensitive identifiers should be suppressed")
	})

	t.Run("preserves real value changes", func(t *testing.T) {
		oldProps := resource.PropertyMap{
			"engine": resource.NewStringProperty("aurora-mysql"),
		}
		newProps := resource.PropertyMap{
			"engine": resource.NewStringProperty("aurora-postgresql"),
		}
		diff := oldProps.Diff(newProps)
		require.NotNil(t, diff, "diff should not be nil")

		result := suppressPropertyTransformDiffs(resourceToken, spec, diff, oldProps, NewTransformCache())

		assert.Len(t, result.Updates, 1, "Real engine change should be preserved")
	})
}

// TestIntegration_EC2_SecurityGroup_IpProtocolTransforms tests the EC2 SecurityGroup
// transforms which normalize numeric IP protocol numbers to their string equivalents.
//
// CloudFormation schema excerpt:
//
//	"propertyTransform": {
//	  "/properties/SecurityGroupEgress/*/IpProtocol":
//	    "($mapVal := $lookup({'1': 'icmp','6': 'tcp','17': 'udp','58': 'icmpv6'}, IpProtocol);
//	     $mapVal ? $mapVal : $lowercase(IpProtocol))"
//	}
func TestIntegration_EC2_SecurityGroup_IpProtocolTransforms(t *testing.T) {
	// Use unique token to avoid cache pollution from other tests
	resourceToken := "integration-test:ec2:SecurityGroup-protocol"

	// Simplified version of the actual transform - the real one is more complex
	// but this captures the essence of protocol normalization
	spec := &metadata.CloudAPIResource{
		CfType: "AWS::EC2::SecurityGroup",
		PropertyTransforms: map[string]string{
			// Simplified: just lowercase for testing
			"securityGroupEgress/*/ipProtocol":  "$lowercase(IpProtocol)",
			"securityGroupIngress/*/ipProtocol": "$lowercase(IpProtocol)",
		},
	}

	t.Run("suppresses case differences in protocol names", func(t *testing.T) {
		oldProps := resource.PropertyMap{
			"securityGroupEgress": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"ipProtocol": resource.NewStringProperty("TCP"),
					"fromPort":   resource.NewNumberProperty(443),
					"toPort":     resource.NewNumberProperty(443),
				}),
			}),
		}
		newProps := resource.PropertyMap{
			"securityGroupEgress": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"ipProtocol": resource.NewStringProperty("tcp"),
					"fromPort":   resource.NewNumberProperty(443),
					"toPort":     resource.NewNumberProperty(443),
				}),
			}),
		}
		diff := oldProps.Diff(newProps)
		require.NotNil(t, diff, "diff should not be nil")

		result := suppressPropertyTransformDiffs(resourceToken, spec, diff, oldProps, NewTransformCache())

		assert.Empty(t, result.Updates, "Protocol case difference should be suppressed")
	})

	t.Run("preserves real protocol changes", func(t *testing.T) {
		oldProps := resource.PropertyMap{
			"securityGroupEgress": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"ipProtocol": resource.NewStringProperty("tcp"),
					"fromPort":   resource.NewNumberProperty(443),
					"toPort":     resource.NewNumberProperty(443),
				}),
			}),
		}
		newProps := resource.PropertyMap{
			"securityGroupEgress": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"ipProtocol": resource.NewStringProperty("udp"),
					"fromPort":   resource.NewNumberProperty(443),
					"toPort":     resource.NewNumberProperty(443),
				}),
			}),
		}
		diff := oldProps.Diff(newProps)
		require.NotNil(t, diff, "diff should not be nil")

		result := suppressPropertyTransformDiffs(resourceToken, spec, diff, oldProps, NewTransformCache())

		assert.Len(t, result.Updates, 1, "Real protocol change should be preserved")
	})
}

// TestIntegration_EFS_FileSystem_ReplicationTransform tests the EFS FileSystem transform
// that handles the DISABLED -> REPLICATING transition.
//
// CloudFormation schema excerpt:
//
//	"propertyTransform": {
//	  "/properties/FileSystemProtection/ReplicationOverwriteProtection":
//	    "$uppercase(FileSystemProtection.ReplicationOverwriteProtection)='DISABLED' ?
//	     'REPLICATING' : $uppercase(FileSystemProtection.ReplicationOverwriteProtection)"
//	}
func TestIntegration_EFS_FileSystem_ReplicationTransform(t *testing.T) {
	// Use unique token to avoid cache pollution from other tests
	resourceToken := "integration-test:efs:FileSystem-replication"

	spec := &metadata.CloudAPIResource{
		CfType: "AWS::EFS::FileSystem",
		PropertyTransforms: map[string]string{
			"fileSystemProtection/replicationOverwriteProtection": "$uppercase(ReplicationOverwriteProtection)='DISABLED' ? 'REPLICATING' : $uppercase(ReplicationOverwriteProtection)",
		},
	}

	t.Run("suppresses DISABLED to REPLICATING transition", func(t *testing.T) {
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

		result := suppressPropertyTransformDiffs(resourceToken, spec, diff, oldProps, NewTransformCache())

		assert.Empty(t, result.Updates, "DISABLED -> REPLICATING transition should be suppressed")
	})

	t.Run("preserves ENABLED to DISABLED change", func(t *testing.T) {
		oldProps := resource.PropertyMap{
			"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
				"replicationOverwriteProtection": resource.NewStringProperty("ENABLED"),
			}),
		}
		newProps := resource.PropertyMap{
			"fileSystemProtection": resource.NewObjectProperty(resource.PropertyMap{
				"replicationOverwriteProtection": resource.NewStringProperty("DISABLED"),
			}),
		}
		diff := oldProps.Diff(newProps)
		require.NotNil(t, diff, "diff should not be nil")

		result := suppressPropertyTransformDiffs(resourceToken, spec, diff, oldProps, NewTransformCache())

		assert.Len(t, result.Updates, 1, "ENABLED -> DISABLED change should be preserved")
	})
}

// TestIntegration_TransformCacheThreadSafety verifies the transform cache
// handles concurrent access correctly and returns the same cached instance.
//
// Note: Run with -race flag to detect data races: go test -race ./...
func TestIntegration_TransformCacheThreadSafety(t *testing.T) {
	spec := &metadata.CloudAPIResource{
		CfType: "AWS::Test::Resource",
		PropertyTransforms: map[string]string{
			"name": "$lowercase(Name)",
		},
	}

	cache := &TransformCache{
		transforms: make(map[string][]CompiledTransform),
	}

	// Collect results to verify all goroutines get the same cached instance
	results := make(chan []CompiledTransform, 10)

	// Concurrent access from multiple goroutines
	for i := 0; i < 10; i++ {
		go func() {
			// Use unique token to avoid cache pollution from other tests
			transforms := cache.GetOrCompile("integration-test:test:Resource-threadsafety", spec)
			results <- transforms
		}()
	}

	// Collect all results
	var allResults [][]CompiledTransform
	for i := 0; i < 10; i++ {
		allResults = append(allResults, <-results)
	}

	// Verify all goroutines got results
	require.Len(t, allResults, 10)

	// Verify all results have correct length
	for i, transforms := range allResults {
		assert.Len(t, transforms, 1, "goroutine %d should get 1 transform", i)
	}

	// Verify all goroutines received the SAME cached slice (pointer equality).
	// This confirms the double-check locking in GetOrCompile works correctly
	// and all concurrent callers receive the same cached instance.
	first := allResults[0]
	for i := 1; i < len(allResults); i++ {
		// Compare the Expression pointer to verify same compiled transform instance
		assert.Same(t, first[0].Expression, allResults[i][0].Expression,
			"goroutine %d should receive same cached Expression instance", i)
	}
}

// TestIntegration_TransformCacheDifferentResourceTypes verifies the transform cache
// keeps different resource types isolated from each other.
func TestIntegration_TransformCacheDifferentResourceTypes(t *testing.T) {
	rdsSpec := &metadata.CloudAPIResource{
		CfType: "AWS::RDS::DBCluster",
		PropertyTransforms: map[string]string{
			"engine": "$lowercase(Engine)",
		},
	}

	ecsSpec := &metadata.CloudAPIResource{
		CfType: "AWS::ECS::Cluster",
		PropertyTransforms: map[string]string{
			"clusterName": "$uppercase(ClusterName)",
		},
	}

	cache := &TransformCache{
		transforms: make(map[string][]CompiledTransform),
	}

	// Get transforms for RDS
	rdsTransforms := cache.GetOrCompile("integration-test:rds:DBCluster-isolation", rdsSpec)
	require.Len(t, rdsTransforms, 1)
	assert.Equal(t, "engine", rdsTransforms[0].Path)

	// Get transforms for ECS - should be separate
	ecsTransforms := cache.GetOrCompile("integration-test:ecs:Cluster-isolation", ecsSpec)
	require.Len(t, ecsTransforms, 1)
	assert.Equal(t, "clusterName", ecsTransforms[0].Path)

	// Verify they're different
	assert.NotEqual(t, rdsTransforms[0].Path, ecsTransforms[0].Path)
}

// TestIntegration_ComplexLookupTransform tests a more complex transform with $lookup.
func TestIntegration_ComplexLookupTransform(t *testing.T) {
	// This tests a simplified version of the SecurityGroup IpProtocol transform
	spec := &metadata.CloudAPIResource{
		CfType: "AWS::EC2::SecurityGroup",
		PropertyTransforms: map[string]string{
			// Test $lookup functionality
			"protocol": "$lookup({'6': 'tcp', '17': 'udp'}, Protocol)",
		},
	}

	transforms := GlobalTransformCache.GetOrCompile("test-lookup", spec)
	require.Len(t, transforms, 1)

	t.Run("lookup transforms numeric to string", func(t *testing.T) {
		result, err := EvaluateTransform(transforms[0], "6", map[string]interface{}{
			"Protocol": "6",
		})
		require.NoError(t, err)
		assert.Equal(t, "tcp", result)
	})

	t.Run("lookup returns original value for unknown key", func(t *testing.T) {
		result, err := EvaluateTransform(transforms[0], "99", map[string]interface{}{
			"Protocol": "99",
		})
		// JSONata $lookup returns undefined for missing keys, which we handle gracefully
		require.NoError(t, err)
		// Result should be the original value since lookup returned undefined
		assert.Equal(t, "99", result)
	})
}

// TestIntegration_MultipleArrayElements tests transforms applied to multiple array elements.
func TestIntegration_MultipleArrayElements(t *testing.T) {
	// Use unique token to avoid cache pollution from other tests
	resourceToken := "integration-test:ec2:SecurityGroup-array"

	spec := &metadata.CloudAPIResource{
		CfType: "AWS::EC2::SecurityGroup",
		PropertyTransforms: map[string]string{
			"rules/*/protocol": "$lowercase(Protocol)",
		},
	}

	oldProps := resource.PropertyMap{
		"rules": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"protocol": resource.NewStringProperty("TCP"),
				"port":     resource.NewNumberProperty(80),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"protocol": resource.NewStringProperty("UDP"),
				"port":     resource.NewNumberProperty(53),
			}),
		}),
	}
	newProps := resource.PropertyMap{
		"rules": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"protocol": resource.NewStringProperty("tcp"),
				"port":     resource.NewNumberProperty(80),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"protocol": resource.NewStringProperty("udp"),
				"port":     resource.NewNumberProperty(53),
			}),
		}),
	}
	diff := oldProps.Diff(newProps)
	require.NotNil(t, diff, "diff should not be nil")

	result := suppressPropertyTransformDiffs(resourceToken, spec, diff, oldProps, NewTransformCache())

	assert.Empty(t, result.Updates, "Case changes across multiple array elements should be suppressed")
}

// TestIntegration_NoTransformsNoEffect verifies that resources without transforms
// are handled correctly (passthrough behavior).
func TestIntegration_NoTransformsNoEffect(t *testing.T) {
	// Use unique token to avoid cache pollution from other tests
	resourceToken := "integration-test:s3:Bucket-notransforms"

	spec := &metadata.CloudAPIResource{
		CfType:             "AWS::S3::Bucket",
		PropertyTransforms: nil, // No transforms
	}

	oldProps := resource.PropertyMap{
		"bucketName": resource.NewStringProperty("old-bucket"),
	}
	newProps := resource.PropertyMap{
		"bucketName": resource.NewStringProperty("new-bucket"),
	}
	diff := oldProps.Diff(newProps)
	require.NotNil(t, diff, "diff should not be nil")

	result := suppressPropertyTransformDiffs(resourceToken, spec, diff, oldProps, NewTransformCache())

	assert.Len(t, result.Updates, 1, "Diff should be preserved when no transforms exist")
}
