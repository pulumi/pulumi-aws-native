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

	subnets := make([]resource.PropertyValue, int(count.NumberValue()))
	startPrefixLen, _ := network.Mask.Size()

	prefixLen := int(cidrBits.NumberValue()) + startPrefixLen
	if prefixLen > len(network.IP)*8 {
		protocol := "IP"
		switch len(network.IP) * 8 {
		case 32:
			protocol = "IPv4"
		case 128:
			protocol = "IPv6"
		}
		return nil, fmt.Errorf("cidrBits %d would extend prefix to %d bits, which is too long for an %s address", int(cidrBits.NumberValue()), prefixLen, protocol)
	}

	current, ok := gocidr.PreviousSubnet(network, prefixLen)
	if ok || !network.Contains(current.IP) {
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
