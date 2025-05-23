// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::Instance
type Instance struct {
	pulumi.CustomResourceState

	// This property is reserved for internal use. If you use it, the stack fails with this error: Bad property set: [Testing this property] (Service: AmazonEC2; Status Code: 400; Error Code: InvalidParameterCombination; Request ID: 0XXXXXX-49c7-4b40-8bcc-76885dcXXXXX).
	AdditionalInfo pulumi.StringPtrOutput `pulumi:"additionalInfo"`
	// Indicates whether the instance is associated with a dedicated host. If you want the instance to always restart on the same host on which it was launched, specify host. If you want the instance to restart on any available host, but try to launch onto the last host it ran on (on a best-effort basis), specify default.
	Affinity InstanceAffinityPtrOutput `pulumi:"affinity"`
	// The Availability Zone of the instance.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// The block device mapping entries that defines the block devices to attach to the instance at launch.
	BlockDeviceMappings InstanceBlockDeviceMappingArrayOutput `pulumi:"blockDeviceMappings"`
	// The CPU options for the instance.
	CpuOptions CpuOptionsPropertiesPtrOutput `pulumi:"cpuOptions"`
	// The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited.
	CreditSpecification CreditSpecificationPropertiesPtrOutput `pulumi:"creditSpecification"`
	// If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can.
	DisableApiTermination pulumi.BoolPtrOutput `pulumi:"disableApiTermination"`
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	EbsOptimized pulumi.BoolPtrOutput `pulumi:"ebsOptimized"`
	// An elastic GPU to associate with the instance.
	ElasticGpuSpecifications InstanceElasticGpuSpecificationArrayOutput `pulumi:"elasticGpuSpecifications"`
	// An elastic inference accelerator to associate with the instance.
	ElasticInferenceAccelerators InstanceElasticInferenceAcceleratorArrayOutput `pulumi:"elasticInferenceAccelerators"`
	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	EnclaveOptions EnclaveOptionsPropertiesPtrOutput `pulumi:"enclaveOptions"`
	// Indicates whether an instance is enabled for hibernation.
	HibernationOptions HibernationOptionsPropertiesPtrOutput `pulumi:"hibernationOptions"`
	// If you specify host for the Affinity property, the ID of a dedicated host that the instance is associated with. If you don't specify an ID, Amazon EC2 launches the instance onto any available, compatible dedicated host in your account.
	HostId pulumi.StringPtrOutput `pulumi:"hostId"`
	// The ARN of the host resource group in which to launch the instances. If you specify a host resource group ARN, omit the Tenancy parameter or set it to host.
	HostResourceGroupArn pulumi.StringPtrOutput `pulumi:"hostResourceGroupArn"`
	// The IAM instance profile.
	IamInstanceProfile pulumi.StringPtrOutput `pulumi:"iamInstanceProfile"`
	// The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template.
	ImageId pulumi.StringPtrOutput `pulumi:"imageId"`
	// The EC2 Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	InstanceInitiatedShutdownBehavior pulumi.StringPtrOutput `pulumi:"instanceInitiatedShutdownBehavior"`
	// The instance type.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// [EC2-VPC] The number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	Ipv6AddressCount pulumi.IntPtrOutput `pulumi:"ipv6AddressCount"`
	// [EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface.
	Ipv6Addresses InstanceIpv6AddressArrayOutput `pulumi:"ipv6Addresses"`
	// The ID of the kernel.
	KernelId pulumi.StringPtrOutput `pulumi:"kernelId"`
	// The name of the key pair.
	KeyName pulumi.StringPtrOutput `pulumi:"keyName"`
	// The launch template to use to launch the instances.
	LaunchTemplate InstanceLaunchTemplateSpecificationPtrOutput `pulumi:"launchTemplate"`
	// The license configurations.
	LicenseSpecifications InstanceLicenseSpecificationArrayOutput `pulumi:"licenseSpecifications"`
	// The metadata options for the instance
	MetadataOptions InstanceMetadataOptionsPtrOutput `pulumi:"metadataOptions"`
	// Specifies whether detailed monitoring is enabled for the instance.
	Monitoring pulumi.BoolPtrOutput `pulumi:"monitoring"`
	// The network interfaces to associate with the instance.
	NetworkInterfaces InstanceNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// The name of an existing placement group that you want to launch the instance into (cluster | partition | spread).
	PlacementGroupName pulumi.StringPtrOutput `pulumi:"placementGroupName"`
	// The private DNS name of the specified instance. For example: ip-10-24-34-0.ec2.internal.
	PrivateDnsName pulumi.StringOutput `pulumi:"privateDnsName"`
	// The options for the instance hostname.
	PrivateDnsNameOptions InstancePrivateDnsNameOptionsPtrOutput `pulumi:"privateDnsNameOptions"`
	// The private IP address of the specified instance. For example: 10.24.34.0.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// [EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet.
	PrivateIpAddress pulumi.StringPtrOutput `pulumi:"privateIpAddress"`
	// Indicates whether to assign the tags from the instance to all of the volumes attached to the instance at launch. If you specify true and you assign tags to the instance, those tags are automatically assigned to all of the volumes that you attach to the instance at launch. If you specify false, those tags are not assigned to the attached volumes.
	PropagateTagsToVolumeOnCreation pulumi.BoolPtrOutput `pulumi:"propagateTagsToVolumeOnCreation"`
	// The public DNS name of the specified instance. For example: ec2-107-20-50-45.compute-1.amazonaws.com.
	PublicDnsName pulumi.StringOutput `pulumi:"publicDnsName"`
	// The public IP address of the specified instance. For example: 192.0.2.0.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
	// The ID of the RAM disk to select.
	RamdiskId pulumi.StringPtrOutput `pulumi:"ramdiskId"`
	// The IDs of the security groups.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// the names of the security groups. For a nondefault VPC, you must use security group IDs instead.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Specifies whether to enable an instance launched in a VPC to perform NAT.
	SourceDestCheck pulumi.BoolPtrOutput `pulumi:"sourceDestCheck"`
	// The SSM document and parameter values in AWS Systems Manager to associate with this instance.
	SsmAssociations InstanceSsmAssociationArrayOutput `pulumi:"ssmAssociations"`
	// The current state of the instance.
	State InstanceStateTypeOutput `pulumi:"state"`
	// [EC2-VPC] The ID of the subnet to launch the instance into.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// The tags to add to the instance.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware.
	Tenancy pulumi.StringPtrOutput `pulumi:"tenancy"`
	// The user data to make available to the instance.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// The volumes to attach to the instance.
	Volumes InstanceVolumeArrayOutput `pulumi:"volumes"`
	// The ID of the VPC that the instance is running in.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"availabilityZone",
		"cpuOptions",
		"elasticGpuSpecifications[*]",
		"elasticInferenceAccelerators[*]",
		"enclaveOptions",
		"hibernationOptions",
		"hostResourceGroupArn",
		"imageId",
		"ipv6AddressCount",
		"ipv6Addresses[*]",
		"keyName",
		"launchTemplate",
		"licenseSpecifications[*]",
		"networkInterfaces[*]",
		"placementGroupName",
		"privateIpAddress",
		"securityGroups[*]",
		"subnetId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("aws-native:ec2:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("aws-native:ec2:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
}

type InstanceState struct {
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// This property is reserved for internal use. If you use it, the stack fails with this error: Bad property set: [Testing this property] (Service: AmazonEC2; Status Code: 400; Error Code: InvalidParameterCombination; Request ID: 0XXXXXX-49c7-4b40-8bcc-76885dcXXXXX).
	AdditionalInfo *string `pulumi:"additionalInfo"`
	// Indicates whether the instance is associated with a dedicated host. If you want the instance to always restart on the same host on which it was launched, specify host. If you want the instance to restart on any available host, but try to launch onto the last host it ran on (on a best-effort basis), specify default.
	Affinity *InstanceAffinity `pulumi:"affinity"`
	// The Availability Zone of the instance.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The block device mapping entries that defines the block devices to attach to the instance at launch.
	BlockDeviceMappings []InstanceBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// The CPU options for the instance.
	CpuOptions *CpuOptionsProperties `pulumi:"cpuOptions"`
	// The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited.
	CreditSpecification *CreditSpecificationProperties `pulumi:"creditSpecification"`
	// If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can.
	DisableApiTermination *bool `pulumi:"disableApiTermination"`
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	EbsOptimized *bool `pulumi:"ebsOptimized"`
	// An elastic GPU to associate with the instance.
	ElasticGpuSpecifications []InstanceElasticGpuSpecification `pulumi:"elasticGpuSpecifications"`
	// An elastic inference accelerator to associate with the instance.
	ElasticInferenceAccelerators []InstanceElasticInferenceAccelerator `pulumi:"elasticInferenceAccelerators"`
	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	EnclaveOptions *EnclaveOptionsProperties `pulumi:"enclaveOptions"`
	// Indicates whether an instance is enabled for hibernation.
	HibernationOptions *HibernationOptionsProperties `pulumi:"hibernationOptions"`
	// If you specify host for the Affinity property, the ID of a dedicated host that the instance is associated with. If you don't specify an ID, Amazon EC2 launches the instance onto any available, compatible dedicated host in your account.
	HostId *string `pulumi:"hostId"`
	// The ARN of the host resource group in which to launch the instances. If you specify a host resource group ARN, omit the Tenancy parameter or set it to host.
	HostResourceGroupArn *string `pulumi:"hostResourceGroupArn"`
	// The IAM instance profile.
	IamInstanceProfile *string `pulumi:"iamInstanceProfile"`
	// The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template.
	ImageId *string `pulumi:"imageId"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	InstanceInitiatedShutdownBehavior *string `pulumi:"instanceInitiatedShutdownBehavior"`
	// The instance type.
	InstanceType *string `pulumi:"instanceType"`
	// [EC2-VPC] The number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	Ipv6AddressCount *int `pulumi:"ipv6AddressCount"`
	// [EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface.
	Ipv6Addresses []InstanceIpv6Address `pulumi:"ipv6Addresses"`
	// The ID of the kernel.
	KernelId *string `pulumi:"kernelId"`
	// The name of the key pair.
	KeyName *string `pulumi:"keyName"`
	// The launch template to use to launch the instances.
	LaunchTemplate *InstanceLaunchTemplateSpecification `pulumi:"launchTemplate"`
	// The license configurations.
	LicenseSpecifications []InstanceLicenseSpecification `pulumi:"licenseSpecifications"`
	// The metadata options for the instance
	MetadataOptions *InstanceMetadataOptions `pulumi:"metadataOptions"`
	// Specifies whether detailed monitoring is enabled for the instance.
	Monitoring *bool `pulumi:"monitoring"`
	// The network interfaces to associate with the instance.
	NetworkInterfaces []InstanceNetworkInterface `pulumi:"networkInterfaces"`
	// The name of an existing placement group that you want to launch the instance into (cluster | partition | spread).
	PlacementGroupName *string `pulumi:"placementGroupName"`
	// The options for the instance hostname.
	PrivateDnsNameOptions *InstancePrivateDnsNameOptions `pulumi:"privateDnsNameOptions"`
	// [EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// Indicates whether to assign the tags from the instance to all of the volumes attached to the instance at launch. If you specify true and you assign tags to the instance, those tags are automatically assigned to all of the volumes that you attach to the instance at launch. If you specify false, those tags are not assigned to the attached volumes.
	PropagateTagsToVolumeOnCreation *bool `pulumi:"propagateTagsToVolumeOnCreation"`
	// The ID of the RAM disk to select.
	RamdiskId *string `pulumi:"ramdiskId"`
	// The IDs of the security groups.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// the names of the security groups. For a nondefault VPC, you must use security group IDs instead.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Specifies whether to enable an instance launched in a VPC to perform NAT.
	SourceDestCheck *bool `pulumi:"sourceDestCheck"`
	// The SSM document and parameter values in AWS Systems Manager to associate with this instance.
	SsmAssociations []InstanceSsmAssociation `pulumi:"ssmAssociations"`
	// [EC2-VPC] The ID of the subnet to launch the instance into.
	SubnetId *string `pulumi:"subnetId"`
	// The tags to add to the instance.
	Tags []aws.Tag `pulumi:"tags"`
	// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware.
	Tenancy *string `pulumi:"tenancy"`
	// The user data to make available to the instance.
	UserData *string `pulumi:"userData"`
	// The volumes to attach to the instance.
	Volumes []InstanceVolume `pulumi:"volumes"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// This property is reserved for internal use. If you use it, the stack fails with this error: Bad property set: [Testing this property] (Service: AmazonEC2; Status Code: 400; Error Code: InvalidParameterCombination; Request ID: 0XXXXXX-49c7-4b40-8bcc-76885dcXXXXX).
	AdditionalInfo pulumi.StringPtrInput
	// Indicates whether the instance is associated with a dedicated host. If you want the instance to always restart on the same host on which it was launched, specify host. If you want the instance to restart on any available host, but try to launch onto the last host it ran on (on a best-effort basis), specify default.
	Affinity InstanceAffinityPtrInput
	// The Availability Zone of the instance.
	AvailabilityZone pulumi.StringPtrInput
	// The block device mapping entries that defines the block devices to attach to the instance at launch.
	BlockDeviceMappings InstanceBlockDeviceMappingArrayInput
	// The CPU options for the instance.
	CpuOptions CpuOptionsPropertiesPtrInput
	// The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited.
	CreditSpecification CreditSpecificationPropertiesPtrInput
	// If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can.
	DisableApiTermination pulumi.BoolPtrInput
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	EbsOptimized pulumi.BoolPtrInput
	// An elastic GPU to associate with the instance.
	ElasticGpuSpecifications InstanceElasticGpuSpecificationArrayInput
	// An elastic inference accelerator to associate with the instance.
	ElasticInferenceAccelerators InstanceElasticInferenceAcceleratorArrayInput
	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	EnclaveOptions EnclaveOptionsPropertiesPtrInput
	// Indicates whether an instance is enabled for hibernation.
	HibernationOptions HibernationOptionsPropertiesPtrInput
	// If you specify host for the Affinity property, the ID of a dedicated host that the instance is associated with. If you don't specify an ID, Amazon EC2 launches the instance onto any available, compatible dedicated host in your account.
	HostId pulumi.StringPtrInput
	// The ARN of the host resource group in which to launch the instances. If you specify a host resource group ARN, omit the Tenancy parameter or set it to host.
	HostResourceGroupArn pulumi.StringPtrInput
	// The IAM instance profile.
	IamInstanceProfile pulumi.StringPtrInput
	// The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template.
	ImageId pulumi.StringPtrInput
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	InstanceInitiatedShutdownBehavior pulumi.StringPtrInput
	// The instance type.
	InstanceType pulumi.StringPtrInput
	// [EC2-VPC] The number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	Ipv6AddressCount pulumi.IntPtrInput
	// [EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface.
	Ipv6Addresses InstanceIpv6AddressArrayInput
	// The ID of the kernel.
	KernelId pulumi.StringPtrInput
	// The name of the key pair.
	KeyName pulumi.StringPtrInput
	// The launch template to use to launch the instances.
	LaunchTemplate InstanceLaunchTemplateSpecificationPtrInput
	// The license configurations.
	LicenseSpecifications InstanceLicenseSpecificationArrayInput
	// The metadata options for the instance
	MetadataOptions InstanceMetadataOptionsPtrInput
	// Specifies whether detailed monitoring is enabled for the instance.
	Monitoring pulumi.BoolPtrInput
	// The network interfaces to associate with the instance.
	NetworkInterfaces InstanceNetworkInterfaceArrayInput
	// The name of an existing placement group that you want to launch the instance into (cluster | partition | spread).
	PlacementGroupName pulumi.StringPtrInput
	// The options for the instance hostname.
	PrivateDnsNameOptions InstancePrivateDnsNameOptionsPtrInput
	// [EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet.
	PrivateIpAddress pulumi.StringPtrInput
	// Indicates whether to assign the tags from the instance to all of the volumes attached to the instance at launch. If you specify true and you assign tags to the instance, those tags are automatically assigned to all of the volumes that you attach to the instance at launch. If you specify false, those tags are not assigned to the attached volumes.
	PropagateTagsToVolumeOnCreation pulumi.BoolPtrInput
	// The ID of the RAM disk to select.
	RamdiskId pulumi.StringPtrInput
	// The IDs of the security groups.
	SecurityGroupIds pulumi.StringArrayInput
	// the names of the security groups. For a nondefault VPC, you must use security group IDs instead.
	SecurityGroups pulumi.StringArrayInput
	// Specifies whether to enable an instance launched in a VPC to perform NAT.
	SourceDestCheck pulumi.BoolPtrInput
	// The SSM document and parameter values in AWS Systems Manager to associate with this instance.
	SsmAssociations InstanceSsmAssociationArrayInput
	// [EC2-VPC] The ID of the subnet to launch the instance into.
	SubnetId pulumi.StringPtrInput
	// The tags to add to the instance.
	Tags aws.TagArrayInput
	// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware.
	Tenancy pulumi.StringPtrInput
	// The user data to make available to the instance.
	UserData pulumi.StringPtrInput
	// The volumes to attach to the instance.
	Volumes InstanceVolumeArrayInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// This property is reserved for internal use. If you use it, the stack fails with this error: Bad property set: [Testing this property] (Service: AmazonEC2; Status Code: 400; Error Code: InvalidParameterCombination; Request ID: 0XXXXXX-49c7-4b40-8bcc-76885dcXXXXX).
func (o InstanceOutput) AdditionalInfo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.AdditionalInfo }).(pulumi.StringPtrOutput)
}

