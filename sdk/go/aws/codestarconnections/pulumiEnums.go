// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codestarconnections

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The name of the external provider where your third-party code repository is configured.
type RepositoryLinkProviderType string

const (
	RepositoryLinkProviderTypeGitHub            = RepositoryLinkProviderType("GitHub")
	RepositoryLinkProviderTypeBitbucket         = RepositoryLinkProviderType("Bitbucket")
	RepositoryLinkProviderTypeGitHubEnterprise  = RepositoryLinkProviderType("GitHubEnterprise")
	RepositoryLinkProviderTypeGitLab            = RepositoryLinkProviderType("GitLab")
	RepositoryLinkProviderTypeGitLabSelfManaged = RepositoryLinkProviderType("GitLabSelfManaged")
)

type RepositoryLinkProviderTypeOutput struct{ *pulumi.OutputState }

func (RepositoryLinkProviderTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryLinkProviderType)(nil)).Elem()
}

func (o RepositoryLinkProviderTypeOutput) ToRepositoryLinkProviderTypeOutput() RepositoryLinkProviderTypeOutput {
	return o
}

func (o RepositoryLinkProviderTypeOutput) ToRepositoryLinkProviderTypeOutputWithContext(ctx context.Context) RepositoryLinkProviderTypeOutput {
	return o
}

func (o RepositoryLinkProviderTypeOutput) ToRepositoryLinkProviderTypePtrOutput() RepositoryLinkProviderTypePtrOutput {
	return o.ToRepositoryLinkProviderTypePtrOutputWithContext(context.Background())
}

func (o RepositoryLinkProviderTypeOutput) ToRepositoryLinkProviderTypePtrOutputWithContext(ctx context.Context) RepositoryLinkProviderTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RepositoryLinkProviderType) *RepositoryLinkProviderType {
		return &v
	}).(RepositoryLinkProviderTypePtrOutput)
}

func (o RepositoryLinkProviderTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RepositoryLinkProviderTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RepositoryLinkProviderType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RepositoryLinkProviderTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RepositoryLinkProviderTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RepositoryLinkProviderType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RepositoryLinkProviderTypePtrOutput struct{ *pulumi.OutputState }

func (RepositoryLinkProviderTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryLinkProviderType)(nil)).Elem()
}

func (o RepositoryLinkProviderTypePtrOutput) ToRepositoryLinkProviderTypePtrOutput() RepositoryLinkProviderTypePtrOutput {
	return o
}

func (o RepositoryLinkProviderTypePtrOutput) ToRepositoryLinkProviderTypePtrOutputWithContext(ctx context.Context) RepositoryLinkProviderTypePtrOutput {
	return o
}

func (o RepositoryLinkProviderTypePtrOutput) Elem() RepositoryLinkProviderTypeOutput {
	return o.ApplyT(func(v *RepositoryLinkProviderType) RepositoryLinkProviderType {
		if v != nil {
			return *v
		}
		var ret RepositoryLinkProviderType
		return ret
	}).(RepositoryLinkProviderTypeOutput)
}

func (o RepositoryLinkProviderTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RepositoryLinkProviderTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RepositoryLinkProviderType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// The name of the external provider where your third-party code repository is configured.
type SyncConfigurationProviderType string

const (
	SyncConfigurationProviderTypeGitHub            = SyncConfigurationProviderType("GitHub")
	SyncConfigurationProviderTypeBitbucket         = SyncConfigurationProviderType("Bitbucket")
	SyncConfigurationProviderTypeGitHubEnterprise  = SyncConfigurationProviderType("GitHubEnterprise")
	SyncConfigurationProviderTypeGitLab            = SyncConfigurationProviderType("GitLab")
	SyncConfigurationProviderTypeGitLabSelfManaged = SyncConfigurationProviderType("GitLabSelfManaged")
)

type SyncConfigurationProviderTypeOutput struct{ *pulumi.OutputState }

func (SyncConfigurationProviderTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncConfigurationProviderType)(nil)).Elem()
}

func (o SyncConfigurationProviderTypeOutput) ToSyncConfigurationProviderTypeOutput() SyncConfigurationProviderTypeOutput {
	return o
}

func (o SyncConfigurationProviderTypeOutput) ToSyncConfigurationProviderTypeOutputWithContext(ctx context.Context) SyncConfigurationProviderTypeOutput {
	return o
}

func (o SyncConfigurationProviderTypeOutput) ToSyncConfigurationProviderTypePtrOutput() SyncConfigurationProviderTypePtrOutput {
	return o.ToSyncConfigurationProviderTypePtrOutputWithContext(context.Background())
}

