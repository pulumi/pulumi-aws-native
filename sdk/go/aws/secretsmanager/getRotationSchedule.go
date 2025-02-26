// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SecretsManager::RotationSchedule
func LookupRotationSchedule(ctx *pulumi.Context, args *LookupRotationScheduleArgs, opts ...pulumi.InvokeOption) (*LookupRotationScheduleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRotationScheduleResult
	err := ctx.Invoke("aws-native:secretsmanager:getRotationSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRotationScheduleArgs struct {
	// The ARN of the secret.
	Id string `pulumi:"id"`
}

type LookupRotationScheduleResult struct {
	// The ARN of the secret.
	Id *string `pulumi:"id"`
	// The ARN of an existing Lambda rotation function. To specify a rotation function that is also defined in this template, use the Ref function.
	RotationLambdaArn *string `pulumi:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret.
	RotationRules *RotationScheduleRotationRules `pulumi:"rotationRules"`
}

func LookupRotationScheduleOutput(ctx *pulumi.Context, args LookupRotationScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupRotationScheduleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRotationScheduleResultOutput, error) {
			args := v.(LookupRotationScheduleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:secretsmanager:getRotationSchedule", args, LookupRotationScheduleResultOutput{}, options).(LookupRotationScheduleResultOutput), nil
		}).(LookupRotationScheduleResultOutput)
}

type LookupRotationScheduleOutputArgs struct {
	// The ARN of the secret.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupRotationScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRotationScheduleArgs)(nil)).Elem()
}

type LookupRotationScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupRotationScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRotationScheduleResult)(nil)).Elem()
}

func (o LookupRotationScheduleResultOutput) ToLookupRotationScheduleResultOutput() LookupRotationScheduleResultOutput {
	return o
}

func (o LookupRotationScheduleResultOutput) ToLookupRotationScheduleResultOutputWithContext(ctx context.Context) LookupRotationScheduleResultOutput {
	return o
}

// The ARN of the secret.
func (o LookupRotationScheduleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRotationScheduleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The ARN of an existing Lambda rotation function. To specify a rotation function that is also defined in this template, use the Ref function.
func (o LookupRotationScheduleResultOutput) RotationLambdaArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRotationScheduleResult) *string { return v.RotationLambdaArn }).(pulumi.StringPtrOutput)
}

// A structure that defines the rotation configuration for this secret.
func (o LookupRotationScheduleResultOutput) RotationRules() RotationScheduleRotationRulesPtrOutput {
	return o.ApplyT(func(v LookupRotationScheduleResult) *RotationScheduleRotationRules { return v.RotationRules }).(RotationScheduleRotationRulesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRotationScheduleResultOutput{})
}
