package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/cognitoidentity"
)

func (g *Getter) getCognitoIdentityPoolAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.cognitoidentity.DescribeIdentityPoolWithContext(ctx, &cognitoidentity.DescribeIdentityPoolInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Name": nil,
	}
	return attrs, nil
}
