// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspaces

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The association status of the connection alias.
type ConnectionAliasAssociationAssociationStatus string

const (
	ConnectionAliasAssociationAssociationStatusNotAssociated               = ConnectionAliasAssociationAssociationStatus("NOT_ASSOCIATED")
	ConnectionAliasAssociationAssociationStatusPendingAssociation          = ConnectionAliasAssociationAssociationStatus("PENDING_ASSOCIATION")
	ConnectionAliasAssociationAssociationStatusAssociatedWithOwnerAccount  = ConnectionAliasAssociationAssociationStatus("ASSOCIATED_WITH_OWNER_ACCOUNT")
	ConnectionAliasAssociationAssociationStatusAssociatedWithSharedAccount = ConnectionAliasAssociationAssociationStatus("ASSOCIATED_WITH_SHARED_ACCOUNT")
	ConnectionAliasAssociationAssociationStatusPendingDisassociation       = ConnectionAliasAssociationAssociationStatus("PENDING_DISASSOCIATION")
)

type ConnectionAliasAssociationAssociationStatusOutput struct{ *pulumi.OutputState }

func (ConnectionAliasAssociationAssociationStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionAliasAssociationAssociationStatus)(nil)).Elem()
}

func (o ConnectionAliasAssociationAssociationStatusOutput) ToConnectionAliasAssociationAssociationStatusOutput() ConnectionAliasAssociationAssociationStatusOutput {
	return o
}

func (o ConnectionAliasAssociationAssociationStatusOutput) ToConnectionAliasAssociationAssociationStatusOutputWithContext(ctx context.Context) ConnectionAliasAssociationAssociationStatusOutput {
	return o
}

func (o ConnectionAliasAssociationAssociationStatusOutput) ToConnectionAliasAssociationAssociationStatusPtrOutput() ConnectionAliasAssociationAssociationStatusPtrOutput {
	return o.ToConnectionAliasAssociationAssociationStatusPtrOutputWithContext(context.Background())
}

func (o ConnectionAliasAssociationAssociationStatusOutput) ToConnectionAliasAssociationAssociationStatusPtrOutputWithContext(ctx context.Context) ConnectionAliasAssociationAssociationStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionAliasAssociationAssociationStatus) *ConnectionAliasAssociationAssociationStatus {
		return &v
	}).(ConnectionAliasAssociationAssociationStatusPtrOutput)
}

func (o ConnectionAliasAssociationAssociationStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionAliasAssociationAssociationStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionAliasAssociationAssociationStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionAliasAssociationAssociationStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionAliasAssociationAssociationStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionAliasAssociationAssociationStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionAliasAssociationAssociationStatusPtrOutput struct{ *pulumi.OutputState }

func (ConnectionAliasAssociationAssociationStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionAliasAssociationAssociationStatus)(nil)).Elem()
}

func (o ConnectionAliasAssociationAssociationStatusPtrOutput) ToConnectionAliasAssociationAssociationStatusPtrOutput() ConnectionAliasAssociationAssociationStatusPtrOutput {
	return o
}

func (o ConnectionAliasAssociationAssociationStatusPtrOutput) ToConnectionAliasAssociationAssociationStatusPtrOutputWithContext(ctx context.Context) ConnectionAliasAssociationAssociationStatusPtrOutput {
	return o
}

func (o ConnectionAliasAssociationAssociationStatusPtrOutput) Elem() ConnectionAliasAssociationAssociationStatusOutput {
	return o.ApplyT(func(v *ConnectionAliasAssociationAssociationStatus) ConnectionAliasAssociationAssociationStatus {
		if v != nil {
			return *v
		}
		var ret ConnectionAliasAssociationAssociationStatus
		return ret
	}).(ConnectionAliasAssociationAssociationStatusOutput)
}

func (o ConnectionAliasAssociationAssociationStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionAliasAssociationAssociationStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionAliasAssociationAssociationStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// The current state of the connection alias, returned as a string.
type ConnectionAliasStateEnum string

const (
	ConnectionAliasStateEnumCreating = ConnectionAliasStateEnum("CREATING")
	ConnectionAliasStateEnumCreated  = ConnectionAliasStateEnum("CREATED")
	ConnectionAliasStateEnumDeleting = ConnectionAliasStateEnum("DELETING")
)

type ConnectionAliasStateEnumOutput struct{ *pulumi.OutputState }

func (ConnectionAliasStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionAliasStateEnum)(nil)).Elem()
}

