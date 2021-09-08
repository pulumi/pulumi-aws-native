package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sagemaker"
)

func (g *Getter) getSageMakerEndpointAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.sagemaker.DescribeEndpointWithContext(ctx, &sagemaker.DescribeEndpointInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"EndpointName": nil,
	}
	return attrs, nil
}
