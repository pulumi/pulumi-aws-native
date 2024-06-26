// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::UserPoolRiskConfigurationAttachment
func LookupUserPoolRiskConfigurationAttachment(ctx *pulumi.Context, args *LookupUserPoolRiskConfigurationAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolRiskConfigurationAttachmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserPoolRiskConfigurationAttachmentResult
	err := ctx.Invoke("aws-native:cognito:getUserPoolRiskConfigurationAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolRiskConfigurationAttachmentArgs struct {
	// The app client ID. You can specify the risk configuration for a single client (with a specific ClientId) or for all clients (by setting the ClientId to `ALL` ).
	ClientId string `pulumi:"clientId"`
	// The user pool ID.
	UserPoolId string `pulumi:"userPoolId"`
}

type LookupUserPoolRiskConfigurationAttachmentResult struct {
	// The account takeover risk configuration object, including the `NotifyConfiguration` object and `Actions` to take if there is an account takeover.
	AccountTakeoverRiskConfiguration *UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationType `pulumi:"accountTakeoverRiskConfiguration"`
	// The compromised credentials risk configuration object, including the `EventFilter` and the `EventAction` .
	CompromisedCredentialsRiskConfiguration *UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationType `pulumi:"compromisedCredentialsRiskConfiguration"`
	// The configuration to override the risk decision.
	RiskExceptionConfiguration *UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationType `pulumi:"riskExceptionConfiguration"`
}

func LookupUserPoolRiskConfigurationAttachmentOutput(ctx *pulumi.Context, args LookupUserPoolRiskConfigurationAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolRiskConfigurationAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolRiskConfigurationAttachmentResult, error) {
			args := v.(LookupUserPoolRiskConfigurationAttachmentArgs)
			r, err := LookupUserPoolRiskConfigurationAttachment(ctx, &args, opts...)
			var s LookupUserPoolRiskConfigurationAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserPoolRiskConfigurationAttachmentResultOutput)
}

type LookupUserPoolRiskConfigurationAttachmentOutputArgs struct {
	// The app client ID. You can specify the risk configuration for a single client (with a specific ClientId) or for all clients (by setting the ClientId to `ALL` ).
	ClientId pulumi.StringInput `pulumi:"clientId"`
	// The user pool ID.
	UserPoolId pulumi.StringInput `pulumi:"userPoolId"`
}

func (LookupUserPoolRiskConfigurationAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolRiskConfigurationAttachmentArgs)(nil)).Elem()
}

type LookupUserPoolRiskConfigurationAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupUserPoolRiskConfigurationAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolRiskConfigurationAttachmentResult)(nil)).Elem()
}

func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) ToLookupUserPoolRiskConfigurationAttachmentResultOutput() LookupUserPoolRiskConfigurationAttachmentResultOutput {
	return o
}

func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) ToLookupUserPoolRiskConfigurationAttachmentResultOutputWithContext(ctx context.Context) LookupUserPoolRiskConfigurationAttachmentResultOutput {
	return o
}

// The account takeover risk configuration object, including the `NotifyConfiguration` object and `Actions` to take if there is an account takeover.
func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) AccountTakeoverRiskConfiguration() UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationTypePtrOutput {
	return o.ApplyT(func(v LookupUserPoolRiskConfigurationAttachmentResult) *UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationType {
		return v.AccountTakeoverRiskConfiguration
	}).(UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationTypePtrOutput)
}

// The compromised credentials risk configuration object, including the `EventFilter` and the `EventAction` .
func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) CompromisedCredentialsRiskConfiguration() UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationTypePtrOutput {
	return o.ApplyT(func(v LookupUserPoolRiskConfigurationAttachmentResult) *UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationType {
		return v.CompromisedCredentialsRiskConfiguration
	}).(UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationTypePtrOutput)
}

// The configuration to override the risk decision.
func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) RiskExceptionConfiguration() UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypePtrOutput {
	return o.ApplyT(func(v LookupUserPoolRiskConfigurationAttachmentResult) *UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationType {
		return v.RiskExceptionConfiguration
	}).(UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolRiskConfigurationAttachmentResultOutput{})
}
