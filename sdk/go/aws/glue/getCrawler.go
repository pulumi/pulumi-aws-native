// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::Crawler
func LookupCrawler(ctx *pulumi.Context, args *LookupCrawlerArgs, opts ...pulumi.InvokeOption) (*LookupCrawlerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCrawlerResult
	err := ctx.Invoke("aws-native:glue:getCrawler", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCrawlerArgs struct {
	// The name of the crawler.
	Name string `pulumi:"name"`
}

type LookupCrawlerResult struct {
	// A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler.
	Classifiers []string `pulumi:"classifiers"`
	// Crawler configuration information. This versioned JSON string allows users to specify aspects of a crawler's behavior.
	Configuration *string `pulumi:"configuration"`
	// The name of the SecurityConfiguration structure to be used by this crawler.
	CrawlerSecurityConfiguration *string `pulumi:"crawlerSecurityConfiguration"`
	// The name of the database in which the crawler's output is stored.
	DatabaseName *string `pulumi:"databaseName"`
	// A description of the crawler.
	Description *string `pulumi:"description"`
	// Specifies whether the crawler should use AWS Lake Formation credentials for the crawler instead of the IAM role credentials.
	LakeFormationConfiguration *CrawlerLakeFormationConfiguration `pulumi:"lakeFormationConfiguration"`
	// A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.
	RecrawlPolicy *CrawlerRecrawlPolicy `pulumi:"recrawlPolicy"`
	// The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.
	Role *string `pulumi:"role"`
	// For scheduled crawlers, the schedule when the crawler runs.
	Schedule *CrawlerSchedule `pulumi:"schedule"`
	// The policy that specifies update and delete behaviors for the crawler. The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The `SchemaChangePolicy` does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the `SchemaChangePolicy` on a crawler.
	//
	// The SchemaChangePolicy consists of two components, `UpdateBehavior` and `DeleteBehavior` .
	SchemaChangePolicy *CrawlerSchemaChangePolicy `pulumi:"schemaChangePolicy"`
	// The prefix added to the names of tables that are created.
	TablePrefix *string `pulumi:"tablePrefix"`
	// The tags to use with this crawler.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Crawler` for more information about the expected schema for this property.
	Tags interface{} `pulumi:"tags"`
	// A collection of targets to crawl.
	Targets *CrawlerTargets `pulumi:"targets"`
}

func LookupCrawlerOutput(ctx *pulumi.Context, args LookupCrawlerOutputArgs, opts ...pulumi.InvokeOption) LookupCrawlerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCrawlerResultOutput, error) {
			args := v.(LookupCrawlerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:glue:getCrawler", args, LookupCrawlerResultOutput{}, options).(LookupCrawlerResultOutput), nil
		}).(LookupCrawlerResultOutput)
}

type LookupCrawlerOutputArgs struct {
	// The name of the crawler.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupCrawlerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCrawlerArgs)(nil)).Elem()
}

type LookupCrawlerResultOutput struct{ *pulumi.OutputState }

func (LookupCrawlerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCrawlerResult)(nil)).Elem()
}

func (o LookupCrawlerResultOutput) ToLookupCrawlerResultOutput() LookupCrawlerResultOutput {
	return o
}

func (o LookupCrawlerResultOutput) ToLookupCrawlerResultOutputWithContext(ctx context.Context) LookupCrawlerResultOutput {
	return o
}

// A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler.
func (o LookupCrawlerResultOutput) Classifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCrawlerResult) []string { return v.Classifiers }).(pulumi.StringArrayOutput)
}

// Crawler configuration information. This versioned JSON string allows users to specify aspects of a crawler's behavior.
func (o LookupCrawlerResultOutput) Configuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *string { return v.Configuration }).(pulumi.StringPtrOutput)
}

// The name of the SecurityConfiguration structure to be used by this crawler.
func (o LookupCrawlerResultOutput) CrawlerSecurityConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *string { return v.CrawlerSecurityConfiguration }).(pulumi.StringPtrOutput)
}

// The name of the database in which the crawler's output is stored.
func (o LookupCrawlerResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// A description of the crawler.
func (o LookupCrawlerResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether the crawler should use AWS Lake Formation credentials for the crawler instead of the IAM role credentials.
func (o LookupCrawlerResultOutput) LakeFormationConfiguration() CrawlerLakeFormationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *CrawlerLakeFormationConfiguration { return v.LakeFormationConfiguration }).(CrawlerLakeFormationConfigurationPtrOutput)
}

// A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.
func (o LookupCrawlerResultOutput) RecrawlPolicy() CrawlerRecrawlPolicyPtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *CrawlerRecrawlPolicy { return v.RecrawlPolicy }).(CrawlerRecrawlPolicyPtrOutput)
}

// The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.
func (o LookupCrawlerResultOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *string { return v.Role }).(pulumi.StringPtrOutput)
}

// For scheduled crawlers, the schedule when the crawler runs.
func (o LookupCrawlerResultOutput) Schedule() CrawlerSchedulePtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *CrawlerSchedule { return v.Schedule }).(CrawlerSchedulePtrOutput)
}

// The policy that specifies update and delete behaviors for the crawler. The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The `SchemaChangePolicy` does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the `SchemaChangePolicy` on a crawler.
//
// The SchemaChangePolicy consists of two components, `UpdateBehavior` and `DeleteBehavior` .
func (o LookupCrawlerResultOutput) SchemaChangePolicy() CrawlerSchemaChangePolicyPtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *CrawlerSchemaChangePolicy { return v.SchemaChangePolicy }).(CrawlerSchemaChangePolicyPtrOutput)
}

// The prefix added to the names of tables that are created.
func (o LookupCrawlerResultOutput) TablePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *string { return v.TablePrefix }).(pulumi.StringPtrOutput)
}

// The tags to use with this crawler.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Crawler` for more information about the expected schema for this property.
func (o LookupCrawlerResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupCrawlerResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

// A collection of targets to crawl.
func (o LookupCrawlerResultOutput) Targets() CrawlerTargetsPtrOutput {
	return o.ApplyT(func(v LookupCrawlerResult) *CrawlerTargets { return v.Targets }).(CrawlerTargetsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCrawlerResultOutput{})
}
