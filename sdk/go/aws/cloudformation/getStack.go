// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::CloudFormation::Stack resource nests a stack as a resource in a top-level template.
func LookupStack(ctx *pulumi.Context, args *LookupStackArgs, opts ...pulumi.InvokeOption) (*LookupStackResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStackResult
	err := ctx.Invoke("aws-native:cloudformation:getStack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStackArgs struct {
	// Unique identifier of the stack.
	StackId string `pulumi:"stackId"`
}

type LookupStackResult struct {
	// In some cases, you must explicitly acknowledge that your stack template contains certain capabilities in order for CloudFormation to create the stack.
	//
	// - `CAPABILITY_IAM` and `CAPABILITY_NAMED_IAM`
	//
	// Some stack templates might include resources that can affect permissions in your AWS account ; for example, by creating new AWS Identity and Access Management (IAM) users. For those stacks, you must explicitly acknowledge this by specifying one of these capabilities.
	//
	// The following IAM resources require you to specify either the `CAPABILITY_IAM` or `CAPABILITY_NAMED_IAM` capability.
	//
	// - If you have IAM resources, you can specify either capability.
	// - If you have IAM resources with custom names, you *must* specify `CAPABILITY_NAMED_IAM` .
	// - If you don't specify either of these capabilities, CloudFormation returns an `InsufficientCapabilities` error.
	//
	// If your stack template contains these resources, we recommend that you review all permissions associated with them and edit their permissions if necessary.
	//
	// - [AWS::IAM::AccessKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-accesskey.html)
	// - [AWS::IAM::Group](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-group.html)
	// - [AWS::IAM::InstanceProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-instanceprofile.html)
	// - [AWS::IAM::Policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-policy.html)
	// - [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-role.html)
	// - [AWS::IAM::User](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-user.html)
	// - [AWS::IAM::UserToGroupAddition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-usertogroupaddition.html)
	//
	// For more information, see [Acknowledging IAM resources in CloudFormation templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/control-access-with-iam.html#using-iam-capabilities) in the *AWS CloudFormation User Guide* .
	// - `CAPABILITY_AUTO_EXPAND`
	//
	// Some template contain macros. Macros perform custom processing on templates; this can include simple actions like find-and-replace operations, all the way to extensive transformations of entire templates. Because of this, users typically create a change set from the processed template, so that they can review the changes resulting from the macros before actually creating the stack. If your stack template contains one or more macros, and you choose to create a stack directly from the processed template, without first reviewing the resulting changes in a change set, you must acknowledge this capability. This includes the [AWS::Include](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-include.html) and [AWS::Serverless](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-serverless.html) transforms, which are macros hosted by CloudFormation .
	//
	// If you want to create a stack from a stack template that contains macros *and* nested stacks, you must create the stack directly from the template using this capability.
	//
	// > You should only create stacks directly from a stack template that contains macros if you know what processing the macro performs.
	// >
	// > Each macro relies on an underlying Lambda service function for processing stack templates. Be aware that the Lambda function owner can update the function operation without CloudFormation being notified.
	//
	// For more information, see [Perform custom processing on CloudFormation templates with template macros](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html) in the *AWS CloudFormation User Guide* .
	Capabilities []StackCapabilitiesItem `pulumi:"capabilities"`
	// The unique ID of the change set.
	ChangeSetId *string `pulumi:"changeSetId"`
	// The time at which the stack was created.
	CreationTime *string `pulumi:"creationTime"`
	// A user-defined description associated with the stack.
	Description *string `pulumi:"description"`
	// Set to `true` to disable rollback of the stack if stack creation failed. You can specify either `DisableRollback` or `OnFailure` , but not both.
	//
	// Default: `false`
	DisableRollback *bool `pulumi:"disableRollback"`
	// Whether to enable termination protection on the specified stack. If a user attempts to delete a stack with termination protection enabled, the operation fails and the stack remains unchanged. For more information, see [Protect CloudFormation stacks from being deleted](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-protect-stacks.html) in the *AWS CloudFormation User Guide* . Termination protection is deactivated on stacks by default.
	//
	// For nested stacks, termination protection is set on the root stack and can't be changed directly on the nested stack.
	EnableTerminationProtection *bool `pulumi:"enableTerminationProtection"`
	// The time the stack was last updated. This field will only be returned if the stack has been updated at least once.
	LastUpdateTime *string `pulumi:"lastUpdateTime"`
	// The Amazon SNS topic ARNs to publish stack related events. You can find your Amazon SNS topic ARNs using the Amazon SNS console or your Command Line Interface (CLI).
	NotificationArns []string `pulumi:"notificationArns"`
	// A list of output structures.
	Outputs []StackOutputType `pulumi:"outputs"`
	// The set value pairs that represent the parameters passed to CloudFormation when this nested stack is created. Each parameter has a name corresponding to a parameter defined in the embedded template and a value representing the value that you want to set for the parameter.
	//
	// > If you use the `Ref` function to pass a parameter value to a nested stack, comma-delimited list parameters must be of type `String` . In other words, you can't pass values that are of type `CommaDelimitedList` to nested stacks.
	//
	// Required if the nested stack requires input parameters.
	//
	// Whether an update causes interruptions depends on the resources that are being updated. An update never causes a nested stack to be replaced.
	Parameters map[string]string `pulumi:"parameters"`
	// For nested stacks, the stack ID of the direct parent of this stack. For the first level of nested stacks, the root stack is also the parent stack.
	//
	// For more information, see [Nested stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html) in the *AWS CloudFormation User Guide* .
	ParentId *string `pulumi:"parentId"`
	// The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes to create the stack. CloudFormation uses the role's credentials to make calls on your behalf. CloudFormation always uses this role for all future operations on the stack. Provided that users have permission to operate on the stack, CloudFormation uses this role even if the users don't have permission to pass it. Ensure that the role grants least privilege.
	//
	// If you don't specify a value, CloudFormation uses the role that was previously associated with the stack. If no role is available, CloudFormation uses a temporary session that's generated from your user credentials.
	RoleArn *string `pulumi:"roleArn"`
	// For nested stacks, the stack ID of the top-level stack to which the nested stack ultimately belongs.
	//
	// For more information, see [Nested stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html) in the *AWS CloudFormation User Guide* .
	RootId *string `pulumi:"rootId"`
	// Unique identifier of the stack.
	StackId *string `pulumi:"stackId"`
	// Structure that contains the stack policy body. For more information, see [Prevent updates to stack resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html) in the *AWS CloudFormation User Guide* . You can specify either the `StackPolicyBody` or the `StackPolicyURL` parameter, but not both.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CloudFormation::Stack` for more information about the expected schema for this property.
	StackPolicyBody interface{} `pulumi:"stackPolicyBody"`
	// Current status of the stack.
	StackStatus *StackStatus `pulumi:"stackStatus"`
	// Success/failure message associated with the stack status.
	StackStatusReason *string `pulumi:"stackStatusReason"`
	// Key-value pairs to associate with this stack. CloudFormation also propagates these tags to the resources created in the stack. A maximum number of 50 tags can be specified.
	Tags []aws.Tag `pulumi:"tags"`
	// Structure that contains the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	//
	// Conditional: You must specify either the `TemplateBody` or the `TemplateURL` parameter, but not both.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CloudFormation::Stack` for more information about the expected schema for this property.
	TemplateBody interface{} `pulumi:"templateBody"`
	// The length of time, in minutes, that CloudFormation waits for the nested stack to reach the `CREATE_COMPLETE` state. The default is no timeout. When CloudFormation detects that the nested stack has reached the `CREATE_COMPLETE` state, it marks the nested stack resource as `CREATE_COMPLETE` in the parent stack and resumes creating the parent stack. If the timeout period expires before the nested stack reaches `CREATE_COMPLETE` , CloudFormation marks the nested stack as failed and rolls back both the nested stack and parent stack.
	//
	// Updates aren't supported.
	TimeoutInMinutes *int `pulumi:"timeoutInMinutes"`
}

func LookupStackOutput(ctx *pulumi.Context, args LookupStackOutputArgs, opts ...pulumi.InvokeOption) LookupStackResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStackResultOutput, error) {
			args := v.(LookupStackArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:cloudformation:getStack", args, LookupStackResultOutput{}, options).(LookupStackResultOutput), nil
		}).(LookupStackResultOutput)
}

