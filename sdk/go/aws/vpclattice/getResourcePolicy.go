// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about the resource policy. The resource policy is an IAM policy created by AWS RAM on behalf of the resource owner when they share a resource.
func LookupResourcePolicy(ctx *pulumi.Context, args *LookupResourcePolicyArgs, opts ...pulumi.InvokeOption) (*LookupResourcePolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourcePolicyResult
	err := ctx.Invoke("aws-native:vpclattice:getResourcePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourcePolicyArgs struct {
	// An IAM policy.
	ResourceArn string `pulumi:"resourceArn"`
}

type LookupResourcePolicyResult struct {
	// The Amazon Resource Name (ARN) of the service network or service.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::VpcLattice::ResourcePolicy` for more information about the expected schema for this property.
	Policy interface{} `pulumi:"policy"`
}

func LookupResourcePolicyOutput(ctx *pulumi.Context, args LookupResourcePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupResourcePolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupResourcePolicyResultOutput, error) {
			args := v.(LookupResourcePolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:vpclattice:getResourcePolicy", args, LookupResourcePolicyResultOutput{}, options).(LookupResourcePolicyResultOutput), nil
		}).(LookupResourcePolicyResultOutput)
}

type LookupResourcePolicyOutputArgs struct {
	// An IAM policy.
	ResourceArn pulumi.StringInput `pulumi:"resourceArn"`
}

func (LookupResourcePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePolicyArgs)(nil)).Elem()
}

type LookupResourcePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupResourcePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePolicyResult)(nil)).Elem()
}

func (o LookupResourcePolicyResultOutput) ToLookupResourcePolicyResultOutput() LookupResourcePolicyResultOutput {
	return o
}

func (o LookupResourcePolicyResultOutput) ToLookupResourcePolicyResultOutputWithContext(ctx context.Context) LookupResourcePolicyResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the service network or service.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::VpcLattice::ResourcePolicy` for more information about the expected schema for this property.
func (o LookupResourcePolicyResultOutput) Policy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) interface{} { return v.Policy }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourcePolicyResultOutput{})
}
