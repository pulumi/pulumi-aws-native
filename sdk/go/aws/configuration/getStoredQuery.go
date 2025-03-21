// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::StoredQuery
func LookupStoredQuery(ctx *pulumi.Context, args *LookupStoredQueryArgs, opts ...pulumi.InvokeOption) (*LookupStoredQueryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStoredQueryResult
	err := ctx.Invoke("aws-native:configuration:getStoredQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStoredQueryArgs struct {
	// The name of the query.
	QueryName string `pulumi:"queryName"`
}

type LookupStoredQueryResult struct {
	// Amazon Resource Name (ARN) of the query. For example, arn:partition:service:region:account-id:resource-type/resource-name/resource-id.
	QueryArn *string `pulumi:"queryArn"`
	// A unique description for the query.
	QueryDescription *string `pulumi:"queryDescription"`
	// The expression of the query. For example, `SELECT resourceId, resourceType, supplementaryConfiguration.BucketVersioningConfiguration.status WHERE resourceType = 'AWS::S3::Bucket' AND supplementaryConfiguration.BucketVersioningConfiguration.status = 'Off'.`
	QueryExpression *string `pulumi:"queryExpression"`
	// The ID of the query.
	QueryId *string `pulumi:"queryId"`
	// The tags for the stored query.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupStoredQueryOutput(ctx *pulumi.Context, args LookupStoredQueryOutputArgs, opts ...pulumi.InvokeOption) LookupStoredQueryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStoredQueryResultOutput, error) {
			args := v.(LookupStoredQueryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:configuration:getStoredQuery", args, LookupStoredQueryResultOutput{}, options).(LookupStoredQueryResultOutput), nil
		}).(LookupStoredQueryResultOutput)
}

type LookupStoredQueryOutputArgs struct {
	// The name of the query.
	QueryName pulumi.StringInput `pulumi:"queryName"`
}

func (LookupStoredQueryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStoredQueryArgs)(nil)).Elem()
}

type LookupStoredQueryResultOutput struct{ *pulumi.OutputState }

func (LookupStoredQueryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStoredQueryResult)(nil)).Elem()
}

func (o LookupStoredQueryResultOutput) ToLookupStoredQueryResultOutput() LookupStoredQueryResultOutput {
	return o
}

func (o LookupStoredQueryResultOutput) ToLookupStoredQueryResultOutputWithContext(ctx context.Context) LookupStoredQueryResultOutput {
	return o
}

// Amazon Resource Name (ARN) of the query. For example, arn:partition:service:region:account-id:resource-type/resource-name/resource-id.
func (o LookupStoredQueryResultOutput) QueryArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoredQueryResult) *string { return v.QueryArn }).(pulumi.StringPtrOutput)
}

// A unique description for the query.
func (o LookupStoredQueryResultOutput) QueryDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoredQueryResult) *string { return v.QueryDescription }).(pulumi.StringPtrOutput)
}

// The expression of the query. For example, `SELECT resourceId, resourceType, supplementaryConfiguration.BucketVersioningConfiguration.status WHERE resourceType = 'AWS::S3::Bucket' AND supplementaryConfiguration.BucketVersioningConfiguration.status = 'Off'.`
func (o LookupStoredQueryResultOutput) QueryExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoredQueryResult) *string { return v.QueryExpression }).(pulumi.StringPtrOutput)
}

// The ID of the query.
func (o LookupStoredQueryResultOutput) QueryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoredQueryResult) *string { return v.QueryId }).(pulumi.StringPtrOutput)
}

// The tags for the stored query.
func (o LookupStoredQueryResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupStoredQueryResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStoredQueryResultOutput{})
}
