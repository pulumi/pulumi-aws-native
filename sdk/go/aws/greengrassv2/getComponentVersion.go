// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrassv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for Greengrass component version.
func LookupComponentVersion(ctx *pulumi.Context, args *LookupComponentVersionArgs, opts ...pulumi.InvokeOption) (*LookupComponentVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupComponentVersionResult
	err := ctx.Invoke("aws-native:greengrassv2:getComponentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComponentVersionArgs struct {
	// The ARN of the component version.
	Arn string `pulumi:"arn"`
}

type LookupComponentVersionResult struct {
	// The ARN of the component version.
	Arn *string `pulumi:"arn"`
	// The name of the component.
	ComponentName *string `pulumi:"componentName"`
	// The version of the component.
	ComponentVersion *string `pulumi:"componentVersion"`
	// Application-specific metadata to attach to the component version. You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	Tags map[string]string `pulumi:"tags"`
}

func LookupComponentVersionOutput(ctx *pulumi.Context, args LookupComponentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupComponentVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComponentVersionResultOutput, error) {
			args := v.(LookupComponentVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:greengrassv2:getComponentVersion", args, LookupComponentVersionResultOutput{}, options).(LookupComponentVersionResultOutput), nil
		}).(LookupComponentVersionResultOutput)
}

type LookupComponentVersionOutputArgs struct {
	// The ARN of the component version.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupComponentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentVersionArgs)(nil)).Elem()
}

type LookupComponentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupComponentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentVersionResult)(nil)).Elem()
}

func (o LookupComponentVersionResultOutput) ToLookupComponentVersionResultOutput() LookupComponentVersionResultOutput {
	return o
}

func (o LookupComponentVersionResultOutput) ToLookupComponentVersionResultOutputWithContext(ctx context.Context) LookupComponentVersionResultOutput {
	return o
}

// The ARN of the component version.
func (o LookupComponentVersionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The name of the component.
func (o LookupComponentVersionResultOutput) ComponentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) *string { return v.ComponentName }).(pulumi.StringPtrOutput)
}

// The version of the component.
func (o LookupComponentVersionResultOutput) ComponentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) *string { return v.ComponentVersion }).(pulumi.StringPtrOutput)
}

// Application-specific metadata to attach to the component version. You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
//
// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
func (o LookupComponentVersionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComponentVersionResultOutput{})
}
