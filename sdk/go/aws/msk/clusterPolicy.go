// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::MSK::ClusterPolicy
type ClusterPolicy struct {
	pulumi.CustomResourceState

	// The arn of the cluster for the resource policy.
	ClusterArn pulumi.StringOutput `pulumi:"clusterArn"`
	// The current version of the policy attached to the specified cluster
	CurrentVersion pulumi.StringOutput `pulumi:"currentVersion"`
	// A policy document containing permissions to add to the specified cluster.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::MSK::ClusterPolicy` for more information about the expected schema for this property.
	Policy pulumi.AnyOutput `pulumi:"policy"`
}

// NewClusterPolicy registers a new resource with the given unique name, arguments, and options.
func NewClusterPolicy(ctx *pulumi.Context,
	name string, args *ClusterPolicyArgs, opts ...pulumi.ResourceOption) (*ClusterPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterArn == nil {
		return nil, errors.New("invalid value for required argument 'ClusterArn'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"clusterArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterPolicy
	err := ctx.RegisterResource("aws-native:msk:ClusterPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterPolicy gets an existing ClusterPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterPolicyState, opts ...pulumi.ResourceOption) (*ClusterPolicy, error) {
	var resource ClusterPolicy
	err := ctx.ReadResource("aws-native:msk:ClusterPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterPolicy resources.
type clusterPolicyState struct {
}

type ClusterPolicyState struct {
}

func (ClusterPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPolicyState)(nil)).Elem()
}

type clusterPolicyArgs struct {
	// The arn of the cluster for the resource policy.
	ClusterArn string `pulumi:"clusterArn"`
	// A policy document containing permissions to add to the specified cluster.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::MSK::ClusterPolicy` for more information about the expected schema for this property.
	Policy interface{} `pulumi:"policy"`
}

// The set of arguments for constructing a ClusterPolicy resource.
type ClusterPolicyArgs struct {
	// The arn of the cluster for the resource policy.
	ClusterArn pulumi.StringInput
	// A policy document containing permissions to add to the specified cluster.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::MSK::ClusterPolicy` for more information about the expected schema for this property.
	Policy pulumi.Input
}

func (ClusterPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPolicyArgs)(nil)).Elem()
}

type ClusterPolicyInput interface {
	pulumi.Input

	ToClusterPolicyOutput() ClusterPolicyOutput
	ToClusterPolicyOutputWithContext(ctx context.Context) ClusterPolicyOutput
}

func (*ClusterPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPolicy)(nil)).Elem()
}

func (i *ClusterPolicy) ToClusterPolicyOutput() ClusterPolicyOutput {
	return i.ToClusterPolicyOutputWithContext(context.Background())
}

func (i *ClusterPolicy) ToClusterPolicyOutputWithContext(ctx context.Context) ClusterPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPolicyOutput)
}

type ClusterPolicyOutput struct{ *pulumi.OutputState }

func (ClusterPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPolicy)(nil)).Elem()
}

func (o ClusterPolicyOutput) ToClusterPolicyOutput() ClusterPolicyOutput {
	return o
}

func (o ClusterPolicyOutput) ToClusterPolicyOutputWithContext(ctx context.Context) ClusterPolicyOutput {
	return o
}

// The arn of the cluster for the resource policy.
func (o ClusterPolicyOutput) ClusterArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPolicy) pulumi.StringOutput { return v.ClusterArn }).(pulumi.StringOutput)
}

// The current version of the policy attached to the specified cluster
func (o ClusterPolicyOutput) CurrentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPolicy) pulumi.StringOutput { return v.CurrentVersion }).(pulumi.StringOutput)
}

// A policy document containing permissions to add to the specified cluster.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::MSK::ClusterPolicy` for more information about the expected schema for this property.
func (o ClusterPolicyOutput) Policy() pulumi.AnyOutput {
	return o.ApplyT(func(v *ClusterPolicy) pulumi.AnyOutput { return v.Policy }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPolicyInput)(nil)).Elem(), &ClusterPolicy{})
	pulumi.RegisterOutputType(ClusterPolicyOutput{})
}
