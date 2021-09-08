package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/iam"
)

func (g *Getter) getIAMGroupAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.iam.GetGroupWithContext(ctx, &iam.GetGroupInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
	}
	return attrs, nil
}
