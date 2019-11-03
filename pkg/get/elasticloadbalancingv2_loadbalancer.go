package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/elbv2"
)

func (g *Getter) getElasticLoadBalancingV2LoadBalancerAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.elbv2.DescribeLoadBalancersWithContext(ctx, &elbv2.DescribeLoadBalancersInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"CanonicalHostedZoneID": nil,
		"DNSName": nil,
		"LoadBalancerFullName": nil,
		"LoadBalancerName": nil,
		"SecurityGroups": nil,
	}
	return attrs, nil
}
