package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/sqs"
)

func (g *Getter) getSQSQueueAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.sqs.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"QueueName": nil,
	}
	return attrs, nil
}
