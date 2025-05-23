// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bedrock

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Bedrock::Flow Resource Type
type Flow struct {
	pulumi.CustomResourceState

	// Arn representation of the Flow
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Identifier for a Flow
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// Time Stamp.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A KMS key ARN
	CustomerEncryptionKeyArn pulumi.StringPtrOutput `pulumi:"customerEncryptionKeyArn"`
	// The definition of the nodes and connections between the nodes in the flow.
	Definition FlowDefinitionPtrOutput `pulumi:"definition"`
	// The Amazon S3 location of the flow definition.
	DefinitionS3Location FlowS3LocationPtrOutput `pulumi:"definitionS3Location"`
	// A JSON string containing a Definition with the same schema as the Definition property of this resource
	DefinitionString pulumi.StringPtrOutput `pulumi:"definitionString"`
	// A map that specifies the mappings for placeholder variables in the prompt flow definition. This enables the customer to inject values obtained at runtime. Variables can be template parameter names, resource logical IDs, resource attributes, or a variable in a key-value map. Only supported with the `DefinitionString` and `DefinitionS3Location` fields.
	//
	// Substitutions must follow the syntax: `${key_name}` or `${variable_1,variable_2,...}` .
	DefinitionSubstitutions pulumi.MapOutput `pulumi:"definitionSubstitutions"`
	// Description of the flow
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ARN of a IAM role
	ExecutionRoleArn pulumi.StringOutput `pulumi:"executionRoleArn"`
	// Name for the flow
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the flow. The following statuses are possible:
	//
	// - NotPrepared – The flow has been created or updated, but hasn't been prepared. If you just created the flow, you can't test it. If you updated the flow, the `DRAFT` version won't contain the latest changes for testing. Send a [PrepareFlow](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_PrepareFlow.html) request to package the latest changes into the `DRAFT` version.
	// - Preparing – The flow is being prepared so that the `DRAFT` version contains the latest changes for testing.
	// - Prepared – The flow is prepared and the `DRAFT` version contains the latest changes for testing.
	// - Failed – The last API operation that you invoked on the flow failed. Send a [GetFlow](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_GetFlow.html) request and check the error message in the `validations` field.
	Status FlowStatusOutput `pulumi:"status"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	Tags          pulumi.StringMapOutput `pulumi:"tags"`
	TestAliasTags pulumi.StringMapOutput `pulumi:"testAliasTags"`
	// Time Stamp.
	UpdatedAt   pulumi.StringOutput       `pulumi:"updatedAt"`
	Validations FlowValidationArrayOutput `pulumi:"validations"`
	// Draft Version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewFlow registers a new resource with the given unique name, arguments, and options.
func NewFlow(ctx *pulumi.Context,
	name string, args *FlowArgs, opts ...pulumi.ResourceOption) (*Flow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExecutionRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'ExecutionRoleArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Flow
	err := ctx.RegisterResource("aws-native:bedrock:Flow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlow gets an existing Flow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowState, opts ...pulumi.ResourceOption) (*Flow, error) {
	var resource Flow
	err := ctx.ReadResource("aws-native:bedrock:Flow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flow resources.
type flowState struct {
}

type FlowState struct {
}

func (FlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowState)(nil)).Elem()
}

type flowArgs struct {
	// A KMS key ARN
	CustomerEncryptionKeyArn *string `pulumi:"customerEncryptionKeyArn"`
	// The definition of the nodes and connections between the nodes in the flow.
	Definition *FlowDefinition `pulumi:"definition"`
	// The Amazon S3 location of the flow definition.
	DefinitionS3Location *FlowS3Location `pulumi:"definitionS3Location"`
	// A JSON string containing a Definition with the same schema as the Definition property of this resource
	DefinitionString *string `pulumi:"definitionString"`
	// A map that specifies the mappings for placeholder variables in the prompt flow definition. This enables the customer to inject values obtained at runtime. Variables can be template parameter names, resource logical IDs, resource attributes, or a variable in a key-value map. Only supported with the `DefinitionString` and `DefinitionS3Location` fields.
	//
	// Substitutions must follow the syntax: `${key_name}` or `${variable_1,variable_2,...}` .
	DefinitionSubstitutions map[string]interface{} `pulumi:"definitionSubstitutions"`
	// Description of the flow
	Description *string `pulumi:"description"`
	// ARN of a IAM role
	ExecutionRoleArn string `pulumi:"executionRoleArn"`
	// Name for the flow
	Name *string `pulumi:"name"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	Tags          map[string]string `pulumi:"tags"`
	TestAliasTags map[string]string `pulumi:"testAliasTags"`
}