type LookupStackOutputArgs struct {
	// Unique identifier of the stack.
	StackId pulumi.StringInput `pulumi:"stackId"`
}

func (LookupStackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackArgs)(nil)).Elem()
}

type LookupStackResultOutput struct{ *pulumi.OutputState }

func (LookupStackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackResult)(nil)).Elem()
}

func (o LookupStackResultOutput) ToLookupStackResultOutput() LookupStackResultOutput {
	return o
}

func (o LookupStackResultOutput) ToLookupStackResultOutputWithContext(ctx context.Context) LookupStackResultOutput {
	return o
}

// In some cases, you must explicitly acknowledge that your stack template contains certain capabilities in order for CloudFormation to create the stack.
//
// - `CAPABILITY_IAM` and `CAPABILITY_NAMED_IAM`
//
// Some stack templates might include resources that can affect permissions in your AWS account ; for example, by creating new AWS Identity and Access Management (IAM) users. For those stacks, you must explicitly acknowledge this by specifying one of these capabilities.
//
// The following IAM resources require you to specify either the `CAPABILITY_IAM` or `CAPABILITY_NAMED_IAM` capability.
//
// - If you have IAM resources, you can specify either capability.
// - If you have IAM resources with custom names, you *must* specify `CAPABILITY_NAMED_IAM` .
// - If you don't specify either of these capabilities, CloudFormation returns an `InsufficientCapabilities` error.
//
// If your stack template contains these resources, we recommend that you review all permissions associated with them and edit their permissions if necessary.
//
// - [AWS::IAM::AccessKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-accesskey.html)
// - [AWS::IAM::Group](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-group.html)
// - [AWS::IAM::InstanceProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-instanceprofile.html)
// - [AWS::IAM::Policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-policy.html)
// - [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-role.html)
// - [AWS::IAM::User](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-user.html)
// - [AWS::IAM::UserToGroupAddition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-usertogroupaddition.html)
//
// For more information, see [Acknowledging IAM resources in CloudFormation templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/control-access-with-iam.html#using-iam-capabilities) in the *AWS CloudFormation User Guide* .
// - `CAPABILITY_AUTO_EXPAND`
//
// Some template contain macros. Macros perform custom processing on templates; this can include simple actions like find-and-replace operations, all the way to extensive transformations of entire templates. Because of this, users typically create a change set from the processed template, so that they can review the changes resulting from the macros before actually creating the stack. If your stack template contains one or more macros, and you choose to create a stack directly from the processed template, without first reviewing the resulting changes in a change set, you must acknowledge this capability. This includes the [AWS::Include](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-include.html) and [AWS::Serverless](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-serverless.html) transforms, which are macros hosted by CloudFormation .
//
// If you want to create a stack from a stack template that contains macros *and* nested stacks, you must create the stack directly from the template using this capability.
//
// > You should only create stacks directly from a stack template that contains macros if you know what processing the macro performs.
// >
// > Each macro relies on an underlying Lambda service function for processing stack templates. Be aware that the Lambda function owner can update the function operation without CloudFormation being notified.
//
// For more information, see [Perform custom processing on CloudFormation templates with template macros](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html) in the *AWS CloudFormation User Guide* .
func (o LookupStackResultOutput) Capabilities() StackCapabilitiesItemArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []StackCapabilitiesItem { return v.Capabilities }).(StackCapabilitiesItemArrayOutput)
}