// Indicates whether the instance is associated with a dedicated host. If you want the instance to always restart on the same host on which it was launched, specify host. If you want the instance to restart on any available host, but try to launch onto the last host it ran on (on a best-effort basis), specify default.
func (o InstanceOutput) Affinity() InstanceAffinityPtrOutput {
	return o.ApplyT(func(v *Instance) InstanceAffinityPtrOutput { return v.Affinity }).(InstanceAffinityPtrOutput)
}

// The Availability Zone of the instance.
func (o InstanceOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The block device mapping entries that defines the block devices to attach to the instance at launch.
func (o InstanceOutput) BlockDeviceMappings() InstanceBlockDeviceMappingArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceBlockDeviceMappingArrayOutput { return v.BlockDeviceMappings }).(InstanceBlockDeviceMappingArrayOutput)
}

// The CPU options for the instance.
func (o InstanceOutput) CpuOptions() CpuOptionsPropertiesPtrOutput {
	return o.ApplyT(func(v *Instance) CpuOptionsPropertiesPtrOutput { return v.CpuOptions }).(CpuOptionsPropertiesPtrOutput)
}

// The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited.
func (o InstanceOutput) CreditSpecification() CreditSpecificationPropertiesPtrOutput {
	return o.ApplyT(func(v *Instance) CreditSpecificationPropertiesPtrOutput { return v.CreditSpecification }).(CreditSpecificationPropertiesPtrOutput)
}

