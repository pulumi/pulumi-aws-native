// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS Route53 Recovery Control basic constructs and validation rules.
 */
export function getSafetyRule(args: GetSafetyRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetSafetyRuleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53recoverycontrol:getSafetyRule", {
        "safetyRuleArn": args.safetyRuleArn,
    }, opts);
}

export interface GetSafetyRuleArgs {
    /**
     * The Amazon Resource Name (ARN) of the safety rule.
     */
    safetyRuleArn: string;
}

export interface GetSafetyRuleResult {
    readonly assertionRule?: outputs.route53recoverycontrol.SafetyRuleAssertionRule;
    readonly gatingRule?: outputs.route53recoverycontrol.SafetyRuleGatingRule;
    readonly name?: string;
    /**
     * The Amazon Resource Name (ARN) of the safety rule.
     */
    readonly safetyRuleArn?: string;
    /**
     * The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
     */
    readonly status?: enums.route53recoverycontrol.SafetyRuleStatus;
}

export function getSafetyRuleOutput(args: GetSafetyRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSafetyRuleResult> {
    return pulumi.output(args).apply(a => getSafetyRule(a, opts))
}

export interface GetSafetyRuleOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the safety rule.
     */
    safetyRuleArn: pulumi.Input<string>;
}