package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sagemaker"
)

func (g *Getter) getSageMakerModelAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.sagemaker.DescribeModelWithContext(ctx, &sagemaker.DescribeModelInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ModelName": nil,
	}
	return attrs, nil
}
