// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// resource definition
type SoftwarePackageVersion struct {
	pulumi.CustomResourceState

	Attributes        SoftwarePackageVersionResourceAttributesPtrOutput `pulumi:"attributes"`
	Description       pulumi.StringPtrOutput                            `pulumi:"description"`
	ErrorReason       pulumi.StringOutput                               `pulumi:"errorReason"`
	PackageName       pulumi.StringOutput                               `pulumi:"packageName"`
	PackageVersionArn pulumi.StringOutput                               `pulumi:"packageVersionArn"`
	Status            SoftwarePackageVersionPackageVersionStatusOutput  `pulumi:"status"`
	// An array of key-value pairs to apply to this resource.
	Tags        SoftwarePackageVersionTagArrayOutput `pulumi:"tags"`
	VersionName pulumi.StringPtrOutput               `pulumi:"versionName"`
}

// NewSoftwarePackageVersion registers a new resource with the given unique name, arguments, and options.
func NewSoftwarePackageVersion(ctx *pulumi.Context,
	name string, args *SoftwarePackageVersionArgs, opts ...pulumi.ResourceOption) (*SoftwarePackageVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PackageName == nil {
		return nil, errors.New("invalid value for required argument 'PackageName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"packageName",
		"versionName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SoftwarePackageVersion
	err := ctx.RegisterResource("aws-native:iot:SoftwarePackageVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSoftwarePackageVersion gets an existing SoftwarePackageVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSoftwarePackageVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SoftwarePackageVersionState, opts ...pulumi.ResourceOption) (*SoftwarePackageVersion, error) {
	var resource SoftwarePackageVersion
	err := ctx.ReadResource("aws-native:iot:SoftwarePackageVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SoftwarePackageVersion resources.
type softwarePackageVersionState struct {
}

type SoftwarePackageVersionState struct {
}

func (SoftwarePackageVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*softwarePackageVersionState)(nil)).Elem()
}

type softwarePackageVersionArgs struct {
	Attributes  *SoftwarePackageVersionResourceAttributes `pulumi:"attributes"`
	Description *string                                   `pulumi:"description"`
	PackageName string                                    `pulumi:"packageName"`
	// An array of key-value pairs to apply to this resource.
	Tags        []SoftwarePackageVersionTag `pulumi:"tags"`
	VersionName *string                     `pulumi:"versionName"`
}

// The set of arguments for constructing a SoftwarePackageVersion resource.
type SoftwarePackageVersionArgs struct {
	Attributes  SoftwarePackageVersionResourceAttributesPtrInput
	Description pulumi.StringPtrInput
	PackageName pulumi.StringInput
	// An array of key-value pairs to apply to this resource.
	Tags        SoftwarePackageVersionTagArrayInput
	VersionName pulumi.StringPtrInput
}

func (SoftwarePackageVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*softwarePackageVersionArgs)(nil)).Elem()
}

type SoftwarePackageVersionInput interface {
	pulumi.Input

	ToSoftwarePackageVersionOutput() SoftwarePackageVersionOutput
	ToSoftwarePackageVersionOutputWithContext(ctx context.Context) SoftwarePackageVersionOutput
}

func (*SoftwarePackageVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwarePackageVersion)(nil)).Elem()
}

func (i *SoftwarePackageVersion) ToSoftwarePackageVersionOutput() SoftwarePackageVersionOutput {
	return i.ToSoftwarePackageVersionOutputWithContext(context.Background())
}

func (i *SoftwarePackageVersion) ToSoftwarePackageVersionOutputWithContext(ctx context.Context) SoftwarePackageVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwarePackageVersionOutput)
}

type SoftwarePackageVersionOutput struct{ *pulumi.OutputState }

func (SoftwarePackageVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwarePackageVersion)(nil)).Elem()
}

func (o SoftwarePackageVersionOutput) ToSoftwarePackageVersionOutput() SoftwarePackageVersionOutput {
	return o
}

func (o SoftwarePackageVersionOutput) ToSoftwarePackageVersionOutputWithContext(ctx context.Context) SoftwarePackageVersionOutput {
	return o
}

func (o SoftwarePackageVersionOutput) Attributes() SoftwarePackageVersionResourceAttributesPtrOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) SoftwarePackageVersionResourceAttributesPtrOutput { return v.Attributes }).(SoftwarePackageVersionResourceAttributesPtrOutput)
}

func (o SoftwarePackageVersionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SoftwarePackageVersionOutput) ErrorReason() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringOutput { return v.ErrorReason }).(pulumi.StringOutput)
}

func (o SoftwarePackageVersionOutput) PackageName() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringOutput { return v.PackageName }).(pulumi.StringOutput)
}

func (o SoftwarePackageVersionOutput) PackageVersionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringOutput { return v.PackageVersionArn }).(pulumi.StringOutput)
}

func (o SoftwarePackageVersionOutput) Status() SoftwarePackageVersionPackageVersionStatusOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) SoftwarePackageVersionPackageVersionStatusOutput { return v.Status }).(SoftwarePackageVersionPackageVersionStatusOutput)
}

// An array of key-value pairs to apply to this resource.
func (o SoftwarePackageVersionOutput) Tags() SoftwarePackageVersionTagArrayOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) SoftwarePackageVersionTagArrayOutput { return v.Tags }).(SoftwarePackageVersionTagArrayOutput)
}

func (o SoftwarePackageVersionOutput) VersionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringPtrOutput { return v.VersionName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SoftwarePackageVersionInput)(nil)).Elem(), &SoftwarePackageVersion{})
	pulumi.RegisterOutputType(SoftwarePackageVersionOutput{})
}