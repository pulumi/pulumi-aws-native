// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Amplify::Branch resource creates a new branch within an app.
func LookupBranch(ctx *pulumi.Context, args *LookupBranchArgs, opts ...pulumi.InvokeOption) (*LookupBranchResult, error) {
	var rv LookupBranchResult
	err := ctx.Invoke("aws-native:amplify:getBranch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBranchArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupBranchResult struct {
	Arn                        *string                     `pulumi:"arn"`
	BuildSpec                  *string                     `pulumi:"buildSpec"`
	Description                *string                     `pulumi:"description"`
	EnableAutoBuild            *bool                       `pulumi:"enableAutoBuild"`
	EnablePerformanceMode      *bool                       `pulumi:"enablePerformanceMode"`
	EnablePullRequestPreview   *bool                       `pulumi:"enablePullRequestPreview"`
	EnvironmentVariables       []BranchEnvironmentVariable `pulumi:"environmentVariables"`
	PullRequestEnvironmentName *string                     `pulumi:"pullRequestEnvironmentName"`
	Stage                      *BranchStage                `pulumi:"stage"`
	Tags                       []BranchTag                 `pulumi:"tags"`
}

func LookupBranchOutput(ctx *pulumi.Context, args LookupBranchOutputArgs, opts ...pulumi.InvokeOption) LookupBranchResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBranchResult, error) {
			args := v.(LookupBranchArgs)
			r, err := LookupBranch(ctx, &args, opts...)
			var s LookupBranchResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBranchResultOutput)
}

type LookupBranchOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupBranchOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBranchArgs)(nil)).Elem()
}

type LookupBranchResultOutput struct{ *pulumi.OutputState }

func (LookupBranchResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBranchResult)(nil)).Elem()
}

func (o LookupBranchResultOutput) ToLookupBranchResultOutput() LookupBranchResultOutput {
	return o
}

func (o LookupBranchResultOutput) ToLookupBranchResultOutputWithContext(ctx context.Context) LookupBranchResultOutput {
	return o
}

func (o LookupBranchResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBranchResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupBranchResultOutput) BuildSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBranchResult) *string { return v.BuildSpec }).(pulumi.StringPtrOutput)
}

func (o LookupBranchResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBranchResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupBranchResultOutput) EnableAutoBuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBranchResult) *bool { return v.EnableAutoBuild }).(pulumi.BoolPtrOutput)
}

func (o LookupBranchResultOutput) EnablePerformanceMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBranchResult) *bool { return v.EnablePerformanceMode }).(pulumi.BoolPtrOutput)
}

func (o LookupBranchResultOutput) EnablePullRequestPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBranchResult) *bool { return v.EnablePullRequestPreview }).(pulumi.BoolPtrOutput)
}

func (o LookupBranchResultOutput) EnvironmentVariables() BranchEnvironmentVariableArrayOutput {
	return o.ApplyT(func(v LookupBranchResult) []BranchEnvironmentVariable { return v.EnvironmentVariables }).(BranchEnvironmentVariableArrayOutput)
}

func (o LookupBranchResultOutput) PullRequestEnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBranchResult) *string { return v.PullRequestEnvironmentName }).(pulumi.StringPtrOutput)
}

func (o LookupBranchResultOutput) Stage() BranchStagePtrOutput {
	return o.ApplyT(func(v LookupBranchResult) *BranchStage { return v.Stage }).(BranchStagePtrOutput)
}

func (o LookupBranchResultOutput) Tags() BranchTagArrayOutput {
	return o.ApplyT(func(v LookupBranchResult) []BranchTag { return v.Tags }).(BranchTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBranchResultOutput{})
}