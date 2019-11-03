package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/iot"
)

func (g *Getter) getIoTTopicRuleAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.iot.GetTopicRuleWithContext(ctx, &iot.GetTopicRuleInput{
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
