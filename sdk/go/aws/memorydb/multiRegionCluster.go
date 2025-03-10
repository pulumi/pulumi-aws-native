// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package memorydb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::MemoryDB::Multi Region Cluster resource creates an Amazon MemoryDB Multi Region Cluster.
type MultiRegionCluster struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the multi region cluster.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the multi region cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The engine type used by the multi region cluster.
	Engine pulumi.StringPtrOutput `pulumi:"engine"`
	// The Redis engine version used by the multi region cluster.
	EngineVersion pulumi.StringPtrOutput `pulumi:"engineVersion"`
	// The name of the Global Datastore, it is generated by MemoryDB adding a prefix to MultiRegionClusterNameSuffix.
	MultiRegionClusterName pulumi.StringOutput `pulumi:"multiRegionClusterName"`
	// The name of the Multi Region cluster. This value must be unique as it also serves as the multi region cluster identifier.
	MultiRegionClusterNameSuffix pulumi.StringPtrOutput `pulumi:"multiRegionClusterNameSuffix"`
	// The name of the parameter group associated with the multi region cluster.
	MultiRegionParameterGroupName pulumi.StringPtrOutput `pulumi:"multiRegionParameterGroupName"`
	// The compute and memory capacity of the nodes in the multi region cluster.
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// The number of shards the multi region cluster will contain.
	NumShards pulumi.IntPtrOutput `pulumi:"numShards"`
	// The status of the multi region cluster. For example, Available, Updating, Creating.
	Status pulumi.StringOutput `pulumi:"status"`
	// An array of key-value pairs to apply to this multi region cluster.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// A flag that enables in-transit encryption when set to true.
	//
	// You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
	TlsEnabled pulumi.BoolPtrOutput `pulumi:"tlsEnabled"`
	// An enum string value that determines the update strategy for scaling. Possible values are 'COORDINATED' and 'UNCOORDINATED'. Default is 'COORDINATED'.
	UpdateStrategy MultiRegionClusterUpdateStrategyPtrOutput `pulumi:"updateStrategy"`
}

// NewMultiRegionCluster registers a new resource with the given unique name, arguments, and options.
func NewMultiRegionCluster(ctx *pulumi.Context,
	name string, args *MultiRegionClusterArgs, opts ...pulumi.ResourceOption) (*MultiRegionCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"engineVersion",
		"multiRegionClusterNameSuffix",
		"multiRegionParameterGroupName",
		"tlsEnabled",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MultiRegionCluster
	err := ctx.RegisterResource("aws-native:memorydb:MultiRegionCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMultiRegionCluster gets an existing MultiRegionCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMultiRegionCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MultiRegionClusterState, opts ...pulumi.ResourceOption) (*MultiRegionCluster, error) {
	var resource MultiRegionCluster
	err := ctx.ReadResource("aws-native:memorydb:MultiRegionCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MultiRegionCluster resources.
type multiRegionClusterState struct {
}

type MultiRegionClusterState struct {
}

func (MultiRegionClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*multiRegionClusterState)(nil)).Elem()
}

type multiRegionClusterArgs struct {
	// Description of the multi region cluster.
	Description *string `pulumi:"description"`
	// The engine type used by the multi region cluster.
	Engine *string `pulumi:"engine"`
	// The Redis engine version used by the multi region cluster.
	EngineVersion *string `pulumi:"engineVersion"`
	// The name of the Multi Region cluster. This value must be unique as it also serves as the multi region cluster identifier.
	MultiRegionClusterNameSuffix *string `pulumi:"multiRegionClusterNameSuffix"`
	// The name of the parameter group associated with the multi region cluster.
	MultiRegionParameterGroupName *string `pulumi:"multiRegionParameterGroupName"`
	// The compute and memory capacity of the nodes in the multi region cluster.
	NodeType string `pulumi:"nodeType"`
	// The number of shards the multi region cluster will contain.
	NumShards *int `pulumi:"numShards"`
	// An array of key-value pairs to apply to this multi region cluster.
	Tags []aws.Tag `pulumi:"tags"`
	// A flag that enables in-transit encryption when set to true.
	//
	// You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
	TlsEnabled *bool `pulumi:"tlsEnabled"`
	// An enum string value that determines the update strategy for scaling. Possible values are 'COORDINATED' and 'UNCOORDINATED'. Default is 'COORDINATED'.
	UpdateStrategy *MultiRegionClusterUpdateStrategy `pulumi:"updateStrategy"`
}

// The set of arguments for constructing a MultiRegionCluster resource.
type MultiRegionClusterArgs struct {
	// Description of the multi region cluster.
	Description pulumi.StringPtrInput
	// The engine type used by the multi region cluster.
	Engine pulumi.StringPtrInput
	// The Redis engine version used by the multi region cluster.
	EngineVersion pulumi.StringPtrInput
	// The name of the Multi Region cluster. This value must be unique as it also serves as the multi region cluster identifier.
	MultiRegionClusterNameSuffix pulumi.StringPtrInput
	// The name of the parameter group associated with the multi region cluster.
	MultiRegionParameterGroupName pulumi.StringPtrInput
	// The compute and memory capacity of the nodes in the multi region cluster.
	NodeType pulumi.StringInput
	// The number of shards the multi region cluster will contain.
	NumShards pulumi.IntPtrInput
	// An array of key-value pairs to apply to this multi region cluster.
	Tags aws.TagArrayInput
	// A flag that enables in-transit encryption when set to true.
	//
	// You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
	TlsEnabled pulumi.BoolPtrInput
	// An enum string value that determines the update strategy for scaling. Possible values are 'COORDINATED' and 'UNCOORDINATED'. Default is 'COORDINATED'.
	UpdateStrategy MultiRegionClusterUpdateStrategyPtrInput
}

func (MultiRegionClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*multiRegionClusterArgs)(nil)).Elem()
}

