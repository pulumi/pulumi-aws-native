// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Route53::CidrCollection.
func LookupCidrCollection(ctx *pulumi.Context, args *LookupCidrCollectionArgs, opts ...pulumi.InvokeOption) (*LookupCidrCollectionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCidrCollectionResult
	err := ctx.Invoke("aws-native:route53:getCidrCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCidrCollectionArgs struct {
	// UUID of the CIDR collection.
	Id string `pulumi:"id"`
}

type LookupCidrCollectionResult struct {
	// The Amazon resource name (ARN) to uniquely identify the AWS resource.
	Arn *string `pulumi:"arn"`
	// UUID of the CIDR collection.
	Id *string `pulumi:"id"`
	// A complex type that contains information about the list of CIDR locations.
	Locations []CidrCollectionLocation `pulumi:"locations"`
}

func LookupCidrCollectionOutput(ctx *pulumi.Context, args LookupCidrCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupCidrCollectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCidrCollectionResultOutput, error) {
			args := v.(LookupCidrCollectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:route53:getCidrCollection", args, LookupCidrCollectionResultOutput{}, options).(LookupCidrCollectionResultOutput), nil
		}).(LookupCidrCollectionResultOutput)
}

type LookupCidrCollectionOutputArgs struct {
	// UUID of the CIDR collection.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupCidrCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCidrCollectionArgs)(nil)).Elem()
}

type LookupCidrCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupCidrCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCidrCollectionResult)(nil)).Elem()
}

func (o LookupCidrCollectionResultOutput) ToLookupCidrCollectionResultOutput() LookupCidrCollectionResultOutput {
	return o
}

func (o LookupCidrCollectionResultOutput) ToLookupCidrCollectionResultOutputWithContext(ctx context.Context) LookupCidrCollectionResultOutput {
	return o
}

// The Amazon resource name (ARN) to uniquely identify the AWS resource.
func (o LookupCidrCollectionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCidrCollectionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// UUID of the CIDR collection.
func (o LookupCidrCollectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCidrCollectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// A complex type that contains information about the list of CIDR locations.
func (o LookupCidrCollectionResultOutput) Locations() CidrCollectionLocationArrayOutput {
	return o.ApplyT(func(v LookupCidrCollectionResult) []CidrCollectionLocation { return v.Locations }).(CidrCollectionLocationArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCidrCollectionResultOutput{})
}
