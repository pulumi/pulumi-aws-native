// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package entityresolution

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// IdMappingWorkflow defined in AWS Entity Resolution service
func LookupIdMappingWorkflow(ctx *pulumi.Context, args *LookupIdMappingWorkflowArgs, opts ...pulumi.InvokeOption) (*LookupIdMappingWorkflowResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIdMappingWorkflowResult
	err := ctx.Invoke("aws-native:entityresolution:getIdMappingWorkflow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIdMappingWorkflowArgs struct {
	// The name of the IdMappingWorkflow
	WorkflowName string `pulumi:"workflowName"`
}

type LookupIdMappingWorkflowResult struct {
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the IdMappingWorkflow
	Description         *string                               `pulumi:"description"`
	IdMappingTechniques *IdMappingWorkflowIdMappingTechniques `pulumi:"idMappingTechniques"`
	InputSourceConfig   []IdMappingWorkflowInputSource        `pulumi:"inputSourceConfig"`
	OutputSourceConfig  []IdMappingWorkflowOutputSource       `pulumi:"outputSourceConfig"`
	RoleArn             *string                               `pulumi:"roleArn"`
	Tags                []IdMappingWorkflowTag                `pulumi:"tags"`
	UpdatedAt           *string                               `pulumi:"updatedAt"`
	WorkflowArn         *string                               `pulumi:"workflowArn"`
}

func LookupIdMappingWorkflowOutput(ctx *pulumi.Context, args LookupIdMappingWorkflowOutputArgs, opts ...pulumi.InvokeOption) LookupIdMappingWorkflowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIdMappingWorkflowResult, error) {
			args := v.(LookupIdMappingWorkflowArgs)
			r, err := LookupIdMappingWorkflow(ctx, &args, opts...)
			var s LookupIdMappingWorkflowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIdMappingWorkflowResultOutput)
}

type LookupIdMappingWorkflowOutputArgs struct {
	// The name of the IdMappingWorkflow
	WorkflowName pulumi.StringInput `pulumi:"workflowName"`
}

func (LookupIdMappingWorkflowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdMappingWorkflowArgs)(nil)).Elem()
}

type LookupIdMappingWorkflowResultOutput struct{ *pulumi.OutputState }

func (LookupIdMappingWorkflowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdMappingWorkflowResult)(nil)).Elem()
}

func (o LookupIdMappingWorkflowResultOutput) ToLookupIdMappingWorkflowResultOutput() LookupIdMappingWorkflowResultOutput {
	return o
}

func (o LookupIdMappingWorkflowResultOutput) ToLookupIdMappingWorkflowResultOutputWithContext(ctx context.Context) LookupIdMappingWorkflowResultOutput {
	return o
}

func (o LookupIdMappingWorkflowResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIdMappingWorkflowResult] {
	return pulumix.Output[LookupIdMappingWorkflowResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupIdMappingWorkflowResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdMappingWorkflowResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The description of the IdMappingWorkflow
func (o LookupIdMappingWorkflowResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdMappingWorkflowResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupIdMappingWorkflowResultOutput) IdMappingTechniques() IdMappingWorkflowIdMappingTechniquesPtrOutput {
	return o.ApplyT(func(v LookupIdMappingWorkflowResult) *IdMappingWorkflowIdMappingTechniques {
		return v.IdMappingTechniques
	}).(IdMappingWorkflowIdMappingTechniquesPtrOutput)
}

func (o LookupIdMappingWorkflowResultOutput) InputSourceConfig() IdMappingWorkflowInputSourceArrayOutput {
	return o.ApplyT(func(v LookupIdMappingWorkflowResult) []IdMappingWorkflowInputSource { return v.InputSourceConfig }).(IdMappingWorkflowInputSourceArrayOutput)
}

func (o LookupIdMappingWorkflowResultOutput) OutputSourceConfig() IdMappingWorkflowOutputSourceArrayOutput {
	return o.ApplyT(func(v LookupIdMappingWorkflowResult) []IdMappingWorkflowOutputSource { return v.OutputSourceConfig }).(IdMappingWorkflowOutputSourceArrayOutput)
}

func (o LookupIdMappingWorkflowResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdMappingWorkflowResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupIdMappingWorkflowResultOutput) Tags() IdMappingWorkflowTagArrayOutput {
	return o.ApplyT(func(v LookupIdMappingWorkflowResult) []IdMappingWorkflowTag { return v.Tags }).(IdMappingWorkflowTagArrayOutput)
}

func (o LookupIdMappingWorkflowResultOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdMappingWorkflowResult) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupIdMappingWorkflowResultOutput) WorkflowArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdMappingWorkflowResult) *string { return v.WorkflowArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdMappingWorkflowResultOutput{})
}