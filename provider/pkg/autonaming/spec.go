// Copyright 2024, Pulumi Corporation.

package autonaming

import (
	"strings"

	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func CreateAutoNamingSpec(inputProperties map[string]pschema.PropertySpec, resourceTypeName string, jsonSchemaProperties map[string]*jsschema.Schema, semanticsSpec metadata.SemanticsSpec) *metadata.AutoNamingSpec {
	tryMakeSpecForField := func(cfPropName string) (spec *metadata.AutoNamingSpec) {
		if prop, has := inputProperties[naming.ToSdkName(cfPropName)]; !has || prop.Type != "string" {
			return nil
		}
		jsonSpec := jsonSchemaProperties[cfPropName]
		if jsonSpec == nil {
			return nil
		}
		return &metadata.AutoNamingSpec{
			SdkName:   naming.ToSdkName(cfPropName),
			MinLength: jsonSpec.MinLength.Val,
			MaxLength: jsonSpec.MaxLength.Val,
		}
	}

	// 1. Look for a property named "Name"
	autoNameSpec := tryMakeSpecForField("Name")
	if autoNameSpec == nil {
		// 2. Look for a property named "<ResourceTypeName>Name"
		autoNameSpec = tryMakeSpecForField(resourceTypeName + "Name")
	}
	if autoNameSpec == nil {
		// 3. Look for a property that ends with "*Name" where that prefix is also contained in the resource name.
		// E.g. ScalingPolicy has a field "PolicyName".
		var potentialSpecs []*metadata.AutoNamingSpec
		for k := range jsonSchemaProperties {
			if before, ok := strings.CutSuffix(k, "Name"); ok {
				if strings.HasSuffix(resourceTypeName, before) {
					potentialSpec := tryMakeSpecForField(k)
					if potentialSpec != nil {
						potentialSpecs = append(potentialSpecs, potentialSpec)
					}
				}
			}
		}
		// Only proceed if there is exactly one matched property name to avoid ambiguity.
		if len(potentialSpecs) == 1 {
			autoNameSpec = potentialSpecs[0]
		}
	}

	if autoNameSpec != nil {
		namingTriviaSpec, ok := semanticsSpec.NamingTriviaSpec[autoNameSpec.SdkName]
		if ok {
			autoNameSpec.TriviaSpec = &namingTriviaSpec
		}
	}

	return autoNameSpec
}
