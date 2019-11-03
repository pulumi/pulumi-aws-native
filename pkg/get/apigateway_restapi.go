package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func (g *Getter) getApiGatewayRestApiAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.apigateway.GetRestApiWithContext(ctx, &apigateway.GetRestApiInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"RootResourceId": nil,
	}
	return attrs, nil
}
