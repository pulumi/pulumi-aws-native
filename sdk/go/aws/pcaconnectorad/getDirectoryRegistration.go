// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pcaconnectorad

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::PCAConnectorAD::DirectoryRegistration Resource Type
func LookupDirectoryRegistration(ctx *pulumi.Context, args *LookupDirectoryRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupDirectoryRegistrationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDirectoryRegistrationResult
	err := ctx.Invoke("aws-native:pcaconnectorad:getDirectoryRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDirectoryRegistrationArgs struct {
	// The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html) .
	DirectoryRegistrationArn string `pulumi:"directoryRegistrationArn"`
}

type LookupDirectoryRegistrationResult struct {
	// The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html) .
	DirectoryRegistrationArn *string `pulumi:"directoryRegistrationArn"`
	// Metadata assigned to a directory registration consisting of a key-value pair.
	Tags map[string]string `pulumi:"tags"`
}

func LookupDirectoryRegistrationOutput(ctx *pulumi.Context, args LookupDirectoryRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupDirectoryRegistrationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDirectoryRegistrationResultOutput, error) {
			args := v.(LookupDirectoryRegistrationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:pcaconnectorad:getDirectoryRegistration", args, LookupDirectoryRegistrationResultOutput{}, options).(LookupDirectoryRegistrationResultOutput), nil
		}).(LookupDirectoryRegistrationResultOutput)
}

type LookupDirectoryRegistrationOutputArgs struct {
	// The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html) .
	DirectoryRegistrationArn pulumi.StringInput `pulumi:"directoryRegistrationArn"`
}

func (LookupDirectoryRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDirectoryRegistrationArgs)(nil)).Elem()
}

type LookupDirectoryRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupDirectoryRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDirectoryRegistrationResult)(nil)).Elem()
}

func (o LookupDirectoryRegistrationResultOutput) ToLookupDirectoryRegistrationResultOutput() LookupDirectoryRegistrationResultOutput {
	return o
}

func (o LookupDirectoryRegistrationResultOutput) ToLookupDirectoryRegistrationResultOutputWithContext(ctx context.Context) LookupDirectoryRegistrationResultOutput {
	return o
}

// The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html) .
func (o LookupDirectoryRegistrationResultOutput) DirectoryRegistrationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDirectoryRegistrationResult) *string { return v.DirectoryRegistrationArn }).(pulumi.StringPtrOutput)
}

// Metadata assigned to a directory registration consisting of a key-value pair.
func (o LookupDirectoryRegistrationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDirectoryRegistrationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDirectoryRegistrationResultOutput{})
}
