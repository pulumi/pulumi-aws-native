// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The default version of a resource that has been registered in the CloudFormation Registry.
func LookupResourceDefaultVersion(ctx *pulumi.Context, args *LookupResourceDefaultVersionArgs, opts ...pulumi.InvokeOption) (*LookupResourceDefaultVersionResult, error) {
	var rv LookupResourceDefaultVersionResult
	err := ctx.Invoke("aws-native:cloudformation:getResourceDefaultVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceDefaultVersionArgs struct {
	// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a ResourceDefaultVersion
	Arn string `pulumi:"arn"`
}

type LookupResourceDefaultVersionResult struct {
	// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a ResourceDefaultVersion
	Arn *string `pulumi:"arn"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName *string `pulumi:"typeName"`
	// The Amazon Resource Name (ARN) of the type version.
	TypeVersionArn *string `pulumi:"typeVersionArn"`
	// The ID of an existing version of the resource to set as the default.
	VersionId *string `pulumi:"versionId"`
}

func LookupResourceDefaultVersionOutput(ctx *pulumi.Context, args LookupResourceDefaultVersionOutputArgs, opts ...pulumi.InvokeOption) LookupResourceDefaultVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceDefaultVersionResult, error) {
			args := v.(LookupResourceDefaultVersionArgs)
			r, err := LookupResourceDefaultVersion(ctx, &args, opts...)
			var s LookupResourceDefaultVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceDefaultVersionResultOutput)
}

type LookupResourceDefaultVersionOutputArgs struct {
	// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a ResourceDefaultVersion
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupResourceDefaultVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceDefaultVersionArgs)(nil)).Elem()
}

type LookupResourceDefaultVersionResultOutput struct{ *pulumi.OutputState }

func (LookupResourceDefaultVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceDefaultVersionResult)(nil)).Elem()
}

func (o LookupResourceDefaultVersionResultOutput) ToLookupResourceDefaultVersionResultOutput() LookupResourceDefaultVersionResultOutput {
	return o
}

func (o LookupResourceDefaultVersionResultOutput) ToLookupResourceDefaultVersionResultOutputWithContext(ctx context.Context) LookupResourceDefaultVersionResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a ResourceDefaultVersion
func (o LookupResourceDefaultVersionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceDefaultVersionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The name of the type being registered.
//
// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
func (o LookupResourceDefaultVersionResultOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceDefaultVersionResult) *string { return v.TypeName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the type version.
func (o LookupResourceDefaultVersionResultOutput) TypeVersionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceDefaultVersionResult) *string { return v.TypeVersionArn }).(pulumi.StringPtrOutput)
}

// The ID of an existing version of the resource to set as the default.
func (o LookupResourceDefaultVersionResultOutput) VersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceDefaultVersionResult) *string { return v.VersionId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceDefaultVersionResultOutput{})
}