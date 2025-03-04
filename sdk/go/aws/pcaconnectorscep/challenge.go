// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pcaconnectorscep

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a SCEP Challenge that is used for certificate enrollment
type Challenge struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the challenge.
	ChallengeArn pulumi.StringOutput `pulumi:"challengeArn"`
	// The Amazon Resource Name (ARN) of the connector.
	ConnectorArn pulumi.StringOutput    `pulumi:"connectorArn"`
	Tags         pulumi.StringMapOutput `pulumi:"tags"`
}

// NewChallenge registers a new resource with the given unique name, arguments, and options.
func NewChallenge(ctx *pulumi.Context,
	name string, args *ChallengeArgs, opts ...pulumi.ResourceOption) (*Challenge, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectorArn == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"connectorArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Challenge
	err := ctx.RegisterResource("aws-native:pcaconnectorscep:Challenge", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChallenge gets an existing Challenge resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChallenge(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChallengeState, opts ...pulumi.ResourceOption) (*Challenge, error) {
	var resource Challenge
	err := ctx.ReadResource("aws-native:pcaconnectorscep:Challenge", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Challenge resources.
type challengeState struct {
}

type ChallengeState struct {
}

func (ChallengeState) ElementType() reflect.Type {
	return reflect.TypeOf((*challengeState)(nil)).Elem()
}

type challengeArgs struct {
	// The Amazon Resource Name (ARN) of the connector.
	ConnectorArn string            `pulumi:"connectorArn"`
	Tags         map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Challenge resource.
type ChallengeArgs struct {
	// The Amazon Resource Name (ARN) of the connector.
	ConnectorArn pulumi.StringInput
	Tags         pulumi.StringMapInput
}

func (ChallengeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*challengeArgs)(nil)).Elem()
}

type ChallengeInput interface {
	pulumi.Input

	ToChallengeOutput() ChallengeOutput
	ToChallengeOutputWithContext(ctx context.Context) ChallengeOutput
}

func (*Challenge) ElementType() reflect.Type {
	return reflect.TypeOf((**Challenge)(nil)).Elem()
}

func (i *Challenge) ToChallengeOutput() ChallengeOutput {
	return i.ToChallengeOutputWithContext(context.Background())
}

func (i *Challenge) ToChallengeOutputWithContext(ctx context.Context) ChallengeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChallengeOutput)
}

type ChallengeOutput struct{ *pulumi.OutputState }

func (ChallengeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Challenge)(nil)).Elem()
}

func (o ChallengeOutput) ToChallengeOutput() ChallengeOutput {
	return o
}

func (o ChallengeOutput) ToChallengeOutputWithContext(ctx context.Context) ChallengeOutput {
	return o
}

// The Amazon Resource Name (ARN) of the challenge.
func (o ChallengeOutput) ChallengeArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Challenge) pulumi.StringOutput { return v.ChallengeArn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the connector.
func (o ChallengeOutput) ConnectorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Challenge) pulumi.StringOutput { return v.ConnectorArn }).(pulumi.StringOutput)
}

func (o ChallengeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Challenge) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChallengeInput)(nil)).Elem(), &Challenge{})
	pulumi.RegisterOutputType(ChallengeOutput{})
}
