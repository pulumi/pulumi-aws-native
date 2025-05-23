// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ivs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IVS::RecordingConfiguration
func LookupRecordingConfiguration(ctx *pulumi.Context, args *LookupRecordingConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupRecordingConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRecordingConfigurationResult
	err := ctx.Invoke("aws-native:ivs:getRecordingConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRecordingConfigurationArgs struct {
	// Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.
	Arn string `pulumi:"arn"`
}

type LookupRecordingConfigurationResult struct {
	// Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.
	Arn *string `pulumi:"arn"`
	// Recording Configuration State.
	State *RecordingConfigurationStateEnum `pulumi:"state"`
	// A list of key-value pairs that contain metadata for the asset model.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupRecordingConfigurationOutput(ctx *pulumi.Context, args LookupRecordingConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupRecordingConfigurationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRecordingConfigurationResultOutput, error) {
			args := v.(LookupRecordingConfigurationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ivs:getRecordingConfiguration", args, LookupRecordingConfigurationResultOutput{}, options).(LookupRecordingConfigurationResultOutput), nil
		}).(LookupRecordingConfigurationResultOutput)
}

type LookupRecordingConfigurationOutputArgs struct {
	// Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupRecordingConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordingConfigurationArgs)(nil)).Elem()
}

type LookupRecordingConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupRecordingConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordingConfigurationResult)(nil)).Elem()
}

func (o LookupRecordingConfigurationResultOutput) ToLookupRecordingConfigurationResultOutput() LookupRecordingConfigurationResultOutput {
	return o
}

func (o LookupRecordingConfigurationResultOutput) ToLookupRecordingConfigurationResultOutputWithContext(ctx context.Context) LookupRecordingConfigurationResultOutput {
	return o
}

// Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.
func (o LookupRecordingConfigurationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordingConfigurationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Recording Configuration State.
func (o LookupRecordingConfigurationResultOutput) State() RecordingConfigurationStateEnumPtrOutput {
	return o.ApplyT(func(v LookupRecordingConfigurationResult) *RecordingConfigurationStateEnum { return v.State }).(RecordingConfigurationStateEnumPtrOutput)
}

// A list of key-value pairs that contain metadata for the asset model.
func (o LookupRecordingConfigurationResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupRecordingConfigurationResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRecordingConfigurationResultOutput{})
}
