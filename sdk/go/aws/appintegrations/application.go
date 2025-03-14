// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appintegrations

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS:AppIntegrations::Application
type Application struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the application.
	ApplicationArn pulumi.StringOutput `pulumi:"applicationArn"`
	// Application source config
	ApplicationSourceConfig ApplicationSourceConfigPropertiesOutput `pulumi:"applicationSourceConfig"`
	// The id of the application.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The application description.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the application.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace of the application.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// The configuration of events or requests that the application has access to.
	Permissions pulumi.StringArrayOutput `pulumi:"permissions"`
	// The tags (keys and values) associated with the application.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationSourceConfig == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationSourceConfig'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("aws-native:appintegrations:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("aws-native:appintegrations:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// Application source config
	ApplicationSourceConfig ApplicationSourceConfigProperties `pulumi:"applicationSourceConfig"`
	// The application description.
	Description string `pulumi:"description"`
	// The name of the application.
	Name *string `pulumi:"name"`
	// The namespace of the application.
	Namespace string `pulumi:"namespace"`
	// The configuration of events or requests that the application has access to.
	Permissions []string `pulumi:"permissions"`
	// The tags (keys and values) associated with the application.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// Application source config
	ApplicationSourceConfig ApplicationSourceConfigPropertiesInput
	// The application description.
	Description pulumi.StringInput
	// The name of the application.
	Name pulumi.StringPtrInput
	// The namespace of the application.
	Namespace pulumi.StringInput
	// The configuration of events or requests that the application has access to.
	Permissions pulumi.StringArrayInput
	// The tags (keys and values) associated with the application.
	Tags aws.TagArrayInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

// The Amazon Resource Name (ARN) of the application.
func (o ApplicationOutput) ApplicationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ApplicationArn }).(pulumi.StringOutput)
}

// Application source config
func (o ApplicationOutput) ApplicationSourceConfig() ApplicationSourceConfigPropertiesOutput {
	return o.ApplyT(func(v *Application) ApplicationSourceConfigPropertiesOutput { return v.ApplicationSourceConfig }).(ApplicationSourceConfigPropertiesOutput)
}

// The id of the application.
func (o ApplicationOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The application description.
func (o ApplicationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The name of the application.
func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace of the application.
func (o ApplicationOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// The configuration of events or requests that the application has access to.
func (o ApplicationOutput) Permissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Application) pulumi.StringArrayOutput { return v.Permissions }).(pulumi.StringArrayOutput)
}

// The tags (keys and values) associated with the application.
func (o ApplicationOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Application) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterOutputType(ApplicationOutput{})
}
