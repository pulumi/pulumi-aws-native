// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apprunner

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::AppRunner::Service resource specifies an AppRunner Service.
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceResult
	err := ctx.Invoke("aws-native:apprunner:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	// The Amazon Resource Name (ARN) of the AppRunner Service.
	ServiceArn string `pulumi:"serviceArn"`
}

type LookupServiceResult struct {
	// The settings for the health check that AWS App Runner performs to monitor the health of the App Runner service.
	HealthCheckConfiguration *ServiceHealthCheckConfiguration `pulumi:"healthCheckConfiguration"`
	// The runtime configuration of instances (scaling units) of your service.
	InstanceConfiguration *ServiceInstanceConfiguration `pulumi:"instanceConfiguration"`
	// Configuration settings related to network traffic of the web application that the App Runner service runs.
	NetworkConfiguration *ServiceNetworkConfiguration `pulumi:"networkConfiguration"`
	// The observability configuration of your service.
	ObservabilityConfiguration *ServiceObservabilityConfiguration `pulumi:"observabilityConfiguration"`
	// The Amazon Resource Name (ARN) of the AppRunner Service.
	ServiceArn *string `pulumi:"serviceArn"`
	// The AppRunner Service Id
	ServiceId *string `pulumi:"serviceId"`
	// The Service Url of the AppRunner Service.
	ServiceUrl *string `pulumi:"serviceUrl"`
	// The source to deploy to the App Runner service. It can be a code or an image repository.
	SourceConfiguration *ServiceSourceConfiguration `pulumi:"sourceConfiguration"`
	// AppRunner Service status.
	Status *string `pulumi:"status"`
	// An optional list of metadata items that you can associate with the App Runner service resource. A tag is a key-value pair.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupServiceResultOutput, error) {
			args := v.(LookupServiceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:apprunner:getService", args, LookupServiceResultOutput{}, options).(LookupServiceResultOutput), nil
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	// The Amazon Resource Name (ARN) of the AppRunner Service.
	ServiceArn pulumi.StringInput `pulumi:"serviceArn"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

// The settings for the health check that AWS App Runner performs to monitor the health of the App Runner service.
func (o LookupServiceResultOutput) HealthCheckConfiguration() ServiceHealthCheckConfigurationPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceHealthCheckConfiguration { return v.HealthCheckConfiguration }).(ServiceHealthCheckConfigurationPtrOutput)
}

// The runtime configuration of instances (scaling units) of your service.
func (o LookupServiceResultOutput) InstanceConfiguration() ServiceInstanceConfigurationPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceInstanceConfiguration { return v.InstanceConfiguration }).(ServiceInstanceConfigurationPtrOutput)
}

// Configuration settings related to network traffic of the web application that the App Runner service runs.
func (o LookupServiceResultOutput) NetworkConfiguration() ServiceNetworkConfigurationPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceNetworkConfiguration { return v.NetworkConfiguration }).(ServiceNetworkConfigurationPtrOutput)
}

// The observability configuration of your service.
func (o LookupServiceResultOutput) ObservabilityConfiguration() ServiceObservabilityConfigurationPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceObservabilityConfiguration { return v.ObservabilityConfiguration }).(ServiceObservabilityConfigurationPtrOutput)
}

// The Amazon Resource Name (ARN) of the AppRunner Service.
func (o LookupServiceResultOutput) ServiceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.ServiceArn }).(pulumi.StringPtrOutput)
}

// The AppRunner Service Id
func (o LookupServiceResultOutput) ServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.ServiceId }).(pulumi.StringPtrOutput)
}

// The Service Url of the AppRunner Service.
func (o LookupServiceResultOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

// The source to deploy to the App Runner service. It can be a code or an image repository.
func (o LookupServiceResultOutput) SourceConfiguration() ServiceSourceConfigurationPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceSourceConfiguration { return v.SourceConfiguration }).(ServiceSourceConfigurationPtrOutput)
}

// AppRunner Service status.
func (o LookupServiceResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// An optional list of metadata items that you can associate with the App Runner service resource. A tag is a key-value pair.
func (o LookupServiceResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
