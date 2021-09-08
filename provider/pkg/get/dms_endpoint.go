package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
)

func (g *Getter) getDMSEndpointAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.databasemigrationservice.DescribeEndpointsWithContext(ctx, &databasemigrationservice.DescribeEndpointsInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ExternalId": nil,
	}
	return attrs, nil
}
