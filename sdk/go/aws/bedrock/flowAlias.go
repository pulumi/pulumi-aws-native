// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bedrock

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Bedrock::FlowAlias Resource Type
type FlowAlias struct {
	pulumi.CustomResourceState

	// Arn of the Flow Alias
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Id for a Flow Alias generated at the server side.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The configuration that specifies how nodes in the flow are executed concurrently.
	ConcurrencyConfiguration FlowAliasConcurrencyConfigurationPtrOutput `pulumi:"concurrencyConfiguration"`
	// Time Stamp.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the Resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Arn representation of the Flow
	FlowArn pulumi.StringOutput `pulumi:"flowArn"`
	// Identifier for a flow resource.
	FlowId pulumi.StringOutput `pulumi:"flowId"`
	// Name for a resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Routing configuration for a Flow alias.
	RoutingConfiguration FlowAliasRoutingConfigurationListItemArrayOutput `pulumi:"routingConfiguration"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Time Stamp.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewFlowAlias registers a new resource with the given unique name, arguments, and options.
func NewFlowAlias(ctx *pulumi.Context,
	name string, args *FlowAliasArgs, opts ...pulumi.ResourceOption) (*FlowAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FlowArn == nil {
		return nil, errors.New("invalid value for required argument 'FlowArn'")
	}
	if args.RoutingConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'RoutingConfiguration'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"flowArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FlowAlias
	err := ctx.RegisterResource("aws-native:bedrock:FlowAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowAlias gets an existing FlowAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowAliasState, opts ...pulumi.ResourceOption) (*FlowAlias, error) {
	var resource FlowAlias
	err := ctx.ReadResource("aws-native:bedrock:FlowAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowAlias resources.
type flowAliasState struct {
}

type FlowAliasState struct {
}

func (FlowAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowAliasState)(nil)).Elem()
}

type flowAliasArgs struct {
	// The configuration that specifies how nodes in the flow are executed concurrently.
	ConcurrencyConfiguration *FlowAliasConcurrencyConfiguration `pulumi:"concurrencyConfiguration"`
	// Description of the Resource.
	Description *string `pulumi:"description"`
	// Arn representation of the Flow
	FlowArn string `pulumi:"flowArn"`
	// Name for a resource.
	Name *string `pulumi:"name"`
	// Routing configuration for a Flow alias.
	RoutingConfiguration []FlowAliasRoutingConfigurationListItem `pulumi:"routingConfiguration"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FlowAlias resource.
type FlowAliasArgs struct {
	// The configuration that specifies how nodes in the flow are executed concurrently.
	ConcurrencyConfiguration FlowAliasConcurrencyConfigurationPtrInput
	// Description of the Resource.
	Description pulumi.StringPtrInput
	// Arn representation of the Flow
	FlowArn pulumi.StringInput
	// Name for a resource.
	Name pulumi.StringPtrInput
	// Routing configuration for a Flow alias.
	RoutingConfiguration FlowAliasRoutingConfigurationListItemArrayInput
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	Tags pulumi.StringMapInput
}

func (FlowAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowAliasArgs)(nil)).Elem()
}

type FlowAliasInput interface {
	pulumi.Input

	ToFlowAliasOutput() FlowAliasOutput
	ToFlowAliasOutputWithContext(ctx context.Context) FlowAliasOutput
}

func (*FlowAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowAlias)(nil)).Elem()
}

func (i *FlowAlias) ToFlowAliasOutput() FlowAliasOutput {
	return i.ToFlowAliasOutputWithContext(context.Background())
}

func (i *FlowAlias) ToFlowAliasOutputWithContext(ctx context.Context) FlowAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowAliasOutput)
}

type FlowAliasOutput struct{ *pulumi.OutputState }

func (FlowAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowAlias)(nil)).Elem()
}

func (o FlowAliasOutput) ToFlowAliasOutput() FlowAliasOutput {
	return o
}

func (o FlowAliasOutput) ToFlowAliasOutputWithContext(ctx context.Context) FlowAliasOutput {
	return o
}

// Arn of the Flow Alias
func (o FlowAliasOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowAlias) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Id for a Flow Alias generated at the server side.
func (o FlowAliasOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowAlias) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The configuration that specifies how nodes in the flow are executed concurrently.
func (o FlowAliasOutput) ConcurrencyConfiguration() FlowAliasConcurrencyConfigurationPtrOutput {
	return o.ApplyT(func(v *FlowAlias) FlowAliasConcurrencyConfigurationPtrOutput { return v.ConcurrencyConfiguration }).(FlowAliasConcurrencyConfigurationPtrOutput)
}

// Time Stamp.
func (o FlowAliasOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowAlias) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description of the Resource.
func (o FlowAliasOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowAlias) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Arn representation of the Flow
func (o FlowAliasOutput) FlowArn() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowAlias) pulumi.StringOutput { return v.FlowArn }).(pulumi.StringOutput)
}

// Identifier for a flow resource.
func (o FlowAliasOutput) FlowId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowAlias) pulumi.StringOutput { return v.FlowId }).(pulumi.StringOutput)
}

// Name for a resource.
func (o FlowAliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowAlias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Routing configuration for a Flow alias.
func (o FlowAliasOutput) RoutingConfiguration() FlowAliasRoutingConfigurationListItemArrayOutput {
	return o.ApplyT(func(v *FlowAlias) FlowAliasRoutingConfigurationListItemArrayOutput { return v.RoutingConfiguration }).(FlowAliasRoutingConfigurationListItemArrayOutput)
}

// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
//
// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
func (o FlowAliasOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FlowAlias) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Time Stamp.
func (o FlowAliasOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowAlias) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowAliasInput)(nil)).Elem(), &FlowAlias{})
	pulumi.RegisterOutputType(FlowAliasOutput{})
}
