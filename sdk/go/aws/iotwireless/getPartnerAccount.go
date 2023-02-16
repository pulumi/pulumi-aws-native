// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotwireless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create and manage partner account
func LookupPartnerAccount(ctx *pulumi.Context, args *LookupPartnerAccountArgs, opts ...pulumi.InvokeOption) (*LookupPartnerAccountResult, error) {
	var rv LookupPartnerAccountResult
	err := ctx.Invoke("aws-native:iotwireless:getPartnerAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerAccountArgs struct {
	// The partner account ID to disassociate from the AWS account
	PartnerAccountId string `pulumi:"partnerAccountId"`
}

type LookupPartnerAccountResult struct {
	// Whether the partner account is linked to the AWS account.
	AccountLinked *bool `pulumi:"accountLinked"`
	// PartnerAccount arn. Returned after successful create.
	Arn *string `pulumi:"arn"`
	// The fingerprint of the Sidewalk application server private key.
	Fingerprint *string `pulumi:"fingerprint"`
	// The partner type
	PartnerType *PartnerAccountPartnerType `pulumi:"partnerType"`
	// The Sidewalk account credentials.
	SidewalkResponse *PartnerAccountSidewalkAccountInfoWithFingerprint `pulumi:"sidewalkResponse"`
	// A list of key-value pairs that contain metadata for the destination.
	Tags []PartnerAccountTag `pulumi:"tags"`
}

func LookupPartnerAccountOutput(ctx *pulumi.Context, args LookupPartnerAccountOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerAccountResult, error) {
			args := v.(LookupPartnerAccountArgs)
			r, err := LookupPartnerAccount(ctx, &args, opts...)
			var s LookupPartnerAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerAccountResultOutput)
}

type LookupPartnerAccountOutputArgs struct {
	// The partner account ID to disassociate from the AWS account
	PartnerAccountId pulumi.StringInput `pulumi:"partnerAccountId"`
}

func (LookupPartnerAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerAccountArgs)(nil)).Elem()
}

type LookupPartnerAccountResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerAccountResult)(nil)).Elem()
}

func (o LookupPartnerAccountResultOutput) ToLookupPartnerAccountResultOutput() LookupPartnerAccountResultOutput {
	return o
}

func (o LookupPartnerAccountResultOutput) ToLookupPartnerAccountResultOutputWithContext(ctx context.Context) LookupPartnerAccountResultOutput {
	return o
}

// Whether the partner account is linked to the AWS account.
func (o LookupPartnerAccountResultOutput) AccountLinked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPartnerAccountResult) *bool { return v.AccountLinked }).(pulumi.BoolPtrOutput)
}

// PartnerAccount arn. Returned after successful create.
func (o LookupPartnerAccountResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerAccountResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The fingerprint of the Sidewalk application server private key.
func (o LookupPartnerAccountResultOutput) Fingerprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerAccountResult) *string { return v.Fingerprint }).(pulumi.StringPtrOutput)
}

// The partner type
func (o LookupPartnerAccountResultOutput) PartnerType() PartnerAccountPartnerTypePtrOutput {
	return o.ApplyT(func(v LookupPartnerAccountResult) *PartnerAccountPartnerType { return v.PartnerType }).(PartnerAccountPartnerTypePtrOutput)
}

// The Sidewalk account credentials.
func (o LookupPartnerAccountResultOutput) SidewalkResponse() PartnerAccountSidewalkAccountInfoWithFingerprintPtrOutput {
	return o.ApplyT(func(v LookupPartnerAccountResult) *PartnerAccountSidewalkAccountInfoWithFingerprint {
		return v.SidewalkResponse
	}).(PartnerAccountSidewalkAccountInfoWithFingerprintPtrOutput)
}

// A list of key-value pairs that contain metadata for the destination.
func (o LookupPartnerAccountResultOutput) Tags() PartnerAccountTagArrayOutput {
	return o.ApplyT(func(v LookupPartnerAccountResult) []PartnerAccountTag { return v.Tags }).(PartnerAccountTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerAccountResultOutput{})
}