type MultiRegionClusterInput interface {
	pulumi.Input

	ToMultiRegionClusterOutput() MultiRegionClusterOutput
	ToMultiRegionClusterOutputWithContext(ctx context.Context) MultiRegionClusterOutput
}

func (*MultiRegionCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**MultiRegionCluster)(nil)).Elem()
}

func (i *MultiRegionCluster) ToMultiRegionClusterOutput() MultiRegionClusterOutput {
	return i.ToMultiRegionClusterOutputWithContext(context.Background())
}

func (i *MultiRegionCluster) ToMultiRegionClusterOutputWithContext(ctx context.Context) MultiRegionClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiRegionClusterOutput)
}

type MultiRegionClusterOutput struct{ *pulumi.OutputState }

func (MultiRegionClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MultiRegionCluster)(nil)).Elem()
}

func (o MultiRegionClusterOutput) ToMultiRegionClusterOutput() MultiRegionClusterOutput {
	return o
}

func (o MultiRegionClusterOutput) ToMultiRegionClusterOutputWithContext(ctx context.Context) MultiRegionClusterOutput {
	return o
}

// The Amazon Resource Name (ARN) of the multi region cluster.
func (o MultiRegionClusterOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of the multi region cluster.
func (o MultiRegionClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The engine type used by the multi region cluster.
func (o MultiRegionClusterOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.StringPtrOutput { return v.Engine }).(pulumi.StringPtrOutput)
}

// The Redis engine version used by the multi region cluster.
func (o MultiRegionClusterOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.StringPtrOutput { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

// The name of the Global Datastore, it is generated by MemoryDB adding a prefix to MultiRegionClusterNameSuffix.
func (o MultiRegionClusterOutput) MultiRegionClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.StringOutput { return v.MultiRegionClusterName }).(pulumi.StringOutput)
}

// The name of the Multi Region cluster. This value must be unique as it also serves as the multi region cluster identifier.
func (o MultiRegionClusterOutput) MultiRegionClusterNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.StringPtrOutput { return v.MultiRegionClusterNameSuffix }).(pulumi.StringPtrOutput)
}

// The name of the parameter group associated with the multi region cluster.
func (o MultiRegionClusterOutput) MultiRegionParameterGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.StringPtrOutput { return v.MultiRegionParameterGroupName }).(pulumi.StringPtrOutput)
}

// The compute and memory capacity of the nodes in the multi region cluster.
func (o MultiRegionClusterOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// The number of shards the multi region cluster will contain.
func (o MultiRegionClusterOutput) NumShards() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.IntPtrOutput { return v.NumShards }).(pulumi.IntPtrOutput)
}

// The status of the multi region cluster. For example, Available, Updating, Creating.
func (o MultiRegionClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this multi region cluster.
func (o MultiRegionClusterOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *MultiRegionCluster) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// A flag that enables in-transit encryption when set to true.
//
// You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
func (o MultiRegionClusterOutput) TlsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MultiRegionCluster) pulumi.BoolPtrOutput { return v.TlsEnabled }).(pulumi.BoolPtrOutput)
}

// An enum string value that determines the update strategy for scaling. Possible values are 'COORDINATED' and 'UNCOORDINATED'. Default is 'COORDINATED'.
func (o MultiRegionClusterOutput) UpdateStrategy() MultiRegionClusterUpdateStrategyPtrOutput {
	return o.ApplyT(func(v *MultiRegionCluster) MultiRegionClusterUpdateStrategyPtrOutput { return v.UpdateStrategy }).(MultiRegionClusterUpdateStrategyPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MultiRegionClusterInput)(nil)).Elem(), &MultiRegionCluster{})
	pulumi.RegisterOutputType(MultiRegionClusterOutput{})
}
