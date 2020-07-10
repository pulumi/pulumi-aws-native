package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

func (g *Getter) getCloudTrailTrailAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.cloudtrail.DescribeTrailsWithContext(ctx, &cloudtrail.DescribeTrailsInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"SnsTopicArn": nil,
	}
	return attrs, nil
}
