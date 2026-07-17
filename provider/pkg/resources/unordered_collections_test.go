// Copyright 2026, Pulumi Corporation.

//nolint:goconst // Repeated property names keep nested PropertyMap fixtures readable.
package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/default_tags"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
)

func TestSuppressUnorderedCollectionDiffs(t *testing.T) {
	t.Run("primitive reorder preserves duplicates", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"values": unorderedTestArray("A", "A", "B")}
		newInputs := resource.PropertyMap{"values": unorderedTestArray("B", "A", "A")}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		assert.Nil(t, diff)
	})

	for _, test := range []struct {
		name string
		old  resource.PropertyValue
		new  resource.PropertyValue
	}{
		{name: "element addition", old: unorderedTestArray("A"), new: unorderedTestArray("A", "B")},
		{name: "element removal", old: unorderedTestArray("A", "B"), new: unorderedTestArray("A")},
		{name: "element value change", old: unorderedTestArray("A", "B"), new: unorderedTestArray("A", "C")},
		{name: "duplicate count change", old: unorderedTestArray("A", "A", "B"), new: unorderedTestArray("A", "B", "B")},
		{name: "null is not empty", old: resource.NewNullProperty(), new: unorderedTestArray()},
	} {
		t.Run(test.name, func(t *testing.T) {
			oldInputs := resource.PropertyMap{"values": test.old}
			newInputs := resource.PropertyMap{"values": test.new}
			diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
				UnorderedCollections: []string{"values"},
			}, "aws-native:test:Resource", oldInputs, newInputs)
			require.NotNil(t, diff)
			assert.True(t, diff.AnyChanges())
		})
	}

	t.Run("ordered array remains positional", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"values": unorderedTestArray("A", "B")}
		newInputs := resource.PropertyMap{"values": unorderedTestArray("B", "A")}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{},
			"aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})

	t.Run("multiple unordered roots", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			"first":  unorderedTestArray("A", "B"),
			"second": unorderedTestArray("C", "D"),
		}
		newInputs := resource.PropertyMap{
			"first":  unorderedTestArray("B", "A"),
			"second": unorderedTestArray("D", "C"),
		}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"first", "second"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		assert.Nil(t, diff)
	})

	t.Run("order-only root is suppressed alongside real sibling change", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			"mode":   resource.NewStringProperty("old"),
			"values": unorderedTestArray("A", "B"),
		}
		newInputs := resource.PropertyMap{
			"mode":   resource.NewStringProperty("new"),
			"values": unorderedTestArray("B", "A"),
		}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.Contains(t, diff.Updates, resource.PropertyKey("mode"))
		assert.NotContains(t, diff.Updates, resource.PropertyKey("values"))
	})

	t.Run("nested unordered array under ordered parent", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"groups": unorderedTestObjectArray(
			resource.NewObjectProperty(resource.PropertyMap{
				"name":   resource.NewStringProperty("group"),
				"values": unorderedTestArray("A", "B"),
			}),
		)}
		newInputs := resource.PropertyMap{"groups": unorderedTestObjectArray(
			resource.NewObjectProperty(resource.PropertyMap{
				"name":   resource.NewStringProperty("group"),
				"values": unorderedTestArray("B", "A"),
			}),
		)}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"groups/*/values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		assert.Nil(t, diff)
	})

	t.Run("object reorder", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"values": unorderedTestObjectArray(
			unorderedTestObject("name", "A", "value", "1"),
			unorderedTestObject("name", "B", "value", "2"),
		)}
		newInputs := resource.PropertyMap{"values": unorderedTestObjectArray(
			unorderedTestObject("name", "B", "value", "2"),
			unorderedTestObject("name", "A", "value", "1"),
		)}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		assert.Nil(t, diff)
	})

	t.Run("unknown placeholders preserve the diff", func(t *testing.T) {
		unknown := resource.MakeComputed(resource.NewStringProperty(""))
		oldInputs := resource.PropertyMap{"values": resource.NewArrayProperty([]resource.PropertyValue{
			unknown, resource.NewStringProperty("known"),
		})}
		newInputs := resource.PropertyMap{"values": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewStringProperty("known"), unknown,
		})}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})

	t.Run("unknown inside known output wrapper preserves the diff", func(t *testing.T) {
		wrappedUnknown := resource.NewOutputProperty(resource.Output{
			Known:   true,
			Element: resource.MakeComputed(resource.NewStringProperty("")),
		})
		oldInputs := resource.PropertyMap{"values": resource.NewArrayProperty([]resource.PropertyValue{
			wrappedUnknown, resource.NewStringProperty("known"),
		})}
		newInputs := resource.PropertyMap{"values": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewStringProperty("known"), wrappedUnknown,
		})}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})

	t.Run("element secret wrappers conservatively preserve the diff", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"values": resource.NewArrayProperty([]resource.PropertyValue{
			resource.MakeSecret(resource.NewStringProperty("A")),
			resource.MakeSecret(resource.NewStringProperty("B")),
		})}
		newInputs := resource.PropertyMap{"values": resource.NewArrayProperty([]resource.PropertyValue{
			resource.MakeSecret(resource.NewStringProperty("B")),
			resource.MakeSecret(resource.NewStringProperty("A")),
		})}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})

	t.Run("collection secret wrapper is retained", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			"values": resource.MakeSecret(unorderedTestArray("A", "B")),
		}
		newInputs := resource.PropertyMap{
			"values": resource.MakeSecret(unorderedTestArray("B", "A")),
		}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		assert.Nil(t, diff)
	})

	t.Run("selective element secret moving with its value preserves the diff", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"values": resource.NewArrayProperty([]resource.PropertyValue{
			resource.MakeSecret(resource.NewStringProperty("A")),
			resource.NewStringProperty("B"),
		})}
		newInputs := resource.PropertyMap{"values": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewStringProperty("B"),
			resource.MakeSecret(resource.NewStringProperty("A")),
		})}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})

	t.Run("write-only guard is opaque to positional transforms", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"actions": unorderedTestObjectArray(
			unorderedTestObject("name", "UPPER", "secret", "hidden"),
		)}
		newInputs := resource.PropertyMap{"actions": unorderedTestObjectArray(
			unorderedTestObject("name", "upper", "secret", "hidden"),
		)}
		spec := &metadata.CloudAPIResource{
			UnorderedCollections: []string{"actions"},
			WriteOnly:            []string{"actions/*/secret"},
			PropertyTransforms: map[string]string{
				"actions/*/name": "$lowercase(Name)",
			},
		}
		diff := suppressUnorderedTestDiff(t, spec, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})

	t.Run("whole write-only unordered collection can still be compared", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"values": unorderedTestArray("A", "B")}
		newInputs := resource.PropertyMap{"values": unorderedTestArray("B", "A")}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
			WriteOnly:            []string{"values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		assert.Nil(t, diff)
	})

	t.Run("property transform normalization is conservatively deferred", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"values": unorderedTestObjectArray(
			unorderedTestObject("name", "A", "mode", "UPPER"),
			unorderedTestObject("name", "B", "mode", "LOWER"),
		)}
		newInputs := resource.PropertyMap{"values": unorderedTestObjectArray(
			unorderedTestObject("name", "B", "mode", "lower"),
			unorderedTestObject("name", "A", "mode", "upper"),
		)}
		spec := &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
			PropertyTransforms: map[string]string{
				"values/*/mode": "$lowercase(Mode)",
			},
		}
		diff := suppressUnorderedTestDiff(t, spec, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})

	t.Run("parent unordered array retains ordered child semantics", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"groups": unorderedTestObjectArray(
			resource.NewObjectProperty(resource.PropertyMap{
				"name":   resource.NewStringProperty("A"),
				"values": unorderedTestArray("1", "2"),
			}),
		)}
		newInputs := resource.PropertyMap{"groups": unorderedTestObjectArray(
			resource.NewObjectProperty(resource.PropertyMap{
				"name":   resource.NewStringProperty("A"),
				"values": unorderedTestArray("2", "1"),
			}),
		)}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"groups"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})

	t.Run("nested unordered child composes with unordered parent", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"groups": unorderedTestObjectArray(
			resource.NewObjectProperty(resource.PropertyMap{
				"name":   resource.NewStringProperty("A"),
				"values": unorderedTestArray("1", "2"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"name":   resource.NewStringProperty("B"),
				"values": unorderedTestArray("3", "4"),
			}),
		)}
		newInputs := resource.PropertyMap{"groups": unorderedTestObjectArray(
			resource.NewObjectProperty(resource.PropertyMap{
				"name":   resource.NewStringProperty("B"),
				"values": unorderedTestArray("4", "3"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"name":   resource.NewStringProperty("A"),
				"values": unorderedTestArray("2", "1"),
			}),
		)}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"groups", "groups/*/values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		assert.Nil(t, diff)
	})

	t.Run("order-only create-only children are suppressed", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"values": unorderedTestObjectArray(
			unorderedTestObject("name", "A"), unorderedTestObject("name", "B"),
		)}
		newInputs := resource.PropertyMap{"values": unorderedTestObjectArray(
			unorderedTestObject("name", "B"), unorderedTestObject("name", "A"),
		)}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
			CreateOnly:           []string{"values/*/name"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		assert.Nil(t, diff)
	})

	t.Run("real create-only child change remains", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"values": unorderedTestObjectArray(unorderedTestObject("name", "A"))}
		newInputs := resource.PropertyMap{"values": unorderedTestObjectArray(unorderedTestObject("name", "B"))}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
			CreateOnly:           []string{"values/*/name"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})

	t.Run("collection exceeding comparison budget preserves diff", func(t *testing.T) {
		const tooLarge = 257
		oldValues := make([]resource.PropertyValue, tooLarge)
		newValues := make([]resource.PropertyValue, tooLarge)
		for i := range oldValues {
			oldValues[i] = resource.NewNumberProperty(float64(i))
			newValues[len(newValues)-1-i] = oldValues[i]
		}
		oldInputs := resource.PropertyMap{"values": resource.NewArrayProperty(oldValues)}
		newInputs := resource.PropertyMap{"values": resource.NewArrayProperty(newValues)}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"values"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})

	t.Run("comparison budget is consumed in stable path order", func(t *testing.T) {
		const collectionSize = 182 // Two roots exceed the aggregate 256*256 budget.
		oldValues := make([]resource.PropertyValue, collectionSize)
		newValues := make([]resource.PropertyValue, collectionSize)
		for i := range oldValues {
			oldValues[i] = resource.NewNumberProperty(float64(i))
			newValues[len(newValues)-1-i] = oldValues[i]
		}
		oldArray := resource.NewArrayProperty(oldValues)
		newArray := resource.NewArrayProperty(newValues)
		oldInputs := resource.PropertyMap{"first": oldArray, "second": oldArray}
		newInputs := resource.PropertyMap{"first": newArray, "second": newArray}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"first", "second"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.NotContains(t, diff.Updates, resource.PropertyKey("first"))
		assert.Contains(t, diff.Updates, resource.PropertyKey("second"))
	})

	t.Run("IAM policy normalization composes with reorder", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"policies": unorderedTestObjectArray(
			unorderedTestPolicy("A", resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("s3:GetObject"),
			})),
			unorderedTestPolicy("B", resource.NewStringProperty("s3:PutObject")),
		)}
		newInputs := resource.PropertyMap{"policies": unorderedTestObjectArray(
			unorderedTestPolicy("B", resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("s3:PutObject"),
			})),
			unorderedTestPolicy("A", resource.NewStringProperty("s3:GetObject")),
		)}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			UnorderedCollections: []string{"policies"},
		}, awsNativeIAMRoleToken, oldInputs, newInputs)
		assert.Nil(t, diff)
	})

	t.Run("AWS-managed tags filter before unordered comparison", func(t *testing.T) {
		tag := func(key, value string) resource.PropertyValue {
			return unorderedTestObject("key", key, "value", value)
		}
		oldActual := resource.PropertyMap{"tags": unorderedTestObjectArray(
			tag("Name", "resource"), tag("aws:managed", "generated"), tag("Environment", "prod"),
		)}
		desired := resource.PropertyMap{"tags": unorderedTestObjectArray(
			tag("Environment", "prod"), tag("Name", "resource"),
		)}
		spec := &metadata.CloudAPIResource{
			TagsProperty:         "tags",
			TagsStyle:            default_tags.TagsStyleKeyValueArray,
			UnorderedCollections: []string{"tags"},
		}
		diff := oldActual.Diff(desired)
		require.NotNil(t, diff)
		diff = SuppressAWSManagedDiffsWithContext(
			"aws-native:test:Resource", spec, diff, desired, oldActual, desired, NewTransformCache(),
		)
		assert.Nil(t, diff)
	})

	t.Run("unordered tags preserve secret and plaintext equality after reorder", func(t *testing.T) {
		tag := func(key string, value resource.PropertyValue) resource.PropertyValue {
			return resource.NewObjectProperty(resource.PropertyMap{
				"key": resource.NewStringProperty(key), "value": value,
			})
		}
		actual := resource.PropertyMap{"tags": unorderedTestObjectArray(
			tag("Environment", resource.NewStringProperty("prod")),
			tag("Password", resource.NewStringProperty("secret")),
		)}
		desired := resource.PropertyMap{"tags": unorderedTestObjectArray(
			tag("Password", resource.MakeSecret(resource.NewStringProperty("secret"))),
			tag("Environment", resource.NewStringProperty("prod")),
		)}
		spec := &metadata.CloudAPIResource{
			TagsProperty:         "tags",
			TagsStyle:            default_tags.TagsStyleKeyValueArray,
			UnorderedCollections: []string{"tags"},
		}
		diff := actual.Diff(desired)
		require.NotNil(t, diff)
		diff = SuppressAWSManagedDiffsWithContext(
			"aws-native:test:Resource", spec, diff, desired, actual, desired, NewTransformCache(),
		)
		assert.Nil(t, diff)
	})

	t.Run("unordered tags with unknown values preserve the diff", func(t *testing.T) {
		tag := func(key string, value resource.PropertyValue) resource.PropertyValue {
			return resource.NewObjectProperty(resource.PropertyMap{
				"key": resource.NewStringProperty(key), "value": value,
			})
		}
		unknown := resource.MakeComputed(resource.NewStringProperty(""))
		oldInputs := resource.PropertyMap{"tags": unorderedTestObjectArray(
			tag("A", unknown), tag("B", resource.NewStringProperty("known")),
		)}
		newInputs := resource.PropertyMap{"tags": unorderedTestObjectArray(
			tag("B", resource.NewStringProperty("known")), tag("A", unknown),
		)}
		diff := suppressUnorderedTestDiff(t, &metadata.CloudAPIResource{
			TagsProperty:         "tags",
			TagsStyle:            default_tags.TagsStyleKeyValueArray,
			UnorderedCollections: []string{"tags"},
		}, "aws-native:test:Resource", oldInputs, newInputs)
		require.NotNil(t, diff)
		assert.True(t, diff.AnyChanges())
	})
}

