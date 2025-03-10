// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppConfig::Deployment
type Deployment struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The configuration profile ID.
	ConfigurationProfileId pulumi.StringOutput `pulumi:"configurationProfileId"`
	// The configuration version to deploy. If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label. For all other configurations, you must specify the version number.
	ConfigurationVersion pulumi.StringOutput `pulumi:"configurationVersion"`
	// The sequence number of the deployment.
	DeploymentNumber pulumi.StringOutput `pulumi:"deploymentNumber"`
	// The deployment strategy ID.
	DeploymentStrategyId pulumi.StringOutput `pulumi:"deploymentStrategyId"`
	// A description of the deployment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A map of dynamic extension parameter names to values to pass to associated extensions with `PRE_START_DEPLOYMENT` actions.
	DynamicExtensionParameters DeploymentDynamicExtensionParametersArrayOutput `pulumi:"dynamicExtensionParameters"`
	// The environment ID.
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.
	KmsKeyIdentifier pulumi.StringPtrOutput `pulumi:"kmsKeyIdentifier"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.CreateOnlyTagArrayOutput `pulumi:"tags"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.ConfigurationProfileId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationProfileId'")
	}
	if args.ConfigurationVersion == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationVersion'")
	}
	if args.DeploymentStrategyId == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentStrategyId'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"applicationId",
		"configurationProfileId",
		"configurationVersion",
		"deploymentStrategyId",
		"description",
		"dynamicExtensionParameters[*]",
		"environmentId",
		"kmsKeyIdentifier",
		"tags[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Deployment
	err := ctx.RegisterResource("aws-native:appconfig:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("aws-native:appconfig:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
}

type DeploymentState struct {
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// The configuration profile ID.
	ConfigurationProfileId string `pulumi:"configurationProfileId"`
	// The configuration version to deploy. If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label. For all other configurations, you must specify the version number.
	ConfigurationVersion string `pulumi:"configurationVersion"`
	// The deployment strategy ID.
	DeploymentStrategyId string `pulumi:"deploymentStrategyId"`
	// A description of the deployment.
	Description *string `pulumi:"description"`
	// A map of dynamic extension parameter names to values to pass to associated extensions with `PRE_START_DEPLOYMENT` actions.
	DynamicExtensionParameters []DeploymentDynamicExtensionParameters `pulumi:"dynamicExtensionParameters"`
	// The environment ID.
	EnvironmentId string `pulumi:"environmentId"`
	// The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.
	KmsKeyIdentifier *string `pulumi:"kmsKeyIdentifier"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.CreateOnlyTag `pulumi:"tags"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// The configuration profile ID.
	ConfigurationProfileId pulumi.StringInput
	// The configuration version to deploy. If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label. For all other configurations, you must specify the version number.
	ConfigurationVersion pulumi.StringInput
	// The deployment strategy ID.
	DeploymentStrategyId pulumi.StringInput
	// A description of the deployment.
	Description pulumi.StringPtrInput
	// A map of dynamic extension parameter names to values to pass to associated extensions with `PRE_START_DEPLOYMENT` actions.
	DynamicExtensionParameters DeploymentDynamicExtensionParametersArrayInput
	// The environment ID.
	EnvironmentId pulumi.StringInput
	// The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.
	KmsKeyIdentifier pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.CreateOnlyTagArrayInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}

type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput
}

func (*Deployment) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

type DeploymentOutput struct{ *pulumi.OutputState }

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

// The application ID.
func (o DeploymentOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// The configuration profile ID.
func (o DeploymentOutput) ConfigurationProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.ConfigurationProfileId }).(pulumi.StringOutput)
}

// The configuration version to deploy. If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label. For all other configurations, you must specify the version number.
func (o DeploymentOutput) ConfigurationVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.ConfigurationVersion }).(pulumi.StringOutput)
}

// The sequence number of the deployment.
func (o DeploymentOutput) DeploymentNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.DeploymentNumber }).(pulumi.StringOutput)
}

// The deployment strategy ID.
func (o DeploymentOutput) DeploymentStrategyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.DeploymentStrategyId }).(pulumi.StringOutput)
}

// A description of the deployment.
func (o DeploymentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A map of dynamic extension parameter names to values to pass to associated extensions with `PRE_START_DEPLOYMENT` actions.
func (o DeploymentOutput) DynamicExtensionParameters() DeploymentDynamicExtensionParametersArrayOutput {
	return o.ApplyT(func(v *Deployment) DeploymentDynamicExtensionParametersArrayOutput {
		return v.DynamicExtensionParameters
	}).(DeploymentDynamicExtensionParametersArrayOutput)
}

// The environment ID.
func (o DeploymentOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.
func (o DeploymentOutput) KmsKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.KmsKeyIdentifier }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o DeploymentOutput) Tags() aws.CreateOnlyTagArrayOutput {
	return o.ApplyT(func(v *Deployment) aws.CreateOnlyTagArrayOutput { return v.Tags }).(aws.CreateOnlyTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentInput)(nil)).Elem(), &Deployment{})
	pulumi.RegisterOutputType(DeploymentOutput{})
}
