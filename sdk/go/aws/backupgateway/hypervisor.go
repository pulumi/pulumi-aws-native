// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package backupgateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::BackupGateway::Hypervisor Resource Type
type Hypervisor struct {
	pulumi.CustomResourceState

	Host          pulumi.StringPtrOutput   `pulumi:"host"`
	HypervisorArn pulumi.StringOutput      `pulumi:"hypervisorArn"`
	KmsKeyArn     pulumi.StringPtrOutput   `pulumi:"kmsKeyArn"`
	LogGroupArn   pulumi.StringPtrOutput   `pulumi:"logGroupArn"`
	Name          pulumi.StringPtrOutput   `pulumi:"name"`
	Password      pulumi.StringPtrOutput   `pulumi:"password"`
	Tags          HypervisorTagArrayOutput `pulumi:"tags"`
	Username      pulumi.StringPtrOutput   `pulumi:"username"`
}

// NewHypervisor registers a new resource with the given unique name, arguments, and options.
func NewHypervisor(ctx *pulumi.Context,
	name string, args *HypervisorArgs, opts ...pulumi.ResourceOption) (*Hypervisor, error) {
	if args == nil {
		args = &HypervisorArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Hypervisor
	err := ctx.RegisterResource("aws-native:backupgateway:Hypervisor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHypervisor gets an existing Hypervisor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHypervisor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HypervisorState, opts ...pulumi.ResourceOption) (*Hypervisor, error) {
	var resource Hypervisor
	err := ctx.ReadResource("aws-native:backupgateway:Hypervisor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hypervisor resources.
type hypervisorState struct {
}

type HypervisorState struct {
}

func (HypervisorState) ElementType() reflect.Type {
	return reflect.TypeOf((*hypervisorState)(nil)).Elem()
}

type hypervisorArgs struct {
	Host        *string         `pulumi:"host"`
	KmsKeyArn   *string         `pulumi:"kmsKeyArn"`
	LogGroupArn *string         `pulumi:"logGroupArn"`
	Name        *string         `pulumi:"name"`
	Password    *string         `pulumi:"password"`
	Tags        []HypervisorTag `pulumi:"tags"`
	Username    *string         `pulumi:"username"`
}

// The set of arguments for constructing a Hypervisor resource.
type HypervisorArgs struct {
	Host        pulumi.StringPtrInput
	KmsKeyArn   pulumi.StringPtrInput
	LogGroupArn pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Password    pulumi.StringPtrInput
	Tags        HypervisorTagArrayInput
	Username    pulumi.StringPtrInput
}

func (HypervisorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hypervisorArgs)(nil)).Elem()
}

type HypervisorInput interface {
	pulumi.Input

	ToHypervisorOutput() HypervisorOutput
	ToHypervisorOutputWithContext(ctx context.Context) HypervisorOutput
}

func (*Hypervisor) ElementType() reflect.Type {
	return reflect.TypeOf((**Hypervisor)(nil)).Elem()
}

func (i *Hypervisor) ToHypervisorOutput() HypervisorOutput {
	return i.ToHypervisorOutputWithContext(context.Background())
}

func (i *Hypervisor) ToHypervisorOutputWithContext(ctx context.Context) HypervisorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HypervisorOutput)
}

type HypervisorOutput struct{ *pulumi.OutputState }

func (HypervisorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hypervisor)(nil)).Elem()
}

func (o HypervisorOutput) ToHypervisorOutput() HypervisorOutput {
	return o
}

func (o HypervisorOutput) ToHypervisorOutputWithContext(ctx context.Context) HypervisorOutput {
	return o
}

func (o HypervisorOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hypervisor) pulumi.StringPtrOutput { return v.Host }).(pulumi.StringPtrOutput)
}

func (o HypervisorOutput) HypervisorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Hypervisor) pulumi.StringOutput { return v.HypervisorArn }).(pulumi.StringOutput)
}

func (o HypervisorOutput) KmsKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hypervisor) pulumi.StringPtrOutput { return v.KmsKeyArn }).(pulumi.StringPtrOutput)
}

func (o HypervisorOutput) LogGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hypervisor) pulumi.StringPtrOutput { return v.LogGroupArn }).(pulumi.StringPtrOutput)
}

func (o HypervisorOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hypervisor) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HypervisorOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hypervisor) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o HypervisorOutput) Tags() HypervisorTagArrayOutput {
	return o.ApplyT(func(v *Hypervisor) HypervisorTagArrayOutput { return v.Tags }).(HypervisorTagArrayOutput)
}

func (o HypervisorOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hypervisor) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HypervisorInput)(nil)).Elem(), &Hypervisor{})
	pulumi.RegisterOutputType(HypervisorOutput{})
}