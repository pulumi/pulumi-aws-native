// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pcaconnectorscep

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a SCEP Challenge that is used for certificate enrollment
func LookupChallenge(ctx *pulumi.Context, args *LookupChallengeArgs, opts ...pulumi.InvokeOption) (*LookupChallengeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupChallengeResult
	err := ctx.Invoke("aws-native:pcaconnectorscep:getChallenge", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChallengeArgs struct {
	ChallengeArn string `pulumi:"challengeArn"`
}

type LookupChallengeResult struct {
	ChallengeArn *string           `pulumi:"challengeArn"`
	Tags         map[string]string `pulumi:"tags"`
}

func LookupChallengeOutput(ctx *pulumi.Context, args LookupChallengeOutputArgs, opts ...pulumi.InvokeOption) LookupChallengeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChallengeResult, error) {
			args := v.(LookupChallengeArgs)
			r, err := LookupChallenge(ctx, &args, opts...)
			var s LookupChallengeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChallengeResultOutput)
}

type LookupChallengeOutputArgs struct {
	ChallengeArn pulumi.StringInput `pulumi:"challengeArn"`
}

func (LookupChallengeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChallengeArgs)(nil)).Elem()
}

type LookupChallengeResultOutput struct{ *pulumi.OutputState }

func (LookupChallengeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChallengeResult)(nil)).Elem()
}

func (o LookupChallengeResultOutput) ToLookupChallengeResultOutput() LookupChallengeResultOutput {
	return o
}

func (o LookupChallengeResultOutput) ToLookupChallengeResultOutputWithContext(ctx context.Context) LookupChallengeResultOutput {
	return o
}

func (o LookupChallengeResultOutput) ChallengeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChallengeResult) *string { return v.ChallengeArn }).(pulumi.StringPtrOutput)
}

func (o LookupChallengeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupChallengeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChallengeResultOutput{})
}