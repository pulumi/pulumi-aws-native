// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::TableOptimizer
func LookupTableOptimizer(ctx *pulumi.Context, args *LookupTableOptimizerArgs, opts ...pulumi.InvokeOption) (*LookupTableOptimizerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTableOptimizerResult
	err := ctx.Invoke("aws-native:glue:getTableOptimizer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableOptimizerArgs struct {
	Id string `pulumi:"id"`
}

type LookupTableOptimizerResult struct {
	Id                          *string                      `pulumi:"id"`
	TableOptimizerConfiguration *TableOptimizerConfiguration `pulumi:"tableOptimizerConfiguration"`
}

func LookupTableOptimizerOutput(ctx *pulumi.Context, args LookupTableOptimizerOutputArgs, opts ...pulumi.InvokeOption) LookupTableOptimizerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTableOptimizerResult, error) {
			args := v.(LookupTableOptimizerArgs)
			r, err := LookupTableOptimizer(ctx, &args, opts...)
			var s LookupTableOptimizerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTableOptimizerResultOutput)
}

type LookupTableOptimizerOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTableOptimizerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableOptimizerArgs)(nil)).Elem()
}

type LookupTableOptimizerResultOutput struct{ *pulumi.OutputState }

func (LookupTableOptimizerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableOptimizerResult)(nil)).Elem()
}

func (o LookupTableOptimizerResultOutput) ToLookupTableOptimizerResultOutput() LookupTableOptimizerResultOutput {
	return o
}

func (o LookupTableOptimizerResultOutput) ToLookupTableOptimizerResultOutputWithContext(ctx context.Context) LookupTableOptimizerResultOutput {
	return o
}

func (o LookupTableOptimizerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTableOptimizerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupTableOptimizerResultOutput) TableOptimizerConfiguration() TableOptimizerConfigurationPtrOutput {
	return o.ApplyT(func(v LookupTableOptimizerResult) *TableOptimizerConfiguration { return v.TableOptimizerConfiguration }).(TableOptimizerConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableOptimizerResultOutput{})
}