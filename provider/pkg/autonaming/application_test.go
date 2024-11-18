// Copyright 2016-2021, Pulumi Corporation.

package autonaming

import (
	"fmt"
	"testing"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_getDefaultName(t *testing.T) {
	const sdkName = "autoName"
	tests := []struct {
		name       string
		minLength  int
		maxLength  int
		olds       resource.PropertyMap
		news       resource.PropertyMap
		err        error
		comparison func(t *testing.T, actual resource.PropertyValue) bool
	}{
		{
			name: "Name specified explicitly",
			news: resource.PropertyMap{
				resource.PropertyKey(sdkName): resource.NewStringProperty("newName"),
			},
			comparison: equals(resource.NewStringProperty("newName")),
		},
		{
			name: "Use old name",
			olds: resource.PropertyMap{
				resource.PropertyKey(sdkName): resource.NewStringProperty("oldName"),
			},
			comparison: equals(resource.NewStringProperty("oldName")),
		},
		{
			name:       "Autoname with defaults",
			comparison: within(14, 14),
		},
		{
			name:       "Autoname with constraints on max length",
			maxLength:  12,
			comparison: within(12, 12),
		},
		{
			name:       "Autoname with constraints on min length",
			minLength:  15,
			comparison: within(15, 15),
		},
		{
			name:       "Autoname with max length too small",
			maxLength:  6,
			comparison: within(15, 15),
			err:        fmt.Errorf("failed to auto-generate value for %[1]q. Prefix: \"myName-\" is too large to fix max length constraint of 6. Please provide a value for %[1]q", sdkName),
		},
		{
			name:       "Autoname with constraints on min and max length",
			minLength:  13,
			maxLength:  13,
			comparison: within(13, 13),
		},
	}

	urn := resource.URN("urn:pulumi:dev::test::test-provider:testModule:TestResource::myName")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			autoNamingSpec := &metadata.AutoNamingSpec{
				SdkName:   "autoName",
				MinLength: tt.minLength,
				MaxLength: tt.maxLength,
			}
			got, err := getDefaultName(nil, urn, autoNamingSpec, tt.olds, tt.news, nil)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
				return
			} else {
				require.NoError(t, err)
			}
			if !tt.comparison(t, got) {
				t.Errorf("getDefaultName() = %v for spec: %+v", got, autoNamingSpec)
			}
			t.Logf("getDefaultName() = %v for spec: %+v", got, autoNamingSpec)
		})
	}
}

func Test_getDefaultName_withAutoNameConfig(t *testing.T) {
	const sdkName = "autoName"
	tests := []struct {
		name           string
		resourceName   string
		autoNameConfig AutoNamingConfig
		minLength      int
		maxLength      int
		olds           resource.PropertyMap
		news           resource.PropertyMap
		err            error
		comparison     func(t *testing.T, actual resource.PropertyValue) bool
	}{
		{
			name:         "Name specified explicitly",
			resourceName: "myName",
			news: resource.PropertyMap{
				resource.PropertyKey(sdkName): resource.NewStringProperty("newName"),
			},
			autoNameConfig: AutoNamingConfig{
				AutoTrim:              true,
				RandomSuffixMinLength: 1,
			},
			maxLength:  2,
			comparison: equals(resource.NewStringProperty("newName")),
		},
		{
			name:         "Use old name",
			resourceName: "myName",
			olds: resource.PropertyMap{
				resource.PropertyKey(sdkName): resource.NewStringProperty("oldName"),
			},
			autoNameConfig: AutoNamingConfig{
				AutoTrim:              true,
				RandomSuffixMinLength: 1,
			},
			maxLength:  2,
			comparison: equals(resource.NewStringProperty("oldName")),
		},
		{
			resourceName: "myReallyLongName",
			name:         "Autoname with constraints on max length",
			maxLength:    12,
			autoNameConfig: AutoNamingConfig{
				AutoTrim: true,
			},
			comparison: startsWith("myReagName", 1),
		},
		{
			resourceName: "myReallyLongName",
			name:         "Autoname with odd constraints on max length",
			maxLength:    13,
			autoNameConfig: AutoNamingConfig{
				AutoTrim: true,
			},
			comparison: startsWith("myRealgName", 1),
		},
		{
			name:         "Autoname with max length too small and no auto trim",
			resourceName: "myName",
			maxLength:    4,
			autoNameConfig: AutoNamingConfig{
				RandomSuffixMinLength: 2,
			},
			comparison: within(15, 15),
			err:        fmt.Errorf("failed to auto-generate value for %[1]q. Prefix: \"myName-\" is too large to fix max length constraint of 4. Please provide a value for %[1]q", sdkName),
		},
		{
			name:         "Autoname with constraints on min and max length",
			minLength:    13,
			resourceName: "myReallyLongName",
			autoNameConfig: AutoNamingConfig{
				RandomSuffixMinLength: 2,
				AutoTrim:              true,
			},
			maxLength:  13,
			comparison: startsWith("myReagName", 2),
		},
		{
			name:         "Autoname with long random suffix",
			resourceName: "myReallyLongName",
			autoNameConfig: AutoNamingConfig{
				RandomSuffixMinLength: 14,
				AutoTrim:              true,
			},
			maxLength:  13,
			comparison: within(13, 13),
			err:        fmt.Errorf("failed to auto-generate value for %[1]q. Prefix: \"myReallyLongName-\" is too large to fix max length constraint of 13 with required suffix length 14. Please provide a value for %[1]q", sdkName),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urn := resource.URN(fmt.Sprintf("urn:pulumi:dev::test::test-provider:testModule:TestResource::%s", tt.resourceName))
			autoNamingSpec := &metadata.AutoNamingSpec{
				SdkName:   "autoName",
				MinLength: tt.minLength,
				MaxLength: tt.maxLength,
			}

			autoNameConfig := tt.autoNameConfig
			got, err := getDefaultName(nil, urn, autoNamingSpec, tt.olds, tt.news, &autoNameConfig)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
				return
			} else {
				require.NoError(t, err)
			}
			if !tt.comparison(t, got) {
				t.Errorf("getDefaultName() = %v for spec: %+v", got, autoNamingSpec)
			}
			t.Logf("getDefaultName() = %v for spec: %+v", got, autoNamingSpec)
		})
	}
}

func equals(expected resource.PropertyValue) func(t *testing.T, actual resource.PropertyValue) bool {
	return func(t *testing.T, actual resource.PropertyValue) bool {
		return expected == actual
	}
}

func startsWith(prefix string, randomLen int) func(t *testing.T, actual resource.PropertyValue) bool {
	return func(t *testing.T, actual resource.PropertyValue) bool {
		return assert.Regexp(t, fmt.Sprintf("^%s-[a-z0-9]{%d}", prefix, randomLen), actual.StringValue())
	}
}

func within(min, max int) func(t *testing.T, value resource.PropertyValue) bool {
	return func(t *testing.T, actual resource.PropertyValue) bool {
		l := len(actual.V.(string))
		return min <= l && l <= max
	}
}

func Test_trimName(t *testing.T) {
	assert.Equal(t, "mysupgname", trimName("mysuperlongname", 10))
	assert.Equal(t, "mysupernam", trimName("mysupernam", 10))
}
