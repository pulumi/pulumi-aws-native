// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package detective

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Detective::OrganizationAdmin
type OrganizationAdmin struct {
	pulumi.CustomResourceState

	// The account ID of the account that should be registered as your Organization's delegated administrator for Detective
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The Detective graph ARN
	GraphArn pulumi.StringOutput `pulumi:"graphArn"`
}

// NewOrganizationAdmin registers a new resource with the given unique name, arguments, and options.
func NewOrganizationAdmin(ctx *pulumi.Context,
	name string, args *OrganizationAdminArgs, opts ...pulumi.ResourceOption) (*OrganizationAdmin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"accountId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrganizationAdmin
	err := ctx.RegisterResource("aws-native:detective:OrganizationAdmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationAdmin gets an existing OrganizationAdmin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationAdmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationAdminState, opts ...pulumi.ResourceOption) (*OrganizationAdmin, error) {
	var resource OrganizationAdmin
	err := ctx.ReadResource("aws-native:detective:OrganizationAdmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationAdmin resources.
type organizationAdminState struct {
}

type OrganizationAdminState struct {
}

func (OrganizationAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationAdminState)(nil)).Elem()
}

type organizationAdminArgs struct {
	// The account ID of the account that should be registered as your Organization's delegated administrator for Detective
	AccountId string `pulumi:"accountId"`
}

// The set of arguments for constructing a OrganizationAdmin resource.
type OrganizationAdminArgs struct {
	// The account ID of the account that should be registered as your Organization's delegated administrator for Detective
	AccountId pulumi.StringInput
}

func (OrganizationAdminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationAdminArgs)(nil)).Elem()
}

type OrganizationAdminInput interface {
	pulumi.Input

	ToOrganizationAdminOutput() OrganizationAdminOutput
	ToOrganizationAdminOutputWithContext(ctx context.Context) OrganizationAdminOutput
}

func (*OrganizationAdmin) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationAdmin)(nil)).Elem()
}

func (i *OrganizationAdmin) ToOrganizationAdminOutput() OrganizationAdminOutput {
	return i.ToOrganizationAdminOutputWithContext(context.Background())
}

func (i *OrganizationAdmin) ToOrganizationAdminOutputWithContext(ctx context.Context) OrganizationAdminOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminOutput)
}

type OrganizationAdminOutput struct{ *pulumi.OutputState }

func (OrganizationAdminOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationAdmin)(nil)).Elem()
}

func (o OrganizationAdminOutput) ToOrganizationAdminOutput() OrganizationAdminOutput {
	return o
}

func (o OrganizationAdminOutput) ToOrganizationAdminOutputWithContext(ctx context.Context) OrganizationAdminOutput {
	return o
}

// The account ID of the account that should be registered as your Organization's delegated administrator for Detective
func (o OrganizationAdminOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationAdmin) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The Detective graph ARN
func (o OrganizationAdminOutput) GraphArn() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationAdmin) pulumi.StringOutput { return v.GraphArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationAdminInput)(nil)).Elem(), &OrganizationAdmin{})
	pulumi.RegisterOutputType(OrganizationAdminOutput{})
}
