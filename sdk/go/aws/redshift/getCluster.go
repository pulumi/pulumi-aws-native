// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("aws-native:redshift:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
}

type LookupClusterResult struct {
	// Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default value is True
	AllowVersionUpgrade *bool `pulumi:"allowVersionUpgrade"`
	// The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored. Possible values include the following.
	//
	// enabled - Use AQUA if it is available for the current Region and Amazon Redshift node type.
	// disabled - Don't use AQUA.
	// auto - Amazon Redshift determines whether to use AQUA.
	AquaConfigurationStatus *string `pulumi:"aquaConfigurationStatus"`
	// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Default value is 1
	AutomatedSnapshotRetentionPeriod *int `pulumi:"automatedSnapshotRetentionPeriod"`
	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster modification is complete.
	AvailabilityZoneRelocation *bool `pulumi:"availabilityZoneRelocation"`
	// The availability zone relocation status of the cluster
	AvailabilityZoneRelocationStatus *string `pulumi:"availabilityZoneRelocationStatus"`
	// The Amazon Resource Name (ARN) of the cluster namespace.
	ClusterNamespaceArn *string `pulumi:"clusterNamespaceArn"`
	// The name of the parameter group to be associated with this cluster.
	ClusterParameterGroupName *string `pulumi:"clusterParameterGroupName"`
	// A list of security groups to be associated with this cluster.
	ClusterSecurityGroups []string `pulumi:"clusterSecurityGroups"`
	// The type of the cluster. When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required
	ClusterType *string `pulumi:"clusterType"`
	// The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster.
	ClusterVersion *string `pulumi:"clusterVersion"`
	// A timestamp indicating end time for the deferred maintenance window. If you specify an end time, you can't specify a duration.
	DeferMaintenanceEndTime *string `pulumi:"deferMaintenanceEndTime"`
	// A unique identifier for the deferred maintenance window.
	DeferMaintenanceIdentifier *string `pulumi:"deferMaintenanceIdentifier"`
	// A timestamp indicating the start time for the deferred maintenance window.
	DeferMaintenanceStartTime *string `pulumi:"deferMaintenanceStartTime"`
	// The destination AWS Region that you want to copy snapshots to. Constraints: Must be the name of a valid AWS Region. For more information, see Regions and Endpoints in the Amazon Web Services [https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region] General Reference
	DestinationRegion *string `pulumi:"destinationRegion"`
	// The Elastic IP (EIP) address for the cluster.
	ElasticIp *string `pulumi:"elasticIp"`
	// If true, the data in the cluster is encrypted at rest.
	Encrypted *bool `pulumi:"encrypted"`
	// The connection endpoint.
	Endpoint *ClusterEndpoint `pulumi:"endpoint"`
	// An option that specifies whether to create the cluster with enhanced VPC routing enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see Enhanced VPC Routing in the Amazon Redshift Cluster Management Guide.
	//
	// If this option is true , enhanced VPC routing is enabled.
	//
	// Default: false
	EnhancedVpcRouting *bool `pulumi:"enhancedVpcRouting"`
	// Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM
	HsmClientCertificateIdentifier *string `pulumi:"hsmClientCertificateIdentifier"`
	// Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.
	HsmConfigurationIdentifier *string `pulumi:"hsmConfigurationIdentifier"`
	// A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 50 IAM roles in a single request
	IamRoles []string `pulumi:"iamRoles"`
	// The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies logging information, such as queries and connection attempts, for the specified Amazon Redshift cluster.
	LoggingProperties *ClusterLoggingProperties `pulumi:"loggingProperties"`
	// The name for the maintenance track that you want to assign for the cluster. This name change is asynchronous. The new track name stays in the PendingModifiedValues for the cluster until the next maintenance window. When the maintenance track changes, the cluster is switched to the latest cluster release available for the maintenance track. At this point, the maintenance track name is applied.
	MaintenanceTrackName *string `pulumi:"maintenanceTrackName"`
	// The number of days to retain newly copied snapshots in the destination AWS Region after they are copied from the source AWS Region. If the value is -1, the manual snapshot is retained indefinitely.
	//
	// The value must be either -1 or an integer between 1 and 3,653.
	ManualSnapshotRetentionPeriod *int `pulumi:"manualSnapshotRetentionPeriod"`
	// The Amazon Resource Name (ARN) for the cluster's admin user credentials secret.
	MasterPasswordSecretArn *string `pulumi:"masterPasswordSecretArn"`
	// The ID of the Key Management Service (KMS) key used to encrypt and store the cluster's admin user credentials secret.
	MasterPasswordSecretKmsKeyId *string `pulumi:"masterPasswordSecretKmsKeyId"`
	// A boolean indicating if the redshift cluster is multi-az or not. If you don't provide this parameter or set the value to false, the redshift cluster will be single-az.
	MultiAz *bool `pulumi:"multiAz"`
	// The namespace resource policy document that will be attached to a Redshift cluster.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Redshift::Cluster` for more information about the expected schema for this property.
	NamespaceResourcePolicy interface{} `pulumi:"namespaceResourcePolicy"`
	// The node type to be provisioned for the cluster.Valid Values: ds2.xlarge | ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.large | ra3.4xlarge | ra3.16xlarge
	NodeType *string `pulumi:"nodeType"`
	// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node.
	NumberOfNodes *int `pulumi:"numberOfNodes"`
	// The port number on which the cluster accepts incoming connections. The cluster is accessible only via the JDBC and ODBC connection strings
	Port *int `pulumi:"port"`
	// The weekly time range (in UTC) during which automated cluster maintenance can occur.
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// If true, the cluster can be accessed from a public network.
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// The Redshift operation to be performed. Resource Action supports pause-cluster, resume-cluster, failover-primary-compute APIs
	ResourceAction *string `pulumi:"resourceAction"`
	// The identifier of the database revision. You can retrieve this value from the response to the DescribeClusterDbRevisions request.
	RevisionTarget *string `pulumi:"revisionTarget"`
	// A boolean indicating if we want to rotate Encryption Keys.
	RotateEncryptionKey *bool `pulumi:"rotateEncryptionKey"`
	// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
	SnapshotCopyGrantName *string `pulumi:"snapshotCopyGrantName"`
	// Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.
	SnapshotCopyManual *bool `pulumi:"snapshotCopyManual"`
	// The number of days to retain automated snapshots in the destination region after they are copied from the source region.
	//
	//  Default is 7.
	//
	//  Constraints: Must be at least 1 and no more than 35.
	SnapshotCopyRetentionPeriod *int `pulumi:"snapshotCopyRetentionPeriod"`
	// The list of tags for the cluster parameter group.
	Tags []aws.Tag `pulumi:"tags"`
	// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupClusterResultOutput, error) {
			args := v.(LookupClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:redshift:getCluster", args, LookupClusterResultOutput{}, options).(LookupClusterResultOutput), nil
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	// A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
	ClusterIdentifier pulumi.StringInput `pulumi:"clusterIdentifier"`
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

// Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default value is True
func (o LookupClusterResultOutput) AllowVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.AllowVersionUpgrade }).(pulumi.BoolPtrOutput)
}

