// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package shield

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Authorizes the Shield Response Team (SRT) to use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
func LookupProactiveEngagement(ctx *pulumi.Context, args *LookupProactiveEngagementArgs, opts ...pulumi.InvokeOption) (*LookupProactiveEngagementResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProactiveEngagementResult
	err := ctx.Invoke("aws-native:shield:getProactiveEngagement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProactiveEngagementArgs struct {
	AccountId string `pulumi:"accountId"`
}

type LookupProactiveEngagementResult struct {
	AccountId *string `pulumi:"accountId"`
	// A list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you for escalations to the SRT and to initiate proactive customer support.
	// To enable proactive engagement, the contact list must include at least one phone number.
	EmergencyContactList []ProactiveEngagementEmergencyContact `pulumi:"emergencyContactList"`
	// If `ENABLED`, the Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
	// If `DISABLED`, the SRT will not proactively notify contacts about escalations or to initiate proactive customer support.
	ProactiveEngagementStatus *ProactiveEngagementStatus `pulumi:"proactiveEngagementStatus"`
}

func LookupProactiveEngagementOutput(ctx *pulumi.Context, args LookupProactiveEngagementOutputArgs, opts ...pulumi.InvokeOption) LookupProactiveEngagementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProactiveEngagementResult, error) {
			args := v.(LookupProactiveEngagementArgs)
			r, err := LookupProactiveEngagement(ctx, &args, opts...)
			var s LookupProactiveEngagementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProactiveEngagementResultOutput)
}

type LookupProactiveEngagementOutputArgs struct {
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (LookupProactiveEngagementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProactiveEngagementArgs)(nil)).Elem()
}

type LookupProactiveEngagementResultOutput struct{ *pulumi.OutputState }

func (LookupProactiveEngagementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProactiveEngagementResult)(nil)).Elem()
}

func (o LookupProactiveEngagementResultOutput) ToLookupProactiveEngagementResultOutput() LookupProactiveEngagementResultOutput {
	return o
}

func (o LookupProactiveEngagementResultOutput) ToLookupProactiveEngagementResultOutputWithContext(ctx context.Context) LookupProactiveEngagementResultOutput {
	return o
}

func (o LookupProactiveEngagementResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProactiveEngagementResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// A list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you for escalations to the SRT and to initiate proactive customer support.
// To enable proactive engagement, the contact list must include at least one phone number.
func (o LookupProactiveEngagementResultOutput) EmergencyContactList() ProactiveEngagementEmergencyContactArrayOutput {
	return o.ApplyT(func(v LookupProactiveEngagementResult) []ProactiveEngagementEmergencyContact {
		return v.EmergencyContactList
	}).(ProactiveEngagementEmergencyContactArrayOutput)
}

// If `ENABLED`, the Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
// If `DISABLED`, the SRT will not proactively notify contacts about escalations or to initiate proactive customer support.
func (o LookupProactiveEngagementResultOutput) ProactiveEngagementStatus() ProactiveEngagementStatusPtrOutput {
	return o.ApplyT(func(v LookupProactiveEngagementResult) *ProactiveEngagementStatus { return v.ProactiveEngagementStatus }).(ProactiveEngagementStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProactiveEngagementResultOutput{})
}