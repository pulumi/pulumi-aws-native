// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerprofiles

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A domain defined for 3rd party data source in Profile Service
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainResult
	err := ctx.Invoke("aws-native:customerprofiles:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainArgs struct {
	// The unique name of the domain.
	DomainName string `pulumi:"domainName"`
}

type LookupDomainResult struct {
	// The time of this integration got created
	CreatedAt *string `pulumi:"createdAt"`
	// The URL of the SQS dead letter queue
	DeadLetterQueueUrl *string `pulumi:"deadLetterQueueUrl"`
	// The default encryption key
	DefaultEncryptionKey *string `pulumi:"defaultEncryptionKey"`
	// The default number of days until the data within the domain expires.
	DefaultExpirationDays *int `pulumi:"defaultExpirationDays"`
	// The time of this integration got last updated at
	LastUpdatedAt *string `pulumi:"lastUpdatedAt"`
	// The process of matching duplicate profiles.
	Matching *DomainMatching `pulumi:"matching"`
	// The process of matching duplicate profiles using Rule-Based matching.
	RuleBasedMatching *DomainRuleBasedMatching `pulumi:"ruleBasedMatching"`
	Stats             *DomainStats             `pulumi:"stats"`
	// The tags (keys and values) associated with the domain
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDomainResultOutput, error) {
			args := v.(LookupDomainArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:customerprofiles:getDomain", args, LookupDomainResultOutput{}, options).(LookupDomainResultOutput), nil
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	// The unique name of the domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

// The time of this integration got created
func (o LookupDomainResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The URL of the SQS dead letter queue
func (o LookupDomainResultOutput) DeadLetterQueueUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DeadLetterQueueUrl }).(pulumi.StringPtrOutput)
}

// The default encryption key
func (o LookupDomainResultOutput) DefaultEncryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DefaultEncryptionKey }).(pulumi.StringPtrOutput)
}

// The default number of days until the data within the domain expires.
func (o LookupDomainResultOutput) DefaultExpirationDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *int { return v.DefaultExpirationDays }).(pulumi.IntPtrOutput)
}

// The time of this integration got last updated at
func (o LookupDomainResultOutput) LastUpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.LastUpdatedAt }).(pulumi.StringPtrOutput)
}

// The process of matching duplicate profiles.
func (o LookupDomainResultOutput) Matching() DomainMatchingPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *DomainMatching { return v.Matching }).(DomainMatchingPtrOutput)
}

// The process of matching duplicate profiles using Rule-Based matching.
func (o LookupDomainResultOutput) RuleBasedMatching() DomainRuleBasedMatchingPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *DomainRuleBasedMatching { return v.RuleBasedMatching }).(DomainRuleBasedMatchingPtrOutput)
}

func (o LookupDomainResultOutput) Stats() DomainStatsPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *DomainStats { return v.Stats }).(DomainStatsPtrOutput)
}

// The tags (keys and values) associated with the domain
func (o LookupDomainResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
