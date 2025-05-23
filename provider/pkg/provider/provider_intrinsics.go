// Copyright 2016-2021, Pulumi Corporation.

package provider

import (
	"context"
	"fmt"
	"net"

	gocidr "github.com/apparentlymart/go-cidr/cidr"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func (p *cfnProvider) getAZs(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	region, ok := inputs["region"]
	if !ok {
		region = resource.NewStringProperty(p.region)
	} else if !region.IsString() {
		return nil, fmt.Errorf("'region' must be a string")
	}

	resp, err := p.ec2.DescribeAvailabilityZones(ctx, &ec2.DescribeAvailabilityZonesInput{
		Filters: []types.Filter{{
			Name:   aws.String("region-name"),
			Values: []string{region.StringValue()},
		}},
	})
	if err != nil {
		return nil, err
	}

	azs := make([]string, len(resp.AvailabilityZones))
	for i, az := range resp.AvailabilityZones {
		azs[i] = *az.ZoneName
	}
	return resource.NewPropertyMapFromMap(map[string]interface{}{
		"azs": azs,
	}), nil
}

// Goal is to implement the Fn::Cidr intrinsic function
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-cidr.html
// ipBlock: The user-specified CIDR block to be split into smaller CIDR blocks.
// count: The number of CIDR blocks to generate. Valid range is between 1 and 256.
// cidrBits: The number of subnet bits for the CIDR. For example, specifying a value "8" for this parameter will create a CIDR with a mask of "/24".
//
//	Note: Subnet bits is the inverse of subnet mask. To calculate the required host bits
//	for a given subnet bits, subtract the subnet bits from 32 for ipv4 or 128 for ipv6.
func cidr(inputs resource.PropertyMap) (resource.PropertyMap, error) {
	ipBlock, ok := inputs["ipBlock"]
	if !ok {
		return nil, fmt.Errorf("missing required property 'ipBlock'")
	}
	if !ipBlock.IsString() {
		return nil, fmt.Errorf("'ipBlock' must be a string")
	}

	count, ok := inputs["count"]
	if !ok {
		return nil, fmt.Errorf("missing required property 'count'")
	}
	if !count.IsNumber() || count.NumberValue() < 1 || count.NumberValue() > 256 {
		return nil, fmt.Errorf("'count' must be a number between 1 and 256")
	}

	cidrBits, ok := inputs["cidrBits"]
	if !ok {
		return nil, fmt.Errorf("missing required property 'cidrBits'")
	}
	if !cidrBits.IsNumber() || cidrBits.NumberValue() < 0 {
		return nil, fmt.Errorf("'cidrBits' must be a positive number")
	}

	_, network, err := net.ParseCIDR(ipBlock.StringValue())
	if err != nil {
		return nil, fmt.Errorf("invalid IP block: %s", err)
	}
	protocol := "IP"
	bits := len(network.IP) * 8
	switch bits {
	case 32:
		protocol = "IPv4"
	case 128:
		protocol = "IPv6"
	}

	subnets := make([]resource.PropertyValue, int(count.NumberValue()))
	subnetBits := int(cidrBits.NumberValue())

	if subnetBits > bits {
		return nil, fmt.Errorf("cidrBits %d is more than %d bits for an %s address. \n"+
			"cidrBits is the inverse of subnet mask, e.g. cidrBits=5 would create a subnet mask of '/27'", subnetBits, bits, protocol)
	}

	prefixLen := bits - subnetBits

	current, ok := gocidr.PreviousSubnet(network, prefixLen)
	// ok is true if we have rolled over (which we don't want)
	if ok {
		return nil, fmt.Errorf("not enough remaining address space for a subnet with a prefix of %d bits after %s", len(subnets), current.String())
	}
	for i := 0; i < len(subnets); i++ {
		subnet, ok := gocidr.NextSubnet(current, prefixLen)
		if ok || !network.Contains(subnet.IP) {
			return nil, fmt.Errorf("not enough remaining address space for a subnet with a prefix of %d bits after %s", len(subnets), current.String())
		}
		current = subnet
		subnets[i] = resource.NewStringProperty(subnet.String())
	}

	return resource.PropertyMap{
		"subnets": resource.NewArrayProperty(subnets),
	}, nil
}

func (p *cfnProvider) cidr(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	return cidr(inputs)
}

func (p *cfnProvider) importValue(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	name, ok := inputs["name"]
	if !ok {
		return nil, fmt.Errorf("missing required property 'name'")
	}
	if !name.IsString() {
		return nil, fmt.Errorf("'name' must be a string")
	}

	value, ok := "", false
	paginator := cloudformation.NewListExportsPaginator(p.cfn, &cloudformation.ListExportsInput{})
	for paginator.HasMorePages() && !ok {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, export := range output.Exports {
			if *export.Name == name.StringValue() {
				value, ok = *export.Value, true
				break
			}
		}
	}
	if !ok {
		return nil, fmt.Errorf("unknown export '%s'", name.StringValue())
	}

	return resource.PropertyMap{
		"value": resource.NewStringProperty(value),
	}, nil
}
