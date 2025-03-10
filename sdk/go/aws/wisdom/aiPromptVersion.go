// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wisdom

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Wisdom::AIPromptVersion Resource Type
type AiPromptVersion struct {
	pulumi.CustomResourceState

	// The ARN of the AI prompt.
	AiPromptArn pulumi.StringOutput `pulumi:"aiPromptArn"`
	// The identifier of the Amazon Q in Connect AI prompt.
	AiPromptId        pulumi.StringOutput `pulumi:"aiPromptId"`
	AiPromptVersionId pulumi.StringOutput `pulumi:"aiPromptVersionId"`
	AssistantArn      pulumi.StringOutput `pulumi:"assistantArn"`
	// The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the ARN.
	AssistantId pulumi.StringOutput `pulumi:"assistantId"`
	// The time the AI Prompt version was last modified in seconds.
	ModifiedTimeSeconds pulumi.Float64PtrOutput `pulumi:"modifiedTimeSeconds"`
	// The version number for this AI Prompt version.
	VersionNumber pulumi.Float64Output `pulumi:"versionNumber"`
}

// NewAiPromptVersion registers a new resource with the given unique name, arguments, and options.
func NewAiPromptVersion(ctx *pulumi.Context,
	name string, args *AiPromptVersionArgs, opts ...pulumi.ResourceOption) (*AiPromptVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AiPromptId == nil {
		return nil, errors.New("invalid value for required argument 'AiPromptId'")
	}
	if args.AssistantId == nil {
		return nil, errors.New("invalid value for required argument 'AssistantId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"aiPromptId",
		"assistantId",
		"modifiedTimeSeconds",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AiPromptVersion
	err := ctx.RegisterResource("aws-native:wisdom:AiPromptVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiPromptVersion gets an existing AiPromptVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiPromptVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiPromptVersionState, opts ...pulumi.ResourceOption) (*AiPromptVersion, error) {
	var resource AiPromptVersion
	err := ctx.ReadResource("aws-native:wisdom:AiPromptVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiPromptVersion resources.
type aiPromptVersionState struct {
}

type AiPromptVersionState struct {
}

func (AiPromptVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiPromptVersionState)(nil)).Elem()
}

type aiPromptVersionArgs struct {
	// The identifier of the Amazon Q in Connect AI prompt.
	AiPromptId string `pulumi:"aiPromptId"`
	// The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the ARN.
	AssistantId string `pulumi:"assistantId"`
	// The time the AI Prompt version was last modified in seconds.
	ModifiedTimeSeconds *float64 `pulumi:"modifiedTimeSeconds"`
}

// The set of arguments for constructing a AiPromptVersion resource.
type AiPromptVersionArgs struct {
	// The identifier of the Amazon Q in Connect AI prompt.
	AiPromptId pulumi.StringInput
	// The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the ARN.
	AssistantId pulumi.StringInput
	// The time the AI Prompt version was last modified in seconds.
	ModifiedTimeSeconds pulumi.Float64PtrInput
}

func (AiPromptVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiPromptVersionArgs)(nil)).Elem()
}

type AiPromptVersionInput interface {
	pulumi.Input

	ToAiPromptVersionOutput() AiPromptVersionOutput
	ToAiPromptVersionOutputWithContext(ctx context.Context) AiPromptVersionOutput
}

func (*AiPromptVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**AiPromptVersion)(nil)).Elem()
}

func (i *AiPromptVersion) ToAiPromptVersionOutput() AiPromptVersionOutput {
	return i.ToAiPromptVersionOutputWithContext(context.Background())
}

func (i *AiPromptVersion) ToAiPromptVersionOutputWithContext(ctx context.Context) AiPromptVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiPromptVersionOutput)
}

type AiPromptVersionOutput struct{ *pulumi.OutputState }

func (AiPromptVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiPromptVersion)(nil)).Elem()
}

func (o AiPromptVersionOutput) ToAiPromptVersionOutput() AiPromptVersionOutput {
	return o
}

func (o AiPromptVersionOutput) ToAiPromptVersionOutputWithContext(ctx context.Context) AiPromptVersionOutput {
	return o
}

// The ARN of the AI prompt.
func (o AiPromptVersionOutput) AiPromptArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AiPromptVersion) pulumi.StringOutput { return v.AiPromptArn }).(pulumi.StringOutput)
}

// The identifier of the Amazon Q in Connect AI prompt.
func (o AiPromptVersionOutput) AiPromptId() pulumi.StringOutput {
	return o.ApplyT(func(v *AiPromptVersion) pulumi.StringOutput { return v.AiPromptId }).(pulumi.StringOutput)
}

func (o AiPromptVersionOutput) AiPromptVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AiPromptVersion) pulumi.StringOutput { return v.AiPromptVersionId }).(pulumi.StringOutput)
}

func (o AiPromptVersionOutput) AssistantArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AiPromptVersion) pulumi.StringOutput { return v.AssistantArn }).(pulumi.StringOutput)
}

// The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the ARN.
func (o AiPromptVersionOutput) AssistantId() pulumi.StringOutput {
	return o.ApplyT(func(v *AiPromptVersion) pulumi.StringOutput { return v.AssistantId }).(pulumi.StringOutput)
}

// The time the AI Prompt version was last modified in seconds.
func (o AiPromptVersionOutput) ModifiedTimeSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AiPromptVersion) pulumi.Float64PtrOutput { return v.ModifiedTimeSeconds }).(pulumi.Float64PtrOutput)
}

// The version number for this AI Prompt version.
func (o AiPromptVersionOutput) VersionNumber() pulumi.Float64Output {
	return o.ApplyT(func(v *AiPromptVersion) pulumi.Float64Output { return v.VersionNumber }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AiPromptVersionInput)(nil)).Elem(), &AiPromptVersion{})
	pulumi.RegisterOutputType(AiPromptVersionOutput{})
}
