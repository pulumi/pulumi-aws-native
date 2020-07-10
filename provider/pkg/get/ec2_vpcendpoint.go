package get

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

func (g *Getter) getEC2VPCEndpointAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ec2.DescribeVpcEndpointsWithContext(ctx, &ec2.DescribeVpcEndpointsInput{
		VpcEndpointIds: []*string{aws.String(physicalResourceID)},
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "InvalidVpcEndpointId.NotFound" {
			err = NewNotFoundError(err)
		}
		return nil, err
	}
	if len(resp.VpcEndpoints) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find VPC endpoint %v", physicalResourceID))
	}
	vpcEndpoint := resp.VpcEndpoints[0]

	creationTimestamp := aws.TimeValue(vpcEndpoint.CreationTimestamp).Format(timestampFormat)

	dnsEntries := make([]interface{}, len(vpcEndpoint.DnsEntries))
	for i, entry := range vpcEndpoint.DnsEntries {
		dnsEntries[i] = fmt.Sprintf("%s:%s", aws.StringValue(entry.DnsName), aws.StringValue(entry.HostedZoneId))
	}

	networkInterfaceIDs := make([]interface{}, len(vpcEndpoint.NetworkInterfaceIds))
	for i, id := range vpcEndpoint.NetworkInterfaceIds {
		networkInterfaceIDs[i] = aws.StringValue(id)
	}

	attrs := map[string]interface{}{
		"CreationTimestamp":   creationTimestamp,
		"DnsEntries":          dnsEntries,
		"NetworkInterfaceIds": networkInterfaceIDs,
	}
	return attrs, nil
}
