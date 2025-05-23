// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bedrock

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Bedrock::IntelligentPromptRouter Resource Type
type IntelligentPromptRouter struct {
	pulumi.CustomResourceState

	// Time Stamp
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the Prompt Router.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The default model to use when the routing criteria is not met.
	FallbackModel IntelligentPromptRouterPromptRouterTargetModelOutput `pulumi:"fallbackModel"`
	// List of model configuration
	Models IntelligentPromptRouterPromptRouterTargetModelArrayOutput `pulumi:"models"`
	// Arn of the Prompt Router.
	PromptRouterArn pulumi.StringOutput `pulumi:"promptRouterArn"`
	// Name of the Prompt Router.
	PromptRouterName pulumi.StringOutput `pulumi:"promptRouterName"`
	// Routing criteria for a prompt router.
	RoutingCriteria IntelligentPromptRouterRoutingCriteriaOutput `pulumi:"routingCriteria"`
	// The router's status.
	Status IntelligentPromptRouterPromptRouterStatusOutput `pulumi:"status"`
	// List of Tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The router's type.
	Type IntelligentPromptRouterPromptRouterTypeOutput `pulumi:"type"`
	// Time Stamp
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewIntelligentPromptRouter registers a new resource with the given unique name, arguments, and options.
func NewIntelligentPromptRouter(ctx *pulumi.Context,
	name string, args *IntelligentPromptRouterArgs, opts ...pulumi.ResourceOption) (*IntelligentPromptRouter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FallbackModel == nil {
		return nil, errors.New("invalid value for required argument 'FallbackModel'")
	}
	if args.Models == nil {
		return nil, errors.New("invalid value for required argument 'Models'")
	}
	if args.RoutingCriteria == nil {
		return nil, errors.New("invalid value for required argument 'RoutingCriteria'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"description",
		"fallbackModel",
		"models[*]",
		"promptRouterName",
		"routingCriteria",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IntelligentPromptRouter
	err := ctx.RegisterResource("aws-native:bedrock:IntelligentPromptRouter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntelligentPromptRouter gets an existing IntelligentPromptRouter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntelligentPromptRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntelligentPromptRouterState, opts ...pulumi.ResourceOption) (*IntelligentPromptRouter, error) {
	var resource IntelligentPromptRouter
	err := ctx.ReadResource("aws-native:bedrock:IntelligentPromptRouter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntelligentPromptRouter resources.
type intelligentPromptRouterState struct {
}

type IntelligentPromptRouterState struct {
}

func (IntelligentPromptRouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*intelligentPromptRouterState)(nil)).Elem()
}

type intelligentPromptRouterArgs struct {
	// Description of the Prompt Router.
	Description *string `pulumi:"description"`
	// The default model to use when the routing criteria is not met.
	FallbackModel IntelligentPromptRouterPromptRouterTargetModel `pulumi:"fallbackModel"`
	// List of model configuration
	Models []IntelligentPromptRouterPromptRouterTargetModel `pulumi:"models"`
	// Name of the Prompt Router.
	PromptRouterName *string `pulumi:"promptRouterName"`
	// Routing criteria for a prompt router.
	RoutingCriteria IntelligentPromptRouterRoutingCriteria `pulumi:"routingCriteria"`
	// List of Tags
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a IntelligentPromptRouter resource.
type IntelligentPromptRouterArgs struct {
	// Description of the Prompt Router.
	Description pulumi.StringPtrInput
	// The default model to use when the routing criteria is not met.
	FallbackModel IntelligentPromptRouterPromptRouterTargetModelInput
	// List of model configuration
	Models IntelligentPromptRouterPromptRouterTargetModelArrayInput
	// Name of the Prompt Router.
	PromptRouterName pulumi.StringPtrInput
	// Routing criteria for a prompt router.
	RoutingCriteria IntelligentPromptRouterRoutingCriteriaInput
	// List of Tags
	Tags aws.TagArrayInput
}

func (IntelligentPromptRouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*intelligentPromptRouterArgs)(nil)).Elem()
}

type IntelligentPromptRouterInput interface {
	pulumi.Input

	ToIntelligentPromptRouterOutput() IntelligentPromptRouterOutput
	ToIntelligentPromptRouterOutputWithContext(ctx context.Context) IntelligentPromptRouterOutput
}

func (*IntelligentPromptRouter) ElementType() reflect.Type {
	return reflect.TypeOf((**IntelligentPromptRouter)(nil)).Elem()
}

func (i *IntelligentPromptRouter) ToIntelligentPromptRouterOutput() IntelligentPromptRouterOutput {
	return i.ToIntelligentPromptRouterOutputWithContext(context.Background())
}

func (i *IntelligentPromptRouter) ToIntelligentPromptRouterOutputWithContext(ctx context.Context) IntelligentPromptRouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntelligentPromptRouterOutput)
}

type IntelligentPromptRouterOutput struct{ *pulumi.OutputState }

func (IntelligentPromptRouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntelligentPromptRouter)(nil)).Elem()
}

