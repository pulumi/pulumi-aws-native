// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::SES::MailManagerRuleSet Resource Type
 */
export function getMailManagerRuleSet(args: GetMailManagerRuleSetArgs, opts?: pulumi.InvokeOptions): Promise<GetMailManagerRuleSetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ses:getMailManagerRuleSet", {
        "ruleSetId": args.ruleSetId,
    }, opts);
}

export interface GetMailManagerRuleSetArgs {
    /**
     * The identifier of the rule set.
     */
    ruleSetId: string;
}

export interface GetMailManagerRuleSetResult {
    /**
     * The Amazon Resource Name (ARN) of the rule set resource.
     */
    readonly ruleSetArn?: string;
    /**
     * The identifier of the rule set.
     */
    readonly ruleSetId?: string;
    /**
     * A user-friendly name for the rule set.
     */
    readonly ruleSetName?: string;
    /**
     * Conditional rules that are evaluated for determining actions on email.
     */
    readonly rules?: outputs.ses.MailManagerRuleSetRule[];
    /**
     * The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Definition of AWS::SES::MailManagerRuleSet Resource Type
 */
export function getMailManagerRuleSetOutput(args: GetMailManagerRuleSetOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMailManagerRuleSetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ses:getMailManagerRuleSet", {
        "ruleSetId": args.ruleSetId,
    }, opts);
}

export interface GetMailManagerRuleSetOutputArgs {
    /**
     * The identifier of the rule set.
     */
    ruleSetId: pulumi.Input<string>;
}
