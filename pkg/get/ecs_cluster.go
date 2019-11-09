package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/pkg/errors"
)

func (g *Getter) getECSClusterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ecs.DescribeClustersWithContext(ctx, &ecs.DescribeClustersInput{
		Clusters: []*string{aws.String(physicalResourceID)},
	})
	if err != nil {
		return nil, err
	}
	if len(resp.Failures) != 0 || len(resp.Clusters) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find ECS cluster %v", physicalResourceID))
	}
	cluster := resp.Clusters[0]

	arn := aws.StringValue(cluster.ClusterArn)

	attrs := map[string]interface{}{
		"Arn": arn,
	}
	return attrs, nil
}
