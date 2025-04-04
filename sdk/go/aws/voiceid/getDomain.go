// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package voiceid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::VoiceID::Domain resource specifies an Amazon VoiceID Domain.
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainResult
	err := ctx.Invoke("aws-native:voiceid:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainArgs struct {
	// The identifier of the domain.
	DomainId string `pulumi:"domainId"`
}

type LookupDomainResult struct {
	// The description of the domain.
	Description *string `pulumi:"description"`
	// The identifier of the domain.
	DomainId *string `pulumi:"domainId"`
	// The name for the domain.
	Name *string `pulumi:"name"`
	// The server-side encryption configuration containing the KMS key identifier you want Voice ID to use to encrypt your data.
	ServerSideEncryptionConfiguration *DomainServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	// The tags used to organize, track, or control access for this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDomainResultOutput, error) {
			args := v.(LookupDomainArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:voiceid:getDomain", args, LookupDomainResultOutput{}, options).(LookupDomainResultOutput), nil
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	// The identifier of the domain.
	DomainId pulumi.StringInput `pulumi:"domainId"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

// The description of the domain.
func (o LookupDomainResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the domain.
func (o LookupDomainResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

// The name for the domain.
func (o LookupDomainResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The server-side encryption configuration containing the KMS key identifier you want Voice ID to use to encrypt your data.
func (o LookupDomainResultOutput) ServerSideEncryptionConfiguration() DomainServerSideEncryptionConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *DomainServerSideEncryptionConfiguration {
		return v.ServerSideEncryptionConfiguration
	}).(DomainServerSideEncryptionConfigurationPtrOutput)
}

// The tags used to organize, track, or control access for this resource.
func (o LookupDomainResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
