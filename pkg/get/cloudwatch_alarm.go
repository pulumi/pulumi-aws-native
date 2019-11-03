package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func (g *Getter) getCloudWatchAlarmAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.cloudwatch.DescribeAlarmsWithContext(ctx, &cloudwatch.DescribeAlarmsInput{
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
