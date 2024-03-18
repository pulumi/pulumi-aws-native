package provider

import "github.com/pulumi/pulumi-aws-native/provider/pkg/schema"

type converter interface {
	Serialize(inputsMap map[string]any) (cfType string, payload map[string]interface{})
	Deserialize(inputsMap, resourceState map[string]any) (outputsMap map[string]interface{})
}

func (p *cfnProvider) MakeConverter(resourceToken string) converter {
	switch resourceToken {
	case schema.ExtensionResourceToken:
		return &extensionConverter{}
	default:
		return &sdkConverter{
			resourceToken: resourceToken,
			resourceMap:   p.resourceMap,
		}
	}
}

type sdkConverter struct {
	resourceToken string
	resourceMap   *schema.CloudAPIMetadata
}

func (c *sdkConverter) Serialize(inputsMap map[string]any) (cfType string, payload map[string]interface{}) {
	spec, _ := c.resourceMap.Resources[c.resourceToken]
	payload, _ = schema.SdkToCfn(&spec, c.resourceMap.Types, inputsMap)
	return spec.CfType, payload
}

func (c *sdkConverter) Deserialize(inputsMap, resourceState map[string]any) (outputsMap map[string]interface{}) {
	spec, _ := c.resourceMap.Resources[c.resourceToken]
	outputs := schema.CfnToSdk(resourceState)

	// Write-only properties are not returned in the outputs, so we assume they should have the same value we sent from the inputs.
	if len(spec.WriteOnly) > 0 {
		for _, writeOnlyProp := range spec.WriteOnly {
			if _, ok := outputs[writeOnlyProp]; !ok {
				inputValue, ok := inputsMap[writeOnlyProp]
				if ok {
					outputs[writeOnlyProp] = inputValue
				}
			}
		}
	}
	return outputs
}

type extensionConverter struct {
}

func (c *extensionConverter) Serialize(inputsMap map[string]any) (cfType string, payload map[string]interface{}) {
	return inputsMap["type"].(string), inputsMap["properties"].(map[string]interface{})
}

func (c *extensionConverter) Deserialize(inputsMap, resourceState map[string]any) (outputsMap map[string]interface{}) {
	return map[string]interface{}{
		"outputs": resourceState,
	}
}
