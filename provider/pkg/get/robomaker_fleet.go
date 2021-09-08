package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/robomaker"
)

func (g *Getter) getRoboMakerFleetAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.robomaker.DescribeFleetWithContext(ctx, &robomaker.DescribeFleetInput{
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
