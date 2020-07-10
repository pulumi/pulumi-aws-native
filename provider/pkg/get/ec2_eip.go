package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

func (g *Getter) getEC2EIPAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ec2.DescribeAddressesWithContext(ctx, &ec2.DescribeAddressesInput{
		PublicIps: []*string{aws.String(physicalResourceID)},
	})
	if err != nil {
		awsErr, ok := err.(awserr.Error)
		if ok && (awsErr.Code() == "InvalidAllocationID.NotFound" || awsErr.Code() == "InvalidAddress.NotFound") {
			err = NewNotFoundError(awsErr)
		}
		return nil, err
	}
	if len(resp.Addresses) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find EIP %v", physicalResourceID))
	}

	attrs := map[string]interface{}{
		"AllocationId": aws.StringValue(resp.Addresses[0].AllocationId),
	}
	return attrs, nil
}
