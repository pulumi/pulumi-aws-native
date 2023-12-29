// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::SnapshotBlockPublicAccess
 */
export function getSnapshotBlockPublicAccess(args: GetSnapshotBlockPublicAccessArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotBlockPublicAccessResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getSnapshotBlockPublicAccess", {
        "accountId": args.accountId,
    }, opts);
}

export interface GetSnapshotBlockPublicAccessArgs {
    /**
     * The identifier for the specified AWS account.
     */
    accountId: string;
}

export interface GetSnapshotBlockPublicAccessResult {
    /**
     * The identifier for the specified AWS account.
     */
    readonly accountId?: string;
    /**
     * The state of EBS Snapshot Block Public Access.
     */
    readonly state?: enums.ec2.SnapshotBlockPublicAccessState;
}
/**
 * Resource Type definition for AWS::EC2::SnapshotBlockPublicAccess
 */
export function getSnapshotBlockPublicAccessOutput(args: GetSnapshotBlockPublicAccessOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSnapshotBlockPublicAccessResult> {
    return pulumi.output(args).apply((a: any) => getSnapshotBlockPublicAccess(a, opts))
}

export interface GetSnapshotBlockPublicAccessOutputArgs {
    /**
     * The identifier for the specified AWS account.
     */
    accountId: pulumi.Input<string>;
}