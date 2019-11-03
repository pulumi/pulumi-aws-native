package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sagemaker"
)

func (g *Getter) getSageMakerNotebookInstanceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.sagemaker.DescribeNotebookInstanceWithContext(ctx, &sagemaker.DescribeNotebookInstanceInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"NotebookInstanceName": nil,
	}
	return attrs, nil
}
