package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/rds"
)

func (g *Getter) getRDSDBInstanceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.rds.DescribeDBInstancesWithContext(ctx, &rds.DescribeDBInstancesInput{
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
