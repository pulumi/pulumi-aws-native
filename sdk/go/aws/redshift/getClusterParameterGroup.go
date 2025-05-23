// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Redshift::ClusterParameterGroup
func LookupClusterParameterGroup(ctx *pulumi.Context, args *LookupClusterParameterGroupArgs, opts ...pulumi.InvokeOption) (*LookupClusterParameterGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterParameterGroupResult
	err := ctx.Invoke("aws-native:redshift:getClusterParameterGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterParameterGroupArgs struct {
	// The name of the cluster parameter group.
	ParameterGroupName string `pulumi:"parameterGroupName"`
}

type LookupClusterParameterGroupResult struct {
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupClusterParameterGroupOutput(ctx *pulumi.Context, args LookupClusterParameterGroupOutputArgs, opts ...pulumi.InvokeOption) LookupClusterParameterGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupClusterParameterGroupResultOutput, error) {
			args := v.(LookupClusterParameterGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:redshift:getClusterParameterGroup", args, LookupClusterParameterGroupResultOutput{}, options).(LookupClusterParameterGroupResultOutput), nil
		}).(LookupClusterParameterGroupResultOutput)
}

type LookupClusterParameterGroupOutputArgs struct {
	// The name of the cluster parameter group.
	ParameterGroupName pulumi.StringInput `pulumi:"parameterGroupName"`
}

func (LookupClusterParameterGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterParameterGroupArgs)(nil)).Elem()
}

type LookupClusterParameterGroupResultOutput struct{ *pulumi.OutputState }

func (LookupClusterParameterGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterParameterGroupResult)(nil)).Elem()
}

func (o LookupClusterParameterGroupResultOutput) ToLookupClusterParameterGroupResultOutput() LookupClusterParameterGroupResultOutput {
	return o
}

func (o LookupClusterParameterGroupResultOutput) ToLookupClusterParameterGroupResultOutputWithContext(ctx context.Context) LookupClusterParameterGroupResultOutput {
	return o
}

// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
func (o LookupClusterParameterGroupResultOutput) Parameters() ClusterParameterGroupParameterArrayOutput {
	return o.ApplyT(func(v LookupClusterParameterGroupResult) []ClusterParameterGroupParameter { return v.Parameters }).(ClusterParameterGroupParameterArrayOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupClusterParameterGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupClusterParameterGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterParameterGroupResultOutput{})
}
