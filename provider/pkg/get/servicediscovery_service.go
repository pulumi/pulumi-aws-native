package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/servicediscovery"
)

func (g *Getter) getServiceDiscoveryServiceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.servicediscovery.GetServiceWithContext(ctx, &servicediscovery.GetServiceInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Id": nil,
		"Name": nil,
	}
	return attrs, nil
}
