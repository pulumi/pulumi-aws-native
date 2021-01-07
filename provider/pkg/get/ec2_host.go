package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

func (g *Getter) getEC2HostAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	input := &ec2.DescribeHostsInput{}
	input.HostIds = []*string{aws.String(physicalResourceID)}

	resp, err := g.ec2.DescribeHostsWithContext(ctx, input)
	if err != nil {
		awsErr, ok := err.(awserr.Error)
		if ok && (awsErr.Code() == "InvalidHostID.NotFound" || awsErr.Code() == "InvalidHost.NotFound") {
			err = NewNotFoundError(err)
		}
		return nil, err
	}
	if len(resp.Hosts) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find host %v", physicalResourceID))
	}
	host := resp.Hosts[0]

	attrs := map[string]interface{}{
		"AutoPlacement": aws.StringValue(host.AutoPlacement),
		"HostRecovery": aws.StringValue(host.HostRecovery),
		"InstanceType": aws.StringValue(host.HostProperties.InstanceType),
		"AvailabilityZone": aws.StringValue(host.AvailabilityZone),
	}
	return attrs, nil
}
