// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource schema for EC2 EIP association.
 */
export function getEipAssociation(args: GetEipAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetEipAssociationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getEipAssociation", {
        "id": args.id,
    }, opts);
}

export interface GetEipAssociationArgs {
    /**
     * Composite ID of non-empty properties, to determine the identification.
     */
    id: string;
}

export interface GetEipAssociationResult {
    /**
     * Composite ID of non-empty properties, to determine the identification.
     */
    readonly id?: string;
}
/**
 * Resource schema for EC2 EIP association.
 */
export function getEipAssociationOutput(args: GetEipAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEipAssociationResult> {
    return pulumi.output(args).apply((a: any) => getEipAssociation(a, opts))
}

export interface GetEipAssociationOutputArgs {
    /**
     * Composite ID of non-empty properties, to determine the identification.
     */
    id: pulumi.Input<string>;
}