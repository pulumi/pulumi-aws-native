// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package proton

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema describing various properties for AWS Proton Environment Account Connections resources.
func LookupEnvironmentAccountConnection(ctx *pulumi.Context, args *LookupEnvironmentAccountConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentAccountConnectionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentAccountConnectionResult
	err := ctx.Invoke("aws-native:proton:getEnvironmentAccountConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentAccountConnectionArgs struct {
	// The Amazon Resource Name (ARN) of the environment account connection.
	Arn string `pulumi:"arn"`
}

type LookupEnvironmentAccountConnectionResult struct {
	// The Amazon Resource Name (ARN) of the environment account connection.
	Arn *string `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of an IAM service role in the environment account. AWS Proton uses this role to provision infrastructure resources using CodeBuild-based provisioning in the associated environment account.
	CodebuildRoleArn *string `pulumi:"codebuildRoleArn"`
	// The Amazon Resource Name (ARN) of the IAM service role that AWS Proton uses when provisioning directly defined components in the associated environment account. It determines the scope of infrastructure that a component can provision in the account.
	ComponentRoleArn *string `pulumi:"componentRoleArn"`
	// The environment account that's connected to the environment account connection.
	EnvironmentAccountId *string `pulumi:"environmentAccountId"`
	// The name of the AWS Proton environment that's created in the associated management account.
	EnvironmentName *string `pulumi:"environmentName"`
	// The ID of the environment account connection.
	Id *string `pulumi:"id"`
	// The ID of the management account that accepts or rejects the environment account connection. You create an manage the AWS Proton environment in this account. If the management account accepts the environment account connection, AWS Proton can use the associated IAM role to provision environment infrastructure resources in the associated environment account.
	ManagementAccountId *string `pulumi:"managementAccountId"`
	// The Amazon Resource Name (ARN) of the IAM service role that's created in the environment account. AWS Proton uses this role to provision infrastructure resources in the associated environment account.
	RoleArn *string `pulumi:"roleArn"`
	// The status of the environment account connection.
	Status *EnvironmentAccountConnectionStatus `pulumi:"status"`
	// <p>An optional list of metadata items that you can associate with the Proton environment account connection. A tag is a key-value pair.</p>
	//          <p>For more information, see <a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html">Proton resources and tagging</a> in the
	//         <i>Proton User Guide</i>.</p>
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupEnvironmentAccountConnectionOutput(ctx *pulumi.Context, args LookupEnvironmentAccountConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentAccountConnectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentAccountConnectionResultOutput, error) {
			args := v.(LookupEnvironmentAccountConnectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:proton:getEnvironmentAccountConnection", args, LookupEnvironmentAccountConnectionResultOutput{}, options).(LookupEnvironmentAccountConnectionResultOutput), nil
		}).(LookupEnvironmentAccountConnectionResultOutput)
}

type LookupEnvironmentAccountConnectionOutputArgs struct {
	// The Amazon Resource Name (ARN) of the environment account connection.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupEnvironmentAccountConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentAccountConnectionArgs)(nil)).Elem()
}

type LookupEnvironmentAccountConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentAccountConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentAccountConnectionResult)(nil)).Elem()
}

func (o LookupEnvironmentAccountConnectionResultOutput) ToLookupEnvironmentAccountConnectionResultOutput() LookupEnvironmentAccountConnectionResultOutput {
	return o
}

func (o LookupEnvironmentAccountConnectionResultOutput) ToLookupEnvironmentAccountConnectionResultOutputWithContext(ctx context.Context) LookupEnvironmentAccountConnectionResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the environment account connection.
func (o LookupEnvironmentAccountConnectionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentAccountConnectionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of an IAM service role in the environment account. AWS Proton uses this role to provision infrastructure resources using CodeBuild-based provisioning in the associated environment account.
func (o LookupEnvironmentAccountConnectionResultOutput) CodebuildRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentAccountConnectionResult) *string { return v.CodebuildRoleArn }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the IAM service role that AWS Proton uses when provisioning directly defined components in the associated environment account. It determines the scope of infrastructure that a component can provision in the account.
func (o LookupEnvironmentAccountConnectionResultOutput) ComponentRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentAccountConnectionResult) *string { return v.ComponentRoleArn }).(pulumi.StringPtrOutput)
}

// The environment account that's connected to the environment account connection.
func (o LookupEnvironmentAccountConnectionResultOutput) EnvironmentAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentAccountConnectionResult) *string { return v.EnvironmentAccountId }).(pulumi.StringPtrOutput)
}

// The name of the AWS Proton environment that's created in the associated management account.
func (o LookupEnvironmentAccountConnectionResultOutput) EnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentAccountConnectionResult) *string { return v.EnvironmentName }).(pulumi.StringPtrOutput)
}

// The ID of the environment account connection.
func (o LookupEnvironmentAccountConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentAccountConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The ID of the management account that accepts or rejects the environment account connection. You create an manage the AWS Proton environment in this account. If the management account accepts the environment account connection, AWS Proton can use the associated IAM role to provision environment infrastructure resources in the associated environment account.
func (o LookupEnvironmentAccountConnectionResultOutput) ManagementAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentAccountConnectionResult) *string { return v.ManagementAccountId }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the IAM service role that's created in the environment account. AWS Proton uses this role to provision infrastructure resources in the associated environment account.
func (o LookupEnvironmentAccountConnectionResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentAccountConnectionResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// The status of the environment account connection.
func (o LookupEnvironmentAccountConnectionResultOutput) Status() EnvironmentAccountConnectionStatusPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentAccountConnectionResult) *EnvironmentAccountConnectionStatus { return v.Status }).(EnvironmentAccountConnectionStatusPtrOutput)
}

// <p>An optional list of metadata items that you can associate with the Proton environment account connection. A tag is a key-value pair.</p>
//
//	 <p>For more information, see <a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html">Proton resources and tagging</a> in the
//	<i>Proton User Guide</i>.</p>
func (o LookupEnvironmentAccountConnectionResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupEnvironmentAccountConnectionResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentAccountConnectionResultOutput{})
}
