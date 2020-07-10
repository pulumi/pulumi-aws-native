package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dax"
)

func (g *Getter) getDAXClusterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.dax.DescribeClustersWithContext(ctx, &dax.DescribeClustersInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"ClusterDiscoveryEndpoint": nil,
	}
	return attrs, nil
}
