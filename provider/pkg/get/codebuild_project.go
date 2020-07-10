package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/codebuild"
)

func (g *Getter) getCodeBuildProjectAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.codebuild.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
	}
	return attrs, nil
}
