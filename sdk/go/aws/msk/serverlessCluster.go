// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::MSK::ServerlessCluster
type ServerlessCluster struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput                         `pulumi:"arn"`
	ClientAuthentication ServerlessClusterClientAuthenticationOutput `pulumi:"clientAuthentication"`
	ClusterName          pulumi.StringOutput                         `pulumi:"clusterName"`
	// A key-value pair to associate with a resource.
	Tags       pulumi.AnyOutput                      `pulumi:"tags"`
	VpcConfigs ServerlessClusterVpcConfigArrayOutput `pulumi:"vpcConfigs"`
}

// NewServerlessCluster registers a new resource with the given unique name, arguments, and options.
func NewServerlessCluster(ctx *pulumi.Context,
	name string, args *ServerlessClusterArgs, opts ...pulumi.ResourceOption) (*ServerlessCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientAuthentication == nil {
		return nil, errors.New("invalid value for required argument 'ClientAuthentication'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.VpcConfigs == nil {
		return nil, errors.New("invalid value for required argument 'VpcConfigs'")
	}
	var resource ServerlessCluster
	err := ctx.RegisterResource("aws-native:msk:ServerlessCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerlessCluster gets an existing ServerlessCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerlessCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerlessClusterState, opts ...pulumi.ResourceOption) (*ServerlessCluster, error) {
	var resource ServerlessCluster
	err := ctx.ReadResource("aws-native:msk:ServerlessCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerlessCluster resources.
type serverlessClusterState struct {
}

type ServerlessClusterState struct {
}

func (ServerlessClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessClusterState)(nil)).Elem()
}

type serverlessClusterArgs struct {
	ClientAuthentication ServerlessClusterClientAuthentication `pulumi:"clientAuthentication"`
	ClusterName          string                                `pulumi:"clusterName"`
	// A key-value pair to associate with a resource.
	Tags       interface{}                  `pulumi:"tags"`
	VpcConfigs []ServerlessClusterVpcConfig `pulumi:"vpcConfigs"`
}

// The set of arguments for constructing a ServerlessCluster resource.
type ServerlessClusterArgs struct {
	ClientAuthentication ServerlessClusterClientAuthenticationInput
	ClusterName          pulumi.StringInput
	// A key-value pair to associate with a resource.
	Tags       pulumi.Input
	VpcConfigs ServerlessClusterVpcConfigArrayInput
}

func (ServerlessClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessClusterArgs)(nil)).Elem()
}

type ServerlessClusterInput interface {
	pulumi.Input

	ToServerlessClusterOutput() ServerlessClusterOutput
	ToServerlessClusterOutputWithContext(ctx context.Context) ServerlessClusterOutput
}

func (*ServerlessCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessCluster)(nil)).Elem()
}

func (i *ServerlessCluster) ToServerlessClusterOutput() ServerlessClusterOutput {
	return i.ToServerlessClusterOutputWithContext(context.Background())
}

func (i *ServerlessCluster) ToServerlessClusterOutputWithContext(ctx context.Context) ServerlessClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessClusterOutput)
}

type ServerlessClusterOutput struct{ *pulumi.OutputState }

func (ServerlessClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessCluster)(nil)).Elem()
}

func (o ServerlessClusterOutput) ToServerlessClusterOutput() ServerlessClusterOutput {
	return o
}

func (o ServerlessClusterOutput) ToServerlessClusterOutputWithContext(ctx context.Context) ServerlessClusterOutput {
	return o
}

func (o ServerlessClusterOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessCluster) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ServerlessClusterOutput) ClientAuthentication() ServerlessClusterClientAuthenticationOutput {
	return o.ApplyT(func(v *ServerlessCluster) ServerlessClusterClientAuthenticationOutput { return v.ClientAuthentication }).(ServerlessClusterClientAuthenticationOutput)
}

func (o ServerlessClusterOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessCluster) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// A key-value pair to associate with a resource.
func (o ServerlessClusterOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *ServerlessCluster) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o ServerlessClusterOutput) VpcConfigs() ServerlessClusterVpcConfigArrayOutput {
	return o.ApplyT(func(v *ServerlessCluster) ServerlessClusterVpcConfigArrayOutput { return v.VpcConfigs }).(ServerlessClusterVpcConfigArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessClusterInput)(nil)).Elem(), &ServerlessCluster{})
	pulumi.RegisterOutputType(ServerlessClusterOutput{})
}