// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AutoScaling::AutoScalingGroup
//
// Deprecated: AutoScalingGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type AutoScalingGroup struct {
	pulumi.CustomResourceState

	AutoScalingGroupName             pulumi.StringPtrOutput                                `pulumi:"autoScalingGroupName"`
	AvailabilityZones                pulumi.StringArrayOutput                              `pulumi:"availabilityZones"`
	CapacityRebalance                pulumi.BoolPtrOutput                                  `pulumi:"capacityRebalance"`
	Context                          pulumi.StringPtrOutput                                `pulumi:"context"`
	Cooldown                         pulumi.StringPtrOutput                                `pulumi:"cooldown"`
	DefaultInstanceWarmup            pulumi.IntPtrOutput                                   `pulumi:"defaultInstanceWarmup"`
	DesiredCapacity                  pulumi.StringPtrOutput                                `pulumi:"desiredCapacity"`
	DesiredCapacityType              pulumi.StringPtrOutput                                `pulumi:"desiredCapacityType"`
	HealthCheckGracePeriod           pulumi.IntPtrOutput                                   `pulumi:"healthCheckGracePeriod"`
	HealthCheckType                  pulumi.StringPtrOutput                                `pulumi:"healthCheckType"`
	InstanceId                       pulumi.StringPtrOutput                                `pulumi:"instanceId"`
	LaunchConfigurationName          pulumi.StringPtrOutput                                `pulumi:"launchConfigurationName"`
	LaunchTemplate                   AutoScalingGroupLaunchTemplateSpecificationPtrOutput  `pulumi:"launchTemplate"`
	LaunchTemplateSpecification      pulumi.StringOutput                                   `pulumi:"launchTemplateSpecification"`
	LifecycleHookSpecificationList   AutoScalingGroupLifecycleHookSpecificationArrayOutput `pulumi:"lifecycleHookSpecificationList"`
	LoadBalancerNames                pulumi.StringArrayOutput                              `pulumi:"loadBalancerNames"`
	MaxInstanceLifetime              pulumi.IntPtrOutput                                   `pulumi:"maxInstanceLifetime"`
	MaxSize                          pulumi.StringOutput                                   `pulumi:"maxSize"`
	MetricsCollection                AutoScalingGroupMetricsCollectionArrayOutput          `pulumi:"metricsCollection"`
	MinSize                          pulumi.StringOutput                                   `pulumi:"minSize"`
	MixedInstancesPolicy             AutoScalingGroupMixedInstancesPolicyPtrOutput         `pulumi:"mixedInstancesPolicy"`
	NewInstancesProtectedFromScaleIn pulumi.BoolPtrOutput                                  `pulumi:"newInstancesProtectedFromScaleIn"`
	NotificationConfigurations       AutoScalingGroupNotificationConfigurationArrayOutput  `pulumi:"notificationConfigurations"`
	PlacementGroup                   pulumi.StringPtrOutput                                `pulumi:"placementGroup"`
	ServiceLinkedRoleARN             pulumi.StringPtrOutput                                `pulumi:"serviceLinkedRoleARN"`
	Tags                             AutoScalingGroupTagPropertyArrayOutput                `pulumi:"tags"`
	TargetGroupARNs                  pulumi.StringArrayOutput                              `pulumi:"targetGroupARNs"`
	TerminationPolicies              pulumi.StringArrayOutput                              `pulumi:"terminationPolicies"`
	VPCZoneIdentifier                pulumi.StringArrayOutput                              `pulumi:"vPCZoneIdentifier"`
}

