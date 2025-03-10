// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudWatch::Dashboard
type Dashboard struct {
	pulumi.CustomResourceState

	// The detailed information about the dashboard in JSON format, including the widgets to include and their location on the dashboard
	DashboardBody pulumi.StringOutput `pulumi:"dashboardBody"`
	// The name of the dashboard. The name must be between 1 and 255 characters. If you do not specify a name, one will be generated automatically.
	DashboardName pulumi.StringPtrOutput `pulumi:"dashboardName"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DashboardBody == nil {
		return nil, errors.New("invalid value for required argument 'DashboardBody'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"dashboardName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Dashboard
	err := ctx.RegisterResource("aws-native:cloudwatch:Dashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboard gets an existing Dashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardState, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	var resource Dashboard
	err := ctx.ReadResource("aws-native:cloudwatch:Dashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dashboard resources.
type dashboardState struct {
}

type DashboardState struct {
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	// The detailed information about the dashboard in JSON format, including the widgets to include and their location on the dashboard
	DashboardBody string `pulumi:"dashboardBody"`
	// The name of the dashboard. The name must be between 1 and 255 characters. If you do not specify a name, one will be generated automatically.
	DashboardName *string `pulumi:"dashboardName"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// The detailed information about the dashboard in JSON format, including the widgets to include and their location on the dashboard
	DashboardBody pulumi.StringInput
	// The name of the dashboard. The name must be between 1 and 255 characters. If you do not specify a name, one will be generated automatically.
	DashboardName pulumi.StringPtrInput
}

func (DashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardArgs)(nil)).Elem()
}

type DashboardInput interface {
	pulumi.Input

	ToDashboardOutput() DashboardOutput
	ToDashboardOutputWithContext(ctx context.Context) DashboardOutput
}

func (*Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (i *Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

type DashboardOutput struct{ *pulumi.OutputState }

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

// The detailed information about the dashboard in JSON format, including the widgets to include and their location on the dashboard
func (o DashboardOutput) DashboardBody() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.DashboardBody }).(pulumi.StringOutput)
}

// The name of the dashboard. The name must be between 1 and 255 characters. If you do not specify a name, one will be generated automatically.
func (o DashboardOutput) DashboardName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringPtrOutput { return v.DashboardName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardInput)(nil)).Elem(), &Dashboard{})
	pulumi.RegisterOutputType(DashboardOutput{})
}
