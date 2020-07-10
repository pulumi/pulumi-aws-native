package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/amplify"
)

func (g *Getter) getAmplifyDomainAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.amplify.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"CertificateRecord": nil,
		"DomainName": nil,
		"DomainStatus": nil,
		"StatusReason": nil,
	}
	return attrs, nil
}
