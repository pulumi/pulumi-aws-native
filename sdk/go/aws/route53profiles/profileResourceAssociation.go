// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53profiles

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Route53Profiles::ProfileResourceAssociation
type ProfileResourceAssociation struct {
	pulumi.CustomResourceState

	// Primary Identifier for  Profile Resource Association
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The name of an association between the  Profile and resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the  profile that you associated the resource to that is specified by ResourceArn.
	ProfileId pulumi.StringOutput `pulumi:"profileId"`
	// The arn of the resource that you associated to the  Profile.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// A JSON-formatted string with key-value pairs specifying the properties of the associated resource.
	ResourceProperties pulumi.StringPtrOutput `pulumi:"resourceProperties"`
	// The type of the resource associated to the  Profile.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
}

// NewProfileResourceAssociation registers a new resource with the given unique name, arguments, and options.
func NewProfileResourceAssociation(ctx *pulumi.Context,
	name string, args *ProfileResourceAssociationArgs, opts ...pulumi.ResourceOption) (*ProfileResourceAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileId == nil {
		return nil, errors.New("invalid value for required argument 'ProfileId'")
	}
	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"profileId",
		"resourceArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProfileResourceAssociation
	err := ctx.RegisterResource("aws-native:route53profiles:ProfileResourceAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfileResourceAssociation gets an existing ProfileResourceAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfileResourceAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileResourceAssociationState, opts ...pulumi.ResourceOption) (*ProfileResourceAssociation, error) {
	var resource ProfileResourceAssociation
	err := ctx.ReadResource("aws-native:route53profiles:ProfileResourceAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfileResourceAssociation resources.
type profileResourceAssociationState struct {
}

type ProfileResourceAssociationState struct {
}

func (ProfileResourceAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileResourceAssociationState)(nil)).Elem()
}

type profileResourceAssociationArgs struct {
	// The name of an association between the  Profile and resource.
	Name *string `pulumi:"name"`
	// The ID of the  profile that you associated the resource to that is specified by ResourceArn.
	ProfileId string `pulumi:"profileId"`
	// The arn of the resource that you associated to the  Profile.
	ResourceArn string `pulumi:"resourceArn"`
	// A JSON-formatted string with key-value pairs specifying the properties of the associated resource.
	ResourceProperties *string `pulumi:"resourceProperties"`
}

// The set of arguments for constructing a ProfileResourceAssociation resource.
type ProfileResourceAssociationArgs struct {
	// The name of an association between the  Profile and resource.
	Name pulumi.StringPtrInput
	// The ID of the  profile that you associated the resource to that is specified by ResourceArn.
	ProfileId pulumi.StringInput
	// The arn of the resource that you associated to the  Profile.
	ResourceArn pulumi.StringInput
	// A JSON-formatted string with key-value pairs specifying the properties of the associated resource.
	ResourceProperties pulumi.StringPtrInput
}

func (ProfileResourceAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileResourceAssociationArgs)(nil)).Elem()
}

type ProfileResourceAssociationInput interface {
	pulumi.Input

	ToProfileResourceAssociationOutput() ProfileResourceAssociationOutput
	ToProfileResourceAssociationOutputWithContext(ctx context.Context) ProfileResourceAssociationOutput
}

func (*ProfileResourceAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileResourceAssociation)(nil)).Elem()
}

func (i *ProfileResourceAssociation) ToProfileResourceAssociationOutput() ProfileResourceAssociationOutput {
	return i.ToProfileResourceAssociationOutputWithContext(context.Background())
}

func (i *ProfileResourceAssociation) ToProfileResourceAssociationOutputWithContext(ctx context.Context) ProfileResourceAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileResourceAssociationOutput)
}

type ProfileResourceAssociationOutput struct{ *pulumi.OutputState }

func (ProfileResourceAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileResourceAssociation)(nil)).Elem()
}

func (o ProfileResourceAssociationOutput) ToProfileResourceAssociationOutput() ProfileResourceAssociationOutput {
	return o
}

func (o ProfileResourceAssociationOutput) ToProfileResourceAssociationOutputWithContext(ctx context.Context) ProfileResourceAssociationOutput {
	return o
}

// Primary Identifier for  Profile Resource Association
func (o ProfileResourceAssociationOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileResourceAssociation) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The name of an association between the  Profile and resource.
func (o ProfileResourceAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileResourceAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the  profile that you associated the resource to that is specified by ResourceArn.
func (o ProfileResourceAssociationOutput) ProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileResourceAssociation) pulumi.StringOutput { return v.ProfileId }).(pulumi.StringOutput)
}

// The arn of the resource that you associated to the  Profile.
func (o ProfileResourceAssociationOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileResourceAssociation) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

// A JSON-formatted string with key-value pairs specifying the properties of the associated resource.
func (o ProfileResourceAssociationOutput) ResourceProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfileResourceAssociation) pulumi.StringPtrOutput { return v.ResourceProperties }).(pulumi.StringPtrOutput)
}

// The type of the resource associated to the  Profile.
func (o ProfileResourceAssociationOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileResourceAssociation) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileResourceAssociationInput)(nil)).Elem(), &ProfileResourceAssociation{})
	pulumi.RegisterOutputType(ProfileResourceAssociationOutput{})
}