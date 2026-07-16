// Copyright 2026, Pulumi Corporation.

package resources

import (
	"encoding/base64"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

const (
	awsNativeWafv2WebAclToken    = "aws-native:wafv2:WebAcl"    //nolint:gosec // Resource type token, not a credential.
	awsNativeWafv2RuleGroupToken = "aws-native:wafv2:RuleGroup" //nolint:gosec // Resource type token, not a credential.
	wafv2RulesProperty           = "rules"
	wafv2StatementProperty       = "statement"
	wafv2ByteMatchProperty       = "byteMatchStatement"
	wafv2SearchStringProperty    = "searchString"
	wafv2SearchStringBase64Prop  = "searchStringBase64"
)

// NormalizeWafv2ByteMatchInputs converts the user-friendly searchString form
// into searchStringBase64 for WAFv2 resources.
//
// CloudControl reads ByteMatchStatement values back with both properties set,
// even though its update model requires exactly one. Sending the canonical
// base64 form avoids CloudFormation adding the alias again while validating an
// update. If a user explicitly specifies both properties, this function leaves
// them unchanged so normal validation can report the invalid input.
func NormalizeWafv2ByteMatchInputs(resourceToken string, inputs resource.PropertyMap) {
	if !isWafv2RuleContainer(resourceToken) {
		return
	}
	normalizeWafv2ByteMatchInputValue(resource.NewObjectProperty(inputs))
}

// NormalizeWafv2ByteMatchBaseline removes the redundant alias that
// CloudControl returns from an actual-state baseline. The desired input decides
// which representation to retain, so refresh preserves the user's existing
// representation while checked program inputs can migrate to the canonical
// base64 form without perpetual diffs.
func NormalizeWafv2ByteMatchBaseline(
	resourceToken string,
	actual resource.PropertyMap,
	desired resource.PropertyMap,
) {
	if !isWafv2RuleContainer(resourceToken) {
		return
	}
	normalizeWafv2ByteMatchBaselineValue(
		resource.NewObjectProperty(actual),
		resource.NewObjectProperty(desired),
	)
}

func isWafv2RuleContainer(resourceToken string) bool {
	return resourceToken == awsNativeWafv2WebAclToken || resourceToken == awsNativeWafv2RuleGroupToken
}

func normalizeWafv2ByteMatchInputValue(value resource.PropertyValue) {
	switch {
	case value.IsSecret():
		normalizeWafv2ByteMatchInputValue(value.SecretValue().Element)
	case value.IsArray():
		for _, item := range value.ArrayValue() {
			normalizeWafv2ByteMatchInputValue(item)
		}
	case value.IsObject():
		obj := value.ObjectValue()
		searchString, hasSearchString := obj[wafv2SearchStringProperty]
		_, hasSearchStringBase64 := obj[wafv2SearchStringBase64Prop]
		if hasSearchString && !hasSearchStringBase64 {
			if encoded, ok := base64PropertyValue(searchString); ok {
				obj[wafv2SearchStringBase64Prop] = encoded
				delete(obj, wafv2SearchStringProperty)
			}
		}
		for _, child := range obj {
			normalizeWafv2ByteMatchInputValue(child)
		}
	}
}

func normalizeWafv2ByteMatchBaselineValue(actual, desired resource.PropertyValue) {
	if actual.IsSecret() {
		actual = actual.SecretValue().Element
	}
	if desired.IsSecret() {
		desired = desired.SecretValue().Element
	}

	switch {
	case actual.IsArray():
		actualItems := actual.ArrayValue()
		var desiredItems []resource.PropertyValue
		if desired.IsArray() {
			desiredItems = desired.ArrayValue()
		}
		for i, item := range actualItems {
			desiredItem := resource.PropertyValue{}
			if i < len(desiredItems) {
				desiredItem = desiredItems[i]
			}
			normalizeWafv2ByteMatchBaselineValue(item, desiredItem)
		}
	case actual.IsObject():
		actualObj := actual.ObjectValue()
		desiredObj := resource.PropertyMap{}
		if desired.IsObject() {
			desiredObj = desired.ObjectValue()
		}

		searchString, hasSearchString := actualObj[wafv2SearchStringProperty]
		searchStringBase64, hasSearchStringBase64 := actualObj[wafv2SearchStringBase64Prop]
		if hasSearchString && hasSearchStringBase64 && base64ValuesEquivalent(searchString, searchStringBase64) {
			_, desiredHasSearchString := desiredObj[wafv2SearchStringProperty]
			_, desiredHasSearchStringBase64 := desiredObj[wafv2SearchStringBase64Prop]
			switch {
			case desiredHasSearchStringBase64 && !desiredHasSearchString:
				delete(actualObj, wafv2SearchStringProperty)
			case desiredHasSearchString && !desiredHasSearchStringBase64:
				delete(actualObj, wafv2SearchStringBase64Prop)
			default:
				delete(actualObj, wafv2SearchStringProperty)
			}
		}

		for key, child := range actualObj {
			normalizeWafv2ByteMatchBaselineValue(child, desiredObj[key])
		}
	}
}

func base64PropertyValue(value resource.PropertyValue) (resource.PropertyValue, bool) {
	if value.IsSecret() {
		encoded, ok := base64PropertyValue(value.SecretValue().Element)
		if !ok {
			return resource.PropertyValue{}, false
		}
		return resource.MakeSecret(encoded), true
	}
	if !value.IsString() {
		return resource.PropertyValue{}, false
	}
	return resource.NewStringProperty(base64.StdEncoding.EncodeToString([]byte(value.StringValue()))), true
}

func base64ValuesEquivalent(searchString, searchStringBase64 resource.PropertyValue) bool {
	if searchString.IsSecret() {
		searchString = searchString.SecretValue().Element
	}
	if searchStringBase64.IsSecret() {
		searchStringBase64 = searchStringBase64.SecretValue().Element
	}
	if !searchString.IsString() || !searchStringBase64.IsString() {
		return false
	}
	return base64.StdEncoding.EncodeToString([]byte(searchString.StringValue())) == searchStringBase64.StringValue()
}
