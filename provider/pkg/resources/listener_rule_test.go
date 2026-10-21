//nolint:goconst // Repeated domain values keep membership fixtures readable.
package resources

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func TestAlignListenerRuleTargetGroupsMembership(t *testing.T) {
	group := func(arn string, weight float64) resource.PropertyValue {
		return resource.NewObjectProperty(resource.PropertyMap{
			"targetGroupArn": resource.NewStringProperty(arn),
			"weight":         resource.NewNumberProperty(weight),
		})
	}
	inputs := func(groups ...resource.PropertyValue) resource.PropertyMap {
		return resource.NewPropertyMapFromMap(map[string]interface{}{
			"actions": []interface{}{map[string]interface{}{
				"type": "forward", "order": 2,
				"forwardConfig": map[string]interface{}{
					"targetGroups": resource.NewArrayProperty(groups),
				},
			}},
		})
	}
	blue, green, red := group("blue", 0), group("green", 100), group("red", 10)
	unknown := resource.NewObjectProperty(resource.PropertyMap{
		"targetGroupArn": resource.MakeComputed(resource.NewStringProperty("")),
	})
	secret := resource.NewObjectProperty(resource.PropertyMap{
		"targetGroupArn": resource.MakeSecret(resource.NewStringProperty("blue")),
	})
	groups := func(values ...resource.PropertyValue) []resource.PropertyValue { return values }
	missing := resource.NewObjectProperty(resource.PropertyMap{})
	for _, tc := range []struct {
		name     string
		old      []resource.PropertyValue
		observed []resource.PropertyValue
		want     []resource.PropertyValue
	}{
		{"append new group", groups(blue, green), groups(red, green, blue), groups(blue, green, red)},
		{"omit removed group", groups(blue, green), groups(green), groups(green)},
		{"external replacement", groups(blue, green), groups(red, green), groups(green, red)},
		{"empty previous collection", nil, groups(green, blue), groups(green, blue)},
		{"empty observed collection", groups(blue, green), nil, []resource.PropertyValue{}},
		{"duplicate previous ARN", groups(blue, blue), groups(green, blue), groups(green, blue)},
		{"duplicate observed ARN", groups(blue, green), groups(green, blue, blue), groups(green, blue, blue)},
		{"missing ARN", groups(blue, green), groups(green, missing), groups(green, missing)},
		{"unknown previous ARN", groups(unknown, green), groups(green, blue), groups(green, blue)},
		{"secret observed ARN", groups(blue, green), groups(green, secret), groups(green, secret)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			old, actual := inputs(tc.old...), inputs(tc.observed...)
			baseline := clonePropertyMap(actual)
			alignListenerRuleTargetGroups(old, baseline)
			require.Equal(t, tc.want, listenerRuleForwardConfig(baseline)["targetGroups"].ArrayValue())
			require.Equal(t, inputs(tc.old...), old, "previous inputs are immutable")
			require.Equal(t, inputs(tc.observed...), actual, "outputs are immutable")
		})
	}

	t.Run("find forward action after authentication action reorders", func(t *testing.T) {
		old, actual := inputs(blue, green), inputs(green, blue)
		auth := resource.NewObjectProperty(resource.PropertyMap{"type": resource.NewStringProperty("authenticate-oidc")})
		old["actions"] = resource.NewArrayProperty(append([]resource.PropertyValue{auth}, old["actions"].ArrayValue()...))
		actual["actions"] = resource.NewArrayProperty(append(actual["actions"].ArrayValue(), auth))
		alignListenerRuleTargetGroups(old, actual)
		require.Equal(t, groups(blue, green), listenerRuleForwardConfig(actual)["targetGroups"].ArrayValue())
		require.Equal(t, "forward", actual["actions"].ArrayValue()[0].ObjectValue()["type"].StringValue())
	})
}
