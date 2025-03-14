// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package events

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Events::Connection.
func LookupConnection(ctx *pulumi.Context, args *LookupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupConnectionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectionResult
	err := ctx.Invoke("aws-native:events:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionArgs struct {
	// Name of the connection.
	Name string `pulumi:"name"`
}

type LookupConnectionResult struct {
	// The arn of the connection resource.
	Arn *string `pulumi:"arn"`
	// The authorization parameters to use to authorize with the endpoint.
	//
	// You must include only authorization parameters for the `AuthorizationType` you specify.
	AuthParameters *ConnectionAuthParameters `pulumi:"authParameters"`
	// The type of authorization to use for the connection.
	//
	// > OAUTH tokens are refreshed when a 401 or 407 response is returned.
	AuthorizationType *ConnectionAuthorizationType `pulumi:"authorizationType"`
	// Description of the connection.
	Description *string `pulumi:"description"`
	// The private resource the HTTP request will be sent to.
	InvocationConnectivityParameters *InvocationConnectivityParametersProperties `pulumi:"invocationConnectivityParameters"`
	// The arn of the secrets manager secret created in the customer account.
	SecretArn *string `pulumi:"secretArn"`
}

func LookupConnectionOutput(ctx *pulumi.Context, args LookupConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupConnectionResultOutput, error) {
			args := v.(LookupConnectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:events:getConnection", args, LookupConnectionResultOutput{}, options).(LookupConnectionResultOutput), nil
		}).(LookupConnectionResultOutput)
}

type LookupConnectionOutputArgs struct {
	// Name of the connection.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionArgs)(nil)).Elem()
}

type LookupConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionResult)(nil)).Elem()
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutput() LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutputWithContext(ctx context.Context) LookupConnectionResultOutput {
	return o
}

// The arn of the connection resource.
func (o LookupConnectionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The authorization parameters to use to authorize with the endpoint.
//
// You must include only authorization parameters for the `AuthorizationType` you specify.
func (o LookupConnectionResultOutput) AuthParameters() ConnectionAuthParametersPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *ConnectionAuthParameters { return v.AuthParameters }).(ConnectionAuthParametersPtrOutput)
}

// The type of authorization to use for the connection.
//
// > OAUTH tokens are refreshed when a 401 or 407 response is returned.
func (o LookupConnectionResultOutput) AuthorizationType() ConnectionAuthorizationTypePtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *ConnectionAuthorizationType { return v.AuthorizationType }).(ConnectionAuthorizationTypePtrOutput)
}

// Description of the connection.
func (o LookupConnectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The private resource the HTTP request will be sent to.
func (o LookupConnectionResultOutput) InvocationConnectivityParameters() InvocationConnectivityParametersPropertiesPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *InvocationConnectivityParametersProperties {
		return v.InvocationConnectivityParameters
	}).(InvocationConnectivityParametersPropertiesPtrOutput)
}

// The arn of the secrets manager secret created in the customer account.
func (o LookupConnectionResultOutput) SecretArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.SecretArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionResultOutput{})
}
