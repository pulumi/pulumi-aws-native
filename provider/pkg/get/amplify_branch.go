package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/amplify"
)

func (g *Getter) getAmplifyBranchAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.amplify.GetBranchWithContext(ctx, &amplify.GetBranchInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"BranchName": nil,
	}
	return attrs, nil
}
