package resources

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestCheckpointObject(t *testing.T) {
	t.Parallel()

	inputs := resource.PropertyMap{
		"input1": resource.NewStringProperty("value1"),
		"input2": resource.NewNumberProperty(42),
	}

	outputs := map[string]interface{}{
		"output1": "value1",
		"output2": 42,
	}

	result := CheckpointObject(inputs, outputs)

	// Check if outputs are correctly set
	assert.Equal(t, resource.NewStringProperty("value1"), result["output1"])
	assert.Equal(t, resource.NewNumberProperty(42), result["output2"])

	// Check if __inputs field is correctly set and is a secret
	inputsField, ok := result["__inputs"]
	assert.True(t, ok)
	assert.True(t, inputsField.IsSecret())

	// Check if the secret value contains the correct inputs
	secretInputs := inputsField.SecretValue().Element.ObjectValue()
	assert.Equal(t, inputs, secretInputs)
}

func TestParseCheckpointObject(t *testing.T) {
	t.Parallel()

	inputs := resource.PropertyMap{
		"input1": resource.NewStringProperty("value1"),
		"input2": resource.NewNumberProperty(42),
	}

	// Create a PropertyMap with __inputs as a secret
	obj := resource.PropertyMap{
		"__inputs": resource.MakeSecret(resource.NewObjectProperty(inputs)),
	}

	// Parse the checkpoint object
	parsedInputs := ParseCheckpointObject(obj)

	// Check if the parsed inputs match the original inputs
	assert.Equal(t, inputs, parsedInputs)

	// Test with an object that does not contain __inputs
	objWithoutInputs := resource.PropertyMap{
		"output1": resource.NewStringProperty("value1"),
		"output2": resource.NewNumberProperty(42),
	}

	// Parse the checkpoint object
	parsedInputs = ParseCheckpointObject(objWithoutInputs)

	// Check if the parsed inputs are nil
	assert.Nil(t, parsedInputs)
}

func TestRoundTripCheckpointObject(t *testing.T) {
	t.Parallel()

	inputs := resource.PropertyMap{
		"input1": resource.NewStringProperty("value1"),
		"input2": resource.NewNumberProperty(42),
	}

	outputs := map[string]interface{}{
		"output1": "value1",
		"output2": 42,
	}

	// Create a checkpoint object
	checkpoint := CheckpointObject(inputs, outputs)

	// Parse the checkpoint object
	parsedInputs := ParseCheckpointObject(checkpoint)

	// Check if the parsed inputs match the original inputs
	assert.Equal(t, inputs, parsedInputs)

	// Check if the outputs are still correctly set
	assert.Equal(t, resource.NewStringProperty("value1"), checkpoint["output1"])
	assert.Equal(t, resource.NewNumberProperty(42), checkpoint["output2"])
}

func TestCheckpointPropertyMap(t *testing.T) {
	t.Parallel()

	inputs := resource.PropertyMap{
		"input1": resource.NewStringProperty("value1"),
		"input2": resource.NewNumberProperty(42),
	}

	outputs := resource.PropertyMap{
		"output1": resource.NewStringProperty("value1"),
		"output2": resource.NewNumberProperty(42),
	}

	result := CheckpointPropertyMap(inputs, outputs)

	// Check if outputs are correctly set
	assert.Equal(t, resource.NewStringProperty("value1"), result["output1"])
	assert.Equal(t, resource.NewNumberProperty(42), result["output2"])

	// Check if __inputs field is correctly set and is a secret
	inputsField, ok := result["__inputs"]
	assert.True(t, ok)
	assert.True(t, inputsField.IsSecret())

	// Check if the secret value contains the correct inputs
	secretInputs := inputsField.SecretValue().Element.ObjectValue()
	assert.Equal(t, inputs, secretInputs)
}

func TestCheckpointPropertyMapWithNilOutputs(t *testing.T) {
	t.Parallel()

	inputs := resource.PropertyMap{
		"input1": resource.NewStringProperty("value1"),
		"input2": resource.NewNumberProperty(42),
	}

	result := CheckpointPropertyMap(inputs, nil)

	// Check if __inputs field is correctly set and is a secret
	inputsField, ok := result["__inputs"]
	assert.True(t, ok)
	assert.True(t, inputsField.IsSecret())

	// Check if the secret value contains the correct inputs
	secretInputs := inputsField.SecretValue().Element.ObjectValue()
	assert.Equal(t, inputs, secretInputs)
}