func (o SyncConfigurationProviderTypeOutput) ToSyncConfigurationProviderTypePtrOutputWithContext(ctx context.Context) SyncConfigurationProviderTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncConfigurationProviderType) *SyncConfigurationProviderType {
		return &v
	}).(SyncConfigurationProviderTypePtrOutput)
}

func (o SyncConfigurationProviderTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SyncConfigurationProviderTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncConfigurationProviderType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SyncConfigurationProviderTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncConfigurationProviderTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncConfigurationProviderType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncConfigurationProviderTypePtrOutput struct{ *pulumi.OutputState }

func (SyncConfigurationProviderTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncConfigurationProviderType)(nil)).Elem()
}

func (o SyncConfigurationProviderTypePtrOutput) ToSyncConfigurationProviderTypePtrOutput() SyncConfigurationProviderTypePtrOutput {
	return o
}

func (o SyncConfigurationProviderTypePtrOutput) ToSyncConfigurationProviderTypePtrOutputWithContext(ctx context.Context) SyncConfigurationProviderTypePtrOutput {
	return o
}

func (o SyncConfigurationProviderTypePtrOutput) Elem() SyncConfigurationProviderTypeOutput {
	return o.ApplyT(func(v *SyncConfigurationProviderType) SyncConfigurationProviderType {
		if v != nil {
			return *v
		}
		var ret SyncConfigurationProviderType
		return ret
	}).(SyncConfigurationProviderTypeOutput)
}

func (o SyncConfigurationProviderTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncConfigurationProviderTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SyncConfigurationProviderType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// Whether to enable or disable publishing of deployment status to source providers.
type SyncConfigurationPublishDeploymentStatus string

const (
	SyncConfigurationPublishDeploymentStatusEnabled  = SyncConfigurationPublishDeploymentStatus("ENABLED")
	SyncConfigurationPublishDeploymentStatusDisabled = SyncConfigurationPublishDeploymentStatus("DISABLED")
)

func (SyncConfigurationPublishDeploymentStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncConfigurationPublishDeploymentStatus)(nil)).Elem()
}

func (e SyncConfigurationPublishDeploymentStatus) ToSyncConfigurationPublishDeploymentStatusOutput() SyncConfigurationPublishDeploymentStatusOutput {
	return pulumi.ToOutput(e).(SyncConfigurationPublishDeploymentStatusOutput)
}

func (e SyncConfigurationPublishDeploymentStatus) ToSyncConfigurationPublishDeploymentStatusOutputWithContext(ctx context.Context) SyncConfigurationPublishDeploymentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SyncConfigurationPublishDeploymentStatusOutput)
}

func (e SyncConfigurationPublishDeploymentStatus) ToSyncConfigurationPublishDeploymentStatusPtrOutput() SyncConfigurationPublishDeploymentStatusPtrOutput {
	return e.ToSyncConfigurationPublishDeploymentStatusPtrOutputWithContext(context.Background())
}

func (e SyncConfigurationPublishDeploymentStatus) ToSyncConfigurationPublishDeploymentStatusPtrOutputWithContext(ctx context.Context) SyncConfigurationPublishDeploymentStatusPtrOutput {
	return SyncConfigurationPublishDeploymentStatus(e).ToSyncConfigurationPublishDeploymentStatusOutputWithContext(ctx).ToSyncConfigurationPublishDeploymentStatusPtrOutputWithContext(ctx)
}

func (e SyncConfigurationPublishDeploymentStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncConfigurationPublishDeploymentStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncConfigurationPublishDeploymentStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SyncConfigurationPublishDeploymentStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SyncConfigurationPublishDeploymentStatusOutput struct{ *pulumi.OutputState }

func (SyncConfigurationPublishDeploymentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncConfigurationPublishDeploymentStatus)(nil)).Elem()
}

func (o SyncConfigurationPublishDeploymentStatusOutput) ToSyncConfigurationPublishDeploymentStatusOutput() SyncConfigurationPublishDeploymentStatusOutput {
	return o
}

func (o SyncConfigurationPublishDeploymentStatusOutput) ToSyncConfigurationPublishDeploymentStatusOutputWithContext(ctx context.Context) SyncConfigurationPublishDeploymentStatusOutput {
	return o
}

func (o SyncConfigurationPublishDeploymentStatusOutput) ToSyncConfigurationPublishDeploymentStatusPtrOutput() SyncConfigurationPublishDeploymentStatusPtrOutput {
	return o.ToSyncConfigurationPublishDeploymentStatusPtrOutputWithContext(context.Background())
}

