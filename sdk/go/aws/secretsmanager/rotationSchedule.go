// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SecretsManager::RotationSchedule
type RotationSchedule struct {
	pulumi.CustomResourceState

	// The ARN of the secret.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// Creates a new Lambda rotation function based on one of the Secrets Manager rotation function templates. To use a rotation function that already exists, specify RotationLambdaARN instead.
	HostedRotationLambda RotationScheduleHostedRotationLambdaPtrOutput `pulumi:"hostedRotationLambda"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	RotateImmediatelyOnUpdate pulumi.BoolPtrOutput `pulumi:"rotateImmediatelyOnUpdate"`
	// The ARN of an existing Lambda rotation function. To specify a rotation function that is also defined in this template, use the Ref function.
	RotationLambdaArn pulumi.StringPtrOutput `pulumi:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret.
	RotationRules RotationScheduleRotationRulesPtrOutput `pulumi:"rotationRules"`
	// The ARN or name of the secret to rotate.
	SecretId pulumi.StringOutput `pulumi:"secretId"`
}

// NewRotationSchedule registers a new resource with the given unique name, arguments, and options.
func NewRotationSchedule(ctx *pulumi.Context,
	name string, args *RotationScheduleArgs, opts ...pulumi.ResourceOption) (*RotationSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecretId == nil {
		return nil, errors.New("invalid value for required argument 'SecretId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"secretId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RotationSchedule
	err := ctx.RegisterResource("aws-native:secretsmanager:RotationSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRotationSchedule gets an existing RotationSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRotationSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RotationScheduleState, opts ...pulumi.ResourceOption) (*RotationSchedule, error) {
	var resource RotationSchedule
	err := ctx.ReadResource("aws-native:secretsmanager:RotationSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RotationSchedule resources.
type rotationScheduleState struct {
}

type RotationScheduleState struct {
}

func (RotationScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rotationScheduleState)(nil)).Elem()
}

type rotationScheduleArgs struct {
	// Creates a new Lambda rotation function based on one of the Secrets Manager rotation function templates. To use a rotation function that already exists, specify RotationLambdaARN instead.
	HostedRotationLambda *RotationScheduleHostedRotationLambda `pulumi:"hostedRotationLambda"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	RotateImmediatelyOnUpdate *bool `pulumi:"rotateImmediatelyOnUpdate"`
	// The ARN of an existing Lambda rotation function. To specify a rotation function that is also defined in this template, use the Ref function.
	RotationLambdaArn *string `pulumi:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret.
	RotationRules *RotationScheduleRotationRules `pulumi:"rotationRules"`
	// The ARN or name of the secret to rotate.
	SecretId string `pulumi:"secretId"`
}

// The set of arguments for constructing a RotationSchedule resource.
type RotationScheduleArgs struct {
	// Creates a new Lambda rotation function based on one of the Secrets Manager rotation function templates. To use a rotation function that already exists, specify RotationLambdaARN instead.
	HostedRotationLambda RotationScheduleHostedRotationLambdaPtrInput
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	RotateImmediatelyOnUpdate pulumi.BoolPtrInput
	// The ARN of an existing Lambda rotation function. To specify a rotation function that is also defined in this template, use the Ref function.
	RotationLambdaArn pulumi.StringPtrInput
	// A structure that defines the rotation configuration for this secret.
	RotationRules RotationScheduleRotationRulesPtrInput
	// The ARN or name of the secret to rotate.
	SecretId pulumi.StringInput
}

func (RotationScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rotationScheduleArgs)(nil)).Elem()
}

type RotationScheduleInput interface {
	pulumi.Input

	ToRotationScheduleOutput() RotationScheduleOutput
	ToRotationScheduleOutputWithContext(ctx context.Context) RotationScheduleOutput
}

func (*RotationSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**RotationSchedule)(nil)).Elem()
}

func (i *RotationSchedule) ToRotationScheduleOutput() RotationScheduleOutput {
	return i.ToRotationScheduleOutputWithContext(context.Background())
}

func (i *RotationSchedule) ToRotationScheduleOutputWithContext(ctx context.Context) RotationScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RotationScheduleOutput)
}

type RotationScheduleOutput struct{ *pulumi.OutputState }

func (RotationScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RotationSchedule)(nil)).Elem()
}

func (o RotationScheduleOutput) ToRotationScheduleOutput() RotationScheduleOutput {
	return o
}

func (o RotationScheduleOutput) ToRotationScheduleOutputWithContext(ctx context.Context) RotationScheduleOutput {
	return o
}

// The ARN of the secret.
func (o RotationScheduleOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *RotationSchedule) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// Creates a new Lambda rotation function based on one of the Secrets Manager rotation function templates. To use a rotation function that already exists, specify RotationLambdaARN instead.
func (o RotationScheduleOutput) HostedRotationLambda() RotationScheduleHostedRotationLambdaPtrOutput {
	return o.ApplyT(func(v *RotationSchedule) RotationScheduleHostedRotationLambdaPtrOutput { return v.HostedRotationLambda }).(RotationScheduleHostedRotationLambdaPtrOutput)
}

// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
func (o RotationScheduleOutput) RotateImmediatelyOnUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RotationSchedule) pulumi.BoolPtrOutput { return v.RotateImmediatelyOnUpdate }).(pulumi.BoolPtrOutput)
}

// The ARN of an existing Lambda rotation function. To specify a rotation function that is also defined in this template, use the Ref function.
func (o RotationScheduleOutput) RotationLambdaArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RotationSchedule) pulumi.StringPtrOutput { return v.RotationLambdaArn }).(pulumi.StringPtrOutput)
}

// A structure that defines the rotation configuration for this secret.
func (o RotationScheduleOutput) RotationRules() RotationScheduleRotationRulesPtrOutput {
	return o.ApplyT(func(v *RotationSchedule) RotationScheduleRotationRulesPtrOutput { return v.RotationRules }).(RotationScheduleRotationRulesPtrOutput)
}

// The ARN or name of the secret to rotate.
func (o RotationScheduleOutput) SecretId() pulumi.StringOutput {
	return o.ApplyT(func(v *RotationSchedule) pulumi.StringOutput { return v.SecretId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RotationScheduleInput)(nil)).Elem(), &RotationSchedule{})
	pulumi.RegisterOutputType(RotationScheduleOutput{})
}
