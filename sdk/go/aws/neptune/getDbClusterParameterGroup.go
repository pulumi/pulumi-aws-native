// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Neptune::DBClusterParameterGroup
func LookupDbClusterParameterGroup(ctx *pulumi.Context, args *LookupDbClusterParameterGroupArgs, opts ...pulumi.InvokeOption) (*LookupDbClusterParameterGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDbClusterParameterGroupResult
	err := ctx.Invoke("aws-native:neptune:getDbClusterParameterGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDbClusterParameterGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupDbClusterParameterGroupResult struct {
	Id         *string                      `pulumi:"id"`
	Parameters interface{}                  `pulumi:"parameters"`
	Tags       []DbClusterParameterGroupTag `pulumi:"tags"`
}

func LookupDbClusterParameterGroupOutput(ctx *pulumi.Context, args LookupDbClusterParameterGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDbClusterParameterGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDbClusterParameterGroupResult, error) {
			args := v.(LookupDbClusterParameterGroupArgs)
			r, err := LookupDbClusterParameterGroup(ctx, &args, opts...)
			var s LookupDbClusterParameterGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDbClusterParameterGroupResultOutput)
}

type LookupDbClusterParameterGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDbClusterParameterGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbClusterParameterGroupArgs)(nil)).Elem()
}

type LookupDbClusterParameterGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDbClusterParameterGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbClusterParameterGroupResult)(nil)).Elem()
}

func (o LookupDbClusterParameterGroupResultOutput) ToLookupDbClusterParameterGroupResultOutput() LookupDbClusterParameterGroupResultOutput {
	return o
}

func (o LookupDbClusterParameterGroupResultOutput) ToLookupDbClusterParameterGroupResultOutputWithContext(ctx context.Context) LookupDbClusterParameterGroupResultOutput {
	return o
}

func (o LookupDbClusterParameterGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDbClusterParameterGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupDbClusterParameterGroupResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDbClusterParameterGroupResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupDbClusterParameterGroupResultOutput) Tags() DbClusterParameterGroupTagArrayOutput {
	return o.ApplyT(func(v LookupDbClusterParameterGroupResult) []DbClusterParameterGroupTag { return v.Tags }).(DbClusterParameterGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDbClusterParameterGroupResultOutput{})
}