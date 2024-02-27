package provider

import (
	"fmt"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func mergeDefaultTags(tags resource.PropertyValue, defaultTags map[string]string, tagsStyle schema.TagsStyle) (resource.PropertyValue, error) {
	if len(defaultTags) == 0 {
		return tags, nil
	}
	if tags.IsComputed() {
		return tags, nil
	}

	if tagsStyle.IsStringMap() {
		if tags.IsNull() {
			return resource.NewPropertyValue(defaultTags), nil
		}
		if !tags.IsObject() {
			return resource.NewObjectProperty(nil), fmt.Errorf("expected tags to be an object but found %v", tags.TypeString())
		}
		asObject := tags.ObjectValue()
		for k, v := range defaultTags {
			key := resource.PropertyKey(k)
			if !asObject.HasValue(key) {
				asObject[key] = resource.NewPropertyValue(v)
			}
		}
		return resource.PropertyValue{V: asObject}, nil
	}

	if tagsStyle.IsKeyValueArray() {
		var asArray []resource.PropertyValue

		if tags.IsArray() {
			asArray = tags.ArrayValue()
		} else if !tags.IsNull() {
			return resource.NewArrayProperty(nil), fmt.Errorf("expected tags to be an array but found %v", tags.TypeString())
		}
		existingKeys := make(map[string]bool)
		for _, tag := range asArray {
			if !tag.IsObject() {
				return tags, fmt.Errorf("expected tags to be an array of objects but found %v", tag.TypeString())
			}
			asObject := tag.ObjectValue()
			if !asObject.HasValue("key") || !asObject.HasValue("value") {
				return tags, fmt.Errorf("expected tags to be an array of objects with 'key' and 'value' properties but found %v", tag.TypeString())
			}
			key := asObject["key"]
			if key.IsString() {
				existingKeys[key.StringValue()] = true
			}
		}
		for k, v := range defaultTags {
			if _, ok := existingKeys[k]; !ok {
				asArray = append(asArray, resource.NewObjectProperty(map[resource.PropertyKey]resource.PropertyValue{
					"key":   resource.NewStringProperty(k),
					"value": resource.NewStringProperty(v),
				}))
			}
		}
		return resource.PropertyValue{V: asArray}, nil
	}
	return tags, nil
}
