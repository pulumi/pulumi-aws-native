// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package shield

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Config the role and list of Amazon S3 log buckets used by the Shield Response Team (SRT) to access your AWS account while assisting with attack mitigation.
func LookupDrtAccess(ctx *pulumi.Context, args *LookupDrtAccessArgs, opts ...pulumi.InvokeOption) (*LookupDrtAccessResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDrtAccessResult
	err := ctx.Invoke("aws-native:shield:getDrtAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDrtAccessArgs struct {
	AccountId string `pulumi:"accountId"`
}

type LookupDrtAccessResult struct {
	AccountId *string `pulumi:"accountId"`
	// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources. You can associate up to 10 Amazon S3 buckets with your subscription.
	LogBucketList []string `pulumi:"logBucketList"`
	// Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks. This enables the SRT to inspect your AWS WAF configuration and create or update AWS WAF rules and web ACLs.
	RoleArn *string `pulumi:"roleArn"`
}

func LookupDrtAccessOutput(ctx *pulumi.Context, args LookupDrtAccessOutputArgs, opts ...pulumi.InvokeOption) LookupDrtAccessResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDrtAccessResult, error) {
			args := v.(LookupDrtAccessArgs)
			r, err := LookupDrtAccess(ctx, &args, opts...)
			var s LookupDrtAccessResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDrtAccessResultOutput)
}

type LookupDrtAccessOutputArgs struct {
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (LookupDrtAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDrtAccessArgs)(nil)).Elem()
}

type LookupDrtAccessResultOutput struct{ *pulumi.OutputState }

func (LookupDrtAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDrtAccessResult)(nil)).Elem()
}

func (o LookupDrtAccessResultOutput) ToLookupDrtAccessResultOutput() LookupDrtAccessResultOutput {
	return o
}

func (o LookupDrtAccessResultOutput) ToLookupDrtAccessResultOutputWithContext(ctx context.Context) LookupDrtAccessResultOutput {
	return o
}

func (o LookupDrtAccessResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDrtAccessResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources. You can associate up to 10 Amazon S3 buckets with your subscription.
func (o LookupDrtAccessResultOutput) LogBucketList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDrtAccessResult) []string { return v.LogBucketList }).(pulumi.StringArrayOutput)
}

// Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks. This enables the SRT to inspect your AWS WAF configuration and create or update AWS WAF rules and web ACLs.
func (o LookupDrtAccessResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDrtAccessResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDrtAccessResultOutput{})
}