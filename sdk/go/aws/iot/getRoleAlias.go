// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the AWS::IoT::RoleAlias resource to declare an AWS IoT RoleAlias.
func LookupRoleAlias(ctx *pulumi.Context, args *LookupRoleAliasArgs, opts ...pulumi.InvokeOption) (*LookupRoleAliasResult, error) {
	var rv LookupRoleAliasResult
	err := ctx.Invoke("aws-native:iot:getRoleAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAliasArgs struct {
	RoleAlias string `pulumi:"roleAlias"`
}

type LookupRoleAliasResult struct {
	CredentialDurationSeconds *int           `pulumi:"credentialDurationSeconds"`
	RoleAliasArn              *string        `pulumi:"roleAliasArn"`
	RoleArn                   *string        `pulumi:"roleArn"`
	Tags                      []RoleAliasTag `pulumi:"tags"`
}

func LookupRoleAliasOutput(ctx *pulumi.Context, args LookupRoleAliasOutputArgs, opts ...pulumi.InvokeOption) LookupRoleAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleAliasResult, error) {
			args := v.(LookupRoleAliasArgs)
			r, err := LookupRoleAlias(ctx, &args, opts...)
			var s LookupRoleAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleAliasResultOutput)
}

type LookupRoleAliasOutputArgs struct {
	RoleAlias pulumi.StringInput `pulumi:"roleAlias"`
}

func (LookupRoleAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleAliasArgs)(nil)).Elem()
}

type LookupRoleAliasResultOutput struct{ *pulumi.OutputState }

func (LookupRoleAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleAliasResult)(nil)).Elem()
}

func (o LookupRoleAliasResultOutput) ToLookupRoleAliasResultOutput() LookupRoleAliasResultOutput {
	return o
}

func (o LookupRoleAliasResultOutput) ToLookupRoleAliasResultOutputWithContext(ctx context.Context) LookupRoleAliasResultOutput {
	return o
}

func (o LookupRoleAliasResultOutput) CredentialDurationSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRoleAliasResult) *int { return v.CredentialDurationSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRoleAliasResultOutput) RoleAliasArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAliasResult) *string { return v.RoleAliasArn }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAliasResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAliasResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAliasResultOutput) Tags() RoleAliasTagArrayOutput {
	return o.ApplyT(func(v LookupRoleAliasResult) []RoleAliasTag { return v.Tags }).(RoleAliasTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleAliasResultOutput{})
}