// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshiftserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RedshiftServerless::Snapshot Resource Type.
type Snapshot struct {
	pulumi.CustomResourceState

	// The namespace the snapshot is associated with.
	NamespaceName pulumi.StringPtrOutput `pulumi:"namespaceName"`
	// The owner account of the snapshot.
	OwnerAccount pulumi.StringOutput `pulumi:"ownerAccount"`
	// The retention period of the snapshot.
	RetentionPeriod pulumi.IntPtrOutput `pulumi:"retentionPeriod"`
	// Definition for snapshot resource
	Snapshot SnapshotTypeOutput `pulumi:"snapshot"`
	// The name of the snapshot.
	SnapshotName pulumi.StringOutput `pulumi:"snapshotName"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.CreateOnlyTagArrayOutput `pulumi:"tags"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		args = &SnapshotArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"namespaceName",
		"snapshotName",
		"tags[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("aws-native:redshiftserverless:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("aws-native:redshiftserverless:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
}

type SnapshotState struct {
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// The namespace the snapshot is associated with.
	NamespaceName *string `pulumi:"namespaceName"`
	// The retention period of the snapshot.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The name of the snapshot.
	SnapshotName *string `pulumi:"snapshotName"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.CreateOnlyTag `pulumi:"tags"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// The namespace the snapshot is associated with.
	NamespaceName pulumi.StringPtrInput
	// The retention period of the snapshot.
	RetentionPeriod pulumi.IntPtrInput
	// The name of the snapshot.
	SnapshotName pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.CreateOnlyTagArrayInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// The namespace the snapshot is associated with.
func (o SnapshotOutput) NamespaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.NamespaceName }).(pulumi.StringPtrOutput)
}

// The owner account of the snapshot.
func (o SnapshotOutput) OwnerAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.OwnerAccount }).(pulumi.StringOutput)
}

// The retention period of the snapshot.
func (o SnapshotOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntPtrOutput { return v.RetentionPeriod }).(pulumi.IntPtrOutput)
}

// Definition for snapshot resource
func (o SnapshotOutput) Snapshot() SnapshotTypeOutput {
	return o.ApplyT(func(v *Snapshot) SnapshotTypeOutput { return v.Snapshot }).(SnapshotTypeOutput)
}

// The name of the snapshot.
func (o SnapshotOutput) SnapshotName() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SnapshotName }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o SnapshotOutput) Tags() aws.CreateOnlyTagArrayOutput {
	return o.ApplyT(func(v *Snapshot) aws.CreateOnlyTagArrayOutput { return v.Tags }).(aws.CreateOnlyTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterOutputType(SnapshotOutput{})
}
