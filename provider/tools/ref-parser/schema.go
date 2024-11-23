package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type schema struct {
	TypeName          string         `json:"typeName"`
	Properties        map[string]any `json:"properties"`
	PrimaryIdentifier []string       `json:"primaryIdentifier"`
}

// Read CF schema and find which properties are available for a resource.
func findSchema(schemaAbsPath, resourceID string) (*schema, error) {
	fn := strings.ToLower(strings.ReplaceAll(resourceID, "::", "-") + ".json")
	fn = strings.ReplaceAll(fn, "`", "")
	file := filepath.Join(schemaAbsPath, fn)
	fbytes, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("No schema file for %q: %w", resourceID, err)
	}
	var s schema
	if err := json.Unmarshal(fbytes, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// Read CF schema and find which properties are available for a resource.
func findProperties(schemaAbsPath, resourceID string) ([]string, error) {
	s, err := findSchema(schemaAbsPath, resourceID)
	if err != nil {
		return nil, err
	}
	var res []string
	for x := range s.Properties {
		res = append(res, x)
	}
	sort.Strings(res)
	return res, nil
}
