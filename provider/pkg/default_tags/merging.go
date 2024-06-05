package default_tags

import (
	"fmt"
	"sort"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func MergeDefaultTags(tags resource.PropertyValue, defaultTags map[string]string, tagsStyle TagsStyle) (resource.PropertyValue, error) {
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
		keyProp := resource.PropertyKey("key")
		valueProp := resource.PropertyKey("value")
		if tagsStyle == TagsStyleKeyValueArrayUpperCase {
			keyProp = resource.PropertyKey("Key")
			valueProp = resource.PropertyKey("Value")
		}

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
			if !asObject.HasValue(keyProp) || !asObject.HasValue(valueProp) {
				return tags, fmt.Errorf("expected tags to be an array of objects with '%s' and '%s' properties but found %v", keyProp, valueProp, tag.TypeString())
			}
			key := asObject[keyProp]
			if key.IsString() {
				existingKeys[key.StringValue()] = true
			}
		}

		// Get the sorted list of keys from defaultTags
		var keys []string
		for k := range defaultTags {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			if _, ok := existingKeys[key]; !ok {
				asArray = append(asArray, resource.NewObjectProperty(map[resource.PropertyKey]resource.PropertyValue{
					keyProp:   resource.NewStringProperty(key),
					valueProp: resource.NewStringProperty(defaultTags[key]),
				}))
			}
		}
		return resource.PropertyValue{V: asArray}, nil
	}
	return tags, nil
}
