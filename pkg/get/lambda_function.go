package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/lambda"
)

func (g *Getter) getLambdaFunctionAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.lambda.GetFunctionWithContext(ctx, &lambda.GetFunctionInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
	}
	return attrs, nil
}
