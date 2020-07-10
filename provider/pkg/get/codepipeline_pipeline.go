package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/codepipeline"
)

func (g *Getter) getCodePipelinePipelineAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.codepipeline.GetPipelineWithContext(ctx, &codepipeline.GetPipelineInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Version": nil,
	}
	return attrs, nil
}
