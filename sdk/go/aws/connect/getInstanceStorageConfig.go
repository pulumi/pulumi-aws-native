// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Connect::InstanceStorageConfig
func LookupInstanceStorageConfig(ctx *pulumi.Context, args *LookupInstanceStorageConfigArgs, opts ...pulumi.InvokeOption) (*LookupInstanceStorageConfigResult, error) {
	var rv LookupInstanceStorageConfigResult
	err := ctx.Invoke("aws-native:connect:getInstanceStorageConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceStorageConfigArgs struct {
	AssociationId string `pulumi:"associationId"`
	// Connect Instance ID with which the storage config will be associated
	InstanceArn  string                                           `pulumi:"instanceArn"`
	ResourceType InstanceStorageConfigInstanceStorageResourceType `pulumi:"resourceType"`
}

type LookupInstanceStorageConfigResult struct {
	AssociationId            *string                                        `pulumi:"associationId"`
	KinesisFirehoseConfig    *InstanceStorageConfigKinesisFirehoseConfig    `pulumi:"kinesisFirehoseConfig"`
	KinesisStreamConfig      *InstanceStorageConfigKinesisStreamConfig      `pulumi:"kinesisStreamConfig"`
	KinesisVideoStreamConfig *InstanceStorageConfigKinesisVideoStreamConfig `pulumi:"kinesisVideoStreamConfig"`
	S3Config                 *InstanceStorageConfigS3Config                 `pulumi:"s3Config"`
	StorageType              *InstanceStorageConfigStorageType              `pulumi:"storageType"`
}

func LookupInstanceStorageConfigOutput(ctx *pulumi.Context, args LookupInstanceStorageConfigOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceStorageConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceStorageConfigResult, error) {
			args := v.(LookupInstanceStorageConfigArgs)
			r, err := LookupInstanceStorageConfig(ctx, &args, opts...)
			var s LookupInstanceStorageConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceStorageConfigResultOutput)
}

type LookupInstanceStorageConfigOutputArgs struct {
	AssociationId pulumi.StringInput `pulumi:"associationId"`
	// Connect Instance ID with which the storage config will be associated
	InstanceArn  pulumi.StringInput                                    `pulumi:"instanceArn"`
	ResourceType InstanceStorageConfigInstanceStorageResourceTypeInput `pulumi:"resourceType"`
}

func (LookupInstanceStorageConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceStorageConfigArgs)(nil)).Elem()
}

type LookupInstanceStorageConfigResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceStorageConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceStorageConfigResult)(nil)).Elem()
}

func (o LookupInstanceStorageConfigResultOutput) ToLookupInstanceStorageConfigResultOutput() LookupInstanceStorageConfigResultOutput {
	return o
}

func (o LookupInstanceStorageConfigResultOutput) ToLookupInstanceStorageConfigResultOutputWithContext(ctx context.Context) LookupInstanceStorageConfigResultOutput {
	return o
}

func (o LookupInstanceStorageConfigResultOutput) AssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceStorageConfigResult) *string { return v.AssociationId }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceStorageConfigResultOutput) KinesisFirehoseConfig() InstanceStorageConfigKinesisFirehoseConfigPtrOutput {
	return o.ApplyT(func(v LookupInstanceStorageConfigResult) *InstanceStorageConfigKinesisFirehoseConfig {
		return v.KinesisFirehoseConfig
	}).(InstanceStorageConfigKinesisFirehoseConfigPtrOutput)
}

func (o LookupInstanceStorageConfigResultOutput) KinesisStreamConfig() InstanceStorageConfigKinesisStreamConfigPtrOutput {
	return o.ApplyT(func(v LookupInstanceStorageConfigResult) *InstanceStorageConfigKinesisStreamConfig {
		return v.KinesisStreamConfig
	}).(InstanceStorageConfigKinesisStreamConfigPtrOutput)
}

func (o LookupInstanceStorageConfigResultOutput) KinesisVideoStreamConfig() InstanceStorageConfigKinesisVideoStreamConfigPtrOutput {
	return o.ApplyT(func(v LookupInstanceStorageConfigResult) *InstanceStorageConfigKinesisVideoStreamConfig {
		return v.KinesisVideoStreamConfig
	}).(InstanceStorageConfigKinesisVideoStreamConfigPtrOutput)
}

func (o LookupInstanceStorageConfigResultOutput) S3Config() InstanceStorageConfigS3ConfigPtrOutput {
	return o.ApplyT(func(v LookupInstanceStorageConfigResult) *InstanceStorageConfigS3Config { return v.S3Config }).(InstanceStorageConfigS3ConfigPtrOutput)
}

func (o LookupInstanceStorageConfigResultOutput) StorageType() InstanceStorageConfigStorageTypePtrOutput {
	return o.ApplyT(func(v LookupInstanceStorageConfigResult) *InstanceStorageConfigStorageType { return v.StorageType }).(InstanceStorageConfigStorageTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceStorageConfigResultOutput{})
}