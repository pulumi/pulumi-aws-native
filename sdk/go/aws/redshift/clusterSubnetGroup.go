// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Redshift::ClusterSubnetGroup. Specifies an Amazon Redshift subnet group.
type ClusterSubnetGroup struct {
	pulumi.CustomResourceState

	// This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be "Default".
	ClusterSubnetGroupName pulumi.StringOutput `pulumi:"clusterSubnetGroupName"`
	// The description of the parameter group.
	Description pulumi.StringOutput `pulumi:"description"`
	// The list of VPC subnet IDs
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// The list of tags for the cluster parameter group.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewClusterSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterSubnetGroup(ctx *pulumi.Context,
	name string, args *ClusterSubnetGroupArgs, opts ...pulumi.ResourceOption) (*ClusterSubnetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterSubnetGroup
	err := ctx.RegisterResource("aws-native:redshift:ClusterSubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterSubnetGroup gets an existing ClusterSubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterSubnetGroupState, opts ...pulumi.ResourceOption) (*ClusterSubnetGroup, error) {
	var resource ClusterSubnetGroup
	err := ctx.ReadResource("aws-native:redshift:ClusterSubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterSubnetGroup resources.
type clusterSubnetGroupState struct {
}

type ClusterSubnetGroupState struct {
}

func (ClusterSubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterSubnetGroupState)(nil)).Elem()
}

type clusterSubnetGroupArgs struct {
	// The description of the parameter group.
	Description string `pulumi:"description"`
	// The list of VPC subnet IDs
	SubnetIds []string `pulumi:"subnetIds"`
	// The list of tags for the cluster parameter group.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterSubnetGroup resource.
type ClusterSubnetGroupArgs struct {
	// The description of the parameter group.
	Description pulumi.StringInput
	// The list of VPC subnet IDs
	SubnetIds pulumi.StringArrayInput
	// The list of tags for the cluster parameter group.
	Tags aws.TagArrayInput
}

func (ClusterSubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterSubnetGroupArgs)(nil)).Elem()
}

type ClusterSubnetGroupInput interface {
	pulumi.Input

	ToClusterSubnetGroupOutput() ClusterSubnetGroupOutput
	ToClusterSubnetGroupOutputWithContext(ctx context.Context) ClusterSubnetGroupOutput
}

func (*ClusterSubnetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSubnetGroup)(nil)).Elem()
}

func (i *ClusterSubnetGroup) ToClusterSubnetGroupOutput() ClusterSubnetGroupOutput {
	return i.ToClusterSubnetGroupOutputWithContext(context.Background())
}

func (i *ClusterSubnetGroup) ToClusterSubnetGroupOutputWithContext(ctx context.Context) ClusterSubnetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSubnetGroupOutput)
}

type ClusterSubnetGroupOutput struct{ *pulumi.OutputState }

func (ClusterSubnetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSubnetGroup)(nil)).Elem()
}

func (o ClusterSubnetGroupOutput) ToClusterSubnetGroupOutput() ClusterSubnetGroupOutput {
	return o
}

func (o ClusterSubnetGroupOutput) ToClusterSubnetGroupOutputWithContext(ctx context.Context) ClusterSubnetGroupOutput {
	return o
}

// This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be "Default".
func (o ClusterSubnetGroupOutput) ClusterSubnetGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSubnetGroup) pulumi.StringOutput { return v.ClusterSubnetGroupName }).(pulumi.StringOutput)
}

// The description of the parameter group.
func (o ClusterSubnetGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSubnetGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The list of VPC subnet IDs
func (o ClusterSubnetGroupOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterSubnetGroup) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// The list of tags for the cluster parameter group.
func (o ClusterSubnetGroupOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *ClusterSubnetGroup) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSubnetGroupInput)(nil)).Elem(), &ClusterSubnetGroup{})
	pulumi.RegisterOutputType(ClusterSubnetGroupOutput{})
}
