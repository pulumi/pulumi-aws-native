// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::RDS::DBProxyTargetGroup
func LookupDbProxyTargetGroup(ctx *pulumi.Context, args *LookupDbProxyTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupDbProxyTargetGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDbProxyTargetGroupResult
	err := ctx.Invoke("aws-native:rds:getDbProxyTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDbProxyTargetGroupArgs struct {
	// The Amazon Resource Name (ARN) representing the target group.
	TargetGroupArn string `pulumi:"targetGroupArn"`
}

type LookupDbProxyTargetGroupResult struct {
	// Displays the settings that control the size and behavior of the connection pool associated with a `DBProxyTarget` .
	ConnectionPoolConfigurationInfo *DbProxyTargetGroupConnectionPoolConfigurationInfoFormat `pulumi:"connectionPoolConfigurationInfo"`
	// One or more DB cluster identifiers.
	DbClusterIdentifiers []string `pulumi:"dbClusterIdentifiers"`
	// One or more DB instance identifiers.
	DbInstanceIdentifiers []string `pulumi:"dbInstanceIdentifiers"`
	// The Amazon Resource Name (ARN) representing the target group.
	TargetGroupArn *string `pulumi:"targetGroupArn"`
}

func LookupDbProxyTargetGroupOutput(ctx *pulumi.Context, args LookupDbProxyTargetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDbProxyTargetGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDbProxyTargetGroupResultOutput, error) {
			args := v.(LookupDbProxyTargetGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:rds:getDbProxyTargetGroup", args, LookupDbProxyTargetGroupResultOutput{}, options).(LookupDbProxyTargetGroupResultOutput), nil
		}).(LookupDbProxyTargetGroupResultOutput)
}

type LookupDbProxyTargetGroupOutputArgs struct {
	// The Amazon Resource Name (ARN) representing the target group.
	TargetGroupArn pulumi.StringInput `pulumi:"targetGroupArn"`
}

func (LookupDbProxyTargetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbProxyTargetGroupArgs)(nil)).Elem()
}

type LookupDbProxyTargetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDbProxyTargetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbProxyTargetGroupResult)(nil)).Elem()
}

func (o LookupDbProxyTargetGroupResultOutput) ToLookupDbProxyTargetGroupResultOutput() LookupDbProxyTargetGroupResultOutput {
	return o
}

func (o LookupDbProxyTargetGroupResultOutput) ToLookupDbProxyTargetGroupResultOutputWithContext(ctx context.Context) LookupDbProxyTargetGroupResultOutput {
	return o
}

// Displays the settings that control the size and behavior of the connection pool associated with a `DBProxyTarget` .
func (o LookupDbProxyTargetGroupResultOutput) ConnectionPoolConfigurationInfo() DbProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput {
	return o.ApplyT(func(v LookupDbProxyTargetGroupResult) *DbProxyTargetGroupConnectionPoolConfigurationInfoFormat {
		return v.ConnectionPoolConfigurationInfo
	}).(DbProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput)
}

// One or more DB cluster identifiers.
func (o LookupDbProxyTargetGroupResultOutput) DbClusterIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDbProxyTargetGroupResult) []string { return v.DbClusterIdentifiers }).(pulumi.StringArrayOutput)
}

// One or more DB instance identifiers.
func (o LookupDbProxyTargetGroupResultOutput) DbInstanceIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDbProxyTargetGroupResult) []string { return v.DbInstanceIdentifiers }).(pulumi.StringArrayOutput)
}

// The Amazon Resource Name (ARN) representing the target group.
func (o LookupDbProxyTargetGroupResultOutput) TargetGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDbProxyTargetGroupResult) *string { return v.TargetGroupArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDbProxyTargetGroupResultOutput{})
}
