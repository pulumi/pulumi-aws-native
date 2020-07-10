package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/cloudformation"
)

func (g *Getter) getCloudFormationWaitConditionAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.cloudformation.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Data": nil,
	}
	return attrs, nil
}
