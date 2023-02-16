// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssmcontacts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSMContacts::Contact
type Contact struct {
	pulumi.CustomResourceState

	// Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// The Amazon Resource Name (ARN) of the contact.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
	Plan ContactStageArrayOutput `pulumi:"plan"`
	// Contact type, which specify type of contact. Currently supported values: “PERSONAL”, “SHARED”, “OTHER“.
	Type ContactTypeOutput `pulumi:"type"`
}

// NewContact registers a new resource with the given unique name, arguments, and options.
func NewContact(ctx *pulumi.Context,
	name string, args *ContactArgs, opts ...pulumi.ResourceOption) (*Contact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Contact
	err := ctx.RegisterResource("aws-native:ssmcontacts:Contact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContact gets an existing Contact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactState, opts ...pulumi.ResourceOption) (*Contact, error) {
	var resource Contact
	err := ctx.ReadResource("aws-native:ssmcontacts:Contact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Contact resources.
type contactState struct {
}

type ContactState struct {
}

func (ContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactState)(nil)).Elem()
}

type contactArgs struct {
	// Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.
	Alias string `pulumi:"alias"`
	// Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.
	DisplayName string `pulumi:"displayName"`
	// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
	Plan []ContactStage `pulumi:"plan"`
	// Contact type, which specify type of contact. Currently supported values: “PERSONAL”, “SHARED”, “OTHER“.
	Type ContactType `pulumi:"type"`
}

// The set of arguments for constructing a Contact resource.
type ContactArgs struct {
	// Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.
	Alias pulumi.StringInput
	// Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.
	DisplayName pulumi.StringInput
	// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
	Plan ContactStageArrayInput
	// Contact type, which specify type of contact. Currently supported values: “PERSONAL”, “SHARED”, “OTHER“.
	Type ContactTypeInput
}

func (ContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactArgs)(nil)).Elem()
}

type ContactInput interface {
	pulumi.Input

	ToContactOutput() ContactOutput
	ToContactOutputWithContext(ctx context.Context) ContactOutput
}

func (*Contact) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (i *Contact) ToContactOutput() ContactOutput {
	return i.ToContactOutputWithContext(context.Background())
}

func (i *Contact) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput)
}

type ContactOutput struct{ *pulumi.OutputState }

func (ContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (o ContactOutput) ToContactOutput() ContactOutput {
	return o
}

func (o ContactOutput) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return o
}

// Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.
func (o ContactOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the contact.
func (o ContactOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.
func (o ContactOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
func (o ContactOutput) Plan() ContactStageArrayOutput {
	return o.ApplyT(func(v *Contact) ContactStageArrayOutput { return v.Plan }).(ContactStageArrayOutput)
}

// Contact type, which specify type of contact. Currently supported values: “PERSONAL”, “SHARED”, “OTHER“.
func (o ContactOutput) Type() ContactTypeOutput {
	return o.ApplyT(func(v *Contact) ContactTypeOutput { return v.Type }).(ContactTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContactInput)(nil)).Elem(), &Contact{})
	pulumi.RegisterOutputType(ContactOutput{})
}