// If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can.
func (o InstanceOutput) DisableApiTermination() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.DisableApiTermination }).(pulumi.BoolPtrOutput)
}

// Indicates whether the instance is optimized for Amazon EBS I/O.
func (o InstanceOutput) EbsOptimized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.EbsOptimized }).(pulumi.BoolPtrOutput)
}

// An elastic GPU to associate with the instance.
func (o InstanceOutput) ElasticGpuSpecifications() InstanceElasticGpuSpecificationArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceElasticGpuSpecificationArrayOutput { return v.ElasticGpuSpecifications }).(InstanceElasticGpuSpecificationArrayOutput)
}

// An elastic inference accelerator to associate with the instance.
func (o InstanceOutput) ElasticInferenceAccelerators() InstanceElasticInferenceAcceleratorArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceElasticInferenceAcceleratorArrayOutput {
		return v.ElasticInferenceAccelerators
	}).(InstanceElasticInferenceAcceleratorArrayOutput)
}

// Indicates whether the instance is enabled for AWS Nitro Enclaves.
func (o InstanceOutput) EnclaveOptions() EnclaveOptionsPropertiesPtrOutput {
	return o.ApplyT(func(v *Instance) EnclaveOptionsPropertiesPtrOutput { return v.EnclaveOptions }).(EnclaveOptionsPropertiesPtrOutput)
}