func (o SyncConfigurationPublishDeploymentStatusOutput) ToSyncConfigurationPublishDeploymentStatusPtrOutputWithContext(ctx context.Context) SyncConfigurationPublishDeploymentStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncConfigurationPublishDeploymentStatus) *SyncConfigurationPublishDeploymentStatus {
		return &v
	}).(SyncConfigurationPublishDeploymentStatusPtrOutput)
}

func (o SyncConfigurationPublishDeploymentStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SyncConfigurationPublishDeploymentStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncConfigurationPublishDeploymentStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SyncConfigurationPublishDeploymentStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncConfigurationPublishDeploymentStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncConfigurationPublishDeploymentStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncConfigurationPublishDeploymentStatusPtrOutput struct{ *pulumi.OutputState }

func (SyncConfigurationPublishDeploymentStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncConfigurationPublishDeploymentStatus)(nil)).Elem()
}

func (o SyncConfigurationPublishDeploymentStatusPtrOutput) ToSyncConfigurationPublishDeploymentStatusPtrOutput() SyncConfigurationPublishDeploymentStatusPtrOutput {
	return o
}

func (o SyncConfigurationPublishDeploymentStatusPtrOutput) ToSyncConfigurationPublishDeploymentStatusPtrOutputWithContext(ctx context.Context) SyncConfigurationPublishDeploymentStatusPtrOutput {
	return o
}

func (o SyncConfigurationPublishDeploymentStatusPtrOutput) Elem() SyncConfigurationPublishDeploymentStatusOutput {
	return o.ApplyT(func(v *SyncConfigurationPublishDeploymentStatus) SyncConfigurationPublishDeploymentStatus {
		if v != nil {
			return *v
		}
		var ret SyncConfigurationPublishDeploymentStatus
		return ret
	}).(SyncConfigurationPublishDeploymentStatusOutput)
}

func (o SyncConfigurationPublishDeploymentStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncConfigurationPublishDeploymentStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SyncConfigurationPublishDeploymentStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SyncConfigurationPublishDeploymentStatusInput is an input type that accepts values of the SyncConfigurationPublishDeploymentStatus enum
// A concrete instance of `SyncConfigurationPublishDeploymentStatusInput` can be one of the following:
//
//	SyncConfigurationPublishDeploymentStatusEnabled
//	SyncConfigurationPublishDeploymentStatusDisabled
type SyncConfigurationPublishDeploymentStatusInput interface {
	pulumi.Input

	ToSyncConfigurationPublishDeploymentStatusOutput() SyncConfigurationPublishDeploymentStatusOutput
	ToSyncConfigurationPublishDeploymentStatusOutputWithContext(context.Context) SyncConfigurationPublishDeploymentStatusOutput
}

var syncConfigurationPublishDeploymentStatusPtrType = reflect.TypeOf((**SyncConfigurationPublishDeploymentStatus)(nil)).Elem()

type SyncConfigurationPublishDeploymentStatusPtrInput interface {
	pulumi.Input

	ToSyncConfigurationPublishDeploymentStatusPtrOutput() SyncConfigurationPublishDeploymentStatusPtrOutput
	ToSyncConfigurationPublishDeploymentStatusPtrOutputWithContext(context.Context) SyncConfigurationPublishDeploymentStatusPtrOutput
}

type syncConfigurationPublishDeploymentStatusPtr string

func SyncConfigurationPublishDeploymentStatusPtr(v string) SyncConfigurationPublishDeploymentStatusPtrInput {
	return (*syncConfigurationPublishDeploymentStatusPtr)(&v)
}

func (*syncConfigurationPublishDeploymentStatusPtr) ElementType() reflect.Type {
	return syncConfigurationPublishDeploymentStatusPtrType
}

func (in *syncConfigurationPublishDeploymentStatusPtr) ToSyncConfigurationPublishDeploymentStatusPtrOutput() SyncConfigurationPublishDeploymentStatusPtrOutput {
	return pulumi.ToOutput(in).(SyncConfigurationPublishDeploymentStatusPtrOutput)
}

func (in *syncConfigurationPublishDeploymentStatusPtr) ToSyncConfigurationPublishDeploymentStatusPtrOutputWithContext(ctx context.Context) SyncConfigurationPublishDeploymentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SyncConfigurationPublishDeploymentStatusPtrOutput)
}

// When to trigger Git sync to begin the stack update.
type SyncConfigurationTriggerResourceUpdateOn string

const (
	SyncConfigurationTriggerResourceUpdateOnAnyChange  = SyncConfigurationTriggerResourceUpdateOn("ANY_CHANGE")
	SyncConfigurationTriggerResourceUpdateOnFileChange = SyncConfigurationTriggerResourceUpdateOn("FILE_CHANGE")
)

