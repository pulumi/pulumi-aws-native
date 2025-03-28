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
type DataProtectionSettings struct {
	pulumi.CustomResourceState

	// The additional encryption context of the data protection settings.
	AdditionalEncryptionContext pulumi.StringMapOutput `pulumi:"additionalEncryptionContext"`
	// A list of web portal ARNs that this data protection settings resource is associated with.
	AssociatedPortalArns pulumi.StringArrayOutput `pulumi:"associatedPortalArns"`
	// The creation date timestamp of the data protection settings.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The customer managed key used to encrypt sensitive information in the data protection settings.
	CustomerManagedKey pulumi.StringPtrOutput `pulumi:"customerManagedKey"`
	// The ARN of the data protection settings resource.
	DataProtectionSettingsArn pulumi.StringOutput `pulumi:"dataProtectionSettingsArn"`
	// The description of the data protection settings.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the data protection settings.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The inline redaction configuration for the data protection settings.
	InlineRedactionConfiguration DataProtectionSettingsInlineRedactionConfigurationPtrOutput `pulumi:"inlineRedactionConfiguration"`
	// The tags of the data protection settings.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewDataProtectionSettings registers a new resource with the given unique name, arguments, and options.
func NewDataProtectionSettings(ctx *pulumi.Context,
	name string, args *DataProtectionSettingsArgs, opts ...pulumi.ResourceOption) (*DataProtectionSettings, error) {
	if args == nil {
		args = &DataProtectionSettingsArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"additionalEncryptionContext.*",
		"customerManagedKey",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataProtectionSettings
	err := ctx.RegisterResource("aws-native:workspacesweb:DataProtectionSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataProtectionSettings gets an existing DataProtectionSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataProtectionSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataProtectionSettingsState, opts ...pulumi.ResourceOption) (*DataProtectionSettings, error) {
	var resource DataProtectionSettings
	err := ctx.ReadResource("aws-native:workspacesweb:DataProtectionSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataProtectionSettings resources.
type dataProtectionSettingsState struct {
}

type DataProtectionSettingsState struct {
}

func (DataProtectionSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataProtectionSettingsState)(nil)).Elem()
}

type dataProtectionSettingsArgs struct {
	// The additional encryption context of the data protection settings.
	AdditionalEncryptionContext map[string]string `pulumi:"additionalEncryptionContext"`
	// The customer managed key used to encrypt sensitive information in the data protection settings.
	CustomerManagedKey *string `pulumi:"customerManagedKey"`
	// The description of the data protection settings.
	Description *string `pulumi:"description"`
	// The display name of the data protection settings.
	DisplayName *string `pulumi:"displayName"`
	// The inline redaction configuration for the data protection settings.
	InlineRedactionConfiguration *DataProtectionSettingsInlineRedactionConfiguration `pulumi:"inlineRedactionConfiguration"`
	// The tags of the data protection settings.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a DataProtectionSettings resource.
type DataProtectionSettingsArgs struct {
	// The additional encryption context of the data protection settings.
	AdditionalEncryptionContext pulumi.StringMapInput
	// The customer managed key used to encrypt sensitive information in the data protection settings.
	CustomerManagedKey pulumi.StringPtrInput
	// The description of the data protection settings.
	Description pulumi.StringPtrInput
	// The display name of the data protection settings.
	DisplayName pulumi.StringPtrInput
	// The inline redaction configuration for the data protection settings.
	InlineRedactionConfiguration DataProtectionSettingsInlineRedactionConfigurationPtrInput
	// The tags of the data protection settings.
	Tags aws.TagArrayInput
}

func (DataProtectionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataProtectionSettingsArgs)(nil)).Elem()
}

type DataProtectionSettingsInput interface {
	pulumi.Input

	ToDataProtectionSettingsOutput() DataProtectionSettingsOutput
	ToDataProtectionSettingsOutputWithContext(ctx context.Context) DataProtectionSettingsOutput
}

func (*DataProtectionSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**DataProtectionSettings)(nil)).Elem()
}

func (i *DataProtectionSettings) ToDataProtectionSettingsOutput() DataProtectionSettingsOutput {
	return i.ToDataProtectionSettingsOutputWithContext(context.Background())
}

func (i *DataProtectionSettings) ToDataProtectionSettingsOutputWithContext(ctx context.Context) DataProtectionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataProtectionSettingsOutput)
}

type DataProtectionSettingsOutput struct{ *pulumi.OutputState }

func (DataProtectionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataProtectionSettings)(nil)).Elem()
}

func (o DataProtectionSettingsOutput) ToDataProtectionSettingsOutput() DataProtectionSettingsOutput {
	return o
}

func (o DataProtectionSettingsOutput) ToDataProtectionSettingsOutputWithContext(ctx context.Context) DataProtectionSettingsOutput {
	return o
}

// The additional encryption context of the data protection settings.
func (o DataProtectionSettingsOutput) AdditionalEncryptionContext() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataProtectionSettings) pulumi.StringMapOutput { return v.AdditionalEncryptionContext }).(pulumi.StringMapOutput)
}

// A list of web portal ARNs that this data protection settings resource is associated with.
func (o DataProtectionSettingsOutput) AssociatedPortalArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataProtectionSettings) pulumi.StringArrayOutput { return v.AssociatedPortalArns }).(pulumi.StringArrayOutput)
}

// The creation date timestamp of the data protection settings.
func (o DataProtectionSettingsOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProtectionSettings) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The customer managed key used to encrypt sensitive information in the data protection settings.
func (o DataProtectionSettingsOutput) CustomerManagedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataProtectionSettings) pulumi.StringPtrOutput { return v.CustomerManagedKey }).(pulumi.StringPtrOutput)
}

// The ARN of the data protection settings resource.
func (o DataProtectionSettingsOutput) DataProtectionSettingsArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProtectionSettings) pulumi.StringOutput { return v.DataProtectionSettingsArn }).(pulumi.StringOutput)
}

// The description of the data protection settings.
func (o DataProtectionSettingsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataProtectionSettings) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the data protection settings.
func (o DataProtectionSettingsOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataProtectionSettings) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The inline redaction configuration for the data protection settings.
func (o DataProtectionSettingsOutput) InlineRedactionConfiguration() DataProtectionSettingsInlineRedactionConfigurationPtrOutput {
	return o.ApplyT(func(v *DataProtectionSettings) DataProtectionSettingsInlineRedactionConfigurationPtrOutput {
		return v.InlineRedactionConfiguration
	}).(DataProtectionSettingsInlineRedactionConfigurationPtrOutput)
}

// The tags of the data protection settings.
func (o DataProtectionSettingsOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *DataProtectionSettings) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataProtectionSettingsInput)(nil)).Elem(), &DataProtectionSettings{})
	pulumi.RegisterOutputType(DataProtectionSettingsOutput{})
}
