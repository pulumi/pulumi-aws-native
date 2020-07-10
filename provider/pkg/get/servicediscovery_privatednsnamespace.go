package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/servicediscovery"
)

func (g *Getter) getServiceDiscoveryPrivateDnsNamespaceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.servicediscovery.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Id": nil,
	}
	return attrs, nil
}