func TestHasPerfectMatching(t *testing.T) {
	tests := []struct {
		name       string
		equivalent [][]bool
		expected   bool
	}{
		{
			name: "reassigns an earlier greedy match",
			equivalent: [][]bool{
				{true, true},
				{true, false},
			},
			expected: true,
		},
		{
			name: "rejects when no distinct assignment exists",
			equivalent: [][]bool{
				{true, false},
				{true, false},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, hasPerfectMatching(test.equivalent))
		})
	}
}

func TestSuppressBaselineDiffsPreservesNestedUnorderedOrder(t *testing.T) {
	oldDesired := resource.PropertyMap{
		"config": resource.NewObjectProperty(resource.PropertyMap{
			"mode":  resource.NewStringProperty("old"),
			"items": unorderedTestArray("B", "A"),
		}),
	}
	baseline := resource.PropertyMap{
		"config": resource.NewObjectProperty(resource.PropertyMap{
			"mode":  resource.NewStringProperty("new"),
			"items": unorderedTestArray("A", "B"),
		}),
	}

	result := SuppressBaselineDiffs("aws-native:test:Resource", &metadata.CloudAPIResource{
		UnorderedCollections: []string{"config/items"},
	}, oldDesired, baseline, NewTransformCache())

	expected := resource.PropertyMap{
		"config": resource.NewObjectProperty(resource.PropertyMap{
			"mode":  resource.NewStringProperty("new"),
			"items": unorderedTestArray("B", "A"),
		}),
	}
	assert.True(t, expected.DeepEquals(result), "expected %v, got %v", expected, result)
}

