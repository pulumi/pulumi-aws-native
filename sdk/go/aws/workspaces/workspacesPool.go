// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspaces

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WorkSpaces::WorkspacesPool
type WorkspacesPool struct {
	pulumi.CustomResourceState

	// The persistent application settings for users of the pool.
	ApplicationSettings WorkspacesPoolApplicationSettingsPtrOutput `pulumi:"applicationSettings"`
	// The identifier of the bundle used by the pool.
	BundleId pulumi.StringOutput `pulumi:"bundleId"`
	// Describes the user capacity for the pool.
	Capacity WorkspacesPoolCapacityOutput `pulumi:"capacity"`
	// The time the pool was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the pool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The identifier of the directory used by the pool.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// The Amazon Resource Name (ARN) for the pool.
	PoolArn pulumi.StringOutput `pulumi:"poolArn"`
	// The identifier of the pool.
	PoolId pulumi.StringOutput `pulumi:"poolId"`
	// The name of the pool.
	PoolName pulumi.StringOutput `pulumi:"poolName"`
	// The running mode of the pool.
	RunningMode WorkspacesPoolRunningModePtrOutput `pulumi:"runningMode"`
	Tags        aws.TagArrayOutput                 `pulumi:"tags"`
	// The amount of time that a pool session remains active after users disconnect. If they try to reconnect to the pool session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new pool instance.
	TimeoutSettings WorkspacesPoolTimeoutSettingsPtrOutput `pulumi:"timeoutSettings"`
}

// NewWorkspacesPool registers a new resource with the given unique name, arguments, and options.
func NewWorkspacesPool(ctx *pulumi.Context,
	name string, args *WorkspacesPoolArgs, opts ...pulumi.ResourceOption) (*WorkspacesPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BundleId == nil {
		return nil, errors.New("invalid value for required argument 'BundleId'")
	}
	if args.Capacity == nil {
		return nil, errors.New("invalid value for required argument 'Capacity'")
	}
	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"poolName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WorkspacesPool
	err := ctx.RegisterResource("aws-native:workspaces:WorkspacesPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspacesPool gets an existing WorkspacesPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspacesPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspacesPoolState, opts ...pulumi.ResourceOption) (*WorkspacesPool, error) {
	var resource WorkspacesPool
	err := ctx.ReadResource("aws-native:workspaces:WorkspacesPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspacesPool resources.
type workspacesPoolState struct {
}

type WorkspacesPoolState struct {
}

func (WorkspacesPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspacesPoolState)(nil)).Elem()
}

type workspacesPoolArgs struct {
	// The persistent application settings for users of the pool.
	ApplicationSettings *WorkspacesPoolApplicationSettings `pulumi:"applicationSettings"`
	// The identifier of the bundle used by the pool.
	BundleId string `pulumi:"bundleId"`
	// Describes the user capacity for the pool.
	Capacity WorkspacesPoolCapacity `pulumi:"capacity"`
	// The description of the pool.
	Description *string `pulumi:"description"`
	// The identifier of the directory used by the pool.
	DirectoryId string `pulumi:"directoryId"`
	// The name of the pool.
	PoolName *string `pulumi:"poolName"`
	// The running mode of the pool.
	RunningMode *WorkspacesPoolRunningMode `pulumi:"runningMode"`
	Tags        []aws.Tag                  `pulumi:"tags"`
	// The amount of time that a pool session remains active after users disconnect. If they try to reconnect to the pool session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new pool instance.
	TimeoutSettings *WorkspacesPoolTimeoutSettings `pulumi:"timeoutSettings"`
}

// The set of arguments for constructing a WorkspacesPool resource.
type WorkspacesPoolArgs struct {
	// The persistent application settings for users of the pool.
	ApplicationSettings WorkspacesPoolApplicationSettingsPtrInput
	// The identifier of the bundle used by the pool.
	BundleId pulumi.StringInput
	// Describes the user capacity for the pool.
	Capacity WorkspacesPoolCapacityInput
	// The description of the pool.
	Description pulumi.StringPtrInput
	// The identifier of the directory used by the pool.
	DirectoryId pulumi.StringInput
	// The name of the pool.
	PoolName pulumi.StringPtrInput
	// The running mode of the pool.
	RunningMode WorkspacesPoolRunningModePtrInput
	Tags        aws.TagArrayInput
	// The amount of time that a pool session remains active after users disconnect. If they try to reconnect to the pool session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new pool instance.
	TimeoutSettings WorkspacesPoolTimeoutSettingsPtrInput
}

func (WorkspacesPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspacesPoolArgs)(nil)).Elem()
}

