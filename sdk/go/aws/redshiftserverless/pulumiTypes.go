// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshiftserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NamespaceType struct {
	AdminUsername     *string              `pulumi:"adminUsername"`
	CreationDate      *string              `pulumi:"creationDate"`
	DbName            *string              `pulumi:"dbName"`
	DefaultIamRoleArn *string              `pulumi:"defaultIamRoleArn"`
	IamRoles          []string             `pulumi:"iamRoles"`
	KmsKeyId          *string              `pulumi:"kmsKeyId"`
	LogExports        []NamespaceLogExport `pulumi:"logExports"`
	NamespaceArn      *string              `pulumi:"namespaceArn"`
	NamespaceId       *string              `pulumi:"namespaceId"`
	NamespaceName     *string              `pulumi:"namespaceName"`
	Status            *NamespaceStatus     `pulumi:"status"`
}

type NamespaceTypeOutput struct{ *pulumi.OutputState }

func (NamespaceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceType)(nil)).Elem()
}

func (o NamespaceTypeOutput) ToNamespaceTypeOutput() NamespaceTypeOutput {
	return o
}

func (o NamespaceTypeOutput) ToNamespaceTypeOutputWithContext(ctx context.Context) NamespaceTypeOutput {
	return o
}

func (o NamespaceTypeOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) DbName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.DbName }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) DefaultIamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.DefaultIamRoleArn }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) IamRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamespaceType) []string { return v.IamRoles }).(pulumi.StringArrayOutput)
}

func (o NamespaceTypeOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) LogExports() NamespaceLogExportArrayOutput {
	return o.ApplyT(func(v NamespaceType) []NamespaceLogExport { return v.LogExports }).(NamespaceLogExportArrayOutput)
}

func (o NamespaceTypeOutput) NamespaceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.NamespaceArn }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) NamespaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.NamespaceName }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) Status() NamespaceStatusPtrOutput {
	return o.ApplyT(func(v NamespaceType) *NamespaceStatus { return v.Status }).(NamespaceStatusPtrOutput)
}

type NamespaceTypePtrOutput struct{ *pulumi.OutputState }

func (NamespaceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceType)(nil)).Elem()
}

func (o NamespaceTypePtrOutput) ToNamespaceTypePtrOutput() NamespaceTypePtrOutput {
	return o
}

func (o NamespaceTypePtrOutput) ToNamespaceTypePtrOutputWithContext(ctx context.Context) NamespaceTypePtrOutput {
	return o
}

func (o NamespaceTypePtrOutput) Elem() NamespaceTypeOutput {
	return o.ApplyT(func(v *NamespaceType) NamespaceType {
		if v != nil {
			return *v
		}
		var ret NamespaceType
		return ret
	}).(NamespaceTypeOutput)
}

func (o NamespaceTypePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.CreationDate
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) DbName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.DbName
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) DefaultIamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.DefaultIamRoleArn
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) IamRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NamespaceType) []string {
		if v == nil {
			return nil
		}
		return v.IamRoles
	}).(pulumi.StringArrayOutput)
}

func (o NamespaceTypePtrOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.KmsKeyId
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) LogExports() NamespaceLogExportArrayOutput {
	return o.ApplyT(func(v *NamespaceType) []NamespaceLogExport {
		if v == nil {
			return nil
		}
		return v.LogExports
	}).(NamespaceLogExportArrayOutput)
}

func (o NamespaceTypePtrOutput) NamespaceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.NamespaceArn
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.NamespaceId
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) NamespaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.NamespaceName
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) Status() NamespaceStatusPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *NamespaceStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(NamespaceStatusPtrOutput)
}

type NamespaceTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// NamespaceTagInput is an input type that accepts NamespaceTagArgs and NamespaceTagOutput values.
// You can construct a concrete instance of `NamespaceTagInput` via:
//
//	NamespaceTagArgs{...}
type NamespaceTagInput interface {
	pulumi.Input

	ToNamespaceTagOutput() NamespaceTagOutput
	ToNamespaceTagOutputWithContext(context.Context) NamespaceTagOutput
}

type NamespaceTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (NamespaceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceTag)(nil)).Elem()
}

func (i NamespaceTagArgs) ToNamespaceTagOutput() NamespaceTagOutput {
	return i.ToNamespaceTagOutputWithContext(context.Background())
}

