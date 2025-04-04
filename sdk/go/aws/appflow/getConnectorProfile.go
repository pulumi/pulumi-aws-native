// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appflow

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppFlow::ConnectorProfile
func LookupConnectorProfile(ctx *pulumi.Context, args *LookupConnectorProfileArgs, opts ...pulumi.InvokeOption) (*LookupConnectorProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectorProfileResult
	err := ctx.Invoke("aws-native:appflow:getConnectorProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorProfileArgs struct {
	// The maximum number of items to retrieve in a single batch.
	ConnectorProfileName string `pulumi:"connectorProfileName"`
}

type LookupConnectorProfileResult struct {
	// Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
	ConnectionMode *ConnectorProfileConnectionMode `pulumi:"connectionMode"`
	// Unique identifier for connector profile resources
	ConnectorProfileArn *string `pulumi:"connectorProfileArn"`
	// A unique Arn for Connector-Profile resource
	CredentialsArn *string `pulumi:"credentialsArn"`
}

func LookupConnectorProfileOutput(ctx *pulumi.Context, args LookupConnectorProfileOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupConnectorProfileResultOutput, error) {
			args := v.(LookupConnectorProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:appflow:getConnectorProfile", args, LookupConnectorProfileResultOutput{}, options).(LookupConnectorProfileResultOutput), nil
		}).(LookupConnectorProfileResultOutput)
}

type LookupConnectorProfileOutputArgs struct {
	// The maximum number of items to retrieve in a single batch.
	ConnectorProfileName pulumi.StringInput `pulumi:"connectorProfileName"`
}

func (LookupConnectorProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorProfileArgs)(nil)).Elem()
}

type LookupConnectorProfileResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorProfileResult)(nil)).Elem()
}

func (o LookupConnectorProfileResultOutput) ToLookupConnectorProfileResultOutput() LookupConnectorProfileResultOutput {
	return o
}

func (o LookupConnectorProfileResultOutput) ToLookupConnectorProfileResultOutputWithContext(ctx context.Context) LookupConnectorProfileResultOutput {
	return o
}

// Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
func (o LookupConnectorProfileResultOutput) ConnectionMode() ConnectorProfileConnectionModePtrOutput {
	return o.ApplyT(func(v LookupConnectorProfileResult) *ConnectorProfileConnectionMode { return v.ConnectionMode }).(ConnectorProfileConnectionModePtrOutput)
}

// Unique identifier for connector profile resources
func (o LookupConnectorProfileResultOutput) ConnectorProfileArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorProfileResult) *string { return v.ConnectorProfileArn }).(pulumi.StringPtrOutput)
}

// A unique Arn for Connector-Profile resource
func (o LookupConnectorProfileResultOutput) CredentialsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorProfileResult) *string { return v.CredentialsArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorProfileResultOutput{})
}
