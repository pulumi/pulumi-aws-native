package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/docdb"
)

func (g *Getter) getDocDBDBInstanceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.docdb.DescribeDBInstancesWithContext(ctx, &docdb.DescribeDBInstancesInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Endpoint": nil,
		"Port": nil,
	}
	return attrs, nil
}
