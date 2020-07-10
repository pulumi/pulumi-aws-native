package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sagemaker"
)

func (g *Getter) getSageMakerNotebookInstanceLifecycleConfigAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.sagemaker.DescribeNotebookInstanceLifecycleConfigWithContext(ctx, &sagemaker.DescribeNotebookInstanceLifecycleConfigInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"NotebookInstanceLifecycleConfigName": nil,
	}
	return attrs, nil
}
