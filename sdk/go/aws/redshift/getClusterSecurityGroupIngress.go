// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Redshift::ClusterSecurityGroupIngress
func LookupClusterSecurityGroupIngress(ctx *pulumi.Context, args *LookupClusterSecurityGroupIngressArgs, opts ...pulumi.InvokeOption) (*LookupClusterSecurityGroupIngressResult, error) {
	var rv LookupClusterSecurityGroupIngressResult
	err := ctx.Invoke("aws-native:redshift:getClusterSecurityGroupIngress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterSecurityGroupIngressArgs struct {
	Id string `pulumi:"id"`
}

type LookupClusterSecurityGroupIngressResult struct {
	Id *string `pulumi:"id"`
}

func LookupClusterSecurityGroupIngressOutput(ctx *pulumi.Context, args LookupClusterSecurityGroupIngressOutputArgs, opts ...pulumi.InvokeOption) LookupClusterSecurityGroupIngressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterSecurityGroupIngressResult, error) {
			args := v.(LookupClusterSecurityGroupIngressArgs)
			r, err := LookupClusterSecurityGroupIngress(ctx, &args, opts...)
			var s LookupClusterSecurityGroupIngressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterSecurityGroupIngressResultOutput)
}

type LookupClusterSecurityGroupIngressOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupClusterSecurityGroupIngressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterSecurityGroupIngressArgs)(nil)).Elem()
}

type LookupClusterSecurityGroupIngressResultOutput struct{ *pulumi.OutputState }

func (LookupClusterSecurityGroupIngressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterSecurityGroupIngressResult)(nil)).Elem()
}

func (o LookupClusterSecurityGroupIngressResultOutput) ToLookupClusterSecurityGroupIngressResultOutput() LookupClusterSecurityGroupIngressResultOutput {
	return o
}

func (o LookupClusterSecurityGroupIngressResultOutput) ToLookupClusterSecurityGroupIngressResultOutputWithContext(ctx context.Context) LookupClusterSecurityGroupIngressResultOutput {
	return o
}

func (o LookupClusterSecurityGroupIngressResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterSecurityGroupIngressResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterSecurityGroupIngressResultOutput{})
}