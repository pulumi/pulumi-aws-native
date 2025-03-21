// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3outposts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type Definition for AWS::S3Outposts::BucketPolicy
func LookupBucketPolicy(ctx *pulumi.Context, args *LookupBucketPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBucketPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBucketPolicyResult
	err := ctx.Invoke("aws-native:s3outposts:getBucketPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBucketPolicyArgs struct {
	// The Amazon Resource Name (ARN) of the specified bucket.
	Bucket string `pulumi:"bucket"`
}

type LookupBucketPolicyResult struct {
	// A policy document containing permissions to add to the specified bucket.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3Outposts::BucketPolicy` for more information about the expected schema for this property.
	PolicyDocument interface{} `pulumi:"policyDocument"`
}

func LookupBucketPolicyOutput(ctx *pulumi.Context, args LookupBucketPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBucketPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBucketPolicyResultOutput, error) {
			args := v.(LookupBucketPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:s3outposts:getBucketPolicy", args, LookupBucketPolicyResultOutput{}, options).(LookupBucketPolicyResultOutput), nil
		}).(LookupBucketPolicyResultOutput)
}

type LookupBucketPolicyOutputArgs struct {
	// The Amazon Resource Name (ARN) of the specified bucket.
	Bucket pulumi.StringInput `pulumi:"bucket"`
}

func (LookupBucketPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketPolicyArgs)(nil)).Elem()
}

type LookupBucketPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBucketPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketPolicyResult)(nil)).Elem()
}

func (o LookupBucketPolicyResultOutput) ToLookupBucketPolicyResultOutput() LookupBucketPolicyResultOutput {
	return o
}

func (o LookupBucketPolicyResultOutput) ToLookupBucketPolicyResultOutputWithContext(ctx context.Context) LookupBucketPolicyResultOutput {
	return o
}

// A policy document containing permissions to add to the specified bucket.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3Outposts::BucketPolicy` for more information about the expected schema for this property.
func (o LookupBucketPolicyResultOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupBucketPolicyResult) interface{} { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketPolicyResultOutput{})
}