// The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored. Possible values include the following.
//
// enabled - Use AQUA if it is available for the current Region and Amazon Redshift node type.
// disabled - Don't use AQUA.
// auto - Amazon Redshift determines whether to use AQUA.
func (o LookupClusterResultOutput) AquaConfigurationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.AquaConfigurationStatus }).(pulumi.StringPtrOutput)
}

// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Default value is 1
func (o LookupClusterResultOutput) AutomatedSnapshotRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.AutomatedSnapshotRetentionPeriod }).(pulumi.IntPtrOutput)
}

// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint
func (o LookupClusterResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster modification is complete.
func (o LookupClusterResultOutput) AvailabilityZoneRelocation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.AvailabilityZoneRelocation }).(pulumi.BoolPtrOutput)
}

// The availability zone relocation status of the cluster
func (o LookupClusterResultOutput) AvailabilityZoneRelocationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.AvailabilityZoneRelocationStatus }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the cluster namespace.
func (o LookupClusterResultOutput) ClusterNamespaceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterNamespaceArn }).(pulumi.StringPtrOutput)
}

// The name of the parameter group to be associated with this cluster.
func (o LookupClusterResultOutput) ClusterParameterGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterParameterGroupName }).(pulumi.StringPtrOutput)
}

// A list of security groups to be associated with this cluster.
func (o LookupClusterResultOutput) ClusterSecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.ClusterSecurityGroups }).(pulumi.StringArrayOutput)
}

