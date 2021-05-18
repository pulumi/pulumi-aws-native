package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func (p *cfnProvider) getSSMParameter(ctx context.Context, inputs resource.PropertyMap) (*ssm.Parameter, error) {
	name, ok := inputs["name"]
	if !ok {
		return nil, fmt.Errorf("missing required parameter 'name'")
	}
	if !name.IsString() {
		return nil, fmt.Errorf("'name' must be a string")
	}

	out, err := p.ssm.GetParameterWithContext(ctx, &ssm.GetParameterInput{
		Name: aws.String(name.StringValue()),
	})
	if err != nil {
		return nil, err
	}
	return out.Parameter, nil
}

func (p *cfnProvider) getSSMParameterString(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	param, err := p.getSSMParameter(ctx, inputs)
	if err != nil {
		return nil, err
	}
	if aws.StringValue(param.Type) != "String" {
		return nil, fmt.Errorf("parameter '%v' is not a string", aws.StringValue(param.Name))
	}
	return resource.PropertyMap{
		"value": resource.NewStringProperty(aws.StringValue(param.Value)),
	}, nil
}

func (p *cfnProvider) getSSMParameterList(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	param, err := p.getSSMParameter(ctx, inputs)
	if err != nil {
		return nil, err
	}
	if aws.StringValue(param.Type) != "StringList" {
		return nil, fmt.Errorf("parameter '%v' is not a string list", aws.StringValue(param.Name))
	}
	return resource.NewPropertyMapFromMap(map[string]interface{}{
		"value": strings.Split(aws.StringValue(param.Value), ","),
	}), nil
}
