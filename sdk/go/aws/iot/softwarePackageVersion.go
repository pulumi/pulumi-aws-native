// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// resource definition
type SoftwarePackageVersion struct {
	pulumi.CustomResourceState

	Artifact SoftwarePackageVersionPackageVersionArtifactPtrOutput `pulumi:"artifact"`
	// Metadata that can be used to define a package version’s configuration. For example, the S3 file location, configuration options that are being sent to the device or fleet.
	//
	// The combined size of all the attributes on a package version is limited to 3KB.
	Attributes pulumi.StringMapOutput `pulumi:"attributes"`
	// A summary of the package version being created. This can be used to outline the package's contents or purpose.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Error reason for a package version failure during creation or update.
	ErrorReason pulumi.StringOutput `pulumi:"errorReason"`
	// The name of the associated software package.
	PackageName pulumi.StringOutput `pulumi:"packageName"`
	// The Amazon Resource Name (ARN) for the package.
	PackageVersionArn pulumi.StringOutput `pulumi:"packageVersionArn"`
	// The inline json job document associated with a software package version
	Recipe               pulumi.StringPtrOutput                           `pulumi:"recipe"`
	Sbom                 SoftwarePackageVersionSbomPtrOutput              `pulumi:"sbom"`
	SbomValidationStatus SoftwarePackageVersionSbomValidationStatusOutput `pulumi:"sbomValidationStatus"`
	// The status of the package version. For more information, see [Package version lifecycle](https://docs.aws.amazon.com/iot/latest/developerguide/preparing-to-use-software-package-catalog.html#package-version-lifecycle) .
	Status SoftwarePackageVersionPackageVersionStatusOutput `pulumi:"status"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The name of the new package version.
	VersionName pulumi.StringPtrOutput `pulumi:"versionName"`
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
	Artifact *SoftwarePackageVersionPackageVersionArtifact `pulumi:"artifact"`
	// Metadata that can be used to define a package version’s configuration. For example, the S3 file location, configuration options that are being sent to the device or fleet.
	//
	// The combined size of all the attributes on a package version is limited to 3KB.
	Attributes map[string]string `pulumi:"attributes"`
	// A summary of the package version being created. This can be used to outline the package's contents or purpose.
	Description *string `pulumi:"description"`
	// The name of the associated software package.
	PackageName string `pulumi:"packageName"`
	// The inline json job document associated with a software package version
	Recipe *string                     `pulumi:"recipe"`
	Sbom   *SoftwarePackageVersionSbom `pulumi:"sbom"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The name of the new package version.
	VersionName *string `pulumi:"versionName"`
}

// The set of arguments for constructing a SoftwarePackageVersion resource.
type SoftwarePackageVersionArgs struct {
	Artifact SoftwarePackageVersionPackageVersionArtifactPtrInput
	// Metadata that can be used to define a package version’s configuration. For example, the S3 file location, configuration options that are being sent to the device or fleet.
	//
	// The combined size of all the attributes on a package version is limited to 3KB.
	Attributes pulumi.StringMapInput
	// A summary of the package version being created. This can be used to outline the package's contents or purpose.
	Description pulumi.StringPtrInput
	// The name of the associated software package.
	PackageName pulumi.StringInput
	// The inline json job document associated with a software package version
	Recipe pulumi.StringPtrInput
	Sbom   SoftwarePackageVersionSbomPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayInput
	// The name of the new package version.
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

func (o SoftwarePackageVersionOutput) Artifact() SoftwarePackageVersionPackageVersionArtifactPtrOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) SoftwarePackageVersionPackageVersionArtifactPtrOutput {
		return v.Artifact
	}).(SoftwarePackageVersionPackageVersionArtifactPtrOutput)
}

// Metadata that can be used to define a package version’s configuration. For example, the S3 file location, configuration options that are being sent to the device or fleet.
//
// The combined size of all the attributes on a package version is limited to 3KB.
func (o SoftwarePackageVersionOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringMapOutput { return v.Attributes }).(pulumi.StringMapOutput)
}

// A summary of the package version being created. This can be used to outline the package's contents or purpose.
func (o SoftwarePackageVersionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Error reason for a package version failure during creation or update.
func (o SoftwarePackageVersionOutput) ErrorReason() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringOutput { return v.ErrorReason }).(pulumi.StringOutput)
}

// The name of the associated software package.
func (o SoftwarePackageVersionOutput) PackageName() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringOutput { return v.PackageName }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) for the package.
func (o SoftwarePackageVersionOutput) PackageVersionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringOutput { return v.PackageVersionArn }).(pulumi.StringOutput)
}

// The inline json job document associated with a software package version
func (o SoftwarePackageVersionOutput) Recipe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringPtrOutput { return v.Recipe }).(pulumi.StringPtrOutput)
}

func (o SoftwarePackageVersionOutput) Sbom() SoftwarePackageVersionSbomPtrOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) SoftwarePackageVersionSbomPtrOutput { return v.Sbom }).(SoftwarePackageVersionSbomPtrOutput)
}

func (o SoftwarePackageVersionOutput) SbomValidationStatus() SoftwarePackageVersionSbomValidationStatusOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) SoftwarePackageVersionSbomValidationStatusOutput {
		return v.SbomValidationStatus
	}).(SoftwarePackageVersionSbomValidationStatusOutput)
}

// The status of the package version. For more information, see [Package version lifecycle](https://docs.aws.amazon.com/iot/latest/developerguide/preparing-to-use-software-package-catalog.html#package-version-lifecycle) .
func (o SoftwarePackageVersionOutput) Status() SoftwarePackageVersionPackageVersionStatusOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) SoftwarePackageVersionPackageVersionStatusOutput { return v.Status }).(SoftwarePackageVersionPackageVersionStatusOutput)
}

// An array of key-value pairs to apply to this resource.
func (o SoftwarePackageVersionOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The name of the new package version.
func (o SoftwarePackageVersionOutput) VersionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoftwarePackageVersion) pulumi.StringPtrOutput { return v.VersionName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SoftwarePackageVersionInput)(nil)).Elem(), &SoftwarePackageVersion{})
	pulumi.RegisterOutputType(SoftwarePackageVersionOutput{})
}
