// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticBeanstalk::ConfigurationTemplate
//
// Deprecated: ConfigurationTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ConfigurationTemplate struct {
	pulumi.CustomResourceState

	ApplicationName     pulumi.StringOutput                                        `pulumi:"applicationName"`
	Description         pulumi.StringPtrOutput                                     `pulumi:"description"`
	EnvironmentId       pulumi.StringPtrOutput                                     `pulumi:"environmentId"`
	OptionSettings      ConfigurationTemplateConfigurationOptionSettingArrayOutput `pulumi:"optionSettings"`
	PlatformArn         pulumi.StringPtrOutput                                     `pulumi:"platformArn"`
	SolutionStackName   pulumi.StringPtrOutput                                     `pulumi:"solutionStackName"`
	SourceConfiguration ConfigurationTemplateSourceConfigurationPtrOutput          `pulumi:"sourceConfiguration"`
}

// NewConfigurationTemplate registers a new resource with the given unique name, arguments, and options.
func NewConfigurationTemplate(ctx *pulumi.Context,
	name string, args *ConfigurationTemplateArgs, opts ...pulumi.ResourceOption) (*ConfigurationTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	var resource ConfigurationTemplate
	err := ctx.RegisterResource("aws-native:elasticbeanstalk:ConfigurationTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationTemplate gets an existing ConfigurationTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationTemplateState, opts ...pulumi.ResourceOption) (*ConfigurationTemplate, error) {
	var resource ConfigurationTemplate
	err := ctx.ReadResource("aws-native:elasticbeanstalk:ConfigurationTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationTemplate resources.
type configurationTemplateState struct {
}

type ConfigurationTemplateState struct {
}

func (ConfigurationTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationTemplateState)(nil)).Elem()
}

type configurationTemplateArgs struct {
	ApplicationName     string                                            `pulumi:"applicationName"`
	Description         *string                                           `pulumi:"description"`
	EnvironmentId       *string                                           `pulumi:"environmentId"`
	OptionSettings      []ConfigurationTemplateConfigurationOptionSetting `pulumi:"optionSettings"`
	PlatformArn         *string                                           `pulumi:"platformArn"`
	SolutionStackName   *string                                           `pulumi:"solutionStackName"`
	SourceConfiguration *ConfigurationTemplateSourceConfiguration         `pulumi:"sourceConfiguration"`
}

// The set of arguments for constructing a ConfigurationTemplate resource.
type ConfigurationTemplateArgs struct {
	ApplicationName     pulumi.StringInput
	Description         pulumi.StringPtrInput
	EnvironmentId       pulumi.StringPtrInput
	OptionSettings      ConfigurationTemplateConfigurationOptionSettingArrayInput
	PlatformArn         pulumi.StringPtrInput
	SolutionStackName   pulumi.StringPtrInput
	SourceConfiguration ConfigurationTemplateSourceConfigurationPtrInput
}

func (ConfigurationTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationTemplateArgs)(nil)).Elem()
}

type ConfigurationTemplateInput interface {
	pulumi.Input

	ToConfigurationTemplateOutput() ConfigurationTemplateOutput
	ToConfigurationTemplateOutputWithContext(ctx context.Context) ConfigurationTemplateOutput
}

func (*ConfigurationTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationTemplate)(nil)).Elem()
}

func (i *ConfigurationTemplate) ToConfigurationTemplateOutput() ConfigurationTemplateOutput {
	return i.ToConfigurationTemplateOutputWithContext(context.Background())
}

func (i *ConfigurationTemplate) ToConfigurationTemplateOutputWithContext(ctx context.Context) ConfigurationTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationTemplateOutput)
}

type ConfigurationTemplateOutput struct{ *pulumi.OutputState }

func (ConfigurationTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationTemplate)(nil)).Elem()
}

func (o ConfigurationTemplateOutput) ToConfigurationTemplateOutput() ConfigurationTemplateOutput {
	return o
}

func (o ConfigurationTemplateOutput) ToConfigurationTemplateOutputWithContext(ctx context.Context) ConfigurationTemplateOutput {
	return o
}

func (o ConfigurationTemplateOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationTemplate) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o ConfigurationTemplateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationTemplate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConfigurationTemplateOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationTemplate) pulumi.StringPtrOutput { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o ConfigurationTemplateOutput) OptionSettings() ConfigurationTemplateConfigurationOptionSettingArrayOutput {
	return o.ApplyT(func(v *ConfigurationTemplate) ConfigurationTemplateConfigurationOptionSettingArrayOutput {
		return v.OptionSettings
	}).(ConfigurationTemplateConfigurationOptionSettingArrayOutput)
}

func (o ConfigurationTemplateOutput) PlatformArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationTemplate) pulumi.StringPtrOutput { return v.PlatformArn }).(pulumi.StringPtrOutput)
}

func (o ConfigurationTemplateOutput) SolutionStackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationTemplate) pulumi.StringPtrOutput { return v.SolutionStackName }).(pulumi.StringPtrOutput)
}

func (o ConfigurationTemplateOutput) SourceConfiguration() ConfigurationTemplateSourceConfigurationPtrOutput {
	return o.ApplyT(func(v *ConfigurationTemplate) ConfigurationTemplateSourceConfigurationPtrOutput {
		return v.SourceConfiguration
	}).(ConfigurationTemplateSourceConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationTemplateInput)(nil)).Elem(), &ConfigurationTemplate{})
	pulumi.RegisterOutputType(ConfigurationTemplateOutput{})
}