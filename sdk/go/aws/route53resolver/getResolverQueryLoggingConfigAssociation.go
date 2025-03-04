// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation.
func LookupResolverQueryLoggingConfigAssociation(ctx *pulumi.Context, args *LookupResolverQueryLoggingConfigAssociationArgs, opts ...pulumi.InvokeOption) (*LookupResolverQueryLoggingConfigAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResolverQueryLoggingConfigAssociationResult
	err := ctx.Invoke("aws-native:route53resolver:getResolverQueryLoggingConfigAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResolverQueryLoggingConfigAssociationArgs struct {
	// Id
	Id string `pulumi:"id"`
}

type LookupResolverQueryLoggingConfigAssociationResult struct {
	// Rfc3339TimeString
	CreationTime *string `pulumi:"creationTime"`
	// ResolverQueryLogConfigAssociationError
	Error *ResolverQueryLoggingConfigAssociationError `pulumi:"error"`
	// ResolverQueryLogConfigAssociationErrorMessage
	ErrorMessage *string `pulumi:"errorMessage"`
	// Id
	Id *string `pulumi:"id"`
	// ResolverQueryLogConfigAssociationStatus
	Status *ResolverQueryLoggingConfigAssociationStatus `pulumi:"status"`
}

func LookupResolverQueryLoggingConfigAssociationOutput(ctx *pulumi.Context, args LookupResolverQueryLoggingConfigAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupResolverQueryLoggingConfigAssociationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupResolverQueryLoggingConfigAssociationResultOutput, error) {
			args := v.(LookupResolverQueryLoggingConfigAssociationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:route53resolver:getResolverQueryLoggingConfigAssociation", args, LookupResolverQueryLoggingConfigAssociationResultOutput{}, options).(LookupResolverQueryLoggingConfigAssociationResultOutput), nil
		}).(LookupResolverQueryLoggingConfigAssociationResultOutput)
}

type LookupResolverQueryLoggingConfigAssociationOutputArgs struct {
	// Id
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupResolverQueryLoggingConfigAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverQueryLoggingConfigAssociationArgs)(nil)).Elem()
}

type LookupResolverQueryLoggingConfigAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupResolverQueryLoggingConfigAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverQueryLoggingConfigAssociationResult)(nil)).Elem()
}

func (o LookupResolverQueryLoggingConfigAssociationResultOutput) ToLookupResolverQueryLoggingConfigAssociationResultOutput() LookupResolverQueryLoggingConfigAssociationResultOutput {
	return o
}

func (o LookupResolverQueryLoggingConfigAssociationResultOutput) ToLookupResolverQueryLoggingConfigAssociationResultOutputWithContext(ctx context.Context) LookupResolverQueryLoggingConfigAssociationResultOutput {
	return o
}

// Rfc3339TimeString
func (o LookupResolverQueryLoggingConfigAssociationResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigAssociationResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// ResolverQueryLogConfigAssociationError
func (o LookupResolverQueryLoggingConfigAssociationResultOutput) Error() ResolverQueryLoggingConfigAssociationErrorPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigAssociationResult) *ResolverQueryLoggingConfigAssociationError {
		return v.Error
	}).(ResolverQueryLoggingConfigAssociationErrorPtrOutput)
}

// ResolverQueryLogConfigAssociationErrorMessage
func (o LookupResolverQueryLoggingConfigAssociationResultOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigAssociationResult) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

// Id
func (o LookupResolverQueryLoggingConfigAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// ResolverQueryLogConfigAssociationStatus
func (o LookupResolverQueryLoggingConfigAssociationResultOutput) Status() ResolverQueryLoggingConfigAssociationStatusPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigAssociationResult) *ResolverQueryLoggingConfigAssociationStatus {
		return v.Status
	}).(ResolverQueryLoggingConfigAssociationStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResolverQueryLoggingConfigAssociationResultOutput{})
}
