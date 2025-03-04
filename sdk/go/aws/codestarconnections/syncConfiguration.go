// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codestarconnections

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for AWS::CodeStarConnections::SyncConfiguration resource which is used to enables an AWS resource to be synchronized from a source-provider.
type SyncConfiguration struct {
	pulumi.CustomResourceState

	// The name of the branch of the repository from which resources are to be synchronized,
	Branch pulumi.StringOutput `pulumi:"branch"`
	// The source provider repository path of the sync configuration file of the respective SyncType.
	ConfigFile pulumi.StringOutput `pulumi:"configFile"`
	// the ID of the entity that owns the repository.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The name of the external provider where your third-party code repository is configured.
	ProviderType SyncConfigurationProviderTypeOutput `pulumi:"providerType"`
	// Whether to enable or disable publishing of deployment status to source providers.
	PublishDeploymentStatus SyncConfigurationPublishDeploymentStatusPtrOutput `pulumi:"publishDeploymentStatus"`
	// A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.
	RepositoryLinkId pulumi.StringOutput `pulumi:"repositoryLinkId"`
	// The name of the repository that is being synced to.
	RepositoryName pulumi.StringOutput `pulumi:"repositoryName"`
	// The name of the resource that is being synchronized to the repository.
	ResourceName pulumi.StringOutput `pulumi:"resourceName"`
	// The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.
	SyncType pulumi.StringOutput `pulumi:"syncType"`
	// When to trigger Git sync to begin the stack update.
	TriggerResourceUpdateOn SyncConfigurationTriggerResourceUpdateOnPtrOutput `pulumi:"triggerResourceUpdateOn"`
}

// NewSyncConfiguration registers a new resource with the given unique name, arguments, and options.
func NewSyncConfiguration(ctx *pulumi.Context,
	name string, args *SyncConfigurationArgs, opts ...pulumi.ResourceOption) (*SyncConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Branch == nil {
		return nil, errors.New("invalid value for required argument 'Branch'")
	}
	if args.ConfigFile == nil {
		return nil, errors.New("invalid value for required argument 'ConfigFile'")
	}
	if args.RepositoryLinkId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryLinkId'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.SyncType == nil {
		return nil, errors.New("invalid value for required argument 'SyncType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"resourceName",
		"syncType",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SyncConfiguration
	err := ctx.RegisterResource("aws-native:codestarconnections:SyncConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncConfiguration gets an existing SyncConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncConfigurationState, opts ...pulumi.ResourceOption) (*SyncConfiguration, error) {
	var resource SyncConfiguration
	err := ctx.ReadResource("aws-native:codestarconnections:SyncConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncConfiguration resources.
type syncConfigurationState struct {
}

type SyncConfigurationState struct {
}

func (SyncConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncConfigurationState)(nil)).Elem()
}

type syncConfigurationArgs struct {
	// The name of the branch of the repository from which resources are to be synchronized,
	Branch string `pulumi:"branch"`
	// The source provider repository path of the sync configuration file of the respective SyncType.
	ConfigFile string `pulumi:"configFile"`
	// Whether to enable or disable publishing of deployment status to source providers.
	PublishDeploymentStatus *SyncConfigurationPublishDeploymentStatus `pulumi:"publishDeploymentStatus"`
	// A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.
	RepositoryLinkId string `pulumi:"repositoryLinkId"`
	// The name of the resource that is being synchronized to the repository.
	ResourceName string `pulumi:"resourceName"`
	// The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.
	RoleArn string `pulumi:"roleArn"`
	// The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.
	SyncType string `pulumi:"syncType"`
	// When to trigger Git sync to begin the stack update.
	TriggerResourceUpdateOn *SyncConfigurationTriggerResourceUpdateOn `pulumi:"triggerResourceUpdateOn"`
}

// The set of arguments for constructing a SyncConfiguration resource.
type SyncConfigurationArgs struct {
	// The name of the branch of the repository from which resources are to be synchronized,
	Branch pulumi.StringInput
	// The source provider repository path of the sync configuration file of the respective SyncType.
	ConfigFile pulumi.StringInput
	// Whether to enable or disable publishing of deployment status to source providers.
	PublishDeploymentStatus SyncConfigurationPublishDeploymentStatusPtrInput
	// A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.
	RepositoryLinkId pulumi.StringInput
	// The name of the resource that is being synchronized to the repository.
	ResourceName pulumi.StringInput
	// The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.
	RoleArn pulumi.StringInput
	// The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.
	SyncType pulumi.StringInput
	// When to trigger Git sync to begin the stack update.
	TriggerResourceUpdateOn SyncConfigurationTriggerResourceUpdateOnPtrInput
}

func (SyncConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncConfigurationArgs)(nil)).Elem()
}

