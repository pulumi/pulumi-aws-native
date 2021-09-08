package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dlm"
)

func (g *Getter) getDLMLifecyclePolicyAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.dlm.GetLifecyclePolicyWithContext(ctx, &dlm.GetLifecyclePolicyInput{
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