// Indicates whether an instance is enabled for hibernation.
func (o InstanceOutput) HibernationOptions() HibernationOptionsPropertiesPtrOutput {
	return o.ApplyT(func(v *Instance) HibernationOptionsPropertiesPtrOutput { return v.HibernationOptions }).(HibernationOptionsPropertiesPtrOutput)
}

// If you specify host for the Affinity property, the ID of a dedicated host that the instance is associated with. If you don't specify an ID, Amazon EC2 launches the instance onto any available, compatible dedicated host in your account.
func (o InstanceOutput) HostId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.HostId }).(pulumi.StringPtrOutput)
}

// The ARN of the host resource group in which to launch the instances. If you specify a host resource group ARN, omit the Tenancy parameter or set it to host.
func (o InstanceOutput) HostResourceGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.HostResourceGroupArn }).(pulumi.StringPtrOutput)
}

// The IAM instance profile.
func (o InstanceOutput) IamInstanceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.IamInstanceProfile }).(pulumi.StringPtrOutput)
}

// The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template.
func (o InstanceOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ImageId }).(pulumi.StringPtrOutput)
}

// The EC2 Instance ID.
func (o InstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
func (o InstanceOutput) InstanceInitiatedShutdownBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceInitiatedShutdownBehavior }).(pulumi.StringPtrOutput)
}

