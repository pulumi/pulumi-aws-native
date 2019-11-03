package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/elasticache"
)

func (g *Getter) getElastiCacheReplicationGroupAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.elasticache.DescribeReplicationGroupsWithContext(ctx, &elasticache.DescribeReplicationGroupsInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ConfigurationEndPoint.Address": nil,
		"ConfigurationEndPoint.Port": nil,
		"PrimaryEndPoint.Address": nil,
		"PrimaryEndPoint.Port": nil,
		"ReadEndPoint.Addresses": nil,
		"ReadEndPoint.Addresses.List": nil,
		"ReadEndPoint.Ports": nil,
		"ReadEndPoint.Ports.List": nil,
	}
	return attrs, nil
}
