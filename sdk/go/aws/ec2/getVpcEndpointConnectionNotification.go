// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPCEndpointConnectionNotification
func LookupVpcEndpointConnectionNotification(ctx *pulumi.Context, args *LookupVpcEndpointConnectionNotificationArgs, opts ...pulumi.InvokeOption) (*LookupVpcEndpointConnectionNotificationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcEndpointConnectionNotificationResult
	err := ctx.Invoke("aws-native:ec2:getVpcEndpointConnectionNotification", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpcEndpointConnectionNotificationArgs struct {
	Id string `pulumi:"id"`
}

type LookupVpcEndpointConnectionNotificationResult struct {
	ConnectionEvents          []string `pulumi:"connectionEvents"`
	ConnectionNotificationArn *string  `pulumi:"connectionNotificationArn"`
	Id                        *string  `pulumi:"id"`
}

func LookupVpcEndpointConnectionNotificationOutput(ctx *pulumi.Context, args LookupVpcEndpointConnectionNotificationOutputArgs, opts ...pulumi.InvokeOption) LookupVpcEndpointConnectionNotificationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcEndpointConnectionNotificationResult, error) {
			args := v.(LookupVpcEndpointConnectionNotificationArgs)
			r, err := LookupVpcEndpointConnectionNotification(ctx, &args, opts...)
			var s LookupVpcEndpointConnectionNotificationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcEndpointConnectionNotificationResultOutput)
}

type LookupVpcEndpointConnectionNotificationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVpcEndpointConnectionNotificationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointConnectionNotificationArgs)(nil)).Elem()
}

type LookupVpcEndpointConnectionNotificationResultOutput struct{ *pulumi.OutputState }

func (LookupVpcEndpointConnectionNotificationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointConnectionNotificationResult)(nil)).Elem()
}

func (o LookupVpcEndpointConnectionNotificationResultOutput) ToLookupVpcEndpointConnectionNotificationResultOutput() LookupVpcEndpointConnectionNotificationResultOutput {
	return o
}

func (o LookupVpcEndpointConnectionNotificationResultOutput) ToLookupVpcEndpointConnectionNotificationResultOutputWithContext(ctx context.Context) LookupVpcEndpointConnectionNotificationResultOutput {
	return o
}

func (o LookupVpcEndpointConnectionNotificationResultOutput) ConnectionEvents() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointConnectionNotificationResult) []string { return v.ConnectionEvents }).(pulumi.StringArrayOutput)
}

func (o LookupVpcEndpointConnectionNotificationResultOutput) ConnectionNotificationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointConnectionNotificationResult) *string { return v.ConnectionNotificationArn }).(pulumi.StringPtrOutput)
}

func (o LookupVpcEndpointConnectionNotificationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointConnectionNotificationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcEndpointConnectionNotificationResultOutput{})
}