func (i NamespaceTagArgs) ToNamespaceTagOutputWithContext(ctx context.Context) NamespaceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceTagOutput)
}

// NamespaceTagArrayInput is an input type that accepts NamespaceTagArray and NamespaceTagArrayOutput values.
// You can construct a concrete instance of `NamespaceTagArrayInput` via:
//
//	NamespaceTagArray{ NamespaceTagArgs{...} }
type NamespaceTagArrayInput interface {
	pulumi.Input

	ToNamespaceTagArrayOutput() NamespaceTagArrayOutput
	ToNamespaceTagArrayOutputWithContext(context.Context) NamespaceTagArrayOutput
}

type NamespaceTagArray []NamespaceTagInput

func (NamespaceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceTag)(nil)).Elem()
}

func (i NamespaceTagArray) ToNamespaceTagArrayOutput() NamespaceTagArrayOutput {
	return i.ToNamespaceTagArrayOutputWithContext(context.Background())
}

func (i NamespaceTagArray) ToNamespaceTagArrayOutputWithContext(ctx context.Context) NamespaceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceTagArrayOutput)
}

type NamespaceTagOutput struct{ *pulumi.OutputState }

func (NamespaceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceTag)(nil)).Elem()
}

func (o NamespaceTagOutput) ToNamespaceTagOutput() NamespaceTagOutput {
	return o
}

func (o NamespaceTagOutput) ToNamespaceTagOutputWithContext(ctx context.Context) NamespaceTagOutput {
	return o
}

func (o NamespaceTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v NamespaceTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o NamespaceTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v NamespaceTag) string { return v.Value }).(pulumi.StringOutput)
}

type NamespaceTagArrayOutput struct{ *pulumi.OutputState }

func (NamespaceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceTag)(nil)).Elem()
}

func (o NamespaceTagArrayOutput) ToNamespaceTagArrayOutput() NamespaceTagArrayOutput {
	return o
}

func (o NamespaceTagArrayOutput) ToNamespaceTagArrayOutputWithContext(ctx context.Context) NamespaceTagArrayOutput {
	return o
}

func (o NamespaceTagArrayOutput) Index(i pulumi.IntInput) NamespaceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamespaceTag {
		return vs[0].([]NamespaceTag)[vs[1].(int)]
	}).(NamespaceTagOutput)
}

type WorkgroupType struct {
	BaseCapacity       *int                       `pulumi:"baseCapacity"`
	ConfigParameters   []WorkgroupConfigParameter `pulumi:"configParameters"`
	CreationDate       *string                    `pulumi:"creationDate"`
	Endpoint           *WorkgroupEndpoint         `pulumi:"endpoint"`
	EnhancedVpcRouting *bool                      `pulumi:"enhancedVpcRouting"`
	NamespaceName      *string                    `pulumi:"namespaceName"`
	PubliclyAccessible *bool                      `pulumi:"publiclyAccessible"`
	SecurityGroupIds   []string                   `pulumi:"securityGroupIds"`
	Status             *WorkgroupStatus           `pulumi:"status"`
	SubnetIds          []string                   `pulumi:"subnetIds"`
	WorkgroupArn       *string                    `pulumi:"workgroupArn"`
	WorkgroupId        *string                    `pulumi:"workgroupId"`
	WorkgroupName      *string                    `pulumi:"workgroupName"`
}

type WorkgroupTypeOutput struct{ *pulumi.OutputState }

func (WorkgroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkgroupType)(nil)).Elem()
}

func (o WorkgroupTypeOutput) ToWorkgroupTypeOutput() WorkgroupTypeOutput {
	return o
}

func (o WorkgroupTypeOutput) ToWorkgroupTypeOutputWithContext(ctx context.Context) WorkgroupTypeOutput {
	return o
}

func (o WorkgroupTypeOutput) BaseCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkgroupType) *int { return v.BaseCapacity }).(pulumi.IntPtrOutput)
}

func (o WorkgroupTypeOutput) ConfigParameters() WorkgroupConfigParameterArrayOutput {
	return o.ApplyT(func(v WorkgroupType) []WorkgroupConfigParameter { return v.ConfigParameters }).(WorkgroupConfigParameterArrayOutput)
}

