package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/codepipeline"
)

func (g *Getter) getCodePipelineWebhookAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.codepipeline.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Url": nil,
	}
	return attrs, nil
}
