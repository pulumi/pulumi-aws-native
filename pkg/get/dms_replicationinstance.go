package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
)

func (g *Getter) getDMSReplicationInstanceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.databasemigrationservice.DescribeReplicationInstancesWithContext(ctx, &databasemigrationservice.DescribeReplicationInstancesInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ReplicationInstancePrivateIpAddresses": nil,
		"ReplicationInstancePublicIpAddresses": nil,
	}
	return attrs, nil
}