// The unique ID of the change set.
func (o LookupStackResultOutput) ChangeSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.ChangeSetId }).(pulumi.StringPtrOutput)
}

// The time at which the stack was created.
func (o LookupStackResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// A user-defined description associated with the stack.
func (o LookupStackResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Set to `true` to disable rollback of the stack if stack creation failed. You can specify either `DisableRollback` or `OnFailure` , but not both.
//
// Default: `false`
func (o LookupStackResultOutput) DisableRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *bool { return v.DisableRollback }).(pulumi.BoolPtrOutput)
}

// Whether to enable termination protection on the specified stack. If a user attempts to delete a stack with termination protection enabled, the operation fails and the stack remains unchanged. For more information, see [Protect CloudFormation stacks from being deleted](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-protect-stacks.html) in the *AWS CloudFormation User Guide* . Termination protection is deactivated on stacks by default.
//
// For nested stacks, termination protection is set on the root stack and can't be changed directly on the nested stack.
func (o LookupStackResultOutput) EnableTerminationProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *bool { return v.EnableTerminationProtection }).(pulumi.BoolPtrOutput)
}

// The time the stack was last updated. This field will only be returned if the stack has been updated at least once.
func (o LookupStackResultOutput) LastUpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.LastUpdateTime }).(pulumi.StringPtrOutput)
}

