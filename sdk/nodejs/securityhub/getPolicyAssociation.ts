// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::SecurityHub::PolicyAssociation resource represents the AWS Security Hub Central Configuration Policy associations in your Target. Only the AWS Security Hub delegated administrator can create the resouce from the home region.
 */
export function getPolicyAssociation(args: GetPolicyAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyAssociationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:securityhub:getPolicyAssociation", {
        "associationIdentifier": args.associationIdentifier,
    }, opts);
}

export interface GetPolicyAssociationArgs {
    /**
     * A unique identifier to indicates if the target has an association
     */
    associationIdentifier: string;
}

export interface GetPolicyAssociationResult {
    /**
     * A unique identifier to indicates if the target has an association
     */
    readonly associationIdentifier?: string;
    /**
     * The current status of the association between the specified target and the configuration
     */
    readonly associationStatus?: enums.securityhub.PolicyAssociationAssociationStatus;
    /**
     * An explanation for a FAILED value for AssociationStatus
     */
    readonly associationStatusMessage?: string;
    /**
     * Indicates whether the association between the specified target and the configuration was directly applied by the Security Hub delegated administrator or inherited from a parent
     */
    readonly associationType?: enums.securityhub.PolicyAssociationAssociationType;
    /**
     * The universally unique identifier (UUID) of the configuration policy or a value of SELF_MANAGED_SECURITY_HUB for a self-managed configuration
     */
    readonly configurationPolicyId?: string;
    /**
     * The date and time, in UTC and ISO 8601 format, that the configuration policy association was last updated
     */
    readonly updatedAt?: string;
}
/**
 * The AWS::SecurityHub::PolicyAssociation resource represents the AWS Security Hub Central Configuration Policy associations in your Target. Only the AWS Security Hub delegated administrator can create the resouce from the home region.
 */
export function getPolicyAssociationOutput(args: GetPolicyAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyAssociationResult> {
    return pulumi.output(args).apply((a: any) => getPolicyAssociation(a, opts))
}

export interface GetPolicyAssociationOutputArgs {
    /**
     * A unique identifier to indicates if the target has an association
     */
    associationIdentifier: pulumi.Input<string>;
}