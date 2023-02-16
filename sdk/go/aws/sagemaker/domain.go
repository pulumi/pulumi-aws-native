// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::Domain
type Domain struct {
	pulumi.CustomResourceState

	// Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.
	AppNetworkAccessType DomainAppNetworkAccessTypePtrOutput `pulumi:"appNetworkAccessType"`
	// The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.
	AppSecurityGroupManagement DomainAppSecurityGroupManagementPtrOutput `pulumi:"appSecurityGroupManagement"`
	// The mode of authentication that members use to access the domain.
	AuthMode DomainAuthModeOutput `pulumi:"authMode"`
	// The default space settings.
	DefaultSpaceSettings DomainDefaultSpaceSettingsPtrOutput `pulumi:"defaultSpaceSettings"`
	// The default user settings.
	DefaultUserSettings DomainUserSettingsOutput `pulumi:"defaultUserSettings"`
	// The Amazon Resource Name (ARN) of the created domain.
	DomainArn pulumi.StringOutput `pulumi:"domainArn"`
	// The domain name.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// A name for the domain.
	DomainName     pulumi.StringOutput     `pulumi:"domainName"`
	DomainSettings DomainSettingsPtrOutput `pulumi:"domainSettings"`
	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId pulumi.StringOutput `pulumi:"homeEfsFileSystemId"`
	// SageMaker uses AWS KMS to encrypt the EFS volume attached to the domain with an AWS managed customer master key (CMK) by default.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
	SecurityGroupIdForDomainBoundary pulumi.StringOutput `pulumi:"securityGroupIdForDomainBoundary"`
	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceId pulumi.StringOutput `pulumi:"singleSignOnManagedApplicationInstanceId"`
	// The VPC subnets that Studio uses for communication.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A list of tags to apply to the user profile.
	Tags DomainTagArrayOutput `pulumi:"tags"`
	// The URL to the created domain.
	Url pulumi.StringOutput `pulumi:"url"`
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthMode == nil {
		return nil, errors.New("invalid value for required argument 'AuthMode'")
	}
	if args.DefaultUserSettings == nil {
		return nil, errors.New("invalid value for required argument 'DefaultUserSettings'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource Domain
	err := ctx.RegisterResource("aws-native:sagemaker:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("aws-native:sagemaker:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
}

type DomainState struct {
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.
	AppNetworkAccessType *DomainAppNetworkAccessType `pulumi:"appNetworkAccessType"`
	// The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.
	AppSecurityGroupManagement *DomainAppSecurityGroupManagement `pulumi:"appSecurityGroupManagement"`
	// The mode of authentication that members use to access the domain.
	AuthMode DomainAuthMode `pulumi:"authMode"`
	// The default space settings.
	DefaultSpaceSettings *DomainDefaultSpaceSettings `pulumi:"defaultSpaceSettings"`
	// The default user settings.
	DefaultUserSettings DomainUserSettings `pulumi:"defaultUserSettings"`
	// A name for the domain.
	DomainName     *string         `pulumi:"domainName"`
	DomainSettings *DomainSettings `pulumi:"domainSettings"`
	// SageMaker uses AWS KMS to encrypt the EFS volume attached to the domain with an AWS managed customer master key (CMK) by default.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The VPC subnets that Studio uses for communication.
	SubnetIds []string `pulumi:"subnetIds"`
	// A list of tags to apply to the user profile.
	Tags []DomainTag `pulumi:"tags"`
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.
	AppNetworkAccessType DomainAppNetworkAccessTypePtrInput
	// The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.
	AppSecurityGroupManagement DomainAppSecurityGroupManagementPtrInput
	// The mode of authentication that members use to access the domain.
	AuthMode DomainAuthModeInput
	// The default space settings.
	DefaultSpaceSettings DomainDefaultSpaceSettingsPtrInput
	// The default user settings.
	DefaultUserSettings DomainUserSettingsInput
	// A name for the domain.
	DomainName     pulumi.StringPtrInput
	DomainSettings DomainSettingsPtrInput
	// SageMaker uses AWS KMS to encrypt the EFS volume attached to the domain with an AWS managed customer master key (CMK) by default.
	KmsKeyId pulumi.StringPtrInput
	// The VPC subnets that Studio uses for communication.
	SubnetIds pulumi.StringArrayInput
	// A list of tags to apply to the user profile.
	Tags DomainTagArrayInput
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	VpcId pulumi.StringInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.
func (o DomainOutput) AppNetworkAccessType() DomainAppNetworkAccessTypePtrOutput {
	return o.ApplyT(func(v *Domain) DomainAppNetworkAccessTypePtrOutput { return v.AppNetworkAccessType }).(DomainAppNetworkAccessTypePtrOutput)
}

// The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.
func (o DomainOutput) AppSecurityGroupManagement() DomainAppSecurityGroupManagementPtrOutput {
	return o.ApplyT(func(v *Domain) DomainAppSecurityGroupManagementPtrOutput { return v.AppSecurityGroupManagement }).(DomainAppSecurityGroupManagementPtrOutput)
}

// The mode of authentication that members use to access the domain.
func (o DomainOutput) AuthMode() DomainAuthModeOutput {
	return o.ApplyT(func(v *Domain) DomainAuthModeOutput { return v.AuthMode }).(DomainAuthModeOutput)
}

// The default space settings.
func (o DomainOutput) DefaultSpaceSettings() DomainDefaultSpaceSettingsPtrOutput {
	return o.ApplyT(func(v *Domain) DomainDefaultSpaceSettingsPtrOutput { return v.DefaultSpaceSettings }).(DomainDefaultSpaceSettingsPtrOutput)
}

// The default user settings.
func (o DomainOutput) DefaultUserSettings() DomainUserSettingsOutput {
	return o.ApplyT(func(v *Domain) DomainUserSettingsOutput { return v.DefaultUserSettings }).(DomainUserSettingsOutput)
}

// The Amazon Resource Name (ARN) of the created domain.
func (o DomainOutput) DomainArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainArn }).(pulumi.StringOutput)
}

// The domain name.
func (o DomainOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// A name for the domain.
func (o DomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o DomainOutput) DomainSettings() DomainSettingsPtrOutput {
	return o.ApplyT(func(v *Domain) DomainSettingsPtrOutput { return v.DomainSettings }).(DomainSettingsPtrOutput)
}

// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
func (o DomainOutput) HomeEfsFileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.HomeEfsFileSystemId }).(pulumi.StringOutput)
}

// SageMaker uses AWS KMS to encrypt the EFS volume attached to the domain with an AWS managed customer master key (CMK) by default.
func (o DomainOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
func (o DomainOutput) SecurityGroupIdForDomainBoundary() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.SecurityGroupIdForDomainBoundary }).(pulumi.StringOutput)
}

// The SSO managed application instance ID.
func (o DomainOutput) SingleSignOnManagedApplicationInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.SingleSignOnManagedApplicationInstanceId }).(pulumi.StringOutput)
}

// The VPC subnets that Studio uses for communication.
func (o DomainOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// A list of tags to apply to the user profile.
func (o DomainOutput) Tags() DomainTagArrayOutput {
	return o.ApplyT(func(v *Domain) DomainTagArrayOutput { return v.Tags }).(DomainTagArrayOutput)
}

// The URL to the created domain.
func (o DomainOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
func (o DomainOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterOutputType(DomainOutput{})
}