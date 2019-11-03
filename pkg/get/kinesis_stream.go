package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/kinesis"
)

func (g *Getter) getKinesisStreamAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.kinesis.DescribeStreamWithContext(ctx, &kinesis.DescribeStreamInput{
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
