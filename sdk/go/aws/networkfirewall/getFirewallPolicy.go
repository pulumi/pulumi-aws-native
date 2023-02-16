// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::NetworkFirewall::FirewallPolicy
func LookupFirewallPolicy(ctx *pulumi.Context, args *LookupFirewallPolicyArgs, opts ...pulumi.InvokeOption) (*LookupFirewallPolicyResult, error) {
	var rv LookupFirewallPolicyResult
	err := ctx.Invoke("aws-native:networkfirewall:getFirewallPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallPolicyArgs struct {
	FirewallPolicyArn string `pulumi:"firewallPolicyArn"`
}

type LookupFirewallPolicyResult struct {
	Description       *string             `pulumi:"description"`
	FirewallPolicy    *FirewallPolicyType `pulumi:"firewallPolicy"`
	FirewallPolicyArn *string             `pulumi:"firewallPolicyArn"`
	FirewallPolicyId  *string             `pulumi:"firewallPolicyId"`
	Tags              []FirewallPolicyTag `pulumi:"tags"`
}

func LookupFirewallPolicyOutput(ctx *pulumi.Context, args LookupFirewallPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallPolicyResult, error) {
			args := v.(LookupFirewallPolicyArgs)
			r, err := LookupFirewallPolicy(ctx, &args, opts...)
			var s LookupFirewallPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallPolicyResultOutput)
}

type LookupFirewallPolicyOutputArgs struct {
	FirewallPolicyArn pulumi.StringInput `pulumi:"firewallPolicyArn"`
}

func (LookupFirewallPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyArgs)(nil)).Elem()
}

type LookupFirewallPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyResult)(nil)).Elem()
}

func (o LookupFirewallPolicyResultOutput) ToLookupFirewallPolicyResultOutput() LookupFirewallPolicyResultOutput {
	return o
}

func (o LookupFirewallPolicyResultOutput) ToLookupFirewallPolicyResultOutputWithContext(ctx context.Context) LookupFirewallPolicyResultOutput {
	return o
}

func (o LookupFirewallPolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicyResultOutput) FirewallPolicy() FirewallPolicyTypePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *FirewallPolicyType { return v.FirewallPolicy }).(FirewallPolicyTypePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) FirewallPolicyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *string { return v.FirewallPolicyArn }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicyResultOutput) FirewallPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *string { return v.FirewallPolicyId }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Tags() FirewallPolicyTagArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) []FirewallPolicyTag { return v.Tags }).(FirewallPolicyTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallPolicyResultOutput{})
}