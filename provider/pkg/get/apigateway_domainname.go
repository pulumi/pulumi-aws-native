package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func (g *Getter) getApiGatewayDomainNameAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.apigateway.GetDomainNameWithContext(ctx, &apigateway.GetDomainNameInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"DistributionDomainName": nil,
		"DistributionHostedZoneId": nil,
		"RegionalDomainName": nil,
		"RegionalHostedZoneId": nil,
	}
	return attrs, nil
}
