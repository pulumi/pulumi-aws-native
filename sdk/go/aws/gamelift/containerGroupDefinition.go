// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::GameLift::ContainerGroupDefinition resource creates an Amazon GameLift container group definition.
type ContainerGroupDefinition struct {
	pulumi.CustomResourceState

	// A collection of container definitions that define the containers in this group.
	ContainerDefinitions ContainerGroupDefinitionContainerDefinitionArrayOutput `pulumi:"containerDefinitions"`
	// The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift container group resource and uniquely identifies it across all AWS Regions.
	ContainerGroupDefinitionArn pulumi.StringOutput `pulumi:"containerGroupDefinitionArn"`
	// A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example "1469498468.057").
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// A descriptive label for the container group definition.
	Name pulumi.StringOutput `pulumi:"name"`
	// The operating system of the container group
	OperatingSystem ContainerGroupDefinitionOperatingSystemOutput `pulumi:"operatingSystem"`
	// Specifies whether the container group includes replica or daemon containers.
	SchedulingStrategy ContainerGroupDefinitionSchedulingStrategyPtrOutput `pulumi:"schedulingStrategy"`
	// A specific ContainerGroupDefinition version to be updated
	SourceVersionNumber pulumi.IntPtrOutput `pulumi:"sourceVersionNumber"`
	// A string indicating ContainerGroupDefinition status.
	Status ContainerGroupDefinitionStatusOutput `pulumi:"status"`
	// A string indicating the reason for ContainerGroupDefinition status.
	StatusReason pulumi.StringOutput `pulumi:"statusReason"`
	// A collection of support container definitions that define the containers in this group.
	SupportContainerDefinitions pulumi.ArrayOutput `pulumi:"supportContainerDefinitions"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The maximum number of CPU units reserved for this container group. The value is expressed as an integer amount of CPU units. (1 vCPU is equal to 1024 CPU units.)
	TotalCpuLimit pulumi.IntOutput `pulumi:"totalCpuLimit"`
	// The maximum amount of memory (in MiB) to allocate for this container group.
	TotalMemoryLimit pulumi.IntOutput `pulumi:"totalMemoryLimit"`
}

// NewContainerGroupDefinition registers a new resource with the given unique name, arguments, and options.
func NewContainerGroupDefinition(ctx *pulumi.Context,
	name string, args *ContainerGroupDefinitionArgs, opts ...pulumi.ResourceOption) (*ContainerGroupDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'ContainerDefinitions'")
	}
	if args.OperatingSystem == nil {
		return nil, errors.New("invalid value for required argument 'OperatingSystem'")
	}
	if args.TotalCpuLimit == nil {
		return nil, errors.New("invalid value for required argument 'TotalCpuLimit'")
	}
	if args.TotalMemoryLimit == nil {
		return nil, errors.New("invalid value for required argument 'TotalMemoryLimit'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"containerDefinitions[*]",
		"name",
		"operatingSystem",
		"schedulingStrategy",
		"totalCpuLimit",
		"totalMemoryLimit",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerGroupDefinition
	err := ctx.RegisterResource("aws-native:gamelift:ContainerGroupDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerGroupDefinition gets an existing ContainerGroupDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerGroupDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerGroupDefinitionState, opts ...pulumi.ResourceOption) (*ContainerGroupDefinition, error) {
	var resource ContainerGroupDefinition
	err := ctx.ReadResource("aws-native:gamelift:ContainerGroupDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerGroupDefinition resources.
type containerGroupDefinitionState struct {
}

type ContainerGroupDefinitionState struct {
}

func (ContainerGroupDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerGroupDefinitionState)(nil)).Elem()
}

type containerGroupDefinitionArgs struct {
	// A collection of container definitions that define the containers in this group.
	ContainerDefinitions []ContainerGroupDefinitionContainerDefinition `pulumi:"containerDefinitions"`
	// A descriptive label for the container group definition.
	Name *string `pulumi:"name"`
	// The operating system of the container group
	OperatingSystem ContainerGroupDefinitionOperatingSystem `pulumi:"operatingSystem"`
	// Specifies whether the container group includes replica or daemon containers.
	SchedulingStrategy *ContainerGroupDefinitionSchedulingStrategy `pulumi:"schedulingStrategy"`
	// A specific ContainerGroupDefinition version to be updated
	SourceVersionNumber *int `pulumi:"sourceVersionNumber"`
	// A collection of support container definitions that define the containers in this group.
	SupportContainerDefinitions []interface{} `pulumi:"supportContainerDefinitions"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The maximum number of CPU units reserved for this container group. The value is expressed as an integer amount of CPU units. (1 vCPU is equal to 1024 CPU units.)
	TotalCpuLimit int `pulumi:"totalCpuLimit"`
	// The maximum amount of memory (in MiB) to allocate for this container group.
	TotalMemoryLimit int `pulumi:"totalMemoryLimit"`
}

