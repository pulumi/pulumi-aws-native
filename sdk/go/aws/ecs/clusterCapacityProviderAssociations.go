// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Associate a set of ECS Capacity Providers with a specified ECS Cluster
//
// ## Example Usage
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			clusterName := cfg.Require("clusterName")
//			_, err := ecs.NewClusterCapacityProviderAssociations(ctx, "clusterCPAssociation", &ecs.ClusterCapacityProviderAssociationsArgs{
//				Cluster: pulumi.String(clusterName),
//				CapacityProviders: pulumi.StringArray{
//					pulumi.String(ecs.ClusterCapacityProviderAssociationsCapacityProviderFargate),
//					pulumi.String(ecs.ClusterCapacityProviderAssociationsCapacityProviderFargateSpot),
//				},
//				DefaultCapacityProviderStrategy: ecs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArray{
//					&ecs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs{
//						Base:             pulumi.Int(2),
//						Weight:           pulumi.Int(1),
//						CapacityProvider: pulumi.String(ecs.ClusterCapacityProviderAssociationsCapacityProviderFargate),
//					},
//					&ecs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs{
//						Base:             pulumi.Int(0),
//						Weight:           pulumi.Int(1),
//						CapacityProvider: pulumi.String(ecs.ClusterCapacityProviderAssociationsCapacityProviderFargateSpot),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			clusterName := cfg.Require("clusterName")
//			_, err := ecs.NewClusterCapacityProviderAssociations(ctx, "clusterCPAssociation", &ecs.ClusterCapacityProviderAssociationsArgs{
//				Cluster: pulumi.String(clusterName),
//				CapacityProviders: pulumi.StringArray{
//					pulumi.String(ecs.ClusterCapacityProviderAssociationsCapacityProviderFargate),
//					pulumi.String(ecs.ClusterCapacityProviderAssociationsCapacityProviderFargateSpot),
//				},
//				DefaultCapacityProviderStrategy: ecs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArray{
//					&ecs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs{
//						Base:             pulumi.Int(2),
//						Weight:           pulumi.Int(1),
//						CapacityProvider: pulumi.String(ecs.ClusterCapacityProviderAssociationsCapacityProviderFargate),
//					},
//					&ecs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs{
//						Base:             pulumi.Int(0),
//						Weight:           pulumi.Int(1),
//						CapacityProvider: pulumi.String(ecs.ClusterCapacityProviderAssociationsCapacityProviderFargateSpot),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ClusterCapacityProviderAssociations struct {
	pulumi.CustomResourceState

	// The capacity providers to associate with the cluster.
	CapacityProviders pulumi.StringArrayOutput `pulumi:"capacityProviders"`
	// The cluster the capacity provider association is the target of.
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// The default capacity provider strategy to associate with the cluster.
	DefaultCapacityProviderStrategy ClusterCapacityProviderAssociationsCapacityProviderStrategyArrayOutput `pulumi:"defaultCapacityProviderStrategy"`
}

