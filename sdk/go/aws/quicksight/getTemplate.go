// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::Template Resource Type.
func LookupTemplate(ctx *pulumi.Context, args *LookupTemplateArgs, opts ...pulumi.InvokeOption) (*LookupTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTemplateResult
	err := ctx.Invoke("aws-native:quicksight:getTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateArgs struct {
	// The ID for the AWS account that the group is in. You use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId string `pulumi:"awsAccountId"`
	// An ID for the template that you want to create. This template is unique per AWS Region ; in each AWS account.
	TemplateId string `pulumi:"templateId"`
}

type LookupTemplateResult struct {
	// <p>The Amazon Resource Name (ARN) of the template.</p>
	Arn *string `pulumi:"arn"`
	// <p>Time when this was created.</p>
	CreatedTime *string `pulumi:"createdTime"`
	// <p>Time when this was last updated.</p>
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// A display name for the template.
	Name *string `pulumi:"name"`
	// A list of resource permissions to be set on the template.
	Permissions []TemplateResourcePermission `pulumi:"permissions"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.
	Tags    []aws.Tag        `pulumi:"tags"`
	Version *TemplateVersion `pulumi:"version"`
}

func LookupTemplateOutput(ctx *pulumi.Context, args LookupTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupTemplateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTemplateResultOutput, error) {
			args := v.(LookupTemplateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:quicksight:getTemplate", args, LookupTemplateResultOutput{}, options).(LookupTemplateResultOutput), nil
		}).(LookupTemplateResultOutput)
}

type LookupTemplateOutputArgs struct {
	// The ID for the AWS account that the group is in. You use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId pulumi.StringInput `pulumi:"awsAccountId"`
	// An ID for the template that you want to create. This template is unique per AWS Region ; in each AWS account.
	TemplateId pulumi.StringInput `pulumi:"templateId"`
}

func (LookupTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateArgs)(nil)).Elem()
}

type LookupTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateResult)(nil)).Elem()
}

func (o LookupTemplateResultOutput) ToLookupTemplateResultOutput() LookupTemplateResultOutput {
	return o
}

func (o LookupTemplateResultOutput) ToLookupTemplateResultOutputWithContext(ctx context.Context) LookupTemplateResultOutput {
	return o
}

// <p>The Amazon Resource Name (ARN) of the template.</p>
func (o LookupTemplateResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// <p>Time when this was created.</p>
func (o LookupTemplateResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// <p>Time when this was last updated.</p>
func (o LookupTemplateResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// A display name for the template.
func (o LookupTemplateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of resource permissions to be set on the template.
func (o LookupTemplateResultOutput) Permissions() TemplateResourcePermissionArrayOutput {
	return o.ApplyT(func(v LookupTemplateResult) []TemplateResourcePermission { return v.Permissions }).(TemplateResourcePermissionArrayOutput)
}

// Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.
func (o LookupTemplateResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupTemplateResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func (o LookupTemplateResultOutput) Version() TemplateVersionPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *TemplateVersion { return v.Version }).(TemplateVersionPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateResultOutput{})
}
