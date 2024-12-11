package schema

import "testing"

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
