// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codestarconnections

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for AWS::CodeStarConnections::SyncConfiguration resource which is used to enables an AWS resource to be synchronized from a source-provider.
func LookupSyncConfiguration(ctx *pulumi.Context, args *LookupSyncConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSyncConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSyncConfigurationResult
	err := ctx.Invoke("aws-native:codestarconnections:getSyncConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncConfigurationArgs struct {
	// The name of the resource that is being synchronized to the repository.
	ResourceName string `pulumi:"resourceName"`
	// The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.
	SyncType string `pulumi:"syncType"`
}

type LookupSyncConfigurationResult struct {
	// The name of the branch of the repository from which resources are to be synchronized,
	Branch *string `pulumi:"branch"`
	// The source provider repository path of the sync configuration file of the respective SyncType.
	ConfigFile *string `pulumi:"configFile"`
	// the ID of the entity that owns the repository.
	OwnerId *string `pulumi:"ownerId"`
	// The name of the external provider where your third-party code repository is configured.
	ProviderType *SyncConfigurationProviderType `pulumi:"providerType"`
	// Whether to enable or disable publishing of deployment status to source providers.
	PublishDeploymentStatus *SyncConfigurationPublishDeploymentStatus `pulumi:"publishDeploymentStatus"`
	// A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.
	RepositoryLinkId *string `pulumi:"repositoryLinkId"`
	// The name of the repository that is being synced to.
	RepositoryName *string `pulumi:"repositoryName"`
	// The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.
	RoleArn *string `pulumi:"roleArn"`
	// When to trigger Git sync to begin the stack update.
	TriggerResourceUpdateOn *SyncConfigurationTriggerResourceUpdateOn `pulumi:"triggerResourceUpdateOn"`
}

func LookupSyncConfigurationOutput(ctx *pulumi.Context, args LookupSyncConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupSyncConfigurationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSyncConfigurationResultOutput, error) {
			args := v.(LookupSyncConfigurationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:codestarconnections:getSyncConfiguration", args, LookupSyncConfigurationResultOutput{}, options).(LookupSyncConfigurationResultOutput), nil
		}).(LookupSyncConfigurationResultOutput)
}

type LookupSyncConfigurationOutputArgs struct {
	// The name of the resource that is being synchronized to the repository.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.
	SyncType pulumi.StringInput `pulumi:"syncType"`
}

func (LookupSyncConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncConfigurationArgs)(nil)).Elem()
}

type LookupSyncConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupSyncConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncConfigurationResult)(nil)).Elem()
}

func (o LookupSyncConfigurationResultOutput) ToLookupSyncConfigurationResultOutput() LookupSyncConfigurationResultOutput {
	return o
}

func (o LookupSyncConfigurationResultOutput) ToLookupSyncConfigurationResultOutputWithContext(ctx context.Context) LookupSyncConfigurationResultOutput {
	return o
}

// The name of the branch of the repository from which resources are to be synchronized,
func (o LookupSyncConfigurationResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncConfigurationResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

// The source provider repository path of the sync configuration file of the respective SyncType.
func (o LookupSyncConfigurationResultOutput) ConfigFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncConfigurationResult) *string { return v.ConfigFile }).(pulumi.StringPtrOutput)
}

// the ID of the entity that owns the repository.
func (o LookupSyncConfigurationResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncConfigurationResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// The name of the external provider where your third-party code repository is configured.
func (o LookupSyncConfigurationResultOutput) ProviderType() SyncConfigurationProviderTypePtrOutput {
	return o.ApplyT(func(v LookupSyncConfigurationResult) *SyncConfigurationProviderType { return v.ProviderType }).(SyncConfigurationProviderTypePtrOutput)
}

// Whether to enable or disable publishing of deployment status to source providers.
func (o LookupSyncConfigurationResultOutput) PublishDeploymentStatus() SyncConfigurationPublishDeploymentStatusPtrOutput {
	return o.ApplyT(func(v LookupSyncConfigurationResult) *SyncConfigurationPublishDeploymentStatus {
		return v.PublishDeploymentStatus
	}).(SyncConfigurationPublishDeploymentStatusPtrOutput)
}

// A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.
func (o LookupSyncConfigurationResultOutput) RepositoryLinkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncConfigurationResult) *string { return v.RepositoryLinkId }).(pulumi.StringPtrOutput)
}

// The name of the repository that is being synced to.
func (o LookupSyncConfigurationResultOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncConfigurationResult) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

// The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.
func (o LookupSyncConfigurationResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncConfigurationResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// When to trigger Git sync to begin the stack update.
func (o LookupSyncConfigurationResultOutput) TriggerResourceUpdateOn() SyncConfigurationTriggerResourceUpdateOnPtrOutput {
	return o.ApplyT(func(v LookupSyncConfigurationResult) *SyncConfigurationTriggerResourceUpdateOn {
		return v.TriggerResourceUpdateOn
	}).(SyncConfigurationTriggerResourceUpdateOnPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSyncConfigurationResultOutput{})
}
