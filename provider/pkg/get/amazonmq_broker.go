package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/mq"
)

func (g *Getter) getAmazonMQBrokerAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.mq.DescribeBrokerWithContext(ctx, &mq.DescribeBrokerInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"AmqpEndpoints": nil,
		"Arn": nil,
		"ConfigurationId": nil,
		"ConfigurationRevision": nil,
		"IpAddresses": nil,
		"MqttEndpoints": nil,
		"OpenWireEndpoints": nil,
		"StompEndpoints": nil,
		"WssEndpoints": nil,
	}
	return attrs, nil
}
