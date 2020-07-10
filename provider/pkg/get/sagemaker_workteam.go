package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sagemaker"
)

func (g *Getter) getSageMakerWorkteamAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.sagemaker.DescribeWorkteamWithContext(ctx, &sagemaker.DescribeWorkteamInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"WorkteamName": nil,
	}
	return attrs, nil
}