// NewClusterCapacityProviderAssociations registers a new resource with the given unique name, arguments, and options.
func NewClusterCapacityProviderAssociations(ctx *pulumi.Context,
	name string, args *ClusterCapacityProviderAssociationsArgs, opts ...pulumi.ResourceOption) (*ClusterCapacityProviderAssociations, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CapacityProviders == nil {
		return nil, errors.New("invalid value for required argument 'CapacityProviders'")
	}
	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
	}
	if args.DefaultCapacityProviderStrategy == nil {
		return nil, errors.New("invalid value for required argument 'DefaultCapacityProviderStrategy'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"cluster",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterCapacityProviderAssociations
	err := ctx.RegisterResource("aws-native:ecs:ClusterCapacityProviderAssociations", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterCapacityProviderAssociations gets an existing ClusterCapacityProviderAssociations resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterCapacityProviderAssociations(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterCapacityProviderAssociationsState, opts ...pulumi.ResourceOption) (*ClusterCapacityProviderAssociations, error) {
	var resource ClusterCapacityProviderAssociations
	err := ctx.ReadResource("aws-native:ecs:ClusterCapacityProviderAssociations", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterCapacityProviderAssociations resources.
type clusterCapacityProviderAssociationsState struct {
}

type ClusterCapacityProviderAssociationsState struct {
}

func (ClusterCapacityProviderAssociationsState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterCapacityProviderAssociationsState)(nil)).Elem()
}

type clusterCapacityProviderAssociationsArgs struct {
	// The capacity providers to associate with the cluster.
	CapacityProviders []string `pulumi:"capacityProviders"`
	// The cluster the capacity provider association is the target of.
	Cluster string `pulumi:"cluster"`
	// The default capacity provider strategy to associate with the cluster.
	DefaultCapacityProviderStrategy []ClusterCapacityProviderAssociationsCapacityProviderStrategy `pulumi:"defaultCapacityProviderStrategy"`
}

// The set of arguments for constructing a ClusterCapacityProviderAssociations resource.
type ClusterCapacityProviderAssociationsArgs struct {
	// The capacity providers to associate with the cluster.
	CapacityProviders pulumi.StringArrayInput
	// The cluster the capacity provider association is the target of.
	Cluster pulumi.StringInput
	// The default capacity provider strategy to associate with the cluster.
	DefaultCapacityProviderStrategy ClusterCapacityProviderAssociationsCapacityProviderStrategyArrayInput
}

func (ClusterCapacityProviderAssociationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterCapacityProviderAssociationsArgs)(nil)).Elem()
}

type ClusterCapacityProviderAssociationsInput interface {
	pulumi.Input

	ToClusterCapacityProviderAssociationsOutput() ClusterCapacityProviderAssociationsOutput
	ToClusterCapacityProviderAssociationsOutputWithContext(ctx context.Context) ClusterCapacityProviderAssociationsOutput
}

func (*ClusterCapacityProviderAssociations) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterCapacityProviderAssociations)(nil)).Elem()
}

func (i *ClusterCapacityProviderAssociations) ToClusterCapacityProviderAssociationsOutput() ClusterCapacityProviderAssociationsOutput {
	return i.ToClusterCapacityProviderAssociationsOutputWithContext(context.Background())
}

func (i *ClusterCapacityProviderAssociations) ToClusterCapacityProviderAssociationsOutputWithContext(ctx context.Context) ClusterCapacityProviderAssociationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCapacityProviderAssociationsOutput)
}

type ClusterCapacityProviderAssociationsOutput struct{ *pulumi.OutputState }

func (ClusterCapacityProviderAssociationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterCapacityProviderAssociations)(nil)).Elem()
}

func (o ClusterCapacityProviderAssociationsOutput) ToClusterCapacityProviderAssociationsOutput() ClusterCapacityProviderAssociationsOutput {
	return o
}

func (o ClusterCapacityProviderAssociationsOutput) ToClusterCapacityProviderAssociationsOutputWithContext(ctx context.Context) ClusterCapacityProviderAssociationsOutput {
	return o
}

// The capacity providers to associate with the cluster.
func (o ClusterCapacityProviderAssociationsOutput) CapacityProviders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterCapacityProviderAssociations) pulumi.StringArrayOutput { return v.CapacityProviders }).(pulumi.StringArrayOutput)
}

// The cluster the capacity provider association is the target of.
func (o ClusterCapacityProviderAssociationsOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterCapacityProviderAssociations) pulumi.StringOutput { return v.Cluster }).(pulumi.StringOutput)
}

// The default capacity provider strategy to associate with the cluster.
func (o ClusterCapacityProviderAssociationsOutput) DefaultCapacityProviderStrategy() ClusterCapacityProviderAssociationsCapacityProviderStrategyArrayOutput {
	return o.ApplyT(func(v *ClusterCapacityProviderAssociations) ClusterCapacityProviderAssociationsCapacityProviderStrategyArrayOutput {
		return v.DefaultCapacityProviderStrategy
	}).(ClusterCapacityProviderAssociationsCapacityProviderStrategyArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterCapacityProviderAssociationsInput)(nil)).Elem(), &ClusterCapacityProviderAssociations{})
	pulumi.RegisterOutputType(ClusterCapacityProviderAssociationsOutput{})
}
