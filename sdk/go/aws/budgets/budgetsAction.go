// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package budgets

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
type BudgetsAction struct {
	pulumi.CustomResourceState

	ActionId         pulumi.StringOutput                 `pulumi:"actionId"`
	ActionThreshold  BudgetsActionActionThresholdOutput  `pulumi:"actionThreshold"`
	ActionType       BudgetsActionActionTypeOutput       `pulumi:"actionType"`
	ApprovalModel    BudgetsActionApprovalModelPtrOutput `pulumi:"approvalModel"`
	BudgetName       pulumi.StringOutput                 `pulumi:"budgetName"`
	Definition       BudgetsActionDefinitionOutput       `pulumi:"definition"`
	ExecutionRoleArn pulumi.StringOutput                 `pulumi:"executionRoleArn"`
	NotificationType BudgetsActionNotificationTypeOutput `pulumi:"notificationType"`
	Subscribers      BudgetsActionSubscriberArrayOutput  `pulumi:"subscribers"`
}

// NewBudgetsAction registers a new resource with the given unique name, arguments, and options.
func NewBudgetsAction(ctx *pulumi.Context,
	name string, args *BudgetsActionArgs, opts ...pulumi.ResourceOption) (*BudgetsAction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionThreshold == nil {
		return nil, errors.New("invalid value for required argument 'ActionThreshold'")
	}
	if args.ActionType == nil {
		return nil, errors.New("invalid value for required argument 'ActionType'")
	}
	if args.BudgetName == nil {
		return nil, errors.New("invalid value for required argument 'BudgetName'")
	}
	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.ExecutionRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'ExecutionRoleArn'")
	}
	if args.NotificationType == nil {
		return nil, errors.New("invalid value for required argument 'NotificationType'")
	}
	if args.Subscribers == nil {
		return nil, errors.New("invalid value for required argument 'Subscribers'")
	}
	var resource BudgetsAction
	err := ctx.RegisterResource("aws-native:budgets:BudgetsAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudgetsAction gets an existing BudgetsAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudgetsAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetsActionState, opts ...pulumi.ResourceOption) (*BudgetsAction, error) {
	var resource BudgetsAction
	err := ctx.ReadResource("aws-native:budgets:BudgetsAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BudgetsAction resources.
type budgetsActionState struct {
}

type BudgetsActionState struct {
}

func (BudgetsActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetsActionState)(nil)).Elem()
}

type budgetsActionArgs struct {
	ActionThreshold  BudgetsActionActionThreshold  `pulumi:"actionThreshold"`
	ActionType       BudgetsActionActionType       `pulumi:"actionType"`
	ApprovalModel    *BudgetsActionApprovalModel   `pulumi:"approvalModel"`
	BudgetName       string                        `pulumi:"budgetName"`
	Definition       BudgetsActionDefinition       `pulumi:"definition"`
	ExecutionRoleArn string                        `pulumi:"executionRoleArn"`
	NotificationType BudgetsActionNotificationType `pulumi:"notificationType"`
	Subscribers      []BudgetsActionSubscriber     `pulumi:"subscribers"`
}

// The set of arguments for constructing a BudgetsAction resource.
type BudgetsActionArgs struct {
	ActionThreshold  BudgetsActionActionThresholdInput
	ActionType       BudgetsActionActionTypeInput
	ApprovalModel    BudgetsActionApprovalModelPtrInput
	BudgetName       pulumi.StringInput
	Definition       BudgetsActionDefinitionInput
	ExecutionRoleArn pulumi.StringInput
	NotificationType BudgetsActionNotificationTypeInput
	Subscribers      BudgetsActionSubscriberArrayInput
}

func (BudgetsActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetsActionArgs)(nil)).Elem()
}

type BudgetsActionInput interface {
	pulumi.Input

	ToBudgetsActionOutput() BudgetsActionOutput
	ToBudgetsActionOutputWithContext(ctx context.Context) BudgetsActionOutput
}

func (*BudgetsAction) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetsAction)(nil)).Elem()
}

func (i *BudgetsAction) ToBudgetsActionOutput() BudgetsActionOutput {
	return i.ToBudgetsActionOutputWithContext(context.Background())
}

func (i *BudgetsAction) ToBudgetsActionOutputWithContext(ctx context.Context) BudgetsActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetsActionOutput)
}

type BudgetsActionOutput struct{ *pulumi.OutputState }

func (BudgetsActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetsAction)(nil)).Elem()
}

func (o BudgetsActionOutput) ToBudgetsActionOutput() BudgetsActionOutput {
	return o
}

func (o BudgetsActionOutput) ToBudgetsActionOutputWithContext(ctx context.Context) BudgetsActionOutput {
	return o
}

func (o BudgetsActionOutput) ActionId() pulumi.StringOutput {
	return o.ApplyT(func(v *BudgetsAction) pulumi.StringOutput { return v.ActionId }).(pulumi.StringOutput)
}

func (o BudgetsActionOutput) ActionThreshold() BudgetsActionActionThresholdOutput {
	return o.ApplyT(func(v *BudgetsAction) BudgetsActionActionThresholdOutput { return v.ActionThreshold }).(BudgetsActionActionThresholdOutput)
}

func (o BudgetsActionOutput) ActionType() BudgetsActionActionTypeOutput {
	return o.ApplyT(func(v *BudgetsAction) BudgetsActionActionTypeOutput { return v.ActionType }).(BudgetsActionActionTypeOutput)
}

func (o BudgetsActionOutput) ApprovalModel() BudgetsActionApprovalModelPtrOutput {
	return o.ApplyT(func(v *BudgetsAction) BudgetsActionApprovalModelPtrOutput { return v.ApprovalModel }).(BudgetsActionApprovalModelPtrOutput)
}

func (o BudgetsActionOutput) BudgetName() pulumi.StringOutput {
	return o.ApplyT(func(v *BudgetsAction) pulumi.StringOutput { return v.BudgetName }).(pulumi.StringOutput)
}

func (o BudgetsActionOutput) Definition() BudgetsActionDefinitionOutput {
	return o.ApplyT(func(v *BudgetsAction) BudgetsActionDefinitionOutput { return v.Definition }).(BudgetsActionDefinitionOutput)
}

func (o BudgetsActionOutput) ExecutionRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *BudgetsAction) pulumi.StringOutput { return v.ExecutionRoleArn }).(pulumi.StringOutput)
}

func (o BudgetsActionOutput) NotificationType() BudgetsActionNotificationTypeOutput {
	return o.ApplyT(func(v *BudgetsAction) BudgetsActionNotificationTypeOutput { return v.NotificationType }).(BudgetsActionNotificationTypeOutput)
}

func (o BudgetsActionOutput) Subscribers() BudgetsActionSubscriberArrayOutput {
	return o.ApplyT(func(v *BudgetsAction) BudgetsActionSubscriberArrayOutput { return v.Subscribers }).(BudgetsActionSubscriberArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetsActionInput)(nil)).Elem(), &BudgetsAction{})
	pulumi.RegisterOutputType(BudgetsActionOutput{})
}