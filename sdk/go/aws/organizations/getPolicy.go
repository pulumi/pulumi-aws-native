// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Policies in AWS Organizations enable you to manage different features of the AWS accounts in your organization.  You can use policies when all features are enabled in your organization.
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyResult
	err := ctx.Invoke("aws-native:organizations:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyArgs struct {
	// Id of the Policy
	Id string `pulumi:"id"`
}

type LookupPolicyResult struct {
	// ARN of the Policy
	Arn *string `pulumi:"arn"`
	// A boolean value that indicates whether the specified policy is an AWS managed policy. If true, then you can attach the policy to roots, OUs, or accounts, but you cannot edit it.
	AwsManaged *bool `pulumi:"awsManaged"`
	// The Policy text content. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Organizations::Policy` for more information about the expected schema for this property.
	Content interface{} `pulumi:"content"`
	// Human readable description of the policy
	Description *string `pulumi:"description"`
	// Id of the Policy
	Id *string `pulumi:"id"`
	// Name of the Policy
	Name *string `pulumi:"name"`
	// A list of tags that you want to attach to the newly created policy. For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null.
	Tags []aws.Tag `pulumi:"tags"`
	// List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to
	TargetIds []string `pulumi:"targetIds"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPolicyResultOutput, error) {
			args := v.(LookupPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:organizations:getPolicy", args, LookupPolicyResultOutput{}, options).(LookupPolicyResultOutput), nil
		}).(LookupPolicyResultOutput)
}

type LookupPolicyOutputArgs struct {
	// Id of the Policy
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

// ARN of the Policy
func (o LookupPolicyResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// A boolean value that indicates whether the specified policy is an AWS managed policy. If true, then you can attach the policy to roots, OUs, or accounts, but you cannot edit it.
func (o LookupPolicyResultOutput) AwsManaged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *bool { return v.AwsManaged }).(pulumi.BoolPtrOutput)
}

// The Policy text content. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Organizations::Policy` for more information about the expected schema for this property.
func (o LookupPolicyResultOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyResult) interface{} { return v.Content }).(pulumi.AnyOutput)
}

// Human readable description of the policy
func (o LookupPolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Id of the Policy
func (o LookupPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of the Policy
func (o LookupPolicyResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of tags that you want to attach to the newly created policy. For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null.
func (o LookupPolicyResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to
func (o LookupPolicyResultOutput) TargetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.TargetIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
