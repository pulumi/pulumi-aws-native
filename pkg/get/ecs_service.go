package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ecs"
)

func (g *Getter) getECSServiceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ecs.DescribeServicesWithContext(ctx, &ecs.DescribeServicesInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Name": nil,
	}
	return attrs, nil
}
