// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::ConnectorDefinition
//
// Deprecated: ConnectorDefinition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ConnectorDefinition struct {
	pulumi.CustomResourceState

	Arn              pulumi.StringOutput                     `pulumi:"arn"`
	InitialVersion   ConnectorDefinitionVersionTypePtrOutput `pulumi:"initialVersion"`
	LatestVersionArn pulumi.StringOutput                     `pulumi:"latestVersionArn"`
	Name             pulumi.StringOutput                     `pulumi:"name"`
	Tags             pulumi.AnyOutput                        `pulumi:"tags"`
}

// NewConnectorDefinition registers a new resource with the given unique name, arguments, and options.
func NewConnectorDefinition(ctx *pulumi.Context,
	name string, args *ConnectorDefinitionArgs, opts ...pulumi.ResourceOption) (*ConnectorDefinition, error) {
	if args == nil {
		args = &ConnectorDefinitionArgs{}
	}

	var resource ConnectorDefinition
	err := ctx.RegisterResource("aws-native:greengrass:ConnectorDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectorDefinition gets an existing ConnectorDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectorDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorDefinitionState, opts ...pulumi.ResourceOption) (*ConnectorDefinition, error) {
	var resource ConnectorDefinition
	err := ctx.ReadResource("aws-native:greengrass:ConnectorDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectorDefinition resources.
type connectorDefinitionState struct {
}

type ConnectorDefinitionState struct {
}

func (ConnectorDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorDefinitionState)(nil)).Elem()
}

type connectorDefinitionArgs struct {
	InitialVersion *ConnectorDefinitionVersionType `pulumi:"initialVersion"`
	Name           *string                         `pulumi:"name"`
	Tags           interface{}                     `pulumi:"tags"`
}

// The set of arguments for constructing a ConnectorDefinition resource.
type ConnectorDefinitionArgs struct {
	InitialVersion ConnectorDefinitionVersionTypePtrInput
	Name           pulumi.StringPtrInput
	Tags           pulumi.Input
}

func (ConnectorDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorDefinitionArgs)(nil)).Elem()
}

type ConnectorDefinitionInput interface {
	pulumi.Input

	ToConnectorDefinitionOutput() ConnectorDefinitionOutput
	ToConnectorDefinitionOutputWithContext(ctx context.Context) ConnectorDefinitionOutput
}

func (*ConnectorDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorDefinition)(nil)).Elem()
}

func (i *ConnectorDefinition) ToConnectorDefinitionOutput() ConnectorDefinitionOutput {
	return i.ToConnectorDefinitionOutputWithContext(context.Background())
}

func (i *ConnectorDefinition) ToConnectorDefinitionOutputWithContext(ctx context.Context) ConnectorDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorDefinitionOutput)
}

type ConnectorDefinitionOutput struct{ *pulumi.OutputState }

func (ConnectorDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorDefinition)(nil)).Elem()
}

func (o ConnectorDefinitionOutput) ToConnectorDefinitionOutput() ConnectorDefinitionOutput {
	return o
}

func (o ConnectorDefinitionOutput) ToConnectorDefinitionOutputWithContext(ctx context.Context) ConnectorDefinitionOutput {
	return o
}

func (o ConnectorDefinitionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorDefinition) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ConnectorDefinitionOutput) InitialVersion() ConnectorDefinitionVersionTypePtrOutput {
	return o.ApplyT(func(v *ConnectorDefinition) ConnectorDefinitionVersionTypePtrOutput { return v.InitialVersion }).(ConnectorDefinitionVersionTypePtrOutput)
}

func (o ConnectorDefinitionOutput) LatestVersionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorDefinition) pulumi.StringOutput { return v.LatestVersionArn }).(pulumi.StringOutput)
}

func (o ConnectorDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectorDefinitionOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *ConnectorDefinition) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorDefinitionInput)(nil)).Elem(), &ConnectorDefinition{})
	pulumi.RegisterOutputType(ConnectorDefinitionOutput{})
}