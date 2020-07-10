package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/redshift"
)

func (g *Getter) getRedshiftClusterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.redshift.DescribeClustersWithContext(ctx, &redshift.DescribeClustersInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Endpoint.Address": nil,
		"Endpoint.Port": nil,
	}
	return attrs, nil
}
