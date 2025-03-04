// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package medialive

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::MediaLive::SignalMap Resource Type
func LookupSignalMap(ctx *pulumi.Context, args *LookupSignalMapArgs, opts ...pulumi.InvokeOption) (*LookupSignalMapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSignalMapResult
	err := ctx.Invoke("aws-native:medialive:getSignalMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalMapArgs struct {
	Identifier string `pulumi:"identifier"`
}

type LookupSignalMapResult struct {
	// A signal map's ARN (Amazon Resource Name)
	Arn *string `pulumi:"arn"`
	// An alarm template group's id.
	CloudWatchAlarmTemplateGroupIds []string `pulumi:"cloudWatchAlarmTemplateGroupIds"`
	// The date and time of resource creation.
	CreatedAt *string `pulumi:"createdAt"`
	// A resource's optional description.
	Description *string `pulumi:"description"`
	// A top-level supported AWS resource ARN to discovery a signal map from.
	DiscoveryEntryPointArn *string `pulumi:"discoveryEntryPointArn"`
	// Error message associated with a failed creation or failed update attempt of a signal map.
	ErrorMessage *string `pulumi:"errorMessage"`
	// An eventbridge rule template group's id.
	EventBridgeRuleTemplateGroupIds []string                          `pulumi:"eventBridgeRuleTemplateGroupIds"`
	FailedMediaResourceMap          map[string]SignalMapMediaResource `pulumi:"failedMediaResourceMap"`
	// A signal map's id.
	Id         *string `pulumi:"id"`
	Identifier *string `pulumi:"identifier"`
	// The date and time of latest discovery.
	LastDiscoveredAt                *string                               `pulumi:"lastDiscoveredAt"`
	LastSuccessfulMonitorDeployment *SignalMapSuccessfulMonitorDeployment `pulumi:"lastSuccessfulMonitorDeployment"`
	MediaResourceMap                map[string]SignalMapMediaResource     `pulumi:"mediaResourceMap"`
	// The date and time of latest resource modification.
	ModifiedAt *string `pulumi:"modifiedAt"`
	// If true, there are pending monitor changes for this signal map that can be deployed.
	MonitorChangesPendingDeployment *bool                       `pulumi:"monitorChangesPendingDeployment"`
	MonitorDeployment               *SignalMapMonitorDeployment `pulumi:"monitorDeployment"`
	// A resource's name. Names must be unique within the scope of a resource type in a specific region.
	Name *string `pulumi:"name"`
	// A signal map's current status, which is dependent on its lifecycle actions or associated jobs.
	Status *SignalMapStatus `pulumi:"status"`
}

func LookupSignalMapOutput(ctx *pulumi.Context, args LookupSignalMapOutputArgs, opts ...pulumi.InvokeOption) LookupSignalMapResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSignalMapResultOutput, error) {
			args := v.(LookupSignalMapArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:medialive:getSignalMap", args, LookupSignalMapResultOutput{}, options).(LookupSignalMapResultOutput), nil
		}).(LookupSignalMapResultOutput)
}

type LookupSignalMapOutputArgs struct {
	Identifier pulumi.StringInput `pulumi:"identifier"`
}

func (LookupSignalMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalMapArgs)(nil)).Elem()
}

type LookupSignalMapResultOutput struct{ *pulumi.OutputState }

func (LookupSignalMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalMapResult)(nil)).Elem()
}

func (o LookupSignalMapResultOutput) ToLookupSignalMapResultOutput() LookupSignalMapResultOutput {
	return o
}

func (o LookupSignalMapResultOutput) ToLookupSignalMapResultOutputWithContext(ctx context.Context) LookupSignalMapResultOutput {
	return o
}

// A signal map's ARN (Amazon Resource Name)
func (o LookupSignalMapResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// An alarm template group's id.
func (o LookupSignalMapResultOutput) CloudWatchAlarmTemplateGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSignalMapResult) []string { return v.CloudWatchAlarmTemplateGroupIds }).(pulumi.StringArrayOutput)
}

// The date and time of resource creation.
func (o LookupSignalMapResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// A resource's optional description.
func (o LookupSignalMapResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A top-level supported AWS resource ARN to discovery a signal map from.
func (o LookupSignalMapResultOutput) DiscoveryEntryPointArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *string { return v.DiscoveryEntryPointArn }).(pulumi.StringPtrOutput)
}

// Error message associated with a failed creation or failed update attempt of a signal map.
func (o LookupSignalMapResultOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

// An eventbridge rule template group's id.
func (o LookupSignalMapResultOutput) EventBridgeRuleTemplateGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSignalMapResult) []string { return v.EventBridgeRuleTemplateGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupSignalMapResultOutput) FailedMediaResourceMap() SignalMapMediaResourceMapOutput {
	return o.ApplyT(func(v LookupSignalMapResult) map[string]SignalMapMediaResource { return v.FailedMediaResourceMap }).(SignalMapMediaResourceMapOutput)
}

// A signal map's id.
func (o LookupSignalMapResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSignalMapResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// The date and time of latest discovery.
func (o LookupSignalMapResultOutput) LastDiscoveredAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *string { return v.LastDiscoveredAt }).(pulumi.StringPtrOutput)
}

func (o LookupSignalMapResultOutput) LastSuccessfulMonitorDeployment() SignalMapSuccessfulMonitorDeploymentPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *SignalMapSuccessfulMonitorDeployment {
		return v.LastSuccessfulMonitorDeployment
	}).(SignalMapSuccessfulMonitorDeploymentPtrOutput)
}

func (o LookupSignalMapResultOutput) MediaResourceMap() SignalMapMediaResourceMapOutput {
	return o.ApplyT(func(v LookupSignalMapResult) map[string]SignalMapMediaResource { return v.MediaResourceMap }).(SignalMapMediaResourceMapOutput)
}

// The date and time of latest resource modification.
func (o LookupSignalMapResultOutput) ModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *string { return v.ModifiedAt }).(pulumi.StringPtrOutput)
}

// If true, there are pending monitor changes for this signal map that can be deployed.
func (o LookupSignalMapResultOutput) MonitorChangesPendingDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *bool { return v.MonitorChangesPendingDeployment }).(pulumi.BoolPtrOutput)
}

func (o LookupSignalMapResultOutput) MonitorDeployment() SignalMapMonitorDeploymentPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *SignalMapMonitorDeployment { return v.MonitorDeployment }).(SignalMapMonitorDeploymentPtrOutput)
}

// A resource's name. Names must be unique within the scope of a resource type in a specific region.
func (o LookupSignalMapResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A signal map's current status, which is dependent on its lifecycle actions or associated jobs.
func (o LookupSignalMapResultOutput) Status() SignalMapStatusPtrOutput {
	return o.ApplyT(func(v LookupSignalMapResult) *SignalMapStatus { return v.Status }).(SignalMapStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalMapResultOutput{})
}
