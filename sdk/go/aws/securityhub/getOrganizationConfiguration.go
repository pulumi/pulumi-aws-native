// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::SecurityHub::OrganizationConfiguration resource represents the configuration of your organization in Security Hub. Only the Security Hub administrator account can create Organization Configuration resource in each region and can opt-in to Central Configuration only in the aggregation region of FindingAggregator.
func LookupOrganizationConfiguration(ctx *pulumi.Context, args *LookupOrganizationConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrganizationConfigurationResult
	err := ctx.Invoke("aws-native:securityhub:getOrganizationConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationConfigurationArgs struct {
	// The identifier of the OrganizationConfiguration being created and assigned as the unique identifier.
	OrganizationConfigurationIdentifier string `pulumi:"organizationConfigurationIdentifier"`
}

type LookupOrganizationConfigurationResult struct {
	// Whether to automatically enable Security Hub in new member accounts when they join the organization.
	AutoEnable *bool `pulumi:"autoEnable"`
	// Whether to automatically enable Security Hub default standards in new member accounts when they join the organization.
	AutoEnableStandards *OrganizationConfigurationAutoEnableStandards `pulumi:"autoEnableStandards"`
	// Indicates whether the organization uses local or central configuration.
	ConfigurationType *OrganizationConfigurationConfigurationType `pulumi:"configurationType"`
	// Whether the maximum number of allowed member accounts are already associated with the Security Hub administrator account.
	MemberAccountLimitReached *bool `pulumi:"memberAccountLimitReached"`
	// The identifier of the OrganizationConfiguration being created and assigned as the unique identifier.
	OrganizationConfigurationIdentifier *string `pulumi:"organizationConfigurationIdentifier"`
	// Describes whether central configuration could be enabled as the ConfigurationType for the organization.
	Status *OrganizationConfigurationStatus `pulumi:"status"`
	// Provides an explanation if the value of Status is equal to FAILED when ConfigurationType is equal to CENTRAL.
	StatusMessage *string `pulumi:"statusMessage"`
}

func LookupOrganizationConfigurationOutput(ctx *pulumi.Context, args LookupOrganizationConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationConfigurationResult, error) {
			args := v.(LookupOrganizationConfigurationArgs)
			r, err := LookupOrganizationConfiguration(ctx, &args, opts...)
			var s LookupOrganizationConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationConfigurationResultOutput)
}

type LookupOrganizationConfigurationOutputArgs struct {
	// The identifier of the OrganizationConfiguration being created and assigned as the unique identifier.
	OrganizationConfigurationIdentifier pulumi.StringInput `pulumi:"organizationConfigurationIdentifier"`
}

func (LookupOrganizationConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationConfigurationArgs)(nil)).Elem()
}

type LookupOrganizationConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationConfigurationResult)(nil)).Elem()
}

func (o LookupOrganizationConfigurationResultOutput) ToLookupOrganizationConfigurationResultOutput() LookupOrganizationConfigurationResultOutput {
	return o
}

func (o LookupOrganizationConfigurationResultOutput) ToLookupOrganizationConfigurationResultOutputWithContext(ctx context.Context) LookupOrganizationConfigurationResultOutput {
	return o
}

// Whether to automatically enable Security Hub in new member accounts when they join the organization.
func (o LookupOrganizationConfigurationResultOutput) AutoEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigurationResult) *bool { return v.AutoEnable }).(pulumi.BoolPtrOutput)
}

// Whether to automatically enable Security Hub default standards in new member accounts when they join the organization.
func (o LookupOrganizationConfigurationResultOutput) AutoEnableStandards() OrganizationConfigurationAutoEnableStandardsPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigurationResult) *OrganizationConfigurationAutoEnableStandards {
		return v.AutoEnableStandards
	}).(OrganizationConfigurationAutoEnableStandardsPtrOutput)
}

// Indicates whether the organization uses local or central configuration.
func (o LookupOrganizationConfigurationResultOutput) ConfigurationType() OrganizationConfigurationConfigurationTypePtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigurationResult) *OrganizationConfigurationConfigurationType {
		return v.ConfigurationType
	}).(OrganizationConfigurationConfigurationTypePtrOutput)
}

// Whether the maximum number of allowed member accounts are already associated with the Security Hub administrator account.
func (o LookupOrganizationConfigurationResultOutput) MemberAccountLimitReached() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigurationResult) *bool { return v.MemberAccountLimitReached }).(pulumi.BoolPtrOutput)
}

// The identifier of the OrganizationConfiguration being created and assigned as the unique identifier.
func (o LookupOrganizationConfigurationResultOutput) OrganizationConfigurationIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigurationResult) *string { return v.OrganizationConfigurationIdentifier }).(pulumi.StringPtrOutput)
}

// Describes whether central configuration could be enabled as the ConfigurationType for the organization.
func (o LookupOrganizationConfigurationResultOutput) Status() OrganizationConfigurationStatusPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigurationResult) *OrganizationConfigurationStatus { return v.Status }).(OrganizationConfigurationStatusPtrOutput)
}

// Provides an explanation if the value of Status is equal to FAILED when ConfigurationType is equal to CENTRAL.
func (o LookupOrganizationConfigurationResultOutput) StatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConfigurationResult) *string { return v.StatusMessage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationConfigurationResultOutput{})
}