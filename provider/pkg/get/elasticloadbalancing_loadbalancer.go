package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/elb"
)

func (g *Getter) getElasticLoadBalancingLoadBalancerAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.elb.DescribeLoadBalancersWithContext(ctx, &elb.DescribeLoadBalancersInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"CanonicalHostedZoneName": nil,
		"CanonicalHostedZoneNameID": nil,
		"DNSName": nil,
		"SourceSecurityGroup.GroupName": nil,
		"SourceSecurityGroup.OwnerAlias": nil,
	}
	return attrs, nil
}
