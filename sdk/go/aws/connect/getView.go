// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Connect::View
func LookupView(ctx *pulumi.Context, args *LookupViewArgs, opts ...pulumi.InvokeOption) (*LookupViewResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupViewResult
	err := ctx.Invoke("aws-native:connect:getView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewArgs struct {
	// The Amazon Resource Name (ARN) of the view.
	ViewArn string `pulumi:"viewArn"`
}

type LookupViewResult struct {
	// The actions of the view in an array.
	Actions []string `pulumi:"actions"`
	// The description of the view.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the instance.
	InstanceArn *string `pulumi:"instanceArn"`
	// The name of the view.
	Name *string `pulumi:"name"`
	// One or more tags.
	Tags []ViewTag `pulumi:"tags"`
	// The template of the view as JSON.
	Template interface{} `pulumi:"template"`
	// The Amazon Resource Name (ARN) of the view.
	ViewArn *string `pulumi:"viewArn"`
	// The view content hash.
	ViewContentSha256 *string `pulumi:"viewContentSha256"`
	// The view id of the view.
	ViewId *string `pulumi:"viewId"`
}

func LookupViewOutput(ctx *pulumi.Context, args LookupViewOutputArgs, opts ...pulumi.InvokeOption) LookupViewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupViewResult, error) {
			args := v.(LookupViewArgs)
			r, err := LookupView(ctx, &args, opts...)
			var s LookupViewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupViewResultOutput)
}

type LookupViewOutputArgs struct {
	// The Amazon Resource Name (ARN) of the view.
	ViewArn pulumi.StringInput `pulumi:"viewArn"`
}

func (LookupViewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewArgs)(nil)).Elem()
}

type LookupViewResultOutput struct{ *pulumi.OutputState }

func (LookupViewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewResult)(nil)).Elem()
}

func (o LookupViewResultOutput) ToLookupViewResultOutput() LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) ToLookupViewResultOutputWithContext(ctx context.Context) LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupViewResult] {
	return pulumix.Output[LookupViewResult]{
		OutputState: o.OutputState,
	}
}

// The actions of the view in an array.
func (o LookupViewResultOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupViewResult) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

// The description of the view.
func (o LookupViewResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the instance.
func (o LookupViewResultOutput) InstanceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.InstanceArn }).(pulumi.StringPtrOutput)
}

// The name of the view.
func (o LookupViewResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// One or more tags.
func (o LookupViewResultOutput) Tags() ViewTagArrayOutput {
	return o.ApplyT(func(v LookupViewResult) []ViewTag { return v.Tags }).(ViewTagArrayOutput)
}

// The template of the view as JSON.
func (o LookupViewResultOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupViewResult) interface{} { return v.Template }).(pulumi.AnyOutput)
}

// The Amazon Resource Name (ARN) of the view.
func (o LookupViewResultOutput) ViewArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.ViewArn }).(pulumi.StringPtrOutput)
}

// The view content hash.
func (o LookupViewResultOutput) ViewContentSha256() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.ViewContentSha256 }).(pulumi.StringPtrOutput)
}

// The view id of the view.
func (o LookupViewResultOutput) ViewId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.ViewId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupViewResultOutput{})
}