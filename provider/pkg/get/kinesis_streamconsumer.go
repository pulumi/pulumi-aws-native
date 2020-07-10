package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/kinesis"
)

func (g *Getter) getKinesisStreamConsumerAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.kinesis.DescribeStreamConsumerWithContext(ctx, &kinesis.DescribeStreamConsumerInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ConsumerARN": nil,
		"ConsumerCreationTimestamp": nil,
		"ConsumerName": nil,
		"ConsumerStatus": nil,
		"StreamARN": nil,
	}
	return attrs, nil
}
