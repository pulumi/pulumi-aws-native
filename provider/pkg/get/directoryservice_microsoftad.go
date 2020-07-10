package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/directoryservice"
)

func (g *Getter) getDirectoryServiceMicrosoftADAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.directoryservice.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Alias": nil,
		"DnsIpAddresses": nil,
	}
	return attrs, nil
}
