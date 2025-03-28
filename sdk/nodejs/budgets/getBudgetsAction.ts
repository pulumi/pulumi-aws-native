// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export function getBudgetsAction(args: GetBudgetsActionArgs, opts?: pulumi.InvokeOptions): Promise<GetBudgetsActionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:budgets:getBudgetsAction", {
        "actionId": args.actionId,
        "budgetName": args.budgetName,
    }, opts);
}

export interface GetBudgetsActionArgs {
    /**
     * A system-generated universally unique identifier (UUID) for the action.
     */
    actionId: string;
    /**
     * A string that represents the budget name. ":" and "\" characters aren't allowed.
     */
    budgetName: string;
}

export interface GetBudgetsActionResult {
    /**
     * A system-generated universally unique identifier (UUID) for the action.
     */
    readonly actionId?: string;
    /**
     * The trigger threshold of the action.
     */
    readonly actionThreshold?: outputs.budgets.BudgetsActionActionThreshold;
    /**
     * This specifies if the action needs manual or automatic approval.
     */
    readonly approvalModel?: enums.budgets.BudgetsActionApprovalModel;
    /**
     * Specifies all of the type-specific parameters.
     */
    readonly definition?: outputs.budgets.BudgetsActionDefinition;
    /**
     * The role passed for action execution and reversion. Roles and actions must be in the same account.
     */
    readonly executionRoleArn?: string;
    /**
     * The type of a notification.
     */
    readonly notificationType?: enums.budgets.BudgetsActionNotificationType;
    /**
     * An optional list of tags to associate with the specified budget action. Each tag consists of a key and a value, and each key must be unique for the resource.
     */
    readonly resourceTags?: outputs.Tag[];
    /**
     * A list of subscribers.
     */
    readonly subscribers?: outputs.budgets.BudgetsActionSubscriber[];
}
/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export function getBudgetsActionOutput(args: GetBudgetsActionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBudgetsActionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:budgets:getBudgetsAction", {
        "actionId": args.actionId,
        "budgetName": args.budgetName,
    }, opts);
}

export interface GetBudgetsActionOutputArgs {
    /**
     * A system-generated universally unique identifier (UUID) for the action.
     */
    actionId: pulumi.Input<string>;
    /**
     * A string that represents the budget name. ":" and "\" characters aren't allowed.
     */
    budgetName: pulumi.Input<string>;
}
