// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WAFRegional::RateBasedRule
//
// Deprecated: RateBasedRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type RateBasedRule struct {
	pulumi.CustomResourceState

	MatchPredicates RateBasedRulePredicateArrayOutput `pulumi:"matchPredicates"`
	MetricName      pulumi.StringOutput               `pulumi:"metricName"`
	Name            pulumi.StringOutput               `pulumi:"name"`
	RateKey         pulumi.StringOutput               `pulumi:"rateKey"`
	RateLimit       pulumi.IntOutput                  `pulumi:"rateLimit"`
}

// NewRateBasedRule registers a new resource with the given unique name, arguments, and options.
func NewRateBasedRule(ctx *pulumi.Context,
	name string, args *RateBasedRuleArgs, opts ...pulumi.ResourceOption) (*RateBasedRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	if args.RateKey == nil {
		return nil, errors.New("invalid value for required argument 'RateKey'")
	}
	if args.RateLimit == nil {
		return nil, errors.New("invalid value for required argument 'RateLimit'")
	}
	var resource RateBasedRule
	err := ctx.RegisterResource("aws-native:wafregional:RateBasedRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRateBasedRule gets an existing RateBasedRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRateBasedRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RateBasedRuleState, opts ...pulumi.ResourceOption) (*RateBasedRule, error) {
	var resource RateBasedRule
	err := ctx.ReadResource("aws-native:wafregional:RateBasedRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RateBasedRule resources.
type rateBasedRuleState struct {
}

type RateBasedRuleState struct {
}

func (RateBasedRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rateBasedRuleState)(nil)).Elem()
}

type rateBasedRuleArgs struct {
	MatchPredicates []RateBasedRulePredicate `pulumi:"matchPredicates"`
	MetricName      string                   `pulumi:"metricName"`
	Name            *string                  `pulumi:"name"`
	RateKey         string                   `pulumi:"rateKey"`
	RateLimit       int                      `pulumi:"rateLimit"`
}

// The set of arguments for constructing a RateBasedRule resource.
type RateBasedRuleArgs struct {
	MatchPredicates RateBasedRulePredicateArrayInput
	MetricName      pulumi.StringInput
	Name            pulumi.StringPtrInput
	RateKey         pulumi.StringInput
	RateLimit       pulumi.IntInput
}

func (RateBasedRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rateBasedRuleArgs)(nil)).Elem()
}

type RateBasedRuleInput interface {
	pulumi.Input

	ToRateBasedRuleOutput() RateBasedRuleOutput
	ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput
}

func (*RateBasedRule) ElementType() reflect.Type {
	return reflect.TypeOf((**RateBasedRule)(nil)).Elem()
}

func (i *RateBasedRule) ToRateBasedRuleOutput() RateBasedRuleOutput {
	return i.ToRateBasedRuleOutputWithContext(context.Background())
}

func (i *RateBasedRule) ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateBasedRuleOutput)
}

type RateBasedRuleOutput struct{ *pulumi.OutputState }

func (RateBasedRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RateBasedRule)(nil)).Elem()
}

func (o RateBasedRuleOutput) ToRateBasedRuleOutput() RateBasedRuleOutput {
	return o
}

func (o RateBasedRuleOutput) ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput {
	return o
}

func (o RateBasedRuleOutput) MatchPredicates() RateBasedRulePredicateArrayOutput {
	return o.ApplyT(func(v *RateBasedRule) RateBasedRulePredicateArrayOutput { return v.MatchPredicates }).(RateBasedRulePredicateArrayOutput)
}

func (o RateBasedRuleOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v *RateBasedRule) pulumi.StringOutput { return v.MetricName }).(pulumi.StringOutput)
}

func (o RateBasedRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RateBasedRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RateBasedRuleOutput) RateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *RateBasedRule) pulumi.StringOutput { return v.RateKey }).(pulumi.StringOutput)
}

func (o RateBasedRuleOutput) RateLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *RateBasedRule) pulumi.IntOutput { return v.RateLimit }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RateBasedRuleInput)(nil)).Elem(), &RateBasedRule{})
	pulumi.RegisterOutputType(RateBasedRuleOutput{})
}