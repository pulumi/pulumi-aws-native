// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codeconnections

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for AWS::CodeConnections::Connection resource which can be used to connect external source providers with other AWS services (i.e. AWS CodePipeline)
type Connection struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the  connection. The ARN is used as the connection reference when the connection is shared between AWS services.
	ConnectionArn pulumi.StringOutput `pulumi:"connectionArn"`
	// The name of the connection. Connection names must be unique in an AWS user account.
	ConnectionName pulumi.StringOutput `pulumi:"connectionName"`
	// The current status of the connection.
	ConnectionStatus pulumi.StringOutput `pulumi:"connectionStatus"`
	// The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.
	HostArn pulumi.StringPtrOutput `pulumi:"hostArn"`
	// The name of the external provider where your third-party code repository is configured. For Bitbucket, this is the account ID of the owner of the Bitbucket repository.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
	// The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.
	ProviderType pulumi.StringPtrOutput `pulumi:"providerType"`
	// Specifies the tags applied to a connection.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		args = &ConnectionArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"connectionName",
		"hostArn",
		"providerType",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Connection
	err := ctx.RegisterResource("aws-native:codeconnections:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("aws-native:codeconnections:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
}

type ConnectionState struct {
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// The name of the connection. Connection names must be unique in an AWS user account.
	ConnectionName *string `pulumi:"connectionName"`
	// The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.
	HostArn *string `pulumi:"hostArn"`
	// The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.
	ProviderType *string `pulumi:"providerType"`
	// Specifies the tags applied to a connection.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// The name of the connection. Connection names must be unique in an AWS user account.
	ConnectionName pulumi.StringPtrInput
	// The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.
	HostArn pulumi.StringPtrInput
	// The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.
	ProviderType pulumi.StringPtrInput
	// Specifies the tags applied to a connection.
	Tags aws.TagArrayInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

// The Amazon Resource Name (ARN) of the  connection. The ARN is used as the connection reference when the connection is shared between AWS services.
func (o ConnectionOutput) ConnectionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ConnectionArn }).(pulumi.StringOutput)
}

// The name of the connection. Connection names must be unique in an AWS user account.
func (o ConnectionOutput) ConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ConnectionName }).(pulumi.StringOutput)
}

// The current status of the connection.
func (o ConnectionOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.
func (o ConnectionOutput) HostArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.HostArn }).(pulumi.StringPtrOutput)
}

// The name of the external provider where your third-party code repository is configured. For Bitbucket, this is the account ID of the owner of the Bitbucket repository.
func (o ConnectionOutput) OwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.OwnerAccountId }).(pulumi.StringOutput)
}

// The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.
func (o ConnectionOutput) ProviderType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.ProviderType }).(pulumi.StringPtrOutput)
}

// Specifies the tags applied to a connection.
func (o ConnectionOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Connection) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterOutputType(ConnectionOutput{})
}