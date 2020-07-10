package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/iot"
)

func (g *Getter) getIoTPolicyAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.iot.GetPolicyWithContext(ctx, &iot.GetPolicyInput{
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
