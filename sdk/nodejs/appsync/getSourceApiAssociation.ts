// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppSync::SourceApiAssociation
 */
export function getSourceApiAssociation(args: GetSourceApiAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceApiAssociationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:appsync:getSourceApiAssociation", {
        "associationArn": args.associationArn,
    }, opts);
}

export interface GetSourceApiAssociationArgs {
    /**
     * ARN of the SourceApiAssociation.
     */
    associationArn: string;
}

export interface GetSourceApiAssociationResult {
    /**
     * ARN of the SourceApiAssociation.
     */
    readonly associationArn?: string;
    /**
     * Id of the SourceApiAssociation.
     */
    readonly associationId?: string;
    /**
     * Description of the SourceApiAssociation.
     */
    readonly description?: string;
    /**
     * Date of last schema successful merge.
     */
    readonly lastSuccessfulMergeDate?: string;
    /**
     * ARN of the Merged API in the association.
     */
    readonly mergedApiArn?: string;
    /**
     * GraphQLApiId of the Merged API in the association.
     */
    readonly mergedApiId?: string;
    /**
     * ARN of the source API in the association.
     */
    readonly sourceApiArn?: string;
    /**
     * Customized configuration for SourceApiAssociation.
     */
    readonly sourceApiAssociationConfig?: outputs.appsync.SourceApiAssociationConfig;
    /**
     * Current status of SourceApiAssociation.
     */
    readonly sourceApiAssociationStatus?: enums.appsync.SourceApiAssociationStatus;
    /**
     * Current SourceApiAssociation status details.
     */
    readonly sourceApiAssociationStatusDetail?: string;
    /**
     * GraphQLApiId of the source API in the association.
     */
    readonly sourceApiId?: string;
}
/**
 * Resource Type definition for AWS::AppSync::SourceApiAssociation
 */
export function getSourceApiAssociationOutput(args: GetSourceApiAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceApiAssociationResult> {
    return pulumi.output(args).apply((a: any) => getSourceApiAssociation(a, opts))
}

export interface GetSourceApiAssociationOutputArgs {
    /**
     * ARN of the SourceApiAssociation.
     */
    associationArn: pulumi.Input<string>;
}