func (SyncConfigurationTriggerResourceUpdateOn) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncConfigurationTriggerResourceUpdateOn)(nil)).Elem()
}

func (e SyncConfigurationTriggerResourceUpdateOn) ToSyncConfigurationTriggerResourceUpdateOnOutput() SyncConfigurationTriggerResourceUpdateOnOutput {
	return pulumi.ToOutput(e).(SyncConfigurationTriggerResourceUpdateOnOutput)
}

func (e SyncConfigurationTriggerResourceUpdateOn) ToSyncConfigurationTriggerResourceUpdateOnOutputWithContext(ctx context.Context) SyncConfigurationTriggerResourceUpdateOnOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SyncConfigurationTriggerResourceUpdateOnOutput)
}

func (e SyncConfigurationTriggerResourceUpdateOn) ToSyncConfigurationTriggerResourceUpdateOnPtrOutput() SyncConfigurationTriggerResourceUpdateOnPtrOutput {
	return e.ToSyncConfigurationTriggerResourceUpdateOnPtrOutputWithContext(context.Background())
}

func (e SyncConfigurationTriggerResourceUpdateOn) ToSyncConfigurationTriggerResourceUpdateOnPtrOutputWithContext(ctx context.Context) SyncConfigurationTriggerResourceUpdateOnPtrOutput {
	return SyncConfigurationTriggerResourceUpdateOn(e).ToSyncConfigurationTriggerResourceUpdateOnOutputWithContext(ctx).ToSyncConfigurationTriggerResourceUpdateOnPtrOutputWithContext(ctx)
}

func (e SyncConfigurationTriggerResourceUpdateOn) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncConfigurationTriggerResourceUpdateOn) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncConfigurationTriggerResourceUpdateOn) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SyncConfigurationTriggerResourceUpdateOn) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SyncConfigurationTriggerResourceUpdateOnOutput struct{ *pulumi.OutputState }

func (SyncConfigurationTriggerResourceUpdateOnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncConfigurationTriggerResourceUpdateOn)(nil)).Elem()
}

func (o SyncConfigurationTriggerResourceUpdateOnOutput) ToSyncConfigurationTriggerResourceUpdateOnOutput() SyncConfigurationTriggerResourceUpdateOnOutput {
	return o
}

func (o SyncConfigurationTriggerResourceUpdateOnOutput) ToSyncConfigurationTriggerResourceUpdateOnOutputWithContext(ctx context.Context) SyncConfigurationTriggerResourceUpdateOnOutput {
	return o
}

func (o SyncConfigurationTriggerResourceUpdateOnOutput) ToSyncConfigurationTriggerResourceUpdateOnPtrOutput() SyncConfigurationTriggerResourceUpdateOnPtrOutput {
	return o.ToSyncConfigurationTriggerResourceUpdateOnPtrOutputWithContext(context.Background())
}

func (o SyncConfigurationTriggerResourceUpdateOnOutput) ToSyncConfigurationTriggerResourceUpdateOnPtrOutputWithContext(ctx context.Context) SyncConfigurationTriggerResourceUpdateOnPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncConfigurationTriggerResourceUpdateOn) *SyncConfigurationTriggerResourceUpdateOn {
		return &v
	}).(SyncConfigurationTriggerResourceUpdateOnPtrOutput)
}

func (o SyncConfigurationTriggerResourceUpdateOnOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SyncConfigurationTriggerResourceUpdateOnOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncConfigurationTriggerResourceUpdateOn) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SyncConfigurationTriggerResourceUpdateOnOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncConfigurationTriggerResourceUpdateOnOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncConfigurationTriggerResourceUpdateOn) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncConfigurationTriggerResourceUpdateOnPtrOutput struct{ *pulumi.OutputState }

func (SyncConfigurationTriggerResourceUpdateOnPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncConfigurationTriggerResourceUpdateOn)(nil)).Elem()
}

func (o SyncConfigurationTriggerResourceUpdateOnPtrOutput) ToSyncConfigurationTriggerResourceUpdateOnPtrOutput() SyncConfigurationTriggerResourceUpdateOnPtrOutput {
	return o
}

func (o SyncConfigurationTriggerResourceUpdateOnPtrOutput) ToSyncConfigurationTriggerResourceUpdateOnPtrOutputWithContext(ctx context.Context) SyncConfigurationTriggerResourceUpdateOnPtrOutput {
	return o
}

