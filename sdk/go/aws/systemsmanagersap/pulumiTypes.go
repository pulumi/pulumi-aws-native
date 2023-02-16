// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package systemsmanagersap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationCredential struct {
	CredentialType *ApplicationCredentialCredentialType `pulumi:"credentialType"`
	DatabaseName   *string                              `pulumi:"databaseName"`
	SecretId       *string                              `pulumi:"secretId"`
}

// ApplicationCredentialInput is an input type that accepts ApplicationCredentialArgs and ApplicationCredentialOutput values.
// You can construct a concrete instance of `ApplicationCredentialInput` via:
//
//	ApplicationCredentialArgs{...}
type ApplicationCredentialInput interface {
	pulumi.Input

	ToApplicationCredentialOutput() ApplicationCredentialOutput
	ToApplicationCredentialOutputWithContext(context.Context) ApplicationCredentialOutput
}

type ApplicationCredentialArgs struct {
	CredentialType ApplicationCredentialCredentialTypePtrInput `pulumi:"credentialType"`
	DatabaseName   pulumi.StringPtrInput                       `pulumi:"databaseName"`
	SecretId       pulumi.StringPtrInput                       `pulumi:"secretId"`
}

func (ApplicationCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationCredential)(nil)).Elem()
}

func (i ApplicationCredentialArgs) ToApplicationCredentialOutput() ApplicationCredentialOutput {
	return i.ToApplicationCredentialOutputWithContext(context.Background())
}

func (i ApplicationCredentialArgs) ToApplicationCredentialOutputWithContext(ctx context.Context) ApplicationCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialOutput)
}

// ApplicationCredentialArrayInput is an input type that accepts ApplicationCredentialArray and ApplicationCredentialArrayOutput values.
// You can construct a concrete instance of `ApplicationCredentialArrayInput` via:
//
//	ApplicationCredentialArray{ ApplicationCredentialArgs{...} }
type ApplicationCredentialArrayInput interface {
	pulumi.Input

	ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput
	ToApplicationCredentialArrayOutputWithContext(context.Context) ApplicationCredentialArrayOutput
}

type ApplicationCredentialArray []ApplicationCredentialInput

func (ApplicationCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationCredential)(nil)).Elem()
}

func (i ApplicationCredentialArray) ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput {
	return i.ToApplicationCredentialArrayOutputWithContext(context.Background())
}

func (i ApplicationCredentialArray) ToApplicationCredentialArrayOutputWithContext(ctx context.Context) ApplicationCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialArrayOutput)
}

type ApplicationCredentialOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationCredential)(nil)).Elem()
}

func (o ApplicationCredentialOutput) ToApplicationCredentialOutput() ApplicationCredentialOutput {
	return o
}

func (o ApplicationCredentialOutput) ToApplicationCredentialOutputWithContext(ctx context.Context) ApplicationCredentialOutput {
	return o
}

func (o ApplicationCredentialOutput) CredentialType() ApplicationCredentialCredentialTypePtrOutput {
	return o.ApplyT(func(v ApplicationCredential) *ApplicationCredentialCredentialType { return v.CredentialType }).(ApplicationCredentialCredentialTypePtrOutput)
}

func (o ApplicationCredentialOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationCredential) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o ApplicationCredentialOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationCredential) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

type ApplicationCredentialArrayOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationCredential)(nil)).Elem()
}

func (o ApplicationCredentialArrayOutput) ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput {
	return o
}

func (o ApplicationCredentialArrayOutput) ToApplicationCredentialArrayOutputWithContext(ctx context.Context) ApplicationCredentialArrayOutput {
	return o
}

func (o ApplicationCredentialArrayOutput) Index(i pulumi.IntInput) ApplicationCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationCredential {
		return vs[0].([]ApplicationCredential)[vs[1].(int)]
	}).(ApplicationCredentialOutput)
}

// A key-value pair to associate with a resource.
type ApplicationTag struct {
	// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// ApplicationTagInput is an input type that accepts ApplicationTagArgs and ApplicationTagOutput values.
// You can construct a concrete instance of `ApplicationTagInput` via:
//
//	ApplicationTagArgs{...}
type ApplicationTagInput interface {
	pulumi.Input

	ToApplicationTagOutput() ApplicationTagOutput
	ToApplicationTagOutputWithContext(context.Context) ApplicationTagOutput
}

// A key-value pair to associate with a resource.
type ApplicationTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ApplicationTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTag)(nil)).Elem()
}

func (i ApplicationTagArgs) ToApplicationTagOutput() ApplicationTagOutput {
	return i.ToApplicationTagOutputWithContext(context.Background())
}

func (i ApplicationTagArgs) ToApplicationTagOutputWithContext(ctx context.Context) ApplicationTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTagOutput)
}

// ApplicationTagArrayInput is an input type that accepts ApplicationTagArray and ApplicationTagArrayOutput values.
// You can construct a concrete instance of `ApplicationTagArrayInput` via:
//
//	ApplicationTagArray{ ApplicationTagArgs{...} }
type ApplicationTagArrayInput interface {
	pulumi.Input

	ToApplicationTagArrayOutput() ApplicationTagArrayOutput
	ToApplicationTagArrayOutputWithContext(context.Context) ApplicationTagArrayOutput
}

type ApplicationTagArray []ApplicationTagInput

func (ApplicationTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationTag)(nil)).Elem()
}

func (i ApplicationTagArray) ToApplicationTagArrayOutput() ApplicationTagArrayOutput {
	return i.ToApplicationTagArrayOutputWithContext(context.Background())
}

func (i ApplicationTagArray) ToApplicationTagArrayOutputWithContext(ctx context.Context) ApplicationTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTagArrayOutput)
}

// A key-value pair to associate with a resource.
type ApplicationTagOutput struct{ *pulumi.OutputState }

func (ApplicationTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTag)(nil)).Elem()
}

func (o ApplicationTagOutput) ToApplicationTagOutput() ApplicationTagOutput {
	return o
}

func (o ApplicationTagOutput) ToApplicationTagOutputWithContext(ctx context.Context) ApplicationTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o ApplicationTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o ApplicationTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationTag) string { return v.Value }).(pulumi.StringOutput)
}

type ApplicationTagArrayOutput struct{ *pulumi.OutputState }

func (ApplicationTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationTag)(nil)).Elem()
}

func (o ApplicationTagArrayOutput) ToApplicationTagArrayOutput() ApplicationTagArrayOutput {
	return o
}

func (o ApplicationTagArrayOutput) ToApplicationTagArrayOutputWithContext(ctx context.Context) ApplicationTagArrayOutput {
	return o
}

func (o ApplicationTagArrayOutput) Index(i pulumi.IntInput) ApplicationTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationTag {
		return vs[0].([]ApplicationTag)[vs[1].(int)]
	}).(ApplicationTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCredentialInput)(nil)).Elem(), ApplicationCredentialArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCredentialArrayInput)(nil)).Elem(), ApplicationCredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationTagInput)(nil)).Elem(), ApplicationTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationTagArrayInput)(nil)).Elem(), ApplicationTagArray{})
	pulumi.RegisterOutputType(ApplicationCredentialOutput{})
	pulumi.RegisterOutputType(ApplicationCredentialArrayOutput{})
	pulumi.RegisterOutputType(ApplicationTagOutput{})
	pulumi.RegisterOutputType(ApplicationTagArrayOutput{})
}