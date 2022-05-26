// Copyright 2016-2021, Pulumi Corporation.

package provider

import (
	"fmt"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/require"
	"testing"
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
		comparison func(actual resource.PropertyValue) bool
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
			autoNamingSpec := &schema.AutoNamingSpec{
				SdkName:   "autoName",
				MinLength: tt.minLength,
				MaxLength: tt.maxLength,
			}
			got, err := getDefaultName(urn, autoNamingSpec, tt.olds, tt.news)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
				return
			} else {
				require.NoError(t, err)
			}
			if !tt.comparison(got) {
				t.Errorf("getDefaultName() = %v for spec: %+v", got, autoNamingSpec)
			}
			t.Logf("getDefaultName() = %v for spec: %+v", got, autoNamingSpec)
		})
	}
}

func equals(expected resource.PropertyValue) func(resource.PropertyValue) bool {
	return func(actual resource.PropertyValue) bool {
		return expected == actual
	}
}

func within(min, max int) func(value resource.PropertyValue) bool {
	return func(actual resource.PropertyValue) bool {
		l := len(actual.V.(string))
		return min <= l && l <= max
	}
}
