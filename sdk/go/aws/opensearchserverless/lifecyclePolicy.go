// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearchserverless

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Amazon OpenSearchServerless lifecycle policy resource
type LifecyclePolicy struct {
	pulumi.CustomResourceState

	// The description of the policy
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the policy
	Name pulumi.StringOutput `pulumi:"name"`
	// The JSON policy document that is the content for the policy
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The type of lifecycle policy.
	Type LifecyclePolicyTypeOutput `pulumi:"type"`
}

// NewLifecyclePolicy registers a new resource with the given unique name, arguments, and options.
func NewLifecyclePolicy(ctx *pulumi.Context,
	name string, args *LifecyclePolicyArgs, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"type",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LifecyclePolicy
	err := ctx.RegisterResource("aws-native:opensearchserverless:LifecyclePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLifecyclePolicy gets an existing LifecyclePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLifecyclePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LifecyclePolicyState, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	var resource LifecyclePolicy
	err := ctx.ReadResource("aws-native:opensearchserverless:LifecyclePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LifecyclePolicy resources.
type lifecyclePolicyState struct {
}

type LifecyclePolicyState struct {
}

func (LifecyclePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyState)(nil)).Elem()
}

type lifecyclePolicyArgs struct {
	// The description of the policy
	Description *string `pulumi:"description"`
	// The name of the policy
	Name *string `pulumi:"name"`
	// The JSON policy document that is the content for the policy
	Policy string `pulumi:"policy"`
	// The type of lifecycle policy.
	Type LifecyclePolicyType `pulumi:"type"`
}

// The set of arguments for constructing a LifecyclePolicy resource.
type LifecyclePolicyArgs struct {
	// The description of the policy
	Description pulumi.StringPtrInput
	// The name of the policy
	Name pulumi.StringPtrInput
	// The JSON policy document that is the content for the policy
	Policy pulumi.StringInput
	// The type of lifecycle policy.
	Type LifecyclePolicyTypeInput
}

func (LifecyclePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyArgs)(nil)).Elem()
}

type LifecyclePolicyInput interface {
	pulumi.Input

	ToLifecyclePolicyOutput() LifecyclePolicyOutput
	ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput
}

func (*LifecyclePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**LifecyclePolicy)(nil)).Elem()
}

func (i *LifecyclePolicy) ToLifecyclePolicyOutput() LifecyclePolicyOutput {
	return i.ToLifecyclePolicyOutputWithContext(context.Background())
}

func (i *LifecyclePolicy) ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyOutput)
}

type LifecyclePolicyOutput struct{ *pulumi.OutputState }

func (LifecyclePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LifecyclePolicy)(nil)).Elem()
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyOutput() LifecyclePolicyOutput {
	return o
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput {
	return o
}

// The description of the policy
func (o LifecyclePolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the policy
func (o LifecyclePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The JSON policy document that is the content for the policy
func (o LifecyclePolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// The type of lifecycle policy.
func (o LifecyclePolicyOutput) Type() LifecyclePolicyTypeOutput {
	return o.ApplyT(func(v *LifecyclePolicy) LifecyclePolicyTypeOutput { return v.Type }).(LifecyclePolicyTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LifecyclePolicyInput)(nil)).Elem(), &LifecyclePolicy{})
	pulumi.RegisterOutputType(LifecyclePolicyOutput{})
}