func (o WorkgroupTypeOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupType) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o WorkgroupTypeOutput) Endpoint() WorkgroupEndpointPtrOutput {
	return o.ApplyT(func(v WorkgroupType) *WorkgroupEndpoint { return v.Endpoint }).(WorkgroupEndpointPtrOutput)
}

func (o WorkgroupTypeOutput) EnhancedVpcRouting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkgroupType) *bool { return v.EnhancedVpcRouting }).(pulumi.BoolPtrOutput)
}

func (o WorkgroupTypeOutput) NamespaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupType) *string { return v.NamespaceName }).(pulumi.StringPtrOutput)
}

func (o WorkgroupTypeOutput) PubliclyAccessible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkgroupType) *bool { return v.PubliclyAccessible }).(pulumi.BoolPtrOutput)
}

func (o WorkgroupTypeOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkgroupType) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o WorkgroupTypeOutput) Status() WorkgroupStatusPtrOutput {
	return o.ApplyT(func(v WorkgroupType) *WorkgroupStatus { return v.Status }).(WorkgroupStatusPtrOutput)
}

func (o WorkgroupTypeOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkgroupType) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o WorkgroupTypeOutput) WorkgroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupType) *string { return v.WorkgroupArn }).(pulumi.StringPtrOutput)
}

func (o WorkgroupTypeOutput) WorkgroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupType) *string { return v.WorkgroupId }).(pulumi.StringPtrOutput)
}

func (o WorkgroupTypeOutput) WorkgroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupType) *string { return v.WorkgroupName }).(pulumi.StringPtrOutput)
}

type WorkgroupTypePtrOutput struct{ *pulumi.OutputState }

func (WorkgroupTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkgroupType)(nil)).Elem()
}

func (o WorkgroupTypePtrOutput) ToWorkgroupTypePtrOutput() WorkgroupTypePtrOutput {
	return o
}

func (o WorkgroupTypePtrOutput) ToWorkgroupTypePtrOutputWithContext(ctx context.Context) WorkgroupTypePtrOutput {
	return o
}

func (o WorkgroupTypePtrOutput) Elem() WorkgroupTypeOutput {
	return o.ApplyT(func(v *WorkgroupType) WorkgroupType {
		if v != nil {
			return *v
		}
		var ret WorkgroupType
		return ret
	}).(WorkgroupTypeOutput)
}

func (o WorkgroupTypePtrOutput) BaseCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkgroupType) *int {
		if v == nil {
			return nil
		}
		return v.BaseCapacity
	}).(pulumi.IntPtrOutput)
}

func (o WorkgroupTypePtrOutput) ConfigParameters() WorkgroupConfigParameterArrayOutput {
	return o.ApplyT(func(v *WorkgroupType) []WorkgroupConfigParameter {
		if v == nil {
			return nil
		}
		return v.ConfigParameters
	}).(WorkgroupConfigParameterArrayOutput)
}

func (o WorkgroupTypePtrOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkgroupType) *string {
		if v == nil {
			return nil
		}
		return v.CreationDate
	}).(pulumi.StringPtrOutput)
}

func (o WorkgroupTypePtrOutput) Endpoint() WorkgroupEndpointPtrOutput {
	return o.ApplyT(func(v *WorkgroupType) *WorkgroupEndpoint {
		if v == nil {
			return nil
		}
		return v.Endpoint
	}).(WorkgroupEndpointPtrOutput)
}

func (o WorkgroupTypePtrOutput) EnhancedVpcRouting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkgroupType) *bool {
		if v == nil {
			return nil
		}
		return v.EnhancedVpcRouting
	}).(pulumi.BoolPtrOutput)
}

func (o WorkgroupTypePtrOutput) NamespaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkgroupType) *string {
		if v == nil {
			return nil
		}
		return v.NamespaceName
	}).(pulumi.StringPtrOutput)
}

func (o WorkgroupTypePtrOutput) PubliclyAccessible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkgroupType) *bool {
		if v == nil {
			return nil
		}
		return v.PubliclyAccessible
	}).(pulumi.BoolPtrOutput)
}

func (o WorkgroupTypePtrOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkgroupType) []string {
		if v == nil {
			return nil
		}
		return v.SecurityGroupIds
	}).(pulumi.StringArrayOutput)
}

