// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * MatchingWorkflow defined in AWS Entity Resolution service
 */
export function getMatchingWorkflow(args: GetMatchingWorkflowArgs, opts?: pulumi.InvokeOptions): Promise<GetMatchingWorkflowResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:entityresolution:getMatchingWorkflow", {
        "workflowName": args.workflowName,
    }, opts);
}

export interface GetMatchingWorkflowArgs {
    /**
     * The name of the MatchingWorkflow
     */
    workflowName: string;
}

export interface GetMatchingWorkflowResult {
    readonly createdAt?: string;
    /**
     * The description of the MatchingWorkflow
     */
    readonly description?: string;
    readonly inputSourceConfig?: outputs.entityresolution.MatchingWorkflowInputSource[];
    readonly outputSourceConfig?: outputs.entityresolution.MatchingWorkflowOutputSource[];
    readonly resolutionTechniques?: outputs.entityresolution.MatchingWorkflowResolutionTechniques;
    readonly roleArn?: string;
    readonly tags?: outputs.entityresolution.MatchingWorkflowTag[];
    readonly updatedAt?: string;
    readonly workflowArn?: string;
}
/**
 * MatchingWorkflow defined in AWS Entity Resolution service
 */
export function getMatchingWorkflowOutput(args: GetMatchingWorkflowOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMatchingWorkflowResult> {
    return pulumi.output(args).apply((a: any) => getMatchingWorkflow(a, opts))
}

export interface GetMatchingWorkflowOutputArgs {
    /**
     * The name of the MatchingWorkflow
     */
    workflowName: pulumi.Input<string>;
}