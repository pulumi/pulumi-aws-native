// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package budgets

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
func LookupBudgetsAction(ctx *pulumi.Context, args *LookupBudgetsActionArgs, opts ...pulumi.InvokeOption) (*LookupBudgetsActionResult, error) {
	var rv LookupBudgetsActionResult
	err := ctx.Invoke("aws-native:budgets:getBudgetsAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBudgetsActionArgs struct {
	ActionId   string `pulumi:"actionId"`
	BudgetName string `pulumi:"budgetName"`
}

type LookupBudgetsActionResult struct {
	ActionId         *string                        `pulumi:"actionId"`
	ActionThreshold  *BudgetsActionActionThreshold  `pulumi:"actionThreshold"`
	ApprovalModel    *BudgetsActionApprovalModel    `pulumi:"approvalModel"`
	Definition       *BudgetsActionDefinition       `pulumi:"definition"`
	ExecutionRoleArn *string                        `pulumi:"executionRoleArn"`
	NotificationType *BudgetsActionNotificationType `pulumi:"notificationType"`
	Subscribers      []BudgetsActionSubscriber      `pulumi:"subscribers"`
}

func LookupBudgetsActionOutput(ctx *pulumi.Context, args LookupBudgetsActionOutputArgs, opts ...pulumi.InvokeOption) LookupBudgetsActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBudgetsActionResult, error) {
			args := v.(LookupBudgetsActionArgs)
			r, err := LookupBudgetsAction(ctx, &args, opts...)
			var s LookupBudgetsActionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBudgetsActionResultOutput)
}

type LookupBudgetsActionOutputArgs struct {
	ActionId   pulumi.StringInput `pulumi:"actionId"`
	BudgetName pulumi.StringInput `pulumi:"budgetName"`
}

func (LookupBudgetsActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetsActionArgs)(nil)).Elem()
}

type LookupBudgetsActionResultOutput struct{ *pulumi.OutputState }

func (LookupBudgetsActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetsActionResult)(nil)).Elem()
}

func (o LookupBudgetsActionResultOutput) ToLookupBudgetsActionResultOutput() LookupBudgetsActionResultOutput {
	return o
}

func (o LookupBudgetsActionResultOutput) ToLookupBudgetsActionResultOutputWithContext(ctx context.Context) LookupBudgetsActionResultOutput {
	return o
}

func (o LookupBudgetsActionResultOutput) ActionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBudgetsActionResult) *string { return v.ActionId }).(pulumi.StringPtrOutput)
}

func (o LookupBudgetsActionResultOutput) ActionThreshold() BudgetsActionActionThresholdPtrOutput {
	return o.ApplyT(func(v LookupBudgetsActionResult) *BudgetsActionActionThreshold { return v.ActionThreshold }).(BudgetsActionActionThresholdPtrOutput)
}

func (o LookupBudgetsActionResultOutput) ApprovalModel() BudgetsActionApprovalModelPtrOutput {
	return o.ApplyT(func(v LookupBudgetsActionResult) *BudgetsActionApprovalModel { return v.ApprovalModel }).(BudgetsActionApprovalModelPtrOutput)
}

func (o LookupBudgetsActionResultOutput) Definition() BudgetsActionDefinitionPtrOutput {
	return o.ApplyT(func(v LookupBudgetsActionResult) *BudgetsActionDefinition { return v.Definition }).(BudgetsActionDefinitionPtrOutput)
}

func (o LookupBudgetsActionResultOutput) ExecutionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBudgetsActionResult) *string { return v.ExecutionRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupBudgetsActionResultOutput) NotificationType() BudgetsActionNotificationTypePtrOutput {
	return o.ApplyT(func(v LookupBudgetsActionResult) *BudgetsActionNotificationType { return v.NotificationType }).(BudgetsActionNotificationTypePtrOutput)
}

func (o LookupBudgetsActionResultOutput) Subscribers() BudgetsActionSubscriberArrayOutput {
	return o.ApplyT(func(v LookupBudgetsActionResult) []BudgetsActionSubscriber { return v.Subscribers }).(BudgetsActionSubscriberArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBudgetsActionResultOutput{})
}