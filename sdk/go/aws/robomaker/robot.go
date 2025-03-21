// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.
type Robot struct {
	pulumi.CustomResourceState

	// The target architecture of the robot.
	Architecture RobotArchitectureOutput `pulumi:"architecture"`
	// The Amazon Resource Name (ARN) of the robot.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of the fleet.
	Fleet pulumi.StringPtrOutput `pulumi:"fleet"`
	// The Greengrass group id.
	GreengrassGroupId pulumi.StringOutput `pulumi:"greengrassGroupId"`
	// The name for the robot.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A map that contains tag keys and tag values that are attached to the robot.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRobot registers a new resource with the given unique name, arguments, and options.
func NewRobot(ctx *pulumi.Context,
	name string, args *RobotArgs, opts ...pulumi.ResourceOption) (*Robot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Architecture == nil {
		return nil, errors.New("invalid value for required argument 'Architecture'")
	}
	if args.GreengrassGroupId == nil {
		return nil, errors.New("invalid value for required argument 'GreengrassGroupId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"architecture",
		"fleet",
		"greengrassGroupId",
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Robot
	err := ctx.RegisterResource("aws-native:robomaker:Robot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRobot gets an existing Robot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRobot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RobotState, opts ...pulumi.ResourceOption) (*Robot, error) {
	var resource Robot
	err := ctx.ReadResource("aws-native:robomaker:Robot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Robot resources.
type robotState struct {
}

type RobotState struct {
}

func (RobotState) ElementType() reflect.Type {
	return reflect.TypeOf((*robotState)(nil)).Elem()
}

type robotArgs struct {
	// The target architecture of the robot.
	Architecture RobotArchitecture `pulumi:"architecture"`
	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string `pulumi:"fleet"`
	// The Greengrass group id.
	GreengrassGroupId string `pulumi:"greengrassGroupId"`
	// The name for the robot.
	Name *string `pulumi:"name"`
	// A map that contains tag keys and tag values that are attached to the robot.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Robot resource.
type RobotArgs struct {
	// The target architecture of the robot.
	Architecture RobotArchitectureInput
	// The Amazon Resource Name (ARN) of the fleet.
	Fleet pulumi.StringPtrInput
	// The Greengrass group id.
	GreengrassGroupId pulumi.StringInput
	// The name for the robot.
	Name pulumi.StringPtrInput
	// A map that contains tag keys and tag values that are attached to the robot.
	Tags pulumi.StringMapInput
}

func (RobotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*robotArgs)(nil)).Elem()
}

type RobotInput interface {
	pulumi.Input

	ToRobotOutput() RobotOutput
	ToRobotOutputWithContext(ctx context.Context) RobotOutput
}

func (*Robot) ElementType() reflect.Type {
	return reflect.TypeOf((**Robot)(nil)).Elem()
}

func (i *Robot) ToRobotOutput() RobotOutput {
	return i.ToRobotOutputWithContext(context.Background())
}

func (i *Robot) ToRobotOutputWithContext(ctx context.Context) RobotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotOutput)
}

type RobotOutput struct{ *pulumi.OutputState }

func (RobotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Robot)(nil)).Elem()
}

func (o RobotOutput) ToRobotOutput() RobotOutput {
	return o
}

func (o RobotOutput) ToRobotOutputWithContext(ctx context.Context) RobotOutput {
	return o
}

// The target architecture of the robot.
func (o RobotOutput) Architecture() RobotArchitectureOutput {
	return o.ApplyT(func(v *Robot) RobotArchitectureOutput { return v.Architecture }).(RobotArchitectureOutput)
}

// The Amazon Resource Name (ARN) of the robot.
func (o RobotOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Robot) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the fleet.
func (o RobotOutput) Fleet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Robot) pulumi.StringPtrOutput { return v.Fleet }).(pulumi.StringPtrOutput)
}

// The Greengrass group id.
func (o RobotOutput) GreengrassGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Robot) pulumi.StringOutput { return v.GreengrassGroupId }).(pulumi.StringOutput)
}

// The name for the robot.
func (o RobotOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Robot) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A map that contains tag keys and tag values that are attached to the robot.
func (o RobotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Robot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RobotInput)(nil)).Elem(), &Robot{})
	pulumi.RegisterOutputType(RobotOutput{})
}
