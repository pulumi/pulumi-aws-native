package get

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/pkg/errors"
)

func (g *Getter) getElasticLoadBalancingV2LoadBalancerAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.elbv2.DescribeLoadBalancersWithContext(ctx, &elbv2.DescribeLoadBalancersInput{
		LoadBalancerArns: []*string{aws.String(physicalResourceID)},
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "LoadBalancerNotFound" {
			return nil, NewNotFoundError(err)
		}
		return nil, err
	}
	if len(resp.LoadBalancers) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find load balancer %v", physicalResourceID))
	}
	loadBalancer := resp.LoadBalancers[0]

	canonicalHostedZoneID := aws.StringValue(loadBalancer.CanonicalHostedZoneId)
	dnsName := aws.StringValue(loadBalancer.DNSName)

	loadBalancerFullName := ""
	if components := strings.Split(aws.StringValue(loadBalancer.LoadBalancerArn), ":loadbalancer/"); len(components) == 2 {
		loadBalancerFullName = components[1]
	}

	loadBalancerName := aws.StringValue(loadBalancer.LoadBalancerName)

	securityGroups := make([]interface{}, len(loadBalancer.SecurityGroups))
	for i, sg := range loadBalancer.SecurityGroups {
		securityGroups[i] = aws.StringValue(sg)
	}

	attrs := map[string]interface{}{
		"CanonicalHostedZoneID": canonicalHostedZoneID,
		"DNSName":               dnsName,
		"LoadBalancerFullName":  loadBalancerFullName,
		"LoadBalancerName":      loadBalancerName,
		"SecurityGroups":        securityGroups,
	}
	return attrs, nil
}
