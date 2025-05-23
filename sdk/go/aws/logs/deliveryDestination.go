// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This structure contains information about one delivery destination in your account.
//
// A delivery destination is an AWS resource that represents an AWS service that logs can be sent to CloudWatch Logs, Amazon S3, are supported as Kinesis Data Firehose delivery destinations.
type DeliveryDestination struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) that uniquely identifies this delivery destination.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 51200
	DeliveryDestinationPolicy DeliveryDestinationDestinationPolicyPtrOutput `pulumi:"deliveryDestinationPolicy"`
	// Displays whether this delivery destination is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.
	DeliveryDestinationType pulumi.StringOutput `pulumi:"deliveryDestinationType"`
	// The ARN of the Amazon Web Services destination that this delivery destination represents. That Amazon Web Services destination can be a log group in CloudWatch Logs, an Amazon S3 bucket, or a delivery stream in Firehose.
	DestinationResourceArn pulumi.StringPtrOutput `pulumi:"destinationResourceArn"`
	// The name of this delivery destination.
	Name pulumi.StringOutput `pulumi:"name"`
	// The format of the logs that are sent to this delivery destination.
	OutputFormat pulumi.StringPtrOutput `pulumi:"outputFormat"`
	// The tags that have been assigned to this delivery destination.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewDeliveryDestination registers a new resource with the given unique name, arguments, and options.
func NewDeliveryDestination(ctx *pulumi.Context,
	name string, args *DeliveryDestinationArgs, opts ...pulumi.ResourceOption) (*DeliveryDestination, error) {
	if args == nil {
		args = &DeliveryDestinationArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"destinationResourceArn",
		"name",
		"outputFormat",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DeliveryDestination
	err := ctx.RegisterResource("aws-native:logs:DeliveryDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeliveryDestination gets an existing DeliveryDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeliveryDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeliveryDestinationState, opts ...pulumi.ResourceOption) (*DeliveryDestination, error) {
	var resource DeliveryDestination
	err := ctx.ReadResource("aws-native:logs:DeliveryDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeliveryDestination resources.
type deliveryDestinationState struct {
}

type DeliveryDestinationState struct {
}

func (DeliveryDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryDestinationState)(nil)).Elem()
}

type deliveryDestinationArgs struct {
	// IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 51200
	DeliveryDestinationPolicy *DeliveryDestinationDestinationPolicy `pulumi:"deliveryDestinationPolicy"`
	// The ARN of the Amazon Web Services destination that this delivery destination represents. That Amazon Web Services destination can be a log group in CloudWatch Logs, an Amazon S3 bucket, or a delivery stream in Firehose.
	DestinationResourceArn *string `pulumi:"destinationResourceArn"`
	// The name of this delivery destination.
	Name *string `pulumi:"name"`
	// The format of the logs that are sent to this delivery destination.
	OutputFormat *string `pulumi:"outputFormat"`
	// The tags that have been assigned to this delivery destination.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a DeliveryDestination resource.
type DeliveryDestinationArgs struct {
	// IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 51200
	DeliveryDestinationPolicy DeliveryDestinationDestinationPolicyPtrInput
	// The ARN of the Amazon Web Services destination that this delivery destination represents. That Amazon Web Services destination can be a log group in CloudWatch Logs, an Amazon S3 bucket, or a delivery stream in Firehose.
	DestinationResourceArn pulumi.StringPtrInput
	// The name of this delivery destination.
	Name pulumi.StringPtrInput
	// The format of the logs that are sent to this delivery destination.
	OutputFormat pulumi.StringPtrInput
	// The tags that have been assigned to this delivery destination.
	Tags aws.TagArrayInput
}

func (DeliveryDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryDestinationArgs)(nil)).Elem()
}

type DeliveryDestinationInput interface {
	pulumi.Input

	ToDeliveryDestinationOutput() DeliveryDestinationOutput
	ToDeliveryDestinationOutputWithContext(ctx context.Context) DeliveryDestinationOutput
}

func (*DeliveryDestination) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryDestination)(nil)).Elem()
}

func (i *DeliveryDestination) ToDeliveryDestinationOutput() DeliveryDestinationOutput {
	return i.ToDeliveryDestinationOutputWithContext(context.Background())
}

func (i *DeliveryDestination) ToDeliveryDestinationOutputWithContext(ctx context.Context) DeliveryDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryDestinationOutput)
}

type DeliveryDestinationOutput struct{ *pulumi.OutputState }

func (DeliveryDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryDestination)(nil)).Elem()
}

func (o DeliveryDestinationOutput) ToDeliveryDestinationOutput() DeliveryDestinationOutput {
	return o
}

func (o DeliveryDestinationOutput) ToDeliveryDestinationOutputWithContext(ctx context.Context) DeliveryDestinationOutput {
	return o
}

// The Amazon Resource Name (ARN) that uniquely identifies this delivery destination.
func (o DeliveryDestinationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryDestination) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.
//
// The policy must be in JSON string format.
//
// Length Constraints: Maximum length of 51200
func (o DeliveryDestinationOutput) DeliveryDestinationPolicy() DeliveryDestinationDestinationPolicyPtrOutput {
	return o.ApplyT(func(v *DeliveryDestination) DeliveryDestinationDestinationPolicyPtrOutput {
		return v.DeliveryDestinationPolicy
	}).(DeliveryDestinationDestinationPolicyPtrOutput)
}

// Displays whether this delivery destination is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.
func (o DeliveryDestinationOutput) DeliveryDestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryDestination) pulumi.StringOutput { return v.DeliveryDestinationType }).(pulumi.StringOutput)
}

// The ARN of the Amazon Web Services destination that this delivery destination represents. That Amazon Web Services destination can be a log group in CloudWatch Logs, an Amazon S3 bucket, or a delivery stream in Firehose.
func (o DeliveryDestinationOutput) DestinationResourceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeliveryDestination) pulumi.StringPtrOutput { return v.DestinationResourceArn }).(pulumi.StringPtrOutput)
}

// The name of this delivery destination.
func (o DeliveryDestinationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryDestination) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The format of the logs that are sent to this delivery destination.
func (o DeliveryDestinationOutput) OutputFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeliveryDestination) pulumi.StringPtrOutput { return v.OutputFormat }).(pulumi.StringPtrOutput)
}

// The tags that have been assigned to this delivery destination.
func (o DeliveryDestinationOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *DeliveryDestination) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeliveryDestinationInput)(nil)).Elem(), &DeliveryDestination{})
	pulumi.RegisterOutputType(DeliveryDestinationOutput{})
}
