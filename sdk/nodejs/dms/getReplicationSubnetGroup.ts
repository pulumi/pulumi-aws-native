// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DMS::ReplicationSubnetGroup
 */
export function getReplicationSubnetGroup(args: GetReplicationSubnetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationSubnetGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:dms:getReplicationSubnetGroup", {
        "id": args.id,
    }, opts);
}

export interface GetReplicationSubnetGroupArgs {
    id: string;
}

export interface GetReplicationSubnetGroupResult {
    readonly id?: string;
    readonly replicationSubnetGroupDescription?: string;
    readonly subnetIds?: string[];
    readonly tags?: outputs.dms.ReplicationSubnetGroupTag[];
}

export function getReplicationSubnetGroupOutput(args: GetReplicationSubnetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReplicationSubnetGroupResult> {
    return pulumi.output(args).apply(a => getReplicationSubnetGroup(a, opts))
}

export interface GetReplicationSubnetGroupOutputArgs {
    id: pulumi.Input<string>;
}