// NewAutoScalingGroup registers a new resource with the given unique name, arguments, and options.
func NewAutoScalingGroup(ctx *pulumi.Context,
	name string, args *AutoScalingGroupArgs, opts ...pulumi.ResourceOption) (*AutoScalingGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxSize == nil {
		return nil, errors.New("invalid value for required argument 'MaxSize'")
	}
	if args.MinSize == nil {
		return nil, errors.New("invalid value for required argument 'MinSize'")
	}
	var resource AutoScalingGroup
	err := ctx.RegisterResource("aws-native:autoscaling:AutoScalingGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoScalingGroup gets an existing AutoScalingGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoScalingGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoScalingGroupState, opts ...pulumi.ResourceOption) (*AutoScalingGroup, error) {
	var resource AutoScalingGroup
	err := ctx.ReadResource("aws-native:autoscaling:AutoScalingGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoScalingGroup resources.
type autoScalingGroupState struct {
}

type AutoScalingGroupState struct {
}

func (AutoScalingGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoScalingGroupState)(nil)).Elem()
}

type autoScalingGroupArgs struct {
	AutoScalingGroupName             *string                                      `pulumi:"autoScalingGroupName"`
	AvailabilityZones                []string                                     `pulumi:"availabilityZones"`
	CapacityRebalance                *bool                                        `pulumi:"capacityRebalance"`
	Context                          *string                                      `pulumi:"context"`
	Cooldown                         *string                                      `pulumi:"cooldown"`
	DefaultInstanceWarmup            *int                                         `pulumi:"defaultInstanceWarmup"`
	DesiredCapacity                  *string                                      `pulumi:"desiredCapacity"`
	DesiredCapacityType              *string                                      `pulumi:"desiredCapacityType"`
	HealthCheckGracePeriod           *int                                         `pulumi:"healthCheckGracePeriod"`
	HealthCheckType                  *string                                      `pulumi:"healthCheckType"`
	InstanceId                       *string                                      `pulumi:"instanceId"`
	LaunchConfigurationName          *string                                      `pulumi:"launchConfigurationName"`
	LaunchTemplate                   *AutoScalingGroupLaunchTemplateSpecification `pulumi:"launchTemplate"`
	LifecycleHookSpecificationList   []AutoScalingGroupLifecycleHookSpecification `pulumi:"lifecycleHookSpecificationList"`
	LoadBalancerNames                []string                                     `pulumi:"loadBalancerNames"`
	MaxInstanceLifetime              *int                                         `pulumi:"maxInstanceLifetime"`
	MaxSize                          string                                       `pulumi:"maxSize"`
	MetricsCollection                []AutoScalingGroupMetricsCollection          `pulumi:"metricsCollection"`
	MinSize                          string                                       `pulumi:"minSize"`
	MixedInstancesPolicy             *AutoScalingGroupMixedInstancesPolicy        `pulumi:"mixedInstancesPolicy"`
	NewInstancesProtectedFromScaleIn *bool                                        `pulumi:"newInstancesProtectedFromScaleIn"`
	NotificationConfigurations       []AutoScalingGroupNotificationConfiguration  `pulumi:"notificationConfigurations"`
	PlacementGroup                   *string                                      `pulumi:"placementGroup"`
	ServiceLinkedRoleARN             *string                                      `pulumi:"serviceLinkedRoleARN"`
	Tags                             []AutoScalingGroupTagProperty                `pulumi:"tags"`
	TargetGroupARNs                  []string                                     `pulumi:"targetGroupARNs"`
	TerminationPolicies              []string                                     `pulumi:"terminationPolicies"`
	VPCZoneIdentifier                []string                                     `pulumi:"vPCZoneIdentifier"`
}

// The set of arguments for constructing a AutoScalingGroup resource.
type AutoScalingGroupArgs struct {
	AutoScalingGroupName             pulumi.StringPtrInput
	AvailabilityZones                pulumi.StringArrayInput
	CapacityRebalance                pulumi.BoolPtrInput
	Context                          pulumi.StringPtrInput
	Cooldown                         pulumi.StringPtrInput
	DefaultInstanceWarmup            pulumi.IntPtrInput
	DesiredCapacity                  pulumi.StringPtrInput
	DesiredCapacityType              pulumi.StringPtrInput
	HealthCheckGracePeriod           pulumi.IntPtrInput
	HealthCheckType                  pulumi.StringPtrInput
	InstanceId                       pulumi.StringPtrInput
	LaunchConfigurationName          pulumi.StringPtrInput
	LaunchTemplate                   AutoScalingGroupLaunchTemplateSpecificationPtrInput
	LifecycleHookSpecificationList   AutoScalingGroupLifecycleHookSpecificationArrayInput
	LoadBalancerNames                pulumi.StringArrayInput
	MaxInstanceLifetime              pulumi.IntPtrInput
	MaxSize                          pulumi.StringInput
	MetricsCollection                AutoScalingGroupMetricsCollectionArrayInput
	MinSize                          pulumi.StringInput
	MixedInstancesPolicy             AutoScalingGroupMixedInstancesPolicyPtrInput
	NewInstancesProtectedFromScaleIn pulumi.BoolPtrInput
	NotificationConfigurations       AutoScalingGroupNotificationConfigurationArrayInput
	PlacementGroup                   pulumi.StringPtrInput
	ServiceLinkedRoleARN             pulumi.StringPtrInput
	Tags                             AutoScalingGroupTagPropertyArrayInput
	TargetGroupARNs                  pulumi.StringArrayInput
	TerminationPolicies              pulumi.StringArrayInput
	VPCZoneIdentifier                pulumi.StringArrayInput
}

func (AutoScalingGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoScalingGroupArgs)(nil)).Elem()
}