// The instance type.
func (o InstanceOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// [EC2-VPC] The number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
func (o InstanceOutput) Ipv6AddressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Ipv6AddressCount }).(pulumi.IntPtrOutput)
}

// [EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface.
func (o InstanceOutput) Ipv6Addresses() InstanceIpv6AddressArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceIpv6AddressArrayOutput { return v.Ipv6Addresses }).(InstanceIpv6AddressArrayOutput)
}

// The ID of the kernel.
func (o InstanceOutput) KernelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.KernelId }).(pulumi.StringPtrOutput)
}

// The name of the key pair.
func (o InstanceOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.KeyName }).(pulumi.StringPtrOutput)
}

// The launch template to use to launch the instances.
func (o InstanceOutput) LaunchTemplate() InstanceLaunchTemplateSpecificationPtrOutput {
	return o.ApplyT(func(v *Instance) InstanceLaunchTemplateSpecificationPtrOutput { return v.LaunchTemplate }).(InstanceLaunchTemplateSpecificationPtrOutput)
}

// The license configurations.
func (o InstanceOutput) LicenseSpecifications() InstanceLicenseSpecificationArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceLicenseSpecificationArrayOutput { return v.LicenseSpecifications }).(InstanceLicenseSpecificationArrayOutput)
}

