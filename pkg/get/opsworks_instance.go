package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/opsworks"
)

func (g *Getter) getOpsWorksInstanceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.opsworks.DescribeInstancesWithContext(ctx, &opsworks.DescribeInstancesInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"AvailabilityZone": nil,
		"PrivateDnsName": nil,
		"PrivateIp": nil,
		"PublicDnsName": nil,
		"PublicIp": nil,
	}
	return attrs, nil
}
