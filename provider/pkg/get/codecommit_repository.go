package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/codecommit"
)

func (g *Getter) getCodeCommitRepositoryAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.codecommit.GetRepositoryWithContext(ctx, &codecommit.GetRepositoryInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"CloneUrlHttp": nil,
		"CloneUrlSsh": nil,
		"Name": nil,
	}
	return attrs, nil
}
