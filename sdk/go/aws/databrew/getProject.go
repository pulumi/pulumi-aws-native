// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databrew

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataBrew::Project.
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("aws-native:databrew:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	// Project name
	Name string `pulumi:"name"`
}

type LookupProjectResult struct {
	// Dataset name
	DatasetName *string `pulumi:"datasetName"`
	// Recipe name
	RecipeName *string `pulumi:"recipeName"`
	// Role arn
	RoleArn *string `pulumi:"roleArn"`
	// Sample
	Sample *ProjectSample `pulumi:"sample"`
	// Metadata tags that have been applied to the project.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectResultOutput, error) {
			args := v.(LookupProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:databrew:getProject", args, LookupProjectResultOutput{}, options).(LookupProjectResultOutput), nil
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	// Project name
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// Dataset name
func (o LookupProjectResultOutput) DatasetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.DatasetName }).(pulumi.StringPtrOutput)
}

// Recipe name
func (o LookupProjectResultOutput) RecipeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.RecipeName }).(pulumi.StringPtrOutput)
}

// Role arn
func (o LookupProjectResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// Sample
func (o LookupProjectResultOutput) Sample() ProjectSamplePtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectSample { return v.Sample }).(ProjectSamplePtrOutput)
}

// Metadata tags that have been applied to the project.
func (o LookupProjectResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
