// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package omics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Omics::Workflow Resource Type
func LookupWorkflow(ctx *pulumi.Context, args *LookupWorkflowArgs, opts ...pulumi.InvokeOption) (*LookupWorkflowResult, error) {
	var rv LookupWorkflowResult
	err := ctx.Invoke("aws-native:omics:getWorkflow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkflowArgs struct {
	Id string `pulumi:"id"`
}

type LookupWorkflowResult struct {
	Arn          *string         `pulumi:"arn"`
	CreationTime *string         `pulumi:"creationTime"`
	Description  *string         `pulumi:"description"`
	Id           *string         `pulumi:"id"`
	Name         *string         `pulumi:"name"`
	Status       *WorkflowStatus `pulumi:"status"`
	Tags         *WorkflowTagMap `pulumi:"tags"`
	Type         *WorkflowType   `pulumi:"type"`
}

func LookupWorkflowOutput(ctx *pulumi.Context, args LookupWorkflowOutputArgs, opts ...pulumi.InvokeOption) LookupWorkflowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkflowResult, error) {
			args := v.(LookupWorkflowArgs)
			r, err := LookupWorkflow(ctx, &args, opts...)
			var s LookupWorkflowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkflowResultOutput)
}

type LookupWorkflowOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupWorkflowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowArgs)(nil)).Elem()
}

type LookupWorkflowResultOutput struct{ *pulumi.OutputState }

func (LookupWorkflowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowResult)(nil)).Elem()
}

func (o LookupWorkflowResultOutput) ToLookupWorkflowResultOutput() LookupWorkflowResultOutput {
	return o
}

func (o LookupWorkflowResultOutput) ToLookupWorkflowResultOutputWithContext(ctx context.Context) LookupWorkflowResultOutput {
	return o
}

func (o LookupWorkflowResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Status() WorkflowStatusPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *WorkflowStatus { return v.Status }).(WorkflowStatusPtrOutput)
}

func (o LookupWorkflowResultOutput) Tags() WorkflowTagMapPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *WorkflowTagMap { return v.Tags }).(WorkflowTagMapPtrOutput)
}

func (o LookupWorkflowResultOutput) Type() WorkflowTypePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *WorkflowType { return v.Type }).(WorkflowTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkflowResultOutput{})
}