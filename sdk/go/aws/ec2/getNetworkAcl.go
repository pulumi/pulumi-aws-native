// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies a network ACL for your VPC.
func LookupNetworkAcl(ctx *pulumi.Context, args *LookupNetworkAclArgs, opts ...pulumi.InvokeOption) (*LookupNetworkAclResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkAclResult
	err := ctx.Invoke("aws-native:ec2:getNetworkAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkAclArgs struct {
	// The ID of the network ACL.
	Id string `pulumi:"id"`
}

type LookupNetworkAclResult struct {
	// The ID of the network ACL.
	Id *string `pulumi:"id"`
	// The tags for the network ACL.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupNetworkAclOutput(ctx *pulumi.Context, args LookupNetworkAclOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkAclResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkAclResultOutput, error) {
			args := v.(LookupNetworkAclArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupNetworkAclResult
			secret, err := ctx.InvokePackageRaw("aws-native:ec2:getNetworkAcl", args, &rv, "", opts...)
			if err != nil {
				return LookupNetworkAclResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupNetworkAclResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupNetworkAclResultOutput), nil
			}
			return output, nil
		}).(LookupNetworkAclResultOutput)
}

type LookupNetworkAclOutputArgs struct {
	// The ID of the network ACL.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupNetworkAclOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkAclArgs)(nil)).Elem()
}

type LookupNetworkAclResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkAclResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkAclResult)(nil)).Elem()
}

func (o LookupNetworkAclResultOutput) ToLookupNetworkAclResultOutput() LookupNetworkAclResultOutput {
	return o
}

func (o LookupNetworkAclResultOutput) ToLookupNetworkAclResultOutputWithContext(ctx context.Context) LookupNetworkAclResultOutput {
	return o
}

// The ID of the network ACL.
func (o LookupNetworkAclResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkAclResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The tags for the network ACL.
func (o LookupNetworkAclResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupNetworkAclResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkAclResultOutput{})
}
