package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/neptune"
)

func (g *Getter) getNeptuneDBClusterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.neptune.DescribeDBClustersWithContext(ctx, &neptune.DescribeDBClustersInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ClusterResourceId": nil,
		"Endpoint": nil,
		"Port": nil,
		"ReadEndpoint": nil,
	}
	return attrs, nil
}