func (o SyncConfigurationTriggerResourceUpdateOnPtrOutput) Elem() SyncConfigurationTriggerResourceUpdateOnOutput {
	return o.ApplyT(func(v *SyncConfigurationTriggerResourceUpdateOn) SyncConfigurationTriggerResourceUpdateOn {
		if v != nil {
			return *v
		}
		var ret SyncConfigurationTriggerResourceUpdateOn
		return ret
	}).(SyncConfigurationTriggerResourceUpdateOnOutput)
}

func (o SyncConfigurationTriggerResourceUpdateOnPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncConfigurationTriggerResourceUpdateOnPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SyncConfigurationTriggerResourceUpdateOn) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SyncConfigurationTriggerResourceUpdateOnInput is an input type that accepts values of the SyncConfigurationTriggerResourceUpdateOn enum
// A concrete instance of `SyncConfigurationTriggerResourceUpdateOnInput` can be one of the following:
//
//	SyncConfigurationTriggerResourceUpdateOnAnyChange
//	SyncConfigurationTriggerResourceUpdateOnFileChange
type SyncConfigurationTriggerResourceUpdateOnInput interface {
	pulumi.Input

	ToSyncConfigurationTriggerResourceUpdateOnOutput() SyncConfigurationTriggerResourceUpdateOnOutput
	ToSyncConfigurationTriggerResourceUpdateOnOutputWithContext(context.Context) SyncConfigurationTriggerResourceUpdateOnOutput
}

var syncConfigurationTriggerResourceUpdateOnPtrType = reflect.TypeOf((**SyncConfigurationTriggerResourceUpdateOn)(nil)).Elem()

type SyncConfigurationTriggerResourceUpdateOnPtrInput interface {
	pulumi.Input

	ToSyncConfigurationTriggerResourceUpdateOnPtrOutput() SyncConfigurationTriggerResourceUpdateOnPtrOutput
	ToSyncConfigurationTriggerResourceUpdateOnPtrOutputWithContext(context.Context) SyncConfigurationTriggerResourceUpdateOnPtrOutput
}

type syncConfigurationTriggerResourceUpdateOnPtr string

func SyncConfigurationTriggerResourceUpdateOnPtr(v string) SyncConfigurationTriggerResourceUpdateOnPtrInput {
	return (*syncConfigurationTriggerResourceUpdateOnPtr)(&v)
}

func (*syncConfigurationTriggerResourceUpdateOnPtr) ElementType() reflect.Type {
	return syncConfigurationTriggerResourceUpdateOnPtrType
}

func (in *syncConfigurationTriggerResourceUpdateOnPtr) ToSyncConfigurationTriggerResourceUpdateOnPtrOutput() SyncConfigurationTriggerResourceUpdateOnPtrOutput {
	return pulumi.ToOutput(in).(SyncConfigurationTriggerResourceUpdateOnPtrOutput)
}

func (in *syncConfigurationTriggerResourceUpdateOnPtr) ToSyncConfigurationTriggerResourceUpdateOnPtrOutputWithContext(ctx context.Context) SyncConfigurationTriggerResourceUpdateOnPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SyncConfigurationTriggerResourceUpdateOnPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncConfigurationPublishDeploymentStatusInput)(nil)).Elem(), SyncConfigurationPublishDeploymentStatus("ENABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*SyncConfigurationPublishDeploymentStatusPtrInput)(nil)).Elem(), SyncConfigurationPublishDeploymentStatus("ENABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*SyncConfigurationTriggerResourceUpdateOnInput)(nil)).Elem(), SyncConfigurationTriggerResourceUpdateOn("ANY_CHANGE"))
	pulumi.RegisterInputType(reflect.TypeOf((*SyncConfigurationTriggerResourceUpdateOnPtrInput)(nil)).Elem(), SyncConfigurationTriggerResourceUpdateOn("ANY_CHANGE"))
	pulumi.RegisterOutputType(RepositoryLinkProviderTypeOutput{})
	pulumi.RegisterOutputType(RepositoryLinkProviderTypePtrOutput{})
	pulumi.RegisterOutputType(SyncConfigurationProviderTypeOutput{})
	pulumi.RegisterOutputType(SyncConfigurationProviderTypePtrOutput{})
	pulumi.RegisterOutputType(SyncConfigurationPublishDeploymentStatusOutput{})
	pulumi.RegisterOutputType(SyncConfigurationPublishDeploymentStatusPtrOutput{})
	pulumi.RegisterOutputType(SyncConfigurationTriggerResourceUpdateOnOutput{})
	pulumi.RegisterOutputType(SyncConfigurationTriggerResourceUpdateOnPtrOutput{})
}
