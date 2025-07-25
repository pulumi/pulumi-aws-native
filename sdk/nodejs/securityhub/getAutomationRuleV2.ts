// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::SecurityHub::AutomationRuleV2
 */
export function getAutomationRuleV2(args: GetAutomationRuleV2Args, opts?: pulumi.InvokeOptions): Promise<GetAutomationRuleV2Result> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:securityhub:getAutomationRuleV2", {
        "ruleArn": args.ruleArn,
    }, opts);
}

export interface GetAutomationRuleV2Args {
    /**
     * The ARN of the automation rule
     */
    ruleArn: string;
}

export interface GetAutomationRuleV2Result {
    /**
     * A list of actions to be performed when the rule criteria is met
     */
    readonly actions?: outputs.securityhub.AutomationRuleV2AutomationRulesActionV2[];
    /**
     * The timestamp when the V2 automation rule was created.
     */
    readonly createdAt?: string;
    /**
     * The filtering type and configuration of the automation rule.
     */
    readonly criteria?: outputs.securityhub.AutomationRuleV2Criteria;
    /**
     * A description of the automation rule
     */
    readonly description?: string;
    /**
     * The ARN of the automation rule
     */
    readonly ruleArn?: string;
    /**
     * The ID of the automation rule
     */
    readonly ruleId?: string;
    /**
     * The name of the automation rule
     */
    readonly ruleName?: string;
    /**
     * The value for the rule priority
     */
    readonly ruleOrder?: number;
    /**
     * The status of the automation rule
     */
    readonly ruleStatus?: enums.securityhub.AutomationRuleV2RuleStatus;
    /**
     * A list of key-value pairs associated with the V2 automation rule.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The timestamp when the V2 automation rule was updated.
     */
    readonly updatedAt?: string;
}
/**
 * Resource schema for AWS::SecurityHub::AutomationRuleV2
 */
export function getAutomationRuleV2Output(args: GetAutomationRuleV2OutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAutomationRuleV2Result> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:securityhub:getAutomationRuleV2", {
        "ruleArn": args.ruleArn,
    }, opts);
}

export interface GetAutomationRuleV2OutputArgs {
    /**
     * The ARN of the automation rule
     */
    ruleArn: pulumi.Input<string>;
}
