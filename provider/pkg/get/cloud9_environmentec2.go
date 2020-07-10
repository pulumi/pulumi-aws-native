package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/cloud9"
)

func (g *Getter) getCloud9EnvironmentEC2Attributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.cloud9.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Name": nil,
	}
	return attrs, nil
}
