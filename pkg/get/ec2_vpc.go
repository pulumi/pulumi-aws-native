package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

func (g *Getter) getEC2VPCAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ec2.DescribeVpcsWithContext(ctx, &ec2.DescribeVpcsInput{
		VpcIds: []*string{aws.String(physicalResourceID)},
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "InvalidVpcID.NotFound" {
			err = NewNotFoundError(err)
		}
		return nil, err
	}
	if len(resp.Vpcs) != 1 {
		return nil, NewNotFoundError(errors.Errorf("could not find VPC %v", physicalResourceID))
	}
	vpc := resp.Vpcs[0]

	var defaultNetworkAcl string
	aclsResp, err := g.ec2.DescribeNetworkAclsWithContext(ctx, &ec2.DescribeNetworkAclsInput{
		Filters: []*ec2.Filter{
			{Name: aws.String("default"), Values: []*string{aws.String("true")}},
			{Name: aws.String("vpc-id"), Values: []*string{aws.String(physicalResourceID)}},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(aclsResp.NetworkAcls) == 1 {
		defaultNetworkAcl = aws.StringValue(aclsResp.NetworkAcls[0].NetworkAclId)
	}

	var defaultSecurityGroup string
	sgsResp, err := g.ec2.DescribeSecurityGroupsWithContext(ctx, &ec2.DescribeSecurityGroupsInput{
		Filters: []*ec2.Filter{
			{Name: aws.String("group-name"), Values: []*string{aws.String("default")}},
			{Name: aws.String("vpc-id"), Values: []*string{aws.String(physicalResourceID)}},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(sgsResp.SecurityGroups) == 1 {
		defaultSecurityGroup = aws.StringValue(sgsResp.SecurityGroups[0].GroupId)
	}

	var cidrBlockAssociations []interface{}
	for _, association := range vpc.CidrBlockAssociationSet {
		cidrBlockAssociations = append(cidrBlockAssociations, aws.StringValue(association.CidrBlock))
	}

	var ipv6CidrBlocks []interface{}
	for _, association := range vpc.Ipv6CidrBlockAssociationSet {
		ipv6CidrBlocks = append(ipv6CidrBlocks, aws.StringValue(association.Ipv6CidrBlock))
	}

	attrs := map[string]interface{}{
		"CidrBlock":             aws.StringValue(vpc.CidrBlock),
		"CidrBlockAssociations": cidrBlockAssociations,
		"DefaultNetworkAcl":     defaultNetworkAcl,
		"DefaultSecurityGroup":  defaultSecurityGroup,
		"Ipv6CidrBlocks":        ipv6CidrBlocks,
	}
	return attrs, nil
}
