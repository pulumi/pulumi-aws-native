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

// Definition of the AWS::QuickSight::Folder Resource Type.
func LookupFolder(ctx *pulumi.Context, args *LookupFolderArgs, opts ...pulumi.InvokeOption) (*LookupFolderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFolderResult
	err := ctx.Invoke("aws-native:quicksight:getFolder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFolderArgs struct {
	// The ID for the AWS account where you want to create the folder.
	AwsAccountId string `pulumi:"awsAccountId"`
	// The ID of the folder.
	FolderId string `pulumi:"folderId"`
}

type LookupFolderResult struct {
	// <p>The Amazon Resource Name (ARN) for the folder.</p>
	Arn *string `pulumi:"arn"`
	// <p>The time that the folder was created.</p>
	CreatedTime *string `pulumi:"createdTime"`
	// <p>The time that the folder was last updated.</p>
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// A display name for the folder.
	Name *string `pulumi:"name"`
	// A structure that describes the principals and the resource-level permissions of a folder.
	//
	// To specify no permissions, omit `Permissions` .
	Permissions []FolderResourcePermission `pulumi:"permissions"`
	// A list of tags for the folders that you want to apply overrides to.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupFolderOutput(ctx *pulumi.Context, args LookupFolderOutputArgs, opts ...pulumi.InvokeOption) LookupFolderResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFolderResultOutput, error) {
			args := v.(LookupFolderArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:quicksight:getFolder", args, LookupFolderResultOutput{}, options).(LookupFolderResultOutput), nil
		}).(LookupFolderResultOutput)
}

type LookupFolderOutputArgs struct {
	// The ID for the AWS account where you want to create the folder.
	AwsAccountId pulumi.StringInput `pulumi:"awsAccountId"`
	// The ID of the folder.
	FolderId pulumi.StringInput `pulumi:"folderId"`
}

func (LookupFolderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderArgs)(nil)).Elem()
}

type LookupFolderResultOutput struct{ *pulumi.OutputState }

func (LookupFolderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderResult)(nil)).Elem()
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutput() LookupFolderResultOutput {
	return o
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutputWithContext(ctx context.Context) LookupFolderResultOutput {
	return o
}

// <p>The Amazon Resource Name (ARN) for the folder.</p>
func (o LookupFolderResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFolderResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// <p>The time that the folder was created.</p>
func (o LookupFolderResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFolderResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// <p>The time that the folder was last updated.</p>
func (o LookupFolderResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFolderResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// A display name for the folder.
func (o LookupFolderResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFolderResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A structure that describes the principals and the resource-level permissions of a folder.
//
// To specify no permissions, omit `Permissions` .
func (o LookupFolderResultOutput) Permissions() FolderResourcePermissionArrayOutput {
	return o.ApplyT(func(v LookupFolderResult) []FolderResourcePermission { return v.Permissions }).(FolderResourcePermissionArrayOutput)
}

// A list of tags for the folders that you want to apply overrides to.
func (o LookupFolderResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupFolderResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderResultOutput{})
}
