// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceCatalog::LaunchRoleConstraint
//
// Deprecated: LaunchRoleConstraint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type LaunchRoleConstraint struct {
	pulumi.CustomResourceState

	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	Description    pulumi.StringPtrOutput `pulumi:"description"`
	LocalRoleName  pulumi.StringPtrOutput `pulumi:"localRoleName"`
	PortfolioId    pulumi.StringOutput    `pulumi:"portfolioId"`
	ProductId      pulumi.StringOutput    `pulumi:"productId"`
	RoleArn        pulumi.StringPtrOutput `pulumi:"roleArn"`
}

// NewLaunchRoleConstraint registers a new resource with the given unique name, arguments, and options.
func NewLaunchRoleConstraint(ctx *pulumi.Context,
	name string, args *LaunchRoleConstraintArgs, opts ...pulumi.ResourceOption) (*LaunchRoleConstraint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortfolioId == nil {
		return nil, errors.New("invalid value for required argument 'PortfolioId'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	var resource LaunchRoleConstraint
	err := ctx.RegisterResource("aws-native:servicecatalog:LaunchRoleConstraint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchRoleConstraint gets an existing LaunchRoleConstraint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchRoleConstraint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchRoleConstraintState, opts ...pulumi.ResourceOption) (*LaunchRoleConstraint, error) {
	var resource LaunchRoleConstraint
	err := ctx.ReadResource("aws-native:servicecatalog:LaunchRoleConstraint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchRoleConstraint resources.
type launchRoleConstraintState struct {
}

type LaunchRoleConstraintState struct {
}

func (LaunchRoleConstraintState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchRoleConstraintState)(nil)).Elem()
}

type launchRoleConstraintArgs struct {
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	Description    *string `pulumi:"description"`
	LocalRoleName  *string `pulumi:"localRoleName"`
	PortfolioId    string  `pulumi:"portfolioId"`
	ProductId      string  `pulumi:"productId"`
	RoleArn        *string `pulumi:"roleArn"`
}

// The set of arguments for constructing a LaunchRoleConstraint resource.
type LaunchRoleConstraintArgs struct {
	AcceptLanguage pulumi.StringPtrInput
	Description    pulumi.StringPtrInput
	LocalRoleName  pulumi.StringPtrInput
	PortfolioId    pulumi.StringInput
	ProductId      pulumi.StringInput
	RoleArn        pulumi.StringPtrInput
}

func (LaunchRoleConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchRoleConstraintArgs)(nil)).Elem()
}

type LaunchRoleConstraintInput interface {
	pulumi.Input

	ToLaunchRoleConstraintOutput() LaunchRoleConstraintOutput
	ToLaunchRoleConstraintOutputWithContext(ctx context.Context) LaunchRoleConstraintOutput
}

func (*LaunchRoleConstraint) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchRoleConstraint)(nil)).Elem()
}

func (i *LaunchRoleConstraint) ToLaunchRoleConstraintOutput() LaunchRoleConstraintOutput {
	return i.ToLaunchRoleConstraintOutputWithContext(context.Background())
}

func (i *LaunchRoleConstraint) ToLaunchRoleConstraintOutputWithContext(ctx context.Context) LaunchRoleConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchRoleConstraintOutput)
}

type LaunchRoleConstraintOutput struct{ *pulumi.OutputState }

func (LaunchRoleConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchRoleConstraint)(nil)).Elem()
}

func (o LaunchRoleConstraintOutput) ToLaunchRoleConstraintOutput() LaunchRoleConstraintOutput {
	return o
}

func (o LaunchRoleConstraintOutput) ToLaunchRoleConstraintOutputWithContext(ctx context.Context) LaunchRoleConstraintOutput {
	return o
}

func (o LaunchRoleConstraintOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchRoleConstraint) pulumi.StringPtrOutput { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

func (o LaunchRoleConstraintOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchRoleConstraint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LaunchRoleConstraintOutput) LocalRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchRoleConstraint) pulumi.StringPtrOutput { return v.LocalRoleName }).(pulumi.StringPtrOutput)
}

func (o LaunchRoleConstraintOutput) PortfolioId() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchRoleConstraint) pulumi.StringOutput { return v.PortfolioId }).(pulumi.StringOutput)
}

func (o LaunchRoleConstraintOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchRoleConstraint) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

func (o LaunchRoleConstraintOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchRoleConstraint) pulumi.StringPtrOutput { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchRoleConstraintInput)(nil)).Elem(), &LaunchRoleConstraint{})
	pulumi.RegisterOutputType(LaunchRoleConstraintOutput{})
}