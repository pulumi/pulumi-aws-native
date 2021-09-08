package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/pkg/errors"
)

func (g *Getter) getLogsLogGroupAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	var arn string
	err := g.cloudwatchlogs.DescribeLogGroupsPagesWithContext(ctx, &cloudwatchlogs.DescribeLogGroupsInput{
		LogGroupNamePrefix: aws.String(physicalResourceID),
	}, func(resp *cloudwatchlogs.DescribeLogGroupsOutput, _ bool) bool {
		for _, lg := range resp.LogGroups {
			if aws.StringValue(lg.LogGroupName) == physicalResourceID {
				arn = aws.StringValue(lg.Arn)
				return false
			}
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	if arn == "" {
		return nil, NewNotFoundError(errors.Errorf("could not find log group %v", physicalResourceID))
	}

	attrs := map[string]interface{}{
		"Arn": arn,
	}
	return attrs, nil
}
