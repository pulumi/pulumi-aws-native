// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains a list of Regular expressions based on the provided inputs. RegexPatternSet can be used with other WAF entities with RegexPatternSetReferenceStatement to perform other actions .
type RegexPatternSet struct {
	pulumi.CustomResourceState

	// ARN of the WAF entity.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the entity.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the RegexPatternSet.
	Name                  pulumi.StringPtrOutput   `pulumi:"name"`
	RegularExpressionList pulumi.StringArrayOutput `pulumi:"regularExpressionList"`
	// Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.
	Scope RegexPatternSetScopeOutput    `pulumi:"scope"`
	Tags  RegexPatternSetTagArrayOutput `pulumi:"tags"`
}

// NewRegexPatternSet registers a new resource with the given unique name, arguments, and options.
func NewRegexPatternSet(ctx *pulumi.Context,
	name string, args *RegexPatternSetArgs, opts ...pulumi.ResourceOption) (*RegexPatternSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegularExpressionList == nil {
		return nil, errors.New("invalid value for required argument 'RegularExpressionList'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource RegexPatternSet
	err := ctx.RegisterResource("aws-native:wafv2:RegexPatternSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegexPatternSet gets an existing RegexPatternSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexPatternSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegexPatternSetState, opts ...pulumi.ResourceOption) (*RegexPatternSet, error) {
	var resource RegexPatternSet
	err := ctx.ReadResource("aws-native:wafv2:RegexPatternSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegexPatternSet resources.
type regexPatternSetState struct {
}

type RegexPatternSetState struct {
}

func (RegexPatternSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*regexPatternSetState)(nil)).Elem()
}

type regexPatternSetArgs struct {
	// Description of the entity.
	Description *string `pulumi:"description"`
	// Name of the RegexPatternSet.
	Name                  *string  `pulumi:"name"`
	RegularExpressionList []string `pulumi:"regularExpressionList"`
	// Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.
	Scope RegexPatternSetScope `pulumi:"scope"`
	Tags  []RegexPatternSetTag `pulumi:"tags"`
}

// The set of arguments for constructing a RegexPatternSet resource.
type RegexPatternSetArgs struct {
	// Description of the entity.
	Description pulumi.StringPtrInput
	// Name of the RegexPatternSet.
	Name                  pulumi.StringPtrInput
	RegularExpressionList pulumi.StringArrayInput
	// Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.
	Scope RegexPatternSetScopeInput
	Tags  RegexPatternSetTagArrayInput
}

func (RegexPatternSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regexPatternSetArgs)(nil)).Elem()
}

type RegexPatternSetInput interface {
	pulumi.Input

	ToRegexPatternSetOutput() RegexPatternSetOutput
	ToRegexPatternSetOutputWithContext(ctx context.Context) RegexPatternSetOutput
}

func (*RegexPatternSet) ElementType() reflect.Type {
	return reflect.TypeOf((**RegexPatternSet)(nil)).Elem()
}

func (i *RegexPatternSet) ToRegexPatternSetOutput() RegexPatternSetOutput {
	return i.ToRegexPatternSetOutputWithContext(context.Background())
}

func (i *RegexPatternSet) ToRegexPatternSetOutputWithContext(ctx context.Context) RegexPatternSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexPatternSetOutput)
}

type RegexPatternSetOutput struct{ *pulumi.OutputState }

func (RegexPatternSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegexPatternSet)(nil)).Elem()
}

func (o RegexPatternSetOutput) ToRegexPatternSetOutput() RegexPatternSetOutput {
	return o
}

func (o RegexPatternSetOutput) ToRegexPatternSetOutputWithContext(ctx context.Context) RegexPatternSetOutput {
	return o
}

// ARN of the WAF entity.
func (o RegexPatternSetOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *RegexPatternSet) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of the entity.
func (o RegexPatternSetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegexPatternSet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the RegexPatternSet.
func (o RegexPatternSetOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegexPatternSet) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RegexPatternSetOutput) RegularExpressionList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RegexPatternSet) pulumi.StringArrayOutput { return v.RegularExpressionList }).(pulumi.StringArrayOutput)
}

// Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.
func (o RegexPatternSetOutput) Scope() RegexPatternSetScopeOutput {
	return o.ApplyT(func(v *RegexPatternSet) RegexPatternSetScopeOutput { return v.Scope }).(RegexPatternSetScopeOutput)
}

func (o RegexPatternSetOutput) Tags() RegexPatternSetTagArrayOutput {
	return o.ApplyT(func(v *RegexPatternSet) RegexPatternSetTagArrayOutput { return v.Tags }).(RegexPatternSetTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegexPatternSetInput)(nil)).Elem(), &RegexPatternSet{})
	pulumi.RegisterOutputType(RegexPatternSetOutput{})
}