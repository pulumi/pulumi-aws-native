// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies an Amazon Redshift subnet group.
func LookupClusterSubnetGroup(ctx *pulumi.Context, args *LookupClusterSubnetGroupArgs, opts ...pulumi.InvokeOption) (*LookupClusterSubnetGroupResult, error) {
	var rv LookupClusterSubnetGroupResult
	err := ctx.Invoke("aws-native:redshift:getClusterSubnetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterSubnetGroupArgs struct {
	// This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be "Default".
	ClusterSubnetGroupName string `pulumi:"clusterSubnetGroupName"`
}

type LookupClusterSubnetGroupResult struct {
	// This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be "Default".
	ClusterSubnetGroupName *string `pulumi:"clusterSubnetGroupName"`
	// The description of the parameter group.
	Description *string `pulumi:"description"`
	// The list of VPC subnet IDs
	SubnetIds []string `pulumi:"subnetIds"`
}

func LookupClusterSubnetGroupOutput(ctx *pulumi.Context, args LookupClusterSubnetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupClusterSubnetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterSubnetGroupResult, error) {
			args := v.(LookupClusterSubnetGroupArgs)
			r, err := LookupClusterSubnetGroup(ctx, &args, opts...)
			var s LookupClusterSubnetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterSubnetGroupResultOutput)
}

type LookupClusterSubnetGroupOutputArgs struct {
	// This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be "Default".
	ClusterSubnetGroupName pulumi.StringInput `pulumi:"clusterSubnetGroupName"`
}

func (LookupClusterSubnetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterSubnetGroupArgs)(nil)).Elem()
}

type LookupClusterSubnetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupClusterSubnetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterSubnetGroupResult)(nil)).Elem()
}

func (o LookupClusterSubnetGroupResultOutput) ToLookupClusterSubnetGroupResultOutput() LookupClusterSubnetGroupResultOutput {
	return o
}

func (o LookupClusterSubnetGroupResultOutput) ToLookupClusterSubnetGroupResultOutputWithContext(ctx context.Context) LookupClusterSubnetGroupResultOutput {
	return o
}

// This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be "Default".
func (o LookupClusterSubnetGroupResultOutput) ClusterSubnetGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterSubnetGroupResult) *string { return v.ClusterSubnetGroupName }).(pulumi.StringPtrOutput)
}

// The description of the parameter group.
func (o LookupClusterSubnetGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterSubnetGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The list of VPC subnet IDs
func (o LookupClusterSubnetGroupResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterSubnetGroupResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterSubnetGroupResultOutput{})
}