package resources

import "github.com/pulumi/pulumi/sdk/v3/go/common/resource"

// alignListenerRuleTargetGroups keeps observed target-group values associated
// with their ARNs while restoring the previous input order. The engine applies
// ignoreChanges by array index, so an AWS-only reorder must not become the next
// input baseline. This operates on a cloned baseline, never on live outputs.
func alignListenerRuleTargetGroups(oldInputs, baseline resource.PropertyMap) {
	oldConfig := listenerRuleForwardConfig(oldInputs)
	actualConfig := listenerRuleForwardConfig(baseline)
	if oldConfig == nil || actualConfig == nil {
		return
	}
	oldGroups := oldConfig["targetGroups"].ArrayValue()
	actualGroups := actualConfig["targetGroups"].ArrayValue()
	oldARNs, ok := listenerRuleTargetGroupARNs(oldGroups)
	if !ok {
		return
	}
	actualARNs, ok := listenerRuleTargetGroupARNs(actualGroups)
	if !ok {
		return
	}

	byARN := make(map[string]resource.PropertyValue, len(actualGroups))
	for i, arn := range actualARNs {
		byARN[arn] = actualGroups[i]
	}
	aligned := make([]resource.PropertyValue, 0, len(actualGroups))
	for _, arn := range oldARNs {
		if group, exists := byARN[arn]; exists {
			aligned = append(aligned, group)
			delete(byARN, arn)
		}
	}
	// Removed groups are omitted. New and externally replaced ARNs have no
	// previous identity, so retain them in observed order after the survivors.
	for i, arn := range actualARNs {
		if _, exists := byARN[arn]; exists {
			aligned = append(aligned, actualGroups[i])
		}
	}
	actualConfig["targetGroups"] = resource.NewArrayProperty(aligned)
}

// A ListenerRule has exactly one routing action. Find its forward action by
// type rather than assuming that authentication and routing actions retain
// their positions. Ambiguous or opaque shapes retain their observed ordering.
func listenerRuleForwardConfig(inputs resource.PropertyMap) resource.PropertyMap {
	actions := inputs["actions"]
	if !actions.IsArray() {
		return nil
	}
	var config resource.PropertyMap
	for _, action := range actions.ArrayValue() {
		if !action.IsObject() {
			return nil
		}
		obj := action.ObjectValue()
		if !obj["type"].IsString() {
			return nil
		}
		if obj["type"].StringValue() != "forward" {
			continue
		}
		if config != nil || !obj["forwardConfig"].IsObject() {
			return nil
		}
		config = obj["forwardConfig"].ObjectValue()
		if !config["targetGroups"].IsArray() {
			return nil
		}
	}
	return config
}

func listenerRuleTargetGroupARNs(groups []resource.PropertyValue) ([]string, bool) {
	arns := make([]string, len(groups))
	seen := make(map[string]bool, len(groups))
	for i, group := range groups {
		if !group.IsObject() {
			return nil, false
		}
		arn := group.ObjectValue()["targetGroupArn"]
		if !arn.IsString() || arn.StringValue() == "" || seen[arn.StringValue()] {
			return nil, false
		}
		arns[i] = arn.StringValue()
		seen[arns[i]] = true
	}
	return arns, true
}
