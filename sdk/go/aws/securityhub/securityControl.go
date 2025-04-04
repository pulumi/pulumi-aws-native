// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A security control in Security Hub describes a security best practice related to a specific resource.
type SecurityControl struct {
	pulumi.CustomResourceState

	// The most recent reason for updating the customizable properties of a security control. This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.
	LastUpdateReason pulumi.StringPtrOutput `pulumi:"lastUpdateReason"`
	// An object that identifies the name of a control parameter, its current value, and whether it has been customized.
	Parameters SecurityControlParameterConfigurationMapOutput `pulumi:"parameters"`
	// The Amazon Resource Name (ARN) for a security control across standards, such as `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1`. This parameter doesn't mention a specific standard.
	SecurityControlArn pulumi.StringPtrOutput `pulumi:"securityControlArn"`
	// The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.
	SecurityControlId pulumi.StringPtrOutput `pulumi:"securityControlId"`
}

// NewSecurityControl registers a new resource with the given unique name, arguments, and options.
func NewSecurityControl(ctx *pulumi.Context,
	name string, args *SecurityControlArgs, opts ...pulumi.ResourceOption) (*SecurityControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"securityControlId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityControl
	err := ctx.RegisterResource("aws-native:securityhub:SecurityControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityControl gets an existing SecurityControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityControlState, opts ...pulumi.ResourceOption) (*SecurityControl, error) {
	var resource SecurityControl
	err := ctx.ReadResource("aws-native:securityhub:SecurityControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityControl resources.
type securityControlState struct {
}

type SecurityControlState struct {
}

func (SecurityControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityControlState)(nil)).Elem()
}

type securityControlArgs struct {
	// The most recent reason for updating the customizable properties of a security control. This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.
	LastUpdateReason *string `pulumi:"lastUpdateReason"`
	// An object that identifies the name of a control parameter, its current value, and whether it has been customized.
	Parameters map[string]SecurityControlParameterConfiguration `pulumi:"parameters"`
	// The Amazon Resource Name (ARN) for a security control across standards, such as `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1`. This parameter doesn't mention a specific standard.
	SecurityControlArn *string `pulumi:"securityControlArn"`
	// The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.
	SecurityControlId *string `pulumi:"securityControlId"`
}

// The set of arguments for constructing a SecurityControl resource.
type SecurityControlArgs struct {
	// The most recent reason for updating the customizable properties of a security control. This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.
	LastUpdateReason pulumi.StringPtrInput
	// An object that identifies the name of a control parameter, its current value, and whether it has been customized.
	Parameters SecurityControlParameterConfigurationMapInput
	// The Amazon Resource Name (ARN) for a security control across standards, such as `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1`. This parameter doesn't mention a specific standard.
	SecurityControlArn pulumi.StringPtrInput
	// The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.
	SecurityControlId pulumi.StringPtrInput
}

func (SecurityControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityControlArgs)(nil)).Elem()
}

type SecurityControlInput interface {
	pulumi.Input

	ToSecurityControlOutput() SecurityControlOutput
	ToSecurityControlOutputWithContext(ctx context.Context) SecurityControlOutput
}

func (*SecurityControl) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityControl)(nil)).Elem()
}

func (i *SecurityControl) ToSecurityControlOutput() SecurityControlOutput {
	return i.ToSecurityControlOutputWithContext(context.Background())
}

func (i *SecurityControl) ToSecurityControlOutputWithContext(ctx context.Context) SecurityControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityControlOutput)
}

type SecurityControlOutput struct{ *pulumi.OutputState }

func (SecurityControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityControl)(nil)).Elem()
}

func (o SecurityControlOutput) ToSecurityControlOutput() SecurityControlOutput {
	return o
}

func (o SecurityControlOutput) ToSecurityControlOutputWithContext(ctx context.Context) SecurityControlOutput {
	return o
}

// The most recent reason for updating the customizable properties of a security control. This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.
func (o SecurityControlOutput) LastUpdateReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityControl) pulumi.StringPtrOutput { return v.LastUpdateReason }).(pulumi.StringPtrOutput)
}

// An object that identifies the name of a control parameter, its current value, and whether it has been customized.
func (o SecurityControlOutput) Parameters() SecurityControlParameterConfigurationMapOutput {
	return o.ApplyT(func(v *SecurityControl) SecurityControlParameterConfigurationMapOutput { return v.Parameters }).(SecurityControlParameterConfigurationMapOutput)
}

// The Amazon Resource Name (ARN) for a security control across standards, such as `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1`. This parameter doesn't mention a specific standard.
func (o SecurityControlOutput) SecurityControlArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityControl) pulumi.StringPtrOutput { return v.SecurityControlArn }).(pulumi.StringPtrOutput)
}

// The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.
func (o SecurityControlOutput) SecurityControlId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityControl) pulumi.StringPtrOutput { return v.SecurityControlId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityControlInput)(nil)).Elem(), &SecurityControl{})
	pulumi.RegisterOutputType(SecurityControlOutput{})
}