func TestSuppressBaselineDiffsAdoptsRealUnorderedMembershipChange(t *testing.T) {
	oldDesired := resource.PropertyMap{
		"config": resource.NewObjectProperty(resource.PropertyMap{
			"mode":  resource.NewStringProperty("old"),
			"items": unorderedTestArray("B", "A"),
		}),
	}
	baseline := resource.PropertyMap{
		"config": resource.NewObjectProperty(resource.PropertyMap{
			"mode":  resource.NewStringProperty("new"),
			"items": unorderedTestArray("A", "C"),
		}),
	}

	result := SuppressBaselineDiffs("aws-native:test:Resource", &metadata.CloudAPIResource{
		UnorderedCollections: []string{"config/items"},
	}, oldDesired, baseline, NewTransformCache())

	assert.True(t, baseline.DeepEquals(result), "expected %v, got %v", baseline, result)
}

func TestSuppressBaselineDiffsPreservesNestedOrderInsideArray(t *testing.T) {
	oldDesired := resource.PropertyMap{
		"groups": unorderedTestObjectArray(resource.NewObjectProperty(resource.PropertyMap{
			"mode":   resource.NewStringProperty("old"),
			"values": unorderedTestArray("B", "A"),
		})),
	}
	baseline := resource.PropertyMap{
		"groups": unorderedTestObjectArray(resource.NewObjectProperty(resource.PropertyMap{
			"mode":   resource.NewStringProperty("new"),
			"values": unorderedTestArray("A", "B"),
		})),
	}

	result := SuppressBaselineDiffs("aws-native:test:Resource", &metadata.CloudAPIResource{
		UnorderedCollections: []string{"groups/*/values"},
	}, oldDesired, baseline, NewTransformCache())

	expected := resource.PropertyMap{
		"groups": unorderedTestObjectArray(resource.NewObjectProperty(resource.PropertyMap{
			"mode":   resource.NewStringProperty("new"),
			"values": unorderedTestArray("B", "A"),
		})),
	}
	assert.True(t, expected.DeepEquals(result), "expected %v, got %v", expected, result)
}

