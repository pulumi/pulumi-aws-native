package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

func (g *Getter) getEC2SubnetNetworkAclAssociationAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ec2.DescribeNetworkAclsWithContext(ctx, &ec2.DescribeNetworkAclsInput{
		Filters: []*ec2.Filter{{Name: aws.String("association.association-id"), Values: []*string{aws.String(physicalResourceID)}}},
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "InvalidNetworkAclID.NotFound" {
			err = NewNotFoundError(err)
		}
		return nil, err
	}
	if len(resp.NetworkAcls) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find network ACL association %v", physicalResourceID))
	}

	var subnetID string
	for _, association := range resp.NetworkAcls[0].Associations {
		if aws.StringValue(association.NetworkAclAssociationId) == physicalResourceID {
			subnetID = aws.StringValue(association.SubnetId)
			break
		}
	}

	// From https://docs.aws.amazon.com/en_pv/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html: Returns the value of this object's SubnetId property.
	attrs := map[string]interface{}{
		"AssociationId": subnetID,
	}
	return attrs, nil
}
