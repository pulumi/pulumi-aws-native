// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deadline

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Deadline::Monitor Resource Type
type Monitor struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the monitor.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the monitor that displays on the Deadline Cloud console.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The Amazon Resource Name (ARN) that the IAM Identity Center assigned to the monitor when it was created.
	IdentityCenterApplicationArn pulumi.StringOutput `pulumi:"identityCenterApplicationArn"`
	// The Amazon Resource Name (ARN) of the IAM Identity Center instance responsible for authenticating monitor users.
	IdentityCenterInstanceArn pulumi.StringOutput `pulumi:"identityCenterInstanceArn"`
	// The unique identifier for the monitor.
	MonitorId pulumi.StringOutput `pulumi:"monitorId"`
	// The Amazon Resource Name (ARN) of the IAM role for the monitor. Users of the monitor use this role to access Deadline Cloud resources.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The subdomain used for the monitor URL. The full URL of the monitor is subdomain.Region.deadlinecloud.amazonaws.com.
	Subdomain pulumi.StringOutput `pulumi:"subdomain"`
	// The complete URL of the monitor. The full URL of the monitor is subdomain.Region.deadlinecloud.amazonaws.com.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewMonitor registers a new resource with the given unique name, arguments, and options.
func NewMonitor(ctx *pulumi.Context,
	name string, args *MonitorArgs, opts ...pulumi.ResourceOption) (*Monitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.IdentityCenterInstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'IdentityCenterInstanceArn'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.Subdomain == nil {
		return nil, errors.New("invalid value for required argument 'Subdomain'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"identityCenterInstanceArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Monitor
	err := ctx.RegisterResource("aws-native:deadline:Monitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitor gets an existing Monitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorState, opts ...pulumi.ResourceOption) (*Monitor, error) {
	var resource Monitor
	err := ctx.ReadResource("aws-native:deadline:Monitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Monitor resources.
type monitorState struct {
}

type MonitorState struct {
}

func (MonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorState)(nil)).Elem()
}

type monitorArgs struct {
	// The name of the monitor that displays on the Deadline Cloud console.
	DisplayName string `pulumi:"displayName"`
	// The Amazon Resource Name (ARN) of the IAM Identity Center instance responsible for authenticating monitor users.
	IdentityCenterInstanceArn string `pulumi:"identityCenterInstanceArn"`
	// The Amazon Resource Name (ARN) of the IAM role for the monitor. Users of the monitor use this role to access Deadline Cloud resources.
	RoleArn string `pulumi:"roleArn"`
	// The subdomain used for the monitor URL. The full URL of the monitor is subdomain.Region.deadlinecloud.amazonaws.com.
	Subdomain string `pulumi:"subdomain"`
}

// The set of arguments for constructing a Monitor resource.
type MonitorArgs struct {
	// The name of the monitor that displays on the Deadline Cloud console.
	DisplayName pulumi.StringInput
	// The Amazon Resource Name (ARN) of the IAM Identity Center instance responsible for authenticating monitor users.
	IdentityCenterInstanceArn pulumi.StringInput
	// The Amazon Resource Name (ARN) of the IAM role for the monitor. Users of the monitor use this role to access Deadline Cloud resources.
	RoleArn pulumi.StringInput
	// The subdomain used for the monitor URL. The full URL of the monitor is subdomain.Region.deadlinecloud.amazonaws.com.
	Subdomain pulumi.StringInput
}

func (MonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorArgs)(nil)).Elem()
}

type MonitorInput interface {
	pulumi.Input

	ToMonitorOutput() MonitorOutput
	ToMonitorOutputWithContext(ctx context.Context) MonitorOutput
}

func (*Monitor) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (i *Monitor) ToMonitorOutput() MonitorOutput {
	return i.ToMonitorOutputWithContext(context.Background())
}

func (i *Monitor) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorOutput)
}

type MonitorOutput struct{ *pulumi.OutputState }

func (MonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (o MonitorOutput) ToMonitorOutput() MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return o
}

// The Amazon Resource Name (ARN) of the monitor.
func (o MonitorOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name of the monitor that displays on the Deadline Cloud console.
func (o MonitorOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) that the IAM Identity Center assigned to the monitor when it was created.
func (o MonitorOutput) IdentityCenterApplicationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.IdentityCenterApplicationArn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the IAM Identity Center instance responsible for authenticating monitor users.
func (o MonitorOutput) IdentityCenterInstanceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.IdentityCenterInstanceArn }).(pulumi.StringOutput)
}

// The unique identifier for the monitor.
func (o MonitorOutput) MonitorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.MonitorId }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the IAM role for the monitor. Users of the monitor use this role to access Deadline Cloud resources.
func (o MonitorOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// The subdomain used for the monitor URL. The full URL of the monitor is subdomain.Region.deadlinecloud.amazonaws.com.
func (o MonitorOutput) Subdomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Subdomain }).(pulumi.StringOutput)
}

// The complete URL of the monitor. The full URL of the monitor is subdomain.Region.deadlinecloud.amazonaws.com.
func (o MonitorOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorInput)(nil)).Elem(), &Monitor{})
	pulumi.RegisterOutputType(MonitorOutput{})
}