func TestSuppressBaselineDiffsMaterializesArrayStructuralChanges(t *testing.T) {
	item := func(name string, values ...string) resource.PropertyValue {
		return resource.NewObjectProperty(resource.PropertyMap{
			"name":   resource.NewStringProperty(name),
			"values": unorderedTestArray(values...),
		})
	}
	tests := []struct {
		name     string
		old      resource.PropertyMap
		baseline resource.PropertyMap
		expected resource.PropertyMap
	}{
		{
			name: "addition",
			old: resource.PropertyMap{"groups": unorderedTestObjectArray(
				item("first", "B", "A"),
			)},
			baseline: resource.PropertyMap{"groups": unorderedTestObjectArray(
				item("first", "A", "B"), item("second", "C"),
			)},
			expected: resource.PropertyMap{"groups": unorderedTestObjectArray(
				item("first", "B", "A"), item("second", "C"),
			)},
		},
		{
			name: "deletion",
			old: resource.PropertyMap{"groups": unorderedTestObjectArray(
				item("first", "B", "A"), item("second", "C"),
			)},
			baseline: resource.PropertyMap{"groups": unorderedTestObjectArray(
				item("first", "A", "B"),
			)},
			expected: resource.PropertyMap{"groups": unorderedTestObjectArray(
				item("first", "B", "A"),
			)},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SuppressBaselineDiffs("aws-native:test:Resource", &metadata.CloudAPIResource{
				UnorderedCollections: []string{"groups/*/values"},
			}, test.old, test.baseline, NewTransformCache())
			assert.True(t, test.expected.DeepEquals(result), "expected %v, got %v", test.expected, result)
		})
	}
}