// The set of arguments for constructing a ContainerGroupDefinition resource.
type ContainerGroupDefinitionArgs struct {
	// A collection of container definitions that define the containers in this group.
	ContainerDefinitions ContainerGroupDefinitionContainerDefinitionArrayInput
	// A descriptive label for the container group definition.
	Name pulumi.StringPtrInput
	// The operating system of the container group
	OperatingSystem ContainerGroupDefinitionOperatingSystemInput
	// Specifies whether the container group includes replica or daemon containers.
	SchedulingStrategy ContainerGroupDefinitionSchedulingStrategyPtrInput
	// A specific ContainerGroupDefinition version to be updated
	SourceVersionNumber pulumi.IntPtrInput
	// A collection of support container definitions that define the containers in this group.
	SupportContainerDefinitions pulumi.ArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayInput
	// The maximum number of CPU units reserved for this container group. The value is expressed as an integer amount of CPU units. (1 vCPU is equal to 1024 CPU units.)
	TotalCpuLimit pulumi.IntInput
	// The maximum amount of memory (in MiB) to allocate for this container group.
	TotalMemoryLimit pulumi.IntInput
}

func (ContainerGroupDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerGroupDefinitionArgs)(nil)).Elem()
}

type ContainerGroupDefinitionInput interface {
	pulumi.Input

	ToContainerGroupDefinitionOutput() ContainerGroupDefinitionOutput
	ToContainerGroupDefinitionOutputWithContext(ctx context.Context) ContainerGroupDefinitionOutput
}

func (*ContainerGroupDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupDefinition)(nil)).Elem()
}

func (i *ContainerGroupDefinition) ToContainerGroupDefinitionOutput() ContainerGroupDefinitionOutput {
	return i.ToContainerGroupDefinitionOutputWithContext(context.Background())
}

func (i *ContainerGroupDefinition) ToContainerGroupDefinitionOutputWithContext(ctx context.Context) ContainerGroupDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupDefinitionOutput)
}

type ContainerGroupDefinitionOutput struct{ *pulumi.OutputState }

func (ContainerGroupDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupDefinition)(nil)).Elem()
}

func (o ContainerGroupDefinitionOutput) ToContainerGroupDefinitionOutput() ContainerGroupDefinitionOutput {
	return o
}

func (o ContainerGroupDefinitionOutput) ToContainerGroupDefinitionOutputWithContext(ctx context.Context) ContainerGroupDefinitionOutput {
	return o
}

// A collection of container definitions that define the containers in this group.
func (o ContainerGroupDefinitionOutput) ContainerDefinitions() ContainerGroupDefinitionContainerDefinitionArrayOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) ContainerGroupDefinitionContainerDefinitionArrayOutput {
		return v.ContainerDefinitions
	}).(ContainerGroupDefinitionContainerDefinitionArrayOutput)
}

// The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift container group resource and uniquely identifies it across all AWS Regions.
func (o ContainerGroupDefinitionOutput) ContainerGroupDefinitionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) pulumi.StringOutput { return v.ContainerGroupDefinitionArn }).(pulumi.StringOutput)
}

// A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example "1469498468.057").
func (o ContainerGroupDefinitionOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// A descriptive label for the container group definition.
func (o ContainerGroupDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The operating system of the container group
func (o ContainerGroupDefinitionOutput) OperatingSystem() ContainerGroupDefinitionOperatingSystemOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) ContainerGroupDefinitionOperatingSystemOutput {
		return v.OperatingSystem
	}).(ContainerGroupDefinitionOperatingSystemOutput)
}

// Specifies whether the container group includes replica or daemon containers.
func (o ContainerGroupDefinitionOutput) SchedulingStrategy() ContainerGroupDefinitionSchedulingStrategyPtrOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) ContainerGroupDefinitionSchedulingStrategyPtrOutput {
		return v.SchedulingStrategy
	}).(ContainerGroupDefinitionSchedulingStrategyPtrOutput)
}

// A specific ContainerGroupDefinition version to be updated
func (o ContainerGroupDefinitionOutput) SourceVersionNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) pulumi.IntPtrOutput { return v.SourceVersionNumber }).(pulumi.IntPtrOutput)
}

// A string indicating ContainerGroupDefinition status.
func (o ContainerGroupDefinitionOutput) Status() ContainerGroupDefinitionStatusOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) ContainerGroupDefinitionStatusOutput { return v.Status }).(ContainerGroupDefinitionStatusOutput)
}

// A string indicating the reason for ContainerGroupDefinition status.
func (o ContainerGroupDefinitionOutput) StatusReason() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) pulumi.StringOutput { return v.StatusReason }).(pulumi.StringOutput)
}

// A collection of support container definitions that define the containers in this group.
func (o ContainerGroupDefinitionOutput) SupportContainerDefinitions() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) pulumi.ArrayOutput { return v.SupportContainerDefinitions }).(pulumi.ArrayOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ContainerGroupDefinitionOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The maximum number of CPU units reserved for this container group. The value is expressed as an integer amount of CPU units. (1 vCPU is equal to 1024 CPU units.)
func (o ContainerGroupDefinitionOutput) TotalCpuLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) pulumi.IntOutput { return v.TotalCpuLimit }).(pulumi.IntOutput)
}

// The maximum amount of memory (in MiB) to allocate for this container group.
func (o ContainerGroupDefinitionOutput) TotalMemoryLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *ContainerGroupDefinition) pulumi.IntOutput { return v.TotalMemoryLimit }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupDefinitionInput)(nil)).Elem(), &ContainerGroupDefinition{})
	pulumi.RegisterOutputType(ContainerGroupDefinitionOutput{})
}