type SyncConfigurationInput interface {
	pulumi.Input

	ToSyncConfigurationOutput() SyncConfigurationOutput
	ToSyncConfigurationOutputWithContext(ctx context.Context) SyncConfigurationOutput
}

func (*SyncConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncConfiguration)(nil)).Elem()
}

func (i *SyncConfiguration) ToSyncConfigurationOutput() SyncConfigurationOutput {
	return i.ToSyncConfigurationOutputWithContext(context.Background())
}

func (i *SyncConfiguration) ToSyncConfigurationOutputWithContext(ctx context.Context) SyncConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncConfigurationOutput)
}

type SyncConfigurationOutput struct{ *pulumi.OutputState }

func (SyncConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncConfiguration)(nil)).Elem()
}

func (o SyncConfigurationOutput) ToSyncConfigurationOutput() SyncConfigurationOutput {
	return o
}

func (o SyncConfigurationOutput) ToSyncConfigurationOutputWithContext(ctx context.Context) SyncConfigurationOutput {
	return o
}

// The name of the branch of the repository from which resources are to be synchronized,
func (o SyncConfigurationOutput) Branch() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncConfiguration) pulumi.StringOutput { return v.Branch }).(pulumi.StringOutput)
}

// The source provider repository path of the sync configuration file of the respective SyncType.
func (o SyncConfigurationOutput) ConfigFile() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncConfiguration) pulumi.StringOutput { return v.ConfigFile }).(pulumi.StringOutput)
}

// the ID of the entity that owns the repository.
func (o SyncConfigurationOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncConfiguration) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The name of the external provider where your third-party code repository is configured.
func (o SyncConfigurationOutput) ProviderType() SyncConfigurationProviderTypeOutput {
	return o.ApplyT(func(v *SyncConfiguration) SyncConfigurationProviderTypeOutput { return v.ProviderType }).(SyncConfigurationProviderTypeOutput)
}

// Whether to enable or disable publishing of deployment status to source providers.
func (o SyncConfigurationOutput) PublishDeploymentStatus() SyncConfigurationPublishDeploymentStatusPtrOutput {
	return o.ApplyT(func(v *SyncConfiguration) SyncConfigurationPublishDeploymentStatusPtrOutput {
		return v.PublishDeploymentStatus
	}).(SyncConfigurationPublishDeploymentStatusPtrOutput)
}

// A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.
func (o SyncConfigurationOutput) RepositoryLinkId() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncConfiguration) pulumi.StringOutput { return v.RepositoryLinkId }).(pulumi.StringOutput)
}

// The name of the repository that is being synced to.
func (o SyncConfigurationOutput) RepositoryName() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncConfiguration) pulumi.StringOutput { return v.RepositoryName }).(pulumi.StringOutput)
}

// The name of the resource that is being synchronized to the repository.
func (o SyncConfigurationOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncConfiguration) pulumi.StringOutput { return v.ResourceName }).(pulumi.StringOutput)
}

// The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.
func (o SyncConfigurationOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncConfiguration) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.
func (o SyncConfigurationOutput) SyncType() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncConfiguration) pulumi.StringOutput { return v.SyncType }).(pulumi.StringOutput)
}

// When to trigger Git sync to begin the stack update.
func (o SyncConfigurationOutput) TriggerResourceUpdateOn() SyncConfigurationTriggerResourceUpdateOnPtrOutput {
	return o.ApplyT(func(v *SyncConfiguration) SyncConfigurationTriggerResourceUpdateOnPtrOutput {
		return v.TriggerResourceUpdateOn
	}).(SyncConfigurationTriggerResourceUpdateOnPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncConfigurationInput)(nil)).Elem(), &SyncConfiguration{})
	pulumi.RegisterOutputType(SyncConfigurationOutput{})
}
