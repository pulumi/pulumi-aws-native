package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/autoscalingplans"
)

func (g *Getter) getAutoScalingPlansScalingPlanAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.autoscalingplans.DescribeScalingPlansWithContext(ctx, &autoscalingplans.DescribeScalingPlansInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ScalingPlanName": nil,
		"ScalingPlanVersion": nil,
	}
	return attrs, nil
}
