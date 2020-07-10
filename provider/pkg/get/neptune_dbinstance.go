package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/neptune"
)

func (g *Getter) getNeptuneDBInstanceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.neptune.DescribeDBInstancesWithContext(ctx, &neptune.DescribeDBInstancesInput{
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
