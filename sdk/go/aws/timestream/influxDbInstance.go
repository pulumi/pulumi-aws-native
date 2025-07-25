// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package timestream

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Timestream::InfluxDBInstance resource creates an InfluxDB instance.
type InfluxDbInstance struct {
	pulumi.CustomResourceState

	// The allocated storage for the InfluxDB instance.
	AllocatedStorage pulumi.IntPtrOutput `pulumi:"allocatedStorage"`
	// The Amazon Resource Name (ARN) that is associated with the InfluxDB instance.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Availability Zone (AZ) where the InfluxDB instance is created.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The service generated unique identifier for InfluxDB instance.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The bucket for the InfluxDB instance.
	Bucket pulumi.StringPtrOutput `pulumi:"bucket"`
	// The compute instance of the InfluxDB instance.
	DbInstanceType InfluxDbInstanceDbInstanceTypePtrOutput `pulumi:"dbInstanceType"`
	// The name of an existing InfluxDB parameter group.
	DbParameterGroupIdentifier pulumi.StringPtrOutput `pulumi:"dbParameterGroupIdentifier"`
	// The storage type of the InfluxDB instance.
	DbStorageType InfluxDbInstanceDbStorageTypePtrOutput `pulumi:"dbStorageType"`
	// Deployment type of the InfluxDB Instance.
	DeploymentType InfluxDbInstanceDeploymentTypePtrOutput `pulumi:"deploymentType"`
	// The connection endpoint for the InfluxDB instance.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The Auth parameters secret Amazon Resource name (ARN) that is associated with the InfluxDB instance.
	InfluxAuthParametersSecretArn pulumi.StringOutput `pulumi:"influxAuthParametersSecretArn"`
	// Configuration for sending logs to customer account from the InfluxDB instance.
	LogDeliveryConfiguration LogDeliveryConfigurationPropertiesPtrOutput `pulumi:"logDeliveryConfiguration"`
	// The unique name that is associated with the InfluxDB instance.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Network type of the InfluxDB Instance.
	NetworkType InfluxDbInstanceNetworkTypePtrOutput `pulumi:"networkType"`
	// The organization for the InfluxDB instance.
	Organization pulumi.StringPtrOutput `pulumi:"organization"`
	// The password for the InfluxDB instance.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The port number on which InfluxDB accepts connections.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// Attach a public IP to the customer ENI.
	PubliclyAccessible pulumi.BoolPtrOutput `pulumi:"publiclyAccessible"`
	// The Secondary Availability Zone (AZ) where the InfluxDB instance is created, if DeploymentType is set as WITH_MULTIAZ_STANDBY.
	SecondaryAvailabilityZone pulumi.StringOutput `pulumi:"secondaryAvailabilityZone"`
	// Status of the InfluxDB Instance.
	Status InfluxDbInstanceStatusOutput `pulumi:"status"`
	// An arbitrary set of tags (key-value pairs) for this DB instance.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The username for the InfluxDB instance.
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// A list of Amazon EC2 VPC security groups to associate with this InfluxDB instance.
	VpcSecurityGroupIds pulumi.StringArrayOutput `pulumi:"vpcSecurityGroupIds"`
	// A list of EC2 subnet IDs for this InfluxDB instance.
	VpcSubnetIds pulumi.StringArrayOutput `pulumi:"vpcSubnetIds"`
}

