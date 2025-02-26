// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::MSK::Cluster
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("aws-native:msk:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupClusterResult struct {
	Arn                  *string                      `pulumi:"arn"`
	BrokerNodeGroupInfo  *ClusterBrokerNodeGroupInfo  `pulumi:"brokerNodeGroupInfo"`
	ClientAuthentication *ClusterClientAuthentication `pulumi:"clientAuthentication"`
	ConfigurationInfo    *ClusterConfigurationInfo    `pulumi:"configurationInfo"`
	// The current version of the MSK cluster
	CurrentVersion      *string                    `pulumi:"currentVersion"`
	EncryptionInfo      *ClusterEncryptionInfo     `pulumi:"encryptionInfo"`
	EnhancedMonitoring  *ClusterEnhancedMonitoring `pulumi:"enhancedMonitoring"`
	KafkaVersion        *string                    `pulumi:"kafkaVersion"`
	LoggingInfo         *ClusterLoggingInfo        `pulumi:"loggingInfo"`
	NumberOfBrokerNodes *int                       `pulumi:"numberOfBrokerNodes"`
	OpenMonitoring      *ClusterOpenMonitoring     `pulumi:"openMonitoring"`
	StorageMode         *ClusterStorageMode        `pulumi:"storageMode"`
	// A key-value pair to associate with a resource.
	Tags map[string]string `pulumi:"tags"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupClusterResultOutput, error) {
			args := v.(LookupClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:msk:getCluster", args, LookupClusterResultOutput{}, options).(LookupClusterResultOutput), nil
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) BrokerNodeGroupInfo() ClusterBrokerNodeGroupInfoPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterBrokerNodeGroupInfo { return v.BrokerNodeGroupInfo }).(ClusterBrokerNodeGroupInfoPtrOutput)
}

func (o LookupClusterResultOutput) ClientAuthentication() ClusterClientAuthenticationPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterClientAuthentication { return v.ClientAuthentication }).(ClusterClientAuthenticationPtrOutput)
}

func (o LookupClusterResultOutput) ConfigurationInfo() ClusterConfigurationInfoPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterConfigurationInfo { return v.ConfigurationInfo }).(ClusterConfigurationInfoPtrOutput)
}

// The current version of the MSK cluster
func (o LookupClusterResultOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) EncryptionInfo() ClusterEncryptionInfoPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterEncryptionInfo { return v.EncryptionInfo }).(ClusterEncryptionInfoPtrOutput)
}

func (o LookupClusterResultOutput) EnhancedMonitoring() ClusterEnhancedMonitoringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterEnhancedMonitoring { return v.EnhancedMonitoring }).(ClusterEnhancedMonitoringPtrOutput)
}

func (o LookupClusterResultOutput) KafkaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.KafkaVersion }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) LoggingInfo() ClusterLoggingInfoPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterLoggingInfo { return v.LoggingInfo }).(ClusterLoggingInfoPtrOutput)
}

func (o LookupClusterResultOutput) NumberOfBrokerNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.NumberOfBrokerNodes }).(pulumi.IntPtrOutput)
}

func (o LookupClusterResultOutput) OpenMonitoring() ClusterOpenMonitoringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterOpenMonitoring { return v.OpenMonitoring }).(ClusterOpenMonitoringPtrOutput)
}

func (o LookupClusterResultOutput) StorageMode() ClusterStorageModePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterStorageMode { return v.StorageMode }).(ClusterStorageModePtrOutput)
}

// A key-value pair to associate with a resource.
func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
