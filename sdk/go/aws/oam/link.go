// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Oam::Link Resource Type
type Link struct {
	pulumi.CustomResourceState

	// The ARN of the link. For example, `arn:aws:oam:us-west-1:111111111111:link:abcd1234-a123-456a-a12b-a123b456c789`
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The friendly human-readable name used to identify this source account when it is viewed from the monitoring account. For example, `my-account1` .
	Label pulumi.StringOutput `pulumi:"label"`
	// Specify a friendly human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
	//
	// You can include the following variables in your template:
	//
	// - `$AccountName` is the name of the account
	// - `$AccountEmail` is a globally-unique email address, which includes the email domain, such as `mariagarcia@example.com`
	// - `$AccountEmailNoDomain` is an email address without the domain name, such as `mariagarcia`
	//
	// > In the  and  Regions, the only supported option is to use custom labels, and the `$AccountName` , `$AccountEmail` , and `$AccountEmailNoDomain` variables all resolve as *account-id* instead of the specified variable.
	LabelTemplate pulumi.StringPtrOutput `pulumi:"labelTemplate"`
	// Use this structure to optionally create filters that specify that only some metric namespaces or log groups are to be shared from the source account to the monitoring account.
	LinkConfiguration LinkConfigurationPtrOutput `pulumi:"linkConfiguration"`
	// An array of strings that define which types of data that the source account shares with the monitoring account. Valid values are `AWS::CloudWatch::Metric | AWS::Logs::LogGroup | AWS::XRay::Trace | AWS::ApplicationInsights::Application | AWS::InternetMonitor::Monitor` .
	ResourceTypes LinkResourceTypeArrayOutput `pulumi:"resourceTypes"`
	// The ARN of the sink in the monitoring account that you want to link to. You can use [ListSinks](https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html) to find the ARNs of sinks.
	SinkIdentifier pulumi.StringOutput `pulumi:"sinkIdentifier"`
	// Tags to apply to the link
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewLink registers a new resource with the given unique name, arguments, and options.
func NewLink(ctx *pulumi.Context,
	name string, args *LinkArgs, opts ...pulumi.ResourceOption) (*Link, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceTypes == nil {
		return nil, errors.New("invalid value for required argument 'ResourceTypes'")
	}
	if args.SinkIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'SinkIdentifier'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"labelTemplate",
		"sinkIdentifier",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Link
	err := ctx.RegisterResource("aws-native:oam:Link", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLink gets an existing Link resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkState, opts ...pulumi.ResourceOption) (*Link, error) {
	var resource Link
	err := ctx.ReadResource("aws-native:oam:Link", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Link resources.
type linkState struct {
}

type LinkState struct {
}

func (LinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkState)(nil)).Elem()
}

type linkArgs struct {
	// Specify a friendly human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
	//
	// You can include the following variables in your template:
	//
	// - `$AccountName` is the name of the account
	// - `$AccountEmail` is a globally-unique email address, which includes the email domain, such as `mariagarcia@example.com`
	// - `$AccountEmailNoDomain` is an email address without the domain name, such as `mariagarcia`
	//
	// > In the  and  Regions, the only supported option is to use custom labels, and the `$AccountName` , `$AccountEmail` , and `$AccountEmailNoDomain` variables all resolve as *account-id* instead of the specified variable.
	LabelTemplate *string `pulumi:"labelTemplate"`
	// Use this structure to optionally create filters that specify that only some metric namespaces or log groups are to be shared from the source account to the monitoring account.
	LinkConfiguration *LinkConfiguration `pulumi:"linkConfiguration"`
	// An array of strings that define which types of data that the source account shares with the monitoring account. Valid values are `AWS::CloudWatch::Metric | AWS::Logs::LogGroup | AWS::XRay::Trace | AWS::ApplicationInsights::Application | AWS::InternetMonitor::Monitor` .
	ResourceTypes []LinkResourceType `pulumi:"resourceTypes"`
	// The ARN of the sink in the monitoring account that you want to link to. You can use [ListSinks](https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html) to find the ARNs of sinks.
	SinkIdentifier string `pulumi:"sinkIdentifier"`
	// Tags to apply to the link
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Link resource.
type LinkArgs struct {
	// Specify a friendly human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
	//
	// You can include the following variables in your template:
	//
	// - `$AccountName` is the name of the account
	// - `$AccountEmail` is a globally-unique email address, which includes the email domain, such as `mariagarcia@example.com`
	// - `$AccountEmailNoDomain` is an email address without the domain name, such as `mariagarcia`
	//
	// > In the  and  Regions, the only supported option is to use custom labels, and the `$AccountName` , `$AccountEmail` , and `$AccountEmailNoDomain` variables all resolve as *account-id* instead of the specified variable.
	LabelTemplate pulumi.StringPtrInput
	// Use this structure to optionally create filters that specify that only some metric namespaces or log groups are to be shared from the source account to the monitoring account.
	LinkConfiguration LinkConfigurationPtrInput
	// An array of strings that define which types of data that the source account shares with the monitoring account. Valid values are `AWS::CloudWatch::Metric | AWS::Logs::LogGroup | AWS::XRay::Trace | AWS::ApplicationInsights::Application | AWS::InternetMonitor::Monitor` .
	ResourceTypes LinkResourceTypeArrayInput
	// The ARN of the sink in the monitoring account that you want to link to. You can use [ListSinks](https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html) to find the ARNs of sinks.
	SinkIdentifier pulumi.StringInput
	// Tags to apply to the link
	Tags pulumi.StringMapInput
}

func (LinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkArgs)(nil)).Elem()
}

type LinkInput interface {
	pulumi.Input

	ToLinkOutput() LinkOutput
	ToLinkOutputWithContext(ctx context.Context) LinkOutput
}

func (*Link) ElementType() reflect.Type {
	return reflect.TypeOf((**Link)(nil)).Elem()
}

func (i *Link) ToLinkOutput() LinkOutput {
	return i.ToLinkOutputWithContext(context.Background())
}

func (i *Link) ToLinkOutputWithContext(ctx context.Context) LinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkOutput)
}

