// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.
func LookupRobot(ctx *pulumi.Context, args *LookupRobotArgs, opts ...pulumi.InvokeOption) (*LookupRobotResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRobotResult
	err := ctx.Invoke("aws-native:robomaker:getRobot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRobotArgs struct {
	// The Amazon Resource Name (ARN) of the robot.
	Arn string `pulumi:"arn"`
}

type LookupRobotResult struct {
	// The Amazon Resource Name (ARN) of the robot.
	Arn *string `pulumi:"arn"`
	// A map that contains tag keys and tag values that are attached to the robot.
	Tags map[string]string `pulumi:"tags"`
}

func LookupRobotOutput(ctx *pulumi.Context, args LookupRobotOutputArgs, opts ...pulumi.InvokeOption) LookupRobotResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRobotResultOutput, error) {
			args := v.(LookupRobotArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:robomaker:getRobot", args, LookupRobotResultOutput{}, options).(LookupRobotResultOutput), nil
		}).(LookupRobotResultOutput)
}

type LookupRobotOutputArgs struct {
	// The Amazon Resource Name (ARN) of the robot.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupRobotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRobotArgs)(nil)).Elem()
}

type LookupRobotResultOutput struct{ *pulumi.OutputState }

func (LookupRobotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRobotResult)(nil)).Elem()
}

func (o LookupRobotResultOutput) ToLookupRobotResultOutput() LookupRobotResultOutput {
	return o
}

func (o LookupRobotResultOutput) ToLookupRobotResultOutputWithContext(ctx context.Context) LookupRobotResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the robot.
func (o LookupRobotResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRobotResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// A map that contains tag keys and tag values that are attached to the robot.
func (o LookupRobotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRobotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRobotResultOutput{})
}
