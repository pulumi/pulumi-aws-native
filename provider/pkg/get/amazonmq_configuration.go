package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/mq"
)

func (g *Getter) getAmazonMQConfigurationAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.mq.DescribeConfigurationWithContext(ctx, &mq.DescribeConfigurationInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Id": nil,
		"Revision": nil,
	}
	return attrs, nil
}
