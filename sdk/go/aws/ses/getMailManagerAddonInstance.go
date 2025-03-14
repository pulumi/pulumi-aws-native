// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::SES::MailManagerAddonInstance Resource Type
func LookupMailManagerAddonInstance(ctx *pulumi.Context, args *LookupMailManagerAddonInstanceArgs, opts ...pulumi.InvokeOption) (*LookupMailManagerAddonInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMailManagerAddonInstanceResult
	err := ctx.Invoke("aws-native:ses:getMailManagerAddonInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMailManagerAddonInstanceArgs struct {
	// The unique ID of the Add On instance.
	AddonInstanceId string `pulumi:"addonInstanceId"`
}

type LookupMailManagerAddonInstanceResult struct {
	// The Amazon Resource Name (ARN) of the Add On instance.
	AddonInstanceArn *string `pulumi:"addonInstanceArn"`
	// The unique ID of the Add On instance.
	AddonInstanceId *string `pulumi:"addonInstanceId"`
	// The name of the Add On for the instance.
	AddonName *string `pulumi:"addonName"`
	// The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupMailManagerAddonInstanceOutput(ctx *pulumi.Context, args LookupMailManagerAddonInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupMailManagerAddonInstanceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMailManagerAddonInstanceResultOutput, error) {
			args := v.(LookupMailManagerAddonInstanceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ses:getMailManagerAddonInstance", args, LookupMailManagerAddonInstanceResultOutput{}, options).(LookupMailManagerAddonInstanceResultOutput), nil
		}).(LookupMailManagerAddonInstanceResultOutput)
}

type LookupMailManagerAddonInstanceOutputArgs struct {
	// The unique ID of the Add On instance.
	AddonInstanceId pulumi.StringInput `pulumi:"addonInstanceId"`
}

func (LookupMailManagerAddonInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMailManagerAddonInstanceArgs)(nil)).Elem()
}

type LookupMailManagerAddonInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupMailManagerAddonInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMailManagerAddonInstanceResult)(nil)).Elem()
}

func (o LookupMailManagerAddonInstanceResultOutput) ToLookupMailManagerAddonInstanceResultOutput() LookupMailManagerAddonInstanceResultOutput {
	return o
}

func (o LookupMailManagerAddonInstanceResultOutput) ToLookupMailManagerAddonInstanceResultOutputWithContext(ctx context.Context) LookupMailManagerAddonInstanceResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Add On instance.
func (o LookupMailManagerAddonInstanceResultOutput) AddonInstanceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMailManagerAddonInstanceResult) *string { return v.AddonInstanceArn }).(pulumi.StringPtrOutput)
}

// The unique ID of the Add On instance.
func (o LookupMailManagerAddonInstanceResultOutput) AddonInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMailManagerAddonInstanceResult) *string { return v.AddonInstanceId }).(pulumi.StringPtrOutput)
}

// The name of the Add On for the instance.
func (o LookupMailManagerAddonInstanceResultOutput) AddonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMailManagerAddonInstanceResult) *string { return v.AddonName }).(pulumi.StringPtrOutput)
}

// The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
func (o LookupMailManagerAddonInstanceResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupMailManagerAddonInstanceResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMailManagerAddonInstanceResultOutput{})
}
