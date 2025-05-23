// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspacesweb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::WorkSpacesWeb::DataProtectionSettings Resource Type
func LookupDataProtectionSettings(ctx *pulumi.Context, args *LookupDataProtectionSettingsArgs, opts ...pulumi.InvokeOption) (*LookupDataProtectionSettingsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDataProtectionSettingsResult
	err := ctx.Invoke("aws-native:workspacesweb:getDataProtectionSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataProtectionSettingsArgs struct {
	// The ARN of the data protection settings resource.
	DataProtectionSettingsArn string `pulumi:"dataProtectionSettingsArn"`
}

type LookupDataProtectionSettingsResult struct {
	// A list of web portal ARNs that this data protection settings resource is associated with.
	AssociatedPortalArns []string `pulumi:"associatedPortalArns"`
	// The creation date timestamp of the data protection settings.
	CreationDate *string `pulumi:"creationDate"`
	// The ARN of the data protection settings resource.
	DataProtectionSettingsArn *string `pulumi:"dataProtectionSettingsArn"`
	// The description of the data protection settings.
	Description *string `pulumi:"description"`
	// The display name of the data protection settings.
	DisplayName *string `pulumi:"displayName"`
	// The inline redaction configuration for the data protection settings.
	InlineRedactionConfiguration *DataProtectionSettingsInlineRedactionConfiguration `pulumi:"inlineRedactionConfiguration"`
	// The tags of the data protection settings.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupDataProtectionSettingsOutput(ctx *pulumi.Context, args LookupDataProtectionSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupDataProtectionSettingsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataProtectionSettingsResultOutput, error) {
			args := v.(LookupDataProtectionSettingsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:workspacesweb:getDataProtectionSettings", args, LookupDataProtectionSettingsResultOutput{}, options).(LookupDataProtectionSettingsResultOutput), nil
		}).(LookupDataProtectionSettingsResultOutput)
}

type LookupDataProtectionSettingsOutputArgs struct {
	// The ARN of the data protection settings resource.
	DataProtectionSettingsArn pulumi.StringInput `pulumi:"dataProtectionSettingsArn"`
}

func (LookupDataProtectionSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataProtectionSettingsArgs)(nil)).Elem()
}

type LookupDataProtectionSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupDataProtectionSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataProtectionSettingsResult)(nil)).Elem()
}

func (o LookupDataProtectionSettingsResultOutput) ToLookupDataProtectionSettingsResultOutput() LookupDataProtectionSettingsResultOutput {
	return o
}

func (o LookupDataProtectionSettingsResultOutput) ToLookupDataProtectionSettingsResultOutputWithContext(ctx context.Context) LookupDataProtectionSettingsResultOutput {
	return o
}

// A list of web portal ARNs that this data protection settings resource is associated with.
func (o LookupDataProtectionSettingsResultOutput) AssociatedPortalArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataProtectionSettingsResult) []string { return v.AssociatedPortalArns }).(pulumi.StringArrayOutput)
}

// The creation date timestamp of the data protection settings.
func (o LookupDataProtectionSettingsResultOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataProtectionSettingsResult) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

// The ARN of the data protection settings resource.
func (o LookupDataProtectionSettingsResultOutput) DataProtectionSettingsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataProtectionSettingsResult) *string { return v.DataProtectionSettingsArn }).(pulumi.StringPtrOutput)
}

// The description of the data protection settings.
func (o LookupDataProtectionSettingsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataProtectionSettingsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the data protection settings.
func (o LookupDataProtectionSettingsResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataProtectionSettingsResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The inline redaction configuration for the data protection settings.
func (o LookupDataProtectionSettingsResultOutput) InlineRedactionConfiguration() DataProtectionSettingsInlineRedactionConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDataProtectionSettingsResult) *DataProtectionSettingsInlineRedactionConfiguration {
		return v.InlineRedactionConfiguration
	}).(DataProtectionSettingsInlineRedactionConfigurationPtrOutput)
}

// The tags of the data protection settings.
func (o LookupDataProtectionSettingsResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupDataProtectionSettingsResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataProtectionSettingsResultOutput{})
}
