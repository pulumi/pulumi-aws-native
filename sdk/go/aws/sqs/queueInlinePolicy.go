// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for SQS QueueInlinePolicy
type QueueInlinePolicy struct {
	pulumi.CustomResourceState

	// A policy document that contains permissions to add to the specified SQS queue
	PolicyDocument pulumi.AnyOutput `pulumi:"policyDocument"`
	// The URL of the SQS queue.
	Queue pulumi.StringOutput `pulumi:"queue"`
}

// NewQueueInlinePolicy registers a new resource with the given unique name, arguments, and options.
func NewQueueInlinePolicy(ctx *pulumi.Context,
	name string, args *QueueInlinePolicyArgs, opts ...pulumi.ResourceOption) (*QueueInlinePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	if args.Queue == nil {
		return nil, errors.New("invalid value for required argument 'Queue'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QueueInlinePolicy
	err := ctx.RegisterResource("aws-native:sqs:QueueInlinePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueueInlinePolicy gets an existing QueueInlinePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueueInlinePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueInlinePolicyState, opts ...pulumi.ResourceOption) (*QueueInlinePolicy, error) {
	var resource QueueInlinePolicy
	err := ctx.ReadResource("aws-native:sqs:QueueInlinePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueueInlinePolicy resources.
type queueInlinePolicyState struct {
}

type QueueInlinePolicyState struct {
}

func (QueueInlinePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueInlinePolicyState)(nil)).Elem()
}

type queueInlinePolicyArgs struct {
	// A policy document that contains permissions to add to the specified SQS queue
	PolicyDocument interface{} `pulumi:"policyDocument"`
	// The URL of the SQS queue.
	Queue string `pulumi:"queue"`
}

// The set of arguments for constructing a QueueInlinePolicy resource.
type QueueInlinePolicyArgs struct {
	// A policy document that contains permissions to add to the specified SQS queue
	PolicyDocument pulumi.Input
	// The URL of the SQS queue.
	Queue pulumi.StringInput
}

func (QueueInlinePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueInlinePolicyArgs)(nil)).Elem()
}

type QueueInlinePolicyInput interface {
	pulumi.Input

	ToQueueInlinePolicyOutput() QueueInlinePolicyOutput
	ToQueueInlinePolicyOutputWithContext(ctx context.Context) QueueInlinePolicyOutput
}

func (*QueueInlinePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueInlinePolicy)(nil)).Elem()
}

func (i *QueueInlinePolicy) ToQueueInlinePolicyOutput() QueueInlinePolicyOutput {
	return i.ToQueueInlinePolicyOutputWithContext(context.Background())
}

func (i *QueueInlinePolicy) ToQueueInlinePolicyOutputWithContext(ctx context.Context) QueueInlinePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueInlinePolicyOutput)
}

type QueueInlinePolicyOutput struct{ *pulumi.OutputState }

func (QueueInlinePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueInlinePolicy)(nil)).Elem()
}

func (o QueueInlinePolicyOutput) ToQueueInlinePolicyOutput() QueueInlinePolicyOutput {
	return o
}

func (o QueueInlinePolicyOutput) ToQueueInlinePolicyOutputWithContext(ctx context.Context) QueueInlinePolicyOutput {
	return o
}

// A policy document that contains permissions to add to the specified SQS queue
func (o QueueInlinePolicyOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v *QueueInlinePolicy) pulumi.AnyOutput { return v.PolicyDocument }).(pulumi.AnyOutput)
}

// The URL of the SQS queue.
func (o QueueInlinePolicyOutput) Queue() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueInlinePolicy) pulumi.StringOutput { return v.Queue }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueInlinePolicyInput)(nil)).Elem(), &QueueInlinePolicy{})
	pulumi.RegisterOutputType(QueueInlinePolicyOutput{})
}