type AutoScalingGroupInput interface {
	pulumi.Input

	ToAutoScalingGroupOutput() AutoScalingGroupOutput
	ToAutoScalingGroupOutputWithContext(ctx context.Context) AutoScalingGroupOutput
}

func (*AutoScalingGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalingGroup)(nil)).Elem()
}

func (i *AutoScalingGroup) ToAutoScalingGroupOutput() AutoScalingGroupOutput {
	return i.ToAutoScalingGroupOutputWithContext(context.Background())
}

func (i *AutoScalingGroup) ToAutoScalingGroupOutputWithContext(ctx context.Context) AutoScalingGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalingGroupOutput)
}

type AutoScalingGroupOutput struct{ *pulumi.OutputState }

func (AutoScalingGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalingGroup)(nil)).Elem()
}

func (o AutoScalingGroupOutput) ToAutoScalingGroupOutput() AutoScalingGroupOutput {
	return o
}

func (o AutoScalingGroupOutput) ToAutoScalingGroupOutputWithContext(ctx context.Context) AutoScalingGroupOutput {
	return o
}

func (o AutoScalingGroupOutput) AutoScalingGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringPtrOutput { return v.AutoScalingGroupName }).(pulumi.StringPtrOutput)
}

func (o AutoScalingGroupOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o AutoScalingGroupOutput) CapacityRebalance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.BoolPtrOutput { return v.CapacityRebalance }).(pulumi.BoolPtrOutput)
}

func (o AutoScalingGroupOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringPtrOutput { return v.Context }).(pulumi.StringPtrOutput)
}

func (o AutoScalingGroupOutput) Cooldown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringPtrOutput { return v.Cooldown }).(pulumi.StringPtrOutput)
}

func (o AutoScalingGroupOutput) DefaultInstanceWarmup() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.IntPtrOutput { return v.DefaultInstanceWarmup }).(pulumi.IntPtrOutput)
}

func (o AutoScalingGroupOutput) DesiredCapacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringPtrOutput { return v.DesiredCapacity }).(pulumi.StringPtrOutput)
}

func (o AutoScalingGroupOutput) DesiredCapacityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringPtrOutput { return v.DesiredCapacityType }).(pulumi.StringPtrOutput)
}

func (o AutoScalingGroupOutput) HealthCheckGracePeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.IntPtrOutput { return v.HealthCheckGracePeriod }).(pulumi.IntPtrOutput)
}

func (o AutoScalingGroupOutput) HealthCheckType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringPtrOutput { return v.HealthCheckType }).(pulumi.StringPtrOutput)
}

func (o AutoScalingGroupOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringPtrOutput { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o AutoScalingGroupOutput) LaunchConfigurationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringPtrOutput { return v.LaunchConfigurationName }).(pulumi.StringPtrOutput)
}

