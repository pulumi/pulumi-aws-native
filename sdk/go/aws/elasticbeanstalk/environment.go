// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticBeanstalk::Environment
type Environment struct {
	pulumi.CustomResourceState

	// The name of the application that is associated with this environment.
	ApplicationName pulumi.StringOutput `pulumi:"applicationName"`
	// If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.
	CNAMEPrefix pulumi.StringPtrOutput `pulumi:"cNAMEPrefix"`
	// Your description for this environment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	EndpointURL pulumi.StringOutput    `pulumi:"endpointURL"`
	// A unique name for the environment.
	EnvironmentName pulumi.StringPtrOutput `pulumi:"environmentName"`
	// The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.
	OperationsRole pulumi.StringPtrOutput `pulumi:"operationsRole"`
	// Key-value pairs defining configuration options for this environment, such as the instance type.
	OptionSettings EnvironmentOptionSettingArrayOutput `pulumi:"optionSettings"`
	// The Amazon Resource Name (ARN) of the custom platform to use with the environment.
	PlatformArn pulumi.StringPtrOutput `pulumi:"platformArn"`
	// The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.
	SolutionStackName pulumi.StringPtrOutput `pulumi:"solutionStackName"`
	// Specifies the tags applied to resources in the environment.
	Tags EnvironmentTagArrayOutput `pulumi:"tags"`
	// The name of the Elastic Beanstalk configuration template to use with the environment.
	TemplateName pulumi.StringPtrOutput `pulumi:"templateName"`
	// Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.
	Tier EnvironmentTierPtrOutput `pulumi:"tier"`
	// The name of the application version to deploy.
	VersionLabel pulumi.StringPtrOutput `pulumi:"versionLabel"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	var resource Environment
	err := ctx.RegisterResource("aws-native:elasticbeanstalk:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("aws-native:elasticbeanstalk:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
}

type EnvironmentState struct {
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// The name of the application that is associated with this environment.
	ApplicationName string `pulumi:"applicationName"`
	// If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.
	CNAMEPrefix *string `pulumi:"cNAMEPrefix"`
	// Your description for this environment.
	Description *string `pulumi:"description"`
	// A unique name for the environment.
	EnvironmentName *string `pulumi:"environmentName"`
	// The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.
	OperationsRole *string `pulumi:"operationsRole"`
	// Key-value pairs defining configuration options for this environment, such as the instance type.
	OptionSettings []EnvironmentOptionSetting `pulumi:"optionSettings"`
	// The Amazon Resource Name (ARN) of the custom platform to use with the environment.
	PlatformArn *string `pulumi:"platformArn"`
	// The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.
	SolutionStackName *string `pulumi:"solutionStackName"`
	// Specifies the tags applied to resources in the environment.
	Tags []EnvironmentTag `pulumi:"tags"`
	// The name of the Elastic Beanstalk configuration template to use with the environment.
	TemplateName *string `pulumi:"templateName"`
	// Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.
	Tier *EnvironmentTier `pulumi:"tier"`
	// The name of the application version to deploy.
	VersionLabel *string `pulumi:"versionLabel"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// The name of the application that is associated with this environment.
	ApplicationName pulumi.StringInput
	// If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.
	CNAMEPrefix pulumi.StringPtrInput
	// Your description for this environment.
	Description pulumi.StringPtrInput
	// A unique name for the environment.
	EnvironmentName pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.
	OperationsRole pulumi.StringPtrInput
	// Key-value pairs defining configuration options for this environment, such as the instance type.
	OptionSettings EnvironmentOptionSettingArrayInput
	// The Amazon Resource Name (ARN) of the custom platform to use with the environment.
	PlatformArn pulumi.StringPtrInput
	// The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.
	SolutionStackName pulumi.StringPtrInput
	// Specifies the tags applied to resources in the environment.
	Tags EnvironmentTagArrayInput
	// The name of the Elastic Beanstalk configuration template to use with the environment.
	TemplateName pulumi.StringPtrInput
	// Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.
	Tier EnvironmentTierPtrInput
	// The name of the application version to deploy.
	VersionLabel pulumi.StringPtrInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

// The name of the application that is associated with this environment.
func (o EnvironmentOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

// If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.
func (o EnvironmentOutput) CNAMEPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.CNAMEPrefix }).(pulumi.StringPtrOutput)
}

// Your description for this environment.
func (o EnvironmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnvironmentOutput) EndpointURL() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.EndpointURL }).(pulumi.StringOutput)
}

// A unique name for the environment.
func (o EnvironmentOutput) EnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.EnvironmentName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.
func (o EnvironmentOutput) OperationsRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.OperationsRole }).(pulumi.StringPtrOutput)
}

// Key-value pairs defining configuration options for this environment, such as the instance type.
func (o EnvironmentOutput) OptionSettings() EnvironmentOptionSettingArrayOutput {
	return o.ApplyT(func(v *Environment) EnvironmentOptionSettingArrayOutput { return v.OptionSettings }).(EnvironmentOptionSettingArrayOutput)
}

// The Amazon Resource Name (ARN) of the custom platform to use with the environment.
func (o EnvironmentOutput) PlatformArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.PlatformArn }).(pulumi.StringPtrOutput)
}

// The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.
func (o EnvironmentOutput) SolutionStackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.SolutionStackName }).(pulumi.StringPtrOutput)
}

// Specifies the tags applied to resources in the environment.
func (o EnvironmentOutput) Tags() EnvironmentTagArrayOutput {
	return o.ApplyT(func(v *Environment) EnvironmentTagArrayOutput { return v.Tags }).(EnvironmentTagArrayOutput)
}

// The name of the Elastic Beanstalk configuration template to use with the environment.
func (o EnvironmentOutput) TemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.TemplateName }).(pulumi.StringPtrOutput)
}

// Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.
func (o EnvironmentOutput) Tier() EnvironmentTierPtrOutput {
	return o.ApplyT(func(v *Environment) EnvironmentTierPtrOutput { return v.Tier }).(EnvironmentTierPtrOutput)
}

// The name of the application version to deploy.
func (o EnvironmentOutput) VersionLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.VersionLabel }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentInput)(nil)).Elem(), &Environment{})
	pulumi.RegisterOutputType(EnvironmentOutput{})
}