// The Amazon SNS topic ARNs to publish stack related events. You can find your Amazon SNS topic ARNs using the Amazon SNS console or your Command Line Interface (CLI).
func (o LookupStackResultOutput) NotificationArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []string { return v.NotificationArns }).(pulumi.StringArrayOutput)
}

// A list of output structures.
func (o LookupStackResultOutput) Outputs() StackOutputTypeArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []StackOutputType { return v.Outputs }).(StackOutputTypeArrayOutput)
}

// The set value pairs that represent the parameters passed to CloudFormation when this nested stack is created. Each parameter has a name corresponding to a parameter defined in the embedded template and a value representing the value that you want to set for the parameter.
//
// > If you use the `Ref` function to pass a parameter value to a nested stack, comma-delimited list parameters must be of type `String` . In other words, you can't pass values that are of type `CommaDelimitedList` to nested stacks.
//
// Required if the nested stack requires input parameters.
//
// Whether an update causes interruptions depends on the resources that are being updated. An update never causes a nested stack to be replaced.
func (o LookupStackResultOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStackResult) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

// For nested stacks, the stack ID of the direct parent of this stack. For the first level of nested stacks, the root stack is also the parent stack.
//
// For more information, see [Nested stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html) in the *AWS CloudFormation User Guide* .
func (o LookupStackResultOutput) ParentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.ParentId }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes to create the stack. CloudFormation uses the role's credentials to make calls on your behalf. CloudFormation always uses this role for all future operations on the stack. Provided that users have permission to operate on the stack, CloudFormation uses this role even if the users don't have permission to pass it. Ensure that the role grants least privilege.
//
// If you don't specify a value, CloudFormation uses the role that was previously associated with the stack. If no role is available, CloudFormation uses a temporary session that's generated from your user credentials.
func (o LookupStackResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// For nested stacks, the stack ID of the top-level stack to which the nested stack ultimately belongs.
//
// For more information, see [Nested stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html) in the *AWS CloudFormation User Guide* .
func (o LookupStackResultOutput) RootId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.RootId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the stack.
func (o LookupStackResultOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.StackId }).(pulumi.StringPtrOutput)
}

// Structure that contains the stack policy body. For more information, see [Prevent updates to stack resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html) in the *AWS CloudFormation User Guide* . You can specify either the `StackPolicyBody` or the `StackPolicyURL` parameter, but not both.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CloudFormation::Stack` for more information about the expected schema for this property.
func (o LookupStackResultOutput) StackPolicyBody() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupStackResult) interface{} { return v.StackPolicyBody }).(pulumi.AnyOutput)
}

// Current status of the stack.
func (o LookupStackResultOutput) StackStatus() StackStatusPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *StackStatus { return v.StackStatus }).(StackStatusPtrOutput)
}

// Success/failure message associated with the stack status.
func (o LookupStackResultOutput) StackStatusReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.StackStatusReason }).(pulumi.StringPtrOutput)
}

// Key-value pairs to associate with this stack. CloudFormation also propagates these tags to the resources created in the stack. A maximum number of 50 tags can be specified.
func (o LookupStackResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// Structure that contains the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.
//
// Conditional: You must specify either the `TemplateBody` or the `TemplateURL` parameter, but not both.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CloudFormation::Stack` for more information about the expected schema for this property.
func (o LookupStackResultOutput) TemplateBody() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupStackResult) interface{} { return v.TemplateBody }).(pulumi.AnyOutput)
}

// The length of time, in minutes, that CloudFormation waits for the nested stack to reach the `CREATE_COMPLETE` state. The default is no timeout. When CloudFormation detects that the nested stack has reached the `CREATE_COMPLETE` state, it marks the nested stack resource as `CREATE_COMPLETE` in the parent stack and resumes creating the parent stack. If the timeout period expires before the nested stack reaches `CREATE_COMPLETE` , CloudFormation marks the nested stack as failed and rolls back both the nested stack and parent stack.
//
// Updates aren't supported.
func (o LookupStackResultOutput) TimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *int { return v.TimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStackResultOutput{})
}