func (o WorkgroupTypePtrOutput) Status() WorkgroupStatusPtrOutput {
	return o.ApplyT(func(v *WorkgroupType) *WorkgroupStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(WorkgroupStatusPtrOutput)
}

func (o WorkgroupTypePtrOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkgroupType) []string {
		if v == nil {
			return nil
		}
		return v.SubnetIds
	}).(pulumi.StringArrayOutput)
}

func (o WorkgroupTypePtrOutput) WorkgroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkgroupType) *string {
		if v == nil {
			return nil
		}
		return v.WorkgroupArn
	}).(pulumi.StringPtrOutput)
}

func (o WorkgroupTypePtrOutput) WorkgroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkgroupType) *string {
		if v == nil {
			return nil
		}
		return v.WorkgroupId
	}).(pulumi.StringPtrOutput)
}

func (o WorkgroupTypePtrOutput) WorkgroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkgroupType) *string {
		if v == nil {
			return nil
		}
		return v.WorkgroupName
	}).(pulumi.StringPtrOutput)
}

type WorkgroupConfigParameter struct {
	ParameterKey   *string `pulumi:"parameterKey"`
	ParameterValue *string `pulumi:"parameterValue"`
}

// WorkgroupConfigParameterInput is an input type that accepts WorkgroupConfigParameterArgs and WorkgroupConfigParameterOutput values.
// You can construct a concrete instance of `WorkgroupConfigParameterInput` via:
//
//	WorkgroupConfigParameterArgs{...}
type WorkgroupConfigParameterInput interface {
	pulumi.Input

	ToWorkgroupConfigParameterOutput() WorkgroupConfigParameterOutput
	ToWorkgroupConfigParameterOutputWithContext(context.Context) WorkgroupConfigParameterOutput
}

type WorkgroupConfigParameterArgs struct {
	ParameterKey   pulumi.StringPtrInput `pulumi:"parameterKey"`
	ParameterValue pulumi.StringPtrInput `pulumi:"parameterValue"`
}

func (WorkgroupConfigParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkgroupConfigParameter)(nil)).Elem()
}

func (i WorkgroupConfigParameterArgs) ToWorkgroupConfigParameterOutput() WorkgroupConfigParameterOutput {
	return i.ToWorkgroupConfigParameterOutputWithContext(context.Background())
}

func (i WorkgroupConfigParameterArgs) ToWorkgroupConfigParameterOutputWithContext(ctx context.Context) WorkgroupConfigParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkgroupConfigParameterOutput)
}

// WorkgroupConfigParameterArrayInput is an input type that accepts WorkgroupConfigParameterArray and WorkgroupConfigParameterArrayOutput values.
// You can construct a concrete instance of `WorkgroupConfigParameterArrayInput` via:
//
//	WorkgroupConfigParameterArray{ WorkgroupConfigParameterArgs{...} }
type WorkgroupConfigParameterArrayInput interface {
	pulumi.Input

	ToWorkgroupConfigParameterArrayOutput() WorkgroupConfigParameterArrayOutput
	ToWorkgroupConfigParameterArrayOutputWithContext(context.Context) WorkgroupConfigParameterArrayOutput
}

type WorkgroupConfigParameterArray []WorkgroupConfigParameterInput

func (WorkgroupConfigParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkgroupConfigParameter)(nil)).Elem()
}

func (i WorkgroupConfigParameterArray) ToWorkgroupConfigParameterArrayOutput() WorkgroupConfigParameterArrayOutput {
	return i.ToWorkgroupConfigParameterArrayOutputWithContext(context.Background())
}

func (i WorkgroupConfigParameterArray) ToWorkgroupConfigParameterArrayOutputWithContext(ctx context.Context) WorkgroupConfigParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkgroupConfigParameterArrayOutput)
}

type WorkgroupConfigParameterOutput struct{ *pulumi.OutputState }

func (WorkgroupConfigParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkgroupConfigParameter)(nil)).Elem()
}

func (o WorkgroupConfigParameterOutput) ToWorkgroupConfigParameterOutput() WorkgroupConfigParameterOutput {
	return o
}

func (o WorkgroupConfigParameterOutput) ToWorkgroupConfigParameterOutputWithContext(ctx context.Context) WorkgroupConfigParameterOutput {
	return o
}

