// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource represents an individual schema version of a schema defined in Glue Schema Registry.
func LookupSchemaVersion(ctx *pulumi.Context, args *LookupSchemaVersionArgs, opts ...pulumi.InvokeOption) (*LookupSchemaVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSchemaVersionResult
	err := ctx.Invoke("aws-native:glue:getSchemaVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchemaVersionArgs struct {
	// Represents the version ID associated with the schema version.
	VersionId string `pulumi:"versionId"`
}

type LookupSchemaVersionResult struct {
	// Represents the version ID associated with the schema version.
	VersionId *string `pulumi:"versionId"`
}

func LookupSchemaVersionOutput(ctx *pulumi.Context, args LookupSchemaVersionOutputArgs, opts ...pulumi.InvokeOption) LookupSchemaVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSchemaVersionResultOutput, error) {
			args := v.(LookupSchemaVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:glue:getSchemaVersion", args, LookupSchemaVersionResultOutput{}, options).(LookupSchemaVersionResultOutput), nil
		}).(LookupSchemaVersionResultOutput)
}

type LookupSchemaVersionOutputArgs struct {
	// Represents the version ID associated with the schema version.
	VersionId pulumi.StringInput `pulumi:"versionId"`
}

func (LookupSchemaVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaVersionArgs)(nil)).Elem()
}

type LookupSchemaVersionResultOutput struct{ *pulumi.OutputState }

func (LookupSchemaVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaVersionResult)(nil)).Elem()
}

func (o LookupSchemaVersionResultOutput) ToLookupSchemaVersionResultOutput() LookupSchemaVersionResultOutput {
	return o
}

func (o LookupSchemaVersionResultOutput) ToLookupSchemaVersionResultOutputWithContext(ctx context.Context) LookupSchemaVersionResultOutput {
	return o
}

// Represents the version ID associated with the schema version.
func (o LookupSchemaVersionResultOutput) VersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) *string { return v.VersionId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSchemaVersionResultOutput{})
}
