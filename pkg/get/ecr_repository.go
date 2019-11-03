package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/ecr"
)

func (g *Getter) getECRRepositoryAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.ecr.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
	}
	return attrs, nil
}
