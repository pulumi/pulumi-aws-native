// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTTwinMaker::SyncJob
 */
export function getSyncJob(args: GetSyncJobArgs, opts?: pulumi.InvokeOptions): Promise<GetSyncJobResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iottwinmaker:getSyncJob", {
        "syncSource": args.syncSource,
        "workspaceId": args.workspaceId,
    }, opts);
}

export interface GetSyncJobArgs {
    /**
     * The source of the SyncJob.
     */
    syncSource: string;
    /**
     * The ID of the workspace.
     */
    workspaceId: string;
}

export interface GetSyncJobResult {
    /**
     * The ARN of the SyncJob.
     */
    readonly arn?: string;
    /**
     * The date and time when the sync job was created.
     */
    readonly creationDateTime?: string;
    /**
     * The state of SyncJob.
     */
    readonly state?: string;
    /**
     * The date and time when the sync job was updated.
     */
    readonly updateDateTime?: string;
}

export function getSyncJobOutput(args: GetSyncJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSyncJobResult> {
    return pulumi.output(args).apply(a => getSyncJob(a, opts))
}

export interface GetSyncJobOutputArgs {
    /**
     * The source of the SyncJob.
     */
    syncSource: pulumi.Input<string>;
    /**
     * The ID of the workspace.
     */
    workspaceId: pulumi.Input<string>;
}