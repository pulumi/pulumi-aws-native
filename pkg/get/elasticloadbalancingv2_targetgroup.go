package get

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/pkg/errors"
)

func (g *Getter) getElasticLoadBalancingV2TargetGroupAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.elbv2.DescribeTargetGroupsWithContext(ctx, &elbv2.DescribeTargetGroupsInput{
		TargetGroupArns: []*string{aws.String(physicalResourceID)},
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == elbv2.ErrCodeTargetGroupNotFoundException {
			return nil, NewNotFoundError(err)
		}
		return nil, err
	}
	if len(resp.TargetGroups) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find target group %v", physicalResourceID))
	}
	targetGroup := resp.TargetGroups[0]

	loadBalancerArns := make([]interface{}, len(targetGroup.LoadBalancerArns))
	for i, arn := range targetGroup.LoadBalancerArns {
		loadBalancerArns[i] = aws.StringValue(arn)
	}

	targetGroupFullName := ""
	if components := strings.Split(aws.StringValue(targetGroup.TargetGroupArn), ":targetgroup/"); len(components) == 2 {
		targetGroupFullName = components[1]
	}

	targetGroupName := aws.StringValue(targetGroup.TargetGroupName)

	attrs := map[string]interface{}{
		"LoadBalancerArns":    loadBalancerArns,
		"TargetGroupFullName": targetGroupFullName,
		"TargetGroupName":     targetGroupName,
	}
	return attrs, nil
}
