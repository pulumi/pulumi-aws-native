// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53Resolver::FirewallRuleGroup.
 */
export function getFirewallRuleGroup(args: GetFirewallRuleGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallRuleGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53resolver:getFirewallRuleGroup", {
        "id": args.id,
    }, opts);
}

export interface GetFirewallRuleGroupArgs {
    /**
     * ResourceId
     */
    id: string;
}

export interface GetFirewallRuleGroupResult {
    /**
     * Arn
     */
    readonly arn?: string;
    /**
     * Rfc3339TimeString
     */
    readonly creationTime?: string;
    /**
     * The id of the creator request.
     */
    readonly creatorRequestId?: string;
    /**
     * FirewallRules
     */
    readonly firewallRules?: outputs.route53resolver.FirewallRuleGroupFirewallRule[];
    /**
     * ResourceId
     */
    readonly id?: string;
    /**
     * Rfc3339TimeString
     */
    readonly modificationTime?: string;
    /**
     * AccountId
     */
    readonly ownerId?: string;
    /**
     * Count
     */
    readonly ruleCount?: number;
    /**
     * ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.
     */
    readonly shareStatus?: enums.route53resolver.FirewallRuleGroupShareStatus;
    /**
     * ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.
     */
    readonly status?: enums.route53resolver.FirewallRuleGroupStatus;
    /**
     * FirewallRuleGroupStatus
     */
    readonly statusMessage?: string;
    /**
     * Tags
     */
    readonly tags?: outputs.route53resolver.FirewallRuleGroupTag[];
}

export function getFirewallRuleGroupOutput(args: GetFirewallRuleGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallRuleGroupResult> {
    return pulumi.output(args).apply(a => getFirewallRuleGroup(a, opts))
}

export interface GetFirewallRuleGroupOutputArgs {
    /**
     * ResourceId
     */
    id: pulumi.Input<string>;
}