func (o WorkgroupConfigParameterOutput) ParameterKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupConfigParameter) *string { return v.ParameterKey }).(pulumi.StringPtrOutput)
}

func (o WorkgroupConfigParameterOutput) ParameterValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupConfigParameter) *string { return v.ParameterValue }).(pulumi.StringPtrOutput)
}

type WorkgroupConfigParameterArrayOutput struct{ *pulumi.OutputState }

func (WorkgroupConfigParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkgroupConfigParameter)(nil)).Elem()
}

func (o WorkgroupConfigParameterArrayOutput) ToWorkgroupConfigParameterArrayOutput() WorkgroupConfigParameterArrayOutput {
	return o
}

func (o WorkgroupConfigParameterArrayOutput) ToWorkgroupConfigParameterArrayOutputWithContext(ctx context.Context) WorkgroupConfigParameterArrayOutput {
	return o
}

func (o WorkgroupConfigParameterArrayOutput) Index(i pulumi.IntInput) WorkgroupConfigParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkgroupConfigParameter {
		return vs[0].([]WorkgroupConfigParameter)[vs[1].(int)]
	}).(WorkgroupConfigParameterOutput)
}

type WorkgroupEndpoint struct {
	Address      *string                `pulumi:"address"`
	Port         *int                   `pulumi:"port"`
	VpcEndpoints []WorkgroupVpcEndpoint `pulumi:"vpcEndpoints"`
}

type WorkgroupEndpointOutput struct{ *pulumi.OutputState }

func (WorkgroupEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkgroupEndpoint)(nil)).Elem()
}

func (o WorkgroupEndpointOutput) ToWorkgroupEndpointOutput() WorkgroupEndpointOutput {
	return o
}

func (o WorkgroupEndpointOutput) ToWorkgroupEndpointOutputWithContext(ctx context.Context) WorkgroupEndpointOutput {
	return o
}

func (o WorkgroupEndpointOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupEndpoint) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o WorkgroupEndpointOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkgroupEndpoint) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o WorkgroupEndpointOutput) VpcEndpoints() WorkgroupVpcEndpointArrayOutput {
	return o.ApplyT(func(v WorkgroupEndpoint) []WorkgroupVpcEndpoint { return v.VpcEndpoints }).(WorkgroupVpcEndpointArrayOutput)
}

type WorkgroupEndpointPtrOutput struct{ *pulumi.OutputState }

func (WorkgroupEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkgroupEndpoint)(nil)).Elem()
}

func (o WorkgroupEndpointPtrOutput) ToWorkgroupEndpointPtrOutput() WorkgroupEndpointPtrOutput {
	return o
}

func (o WorkgroupEndpointPtrOutput) ToWorkgroupEndpointPtrOutputWithContext(ctx context.Context) WorkgroupEndpointPtrOutput {
	return o
}

func (o WorkgroupEndpointPtrOutput) Elem() WorkgroupEndpointOutput {
	return o.ApplyT(func(v *WorkgroupEndpoint) WorkgroupEndpoint {
		if v != nil {
			return *v
		}
		var ret WorkgroupEndpoint
		return ret
	}).(WorkgroupEndpointOutput)
}

func (o WorkgroupEndpointPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkgroupEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o WorkgroupEndpointPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkgroupEndpoint) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

func (o WorkgroupEndpointPtrOutput) VpcEndpoints() WorkgroupVpcEndpointArrayOutput {
	return o.ApplyT(func(v *WorkgroupEndpoint) []WorkgroupVpcEndpoint {
		if v == nil {
			return nil
		}
		return v.VpcEndpoints
	}).(WorkgroupVpcEndpointArrayOutput)
}

type WorkgroupNetworkInterface struct {
	AvailabilityZone   *string `pulumi:"availabilityZone"`
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	PrivateIpAddress   *string `pulumi:"privateIpAddress"`
	SubnetId           *string `pulumi:"subnetId"`
}

type WorkgroupNetworkInterfaceOutput struct{ *pulumi.OutputState }

func (WorkgroupNetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkgroupNetworkInterface)(nil)).Elem()
}

func (o WorkgroupNetworkInterfaceOutput) ToWorkgroupNetworkInterfaceOutput() WorkgroupNetworkInterfaceOutput {
	return o
}

