package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/eventbridge"
)

func (g *Getter) getEventsRuleAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.eventbridge.DescribeRuleWithContext(ctx, &eventbridge.DescribeRuleInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
	}
	return attrs, nil
}
