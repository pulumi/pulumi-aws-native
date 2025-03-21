// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticBeanstalk::ApplicationVersion
func LookupApplicationVersion(ctx *pulumi.Context, args *LookupApplicationVersionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationVersionResult
	err := ctx.Invoke("aws-native:elasticbeanstalk:getApplicationVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationVersionArgs struct {
	// The name of the Elastic Beanstalk application that is associated with this application version.
	ApplicationName string `pulumi:"applicationName"`
	Id              string `pulumi:"id"`
}

type LookupApplicationVersionResult struct {
	// A description of this application version.
	Description *string `pulumi:"description"`
	Id          *string `pulumi:"id"`
}

func LookupApplicationVersionOutput(ctx *pulumi.Context, args LookupApplicationVersionOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupApplicationVersionResultOutput, error) {
			args := v.(LookupApplicationVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:elasticbeanstalk:getApplicationVersion", args, LookupApplicationVersionResultOutput{}, options).(LookupApplicationVersionResultOutput), nil
		}).(LookupApplicationVersionResultOutput)
}

type LookupApplicationVersionOutputArgs struct {
	// The name of the Elastic Beanstalk application that is associated with this application version.
	ApplicationName pulumi.StringInput `pulumi:"applicationName"`
	Id              pulumi.StringInput `pulumi:"id"`
}

func (LookupApplicationVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationVersionArgs)(nil)).Elem()
}

type LookupApplicationVersionResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationVersionResult)(nil)).Elem()
}

func (o LookupApplicationVersionResultOutput) ToLookupApplicationVersionResultOutput() LookupApplicationVersionResultOutput {
	return o
}

func (o LookupApplicationVersionResultOutput) ToLookupApplicationVersionResultOutputWithContext(ctx context.Context) LookupApplicationVersionResultOutput {
	return o
}

// A description of this application version.
func (o LookupApplicationVersionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationVersionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationVersionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationVersionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationVersionResultOutput{})
}
