// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WorkSpaces::Workspace
 */
export function getWorkspace(args: GetWorkspaceArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:workspaces:getWorkspace", {
        "id": args.id,
    }, opts);
}

export interface GetWorkspaceArgs {
    id: string;
}

export interface GetWorkspaceResult {
    readonly bundleId?: string;
    readonly directoryId?: string;
    readonly id?: string;
    readonly rootVolumeEncryptionEnabled?: boolean;
    readonly tags?: outputs.workspaces.WorkspaceTag[];
    readonly userVolumeEncryptionEnabled?: boolean;
    readonly volumeEncryptionKey?: string;
    readonly workspaceProperties?: outputs.workspaces.WorkspaceProperties;
}

export function getWorkspaceOutput(args: GetWorkspaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkspaceResult> {
    return pulumi.output(args).apply(a => getWorkspace(a, opts))
}

export interface GetWorkspaceOutputArgs {
    id: pulumi.Input<string>;
}