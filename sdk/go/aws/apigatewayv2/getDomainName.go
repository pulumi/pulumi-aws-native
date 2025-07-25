// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::ApiGatewayV2::DomainName“ resource specifies a custom domain name for your API in Amazon API Gateway (API Gateway).
//
//	You can use a custom domain name to provide a URL that's more intuitive and easier to recall. For more information about using custom domain names, see [Set up Custom Domain Name for an API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html) in the *API Gateway Developer Guide*.
func LookupDomainName(ctx *pulumi.Context, args *LookupDomainNameArgs, opts ...pulumi.InvokeOption) (*LookupDomainNameResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainNameResult
	err := ctx.Invoke("aws-native:apigatewayv2:getDomainName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainNameArgs struct {
	// The custom domain name for your API in Amazon API Gateway. Uppercase letters and the underscore (``_``) character are not supported.
	DomainName string `pulumi:"domainName"`
}

type LookupDomainNameResult struct {
	// Represents an Amazon Resource Name (ARN).
	DomainNameArn *string `pulumi:"domainNameArn"`
	// The domain name configurations.
	DomainNameConfigurations []DomainNameConfiguration `pulumi:"domainNameConfigurations"`
	// The mutual TLS authentication configuration for a custom domain name.
	MutualTlsAuthentication *DomainNameMutualTlsAuthentication `pulumi:"mutualTlsAuthentication"`
	// The domain name associated with the regional endpoint for this custom domain name. You set up this association by adding a DNS record that points the custom domain name to this regional domain name.
	RegionalDomainName *string `pulumi:"regionalDomainName"`
	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	RegionalHostedZoneId *string `pulumi:"regionalHostedZoneId"`
	// The routing mode API Gateway uses to route traffic to your APIs.
	RoutingMode *DomainNameRoutingMode `pulumi:"routingMode"`
	// The collection of tags associated with a domain name.
	Tags map[string]string `pulumi:"tags"`
}

func LookupDomainNameOutput(ctx *pulumi.Context, args LookupDomainNameOutputArgs, opts ...pulumi.InvokeOption) LookupDomainNameResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDomainNameResultOutput, error) {
			args := v.(LookupDomainNameArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:apigatewayv2:getDomainName", args, LookupDomainNameResultOutput{}, options).(LookupDomainNameResultOutput), nil
		}).(LookupDomainNameResultOutput)
}

type LookupDomainNameOutputArgs struct {
	// The custom domain name for your API in Amazon API Gateway. Uppercase letters and the underscore (``_``) character are not supported.
	DomainName pulumi.StringInput `pulumi:"domainName"`
}

func (LookupDomainNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainNameArgs)(nil)).Elem()
}

type LookupDomainNameResultOutput struct{ *pulumi.OutputState }

func (LookupDomainNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainNameResult)(nil)).Elem()
}

func (o LookupDomainNameResultOutput) ToLookupDomainNameResultOutput() LookupDomainNameResultOutput {
	return o
}

func (o LookupDomainNameResultOutput) ToLookupDomainNameResultOutputWithContext(ctx context.Context) LookupDomainNameResultOutput {
	return o
}

// Represents an Amazon Resource Name (ARN).
func (o LookupDomainNameResultOutput) DomainNameArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainNameResult) *string { return v.DomainNameArn }).(pulumi.StringPtrOutput)
}

// The domain name configurations.
func (o LookupDomainNameResultOutput) DomainNameConfigurations() DomainNameConfigurationArrayOutput {
	return o.ApplyT(func(v LookupDomainNameResult) []DomainNameConfiguration { return v.DomainNameConfigurations }).(DomainNameConfigurationArrayOutput)
}

// The mutual TLS authentication configuration for a custom domain name.
func (o LookupDomainNameResultOutput) MutualTlsAuthentication() DomainNameMutualTlsAuthenticationPtrOutput {
	return o.ApplyT(func(v LookupDomainNameResult) *DomainNameMutualTlsAuthentication { return v.MutualTlsAuthentication }).(DomainNameMutualTlsAuthenticationPtrOutput)
}

// The domain name associated with the regional endpoint for this custom domain name. You set up this association by adding a DNS record that points the custom domain name to this regional domain name.
func (o LookupDomainNameResultOutput) RegionalDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainNameResult) *string { return v.RegionalDomainName }).(pulumi.StringPtrOutput)
}

// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
func (o LookupDomainNameResultOutput) RegionalHostedZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainNameResult) *string { return v.RegionalHostedZoneId }).(pulumi.StringPtrOutput)
}

// The routing mode API Gateway uses to route traffic to your APIs.
func (o LookupDomainNameResultOutput) RoutingMode() DomainNameRoutingModePtrOutput {
	return o.ApplyT(func(v LookupDomainNameResult) *DomainNameRoutingMode { return v.RoutingMode }).(DomainNameRoutingModePtrOutput)
}

// The collection of tags associated with a domain name.
func (o LookupDomainNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainNameResultOutput{})
}
