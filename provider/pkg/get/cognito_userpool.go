package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/cognitoidentity"
)

func (g *Getter) getCognitoUserPoolAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.cognitoidentity.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"ProviderName": nil,
		"ProviderURL": nil,
	}
	return attrs, nil
}
