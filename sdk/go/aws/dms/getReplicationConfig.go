// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A replication configuration that you later provide to configure and start a AWS DMS Serverless replication
func LookupReplicationConfig(ctx *pulumi.Context, args *LookupReplicationConfigArgs, opts ...pulumi.InvokeOption) (*LookupReplicationConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationConfigResult
	err := ctx.Invoke("aws-native:dms:getReplicationConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationConfigArgs struct {
	// The Amazon Resource Name (ARN) of the Replication Config
	ReplicationConfigArn string `pulumi:"replicationConfigArn"`
}

type LookupReplicationConfigResult struct {
	// Configuration parameters for provisioning an AWS DMS Serverless replication.
	ComputeConfig *ReplicationConfigComputeConfig `pulumi:"computeConfig"`
	// The Amazon Resource Name (ARN) of the Replication Config
	ReplicationConfigArn *string `pulumi:"replicationConfigArn"`
	// A unique identifier of replication configuration
	ReplicationConfigIdentifier *string `pulumi:"replicationConfigIdentifier"`
	// JSON settings for Servereless replications that are provisioned using this replication configuration
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::DMS::ReplicationConfig` for more information about the expected schema for this property.
	ReplicationSettings interface{} `pulumi:"replicationSettings"`
	// The type of AWS DMS Serverless replication to provision using this replication configuration
	ReplicationType *ReplicationConfigReplicationType `pulumi:"replicationType"`
	// The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration
	SourceEndpointArn *string `pulumi:"sourceEndpointArn"`
	// JSON settings for specifying supplemental data
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::DMS::ReplicationConfig` for more information about the expected schema for this property.
	SupplementalSettings interface{} `pulumi:"supplementalSettings"`
	// JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::DMS::ReplicationConfig` for more information about the expected schema for this property.
	TableMappings interface{} `pulumi:"tableMappings"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>
	Tags []aws.Tag `pulumi:"tags"`
	// The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration
	TargetEndpointArn *string `pulumi:"targetEndpointArn"`
}

func LookupReplicationConfigOutput(ctx *pulumi.Context, args LookupReplicationConfigOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationConfigResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupReplicationConfigResultOutput, error) {
			args := v.(LookupReplicationConfigArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:dms:getReplicationConfig", args, LookupReplicationConfigResultOutput{}, options).(LookupReplicationConfigResultOutput), nil
		}).(LookupReplicationConfigResultOutput)
}

type LookupReplicationConfigOutputArgs struct {
	// The Amazon Resource Name (ARN) of the Replication Config
	ReplicationConfigArn pulumi.StringInput `pulumi:"replicationConfigArn"`
}

func (LookupReplicationConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationConfigArgs)(nil)).Elem()
}

type LookupReplicationConfigResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationConfigResult)(nil)).Elem()
}

func (o LookupReplicationConfigResultOutput) ToLookupReplicationConfigResultOutput() LookupReplicationConfigResultOutput {
	return o
}

func (o LookupReplicationConfigResultOutput) ToLookupReplicationConfigResultOutputWithContext(ctx context.Context) LookupReplicationConfigResultOutput {
	return o
}

// Configuration parameters for provisioning an AWS DMS Serverless replication.
func (o LookupReplicationConfigResultOutput) ComputeConfig() ReplicationConfigComputeConfigPtrOutput {
	return o.ApplyT(func(v LookupReplicationConfigResult) *ReplicationConfigComputeConfig { return v.ComputeConfig }).(ReplicationConfigComputeConfigPtrOutput)
}

// The Amazon Resource Name (ARN) of the Replication Config
func (o LookupReplicationConfigResultOutput) ReplicationConfigArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationConfigResult) *string { return v.ReplicationConfigArn }).(pulumi.StringPtrOutput)
}

// A unique identifier of replication configuration
func (o LookupReplicationConfigResultOutput) ReplicationConfigIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationConfigResult) *string { return v.ReplicationConfigIdentifier }).(pulumi.StringPtrOutput)
}

// JSON settings for Servereless replications that are provisioned using this replication configuration
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::DMS::ReplicationConfig` for more information about the expected schema for this property.
func (o LookupReplicationConfigResultOutput) ReplicationSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupReplicationConfigResult) interface{} { return v.ReplicationSettings }).(pulumi.AnyOutput)
}

// The type of AWS DMS Serverless replication to provision using this replication configuration
func (o LookupReplicationConfigResultOutput) ReplicationType() ReplicationConfigReplicationTypePtrOutput {
	return o.ApplyT(func(v LookupReplicationConfigResult) *ReplicationConfigReplicationType { return v.ReplicationType }).(ReplicationConfigReplicationTypePtrOutput)
}

// The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration
func (o LookupReplicationConfigResultOutput) SourceEndpointArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationConfigResult) *string { return v.SourceEndpointArn }).(pulumi.StringPtrOutput)
}

// JSON settings for specifying supplemental data
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::DMS::ReplicationConfig` for more information about the expected schema for this property.
func (o LookupReplicationConfigResultOutput) SupplementalSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupReplicationConfigResult) interface{} { return v.SupplementalSettings }).(pulumi.AnyOutput)
}

// JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::DMS::ReplicationConfig` for more information about the expected schema for this property.
func (o LookupReplicationConfigResultOutput) TableMappings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupReplicationConfigResult) interface{} { return v.TableMappings }).(pulumi.AnyOutput)
}

// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>
func (o LookupReplicationConfigResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupReplicationConfigResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration
func (o LookupReplicationConfigResultOutput) TargetEndpointArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationConfigResult) *string { return v.TargetEndpointArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationConfigResultOutput{})
}
