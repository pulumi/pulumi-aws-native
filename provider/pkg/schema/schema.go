// Copyright 2016-2021, Pulumi Corporation.

package schema

type PropertySpec struct {
	Documentation     string `json:"Documentation"`
	DuplicatesAllowed bool   `json:"DuplicatesAllowed"`
	ItemType          string `json:"ItemType,omitempty"`
	PrimitiveItemType string `json:"PrimitiveItemType,omitempty"`
	PrimitiveType     string `json:"PrimitiveType,omitempty"`
	Required          bool   `json:"Required"`
	Type              string `json:"Type,omitempty"`
	UpdateType        string `json:"UpdateType"`
}

type PropertyTypeSpec struct {
	Documentation string                  `json:"Documentation"`
	Properties    map[string]PropertySpec `json:"Properties"`
}

type AttributeSpec struct {
	ItemType          string `json:"ItemType,omitempty"`
	PrimitiveItemType string `json:"PrimitiveItemType,omitempty"`
	PrimitiveType     string `json:"PrimitiveType,omitempty"`
	Type              string `json:"Type,omitempty"`
}

type ResourceSpec struct {
	Attributes map[string]AttributeSpec `json:"Attributes"`
	PropertyTypeSpec
}

type CloudFormationSchema struct {
	PropertyTypes                map[string]PropertyTypeSpec `json:"PropertyTypes"`
	ResourceSpecificationVersion string                      `json:"ResourceSpecificationVersion"`
	ResourceTypes                map[string]ResourceSpec     `json:"ResourceTypes"`
}
