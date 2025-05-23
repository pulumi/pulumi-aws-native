// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPCBlockPublicAccessOptions
type VpcBlockPublicAccessOptions struct {
	pulumi.CustomResourceState

	// The identifier for the specified AWS account.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The desired Block Public Access mode for Internet Gateways in your account. We do not allow to create in a off mode as this is the default value
	InternetGatewayBlockMode VpcBlockPublicAccessOptionsInternetGatewayBlockModeOutput `pulumi:"internetGatewayBlockMode"`
}

// NewVpcBlockPublicAccessOptions registers a new resource with the given unique name, arguments, and options.
func NewVpcBlockPublicAccessOptions(ctx *pulumi.Context,
	name string, args *VpcBlockPublicAccessOptionsArgs, opts ...pulumi.ResourceOption) (*VpcBlockPublicAccessOptions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InternetGatewayBlockMode == nil {
		return nil, errors.New("invalid value for required argument 'InternetGatewayBlockMode'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcBlockPublicAccessOptions
	err := ctx.RegisterResource("aws-native:ec2:VpcBlockPublicAccessOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcBlockPublicAccessOptions gets an existing VpcBlockPublicAccessOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcBlockPublicAccessOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcBlockPublicAccessOptionsState, opts ...pulumi.ResourceOption) (*VpcBlockPublicAccessOptions, error) {
	var resource VpcBlockPublicAccessOptions
	err := ctx.ReadResource("aws-native:ec2:VpcBlockPublicAccessOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcBlockPublicAccessOptions resources.
type vpcBlockPublicAccessOptionsState struct {
}

type VpcBlockPublicAccessOptionsState struct {
}

func (VpcBlockPublicAccessOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcBlockPublicAccessOptionsState)(nil)).Elem()
}

type vpcBlockPublicAccessOptionsArgs struct {
	// The desired Block Public Access mode for Internet Gateways in your account. We do not allow to create in a off mode as this is the default value
	InternetGatewayBlockMode VpcBlockPublicAccessOptionsInternetGatewayBlockMode `pulumi:"internetGatewayBlockMode"`
}

// The set of arguments for constructing a VpcBlockPublicAccessOptions resource.
type VpcBlockPublicAccessOptionsArgs struct {
	// The desired Block Public Access mode for Internet Gateways in your account. We do not allow to create in a off mode as this is the default value
	InternetGatewayBlockMode VpcBlockPublicAccessOptionsInternetGatewayBlockModeInput
}

func (VpcBlockPublicAccessOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcBlockPublicAccessOptionsArgs)(nil)).Elem()
}

type VpcBlockPublicAccessOptionsInput interface {
	pulumi.Input

	ToVpcBlockPublicAccessOptionsOutput() VpcBlockPublicAccessOptionsOutput
	ToVpcBlockPublicAccessOptionsOutputWithContext(ctx context.Context) VpcBlockPublicAccessOptionsOutput
}

func (*VpcBlockPublicAccessOptions) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcBlockPublicAccessOptions)(nil)).Elem()
}

func (i *VpcBlockPublicAccessOptions) ToVpcBlockPublicAccessOptionsOutput() VpcBlockPublicAccessOptionsOutput {
	return i.ToVpcBlockPublicAccessOptionsOutputWithContext(context.Background())
}

func (i *VpcBlockPublicAccessOptions) ToVpcBlockPublicAccessOptionsOutputWithContext(ctx context.Context) VpcBlockPublicAccessOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcBlockPublicAccessOptionsOutput)
}

type VpcBlockPublicAccessOptionsOutput struct{ *pulumi.OutputState }

func (VpcBlockPublicAccessOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcBlockPublicAccessOptions)(nil)).Elem()
}

func (o VpcBlockPublicAccessOptionsOutput) ToVpcBlockPublicAccessOptionsOutput() VpcBlockPublicAccessOptionsOutput {
	return o
}

func (o VpcBlockPublicAccessOptionsOutput) ToVpcBlockPublicAccessOptionsOutputWithContext(ctx context.Context) VpcBlockPublicAccessOptionsOutput {
	return o
}

// The identifier for the specified AWS account.
func (o VpcBlockPublicAccessOptionsOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcBlockPublicAccessOptions) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The desired Block Public Access mode for Internet Gateways in your account. We do not allow to create in a off mode as this is the default value
func (o VpcBlockPublicAccessOptionsOutput) InternetGatewayBlockMode() VpcBlockPublicAccessOptionsInternetGatewayBlockModeOutput {
	return o.ApplyT(func(v *VpcBlockPublicAccessOptions) VpcBlockPublicAccessOptionsInternetGatewayBlockModeOutput {
		return v.InternetGatewayBlockMode
	}).(VpcBlockPublicAccessOptionsInternetGatewayBlockModeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcBlockPublicAccessOptionsInput)(nil)).Elem(), &VpcBlockPublicAccessOptions{})
	pulumi.RegisterOutputType(VpcBlockPublicAccessOptionsOutput{})
}
