// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gameliftstreams

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::GameLiftStreams::Application Resource Type
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationResult
	err := ctx.Invoke("aws-native:gameliftstreams:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	// An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the application resource across all AWS Regions. For example:
	//
	// `arn:aws:gameliftstreams:us-west-2:123456789012:application/a-9ZY8X7Wv6` .
	Arn string `pulumi:"arn"`
}

type LookupApplicationResult struct {
	// An Amazon S3 URI to a bucket where you would like Amazon GameLift Streams to save application logs. Required if you specify one or more `ApplicationLogPaths` .
	ApplicationLogOutputUri *string `pulumi:"applicationLogOutputUri"`
	// Locations of log files that your content generates during a stream session. Enter path values that are relative to the `ApplicationSourceUri` location. You can specify up to 10 log paths. Amazon GameLift Streams uploads designated log files to the Amazon S3 bucket that you specify in `ApplicationLogOutputUri` at the end of a stream session. To retrieve stored log files, call [GetStreamSession](https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamSession.html) and get the `LogFileLocationUri` .
	ApplicationLogPaths []string `pulumi:"applicationLogPaths"`
	// An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the application resource across all AWS Regions. For example:
	//
	// `arn:aws:gameliftstreams:us-west-2:123456789012:application/a-9ZY8X7Wv6` .
	Arn *string `pulumi:"arn"`
	// A human-readable label for the application. You can update this value later.
	Description *string `pulumi:"description"`
	// An ID that uniquely identifies the application resource. For example: `a-9ZY8X7Wv6` .
	Id *string `pulumi:"id"`
	// A list of labels to assign to the new application resource. Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
	Tags map[string]string `pulumi:"tags"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupApplicationResultOutput, error) {
			args := v.(LookupApplicationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:gameliftstreams:getApplication", args, LookupApplicationResultOutput{}, options).(LookupApplicationResultOutput), nil
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	// An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the application resource across all AWS Regions. For example:
	//
	// `arn:aws:gameliftstreams:us-west-2:123456789012:application/a-9ZY8X7Wv6` .
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

// An Amazon S3 URI to a bucket where you would like Amazon GameLift Streams to save application logs. Required if you specify one or more `ApplicationLogPaths` .
func (o LookupApplicationResultOutput) ApplicationLogOutputUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.ApplicationLogOutputUri }).(pulumi.StringPtrOutput)
}

// Locations of log files that your content generates during a stream session. Enter path values that are relative to the `ApplicationSourceUri` location. You can specify up to 10 log paths. Amazon GameLift Streams uploads designated log files to the Amazon S3 bucket that you specify in `ApplicationLogOutputUri` at the end of a stream session. To retrieve stored log files, call [GetStreamSession](https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamSession.html) and get the `LogFileLocationUri` .
func (o LookupApplicationResultOutput) ApplicationLogPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []string { return v.ApplicationLogPaths }).(pulumi.StringArrayOutput)
}

// An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the application resource across all AWS Regions. For example:
//
// `arn:aws:gameliftstreams:us-west-2:123456789012:application/a-9ZY8X7Wv6` .
func (o LookupApplicationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// A human-readable label for the application. You can update this value later.
func (o LookupApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// An ID that uniquely identifies the application resource. For example: `a-9ZY8X7Wv6` .
func (o LookupApplicationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// A list of labels to assign to the new application resource. Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
func (o LookupApplicationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
