package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (g *Getter) getEC2NetworkInterfaceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ec2.DescribeNetworkInterfacesWithContext(ctx, &ec2.DescribeNetworkInterfacesInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"PrimaryPrivateIpAddress": nil,
		"SecondaryPrivateIpAddresses": nil,
	}
	return attrs, nil
}