func (o WorkgroupNetworkInterfaceOutput) ToWorkgroupNetworkInterfaceOutputWithContext(ctx context.Context) WorkgroupNetworkInterfaceOutput {
	return o
}

func (o WorkgroupNetworkInterfaceOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupNetworkInterface) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o WorkgroupNetworkInterfaceOutput) NetworkInterfaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupNetworkInterface) *string { return v.NetworkInterfaceId }).(pulumi.StringPtrOutput)
}

func (o WorkgroupNetworkInterfaceOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupNetworkInterface) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o WorkgroupNetworkInterfaceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupNetworkInterface) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type WorkgroupNetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (WorkgroupNetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkgroupNetworkInterface)(nil)).Elem()
}

func (o WorkgroupNetworkInterfaceArrayOutput) ToWorkgroupNetworkInterfaceArrayOutput() WorkgroupNetworkInterfaceArrayOutput {
	return o
}

func (o WorkgroupNetworkInterfaceArrayOutput) ToWorkgroupNetworkInterfaceArrayOutputWithContext(ctx context.Context) WorkgroupNetworkInterfaceArrayOutput {
	return o
}

func (o WorkgroupNetworkInterfaceArrayOutput) Index(i pulumi.IntInput) WorkgroupNetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkgroupNetworkInterface {
		return vs[0].([]WorkgroupNetworkInterface)[vs[1].(int)]
	}).(WorkgroupNetworkInterfaceOutput)
}

type WorkgroupTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// WorkgroupTagInput is an input type that accepts WorkgroupTagArgs and WorkgroupTagOutput values.
// You can construct a concrete instance of `WorkgroupTagInput` via:
//
//	WorkgroupTagArgs{...}
type WorkgroupTagInput interface {
	pulumi.Input

	ToWorkgroupTagOutput() WorkgroupTagOutput
	ToWorkgroupTagOutputWithContext(context.Context) WorkgroupTagOutput
}

type WorkgroupTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (WorkgroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkgroupTag)(nil)).Elem()
}

func (i WorkgroupTagArgs) ToWorkgroupTagOutput() WorkgroupTagOutput {
	return i.ToWorkgroupTagOutputWithContext(context.Background())
}

func (i WorkgroupTagArgs) ToWorkgroupTagOutputWithContext(ctx context.Context) WorkgroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkgroupTagOutput)
}

// WorkgroupTagArrayInput is an input type that accepts WorkgroupTagArray and WorkgroupTagArrayOutput values.
// You can construct a concrete instance of `WorkgroupTagArrayInput` via:
//
//	WorkgroupTagArray{ WorkgroupTagArgs{...} }
type WorkgroupTagArrayInput interface {
	pulumi.Input

	ToWorkgroupTagArrayOutput() WorkgroupTagArrayOutput
	ToWorkgroupTagArrayOutputWithContext(context.Context) WorkgroupTagArrayOutput
}

type WorkgroupTagArray []WorkgroupTagInput

func (WorkgroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkgroupTag)(nil)).Elem()
}

func (i WorkgroupTagArray) ToWorkgroupTagArrayOutput() WorkgroupTagArrayOutput {
	return i.ToWorkgroupTagArrayOutputWithContext(context.Background())
}

func (i WorkgroupTagArray) ToWorkgroupTagArrayOutputWithContext(ctx context.Context) WorkgroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkgroupTagArrayOutput)
}

type WorkgroupTagOutput struct{ *pulumi.OutputState }

func (WorkgroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkgroupTag)(nil)).Elem()
}

func (o WorkgroupTagOutput) ToWorkgroupTagOutput() WorkgroupTagOutput {
	return o
}

func (o WorkgroupTagOutput) ToWorkgroupTagOutputWithContext(ctx context.Context) WorkgroupTagOutput {
	return o
}

func (o WorkgroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v WorkgroupTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o WorkgroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v WorkgroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type WorkgroupTagArrayOutput struct{ *pulumi.OutputState }

func (WorkgroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkgroupTag)(nil)).Elem()
}

func (o WorkgroupTagArrayOutput) ToWorkgroupTagArrayOutput() WorkgroupTagArrayOutput {
	return o
}

func (o WorkgroupTagArrayOutput) ToWorkgroupTagArrayOutputWithContext(ctx context.Context) WorkgroupTagArrayOutput {
	return o
}

