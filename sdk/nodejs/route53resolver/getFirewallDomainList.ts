// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53Resolver::FirewallDomainList.
 */
export function getFirewallDomainList(args: GetFirewallDomainListArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallDomainListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:route53resolver:getFirewallDomainList", {
        "id": args.id,
    }, opts);
}

export interface GetFirewallDomainListArgs {
    /**
     * ResourceId
     */
    id: string;
}

export interface GetFirewallDomainListResult {
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
     * Count
     */
    readonly domainCount?: number;
    /**
     * ResourceId
     */
    readonly id?: string;
    /**
     * ServicePrincipal
     */
    readonly managedOwnerName?: string;
    /**
     * Rfc3339TimeString
     */
    readonly modificationTime?: string;
    /**
     * ResolverFirewallDomainList, possible values are COMPLETE, DELETING, UPDATING, COMPLETE_IMPORT_FAILED, IMPORTING, and INACTIVE_OWNER_ACCOUNT_CLOSED.
     */
    readonly status?: enums.route53resolver.FirewallDomainListStatus;
    /**
     * FirewallDomainListAssociationStatus
     */
    readonly statusMessage?: string;
    /**
     * Tags
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource schema for AWS::Route53Resolver::FirewallDomainList.
 */
export function getFirewallDomainListOutput(args: GetFirewallDomainListOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFirewallDomainListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:route53resolver:getFirewallDomainList", {
        "id": args.id,
    }, opts);
}

export interface GetFirewallDomainListOutputArgs {
    /**
     * ResourceId
     */
    id: pulumi.Input<string>;
}
