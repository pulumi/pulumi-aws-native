// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::LoggerDefinitionVersion
func LookupLoggerDefinitionVersion(ctx *pulumi.Context, args *LookupLoggerDefinitionVersionArgs, opts ...pulumi.InvokeOption) (*LookupLoggerDefinitionVersionResult, error) {
	var rv LookupLoggerDefinitionVersionResult
	err := ctx.Invoke("aws-native:greengrass:getLoggerDefinitionVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoggerDefinitionVersionArgs struct {
	Id string `pulumi:"id"`
}

type LookupLoggerDefinitionVersionResult struct {
	Id *string `pulumi:"id"`
}

func LookupLoggerDefinitionVersionOutput(ctx *pulumi.Context, args LookupLoggerDefinitionVersionOutputArgs, opts ...pulumi.InvokeOption) LookupLoggerDefinitionVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoggerDefinitionVersionResult, error) {
			args := v.(LookupLoggerDefinitionVersionArgs)
			r, err := LookupLoggerDefinitionVersion(ctx, &args, opts...)
			var s LookupLoggerDefinitionVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoggerDefinitionVersionResultOutput)
}

type LookupLoggerDefinitionVersionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLoggerDefinitionVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggerDefinitionVersionArgs)(nil)).Elem()
}

type LookupLoggerDefinitionVersionResultOutput struct{ *pulumi.OutputState }

func (LookupLoggerDefinitionVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggerDefinitionVersionResult)(nil)).Elem()
}

func (o LookupLoggerDefinitionVersionResultOutput) ToLookupLoggerDefinitionVersionResultOutput() LookupLoggerDefinitionVersionResultOutput {
	return o
}

func (o LookupLoggerDefinitionVersionResultOutput) ToLookupLoggerDefinitionVersionResultOutputWithContext(ctx context.Context) LookupLoggerDefinitionVersionResultOutput {
	return o
}

func (o LookupLoggerDefinitionVersionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoggerDefinitionVersionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoggerDefinitionVersionResultOutput{})
}