// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nimblestudio

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::NimbleStudio::StudioComponent
func LookupStudioComponent(ctx *pulumi.Context, args *LookupStudioComponentArgs, opts ...pulumi.InvokeOption) (*LookupStudioComponentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStudioComponentResult
	err := ctx.Invoke("aws-native:nimblestudio:getStudioComponent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStudioComponentArgs struct {
	StudioComponentId string `pulumi:"studioComponentId"`
}

type LookupStudioComponentResult struct {
	Configuration         *StudioComponentConfiguration            `pulumi:"configuration"`
	Description           *string                                  `pulumi:"description"`
	Ec2SecurityGroupIds   []string                                 `pulumi:"ec2SecurityGroupIds"`
	InitializationScripts []StudioComponentInitializationScript    `pulumi:"initializationScripts"`
	Name                  *string                                  `pulumi:"name"`
	ScriptParameters      []StudioComponentScriptParameterKeyValue `pulumi:"scriptParameters"`
	StudioComponentId     *string                                  `pulumi:"studioComponentId"`
	Type                  *string                                  `pulumi:"type"`
}

func LookupStudioComponentOutput(ctx *pulumi.Context, args LookupStudioComponentOutputArgs, opts ...pulumi.InvokeOption) LookupStudioComponentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStudioComponentResultOutput, error) {
			args := v.(LookupStudioComponentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:nimblestudio:getStudioComponent", args, LookupStudioComponentResultOutput{}, options).(LookupStudioComponentResultOutput), nil
		}).(LookupStudioComponentResultOutput)
}

type LookupStudioComponentOutputArgs struct {
	StudioComponentId pulumi.StringInput `pulumi:"studioComponentId"`
}

func (LookupStudioComponentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioComponentArgs)(nil)).Elem()
}

type LookupStudioComponentResultOutput struct{ *pulumi.OutputState }

func (LookupStudioComponentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioComponentResult)(nil)).Elem()
}

func (o LookupStudioComponentResultOutput) ToLookupStudioComponentResultOutput() LookupStudioComponentResultOutput {
	return o
}

func (o LookupStudioComponentResultOutput) ToLookupStudioComponentResultOutputWithContext(ctx context.Context) LookupStudioComponentResultOutput {
	return o
}

func (o LookupStudioComponentResultOutput) Configuration() StudioComponentConfigurationPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *StudioComponentConfiguration { return v.Configuration }).(StudioComponentConfigurationPtrOutput)
}

func (o LookupStudioComponentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupStudioComponentResultOutput) Ec2SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) []string { return v.Ec2SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupStudioComponentResultOutput) InitializationScripts() StudioComponentInitializationScriptArrayOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) []StudioComponentInitializationScript {
		return v.InitializationScripts
	}).(StudioComponentInitializationScriptArrayOutput)
}

func (o LookupStudioComponentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupStudioComponentResultOutput) ScriptParameters() StudioComponentScriptParameterKeyValueArrayOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) []StudioComponentScriptParameterKeyValue {
		return v.ScriptParameters
	}).(StudioComponentScriptParameterKeyValueArrayOutput)
}

func (o LookupStudioComponentResultOutput) StudioComponentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *string { return v.StudioComponentId }).(pulumi.StringPtrOutput)
}

func (o LookupStudioComponentResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStudioComponentResultOutput{})
}
