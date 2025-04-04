// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A zero-ETL integration with Amazon Redshift.
func LookupIntegration(ctx *pulumi.Context, args *LookupIntegrationArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIntegrationResult
	err := ctx.Invoke("aws-native:rds:getIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationArgs struct {
	// The ARN of the integration.
	IntegrationArn string `pulumi:"integrationArn"`
}

type LookupIntegrationResult struct {
	// The time when the integration was created, in Universal Coordinated Time (UTC).
	CreateTime *string `pulumi:"createTime"`
	// Data filters for the integration. These filters determine which tables from the source database are sent to the target Amazon Redshift data warehouse.
	DataFilter *string `pulumi:"dataFilter"`
	// A description of the integration.
	Description *string `pulumi:"description"`
	// The ARN of the integration.
	IntegrationArn *string `pulumi:"integrationArn"`
	// The name of the integration.
	IntegrationName *string `pulumi:"integrationName"`
	// A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupIntegrationOutput(ctx *pulumi.Context, args LookupIntegrationOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIntegrationResultOutput, error) {
			args := v.(LookupIntegrationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:rds:getIntegration", args, LookupIntegrationResultOutput{}, options).(LookupIntegrationResultOutput), nil
		}).(LookupIntegrationResultOutput)
}

type LookupIntegrationOutputArgs struct {
	// The ARN of the integration.
	IntegrationArn pulumi.StringInput `pulumi:"integrationArn"`
}

func (LookupIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationArgs)(nil)).Elem()
}

type LookupIntegrationResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationResult)(nil)).Elem()
}

func (o LookupIntegrationResultOutput) ToLookupIntegrationResultOutput() LookupIntegrationResultOutput {
	return o
}

func (o LookupIntegrationResultOutput) ToLookupIntegrationResultOutputWithContext(ctx context.Context) LookupIntegrationResultOutput {
	return o
}

// The time when the integration was created, in Universal Coordinated Time (UTC).
func (o LookupIntegrationResultOutput) CreateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.CreateTime }).(pulumi.StringPtrOutput)
}

// Data filters for the integration. These filters determine which tables from the source database are sent to the target Amazon Redshift data warehouse.
func (o LookupIntegrationResultOutput) DataFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.DataFilter }).(pulumi.StringPtrOutput)
}

// A description of the integration.
func (o LookupIntegrationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The ARN of the integration.
func (o LookupIntegrationResultOutput) IntegrationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.IntegrationArn }).(pulumi.StringPtrOutput)
}

// The name of the integration.
func (o LookupIntegrationResultOutput) IntegrationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.IntegrationName }).(pulumi.StringPtrOutput)
}

// A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*.
func (o LookupIntegrationResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupIntegrationResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationResultOutput{})
}
