// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotfleetwise

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::IoTFleetWise::Campaign Resource Type
func LookupCampaign(ctx *pulumi.Context, args *LookupCampaignArgs, opts ...pulumi.InvokeOption) (*LookupCampaignResult, error) {
	var rv LookupCampaignResult
	err := ctx.Invoke("aws-native:iotfleetwise:getCampaign", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCampaignArgs struct {
	Name string `pulumi:"name"`
}

type LookupCampaignResult struct {
	Arn                  *string                     `pulumi:"arn"`
	CreationTime         *string                     `pulumi:"creationTime"`
	DataExtraDimensions  []string                    `pulumi:"dataExtraDimensions"`
	Description          *string                     `pulumi:"description"`
	LastModificationTime *string                     `pulumi:"lastModificationTime"`
	SignalsToCollect     []CampaignSignalInformation `pulumi:"signalsToCollect"`
	Status               *CampaignStatus             `pulumi:"status"`
	Tags                 []CampaignTag               `pulumi:"tags"`
}

func LookupCampaignOutput(ctx *pulumi.Context, args LookupCampaignOutputArgs, opts ...pulumi.InvokeOption) LookupCampaignResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCampaignResult, error) {
			args := v.(LookupCampaignArgs)
			r, err := LookupCampaign(ctx, &args, opts...)
			var s LookupCampaignResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCampaignResultOutput)
}

type LookupCampaignOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupCampaignOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCampaignArgs)(nil)).Elem()
}

type LookupCampaignResultOutput struct{ *pulumi.OutputState }

func (LookupCampaignResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCampaignResult)(nil)).Elem()
}

func (o LookupCampaignResultOutput) ToLookupCampaignResultOutput() LookupCampaignResultOutput {
	return o
}

func (o LookupCampaignResultOutput) ToLookupCampaignResultOutputWithContext(ctx context.Context) LookupCampaignResultOutput {
	return o
}

func (o LookupCampaignResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupCampaignResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupCampaignResultOutput) DataExtraDimensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCampaignResult) []string { return v.DataExtraDimensions }).(pulumi.StringArrayOutput)
}

func (o LookupCampaignResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupCampaignResultOutput) LastModificationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.LastModificationTime }).(pulumi.StringPtrOutput)
}

func (o LookupCampaignResultOutput) SignalsToCollect() CampaignSignalInformationArrayOutput {
	return o.ApplyT(func(v LookupCampaignResult) []CampaignSignalInformation { return v.SignalsToCollect }).(CampaignSignalInformationArrayOutput)
}

func (o LookupCampaignResultOutput) Status() CampaignStatusPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *CampaignStatus { return v.Status }).(CampaignStatusPtrOutput)
}

func (o LookupCampaignResultOutput) Tags() CampaignTagArrayOutput {
	return o.ApplyT(func(v LookupCampaignResult) []CampaignTag { return v.Tags }).(CampaignTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCampaignResultOutput{})
}