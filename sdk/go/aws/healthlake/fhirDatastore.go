// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthlake

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// HealthLake FHIR Datastore
type FhirDatastore struct {
	pulumi.CustomResourceState

	CreatedAt                     FhirDatastoreCreatedAtOutput                        `pulumi:"createdAt"`
	DatastoreArn                  pulumi.StringOutput                                 `pulumi:"datastoreArn"`
	DatastoreEndpoint             pulumi.StringOutput                                 `pulumi:"datastoreEndpoint"`
	DatastoreId                   pulumi.StringOutput                                 `pulumi:"datastoreId"`
	DatastoreName                 pulumi.StringPtrOutput                              `pulumi:"datastoreName"`
	DatastoreStatus               FhirDatastoreDatastoreStatusOutput                  `pulumi:"datastoreStatus"`
	DatastoreTypeVersion          FhirDatastoreDatastoreTypeVersionOutput             `pulumi:"datastoreTypeVersion"`
	IdentityProviderConfiguration FhirDatastoreIdentityProviderConfigurationPtrOutput `pulumi:"identityProviderConfiguration"`
	PreloadDataConfig             FhirDatastorePreloadDataConfigPtrOutput             `pulumi:"preloadDataConfig"`
	SseConfiguration              FhirDatastoreSseConfigurationPtrOutput              `pulumi:"sseConfiguration"`
	Tags                          FhirDatastoreTagArrayOutput                         `pulumi:"tags"`
}

// NewFhirDatastore registers a new resource with the given unique name, arguments, and options.
func NewFhirDatastore(ctx *pulumi.Context,
	name string, args *FhirDatastoreArgs, opts ...pulumi.ResourceOption) (*FhirDatastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatastoreTypeVersion == nil {
		return nil, errors.New("invalid value for required argument 'DatastoreTypeVersion'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FhirDatastore
	err := ctx.RegisterResource("aws-native:healthlake:FhirDatastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFhirDatastore gets an existing FhirDatastore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FhirDatastoreState, opts ...pulumi.ResourceOption) (*FhirDatastore, error) {
	var resource FhirDatastore
	err := ctx.ReadResource("aws-native:healthlake:FhirDatastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FhirDatastore resources.
type fhirDatastoreState struct {
}

type FhirDatastoreState struct {
}

func (FhirDatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirDatastoreState)(nil)).Elem()
}

type fhirDatastoreArgs struct {
	DatastoreName                 *string                                     `pulumi:"datastoreName"`
	DatastoreTypeVersion          FhirDatastoreDatastoreTypeVersion           `pulumi:"datastoreTypeVersion"`
	IdentityProviderConfiguration *FhirDatastoreIdentityProviderConfiguration `pulumi:"identityProviderConfiguration"`
	PreloadDataConfig             *FhirDatastorePreloadDataConfig             `pulumi:"preloadDataConfig"`
	SseConfiguration              *FhirDatastoreSseConfiguration              `pulumi:"sseConfiguration"`
	Tags                          []FhirDatastoreTag                          `pulumi:"tags"`
}

// The set of arguments for constructing a FhirDatastore resource.
type FhirDatastoreArgs struct {
	DatastoreName                 pulumi.StringPtrInput
	DatastoreTypeVersion          FhirDatastoreDatastoreTypeVersionInput
	IdentityProviderConfiguration FhirDatastoreIdentityProviderConfigurationPtrInput
	PreloadDataConfig             FhirDatastorePreloadDataConfigPtrInput
	SseConfiguration              FhirDatastoreSseConfigurationPtrInput
	Tags                          FhirDatastoreTagArrayInput
}

func (FhirDatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirDatastoreArgs)(nil)).Elem()
}

type FhirDatastoreInput interface {
	pulumi.Input

	ToFhirDatastoreOutput() FhirDatastoreOutput
	ToFhirDatastoreOutputWithContext(ctx context.Context) FhirDatastoreOutput
}

func (*FhirDatastore) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirDatastore)(nil)).Elem()
}

