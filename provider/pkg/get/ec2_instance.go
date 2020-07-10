package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

func (g *Getter) getEC2InstanceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ec2.DescribeInstancesWithContext(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(physicalResourceID)},
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "InvalidInstanceID.NotFound" {
			err = NewNotFoundError(err)
		}
		return nil, err
	}
	if len(resp.Reservations) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find EC2 instance %v", physicalResourceID))
	}

	reservation := resp.Reservations[0]
	if len(reservation.Instances) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find EC2 instance %v", physicalResourceID))
	}
	instance := reservation.Instances[0]

	var availabilityZone string
	if instance.Placement != nil {
		availabilityZone = aws.StringValue(instance.Placement.AvailabilityZone)
	}

	attrs := map[string]interface{}{
		"AvailabilityZone": availabilityZone,
		"PrivateDnsName":   aws.StringValue(instance.PrivateDnsName),
		"PrivateIp":        aws.StringValue(instance.PrivateIpAddress),
		"PublicDnsName":    aws.StringValue(instance.PublicDnsName),
		"PublicIp":         aws.StringValue(instance.PublicIpAddress),
	}
	return attrs, nil
}