// The type of the cluster. When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required
func (o LookupClusterResultOutput) ClusterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterType }).(pulumi.StringPtrOutput)
}

// The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster.
func (o LookupClusterResultOutput) ClusterVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterVersion }).(pulumi.StringPtrOutput)
}

// A timestamp indicating end time for the deferred maintenance window. If you specify an end time, you can't specify a duration.
func (o LookupClusterResultOutput) DeferMaintenanceEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.DeferMaintenanceEndTime }).(pulumi.StringPtrOutput)
}

// A unique identifier for the deferred maintenance window.
func (o LookupClusterResultOutput) DeferMaintenanceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.DeferMaintenanceIdentifier }).(pulumi.StringPtrOutput)
}

// A timestamp indicating the start time for the deferred maintenance window.
func (o LookupClusterResultOutput) DeferMaintenanceStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.DeferMaintenanceStartTime }).(pulumi.StringPtrOutput)
}

// The destination AWS Region that you want to copy snapshots to. Constraints: Must be the name of a valid AWS Region. For more information, see Regions and Endpoints in the Amazon Web Services [https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region] General Reference
func (o LookupClusterResultOutput) DestinationRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.DestinationRegion }).(pulumi.StringPtrOutput)
}

// The Elastic IP (EIP) address for the cluster.
func (o LookupClusterResultOutput) ElasticIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ElasticIp }).(pulumi.StringPtrOutput)
}

// If true, the data in the cluster is encrypted at rest.
func (o LookupClusterResultOutput) Encrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.Encrypted }).(pulumi.BoolPtrOutput)
}

// The connection endpoint.
func (o LookupClusterResultOutput) Endpoint() ClusterEndpointPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterEndpoint { return v.Endpoint }).(ClusterEndpointPtrOutput)
}

// An option that specifies whether to create the cluster with enhanced VPC routing enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see Enhanced VPC Routing in the Amazon Redshift Cluster Management Guide.
//
// If this option is true , enhanced VPC routing is enabled.
//
// Default: false
func (o LookupClusterResultOutput) EnhancedVpcRouting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EnhancedVpcRouting }).(pulumi.BoolPtrOutput)
}

// Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM
func (o LookupClusterResultOutput) HsmClientCertificateIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.HsmClientCertificateIdentifier }).(pulumi.StringPtrOutput)
}

// Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.
func (o LookupClusterResultOutput) HsmConfigurationIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.HsmConfigurationIdentifier }).(pulumi.StringPtrOutput)
}

// A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 50 IAM roles in a single request
func (o LookupClusterResultOutput) IamRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.IamRoles }).(pulumi.StringArrayOutput)
}

// The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.
func (o LookupClusterResultOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// Specifies logging information, such as queries and connection attempts, for the specified Amazon Redshift cluster.
func (o LookupClusterResultOutput) LoggingProperties() ClusterLoggingPropertiesPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterLoggingProperties { return v.LoggingProperties }).(ClusterLoggingPropertiesPtrOutput)
}