func (o IntelligentPromptRouterOutput) ToIntelligentPromptRouterOutput() IntelligentPromptRouterOutput {
	return o
}

func (o IntelligentPromptRouterOutput) ToIntelligentPromptRouterOutputWithContext(ctx context.Context) IntelligentPromptRouterOutput {
	return o
}

// Time Stamp
func (o IntelligentPromptRouterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description of the Prompt Router.
func (o IntelligentPromptRouterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The default model to use when the routing criteria is not met.
func (o IntelligentPromptRouterOutput) FallbackModel() IntelligentPromptRouterPromptRouterTargetModelOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) IntelligentPromptRouterPromptRouterTargetModelOutput {
		return v.FallbackModel
	}).(IntelligentPromptRouterPromptRouterTargetModelOutput)
}

// List of model configuration
func (o IntelligentPromptRouterOutput) Models() IntelligentPromptRouterPromptRouterTargetModelArrayOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) IntelligentPromptRouterPromptRouterTargetModelArrayOutput {
		return v.Models
	}).(IntelligentPromptRouterPromptRouterTargetModelArrayOutput)
}

// Arn of the Prompt Router.
func (o IntelligentPromptRouterOutput) PromptRouterArn() pulumi.StringOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) pulumi.StringOutput { return v.PromptRouterArn }).(pulumi.StringOutput)
}

// Name of the Prompt Router.
func (o IntelligentPromptRouterOutput) PromptRouterName() pulumi.StringOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) pulumi.StringOutput { return v.PromptRouterName }).(pulumi.StringOutput)
}

// Routing criteria for a prompt router.
func (o IntelligentPromptRouterOutput) RoutingCriteria() IntelligentPromptRouterRoutingCriteriaOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) IntelligentPromptRouterRoutingCriteriaOutput {
		return v.RoutingCriteria
	}).(IntelligentPromptRouterRoutingCriteriaOutput)
}

// The router's status.
func (o IntelligentPromptRouterOutput) Status() IntelligentPromptRouterPromptRouterStatusOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) IntelligentPromptRouterPromptRouterStatusOutput { return v.Status }).(IntelligentPromptRouterPromptRouterStatusOutput)
}

// List of Tags
func (o IntelligentPromptRouterOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The router's type.
func (o IntelligentPromptRouterOutput) Type() IntelligentPromptRouterPromptRouterTypeOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) IntelligentPromptRouterPromptRouterTypeOutput { return v.Type }).(IntelligentPromptRouterPromptRouterTypeOutput)
}

// Time Stamp
func (o IntelligentPromptRouterOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IntelligentPromptRouter) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntelligentPromptRouterInput)(nil)).Elem(), &IntelligentPromptRouter{})
	pulumi.RegisterOutputType(IntelligentPromptRouterOutput{})
}