// The metadata options for the instance
func (o InstanceOutput) MetadataOptions() InstanceMetadataOptionsPtrOutput {
	return o.ApplyT(func(v *Instance) InstanceMetadataOptionsPtrOutput { return v.MetadataOptions }).(InstanceMetadataOptionsPtrOutput)
}

// Specifies whether detailed monitoring is enabled for the instance.
func (o InstanceOutput) Monitoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.Monitoring }).(pulumi.BoolPtrOutput)
}

// The network interfaces to associate with the instance.
func (o InstanceOutput) NetworkInterfaces() InstanceNetworkInterfaceArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceNetworkInterfaceArrayOutput { return v.NetworkInterfaces }).(InstanceNetworkInterfaceArrayOutput)
}

// The name of an existing placement group that you want to launch the instance into (cluster | partition | spread).
func (o InstanceOutput) PlacementGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.PlacementGroupName }).(pulumi.StringPtrOutput)
}

// The private DNS name of the specified instance. For example: ip-10-24-34-0.ec2.internal.
func (o InstanceOutput) PrivateDnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PrivateDnsName }).(pulumi.StringOutput)
}

// The options for the instance hostname.
func (o InstanceOutput) PrivateDnsNameOptions() InstancePrivateDnsNameOptionsPtrOutput {
	return o.ApplyT(func(v *Instance) InstancePrivateDnsNameOptionsPtrOutput { return v.PrivateDnsNameOptions }).(InstancePrivateDnsNameOptionsPtrOutput)
}

