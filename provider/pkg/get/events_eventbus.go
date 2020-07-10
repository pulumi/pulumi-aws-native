package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/eventbridge"
)

func (g *Getter) getEventsEventBusAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.eventbridge.DescribeEventBusWithContext(ctx, &eventbridge.DescribeEventBusInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Name": nil,
		"Policy": nil,
	}
	return attrs, nil
}
