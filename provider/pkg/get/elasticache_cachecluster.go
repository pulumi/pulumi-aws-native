package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/elasticache"
)

func (g *Getter) getElastiCacheCacheClusterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.elasticache.DescribeCacheClustersWithContext(ctx, &elasticache.DescribeCacheClustersInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ConfigurationEndpoint.Address": nil,
		"ConfigurationEndpoint.Port": nil,
		"RedisEndpoint.Address": nil,
		"RedisEndpoint.Port": nil,
	}
	return attrs, nil
}