// The set of arguments for constructing a Flow resource.
type FlowArgs struct {
	// A KMS key ARN
	CustomerEncryptionKeyArn pulumi.StringPtrInput
	// The definition of the nodes and connections between the nodes in the flow.
	Definition FlowDefinitionPtrInput
	// The Amazon S3 location of the flow definition.
	DefinitionS3Location FlowS3LocationPtrInput
	// A JSON string containing a Definition with the same schema as the Definition property of this resource
	DefinitionString pulumi.StringPtrInput
	// A map that specifies the mappings for placeholder variables in the prompt flow definition. This enables the customer to inject values obtained at runtime. Variables can be template parameter names, resource logical IDs, resource attributes, or a variable in a key-value map. Only supported with the `DefinitionString` and `DefinitionS3Location` fields.
	//
	// Substitutions must follow the syntax: `${key_name}` or `${variable_1,variable_2,...}` .
	DefinitionSubstitutions pulumi.MapInput
	// Description of the flow
	Description pulumi.StringPtrInput
	// ARN of a IAM role
	ExecutionRoleArn pulumi.StringInput
	// Name for the flow
	Name pulumi.StringPtrInput
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	Tags          pulumi.StringMapInput
	TestAliasTags pulumi.StringMapInput
}

func (FlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowArgs)(nil)).Elem()
}

type FlowInput interface {
	pulumi.Input

	ToFlowOutput() FlowOutput
	ToFlowOutputWithContext(ctx context.Context) FlowOutput
}

func (*Flow) ElementType() reflect.Type {
	return reflect.TypeOf((**Flow)(nil)).Elem()
}

func (i *Flow) ToFlowOutput() FlowOutput {
	return i.ToFlowOutputWithContext(context.Background())
}

func (i *Flow) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowOutput)
}

type FlowOutput struct{ *pulumi.OutputState }

func (FlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Flow)(nil)).Elem()
}

func (o FlowOutput) ToFlowOutput() FlowOutput {
	return o
}

func (o FlowOutput) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return o
}

// Arn representation of the Flow
func (o FlowOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Identifier for a Flow
func (o FlowOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// Time Stamp.
func (o FlowOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A KMS key ARN
func (o FlowOutput) CustomerEncryptionKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringPtrOutput { return v.CustomerEncryptionKeyArn }).(pulumi.StringPtrOutput)
}

// The definition of the nodes and connections between the nodes in the flow.
func (o FlowOutput) Definition() FlowDefinitionPtrOutput {
	return o.ApplyT(func(v *Flow) FlowDefinitionPtrOutput { return v.Definition }).(FlowDefinitionPtrOutput)
}

// The Amazon S3 location of the flow definition.
func (o FlowOutput) DefinitionS3Location() FlowS3LocationPtrOutput {
	return o.ApplyT(func(v *Flow) FlowS3LocationPtrOutput { return v.DefinitionS3Location }).(FlowS3LocationPtrOutput)
}

// A JSON string containing a Definition with the same schema as the Definition property of this resource
func (o FlowOutput) DefinitionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringPtrOutput { return v.DefinitionString }).(pulumi.StringPtrOutput)
}

// A map that specifies the mappings for placeholder variables in the prompt flow definition. This enables the customer to inject values obtained at runtime. Variables can be template parameter names, resource logical IDs, resource attributes, or a variable in a key-value map. Only supported with the `DefinitionString` and `DefinitionS3Location` fields.
//
// Substitutions must follow the syntax: `${key_name}` or `${variable_1,variable_2,...}` .
func (o FlowOutput) DefinitionSubstitutions() pulumi.MapOutput {
	return o.ApplyT(func(v *Flow) pulumi.MapOutput { return v.DefinitionSubstitutions }).(pulumi.MapOutput)
}

// Description of the flow
func (o FlowOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// ARN of a IAM role
func (o FlowOutput) ExecutionRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.ExecutionRoleArn }).(pulumi.StringOutput)
}

// Name for the flow
func (o FlowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the flow. The following statuses are possible:
//
// - NotPrepared – The flow has been created or updated, but hasn't been prepared. If you just created the flow, you can't test it. If you updated the flow, the `DRAFT` version won't contain the latest changes for testing. Send a [PrepareFlow](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_PrepareFlow.html) request to package the latest changes into the `DRAFT` version.
// - Preparing – The flow is being prepared so that the `DRAFT` version contains the latest changes for testing.
// - Prepared – The flow is prepared and the `DRAFT` version contains the latest changes for testing.
// - Failed – The last API operation that you invoked on the flow failed. Send a [GetFlow](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_GetFlow.html) request and check the error message in the `validations` field.
func (o FlowOutput) Status() FlowStatusOutput {
	return o.ApplyT(func(v *Flow) FlowStatusOutput { return v.Status }).(FlowStatusOutput)
}

// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
//
// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
func (o FlowOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FlowOutput) TestAliasTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringMapOutput { return v.TestAliasTags }).(pulumi.StringMapOutput)
}

// Time Stamp.
func (o FlowOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o FlowOutput) Validations() FlowValidationArrayOutput {
	return o.ApplyT(func(v *Flow) FlowValidationArrayOutput { return v.Validations }).(FlowValidationArrayOutput)
}

// Draft Version.
func (o FlowOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowInput)(nil)).Elem(), &Flow{})
	pulumi.RegisterOutputType(FlowOutput{})
}
