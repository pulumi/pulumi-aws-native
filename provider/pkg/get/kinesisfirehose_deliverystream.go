package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/firehose"
)

func (g *Getter) getKinesisFirehoseDeliveryStreamAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.firehose.DescribeDeliveryStreamWithContext(ctx, &firehose.DescribeDeliveryStreamInput{
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
