// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pcaconnectorad

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a template that defines certificate configurations, both for issuance and client handling
type Template struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html) .
	ConnectorArn pulumi.StringOutput `pulumi:"connectorArn"`
	// Template configuration to define the information included in certificates. Define certificate validity and renewal periods, certificate request handling and enrollment options, key usage extensions, application policies, and cryptography settings.
	Definition pulumi.AnyOutput `pulumi:"definition"`
	// Name of the templates. Template names must be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// This setting allows the major version of a template to be increased automatically. All members of Active Directory groups that are allowed to enroll with a template will receive a new certificate issued using that template.
	ReenrollAllCertificateHolders pulumi.BoolPtrOutput `pulumi:"reenrollAllCertificateHolders"`
	// Metadata assigned to a template consisting of a key-value pair.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html) .
	TemplateArn pulumi.StringOutput `pulumi:"templateArn"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectorArn == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorArn'")
	}
	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"connectorArn",
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Template
	err := ctx.RegisterResource("aws-native:pcaconnectorad:Template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	var resource Template
	err := ctx.ReadResource("aws-native:pcaconnectorad:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
}

type TemplateState struct {
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	// The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html) .
	ConnectorArn string `pulumi:"connectorArn"`
	// Template configuration to define the information included in certificates. Define certificate validity and renewal periods, certificate request handling and enrollment options, key usage extensions, application policies, and cryptography settings.
	Definition interface{} `pulumi:"definition"`
	// Name of the templates. Template names must be unique.
	Name *string `pulumi:"name"`
	// This setting allows the major version of a template to be increased automatically. All members of Active Directory groups that are allowed to enroll with a template will receive a new certificate issued using that template.
	ReenrollAllCertificateHolders *bool `pulumi:"reenrollAllCertificateHolders"`
	// Metadata assigned to a template consisting of a key-value pair.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	// The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html) .
	ConnectorArn pulumi.StringInput
	// Template configuration to define the information included in certificates. Define certificate validity and renewal periods, certificate request handling and enrollment options, key usage extensions, application policies, and cryptography settings.
	Definition pulumi.Input
	// Name of the templates. Template names must be unique.
	Name pulumi.StringPtrInput
	// This setting allows the major version of a template to be increased automatically. All members of Active Directory groups that are allowed to enroll with a template will receive a new certificate issued using that template.
	ReenrollAllCertificateHolders pulumi.BoolPtrInput
	// Metadata assigned to a template consisting of a key-value pair.
	Tags pulumi.StringMapInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}

type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(ctx context.Context) TemplateOutput
}

func (*Template) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (i *Template) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i *Template) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

type TemplateOutput struct{ *pulumi.OutputState }

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

// The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html) .
func (o TemplateOutput) ConnectorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.ConnectorArn }).(pulumi.StringOutput)
}

// Template configuration to define the information included in certificates. Define certificate validity and renewal periods, certificate request handling and enrollment options, key usage extensions, application policies, and cryptography settings.
func (o TemplateOutput) Definition() pulumi.AnyOutput {
	return o.ApplyT(func(v *Template) pulumi.AnyOutput { return v.Definition }).(pulumi.AnyOutput)
}

// Name of the templates. Template names must be unique.
func (o TemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// This setting allows the major version of a template to be increased automatically. All members of Active Directory groups that are allowed to enroll with a template will receive a new certificate issued using that template.
func (o TemplateOutput) ReenrollAllCertificateHolders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.BoolPtrOutput { return v.ReenrollAllCertificateHolders }).(pulumi.BoolPtrOutput)
}

// Metadata assigned to a template consisting of a key-value pair.
func (o TemplateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Template) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html) .
func (o TemplateOutput) TemplateArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateInput)(nil)).Elem(), &Template{})
	pulumi.RegisterOutputType(TemplateOutput{})
}
