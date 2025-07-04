// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datazone

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::DataZone::ProjectProfile Resource Type
type ProjectProfile struct {
	pulumi.CustomResourceState

	// The ID of the project profile.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The timestamp of when the project profile was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The user who created the project profile.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// The description of the project profile.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain ID of the project profile.
	DomainId         pulumi.StringOutput    `pulumi:"domainId"`
	DomainIdentifier pulumi.StringPtrOutput `pulumi:"domainIdentifier"`
	// The domain unit ID of the project profile.
	DomainUnitId              pulumi.StringOutput                               `pulumi:"domainUnitId"`
	DomainUnitIdentifier      pulumi.StringPtrOutput                            `pulumi:"domainUnitIdentifier"`
	EnvironmentConfigurations ProjectProfileEnvironmentConfigurationArrayOutput `pulumi:"environmentConfigurations"`
	Identifier                pulumi.StringOutput                               `pulumi:"identifier"`
	// The timestamp at which a project profile was last updated.
	LastUpdatedAt pulumi.StringOutput `pulumi:"lastUpdatedAt"`
	// The name of a project profile.
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of a project profile.
	Status ProjectProfileStatusPtrOutput `pulumi:"status"`
}

// NewProjectProfile registers a new resource with the given unique name, arguments, and options.
func NewProjectProfile(ctx *pulumi.Context,
	name string, args *ProjectProfileArgs, opts ...pulumi.ResourceOption) (*ProjectProfile, error) {
	if args == nil {
		args = &ProjectProfileArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"domainIdentifier",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectProfile
	err := ctx.RegisterResource("aws-native:datazone:ProjectProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectProfile gets an existing ProjectProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectProfileState, opts ...pulumi.ResourceOption) (*ProjectProfile, error) {
	var resource ProjectProfile
	err := ctx.ReadResource("aws-native:datazone:ProjectProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectProfile resources.
type projectProfileState struct {
}

type ProjectProfileState struct {
}

func (ProjectProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectProfileState)(nil)).Elem()
}

type projectProfileArgs struct {
	// The description of the project profile.
	Description               *string                                  `pulumi:"description"`
	DomainIdentifier          *string                                  `pulumi:"domainIdentifier"`
	DomainUnitIdentifier      *string                                  `pulumi:"domainUnitIdentifier"`
	EnvironmentConfigurations []ProjectProfileEnvironmentConfiguration `pulumi:"environmentConfigurations"`
	// The name of a project profile.
	Name *string `pulumi:"name"`
	// The status of a project profile.
	Status *ProjectProfileStatus `pulumi:"status"`
}

// The set of arguments for constructing a ProjectProfile resource.
type ProjectProfileArgs struct {
	// The description of the project profile.
	Description               pulumi.StringPtrInput
	DomainIdentifier          pulumi.StringPtrInput
	DomainUnitIdentifier      pulumi.StringPtrInput
	EnvironmentConfigurations ProjectProfileEnvironmentConfigurationArrayInput
	// The name of a project profile.
	Name pulumi.StringPtrInput
	// The status of a project profile.
	Status ProjectProfileStatusPtrInput
}

func (ProjectProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectProfileArgs)(nil)).Elem()
}

type ProjectProfileInput interface {
	pulumi.Input

	ToProjectProfileOutput() ProjectProfileOutput
	ToProjectProfileOutputWithContext(ctx context.Context) ProjectProfileOutput
}

func (*ProjectProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectProfile)(nil)).Elem()
}

func (i *ProjectProfile) ToProjectProfileOutput() ProjectProfileOutput {
	return i.ToProjectProfileOutputWithContext(context.Background())
}

func (i *ProjectProfile) ToProjectProfileOutputWithContext(ctx context.Context) ProjectProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectProfileOutput)
}

type ProjectProfileOutput struct{ *pulumi.OutputState }

func (ProjectProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectProfile)(nil)).Elem()
}

func (o ProjectProfileOutput) ToProjectProfileOutput() ProjectProfileOutput {
	return o
}

func (o ProjectProfileOutput) ToProjectProfileOutputWithContext(ctx context.Context) ProjectProfileOutput {
	return o
}

// The ID of the project profile.
func (o ProjectProfileOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The timestamp of when the project profile was created.
func (o ProjectProfileOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The user who created the project profile.
func (o ProjectProfileOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// The description of the project profile.
func (o ProjectProfileOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The domain ID of the project profile.
func (o ProjectProfileOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

func (o ProjectProfileOutput) DomainIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringPtrOutput { return v.DomainIdentifier }).(pulumi.StringPtrOutput)
}

// The domain unit ID of the project profile.
func (o ProjectProfileOutput) DomainUnitId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringOutput { return v.DomainUnitId }).(pulumi.StringOutput)
}

func (o ProjectProfileOutput) DomainUnitIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringPtrOutput { return v.DomainUnitIdentifier }).(pulumi.StringPtrOutput)
}

func (o ProjectProfileOutput) EnvironmentConfigurations() ProjectProfileEnvironmentConfigurationArrayOutput {
	return o.ApplyT(func(v *ProjectProfile) ProjectProfileEnvironmentConfigurationArrayOutput {
		return v.EnvironmentConfigurations
	}).(ProjectProfileEnvironmentConfigurationArrayOutput)
}

func (o ProjectProfileOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// The timestamp at which a project profile was last updated.
func (o ProjectProfileOutput) LastUpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringOutput { return v.LastUpdatedAt }).(pulumi.StringOutput)
}

// The name of a project profile.
func (o ProjectProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of a project profile.
func (o ProjectProfileOutput) Status() ProjectProfileStatusPtrOutput {
	return o.ApplyT(func(v *ProjectProfile) ProjectProfileStatusPtrOutput { return v.Status }).(ProjectProfileStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectProfileInput)(nil)).Elem(), &ProjectProfile{})
	pulumi.RegisterOutputType(ProjectProfileOutput{})
}
