package provider

import (
	"context"
	"fmt"
	"net"

	gocidr "github.com/apparentlymart/go-cidr/cidr"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func (p *cfnProvider) getAZs(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	region, ok := inputs["region"]
	if !ok {
		region = resource.NewStringProperty(p.region)
	} else if !region.IsString() {
		return nil, fmt.Errorf("'region' must be a string")
	}

	resp, err := p.ec2.DescribeAvailabilityZonesWithContext(ctx, &ec2.DescribeAvailabilityZonesInput{
		Filters: []*ec2.Filter{{
			Name:   aws.String("region-name"),
			Values: []*string{aws.String(region.String())},
		}},
	})
	if err != nil {
		return nil, err
	}

	azs := make([]interface{}, len(resp.AvailabilityZones))
	for i, az := range resp.AvailabilityZones {
		azs[i] = aws.StringValue(az.ZoneName)
	}
	return resource.NewPropertyMapFromMap(map[string]interface{}{
		"azs": azs,
	}), nil
}

func (p *cfnProvider) cidr(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	ipBlock, ok := inputs["ipBlock"]
	if !ok {
		return nil, fmt.Errorf("missing required property 'ipBlock'")
	}
	if !ipBlock.IsString() {
		return nil, fmt.Errorf("'ipBlock' must be a string")
	}

	count, ok := inputs["count"]
	if !ok {
		return nil, fmt.Errorf("mising required property 'count'")
	}
	if !count.IsNumber() || count.NumberValue() < 1 || count.NumberValue() > 256 {
		return nil, fmt.Errorf("'count' must be a number between 1 and 256")
	}

	cidrBits, ok := inputs["cidrBits"]
	if !ok {
		return nil, fmt.Errorf("mising required property 'cidrBits'")
	}
	if !cidrBits.IsNumber() || cidrBits.NumberValue() < 0 {
		return nil, fmt.Errorf("'cidrBits' must be a positive number")
	}

	_, network, err := net.ParseCIDR(ipBlock.StringValue())
	if err != nil {
		return nil, fmt.Errorf("invalid IP block: %s", err)
	}

	subnets := make([]resource.PropertyValue, int(count.NumberValue()))
	subnets[0] = resource.NewStringProperty(network.String())

	prefixLen := int(cidrBits.NumberValue())
	for i := 1; i < len(subnets)-1; i++ {
		subnet, ok := gocidr.NextSubnet(network, prefixLen)
		if !ok {
			return nil, fmt.Errorf("could not create %d subnets", len(subnets))
		}
		subnets[i], network = resource.NewStringProperty(subnet.String()), subnet
	}

	return resource.PropertyMap{
		"subnets": resource.NewArrayProperty(subnets),
	}, nil
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
	err := p.cfn.ListExportsPages(&cloudformation.ListExportsInput{}, func(page *cloudformation.ListExportsOutput, lastPage bool) bool {
		for _, export := range page.Exports {
			if aws.StringValue(export.Name) == name.StringValue() {
				value, ok = aws.StringValue(export.Value), true
				return false
			}
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("unknown export '%s'", name.StringValue())
	}

	return resource.PropertyMap{
		"value": resource.NewStringProperty(value),
	}, nil
}