type LinkOutput struct{ *pulumi.OutputState }

func (LinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Link)(nil)).Elem()
}

func (o LinkOutput) ToLinkOutput() LinkOutput {
	return o
}

func (o LinkOutput) ToLinkOutputWithContext(ctx context.Context) LinkOutput {
	return o
}

// The ARN of the link. For example, `arn:aws:oam:us-west-1:111111111111:link:abcd1234-a123-456a-a12b-a123b456c789`
func (o LinkOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The friendly human-readable name used to identify this source account when it is viewed from the monitoring account. For example, `my-account1` .
func (o LinkOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Label }).(pulumi.StringOutput)
}

// Specify a friendly human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
//
// You can include the following variables in your template:
//
// - `$AccountName` is the name of the account
// - `$AccountEmail` is a globally-unique email address, which includes the email domain, such as `mariagarcia@example.com`
// - `$AccountEmailNoDomain` is an email address without the domain name, such as `mariagarcia`
//
// > In the  and  Regions, the only supported option is to use custom labels, and the `$AccountName` , `$AccountEmail` , and `$AccountEmailNoDomain` variables all resolve as *account-id* instead of the specified variable.
func (o LinkOutput) LabelTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Link) pulumi.StringPtrOutput { return v.LabelTemplate }).(pulumi.StringPtrOutput)
}

// Use this structure to optionally create filters that specify that only some metric namespaces or log groups are to be shared from the source account to the monitoring account.
func (o LinkOutput) LinkConfiguration() LinkConfigurationPtrOutput {
	return o.ApplyT(func(v *Link) LinkConfigurationPtrOutput { return v.LinkConfiguration }).(LinkConfigurationPtrOutput)
}

// An array of strings that define which types of data that the source account shares with the monitoring account. Valid values are `AWS::CloudWatch::Metric | AWS::Logs::LogGroup | AWS::XRay::Trace | AWS::ApplicationInsights::Application | AWS::InternetMonitor::Monitor` .
func (o LinkOutput) ResourceTypes() LinkResourceTypeArrayOutput {
	return o.ApplyT(func(v *Link) LinkResourceTypeArrayOutput { return v.ResourceTypes }).(LinkResourceTypeArrayOutput)
}

// The ARN of the sink in the monitoring account that you want to link to. You can use [ListSinks](https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html) to find the ARNs of sinks.
func (o LinkOutput) SinkIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.SinkIdentifier }).(pulumi.StringOutput)
}

// Tags to apply to the link
func (o LinkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Link) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkInput)(nil)).Elem(), &Link{})
	pulumi.RegisterOutputType(LinkOutput{})
}