// The name for the maintenance track that you want to assign for the cluster. This name change is asynchronous. The new track name stays in the PendingModifiedValues for the cluster until the next maintenance window. When the maintenance track changes, the cluster is switched to the latest cluster release available for the maintenance track. At this point, the maintenance track name is applied.
func (o LookupClusterResultOutput) MaintenanceTrackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.MaintenanceTrackName }).(pulumi.StringPtrOutput)
}

// The number of days to retain newly copied snapshots in the destination AWS Region after they are copied from the source AWS Region. If the value is -1, the manual snapshot is retained indefinitely.
//
// The value must be either -1 or an integer between 1 and 3,653.
func (o LookupClusterResultOutput) ManualSnapshotRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.ManualSnapshotRetentionPeriod }).(pulumi.IntPtrOutput)
}

// The Amazon Resource Name (ARN) for the cluster's admin user credentials secret.
func (o LookupClusterResultOutput) MasterPasswordSecretArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.MasterPasswordSecretArn }).(pulumi.StringPtrOutput)
}

// The ID of the Key Management Service (KMS) key used to encrypt and store the cluster's admin user credentials secret.
func (o LookupClusterResultOutput) MasterPasswordSecretKmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.MasterPasswordSecretKmsKeyId }).(pulumi.StringPtrOutput)
}

// A boolean indicating if the redshift cluster is multi-az or not. If you don't provide this parameter or set the value to false, the redshift cluster will be single-az.
func (o LookupClusterResultOutput) MultiAz() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.MultiAz }).(pulumi.BoolPtrOutput)
}

// The namespace resource policy document that will be attached to a Redshift cluster.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Redshift::Cluster` for more information about the expected schema for this property.
func (o LookupClusterResultOutput) NamespaceResourcePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupClusterResult) interface{} { return v.NamespaceResourcePolicy }).(pulumi.AnyOutput)
}

// The node type to be provisioned for the cluster.Valid Values: ds2.xlarge | ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.large | ra3.4xlarge | ra3.16xlarge
func (o LookupClusterResultOutput) NodeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.NodeType }).(pulumi.StringPtrOutput)
}

// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node.
func (o LookupClusterResultOutput) NumberOfNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.NumberOfNodes }).(pulumi.IntPtrOutput)
}

// The port number on which the cluster accepts incoming connections. The cluster is accessible only via the JDBC and ODBC connection strings
func (o LookupClusterResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// The weekly time range (in UTC) during which automated cluster maintenance can occur.
func (o LookupClusterResultOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

// If true, the cluster can be accessed from a public network.
func (o LookupClusterResultOutput) PubliclyAccessible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.PubliclyAccessible }).(pulumi.BoolPtrOutput)
}

// The Redshift operation to be performed. Resource Action supports pause-cluster, resume-cluster, failover-primary-compute APIs
func (o LookupClusterResultOutput) ResourceAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ResourceAction }).(pulumi.StringPtrOutput)
}

// The identifier of the database revision. You can retrieve this value from the response to the DescribeClusterDbRevisions request.
func (o LookupClusterResultOutput) RevisionTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.RevisionTarget }).(pulumi.StringPtrOutput)
}

// A boolean indicating if we want to rotate Encryption Keys.
func (o LookupClusterResultOutput) RotateEncryptionKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.RotateEncryptionKey }).(pulumi.BoolPtrOutput)
}

// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
func (o LookupClusterResultOutput) SnapshotCopyGrantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.SnapshotCopyGrantName }).(pulumi.StringPtrOutput)
}

// Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.
func (o LookupClusterResultOutput) SnapshotCopyManual() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.SnapshotCopyManual }).(pulumi.BoolPtrOutput)
}

// The number of days to retain automated snapshots in the destination region after they are copied from the source region.
//
//	Default is 7.
//
//	Constraints: Must be at least 1 and no more than 35.
func (o LookupClusterResultOutput) SnapshotCopyRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.SnapshotCopyRetentionPeriod }).(pulumi.IntPtrOutput)
}

// The list of tags for the cluster parameter group.
func (o LookupClusterResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
func (o LookupClusterResultOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
