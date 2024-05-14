package schema

import (
	"encoding/json"
	"fmt"
	"os"
)

type Docs struct {
	Types map[string]DocsTypes `json:"Types"`
}

type DocsTypes struct {
	Properties  DocsProperties    `json:"properties"`
	Attributes  map[string]string `json:"attributes"`
	Description string            `json:"description"`
}

type DocsProperties map[string]string

func ReadCloudFormationDocsFile(filePath string) (*Docs, error) {
	docs := Docs{}
	in, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer in.Close()

	dec := json.NewDecoder(in)
	if err := dec.Decode(&docs); err != nil {
		return nil, fmt.Errorf("error decoding: %v", err)
	}
	return &docs, nil
}

func (d *Docs) GetPropertyDesc(typeName, propName string) (string, bool) {
	if typ, exists := d.Types[typeName]; exists {
		if propName == "" {
			return typ.Description, true
		}
		if prop, ok := typ.Properties[propName]; ok {
			return prop, true
		}
		if prop, ok := typ.Attributes[propName]; ok {
			return prop, true
		}
	}
	return "", false
}
