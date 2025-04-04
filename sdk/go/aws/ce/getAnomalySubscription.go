// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ce

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. Create subscription to be notified
func LookupAnomalySubscription(ctx *pulumi.Context, args *LookupAnomalySubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupAnomalySubscriptionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAnomalySubscriptionResult
	err := ctx.Invoke("aws-native:ce:getAnomalySubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnomalySubscriptionArgs struct {
	// The `AnomalySubscription` Amazon Resource Name (ARN).
	SubscriptionArn string `pulumi:"subscriptionArn"`
}

type LookupAnomalySubscriptionResult struct {
	// The accountId
	AccountId *string `pulumi:"accountId"`
	// The frequency at which anomaly reports are sent over email.
	Frequency *AnomalySubscriptionFrequency `pulumi:"frequency"`
	// A list of cost anomaly monitors.
	MonitorArnList []string `pulumi:"monitorArnList"`
	// A list of subscriber
	Subscribers []AnomalySubscriptionSubscriber `pulumi:"subscribers"`
	// The `AnomalySubscription` Amazon Resource Name (ARN).
	SubscriptionArn *string `pulumi:"subscriptionArn"`
	// The name of the subscription.
	SubscriptionName *string `pulumi:"subscriptionName"`
	// The dollar value that triggers a notification if the threshold is exceeded.
	Threshold *float64 `pulumi:"threshold"`
	// An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.
	ThresholdExpression *string `pulumi:"thresholdExpression"`
}

func LookupAnomalySubscriptionOutput(ctx *pulumi.Context, args LookupAnomalySubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupAnomalySubscriptionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAnomalySubscriptionResultOutput, error) {
			args := v.(LookupAnomalySubscriptionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ce:getAnomalySubscription", args, LookupAnomalySubscriptionResultOutput{}, options).(LookupAnomalySubscriptionResultOutput), nil
		}).(LookupAnomalySubscriptionResultOutput)
}

type LookupAnomalySubscriptionOutputArgs struct {
	// The `AnomalySubscription` Amazon Resource Name (ARN).
	SubscriptionArn pulumi.StringInput `pulumi:"subscriptionArn"`
}

func (LookupAnomalySubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnomalySubscriptionArgs)(nil)).Elem()
}

type LookupAnomalySubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupAnomalySubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnomalySubscriptionResult)(nil)).Elem()
}

func (o LookupAnomalySubscriptionResultOutput) ToLookupAnomalySubscriptionResultOutput() LookupAnomalySubscriptionResultOutput {
	return o
}

func (o LookupAnomalySubscriptionResultOutput) ToLookupAnomalySubscriptionResultOutputWithContext(ctx context.Context) LookupAnomalySubscriptionResultOutput {
	return o
}

// The accountId
func (o LookupAnomalySubscriptionResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomalySubscriptionResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// The frequency at which anomaly reports are sent over email.
func (o LookupAnomalySubscriptionResultOutput) Frequency() AnomalySubscriptionFrequencyPtrOutput {
	return o.ApplyT(func(v LookupAnomalySubscriptionResult) *AnomalySubscriptionFrequency { return v.Frequency }).(AnomalySubscriptionFrequencyPtrOutput)
}

// A list of cost anomaly monitors.
func (o LookupAnomalySubscriptionResultOutput) MonitorArnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAnomalySubscriptionResult) []string { return v.MonitorArnList }).(pulumi.StringArrayOutput)
}

// A list of subscriber
func (o LookupAnomalySubscriptionResultOutput) Subscribers() AnomalySubscriptionSubscriberArrayOutput {
	return o.ApplyT(func(v LookupAnomalySubscriptionResult) []AnomalySubscriptionSubscriber { return v.Subscribers }).(AnomalySubscriptionSubscriberArrayOutput)
}

// The `AnomalySubscription` Amazon Resource Name (ARN).
func (o LookupAnomalySubscriptionResultOutput) SubscriptionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomalySubscriptionResult) *string { return v.SubscriptionArn }).(pulumi.StringPtrOutput)
}

// The name of the subscription.
func (o LookupAnomalySubscriptionResultOutput) SubscriptionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomalySubscriptionResult) *string { return v.SubscriptionName }).(pulumi.StringPtrOutput)
}

// The dollar value that triggers a notification if the threshold is exceeded.
func (o LookupAnomalySubscriptionResultOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupAnomalySubscriptionResult) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

// An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.
func (o LookupAnomalySubscriptionResultOutput) ThresholdExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomalySubscriptionResult) *string { return v.ThresholdExpression }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnomalySubscriptionResultOutput{})
}
