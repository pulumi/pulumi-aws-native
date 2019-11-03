package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/lambda"
)

func (g *Getter) getLambdaVersionAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.lambda.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Version": nil,
	}
	return attrs, nil
}
