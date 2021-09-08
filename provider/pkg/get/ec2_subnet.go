package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

func (g *Getter) getEC2SubnetAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ec2.DescribeSubnetsWithContext(ctx, &ec2.DescribeSubnetsInput{
		SubnetIds: []*string{aws.String(physicalResourceID)},
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "InvalidSubnetID.NotFound" {
			err = NewNotFoundError(err)
		}
		return nil, err
	}
	if len(resp.Subnets) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find subnet %v", physicalResourceID))
	}
	subnet := resp.Subnets[0]

	var networkAclAssociationID string
	aclsResp, err := g.ec2.DescribeNetworkAclsWithContext(ctx, &ec2.DescribeNetworkAclsInput{
		Filters: []*ec2.Filter{{Name: aws.String("association.subnet-id"), Values: []*string{aws.String(physicalResourceID)}}},
	})
	if err != nil {
		return nil, err
	}
	if len(aclsResp.NetworkAcls) == 1 {
		for _, association := range aclsResp.NetworkAcls[0].Associations {
			if aws.StringValue(association.SubnetId) == physicalResourceID {
				networkAclAssociationID = aws.StringValue(association.NetworkAclAssociationId)
				break
			}
		}
	}

	var ipv6CidrBlocks []interface{}
	for _, association := range subnet.Ipv6CidrBlockAssociationSet {
		ipv6CidrBlocks = append(ipv6CidrBlocks, aws.StringValue(association.Ipv6CidrBlock))
	}

	attrs := map[string]interface{}{
		"AvailabilityZone":        aws.StringValue(subnet.AvailabilityZone),
		"Ipv6CidrBlocks":          ipv6CidrBlocks,
		"NetworkAclAssociationId": networkAclAssociationID,
		"VpcId":                   aws.StringValue(subnet.VpcId),
	}
	return attrs, nil
}
