// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Publishes new or first hook version to AWS CloudFormation Registry.
func LookupHookVersion(ctx *pulumi.Context, args *LookupHookVersionArgs, opts ...pulumi.InvokeOption) (*LookupHookVersionResult, error) {
	var rv LookupHookVersionResult
	err := ctx.Invoke("aws-native:cloudformation:getHookVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHookVersionArgs struct {
	// The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource
	Arn string `pulumi:"arn"`
}

type LookupHookVersionResult struct {
	// The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource
	Arn *string `pulumi:"arn"`
	// Indicates if this type version is the current default version
	IsDefaultVersion *bool `pulumi:"isDefaultVersion"`
	// The Amazon Resource Name (ARN) of the type without the versionID.
	TypeArn *string `pulumi:"typeArn"`
	// The ID of the version of the type represented by this hook instance.
	VersionId *string `pulumi:"versionId"`
	// The scope at which the type is visible and usable in CloudFormation operations.
	//
	// Valid values include:
	//
	// PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
	//
	// PUBLIC: The type is publically visible and usable within any Amazon account.
	Visibility *HookVersionVisibility `pulumi:"visibility"`
}

func LookupHookVersionOutput(ctx *pulumi.Context, args LookupHookVersionOutputArgs, opts ...pulumi.InvokeOption) LookupHookVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHookVersionResult, error) {
			args := v.(LookupHookVersionArgs)
			r, err := LookupHookVersion(ctx, &args, opts...)
			var s LookupHookVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHookVersionResultOutput)
}

type LookupHookVersionOutputArgs struct {
	// The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupHookVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHookVersionArgs)(nil)).Elem()
}

type LookupHookVersionResultOutput struct{ *pulumi.OutputState }

func (LookupHookVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHookVersionResult)(nil)).Elem()
}

func (o LookupHookVersionResultOutput) ToLookupHookVersionResultOutput() LookupHookVersionResultOutput {
	return o
}

func (o LookupHookVersionResultOutput) ToLookupHookVersionResultOutputWithContext(ctx context.Context) LookupHookVersionResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource
func (o LookupHookVersionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookVersionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Indicates if this type version is the current default version
func (o LookupHookVersionResultOutput) IsDefaultVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHookVersionResult) *bool { return v.IsDefaultVersion }).(pulumi.BoolPtrOutput)
}

// The Amazon Resource Name (ARN) of the type without the versionID.
func (o LookupHookVersionResultOutput) TypeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookVersionResult) *string { return v.TypeArn }).(pulumi.StringPtrOutput)
}

// The ID of the version of the type represented by this hook instance.
func (o LookupHookVersionResultOutput) VersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookVersionResult) *string { return v.VersionId }).(pulumi.StringPtrOutput)
}

// The scope at which the type is visible and usable in CloudFormation operations.
//
// Valid values include:
//
// PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
//
// PUBLIC: The type is publically visible and usable within any Amazon account.
func (o LookupHookVersionResultOutput) Visibility() HookVersionVisibilityPtrOutput {
	return o.ApplyT(func(v LookupHookVersionResult) *HookVersionVisibility { return v.Visibility }).(HookVersionVisibilityPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHookVersionResultOutput{})
}