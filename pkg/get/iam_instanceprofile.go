package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/iam"
)

func (g *Getter) getIAMInstanceProfileAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.iam.GetInstanceProfileWithContext(ctx, &iam.GetInstanceProfileInput{
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
