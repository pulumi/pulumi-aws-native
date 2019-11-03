package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ecs"
)

func (g *Getter) getECSClusterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ecs.DescribeClustersWithContext(ctx, &ecs.DescribeClustersInput{
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
