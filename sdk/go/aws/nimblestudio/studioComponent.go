// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nimblestudio

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::NimbleStudio::StudioComponent
type StudioComponent struct {
	pulumi.CustomResourceState

	Configuration         StudioComponentConfigurationPtrOutput             `pulumi:"configuration"`
	Description           pulumi.StringPtrOutput                            `pulumi:"description"`
	Ec2SecurityGroupIds   pulumi.StringArrayOutput                          `pulumi:"ec2SecurityGroupIds"`
	InitializationScripts StudioComponentInitializationScriptArrayOutput    `pulumi:"initializationScripts"`
	Name                  pulumi.StringOutput                               `pulumi:"name"`
	ScriptParameters      StudioComponentScriptParameterKeyValueArrayOutput `pulumi:"scriptParameters"`
	StudioComponentId     pulumi.StringOutput                               `pulumi:"studioComponentId"`
	StudioId              pulumi.StringOutput                               `pulumi:"studioId"`
	Subtype               pulumi.StringPtrOutput                            `pulumi:"subtype"`
	Tags                  pulumi.StringMapOutput                            `pulumi:"tags"`
	Type                  pulumi.StringOutput                               `pulumi:"type"`
}

// NewStudioComponent registers a new resource with the given unique name, arguments, and options.
func NewStudioComponent(ctx *pulumi.Context,
	name string, args *StudioComponentArgs, opts ...pulumi.ResourceOption) (*StudioComponent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StudioId == nil {
		return nil, errors.New("invalid value for required argument 'StudioId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"studioId",
		"subtype",
		"tags.*",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StudioComponent
	err := ctx.RegisterResource("aws-native:nimblestudio:StudioComponent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStudioComponent gets an existing StudioComponent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStudioComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StudioComponentState, opts ...pulumi.ResourceOption) (*StudioComponent, error) {
	var resource StudioComponent
	err := ctx.ReadResource("aws-native:nimblestudio:StudioComponent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StudioComponent resources.
type studioComponentState struct {
}

type StudioComponentState struct {
}

func (StudioComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*studioComponentState)(nil)).Elem()
}

type studioComponentArgs struct {
	Configuration         *StudioComponentConfiguration            `pulumi:"configuration"`
	Description           *string                                  `pulumi:"description"`
	Ec2SecurityGroupIds   []string                                 `pulumi:"ec2SecurityGroupIds"`
	InitializationScripts []StudioComponentInitializationScript    `pulumi:"initializationScripts"`
	Name                  *string                                  `pulumi:"name"`
	ScriptParameters      []StudioComponentScriptParameterKeyValue `pulumi:"scriptParameters"`
	StudioId              string                                   `pulumi:"studioId"`
	Subtype               *string                                  `pulumi:"subtype"`
	Tags                  map[string]string                        `pulumi:"tags"`
	Type                  string                                   `pulumi:"type"`
}

// The set of arguments for constructing a StudioComponent resource.
type StudioComponentArgs struct {
	Configuration         StudioComponentConfigurationPtrInput
	Description           pulumi.StringPtrInput
	Ec2SecurityGroupIds   pulumi.StringArrayInput
	InitializationScripts StudioComponentInitializationScriptArrayInput
	Name                  pulumi.StringPtrInput
	ScriptParameters      StudioComponentScriptParameterKeyValueArrayInput
	StudioId              pulumi.StringInput
	Subtype               pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
	Type                  pulumi.StringInput
}

func (StudioComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*studioComponentArgs)(nil)).Elem()
}

type StudioComponentInput interface {
	pulumi.Input

	ToStudioComponentOutput() StudioComponentOutput
	ToStudioComponentOutputWithContext(ctx context.Context) StudioComponentOutput
}

func (*StudioComponent) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioComponent)(nil)).Elem()
}

func (i *StudioComponent) ToStudioComponentOutput() StudioComponentOutput {
	return i.ToStudioComponentOutputWithContext(context.Background())
}

func (i *StudioComponent) ToStudioComponentOutputWithContext(ctx context.Context) StudioComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioComponentOutput)
}

type StudioComponentOutput struct{ *pulumi.OutputState }

func (StudioComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioComponent)(nil)).Elem()
}

func (o StudioComponentOutput) ToStudioComponentOutput() StudioComponentOutput {
	return o
}

func (o StudioComponentOutput) ToStudioComponentOutputWithContext(ctx context.Context) StudioComponentOutput {
	return o
}

func (o StudioComponentOutput) Configuration() StudioComponentConfigurationPtrOutput {
	return o.ApplyT(func(v *StudioComponent) StudioComponentConfigurationPtrOutput { return v.Configuration }).(StudioComponentConfigurationPtrOutput)
}

func (o StudioComponentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StudioComponent) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StudioComponentOutput) Ec2SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StudioComponent) pulumi.StringArrayOutput { return v.Ec2SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o StudioComponentOutput) InitializationScripts() StudioComponentInitializationScriptArrayOutput {
	return o.ApplyT(func(v *StudioComponent) StudioComponentInitializationScriptArrayOutput {
		return v.InitializationScripts
	}).(StudioComponentInitializationScriptArrayOutput)
}

func (o StudioComponentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioComponent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StudioComponentOutput) ScriptParameters() StudioComponentScriptParameterKeyValueArrayOutput {
	return o.ApplyT(func(v *StudioComponent) StudioComponentScriptParameterKeyValueArrayOutput { return v.ScriptParameters }).(StudioComponentScriptParameterKeyValueArrayOutput)
}

func (o StudioComponentOutput) StudioComponentId() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioComponent) pulumi.StringOutput { return v.StudioComponentId }).(pulumi.StringOutput)
}

func (o StudioComponentOutput) StudioId() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioComponent) pulumi.StringOutput { return v.StudioId }).(pulumi.StringOutput)
}

func (o StudioComponentOutput) Subtype() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StudioComponent) pulumi.StringPtrOutput { return v.Subtype }).(pulumi.StringPtrOutput)
}

func (o StudioComponentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StudioComponent) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o StudioComponentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioComponent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StudioComponentInput)(nil)).Elem(), &StudioComponent{})
	pulumi.RegisterOutputType(StudioComponentOutput{})
}
