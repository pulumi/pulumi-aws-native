package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (g *Getter) getEC2VPCEndpointAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ec2.DescribeVpcEndpointsWithContext(ctx, &ec2.DescribeVpcEndpointsInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"CreationTimestamp": nil,
		"DnsEntries": nil,
		"NetworkInterfaceIds": nil,
	}
	return attrs, nil
}
