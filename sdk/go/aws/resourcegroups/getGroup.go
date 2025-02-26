// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcegroups

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for ResourceGroups::Group
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupResult
	err := ctx.Invoke("aws-native:resourcegroups:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	// The name of the resource group
	Name string `pulumi:"name"`
}

type LookupGroupResult struct {
	// The Resource Group ARN.
	Arn *string `pulumi:"arn"`
	// The service configuration currently associated with the resource group and in effect for the members of the resource group. A `Configuration` consists of one or more `ConfigurationItem` entries. For information about service configurations for resource groups and how to construct them, see [Service configurations for resource groups](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html) in the *AWS Resource Groups User Guide* .
	//
	// > You can include either a `Configuration` or a `ResourceQuery` , but not both.
	Configuration []GroupConfigurationItem `pulumi:"configuration"`
	// The description of the resource group
	Description *string `pulumi:"description"`
	// The resource query structure that is used to dynamically determine which AWS resources are members of the associated resource group. For more information about queries and how to construct them, see [Build queries and groups in AWS Resource Groups](https://docs.aws.amazon.com//ARG/latest/userguide/gettingstarted-query.html) in the *AWS Resource Groups User Guide*
	//
	// > - You can include either a `ResourceQuery` or a `Configuration` , but not both.
	// > - You can specify the group's membership either by using a `ResourceQuery` or by using a list of `Resources` , but not both.
	ResourceQuery *GroupResourceQuery `pulumi:"resourceQuery"`
	// A list of the Amazon Resource Names (ARNs) of AWS resources that you want to add to the specified group.
	//
	// > - You can specify the group membership either by using a list of `Resources` or by using a `ResourceQuery` , but not both.
	// > - You can include a `Resources` property only if you also specify a `Configuration` property.
	Resources []string `pulumi:"resources"`
	// The tag key and value pairs that are attached to the resource group.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGroupResultOutput, error) {
			args := v.(LookupGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:resourcegroups:getGroup", args, LookupGroupResultOutput{}, options).(LookupGroupResultOutput), nil
		}).(LookupGroupResultOutput)
}

type LookupGroupOutputArgs struct {
	// The name of the resource group
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// The Resource Group ARN.
func (o LookupGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The service configuration currently associated with the resource group and in effect for the members of the resource group. A `Configuration` consists of one or more `ConfigurationItem` entries. For information about service configurations for resource groups and how to construct them, see [Service configurations for resource groups](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html) in the *AWS Resource Groups User Guide* .
//
// > You can include either a `Configuration` or a `ResourceQuery` , but not both.
func (o LookupGroupResultOutput) Configuration() GroupConfigurationItemArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []GroupConfigurationItem { return v.Configuration }).(GroupConfigurationItemArrayOutput)
}

// The description of the resource group
func (o LookupGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The resource query structure that is used to dynamically determine which AWS resources are members of the associated resource group. For more information about queries and how to construct them, see [Build queries and groups in AWS Resource Groups](https://docs.aws.amazon.com//ARG/latest/userguide/gettingstarted-query.html) in the *AWS Resource Groups User Guide*
//
// > - You can include either a `ResourceQuery` or a `Configuration` , but not both.
// > - You can specify the group's membership either by using a `ResourceQuery` or by using a list of `Resources` , but not both.
func (o LookupGroupResultOutput) ResourceQuery() GroupResourceQueryPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *GroupResourceQuery { return v.ResourceQuery }).(GroupResourceQueryPtrOutput)
}

// A list of the Amazon Resource Names (ARNs) of AWS resources that you want to add to the specified group.
//
// > - You can specify the group membership either by using a list of `Resources` or by using a `ResourceQuery` , but not both.
// > - You can include a `Resources` property only if you also specify a `Configuration` property.
func (o LookupGroupResultOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

// The tag key and value pairs that are attached to the resource group.
func (o LookupGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
