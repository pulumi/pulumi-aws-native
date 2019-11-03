package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/elbv2"
)

func (g *Getter) getElasticLoadBalancingV2TargetGroupAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.elbv2.DescribeTargetGroupsWithContext(ctx, &elbv2.DescribeTargetGroupsInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"LoadBalancerArns": nil,
		"TargetGroupFullName": nil,
		"TargetGroupName": nil,
	}
	return attrs, nil
}
