package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/configservice"
)

func (g *Getter) getConfigConfigRuleAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.configservice.DescribeConfigRulesWithContext(ctx, &configservice.DescribeConfigRulesInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Compliance.Type": nil,
		"ConfigRuleId": nil,
	}
	return attrs, nil
}
