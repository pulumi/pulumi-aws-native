// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transfer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Transfer::Certificate
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("aws-native:transfer:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	// A unique identifier for the certificate.
	CertificateId string `pulumi:"certificateId"`
}

type LookupCertificateResult struct {
	// Specifies the active date for the certificate.
	ActiveDate *string `pulumi:"activeDate"`
	// Specifies the unique Amazon Resource Name (ARN) for the agreement.
	Arn *string `pulumi:"arn"`
	// A unique identifier for the certificate.
	CertificateId *string `pulumi:"certificateId"`
	// A textual description for the certificate.
	Description *string `pulumi:"description"`
	// Specifies the inactive date for the certificate.
	InactiveDate *string `pulumi:"inactiveDate"`
	// Specifies the not after date for the certificate.
	NotAfterDate *string `pulumi:"notAfterDate"`
	// Specifies the not before date for the certificate.
	NotBeforeDate *string `pulumi:"notBeforeDate"`
	// Specifies Certificate's serial.
	Serial *string `pulumi:"serial"`
	// A status description for the certificate.
	Status *CertificateStatus `pulumi:"status"`
	// Key-value pairs that can be used to group and search for certificates. Tags are metadata attached to certificates for any purpose.
	Tags []CertificateTag `pulumi:"tags"`
	// Describing the type of certificate. With or without a private key.
	Type *CertificateType `pulumi:"type"`
	// Specifies the usage type for the certificate.
	Usage *CertificateUsage `pulumi:"usage"`
}

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateResult, error) {
			args := v.(LookupCertificateArgs)
			r, err := LookupCertificate(ctx, &args, opts...)
			var s LookupCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCertificateResultOutput)
}

type LookupCertificateOutputArgs struct {
	// A unique identifier for the certificate.
	CertificateId pulumi.StringInput `pulumi:"certificateId"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

// Specifies the active date for the certificate.
func (o LookupCertificateResultOutput) ActiveDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.ActiveDate }).(pulumi.StringPtrOutput)
}

// Specifies the unique Amazon Resource Name (ARN) for the agreement.
func (o LookupCertificateResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// A unique identifier for the certificate.
func (o LookupCertificateResultOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.CertificateId }).(pulumi.StringPtrOutput)
}

// A textual description for the certificate.
func (o LookupCertificateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the inactive date for the certificate.
func (o LookupCertificateResultOutput) InactiveDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.InactiveDate }).(pulumi.StringPtrOutput)
}

// Specifies the not after date for the certificate.
func (o LookupCertificateResultOutput) NotAfterDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.NotAfterDate }).(pulumi.StringPtrOutput)
}

// Specifies the not before date for the certificate.
func (o LookupCertificateResultOutput) NotBeforeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.NotBeforeDate }).(pulumi.StringPtrOutput)
}

// Specifies Certificate's serial.
func (o LookupCertificateResultOutput) Serial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Serial }).(pulumi.StringPtrOutput)
}

// A status description for the certificate.
func (o LookupCertificateResultOutput) Status() CertificateStatusPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *CertificateStatus { return v.Status }).(CertificateStatusPtrOutput)
}

// Key-value pairs that can be used to group and search for certificates. Tags are metadata attached to certificates for any purpose.
func (o LookupCertificateResultOutput) Tags() CertificateTagArrayOutput {
	return o.ApplyT(func(v LookupCertificateResult) []CertificateTag { return v.Tags }).(CertificateTagArrayOutput)
}

// Describing the type of certificate. With or without a private key.
func (o LookupCertificateResultOutput) Type() CertificateTypePtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *CertificateType { return v.Type }).(CertificateTypePtrOutput)
}

// Specifies the usage type for the certificate.
func (o LookupCertificateResultOutput) Usage() CertificateUsagePtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *CertificateUsage { return v.Usage }).(CertificateUsagePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}