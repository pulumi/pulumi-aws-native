// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sso

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for SSO PermissionSet
type PermissionSet struct {
	pulumi.CustomResourceState

	// Specifies the names and paths of the customer managed policies that you have attached to your permission set.
	CustomerManagedPolicyReferences PermissionSetCustomerManagedPolicyReferenceArrayOutput `pulumi:"customerManagedPolicyReferences"`
	// The permission set description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The inline policy to put in permission set.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SSO::PermissionSet` for more information about the expected schema for this property.
	InlinePolicy pulumi.AnyOutput `pulumi:"inlinePolicy"`
	// The sso instance arn that the permission set is owned.
	InstanceArn pulumi.StringOutput `pulumi:"instanceArn"`
	// A structure that stores a list of managed policy ARNs that describe the associated AWS managed policy.
	ManagedPolicies pulumi.StringArrayOutput `pulumi:"managedPolicies"`
	// The name you want to assign to this permission set.
	Name pulumi.StringOutput `pulumi:"name"`
	// The permission set that the policy will be attached to
	PermissionSetArn pulumi.StringOutput `pulumi:"permissionSetArn"`
	// Specifies the configuration of the AWS managed or customer managed policy that you want to set as a permissions boundary. Specify either `CustomerManagedPolicyReference` to use the name and path of a customer managed policy, or `ManagedPolicyArn` to use the ARN of an AWS managed policy. A permissions boundary represents the maximum permissions that any policy can grant your role. For more information, see [Permissions boundaries for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide* .
	//
	// > Policies used as permissions boundaries don't provide permissions. You must also attach an IAM policy to the role. To learn how the effective permissions for a role are evaluated, see [IAM JSON policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) in the *IAM User Guide* .
	PermissionsBoundary PermissionSetPermissionsBoundaryPtrOutput `pulumi:"permissionsBoundary"`
	// The relay state URL that redirect links to any service in the AWS Management Console.
	RelayStateType pulumi.StringPtrOutput `pulumi:"relayStateType"`
	// The length of time that a user can be signed in to an AWS account.
	SessionDuration pulumi.StringPtrOutput `pulumi:"sessionDuration"`
	// The tags to attach to the new `PermissionSet` .
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewPermissionSet registers a new resource with the given unique name, arguments, and options.
func NewPermissionSet(ctx *pulumi.Context,
	name string, args *PermissionSetArgs, opts ...pulumi.ResourceOption) (*PermissionSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'InstanceArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"instanceArn",
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PermissionSet
	err := ctx.RegisterResource("aws-native:sso:PermissionSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermissionSet gets an existing PermissionSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermissionSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionSetState, opts ...pulumi.ResourceOption) (*PermissionSet, error) {
	var resource PermissionSet
	err := ctx.ReadResource("aws-native:sso:PermissionSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PermissionSet resources.
type permissionSetState struct {
}

type PermissionSetState struct {
}

func (PermissionSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionSetState)(nil)).Elem()
}

type permissionSetArgs struct {
	// Specifies the names and paths of the customer managed policies that you have attached to your permission set.
	CustomerManagedPolicyReferences []PermissionSetCustomerManagedPolicyReference `pulumi:"customerManagedPolicyReferences"`
	// The permission set description.
	Description *string `pulumi:"description"`
	// The inline policy to put in permission set.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SSO::PermissionSet` for more information about the expected schema for this property.
	InlinePolicy interface{} `pulumi:"inlinePolicy"`
	// The sso instance arn that the permission set is owned.
	InstanceArn string `pulumi:"instanceArn"`
	// A structure that stores a list of managed policy ARNs that describe the associated AWS managed policy.
	ManagedPolicies []string `pulumi:"managedPolicies"`
	// The name you want to assign to this permission set.
	Name *string `pulumi:"name"`
	// Specifies the configuration of the AWS managed or customer managed policy that you want to set as a permissions boundary. Specify either `CustomerManagedPolicyReference` to use the name and path of a customer managed policy, or `ManagedPolicyArn` to use the ARN of an AWS managed policy. A permissions boundary represents the maximum permissions that any policy can grant your role. For more information, see [Permissions boundaries for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide* .
	//
	// > Policies used as permissions boundaries don't provide permissions. You must also attach an IAM policy to the role. To learn how the effective permissions for a role are evaluated, see [IAM JSON policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) in the *IAM User Guide* .
	PermissionsBoundary *PermissionSetPermissionsBoundary `pulumi:"permissionsBoundary"`
	// The relay state URL that redirect links to any service in the AWS Management Console.
	RelayStateType *string `pulumi:"relayStateType"`
	// The length of time that a user can be signed in to an AWS account.
	SessionDuration *string `pulumi:"sessionDuration"`
	// The tags to attach to the new `PermissionSet` .
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a PermissionSet resource.
type PermissionSetArgs struct {
	// Specifies the names and paths of the customer managed policies that you have attached to your permission set.
	CustomerManagedPolicyReferences PermissionSetCustomerManagedPolicyReferenceArrayInput
	// The permission set description.
	Description pulumi.StringPtrInput
	// The inline policy to put in permission set.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SSO::PermissionSet` for more information about the expected schema for this property.
	InlinePolicy pulumi.Input
	// The sso instance arn that the permission set is owned.
	InstanceArn pulumi.StringInput
	// A structure that stores a list of managed policy ARNs that describe the associated AWS managed policy.
	ManagedPolicies pulumi.StringArrayInput
	// The name you want to assign to this permission set.
	Name pulumi.StringPtrInput
	// Specifies the configuration of the AWS managed or customer managed policy that you want to set as a permissions boundary. Specify either `CustomerManagedPolicyReference` to use the name and path of a customer managed policy, or `ManagedPolicyArn` to use the ARN of an AWS managed policy. A permissions boundary represents the maximum permissions that any policy can grant your role. For more information, see [Permissions boundaries for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide* .
	//
	// > Policies used as permissions boundaries don't provide permissions. You must also attach an IAM policy to the role. To learn how the effective permissions for a role are evaluated, see [IAM JSON policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) in the *IAM User Guide* .
	PermissionsBoundary PermissionSetPermissionsBoundaryPtrInput
	// The relay state URL that redirect links to any service in the AWS Management Console.
	RelayStateType pulumi.StringPtrInput
	// The length of time that a user can be signed in to an AWS account.
	SessionDuration pulumi.StringPtrInput
	// The tags to attach to the new `PermissionSet` .
	Tags aws.TagArrayInput
}

func (PermissionSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionSetArgs)(nil)).Elem()
}

type PermissionSetInput interface {
	pulumi.Input

	ToPermissionSetOutput() PermissionSetOutput
	ToPermissionSetOutputWithContext(ctx context.Context) PermissionSetOutput
}

func (*PermissionSet) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionSet)(nil)).Elem()
}