func (o WorkgroupTagArrayOutput) Index(i pulumi.IntInput) WorkgroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkgroupTag {
		return vs[0].([]WorkgroupTag)[vs[1].(int)]
	}).(WorkgroupTagOutput)
}

type WorkgroupVpcEndpoint struct {
	NetworkInterfaces []WorkgroupNetworkInterface `pulumi:"networkInterfaces"`
	VpcEndpointId     *string                     `pulumi:"vpcEndpointId"`
	VpcId             *string                     `pulumi:"vpcId"`
}

type WorkgroupVpcEndpointOutput struct{ *pulumi.OutputState }

func (WorkgroupVpcEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkgroupVpcEndpoint)(nil)).Elem()
}

func (o WorkgroupVpcEndpointOutput) ToWorkgroupVpcEndpointOutput() WorkgroupVpcEndpointOutput {
	return o
}

func (o WorkgroupVpcEndpointOutput) ToWorkgroupVpcEndpointOutputWithContext(ctx context.Context) WorkgroupVpcEndpointOutput {
	return o
}

func (o WorkgroupVpcEndpointOutput) NetworkInterfaces() WorkgroupNetworkInterfaceArrayOutput {
	return o.ApplyT(func(v WorkgroupVpcEndpoint) []WorkgroupNetworkInterface { return v.NetworkInterfaces }).(WorkgroupNetworkInterfaceArrayOutput)
}

func (o WorkgroupVpcEndpointOutput) VpcEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupVpcEndpoint) *string { return v.VpcEndpointId }).(pulumi.StringPtrOutput)
}

func (o WorkgroupVpcEndpointOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkgroupVpcEndpoint) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

type WorkgroupVpcEndpointArrayOutput struct{ *pulumi.OutputState }

func (WorkgroupVpcEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkgroupVpcEndpoint)(nil)).Elem()
}

func (o WorkgroupVpcEndpointArrayOutput) ToWorkgroupVpcEndpointArrayOutput() WorkgroupVpcEndpointArrayOutput {
	return o
}

func (o WorkgroupVpcEndpointArrayOutput) ToWorkgroupVpcEndpointArrayOutputWithContext(ctx context.Context) WorkgroupVpcEndpointArrayOutput {
	return o
}

func (o WorkgroupVpcEndpointArrayOutput) Index(i pulumi.IntInput) WorkgroupVpcEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkgroupVpcEndpoint {
		return vs[0].([]WorkgroupVpcEndpoint)[vs[1].(int)]
	}).(WorkgroupVpcEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceTagInput)(nil)).Elem(), NamespaceTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceTagArrayInput)(nil)).Elem(), NamespaceTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkgroupConfigParameterInput)(nil)).Elem(), WorkgroupConfigParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkgroupConfigParameterArrayInput)(nil)).Elem(), WorkgroupConfigParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkgroupTagInput)(nil)).Elem(), WorkgroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkgroupTagArrayInput)(nil)).Elem(), WorkgroupTagArray{})
	pulumi.RegisterOutputType(NamespaceTypeOutput{})
	pulumi.RegisterOutputType(NamespaceTypePtrOutput{})
	pulumi.RegisterOutputType(NamespaceTagOutput{})
	pulumi.RegisterOutputType(NamespaceTagArrayOutput{})
	pulumi.RegisterOutputType(WorkgroupTypeOutput{})
	pulumi.RegisterOutputType(WorkgroupTypePtrOutput{})
	pulumi.RegisterOutputType(WorkgroupConfigParameterOutput{})
	pulumi.RegisterOutputType(WorkgroupConfigParameterArrayOutput{})
	pulumi.RegisterOutputType(WorkgroupEndpointOutput{})
	pulumi.RegisterOutputType(WorkgroupEndpointPtrOutput{})
	pulumi.RegisterOutputType(WorkgroupNetworkInterfaceOutput{})
	pulumi.RegisterOutputType(WorkgroupNetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(WorkgroupTagOutput{})
	pulumi.RegisterOutputType(WorkgroupTagArrayOutput{})
	pulumi.RegisterOutputType(WorkgroupVpcEndpointOutput{})
	pulumi.RegisterOutputType(WorkgroupVpcEndpointArrayOutput{})
}