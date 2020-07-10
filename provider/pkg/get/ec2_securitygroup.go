package get

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

func (g *Getter) getEC2SecurityGroupAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	input := &ec2.DescribeSecurityGroupsInput{}
	if isID := strings.HasPrefix(physicalResourceID, "sg-"); isID {
		input.GroupIds = []*string{aws.String(physicalResourceID)}
	} else {
		input.GroupNames = []*string{aws.String(physicalResourceID)}
	}

	resp, err := g.ec2.DescribeSecurityGroupsWithContext(ctx, input)
	if err != nil {
		awsErr, ok := err.(awserr.Error)
		if ok && (awsErr.Code() == "InvalidSecurityGroupID.NotFound" || awsErr.Code() == "InvalidGroup.NotFound") {
			err = NewNotFoundError(err)
		}
		return nil, err
	}
	if len(resp.SecurityGroups) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find security group %v", physicalResourceID))
	}
	securityGroup := resp.SecurityGroups[0]

	attrs := map[string]interface{}{
		"GroupId": aws.StringValue(securityGroup.GroupId),
		"VpcId":   aws.StringValue(securityGroup.VpcId),
	}
	return attrs, nil
}
