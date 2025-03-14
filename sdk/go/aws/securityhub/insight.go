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

// The AWS::SecurityHub::Insight resource represents the AWS Security Hub Insight in your account. An AWS Security Hub insight is a collection of related findings.
type Insight struct {
	pulumi.CustomResourceState

	// One or more attributes used to filter the findings included in the insight
	Filters InsightAwsSecurityFindingFiltersOutput `pulumi:"filters"`
	// The grouping attribute for the insight's findings
	GroupByAttribute pulumi.StringOutput `pulumi:"groupByAttribute"`
	// The ARN of a Security Hub insight
	InsightArn pulumi.StringOutput `pulumi:"insightArn"`
	// The name of a Security Hub insight
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewInsight registers a new resource with the given unique name, arguments, and options.
func NewInsight(ctx *pulumi.Context,
	name string, args *InsightArgs, opts ...pulumi.ResourceOption) (*Insight, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filters == nil {
		return nil, errors.New("invalid value for required argument 'Filters'")
	}
	if args.GroupByAttribute == nil {
		return nil, errors.New("invalid value for required argument 'GroupByAttribute'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Insight
	err := ctx.RegisterResource("aws-native:securityhub:Insight", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInsight gets an existing Insight resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInsight(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InsightState, opts ...pulumi.ResourceOption) (*Insight, error) {
	var resource Insight
	err := ctx.ReadResource("aws-native:securityhub:Insight", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Insight resources.
type insightState struct {
}

type InsightState struct {
}

func (InsightState) ElementType() reflect.Type {
	return reflect.TypeOf((*insightState)(nil)).Elem()
}

type insightArgs struct {
	// One or more attributes used to filter the findings included in the insight
	Filters InsightAwsSecurityFindingFilters `pulumi:"filters"`
	// The grouping attribute for the insight's findings
	GroupByAttribute string `pulumi:"groupByAttribute"`
	// The name of a Security Hub insight
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Insight resource.
type InsightArgs struct {
	// One or more attributes used to filter the findings included in the insight
	Filters InsightAwsSecurityFindingFiltersInput
	// The grouping attribute for the insight's findings
	GroupByAttribute pulumi.StringInput
	// The name of a Security Hub insight
	Name pulumi.StringPtrInput
}

func (InsightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*insightArgs)(nil)).Elem()
}

type InsightInput interface {
	pulumi.Input

	ToInsightOutput() InsightOutput
	ToInsightOutputWithContext(ctx context.Context) InsightOutput
}

func (*Insight) ElementType() reflect.Type {
	return reflect.TypeOf((**Insight)(nil)).Elem()
}

func (i *Insight) ToInsightOutput() InsightOutput {
	return i.ToInsightOutputWithContext(context.Background())
}

func (i *Insight) ToInsightOutputWithContext(ctx context.Context) InsightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightOutput)
}

type InsightOutput struct{ *pulumi.OutputState }

func (InsightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Insight)(nil)).Elem()
}

func (o InsightOutput) ToInsightOutput() InsightOutput {
	return o
}

func (o InsightOutput) ToInsightOutputWithContext(ctx context.Context) InsightOutput {
	return o
}

// One or more attributes used to filter the findings included in the insight
func (o InsightOutput) Filters() InsightAwsSecurityFindingFiltersOutput {
	return o.ApplyT(func(v *Insight) InsightAwsSecurityFindingFiltersOutput { return v.Filters }).(InsightAwsSecurityFindingFiltersOutput)
}

// The grouping attribute for the insight's findings
func (o InsightOutput) GroupByAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *Insight) pulumi.StringOutput { return v.GroupByAttribute }).(pulumi.StringOutput)
}

// The ARN of a Security Hub insight
func (o InsightOutput) InsightArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Insight) pulumi.StringOutput { return v.InsightArn }).(pulumi.StringOutput)
}

// The name of a Security Hub insight
func (o InsightOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Insight) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InsightInput)(nil)).Elem(), &Insight{})
	pulumi.RegisterOutputType(InsightOutput{})
}