func TestCalcPatchWithActualOutputsSuppressesUnorderedReorder(t *testing.T) {
	stringArrayType := pschema.TypeSpec{Type: arrayType, Items: &pschema.TypeSpec{Type: "string"}}
	spec := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"values": {TypeSpec: stringArrayType},
		},
		Outputs: map[string]pschema.PropertySpec{
			"values": {TypeSpec: stringArrayType},
		},
		UnorderedCollections: []string{"values"},
	}
	oldDesired := resource.PropertyMap{"values": unorderedTestArray("B", "A")}
	actualOutputs := resource.PropertyMap{"values": unorderedTestArray("A", "B")}

	patch, err := CalcPatchWithActualOutputs(
		oldDesired, actualOutputs, oldDesired, spec, nil, nil,
		"aws-native:test:Resource", NewTransformCache(),
	)
	require.NoError(t, err)
	assert.Empty(t, patch)

	changedDesired := resource.PropertyMap{"values": unorderedTestArray("B", "C")}
	patch, err = CalcPatchWithActualOutputs(
		oldDesired, actualOutputs, changedDesired, spec, nil, nil,
		"aws-native:test:Resource", NewTransformCache(),
	)
	require.NoError(t, err)
	require.Len(t, patch, 1)
	assert.Equal(t, "replace", patch[0].Operation)
	assert.Equal(t, "/Values", patch[0].Path)
	assert.Equal(t, []interface{}{"B", "C"}, patch[0].Value)

	writeOnlySpec := spec
	writeOnlySpec.Inputs = map[string]pschema.PropertySpec{
		"values":   {TypeSpec: stringArrayType},
		"password": {TypeSpec: pschema.TypeSpec{Type: "string"}},
	}
	writeOnlySpec.WriteOnly = []string{"password"}
	oldDesiredWithPassword := resource.PropertyMap{
		"values":   oldDesired["values"],
		"password": resource.NewStringProperty("secret"),
	}
	patch, err = CalcPatchWithActualOutputs(
		oldDesiredWithPassword, actualOutputs, oldDesiredWithPassword, writeOnlySpec, nil, nil,
		"aws-native:test:Resource", NewTransformCache(),
	)
	require.NoError(t, err)
	require.Len(t, patch, 1)
	assert.Equal(t, "add", patch[0].Operation)
	assert.Equal(t, "/Password", patch[0].Path)
	assert.Equal(t, "secret", patch[0].Value)
}

