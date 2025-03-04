// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::ECR::RepositoryCreationTemplate is used to create repository with configuration from a pre-defined template.
func LookupRepositoryCreationTemplate(ctx *pulumi.Context, args *LookupRepositoryCreationTemplateArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryCreationTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRepositoryCreationTemplateResult
	err := ctx.Invoke("aws-native:ecr:getRepositoryCreationTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRepositoryCreationTemplateArgs struct {
	// The prefix use to match the repository name and apply the template.
	Prefix string `pulumi:"prefix"`
}

type LookupRepositoryCreationTemplateResult struct {
	// A list of enumerable Strings representing the repository creation scenarios that the template will apply towards.
	AppliedFor []RepositoryCreationTemplateAppliedForItem `pulumi:"appliedFor"`
	// Create timestamp of the template.
	CreatedAt *string `pulumi:"createdAt"`
	// The ARN of the role to be assumed by ECR. This role must be in the same account as the registry that you are configuring.
	CustomRoleArn *string `pulumi:"customRoleArn"`
	// The description of the template.
	Description *string `pulumi:"description"`
	// The encryption configuration associated with the repository creation template.
	EncryptionConfiguration *RepositoryCreationTemplateEncryptionConfiguration `pulumi:"encryptionConfiguration"`
	// The image tag mutability setting for the repository.
	ImageTagMutability *RepositoryCreationTemplateImageTagMutability `pulumi:"imageTagMutability"`
	// The JSON lifecycle policy text to apply to the repository. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html
	LifecyclePolicy *string `pulumi:"lifecyclePolicy"`
	// The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html
	RepositoryPolicy *string `pulumi:"repositoryPolicy"`
	// An array of key-value pairs to apply to this resource.
	ResourceTags []RepositoryCreationTemplateTag `pulumi:"resourceTags"`
	// Update timestamp of the template.
	UpdatedAt *string `pulumi:"updatedAt"`
}

func LookupRepositoryCreationTemplateOutput(ctx *pulumi.Context, args LookupRepositoryCreationTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryCreationTemplateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRepositoryCreationTemplateResultOutput, error) {
			args := v.(LookupRepositoryCreationTemplateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ecr:getRepositoryCreationTemplate", args, LookupRepositoryCreationTemplateResultOutput{}, options).(LookupRepositoryCreationTemplateResultOutput), nil
		}).(LookupRepositoryCreationTemplateResultOutput)
}

type LookupRepositoryCreationTemplateOutputArgs struct {
	// The prefix use to match the repository name and apply the template.
	Prefix pulumi.StringInput `pulumi:"prefix"`
}

func (LookupRepositoryCreationTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryCreationTemplateArgs)(nil)).Elem()
}

type LookupRepositoryCreationTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryCreationTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryCreationTemplateResult)(nil)).Elem()
}

func (o LookupRepositoryCreationTemplateResultOutput) ToLookupRepositoryCreationTemplateResultOutput() LookupRepositoryCreationTemplateResultOutput {
	return o
}

func (o LookupRepositoryCreationTemplateResultOutput) ToLookupRepositoryCreationTemplateResultOutputWithContext(ctx context.Context) LookupRepositoryCreationTemplateResultOutput {
	return o
}

// A list of enumerable Strings representing the repository creation scenarios that the template will apply towards.
func (o LookupRepositoryCreationTemplateResultOutput) AppliedFor() RepositoryCreationTemplateAppliedForItemArrayOutput {
	return o.ApplyT(func(v LookupRepositoryCreationTemplateResult) []RepositoryCreationTemplateAppliedForItem {
		return v.AppliedFor
	}).(RepositoryCreationTemplateAppliedForItemArrayOutput)
}

// Create timestamp of the template.
func (o LookupRepositoryCreationTemplateResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryCreationTemplateResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The ARN of the role to be assumed by ECR. This role must be in the same account as the registry that you are configuring.
func (o LookupRepositoryCreationTemplateResultOutput) CustomRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryCreationTemplateResult) *string { return v.CustomRoleArn }).(pulumi.StringPtrOutput)
}

// The description of the template.
func (o LookupRepositoryCreationTemplateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryCreationTemplateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The encryption configuration associated with the repository creation template.
func (o LookupRepositoryCreationTemplateResultOutput) EncryptionConfiguration() RepositoryCreationTemplateEncryptionConfigurationPtrOutput {
	return o.ApplyT(func(v LookupRepositoryCreationTemplateResult) *RepositoryCreationTemplateEncryptionConfiguration {
		return v.EncryptionConfiguration
	}).(RepositoryCreationTemplateEncryptionConfigurationPtrOutput)
}

// The image tag mutability setting for the repository.
func (o LookupRepositoryCreationTemplateResultOutput) ImageTagMutability() RepositoryCreationTemplateImageTagMutabilityPtrOutput {
	return o.ApplyT(func(v LookupRepositoryCreationTemplateResult) *RepositoryCreationTemplateImageTagMutability {
		return v.ImageTagMutability
	}).(RepositoryCreationTemplateImageTagMutabilityPtrOutput)
}

// The JSON lifecycle policy text to apply to the repository. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html
func (o LookupRepositoryCreationTemplateResultOutput) LifecyclePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryCreationTemplateResult) *string { return v.LifecyclePolicy }).(pulumi.StringPtrOutput)
}

// The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html
func (o LookupRepositoryCreationTemplateResultOutput) RepositoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryCreationTemplateResult) *string { return v.RepositoryPolicy }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupRepositoryCreationTemplateResultOutput) ResourceTags() RepositoryCreationTemplateTagArrayOutput {
	return o.ApplyT(func(v LookupRepositoryCreationTemplateResult) []RepositoryCreationTemplateTag { return v.ResourceTags }).(RepositoryCreationTemplateTagArrayOutput)
}

// Update timestamp of the template.
func (o LookupRepositoryCreationTemplateResultOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryCreationTemplateResult) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryCreationTemplateResultOutput{})
}
