// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::NetworkFirewall::LoggingConfiguration
func LookupLoggingConfiguration(ctx *pulumi.Context, args *LookupLoggingConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupLoggingConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLoggingConfigurationResult
	err := ctx.Invoke("aws-native:networkfirewall:getLoggingConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoggingConfigurationArgs struct {
	// The Amazon Resource Name (ARN) of the `Firewall` that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
	FirewallArn string `pulumi:"firewallArn"`
}

type LookupLoggingConfigurationResult struct {
	EnableMonitoringDashboard *bool `pulumi:"enableMonitoringDashboard"`
	// Defines how AWS Network Firewall performs logging for a `Firewall` .
	LoggingConfiguration *LoggingConfigurationType `pulumi:"loggingConfiguration"`
}

func LookupLoggingConfigurationOutput(ctx *pulumi.Context, args LookupLoggingConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupLoggingConfigurationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLoggingConfigurationResultOutput, error) {
			args := v.(LookupLoggingConfigurationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:networkfirewall:getLoggingConfiguration", args, LookupLoggingConfigurationResultOutput{}, options).(LookupLoggingConfigurationResultOutput), nil
		}).(LookupLoggingConfigurationResultOutput)
}

type LookupLoggingConfigurationOutputArgs struct {
	// The Amazon Resource Name (ARN) of the `Firewall` that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
	FirewallArn pulumi.StringInput `pulumi:"firewallArn"`
}

func (LookupLoggingConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggingConfigurationArgs)(nil)).Elem()
}

type LookupLoggingConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupLoggingConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggingConfigurationResult)(nil)).Elem()
}

func (o LookupLoggingConfigurationResultOutput) ToLookupLoggingConfigurationResultOutput() LookupLoggingConfigurationResultOutput {
	return o
}

func (o LookupLoggingConfigurationResultOutput) ToLookupLoggingConfigurationResultOutputWithContext(ctx context.Context) LookupLoggingConfigurationResultOutput {
	return o
}

func (o LookupLoggingConfigurationResultOutput) EnableMonitoringDashboard() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLoggingConfigurationResult) *bool { return v.EnableMonitoringDashboard }).(pulumi.BoolPtrOutput)
}

// Defines how AWS Network Firewall performs logging for a `Firewall` .
func (o LookupLoggingConfigurationResultOutput) LoggingConfiguration() LoggingConfigurationTypePtrOutput {
	return o.ApplyT(func(v LookupLoggingConfigurationResult) *LoggingConfigurationType { return v.LoggingConfiguration }).(LoggingConfigurationTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoggingConfigurationResultOutput{})
}