// NewInfluxDbInstance registers a new resource with the given unique name, arguments, and options.
func NewInfluxDbInstance(ctx *pulumi.Context,
	name string, args *InfluxDbInstanceArgs, opts ...pulumi.ResourceOption) (*InfluxDbInstance, error) {
	if args == nil {
		args = &InfluxDbInstanceArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bucket",
		"name",
		"networkType",
		"organization",
		"password",
		"publiclyAccessible",
		"username",
		"vpcSecurityGroupIds[*]",
		"vpcSubnetIds[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InfluxDbInstance
	err := ctx.RegisterResource("aws-native:timestream:InfluxDbInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInfluxDbInstance gets an existing InfluxDbInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInfluxDbInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InfluxDbInstanceState, opts ...pulumi.ResourceOption) (*InfluxDbInstance, error) {
	var resource InfluxDbInstance
	err := ctx.ReadResource("aws-native:timestream:InfluxDbInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InfluxDbInstance resources.
type influxDbInstanceState struct {
}

type InfluxDbInstanceState struct {
}

func (InfluxDbInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*influxDbInstanceState)(nil)).Elem()
}

type influxDbInstanceArgs struct {
	// The allocated storage for the InfluxDB instance.
	AllocatedStorage *int `pulumi:"allocatedStorage"`
	// The bucket for the InfluxDB instance.
	Bucket *string `pulumi:"bucket"`
	// The compute instance of the InfluxDB instance.
	DbInstanceType *InfluxDbInstanceDbInstanceType `pulumi:"dbInstanceType"`
	// The name of an existing InfluxDB parameter group.
	DbParameterGroupIdentifier *string `pulumi:"dbParameterGroupIdentifier"`
	// The storage type of the InfluxDB instance.
	DbStorageType *InfluxDbInstanceDbStorageType `pulumi:"dbStorageType"`
	// Deployment type of the InfluxDB Instance.
	DeploymentType *InfluxDbInstanceDeploymentType `pulumi:"deploymentType"`
	// Configuration for sending logs to customer account from the InfluxDB instance.
	LogDeliveryConfiguration *LogDeliveryConfigurationProperties `pulumi:"logDeliveryConfiguration"`
	// The unique name that is associated with the InfluxDB instance.
	Name *string `pulumi:"name"`
	// Network type of the InfluxDB Instance.
	NetworkType *InfluxDbInstanceNetworkType `pulumi:"networkType"`
	// The organization for the InfluxDB instance.
	Organization *string `pulumi:"organization"`
	// The password for the InfluxDB instance.
	Password *string `pulumi:"password"`
	// The port number on which InfluxDB accepts connections.
	Port *int `pulumi:"port"`
	// Attach a public IP to the customer ENI.
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// An arbitrary set of tags (key-value pairs) for this DB instance.
	Tags []aws.Tag `pulumi:"tags"`
	// The username for the InfluxDB instance.
	Username *string `pulumi:"username"`
	// A list of Amazon EC2 VPC security groups to associate with this InfluxDB instance.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
	// A list of EC2 subnet IDs for this InfluxDB instance.
	VpcSubnetIds []string `pulumi:"vpcSubnetIds"`
}

// The set of arguments for constructing a InfluxDbInstance resource.
type InfluxDbInstanceArgs struct {
	// The allocated storage for the InfluxDB instance.
	AllocatedStorage pulumi.IntPtrInput
	// The bucket for the InfluxDB instance.
	Bucket pulumi.StringPtrInput
	// The compute instance of the InfluxDB instance.
	DbInstanceType InfluxDbInstanceDbInstanceTypePtrInput
	// The name of an existing InfluxDB parameter group.
	DbParameterGroupIdentifier pulumi.StringPtrInput
	// The storage type of the InfluxDB instance.
	DbStorageType InfluxDbInstanceDbStorageTypePtrInput
	// Deployment type of the InfluxDB Instance.
	DeploymentType InfluxDbInstanceDeploymentTypePtrInput
	// Configuration for sending logs to customer account from the InfluxDB instance.
	LogDeliveryConfiguration LogDeliveryConfigurationPropertiesPtrInput
	// The unique name that is associated with the InfluxDB instance.
	Name pulumi.StringPtrInput
	// Network type of the InfluxDB Instance.
	NetworkType InfluxDbInstanceNetworkTypePtrInput
	// The organization for the InfluxDB instance.
	Organization pulumi.StringPtrInput
	// The password for the InfluxDB instance.
	Password pulumi.StringPtrInput
	// The port number on which InfluxDB accepts connections.
	Port pulumi.IntPtrInput
	// Attach a public IP to the customer ENI.
	PubliclyAccessible pulumi.BoolPtrInput
	// An arbitrary set of tags (key-value pairs) for this DB instance.
	Tags aws.TagArrayInput
	// The username for the InfluxDB instance.
	Username pulumi.StringPtrInput
	// A list of Amazon EC2 VPC security groups to associate with this InfluxDB instance.
	VpcSecurityGroupIds pulumi.StringArrayInput
	// A list of EC2 subnet IDs for this InfluxDB instance.
	VpcSubnetIds pulumi.StringArrayInput
}

func (InfluxDbInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*influxDbInstanceArgs)(nil)).Elem()
}

type InfluxDbInstanceInput interface {
	pulumi.Input

	ToInfluxDbInstanceOutput() InfluxDbInstanceOutput
	ToInfluxDbInstanceOutputWithContext(ctx context.Context) InfluxDbInstanceOutput
}

func (*InfluxDbInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**InfluxDbInstance)(nil)).Elem()
}

func (i *InfluxDbInstance) ToInfluxDbInstanceOutput() InfluxDbInstanceOutput {
	return i.ToInfluxDbInstanceOutputWithContext(context.Background())
}

func (i *InfluxDbInstance) ToInfluxDbInstanceOutputWithContext(ctx context.Context) InfluxDbInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InfluxDbInstanceOutput)
}

type InfluxDbInstanceOutput struct{ *pulumi.OutputState }

func (InfluxDbInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InfluxDbInstance)(nil)).Elem()
}

func (o InfluxDbInstanceOutput) ToInfluxDbInstanceOutput() InfluxDbInstanceOutput {
	return o
}

func (o InfluxDbInstanceOutput) ToInfluxDbInstanceOutputWithContext(ctx context.Context) InfluxDbInstanceOutput {
	return o
}

