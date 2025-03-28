// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// resource definition
func LookupSoftwarePackage(ctx *pulumi.Context, args *LookupSoftwarePackageArgs, opts ...pulumi.InvokeOption) (*LookupSoftwarePackageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSoftwarePackageResult
	err := ctx.Invoke("aws-native:iot:getSoftwarePackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSoftwarePackageArgs struct {
	// The name of the new software package.
	PackageName string `pulumi:"packageName"`
}

type LookupSoftwarePackageResult struct {
	// A summary of the package being created. This can be used to outline the package's contents or purpose.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) for the package.
	PackageArn *string `pulumi:"packageArn"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupSoftwarePackageOutput(ctx *pulumi.Context, args LookupSoftwarePackageOutputArgs, opts ...pulumi.InvokeOption) LookupSoftwarePackageResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSoftwarePackageResultOutput, error) {
			args := v.(LookupSoftwarePackageArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:iot:getSoftwarePackage", args, LookupSoftwarePackageResultOutput{}, options).(LookupSoftwarePackageResultOutput), nil
		}).(LookupSoftwarePackageResultOutput)
}

type LookupSoftwarePackageOutputArgs struct {
	// The name of the new software package.
	PackageName pulumi.StringInput `pulumi:"packageName"`
}

func (LookupSoftwarePackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSoftwarePackageArgs)(nil)).Elem()
}

type LookupSoftwarePackageResultOutput struct{ *pulumi.OutputState }

func (LookupSoftwarePackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSoftwarePackageResult)(nil)).Elem()
}

func (o LookupSoftwarePackageResultOutput) ToLookupSoftwarePackageResultOutput() LookupSoftwarePackageResultOutput {
	return o
}

func (o LookupSoftwarePackageResultOutput) ToLookupSoftwarePackageResultOutputWithContext(ctx context.Context) LookupSoftwarePackageResultOutput {
	return o
}

// A summary of the package being created. This can be used to outline the package's contents or purpose.
func (o LookupSoftwarePackageResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSoftwarePackageResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) for the package.
func (o LookupSoftwarePackageResultOutput) PackageArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSoftwarePackageResult) *string { return v.PackageArn }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupSoftwarePackageResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupSoftwarePackageResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSoftwarePackageResultOutput{})
}
