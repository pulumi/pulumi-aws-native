// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspacesweb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::WorkSpacesWeb::IpAccessSettings Resource Type
func LookupIpAccessSettings(ctx *pulumi.Context, args *LookupIpAccessSettingsArgs, opts ...pulumi.InvokeOption) (*LookupIpAccessSettingsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpAccessSettingsResult
	err := ctx.Invoke("aws-native:workspacesweb:getIpAccessSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpAccessSettingsArgs struct {
	IpAccessSettingsArn string `pulumi:"ipAccessSettingsArn"`
}

type LookupIpAccessSettingsResult struct {
	AssociatedPortalArns []string                 `pulumi:"associatedPortalArns"`
	CreationDate         *string                  `pulumi:"creationDate"`
	Description          *string                  `pulumi:"description"`
	DisplayName          *string                  `pulumi:"displayName"`
	IpAccessSettingsArn  *string                  `pulumi:"ipAccessSettingsArn"`
	IpRules              []IpAccessSettingsIpRule `pulumi:"ipRules"`
	Tags                 []IpAccessSettingsTag    `pulumi:"tags"`
}

func LookupIpAccessSettingsOutput(ctx *pulumi.Context, args LookupIpAccessSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupIpAccessSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpAccessSettingsResult, error) {
			args := v.(LookupIpAccessSettingsArgs)
			r, err := LookupIpAccessSettings(ctx, &args, opts...)
			var s LookupIpAccessSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpAccessSettingsResultOutput)
}

type LookupIpAccessSettingsOutputArgs struct {
	IpAccessSettingsArn pulumi.StringInput `pulumi:"ipAccessSettingsArn"`
}

func (LookupIpAccessSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpAccessSettingsArgs)(nil)).Elem()
}

type LookupIpAccessSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupIpAccessSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpAccessSettingsResult)(nil)).Elem()
}

func (o LookupIpAccessSettingsResultOutput) ToLookupIpAccessSettingsResultOutput() LookupIpAccessSettingsResultOutput {
	return o
}

func (o LookupIpAccessSettingsResultOutput) ToLookupIpAccessSettingsResultOutputWithContext(ctx context.Context) LookupIpAccessSettingsResultOutput {
	return o
}

func (o LookupIpAccessSettingsResultOutput) AssociatedPortalArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIpAccessSettingsResult) []string { return v.AssociatedPortalArns }).(pulumi.StringArrayOutput)
}

func (o LookupIpAccessSettingsResultOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAccessSettingsResult) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o LookupIpAccessSettingsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAccessSettingsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupIpAccessSettingsResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAccessSettingsResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupIpAccessSettingsResultOutput) IpAccessSettingsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAccessSettingsResult) *string { return v.IpAccessSettingsArn }).(pulumi.StringPtrOutput)
}

func (o LookupIpAccessSettingsResultOutput) IpRules() IpAccessSettingsIpRuleArrayOutput {
	return o.ApplyT(func(v LookupIpAccessSettingsResult) []IpAccessSettingsIpRule { return v.IpRules }).(IpAccessSettingsIpRuleArrayOutput)
}

func (o LookupIpAccessSettingsResultOutput) Tags() IpAccessSettingsTagArrayOutput {
	return o.ApplyT(func(v LookupIpAccessSettingsResult) []IpAccessSettingsTag { return v.Tags }).(IpAccessSettingsTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpAccessSettingsResultOutput{})
}