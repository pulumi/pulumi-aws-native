package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/apigatewayv2"
)

func (g *Getter) getApiGatewayV2DomainNameAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.apigatewayv2.GetDomainNameWithContext(ctx, &apigatewayv2.GetDomainNameInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"RegionalDomainName": nil,
		"RegionalHostedZoneId": nil,
	}
	return attrs, nil
}
