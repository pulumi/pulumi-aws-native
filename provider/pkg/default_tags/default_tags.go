package default_tags

import (
	"strings"
)

type TagsStyle string

const (
	tagStyleKeyValueArrayPrefix = "keyValueArray"
	// TagsStyleUnknown indicates we can't identify the style of tags.
	TagsStyleUnknown TagsStyle = ""
	// TagsStyleUntyped indicates the resource has no tags
	TagsStyleNone TagsStyle = "none"
	// TagsStyleUntyped is a style where the tags are represented as "Any" - without a schema.
	TagsStyleUntyped TagsStyle = "untyped"
	// TagsStyleStringMap is a style where the tags are represented as a map of strings.
	TagsStyleStringMap TagsStyle = "stringMap"
	// TagsStyleKeyValueArray is a style where the tags are represented as an array of key-value pairs.
	TagsStyleKeyValueArray TagsStyle = tagStyleKeyValueArrayPrefix
	// TagsStyleKeyValueArrayCreateOnly is a style where the tags are represented as an array of key-value pairs, but
	// the tags are create-only.
	TagsStyleKeyValueArrayCreateOnly TagsStyle = tagStyleKeyValueArrayPrefix + "CreateOnly"
	// TagsStyleKeyValueArrayWithExtraProperties is a style where the tags are represented as an array of key-value pairs
	// but can have extra properties.
	TagsStyleKeyValueArrayWithExtraProperties TagsStyle = tagStyleKeyValueArrayPrefix + "WithExtraProperties"
	// TagsStyleKeyValueArrayWithAlternateType is a style where the tags are represented as an array of key-value pairs
	// but the value can also be of a different type.
	TagsStyleKeyValueArrayWithAlternateType TagsStyle = tagStyleKeyValueArrayPrefix + "WithAlternateType"
	// TagsStyleKeyValueArrayUpperCase is a style where the tags are represented as an array of key-value pairs but the "Key" and "Value" properties are uppercase.
	TagsStyleKeyValueArrayUpperCase TagsStyle = tagStyleKeyValueArrayPrefix + "UpperCase"
	// TagsStyleObjectArrayWithResourceType is a style where the tags are represented as an array of objects with a "resourceType" property
	// for example { tagSpecification: [ { resourceType: 'capacity-reservation' }, tags: [ { key: 'Key', value: 'Value' }]]}
	TagsStyleNestedKeyValueArrayWithResourceType TagsStyle = "nestedKeyValueArrayWithResourceType"
)

func (ts TagsStyle) IsStringMap() bool {
	return ts == TagsStyleStringMap
}

func (ts TagsStyle) IsKeyValueArray() bool {
	return strings.HasPrefix(string(ts), tagStyleKeyValueArrayPrefix)
}

func (ts TagsStyle) IsValid() bool {
	return ts == TagsStyleUnknown || ts == TagsStyleNone || ts == TagsStyleUntyped || ts.IsStringMap() || ts.IsKeyValueArray()
}
