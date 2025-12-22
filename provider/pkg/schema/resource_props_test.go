package schema

import (
	"testing"

	jsschema "github.com/pulumi/jsschema"
)

func TestToSdkName(t *testing.T) {
	cases := map[string]struct {
		input    ResourceProperty
		expected string
	}{
		"simple": {
			input:    "Foo",
			expected: "foo",
		},
		"nestedObject": {
			input:    "Foo/BarBaz",
			expected: "foo/barBaz",
		},
		"nestedArray": {
			input:    "Foo/*/BarBaz",
			expected: "foo/*/barBaz",
		},
		"nestedArrayObject": {
			input:    "Foo/*/Bar/Baz",
			expected: "foo/*/bar/baz",
		},
		"multiNestedArrayObject": {
			input:    "Foo/*/Bar/*/Baz/Hello",
			expected: "foo/*/bar/*/baz/hello",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := tc.input.ToSdkName()
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestCfnPathToSdkPath(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected string
	}{
		"simple property": {
			input:    "/properties/Engine",
			expected: "engine",
		},
		"nested property": {
			input:    "/properties/FileSystemProtection/ReplicationOverwriteProtection",
			expected: "fileSystemProtection/replicationOverwriteProtection",
		},
		"array wildcard": {
			input:    "/properties/SecurityGroupEgress/*/IpProtocol",
			expected: "securityGroupEgress/*/ipProtocol",
		},
		"complex nested array": {
			input:    "/properties/ReplicationConfiguration/Destinations/*/FileSystemId",
			expected: "replicationConfiguration/destinations/*/fileSystemId",
		},
		"no prefix returns empty": {
			input:    "Engine",
			expected: "",
		},
		"wrong prefix returns empty": {
			input:    "/definitions/Engine",
			expected: "",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cfnPathToSdkPath(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestReadPropertyTransforms(t *testing.T) {
	t.Run("returns nil for nil schema", func(t *testing.T) {
		result := readPropertyTransforms(nil)
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})

	t.Run("returns nil for schema without propertyTransform", func(t *testing.T) {
		schema := &jsschema.Schema{
			Extras: map[string]interface{}{},
		}
		result := readPropertyTransforms(schema)
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})

	t.Run("extracts and converts transforms", func(t *testing.T) {
		schema := &jsschema.Schema{
			Extras: map[string]interface{}{
				"propertyTransform": map[string]interface{}{
					"/properties/Engine":                      "$lowercase(Engine)",
					"/properties/SecurityGroupEgress/*/IpProtocol": "$lowercase(IpProtocol)",
				},
			},
		}

		result := readPropertyTransforms(schema)
		if result == nil {
			t.Fatal("expected non-nil result")
		}

		if len(result) != 2 {
			t.Errorf("expected 2 transforms, got %d", len(result))
		}

		if result["engine"] != "$lowercase(Engine)" {
			t.Errorf("expected engine transform, got %v", result["engine"])
		}

		if result["securityGroupEgress/*/ipProtocol"] != "$lowercase(IpProtocol)" {
			t.Errorf("expected securityGroupEgress transform, got %v", result["securityGroupEgress/*/ipProtocol"])
		}
	})

	t.Run("skips invalid paths", func(t *testing.T) {
		schema := &jsschema.Schema{
			Extras: map[string]interface{}{
				"propertyTransform": map[string]interface{}{
					"/properties/Valid":   "$lowercase(Valid)",
					"/definitions/Invalid": "$lowercase(Invalid)", // Invalid path
				},
			},
		}

		result := readPropertyTransforms(schema)
		if result == nil {
			t.Fatal("expected non-nil result")
		}

		if len(result) != 1 {
			t.Errorf("expected 1 transform (invalid path should be skipped), got %d", len(result))
		}
	})
}
