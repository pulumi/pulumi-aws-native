// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Transfer::Workflow
 */
export function getWorkflow(args: GetWorkflowArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkflowResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:transfer:getWorkflow", {
        "workflowId": args.workflowId,
    }, opts);
}

export interface GetWorkflowArgs {
    /**
     * A unique identifier for the workflow.
     */
    workflowId: string;
}

export interface GetWorkflowResult {
    /**
     * Specifies the unique Amazon Resource Name (ARN) for the workflow.
     */
    readonly arn?: string;
    /**
     * Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.
     */
    readonly tags?: outputs.Tag[];
    /**
     * A unique identifier for the workflow.
     */
    readonly workflowId?: string;
}
/**
 * Resource Type definition for AWS::Transfer::Workflow
 */
export function getWorkflowOutput(args: GetWorkflowOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWorkflowResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:transfer:getWorkflow", {
        "workflowId": args.workflowId,
    }, opts);
}

export interface GetWorkflowOutputArgs {
    /**
     * A unique identifier for the workflow.
     */
    workflowId: pulumi.Input<string>;
}