type WorkspacesPoolInput interface {
	pulumi.Input

	ToWorkspacesPoolOutput() WorkspacesPoolOutput
	ToWorkspacesPoolOutputWithContext(ctx context.Context) WorkspacesPoolOutput
}

func (*WorkspacesPool) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacesPool)(nil)).Elem()
}

func (i *WorkspacesPool) ToWorkspacesPoolOutput() WorkspacesPoolOutput {
	return i.ToWorkspacesPoolOutputWithContext(context.Background())
}

func (i *WorkspacesPool) ToWorkspacesPoolOutputWithContext(ctx context.Context) WorkspacesPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacesPoolOutput)
}

type WorkspacesPoolOutput struct{ *pulumi.OutputState }

func (WorkspacesPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacesPool)(nil)).Elem()
}

func (o WorkspacesPoolOutput) ToWorkspacesPoolOutput() WorkspacesPoolOutput {
	return o
}

func (o WorkspacesPoolOutput) ToWorkspacesPoolOutputWithContext(ctx context.Context) WorkspacesPoolOutput {
	return o
}

// The persistent application settings for users of the pool.
func (o WorkspacesPoolOutput) ApplicationSettings() WorkspacesPoolApplicationSettingsPtrOutput {
	return o.ApplyT(func(v *WorkspacesPool) WorkspacesPoolApplicationSettingsPtrOutput { return v.ApplicationSettings }).(WorkspacesPoolApplicationSettingsPtrOutput)
}

// The identifier of the bundle used by the pool.
func (o WorkspacesPoolOutput) BundleId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacesPool) pulumi.StringOutput { return v.BundleId }).(pulumi.StringOutput)
}

// Describes the user capacity for the pool.
func (o WorkspacesPoolOutput) Capacity() WorkspacesPoolCapacityOutput {
	return o.ApplyT(func(v *WorkspacesPool) WorkspacesPoolCapacityOutput { return v.Capacity }).(WorkspacesPoolCapacityOutput)
}

// The time the pool was created.
func (o WorkspacesPoolOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacesPool) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the pool.
func (o WorkspacesPoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspacesPool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the directory used by the pool.
func (o WorkspacesPoolOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacesPool) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) for the pool.
func (o WorkspacesPoolOutput) PoolArn() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacesPool) pulumi.StringOutput { return v.PoolArn }).(pulumi.StringOutput)
}

// The identifier of the pool.
func (o WorkspacesPoolOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacesPool) pulumi.StringOutput { return v.PoolId }).(pulumi.StringOutput)
}

// The name of the pool.
func (o WorkspacesPoolOutput) PoolName() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacesPool) pulumi.StringOutput { return v.PoolName }).(pulumi.StringOutput)
}

// The running mode of the pool.
func (o WorkspacesPoolOutput) RunningMode() WorkspacesPoolRunningModePtrOutput {
	return o.ApplyT(func(v *WorkspacesPool) WorkspacesPoolRunningModePtrOutput { return v.RunningMode }).(WorkspacesPoolRunningModePtrOutput)
}

func (o WorkspacesPoolOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *WorkspacesPool) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The amount of time that a pool session remains active after users disconnect. If they try to reconnect to the pool session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new pool instance.
func (o WorkspacesPoolOutput) TimeoutSettings() WorkspacesPoolTimeoutSettingsPtrOutput {
	return o.ApplyT(func(v *WorkspacesPool) WorkspacesPoolTimeoutSettingsPtrOutput { return v.TimeoutSettings }).(WorkspacesPoolTimeoutSettingsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspacesPoolInput)(nil)).Elem(), &WorkspacesPool{})
	pulumi.RegisterOutputType(WorkspacesPoolOutput{})
}
