// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::GameLift::Location resource creates an Amazon GameLift (GameLift) custom location.
func LookupLocation(ctx *pulumi.Context, args *LookupLocationArgs, opts ...pulumi.InvokeOption) (*LookupLocationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocationResult
	err := ctx.Invoke("aws-native:gamelift:getLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocationArgs struct {
	// A descriptive name for the custom location.
	LocationName string `pulumi:"locationName"`
}

type LookupLocationResult struct {
	// A unique identifier for the custom location. For example, `arn:aws:gamelift:[region]::location/location-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912` .
	LocationArn *string `pulumi:"locationArn"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupLocationOutput(ctx *pulumi.Context, args LookupLocationOutputArgs, opts ...pulumi.InvokeOption) LookupLocationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLocationResultOutput, error) {
			args := v.(LookupLocationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:gamelift:getLocation", args, LookupLocationResultOutput{}, options).(LookupLocationResultOutput), nil
		}).(LookupLocationResultOutput)
}

type LookupLocationOutputArgs struct {
	// A descriptive name for the custom location.
	LocationName pulumi.StringInput `pulumi:"locationName"`
}

func (LookupLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationArgs)(nil)).Elem()
}

type LookupLocationResultOutput struct{ *pulumi.OutputState }

func (LookupLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationResult)(nil)).Elem()
}

func (o LookupLocationResultOutput) ToLookupLocationResultOutput() LookupLocationResultOutput {
	return o
}

func (o LookupLocationResultOutput) ToLookupLocationResultOutputWithContext(ctx context.Context) LookupLocationResultOutput {
	return o
}

// A unique identifier for the custom location. For example, `arn:aws:gamelift:[region]::location/location-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912` .
func (o LookupLocationResultOutput) LocationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationResult) *string { return v.LocationArn }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLocationResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupLocationResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocationResultOutput{})
}