func (i *FhirDatastore) ToFhirDatastoreOutput() FhirDatastoreOutput {
	return i.ToFhirDatastoreOutputWithContext(context.Background())
}

func (i *FhirDatastore) ToFhirDatastoreOutputWithContext(ctx context.Context) FhirDatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirDatastoreOutput)
}

type FhirDatastoreOutput struct{ *pulumi.OutputState }

func (FhirDatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirDatastore)(nil)).Elem()
}

func (o FhirDatastoreOutput) ToFhirDatastoreOutput() FhirDatastoreOutput {
	return o
}

func (o FhirDatastoreOutput) ToFhirDatastoreOutputWithContext(ctx context.Context) FhirDatastoreOutput {
	return o
}

func (o FhirDatastoreOutput) CreatedAt() FhirDatastoreCreatedAtOutput {
	return o.ApplyT(func(v *FhirDatastore) FhirDatastoreCreatedAtOutput { return v.CreatedAt }).(FhirDatastoreCreatedAtOutput)
}

func (o FhirDatastoreOutput) DatastoreArn() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirDatastore) pulumi.StringOutput { return v.DatastoreArn }).(pulumi.StringOutput)
}

func (o FhirDatastoreOutput) DatastoreEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirDatastore) pulumi.StringOutput { return v.DatastoreEndpoint }).(pulumi.StringOutput)
}

func (o FhirDatastoreOutput) DatastoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirDatastore) pulumi.StringOutput { return v.DatastoreId }).(pulumi.StringOutput)
}

func (o FhirDatastoreOutput) DatastoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirDatastore) pulumi.StringPtrOutput { return v.DatastoreName }).(pulumi.StringPtrOutput)
}

func (o FhirDatastoreOutput) DatastoreStatus() FhirDatastoreDatastoreStatusOutput {
	return o.ApplyT(func(v *FhirDatastore) FhirDatastoreDatastoreStatusOutput { return v.DatastoreStatus }).(FhirDatastoreDatastoreStatusOutput)
}

func (o FhirDatastoreOutput) DatastoreTypeVersion() FhirDatastoreDatastoreTypeVersionOutput {
	return o.ApplyT(func(v *FhirDatastore) FhirDatastoreDatastoreTypeVersionOutput { return v.DatastoreTypeVersion }).(FhirDatastoreDatastoreTypeVersionOutput)
}

func (o FhirDatastoreOutput) IdentityProviderConfiguration() FhirDatastoreIdentityProviderConfigurationPtrOutput {
	return o.ApplyT(func(v *FhirDatastore) FhirDatastoreIdentityProviderConfigurationPtrOutput {
		return v.IdentityProviderConfiguration
	}).(FhirDatastoreIdentityProviderConfigurationPtrOutput)
}

func (o FhirDatastoreOutput) PreloadDataConfig() FhirDatastorePreloadDataConfigPtrOutput {
	return o.ApplyT(func(v *FhirDatastore) FhirDatastorePreloadDataConfigPtrOutput { return v.PreloadDataConfig }).(FhirDatastorePreloadDataConfigPtrOutput)
}

func (o FhirDatastoreOutput) SseConfiguration() FhirDatastoreSseConfigurationPtrOutput {
	return o.ApplyT(func(v *FhirDatastore) FhirDatastoreSseConfigurationPtrOutput { return v.SseConfiguration }).(FhirDatastoreSseConfigurationPtrOutput)
}

func (o FhirDatastoreOutput) Tags() FhirDatastoreTagArrayOutput {
	return o.ApplyT(func(v *FhirDatastore) FhirDatastoreTagArrayOutput { return v.Tags }).(FhirDatastoreTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FhirDatastoreInput)(nil)).Elem(), &FhirDatastore{})
	pulumi.RegisterOutputType(FhirDatastoreOutput{})
}