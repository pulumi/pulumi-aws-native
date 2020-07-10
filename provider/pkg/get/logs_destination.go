package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func (g *Getter) getLogsDestinationAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.cloudwatchlogs.DescribeDestinationsWithContext(ctx, &cloudwatchlogs.DescribeDestinationsInput{
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