// The private IP address of the specified instance. For example: 10.24.34.0.
func (o InstanceOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

// [EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet.
func (o InstanceOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

// Indicates whether to assign the tags from the instance to all of the volumes attached to the instance at launch. If you specify true and you assign tags to the instance, those tags are automatically assigned to all of the volumes that you attach to the instance at launch. If you specify false, those tags are not assigned to the attached volumes.
func (o InstanceOutput) PropagateTagsToVolumeOnCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.PropagateTagsToVolumeOnCreation }).(pulumi.BoolPtrOutput)
}

// The public DNS name of the specified instance. For example: ec2-107-20-50-45.compute-1.amazonaws.com.
func (o InstanceOutput) PublicDnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PublicDnsName }).(pulumi.StringOutput)
}

// The public IP address of the specified instance. For example: 192.0.2.0.
func (o InstanceOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PublicIp }).(pulumi.StringOutput)
}

// The ID of the RAM disk to select.
func (o InstanceOutput) RamdiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.RamdiskId }).(pulumi.StringPtrOutput)
}

// The IDs of the security groups.
func (o InstanceOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// the names of the security groups. For a nondefault VPC, you must use security group IDs instead.
func (o InstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Specifies whether to enable an instance launched in a VPC to perform NAT.
func (o InstanceOutput) SourceDestCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.SourceDestCheck }).(pulumi.BoolPtrOutput)
}

// The SSM document and parameter values in AWS Systems Manager to associate with this instance.
func (o InstanceOutput) SsmAssociations() InstanceSsmAssociationArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceSsmAssociationArrayOutput { return v.SsmAssociations }).(InstanceSsmAssociationArrayOutput)
}

// The current state of the instance.
func (o InstanceOutput) State() InstanceStateTypeOutput {
	return o.ApplyT(func(v *Instance) InstanceStateTypeOutput { return v.State }).(InstanceStateTypeOutput)
}

// [EC2-VPC] The ID of the subnet to launch the instance into.
func (o InstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// The tags to add to the instance.
func (o InstanceOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Instance) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware.
func (o InstanceOutput) Tenancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Tenancy }).(pulumi.StringPtrOutput)
}

// The user data to make available to the instance.
func (o InstanceOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

// The volumes to attach to the instance.
func (o InstanceOutput) Volumes() InstanceVolumeArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceVolumeArrayOutput { return v.Volumes }).(InstanceVolumeArrayOutput)
}

// The ID of the VPC that the instance is running in.
func (o InstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterOutputType(InstanceOutput{})
}
