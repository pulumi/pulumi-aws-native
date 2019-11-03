package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sagemaker"
)

func (g *Getter) getSageMakerCodeRepositoryAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.sagemaker.DescribeCodeRepositoryWithContext(ctx, &sagemaker.DescribeCodeRepositoryInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"CodeRepositoryName": nil,
	}
	return attrs, nil
}