func (o AutoScalingGroupOutput) LaunchTemplate() AutoScalingGroupLaunchTemplateSpecificationPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) AutoScalingGroupLaunchTemplateSpecificationPtrOutput {
		return v.LaunchTemplate
	}).(AutoScalingGroupLaunchTemplateSpecificationPtrOutput)
}

func (o AutoScalingGroupOutput) LaunchTemplateSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringOutput { return v.LaunchTemplateSpecification }).(pulumi.StringOutput)
}

func (o AutoScalingGroupOutput) LifecycleHookSpecificationList() AutoScalingGroupLifecycleHookSpecificationArrayOutput {
	return o.ApplyT(func(v *AutoScalingGroup) AutoScalingGroupLifecycleHookSpecificationArrayOutput {
		return v.LifecycleHookSpecificationList
	}).(AutoScalingGroupLifecycleHookSpecificationArrayOutput)
}

func (o AutoScalingGroupOutput) LoadBalancerNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringArrayOutput { return v.LoadBalancerNames }).(pulumi.StringArrayOutput)
}

func (o AutoScalingGroupOutput) MaxInstanceLifetime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.IntPtrOutput { return v.MaxInstanceLifetime }).(pulumi.IntPtrOutput)
}

func (o AutoScalingGroupOutput) MaxSize() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringOutput { return v.MaxSize }).(pulumi.StringOutput)
}

func (o AutoScalingGroupOutput) MetricsCollection() AutoScalingGroupMetricsCollectionArrayOutput {
	return o.ApplyT(func(v *AutoScalingGroup) AutoScalingGroupMetricsCollectionArrayOutput { return v.MetricsCollection }).(AutoScalingGroupMetricsCollectionArrayOutput)
}

func (o AutoScalingGroupOutput) MinSize() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringOutput { return v.MinSize }).(pulumi.StringOutput)
}

func (o AutoScalingGroupOutput) MixedInstancesPolicy() AutoScalingGroupMixedInstancesPolicyPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) AutoScalingGroupMixedInstancesPolicyPtrOutput { return v.MixedInstancesPolicy }).(AutoScalingGroupMixedInstancesPolicyPtrOutput)
}

func (o AutoScalingGroupOutput) NewInstancesProtectedFromScaleIn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.BoolPtrOutput { return v.NewInstancesProtectedFromScaleIn }).(pulumi.BoolPtrOutput)
}

func (o AutoScalingGroupOutput) NotificationConfigurations() AutoScalingGroupNotificationConfigurationArrayOutput {
	return o.ApplyT(func(v *AutoScalingGroup) AutoScalingGroupNotificationConfigurationArrayOutput {
		return v.NotificationConfigurations
	}).(AutoScalingGroupNotificationConfigurationArrayOutput)
}

func (o AutoScalingGroupOutput) PlacementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringPtrOutput { return v.PlacementGroup }).(pulumi.StringPtrOutput)
}

func (o AutoScalingGroupOutput) ServiceLinkedRoleARN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringPtrOutput { return v.ServiceLinkedRoleARN }).(pulumi.StringPtrOutput)
}

func (o AutoScalingGroupOutput) Tags() AutoScalingGroupTagPropertyArrayOutput {
	return o.ApplyT(func(v *AutoScalingGroup) AutoScalingGroupTagPropertyArrayOutput { return v.Tags }).(AutoScalingGroupTagPropertyArrayOutput)
}

func (o AutoScalingGroupOutput) TargetGroupARNs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringArrayOutput { return v.TargetGroupARNs }).(pulumi.StringArrayOutput)
}

func (o AutoScalingGroupOutput) TerminationPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringArrayOutput { return v.TerminationPolicies }).(pulumi.StringArrayOutput)
}

func (o AutoScalingGroupOutput) VPCZoneIdentifier() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoScalingGroup) pulumi.StringArrayOutput { return v.VPCZoneIdentifier }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoScalingGroupInput)(nil)).Elem(), &AutoScalingGroup{})
	pulumi.RegisterOutputType(AutoScalingGroupOutput{})
}