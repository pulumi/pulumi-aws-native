// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package internetmonitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a monitor, which defines the monitoring boundaries for measurements that Internet Monitor publishes information about for an application
type Monitor struct {
	pulumi.CustomResourceState

	CreatedAt                       pulumi.StringOutput                             `pulumi:"createdAt"`
	InternetMeasurementsLogDelivery MonitorInternetMeasurementsLogDeliveryPtrOutput `pulumi:"internetMeasurementsLogDelivery"`
	MaxCityNetworksToMonitor        pulumi.IntPtrOutput                             `pulumi:"maxCityNetworksToMonitor"`
	ModifiedAt                      pulumi.StringOutput                             `pulumi:"modifiedAt"`
	MonitorArn                      pulumi.StringOutput                             `pulumi:"monitorArn"`
	MonitorName                     pulumi.StringOutput                             `pulumi:"monitorName"`
	ProcessingStatus                MonitorProcessingStatusCodeOutput               `pulumi:"processingStatus"`
	ProcessingStatusInfo            pulumi.StringOutput                             `pulumi:"processingStatusInfo"`
	Resources                       pulumi.StringArrayOutput                        `pulumi:"resources"`
	ResourcesToAdd                  pulumi.StringArrayOutput                        `pulumi:"resourcesToAdd"`
	ResourcesToRemove               pulumi.StringArrayOutput                        `pulumi:"resourcesToRemove"`
	Status                          MonitorConfigStatePtrOutput                     `pulumi:"status"`
	Tags                            MonitorTagArrayOutput                           `pulumi:"tags"`
}

// NewMonitor registers a new resource with the given unique name, arguments, and options.
func NewMonitor(ctx *pulumi.Context,
	name string, args *MonitorArgs, opts ...pulumi.ResourceOption) (*Monitor, error) {
	if args == nil {
		args = &MonitorArgs{}
	}

	var resource Monitor
	err := ctx.RegisterResource("aws-native:internetmonitor:Monitor", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:internetmonitor:Monitor", name, id, state, &resource, opts...)
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
	InternetMeasurementsLogDelivery *MonitorInternetMeasurementsLogDelivery `pulumi:"internetMeasurementsLogDelivery"`
	MaxCityNetworksToMonitor        *int                                    `pulumi:"maxCityNetworksToMonitor"`
	MonitorName                     *string                                 `pulumi:"monitorName"`
	Resources                       []string                                `pulumi:"resources"`
	ResourcesToAdd                  []string                                `pulumi:"resourcesToAdd"`
	ResourcesToRemove               []string                                `pulumi:"resourcesToRemove"`
	Status                          *MonitorConfigState                     `pulumi:"status"`
	Tags                            []MonitorTag                            `pulumi:"tags"`
}

// The set of arguments for constructing a Monitor resource.
type MonitorArgs struct {
	InternetMeasurementsLogDelivery MonitorInternetMeasurementsLogDeliveryPtrInput
	MaxCityNetworksToMonitor        pulumi.IntPtrInput
	MonitorName                     pulumi.StringPtrInput
	Resources                       pulumi.StringArrayInput
	ResourcesToAdd                  pulumi.StringArrayInput
	ResourcesToRemove               pulumi.StringArrayInput
	Status                          MonitorConfigStatePtrInput
	Tags                            MonitorTagArrayInput
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

func (o MonitorOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o MonitorOutput) InternetMeasurementsLogDelivery() MonitorInternetMeasurementsLogDeliveryPtrOutput {
	return o.ApplyT(func(v *Monitor) MonitorInternetMeasurementsLogDeliveryPtrOutput {
		return v.InternetMeasurementsLogDelivery
	}).(MonitorInternetMeasurementsLogDeliveryPtrOutput)
}

func (o MonitorOutput) MaxCityNetworksToMonitor() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.IntPtrOutput { return v.MaxCityNetworksToMonitor }).(pulumi.IntPtrOutput)
}

func (o MonitorOutput) ModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.ModifiedAt }).(pulumi.StringOutput)
}

func (o MonitorOutput) MonitorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.MonitorArn }).(pulumi.StringOutput)
}

func (o MonitorOutput) MonitorName() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.MonitorName }).(pulumi.StringOutput)
}

func (o MonitorOutput) ProcessingStatus() MonitorProcessingStatusCodeOutput {
	return o.ApplyT(func(v *Monitor) MonitorProcessingStatusCodeOutput { return v.ProcessingStatus }).(MonitorProcessingStatusCodeOutput)
}

func (o MonitorOutput) ProcessingStatusInfo() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.ProcessingStatusInfo }).(pulumi.StringOutput)
}

func (o MonitorOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringArrayOutput { return v.Resources }).(pulumi.StringArrayOutput)
}

func (o MonitorOutput) ResourcesToAdd() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringArrayOutput { return v.ResourcesToAdd }).(pulumi.StringArrayOutput)
}

func (o MonitorOutput) ResourcesToRemove() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringArrayOutput { return v.ResourcesToRemove }).(pulumi.StringArrayOutput)
}

func (o MonitorOutput) Status() MonitorConfigStatePtrOutput {
	return o.ApplyT(func(v *Monitor) MonitorConfigStatePtrOutput { return v.Status }).(MonitorConfigStatePtrOutput)
}

func (o MonitorOutput) Tags() MonitorTagArrayOutput {
	return o.ApplyT(func(v *Monitor) MonitorTagArrayOutput { return v.Tags }).(MonitorTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorInput)(nil)).Elem(), &Monitor{})
	pulumi.RegisterOutputType(MonitorOutput{})
}