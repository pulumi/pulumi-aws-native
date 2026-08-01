package resources

import (
	"testing"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
)

// consecutiveSecretDepth counts consecutive secret wrappers around a value.
func consecutiveSecretDepth(v resource.PropertyValue) int {
	depth := 0
	for v.IsSecret() {
		depth++
		v = v.SecretValue().Element
	}
	return depth
}

// canarySpec models AWS::Synthetics::Canary, whose runConfig/environmentVariables
// is write-only: it is accepted on create/update but never returned by
// CloudControl, so refresh must restore it from checkpointed old inputs.
func canarySpec() (metadata.CloudAPIResource, map[string]metadata.CloudAPIType) {
	spec := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"name":      {TypeSpec: pschema.TypeSpec{Type: "string"}},
			"runConfig": {TypeSpec: pschema.TypeSpec{Ref: "#/types/aws-native:test:RunConfig"}},
		},
		Outputs: map[string]pschema.PropertySpec{
			"name":      {TypeSpec: pschema.TypeSpec{Type: "string"}},
			"runConfig": {TypeSpec: pschema.TypeSpec{Ref: "#/types/aws-native:test:RunConfig"}},
		},
		WriteOnly: []string{"runConfig/environmentVariables"},
	}
	types := map[string]metadata.CloudAPIType{
		"aws-native:test:RunConfig": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"timeoutInSeconds": {TypeSpec: pschema.TypeSpec{Type: "integer"}},
				"environmentVariables": {TypeSpec: pschema.TypeSpec{
					Type:                 "object",
					AdditionalProperties: &pschema.TypeSpec{Type: "string"},
				}},
			},
		},
	}
	return spec, types
}

// simulateRefreshRead mirrors the standard-resource refresh path of
// cfnProvider.Read for a resource with a secret value under a write-only path.
func simulateRefreshRead(t *testing.T, state resource.PropertyMap) resource.PropertyMap {
	spec, types := canarySpec()
	classifier := NewPathClassifier(&spec, types)

	inputs := ParseCheckpointObject(state)
	require.NotNil(t, inputs)

	// CloudControl never returns write-only properties.
	newStateProps := resource.PropertyMap{
		"name": resource.NewStringProperty("canary"),
		"runConfig": resource.NewObjectProperty(resource.PropertyMap{
			"timeoutInSeconds": resource.NewNumberProperty(60),
		}),
	}
	classifier.AddWriteOnlyOutputFallbacks(newStateProps, inputs)
	PreserveSecretWrappers(newStateProps, inputs)
	baseline := classifier.ActualInputBaselineFromOutputs(inputs, newStateProps, inputs)
	newInputs := SuppressBaselineDiffs("aws-native:test:Canary", &spec, inputs, baseline, NewTransformCache())
	return CheckpointPropertyMap(newInputs, newStateProps)
}

// TestRefreshDoesNotNestWriteOnlySecrets asserts the invariant that repeated
// refreshes keep a secret write-only value at exactly one secret wrapper in
// both output state and checkpointed __inputs. Before the fix, the wrapper
// depth doubled on every refresh (1, 2, 4, 8, 16, ...), and JSON escaping made
// the encrypted state entry roughly double per level.
func TestRefreshDoesNotNestWriteOnlySecrets(t *testing.T) {
	t.Parallel()

	inputs := resource.PropertyMap{
		"name": resource.NewStringProperty("canary"),
		"runConfig": resource.NewObjectProperty(resource.PropertyMap{
			"timeoutInSeconds": resource.NewNumberProperty(60),
			"environmentVariables": resource.NewObjectProperty(resource.PropertyMap{
				"API_KEY": resource.MakeSecret(resource.NewStringProperty("fake-36-character-secret-value-here")),
			}),
		}),
	}
	state := CheckpointPropertyMap(inputs, inputs.Copy())

	for i := 0; i < 20; i++ {
		state = simulateRefreshRead(t, state)

		outLeaf, ok := GetPath(state, "runConfig/environmentVariables/API_KEY")
		require.Truef(t, ok, "iteration %d: output leaf missing", i)
		require.Equalf(t, 1, consecutiveSecretDepth(outLeaf), "iteration %d: output secret depth", i)
		assert.Equal(t, resource.NewStringProperty("fake-36-character-secret-value-here"), outLeaf.SecretValue().Element)

		inLeaf, ok := GetPath(ParseCheckpointObject(state), "runConfig/environmentVariables/API_KEY")
		require.Truef(t, ok, "iteration %d: __inputs leaf missing", i)
		require.Equalf(t, 1, consecutiveSecretDepth(inLeaf), "iteration %d: __inputs secret depth", i)
	}
}

// TestRefreshHealsNestedSecretState asserts that a state poisoned by earlier
// versions, with a deeply nested secret under a write-only path, collapses back
// to a single wrapper on the next refresh instead of doubling further.
func TestRefreshHealsNestedSecretState(t *testing.T) {
	t.Parallel()

	deep := resource.NewStringProperty("fake-36-character-secret-value-here")
	for i := 0; i < 16; i++ {
		deep = resource.MakeSecret(deep)
	}
	poisoned := resource.PropertyMap{
		"name": resource.NewStringProperty("canary"),
		"runConfig": resource.NewObjectProperty(resource.PropertyMap{
			"timeoutInSeconds": resource.NewNumberProperty(60),
			"environmentVariables": resource.NewObjectProperty(resource.PropertyMap{
				"API_KEY": deep,
			}),
		}),
	}
	state := resource.PropertyMap{
		"name": resource.NewStringProperty("canary"),
		"runConfig": resource.NewObjectProperty(resource.PropertyMap{
			"timeoutInSeconds": resource.NewNumberProperty(60),
			"environmentVariables": resource.NewObjectProperty(resource.PropertyMap{
				"API_KEY": deep,
			}),
		}),
		"__inputs": resource.MakeSecret(resource.NewObjectProperty(poisoned)),
	}

	state = simulateRefreshRead(t, state)

	outLeaf, ok := GetPath(state, "runConfig/environmentVariables/API_KEY")
	require.True(t, ok)
	assert.Equal(t, 1, consecutiveSecretDepth(outLeaf))
	assert.Equal(t, resource.NewStringProperty("fake-36-character-secret-value-here"), outLeaf.SecretValue().Element)

	inLeaf, ok := GetPath(ParseCheckpointObject(state), "runConfig/environmentVariables/API_KEY")
	require.True(t, ok)
	assert.Equal(t, 1, consecutiveSecretDepth(inLeaf))
}
