// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotevents

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::IoTEvents::Input resource creates an input. To monitor your devices and processes, they must have a way to get telemetry data into AWS IoT Events. This is done by sending messages as *inputs* to AWS IoT Events. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide*.
func LookupInput(ctx *pulumi.Context, args *LookupInputArgs, opts ...pulumi.InvokeOption) (*LookupInputResult, error) {
	var rv LookupInputResult
	err := ctx.Invoke("aws-native:iotevents:getInput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInputArgs struct {
	// The name of the input.
	InputName string `pulumi:"inputName"`
}

type LookupInputResult struct {
	InputDefinition *InputDefinition `pulumi:"inputDefinition"`
	// A brief description of the input.
	InputDescription *string `pulumi:"inputDescription"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
	Tags []InputTag `pulumi:"tags"`
}

func LookupInputOutput(ctx *pulumi.Context, args LookupInputOutputArgs, opts ...pulumi.InvokeOption) LookupInputResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInputResult, error) {
			args := v.(LookupInputArgs)
			r, err := LookupInput(ctx, &args, opts...)
			var s LookupInputResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInputResultOutput)
}

type LookupInputOutputArgs struct {
	// The name of the input.
	InputName pulumi.StringInput `pulumi:"inputName"`
}

func (LookupInputOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInputArgs)(nil)).Elem()
}

type LookupInputResultOutput struct{ *pulumi.OutputState }

func (LookupInputResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInputResult)(nil)).Elem()
}

func (o LookupInputResultOutput) ToLookupInputResultOutput() LookupInputResultOutput {
	return o
}

func (o LookupInputResultOutput) ToLookupInputResultOutputWithContext(ctx context.Context) LookupInputResultOutput {
	return o
}

func (o LookupInputResultOutput) InputDefinition() InputDefinitionPtrOutput {
	return o.ApplyT(func(v LookupInputResult) *InputDefinition { return v.InputDefinition }).(InputDefinitionPtrOutput)
}

// A brief description of the input.
func (o LookupInputResultOutput) InputDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInputResult) *string { return v.InputDescription }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
//
// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
func (o LookupInputResultOutput) Tags() InputTagArrayOutput {
	return o.ApplyT(func(v LookupInputResult) []InputTag { return v.Tags }).(InputTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInputResultOutput{})
}