func (o ConnectionAliasStateEnumOutput) ToConnectionAliasStateEnumOutput() ConnectionAliasStateEnumOutput {
	return o
}

func (o ConnectionAliasStateEnumOutput) ToConnectionAliasStateEnumOutputWithContext(ctx context.Context) ConnectionAliasStateEnumOutput {
	return o
}

func (o ConnectionAliasStateEnumOutput) ToConnectionAliasStateEnumPtrOutput() ConnectionAliasStateEnumPtrOutput {
	return o.ToConnectionAliasStateEnumPtrOutputWithContext(context.Background())
}

func (o ConnectionAliasStateEnumOutput) ToConnectionAliasStateEnumPtrOutputWithContext(ctx context.Context) ConnectionAliasStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionAliasStateEnum) *ConnectionAliasStateEnum {
		return &v
	}).(ConnectionAliasStateEnumPtrOutput)
}

func (o ConnectionAliasStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionAliasStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionAliasStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionAliasStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionAliasStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionAliasStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionAliasStateEnumPtrOutput struct{ *pulumi.OutputState }

func (ConnectionAliasStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionAliasStateEnum)(nil)).Elem()
}

func (o ConnectionAliasStateEnumPtrOutput) ToConnectionAliasStateEnumPtrOutput() ConnectionAliasStateEnumPtrOutput {
	return o
}

func (o ConnectionAliasStateEnumPtrOutput) ToConnectionAliasStateEnumPtrOutputWithContext(ctx context.Context) ConnectionAliasStateEnumPtrOutput {
	return o
}

func (o ConnectionAliasStateEnumPtrOutput) Elem() ConnectionAliasStateEnumOutput {
	return o.ApplyT(func(v *ConnectionAliasStateEnum) ConnectionAliasStateEnum {
		if v != nil {
			return *v
		}
		var ret ConnectionAliasStateEnum
		return ret
	}).(ConnectionAliasStateEnumOutput)
}

func (o ConnectionAliasStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionAliasStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionAliasStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkspacesPoolApplicationSettingsStatus string

const (
	WorkspacesPoolApplicationSettingsStatusDisabled = WorkspacesPoolApplicationSettingsStatus("DISABLED")
	WorkspacesPoolApplicationSettingsStatusEnabled  = WorkspacesPoolApplicationSettingsStatus("ENABLED")
)

func (WorkspacesPoolApplicationSettingsStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspacesPoolApplicationSettingsStatus)(nil)).Elem()
}

func (e WorkspacesPoolApplicationSettingsStatus) ToWorkspacesPoolApplicationSettingsStatusOutput() WorkspacesPoolApplicationSettingsStatusOutput {
	return pulumi.ToOutput(e).(WorkspacesPoolApplicationSettingsStatusOutput)
}

func (e WorkspacesPoolApplicationSettingsStatus) ToWorkspacesPoolApplicationSettingsStatusOutputWithContext(ctx context.Context) WorkspacesPoolApplicationSettingsStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkspacesPoolApplicationSettingsStatusOutput)
}

func (e WorkspacesPoolApplicationSettingsStatus) ToWorkspacesPoolApplicationSettingsStatusPtrOutput() WorkspacesPoolApplicationSettingsStatusPtrOutput {
	return e.ToWorkspacesPoolApplicationSettingsStatusPtrOutputWithContext(context.Background())
}

func (e WorkspacesPoolApplicationSettingsStatus) ToWorkspacesPoolApplicationSettingsStatusPtrOutputWithContext(ctx context.Context) WorkspacesPoolApplicationSettingsStatusPtrOutput {
	return WorkspacesPoolApplicationSettingsStatus(e).ToWorkspacesPoolApplicationSettingsStatusOutputWithContext(ctx).ToWorkspacesPoolApplicationSettingsStatusPtrOutputWithContext(ctx)
}

func (e WorkspacesPoolApplicationSettingsStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspacesPoolApplicationSettingsStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspacesPoolApplicationSettingsStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkspacesPoolApplicationSettingsStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkspacesPoolApplicationSettingsStatusOutput struct{ *pulumi.OutputState }

func (WorkspacesPoolApplicationSettingsStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspacesPoolApplicationSettingsStatus)(nil)).Elem()
}

func (o WorkspacesPoolApplicationSettingsStatusOutput) ToWorkspacesPoolApplicationSettingsStatusOutput() WorkspacesPoolApplicationSettingsStatusOutput {
	return o
}

func (o WorkspacesPoolApplicationSettingsStatusOutput) ToWorkspacesPoolApplicationSettingsStatusOutputWithContext(ctx context.Context) WorkspacesPoolApplicationSettingsStatusOutput {
	return o
}

func (o WorkspacesPoolApplicationSettingsStatusOutput) ToWorkspacesPoolApplicationSettingsStatusPtrOutput() WorkspacesPoolApplicationSettingsStatusPtrOutput {
	return o.ToWorkspacesPoolApplicationSettingsStatusPtrOutputWithContext(context.Background())
}

func (o WorkspacesPoolApplicationSettingsStatusOutput) ToWorkspacesPoolApplicationSettingsStatusPtrOutputWithContext(ctx context.Context) WorkspacesPoolApplicationSettingsStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspacesPoolApplicationSettingsStatus) *WorkspacesPoolApplicationSettingsStatus {
		return &v
	}).(WorkspacesPoolApplicationSettingsStatusPtrOutput)
}

func (o WorkspacesPoolApplicationSettingsStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkspacesPoolApplicationSettingsStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspacesPoolApplicationSettingsStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkspacesPoolApplicationSettingsStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspacesPoolApplicationSettingsStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspacesPoolApplicationSettingsStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkspacesPoolApplicationSettingsStatusPtrOutput struct{ *pulumi.OutputState }

func (WorkspacesPoolApplicationSettingsStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacesPoolApplicationSettingsStatus)(nil)).Elem()
}

func (o WorkspacesPoolApplicationSettingsStatusPtrOutput) ToWorkspacesPoolApplicationSettingsStatusPtrOutput() WorkspacesPoolApplicationSettingsStatusPtrOutput {
	return o
}

func (o WorkspacesPoolApplicationSettingsStatusPtrOutput) ToWorkspacesPoolApplicationSettingsStatusPtrOutputWithContext(ctx context.Context) WorkspacesPoolApplicationSettingsStatusPtrOutput {
	return o
}

func (o WorkspacesPoolApplicationSettingsStatusPtrOutput) Elem() WorkspacesPoolApplicationSettingsStatusOutput {
	return o.ApplyT(func(v *WorkspacesPoolApplicationSettingsStatus) WorkspacesPoolApplicationSettingsStatus {
		if v != nil {
			return *v
		}
		var ret WorkspacesPoolApplicationSettingsStatus
		return ret
	}).(WorkspacesPoolApplicationSettingsStatusOutput)
}

func (o WorkspacesPoolApplicationSettingsStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspacesPoolApplicationSettingsStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkspacesPoolApplicationSettingsStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// WorkspacesPoolApplicationSettingsStatusInput is an input type that accepts values of the WorkspacesPoolApplicationSettingsStatus enum
// A concrete instance of `WorkspacesPoolApplicationSettingsStatusInput` can be one of the following:
//
//	WorkspacesPoolApplicationSettingsStatusDisabled
//	WorkspacesPoolApplicationSettingsStatusEnabled
type WorkspacesPoolApplicationSettingsStatusInput interface {
	pulumi.Input

	ToWorkspacesPoolApplicationSettingsStatusOutput() WorkspacesPoolApplicationSettingsStatusOutput
	ToWorkspacesPoolApplicationSettingsStatusOutputWithContext(context.Context) WorkspacesPoolApplicationSettingsStatusOutput
}

var workspacesPoolApplicationSettingsStatusPtrType = reflect.TypeOf((**WorkspacesPoolApplicationSettingsStatus)(nil)).Elem()

type WorkspacesPoolApplicationSettingsStatusPtrInput interface {
	pulumi.Input

	ToWorkspacesPoolApplicationSettingsStatusPtrOutput() WorkspacesPoolApplicationSettingsStatusPtrOutput
	ToWorkspacesPoolApplicationSettingsStatusPtrOutputWithContext(context.Context) WorkspacesPoolApplicationSettingsStatusPtrOutput
}

type workspacesPoolApplicationSettingsStatusPtr string

func WorkspacesPoolApplicationSettingsStatusPtr(v string) WorkspacesPoolApplicationSettingsStatusPtrInput {
	return (*workspacesPoolApplicationSettingsStatusPtr)(&v)
}

func (*workspacesPoolApplicationSettingsStatusPtr) ElementType() reflect.Type {
	return workspacesPoolApplicationSettingsStatusPtrType
}

func (in *workspacesPoolApplicationSettingsStatusPtr) ToWorkspacesPoolApplicationSettingsStatusPtrOutput() WorkspacesPoolApplicationSettingsStatusPtrOutput {
	return pulumi.ToOutput(in).(WorkspacesPoolApplicationSettingsStatusPtrOutput)
}

func (in *workspacesPoolApplicationSettingsStatusPtr) ToWorkspacesPoolApplicationSettingsStatusPtrOutputWithContext(ctx context.Context) WorkspacesPoolApplicationSettingsStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkspacesPoolApplicationSettingsStatusPtrOutput)
}

type WorkspacesPoolRunningMode string

const (
	WorkspacesPoolRunningModeAlwaysOn = WorkspacesPoolRunningMode("ALWAYS_ON")
	WorkspacesPoolRunningModeAutoStop = WorkspacesPoolRunningMode("AUTO_STOP")
)

func (WorkspacesPoolRunningMode) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspacesPoolRunningMode)(nil)).Elem()
}

func (e WorkspacesPoolRunningMode) ToWorkspacesPoolRunningModeOutput() WorkspacesPoolRunningModeOutput {
	return pulumi.ToOutput(e).(WorkspacesPoolRunningModeOutput)
}

func (e WorkspacesPoolRunningMode) ToWorkspacesPoolRunningModeOutputWithContext(ctx context.Context) WorkspacesPoolRunningModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkspacesPoolRunningModeOutput)
}

func (e WorkspacesPoolRunningMode) ToWorkspacesPoolRunningModePtrOutput() WorkspacesPoolRunningModePtrOutput {
	return e.ToWorkspacesPoolRunningModePtrOutputWithContext(context.Background())
}

func (e WorkspacesPoolRunningMode) ToWorkspacesPoolRunningModePtrOutputWithContext(ctx context.Context) WorkspacesPoolRunningModePtrOutput {
	return WorkspacesPoolRunningMode(e).ToWorkspacesPoolRunningModeOutputWithContext(ctx).ToWorkspacesPoolRunningModePtrOutputWithContext(ctx)
}

func (e WorkspacesPoolRunningMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspacesPoolRunningMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspacesPoolRunningMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkspacesPoolRunningMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkspacesPoolRunningModeOutput struct{ *pulumi.OutputState }

func (WorkspacesPoolRunningModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspacesPoolRunningMode)(nil)).Elem()
}

func (o WorkspacesPoolRunningModeOutput) ToWorkspacesPoolRunningModeOutput() WorkspacesPoolRunningModeOutput {
	return o
}

func (o WorkspacesPoolRunningModeOutput) ToWorkspacesPoolRunningModeOutputWithContext(ctx context.Context) WorkspacesPoolRunningModeOutput {
	return o
}

func (o WorkspacesPoolRunningModeOutput) ToWorkspacesPoolRunningModePtrOutput() WorkspacesPoolRunningModePtrOutput {
	return o.ToWorkspacesPoolRunningModePtrOutputWithContext(context.Background())
}

func (o WorkspacesPoolRunningModeOutput) ToWorkspacesPoolRunningModePtrOutputWithContext(ctx context.Context) WorkspacesPoolRunningModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspacesPoolRunningMode) *WorkspacesPoolRunningMode {
		return &v
	}).(WorkspacesPoolRunningModePtrOutput)
}

func (o WorkspacesPoolRunningModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkspacesPoolRunningModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspacesPoolRunningMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkspacesPoolRunningModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspacesPoolRunningModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspacesPoolRunningMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkspacesPoolRunningModePtrOutput struct{ *pulumi.OutputState }

func (WorkspacesPoolRunningModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacesPoolRunningMode)(nil)).Elem()
}