func suppressUnorderedTestDiff(
	t *testing.T,
	spec *metadata.CloudAPIResource,
	resourceToken string,
	oldInputs, newInputs resource.PropertyMap,
) *resource.ObjectDiff {
	t.Helper()
	diff := oldInputs.Diff(newInputs)
	require.NotNil(t, diff)
	return SuppressAWSManagedDiffsWithContext(
		resourceToken, spec, diff, oldInputs, oldInputs, newInputs, NewTransformCache(),
	)
}

func unorderedTestArray(values ...string) resource.PropertyValue {
	result := make([]resource.PropertyValue, len(values))
	for i, value := range values {
		result[i] = resource.NewStringProperty(value)
	}
	return resource.NewArrayProperty(result)
}

func unorderedTestObjectArray(values ...resource.PropertyValue) resource.PropertyValue {
	return resource.NewArrayProperty(values)
}

func unorderedTestObject(values ...string) resource.PropertyValue {
	if len(values)%2 != 0 {
		panic("unorderedTestObject requires key/value pairs")
	}
	result := resource.PropertyMap{}
	for len(values) > 0 {
		result[resource.PropertyKey(values[0])] = resource.NewStringProperty(values[1])
		values = values[2:]
	}
	return resource.NewObjectProperty(result)
}

func unorderedTestPolicy(name string, action resource.PropertyValue) resource.PropertyValue {
	return resource.NewObjectProperty(resource.PropertyMap{
		"policyName": resource.NewStringProperty(name),
		"policyDocument": resource.NewObjectProperty(resource.PropertyMap{
			"Version": resource.NewStringProperty("2012-10-17"),
			"Statement": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"Effect":   resource.NewStringProperty("Allow"),
					"Action":   action,
					"Resource": resource.NewStringProperty("*"),
				}),
			}),
		}),
	})
}
