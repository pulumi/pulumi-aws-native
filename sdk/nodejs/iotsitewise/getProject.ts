// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTSiteWise::Project
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotsitewise:getProject", {
        "projectId": args.projectId,
    }, opts);
}

export interface GetProjectArgs {
    /**
     * The ID of the project.
     */
    projectId: string;
}

export interface GetProjectResult {
    /**
     * The IDs of the assets to be associated to the project.
     */
    readonly assetIds?: string[];
    /**
     * The ARN of the project.
     */
    readonly projectArn?: string;
    /**
     * A description for the project.
     */
    readonly projectDescription?: string;
    /**
     * The ID of the project.
     */
    readonly projectId?: string;
    /**
     * A friendly name for the project.
     */
    readonly projectName?: string;
    /**
     * A list of key-value pairs that contain metadata for the project.
     */
    readonly tags?: outputs.iotsitewise.ProjectTag[];
}

export function getProjectOutput(args: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply(a => getProject(a, opts))
}

export interface GetProjectOutputArgs {
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
}