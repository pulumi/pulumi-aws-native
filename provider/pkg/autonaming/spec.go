// Copyright 2024, Pulumi Corporation.

package autonaming

import (
	"strings"

	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func CreateAutoNamingSpec(inputProperties map[string]pschema.PropertySpec, resourceTypeName string, jsonSchemaProperties map[string]*jsschema.Schema, semanticsSpec metadata.SemanticsSpec) *metadata.AutoNamingSpec {
	type names struct {
		cfName  string
		sdkName string
	}
	// Lowercase all keys so we can do case-insensitive lookups.
	fieldsLowered := make(map[string]names, len(jsonSchemaProperties))
	for k := range jsonSchemaProperties {
		fieldsLowered[strings.ToLower(k)] = names{cfName: k}
	}
	for k := range inputProperties {
		lowered := strings.ToLower(k)
		if v, ok := fieldsLowered[lowered]; ok {
			v.sdkName = k
			fieldsLowered[lowered] = v
		} else {
			fieldsLowered[lowered] = names{sdkName: k}
		}
	}

	tryMakeSpecForField := func(loweredName string) (spec *metadata.AutoNamingSpec) {
		sdkName := fieldsLowered[loweredName].sdkName
		if prop, has := inputProperties[sdkName]; !has || prop.Type != "string" {
			return nil
		}
		jsonSpec := jsonSchemaProperties[fieldsLowered[loweredName].cfName]
		if jsonSpec == nil {
			return nil
		}
		return &metadata.AutoNamingSpec{
			SdkName:   sdkName,
			MinLength: jsonSpec.MinLength.Val,
			MaxLength: jsonSpec.MaxLength.Val,
		}
	}

	// 1. Look for a property named "Name"
	autoNameSpec := tryMakeSpecForField("name")
	if autoNameSpec == nil {
		// 2. Look for a property named "<ResourceTypeName>Name"
		autoNameSpec = tryMakeSpecForField(strings.ToLower(resourceTypeName) + "name")
	}
	if autoNameSpec == nil {
		// 3. Look for a property that ends with "*Name" where that prefix is also contained in the resource name.
		// E.g. ScalingPolicy has a field "PolicyName".
		var potentialSpecs []*metadata.AutoNamingSpec
		for k := range fieldsLowered {
			if before, ok := strings.CutSuffix(k, "name"); ok {
				if strings.HasSuffix(strings.ToLower(resourceTypeName), before) {
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