func (i *PermissionSet) ToPermissionSetOutput() PermissionSetOutput {
	return i.ToPermissionSetOutputWithContext(context.Background())
}

func (i *PermissionSet) ToPermissionSetOutputWithContext(ctx context.Context) PermissionSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetOutput)
}

type PermissionSetOutput struct{ *pulumi.OutputState }

func (PermissionSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionSet)(nil)).Elem()
}

func (o PermissionSetOutput) ToPermissionSetOutput() PermissionSetOutput {
	return o
}

func (o PermissionSetOutput) ToPermissionSetOutputWithContext(ctx context.Context) PermissionSetOutput {
	return o
}

// Specifies the names and paths of the customer managed policies that you have attached to your permission set.
func (o PermissionSetOutput) CustomerManagedPolicyReferences() PermissionSetCustomerManagedPolicyReferenceArrayOutput {
	return o.ApplyT(func(v *PermissionSet) PermissionSetCustomerManagedPolicyReferenceArrayOutput {
		return v.CustomerManagedPolicyReferences
	}).(PermissionSetCustomerManagedPolicyReferenceArrayOutput)
}

// The permission set description.
func (o PermissionSetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PermissionSet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The inline policy to put in permission set.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SSO::PermissionSet` for more information about the expected schema for this property.
func (o PermissionSetOutput) InlinePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *PermissionSet) pulumi.AnyOutput { return v.InlinePolicy }).(pulumi.AnyOutput)
}

// The sso instance arn that the permission set is owned.
func (o PermissionSetOutput) InstanceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *PermissionSet) pulumi.StringOutput { return v.InstanceArn }).(pulumi.StringOutput)
}

// A structure that stores a list of managed policy ARNs that describe the associated AWS managed policy.
func (o PermissionSetOutput) ManagedPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PermissionSet) pulumi.StringArrayOutput { return v.ManagedPolicies }).(pulumi.StringArrayOutput)
}

// The name you want to assign to this permission set.
func (o PermissionSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PermissionSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The permission set that the policy will be attached to
func (o PermissionSetOutput) PermissionSetArn() pulumi.StringOutput {
	return o.ApplyT(func(v *PermissionSet) pulumi.StringOutput { return v.PermissionSetArn }).(pulumi.StringOutput)
}

// Specifies the configuration of the AWS managed or customer managed policy that you want to set as a permissions boundary. Specify either `CustomerManagedPolicyReference` to use the name and path of a customer managed policy, or `ManagedPolicyArn` to use the ARN of an AWS managed policy. A permissions boundary represents the maximum permissions that any policy can grant your role. For more information, see [Permissions boundaries for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide* .
//
// > Policies used as permissions boundaries don't provide permissions. You must also attach an IAM policy to the role. To learn how the effective permissions for a role are evaluated, see [IAM JSON policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) in the *IAM User Guide* .
func (o PermissionSetOutput) PermissionsBoundary() PermissionSetPermissionsBoundaryPtrOutput {
	return o.ApplyT(func(v *PermissionSet) PermissionSetPermissionsBoundaryPtrOutput { return v.PermissionsBoundary }).(PermissionSetPermissionsBoundaryPtrOutput)
}

// The relay state URL that redirect links to any service in the AWS Management Console.
func (o PermissionSetOutput) RelayStateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PermissionSet) pulumi.StringPtrOutput { return v.RelayStateType }).(pulumi.StringPtrOutput)
}

// The length of time that a user can be signed in to an AWS account.
func (o PermissionSetOutput) SessionDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PermissionSet) pulumi.StringPtrOutput { return v.SessionDuration }).(pulumi.StringPtrOutput)
}

// The tags to attach to the new `PermissionSet` .
func (o PermissionSetOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *PermissionSet) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionSetInput)(nil)).Elem(), &PermissionSet{})
	pulumi.RegisterOutputType(PermissionSetOutput{})
}
