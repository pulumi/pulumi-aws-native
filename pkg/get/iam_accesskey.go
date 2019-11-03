package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/iam"
)

func (g *Getter) getIAMAccessKeyAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.iam.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"SecretAccessKey": nil,
	}
	return attrs, nil
}
