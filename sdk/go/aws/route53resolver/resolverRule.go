// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Route53Resolver::ResolverRule
type ResolverRule struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the resolver rule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// The name for the Resolver rule
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The ID of the endpoint that the rule is associated with.
	ResolverEndpointId pulumi.StringPtrOutput `pulumi:"resolverEndpointId"`
	// The ID of the endpoint that the rule is associated with.
	ResolverRuleId pulumi.StringOutput `pulumi:"resolverRuleId"`
	// When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.
	RuleType ResolverRuleRuleTypeOutput `pulumi:"ruleType"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
	TargetIps ResolverRuleTargetAddressArrayOutput `pulumi:"targetIps"`
}

// NewResolverRule registers a new resource with the given unique name, arguments, and options.
func NewResolverRule(ctx *pulumi.Context,
	name string, args *ResolverRuleArgs, opts ...pulumi.ResourceOption) (*ResolverRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RuleType == nil {
		return nil, errors.New("invalid value for required argument 'RuleType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"ruleType",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResolverRule
	err := ctx.RegisterResource("aws-native:route53resolver:ResolverRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverRule gets an existing ResolverRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverRuleState, opts ...pulumi.ResourceOption) (*ResolverRule, error) {
	var resource ResolverRule
	err := ctx.ReadResource("aws-native:route53resolver:ResolverRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverRule resources.
type resolverRuleState struct {
}

type ResolverRuleState struct {
}

func (ResolverRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleState)(nil)).Elem()
}

type resolverRuleArgs struct {
	// DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps
	DomainName *string `pulumi:"domainName"`
	// The name for the Resolver rule
	Name *string `pulumi:"name"`
	// The ID of the endpoint that the rule is associated with.
	ResolverEndpointId *string `pulumi:"resolverEndpointId"`
	// When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.
	RuleType ResolverRuleRuleType `pulumi:"ruleType"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
	TargetIps []ResolverRuleTargetAddress `pulumi:"targetIps"`
}

// The set of arguments for constructing a ResolverRule resource.
type ResolverRuleArgs struct {
	// DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps
	DomainName pulumi.StringPtrInput
	// The name for the Resolver rule
	Name pulumi.StringPtrInput
	// The ID of the endpoint that the rule is associated with.
	ResolverEndpointId pulumi.StringPtrInput
	// When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.
	RuleType ResolverRuleRuleTypeInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayInput
	// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
	TargetIps ResolverRuleTargetAddressArrayInput
}

func (ResolverRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleArgs)(nil)).Elem()
}

type ResolverRuleInput interface {
	pulumi.Input

	ToResolverRuleOutput() ResolverRuleOutput
	ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput
}

func (*ResolverRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverRule)(nil)).Elem()
}

func (i *ResolverRule) ToResolverRuleOutput() ResolverRuleOutput {
	return i.ToResolverRuleOutputWithContext(context.Background())
}

func (i *ResolverRule) ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverRuleOutput)
}

type ResolverRuleOutput struct{ *pulumi.OutputState }

func (ResolverRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverRule)(nil)).Elem()
}

func (o ResolverRuleOutput) ToResolverRuleOutput() ResolverRuleOutput {
	return o
}

func (o ResolverRuleOutput) ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput {
	return o
}

// The Amazon Resource Name (ARN) of the resolver rule.
func (o ResolverRuleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverRule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps
func (o ResolverRuleOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverRule) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The name for the Resolver rule
func (o ResolverRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the endpoint that the rule is associated with.
func (o ResolverRuleOutput) ResolverEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverRule) pulumi.StringPtrOutput { return v.ResolverEndpointId }).(pulumi.StringPtrOutput)
}

// The ID of the endpoint that the rule is associated with.
func (o ResolverRuleOutput) ResolverRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverRule) pulumi.StringOutput { return v.ResolverRuleId }).(pulumi.StringOutput)
}

// When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.
func (o ResolverRuleOutput) RuleType() ResolverRuleRuleTypeOutput {
	return o.ApplyT(func(v *ResolverRule) ResolverRuleRuleTypeOutput { return v.RuleType }).(ResolverRuleRuleTypeOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ResolverRuleOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *ResolverRule) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
func (o ResolverRuleOutput) TargetIps() ResolverRuleTargetAddressArrayOutput {
	return o.ApplyT(func(v *ResolverRule) ResolverRuleTargetAddressArrayOutput { return v.TargetIps }).(ResolverRuleTargetAddressArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverRuleInput)(nil)).Elem(), &ResolverRule{})
	pulumi.RegisterOutputType(ResolverRuleOutput{})
}
