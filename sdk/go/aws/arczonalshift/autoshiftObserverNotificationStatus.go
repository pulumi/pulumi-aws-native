// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package arczonalshift

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::ARCZonalShift::AutoshiftObserverNotificationStatus Resource Type
type AutoshiftObserverNotificationStatus struct {
	pulumi.CustomResourceState

	AccountId pulumi.StringOutput                           `pulumi:"accountId"`
	Region    pulumi.StringOutput                           `pulumi:"region"`
	Status    AutoshiftObserverNotificationStatusEnumOutput `pulumi:"status"`
}

// NewAutoshiftObserverNotificationStatus registers a new resource with the given unique name, arguments, and options.
func NewAutoshiftObserverNotificationStatus(ctx *pulumi.Context,
	name string, args *AutoshiftObserverNotificationStatusArgs, opts ...pulumi.ResourceOption) (*AutoshiftObserverNotificationStatus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"status",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AutoshiftObserverNotificationStatus
	err := ctx.RegisterResource("aws-native:arczonalshift:AutoshiftObserverNotificationStatus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoshiftObserverNotificationStatus gets an existing AutoshiftObserverNotificationStatus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoshiftObserverNotificationStatus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoshiftObserverNotificationStatusState, opts ...pulumi.ResourceOption) (*AutoshiftObserverNotificationStatus, error) {
	var resource AutoshiftObserverNotificationStatus
	err := ctx.ReadResource("aws-native:arczonalshift:AutoshiftObserverNotificationStatus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoshiftObserverNotificationStatus resources.
type autoshiftObserverNotificationStatusState struct {
}

type AutoshiftObserverNotificationStatusState struct {
}

func (AutoshiftObserverNotificationStatusState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoshiftObserverNotificationStatusState)(nil)).Elem()
}

type autoshiftObserverNotificationStatusArgs struct {
	Status AutoshiftObserverNotificationStatusEnum `pulumi:"status"`
}

// The set of arguments for constructing a AutoshiftObserverNotificationStatus resource.
type AutoshiftObserverNotificationStatusArgs struct {
	Status AutoshiftObserverNotificationStatusEnumInput
}

func (AutoshiftObserverNotificationStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoshiftObserverNotificationStatusArgs)(nil)).Elem()
}

type AutoshiftObserverNotificationStatusInput interface {
	pulumi.Input

	ToAutoshiftObserverNotificationStatusOutput() AutoshiftObserverNotificationStatusOutput
	ToAutoshiftObserverNotificationStatusOutputWithContext(ctx context.Context) AutoshiftObserverNotificationStatusOutput
}

func (*AutoshiftObserverNotificationStatus) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoshiftObserverNotificationStatus)(nil)).Elem()
}

func (i *AutoshiftObserverNotificationStatus) ToAutoshiftObserverNotificationStatusOutput() AutoshiftObserverNotificationStatusOutput {
	return i.ToAutoshiftObserverNotificationStatusOutputWithContext(context.Background())
}

func (i *AutoshiftObserverNotificationStatus) ToAutoshiftObserverNotificationStatusOutputWithContext(ctx context.Context) AutoshiftObserverNotificationStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoshiftObserverNotificationStatusOutput)
}

type AutoshiftObserverNotificationStatusOutput struct{ *pulumi.OutputState }

func (AutoshiftObserverNotificationStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoshiftObserverNotificationStatus)(nil)).Elem()
}

func (o AutoshiftObserverNotificationStatusOutput) ToAutoshiftObserverNotificationStatusOutput() AutoshiftObserverNotificationStatusOutput {
	return o
}

func (o AutoshiftObserverNotificationStatusOutput) ToAutoshiftObserverNotificationStatusOutputWithContext(ctx context.Context) AutoshiftObserverNotificationStatusOutput {
	return o
}

func (o AutoshiftObserverNotificationStatusOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoshiftObserverNotificationStatus) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o AutoshiftObserverNotificationStatusOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoshiftObserverNotificationStatus) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o AutoshiftObserverNotificationStatusOutput) Status() AutoshiftObserverNotificationStatusEnumOutput {
	return o.ApplyT(func(v *AutoshiftObserverNotificationStatus) AutoshiftObserverNotificationStatusEnumOutput {
		return v.Status
	}).(AutoshiftObserverNotificationStatusEnumOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoshiftObserverNotificationStatusInput)(nil)).Elem(), &AutoshiftObserverNotificationStatus{})
	pulumi.RegisterOutputType(AutoshiftObserverNotificationStatusOutput{})
}