func (o WorkspacesPoolRunningModePtrOutput) ToWorkspacesPoolRunningModePtrOutput() WorkspacesPoolRunningModePtrOutput {
	return o
}

func (o WorkspacesPoolRunningModePtrOutput) ToWorkspacesPoolRunningModePtrOutputWithContext(ctx context.Context) WorkspacesPoolRunningModePtrOutput {
	return o
}

func (o WorkspacesPoolRunningModePtrOutput) Elem() WorkspacesPoolRunningModeOutput {
	return o.ApplyT(func(v *WorkspacesPoolRunningMode) WorkspacesPoolRunningMode {
		if v != nil {
			return *v
		}
		var ret WorkspacesPoolRunningMode
		return ret
	}).(WorkspacesPoolRunningModeOutput)
}

func (o WorkspacesPoolRunningModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspacesPoolRunningModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkspacesPoolRunningMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// WorkspacesPoolRunningModeInput is an input type that accepts values of the WorkspacesPoolRunningMode enum
// A concrete instance of `WorkspacesPoolRunningModeInput` can be one of the following:
//
//	WorkspacesPoolRunningModeAlwaysOn
//	WorkspacesPoolRunningModeAutoStop
type WorkspacesPoolRunningModeInput interface {
	pulumi.Input

	ToWorkspacesPoolRunningModeOutput() WorkspacesPoolRunningModeOutput
	ToWorkspacesPoolRunningModeOutputWithContext(context.Context) WorkspacesPoolRunningModeOutput
}

var workspacesPoolRunningModePtrType = reflect.TypeOf((**WorkspacesPoolRunningMode)(nil)).Elem()

type WorkspacesPoolRunningModePtrInput interface {
	pulumi.Input

	ToWorkspacesPoolRunningModePtrOutput() WorkspacesPoolRunningModePtrOutput
	ToWorkspacesPoolRunningModePtrOutputWithContext(context.Context) WorkspacesPoolRunningModePtrOutput
}

type workspacesPoolRunningModePtr string

func WorkspacesPoolRunningModePtr(v string) WorkspacesPoolRunningModePtrInput {
	return (*workspacesPoolRunningModePtr)(&v)
}

func (*workspacesPoolRunningModePtr) ElementType() reflect.Type {
	return workspacesPoolRunningModePtrType
}

func (in *workspacesPoolRunningModePtr) ToWorkspacesPoolRunningModePtrOutput() WorkspacesPoolRunningModePtrOutput {
	return pulumi.ToOutput(in).(WorkspacesPoolRunningModePtrOutput)
}

func (in *workspacesPoolRunningModePtr) ToWorkspacesPoolRunningModePtrOutputWithContext(ctx context.Context) WorkspacesPoolRunningModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkspacesPoolRunningModePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspacesPoolApplicationSettingsStatusInput)(nil)).Elem(), WorkspacesPoolApplicationSettingsStatus("DISABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspacesPoolApplicationSettingsStatusPtrInput)(nil)).Elem(), WorkspacesPoolApplicationSettingsStatus("DISABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspacesPoolRunningModeInput)(nil)).Elem(), WorkspacesPoolRunningMode("ALWAYS_ON"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspacesPoolRunningModePtrInput)(nil)).Elem(), WorkspacesPoolRunningMode("ALWAYS_ON"))
	pulumi.RegisterOutputType(ConnectionAliasAssociationAssociationStatusOutput{})
	pulumi.RegisterOutputType(ConnectionAliasAssociationAssociationStatusPtrOutput{})
	pulumi.RegisterOutputType(ConnectionAliasStateEnumOutput{})
	pulumi.RegisterOutputType(ConnectionAliasStateEnumPtrOutput{})
	pulumi.RegisterOutputType(WorkspacesPoolApplicationSettingsStatusOutput{})
	pulumi.RegisterOutputType(WorkspacesPoolApplicationSettingsStatusPtrOutput{})
	pulumi.RegisterOutputType(WorkspacesPoolRunningModeOutput{})
	pulumi.RegisterOutputType(WorkspacesPoolRunningModePtrOutput{})
}