// The allocated storage for the InfluxDB instance.
func (o InfluxDbInstanceOutput) AllocatedStorage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.IntPtrOutput { return v.AllocatedStorage }).(pulumi.IntPtrOutput)
}

// The Amazon Resource Name (ARN) that is associated with the InfluxDB instance.
func (o InfluxDbInstanceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Availability Zone (AZ) where the InfluxDB instance is created.
func (o InfluxDbInstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The service generated unique identifier for InfluxDB instance.
func (o InfluxDbInstanceOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The bucket for the InfluxDB instance.
func (o InfluxDbInstanceOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringPtrOutput { return v.Bucket }).(pulumi.StringPtrOutput)
}

// The compute instance of the InfluxDB instance.
func (o InfluxDbInstanceOutput) DbInstanceType() InfluxDbInstanceDbInstanceTypePtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) InfluxDbInstanceDbInstanceTypePtrOutput { return v.DbInstanceType }).(InfluxDbInstanceDbInstanceTypePtrOutput)
}

// The name of an existing InfluxDB parameter group.
func (o InfluxDbInstanceOutput) DbParameterGroupIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringPtrOutput { return v.DbParameterGroupIdentifier }).(pulumi.StringPtrOutput)
}

// The storage type of the InfluxDB instance.
func (o InfluxDbInstanceOutput) DbStorageType() InfluxDbInstanceDbStorageTypePtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) InfluxDbInstanceDbStorageTypePtrOutput { return v.DbStorageType }).(InfluxDbInstanceDbStorageTypePtrOutput)
}

// Deployment type of the InfluxDB Instance.
func (o InfluxDbInstanceOutput) DeploymentType() InfluxDbInstanceDeploymentTypePtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) InfluxDbInstanceDeploymentTypePtrOutput { return v.DeploymentType }).(InfluxDbInstanceDeploymentTypePtrOutput)
}

// The connection endpoint for the InfluxDB instance.
func (o InfluxDbInstanceOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The Auth parameters secret Amazon Resource name (ARN) that is associated with the InfluxDB instance.
func (o InfluxDbInstanceOutput) InfluxAuthParametersSecretArn() pulumi.StringOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringOutput { return v.InfluxAuthParametersSecretArn }).(pulumi.StringOutput)
}

// Configuration for sending logs to customer account from the InfluxDB instance.
func (o InfluxDbInstanceOutput) LogDeliveryConfiguration() LogDeliveryConfigurationPropertiesPtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) LogDeliveryConfigurationPropertiesPtrOutput {
		return v.LogDeliveryConfiguration
	}).(LogDeliveryConfigurationPropertiesPtrOutput)
}

// The unique name that is associated with the InfluxDB instance.
func (o InfluxDbInstanceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Network type of the InfluxDB Instance.
func (o InfluxDbInstanceOutput) NetworkType() InfluxDbInstanceNetworkTypePtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) InfluxDbInstanceNetworkTypePtrOutput { return v.NetworkType }).(InfluxDbInstanceNetworkTypePtrOutput)
}

// The organization for the InfluxDB instance.
func (o InfluxDbInstanceOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringPtrOutput { return v.Organization }).(pulumi.StringPtrOutput)
}

// The password for the InfluxDB instance.
func (o InfluxDbInstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The port number on which InfluxDB accepts connections.
func (o InfluxDbInstanceOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// Attach a public IP to the customer ENI.
func (o InfluxDbInstanceOutput) PubliclyAccessible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.BoolPtrOutput { return v.PubliclyAccessible }).(pulumi.BoolPtrOutput)
}

// The Secondary Availability Zone (AZ) where the InfluxDB instance is created, if DeploymentType is set as WITH_MULTIAZ_STANDBY.
func (o InfluxDbInstanceOutput) SecondaryAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringOutput { return v.SecondaryAvailabilityZone }).(pulumi.StringOutput)
}

// Status of the InfluxDB Instance.
func (o InfluxDbInstanceOutput) Status() InfluxDbInstanceStatusOutput {
	return o.ApplyT(func(v *InfluxDbInstance) InfluxDbInstanceStatusOutput { return v.Status }).(InfluxDbInstanceStatusOutput)
}

// An arbitrary set of tags (key-value pairs) for this DB instance.
func (o InfluxDbInstanceOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *InfluxDbInstance) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The username for the InfluxDB instance.
func (o InfluxDbInstanceOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

// A list of Amazon EC2 VPC security groups to associate with this InfluxDB instance.
func (o InfluxDbInstanceOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringArrayOutput { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

// A list of EC2 subnet IDs for this InfluxDB instance.
func (o InfluxDbInstanceOutput) VpcSubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InfluxDbInstance) pulumi.StringArrayOutput { return v.VpcSubnetIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InfluxDbInstanceInput)(nil)).Elem(), &InfluxDbInstance{})
	pulumi.RegisterOutputType(InfluxDbInstanceOutput{})
}
