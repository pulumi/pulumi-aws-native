// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DataPipeline::Pipeline
 */
export function getPipeline(args: GetPipelineArgs, opts?: pulumi.InvokeOptions): Promise<GetPipelineResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:datapipeline:getPipeline", {
        "id": args.id,
    }, opts);
}

export interface GetPipelineArgs {
    id: string;
}

export interface GetPipelineResult {
    readonly activate?: boolean;
    readonly id?: string;
    readonly parameterObjects?: outputs.datapipeline.PipelineParameterObject[];
    readonly parameterValues?: outputs.datapipeline.PipelineParameterValue[];
    readonly pipelineObjects?: outputs.datapipeline.PipelineObject[];
    readonly pipelineTags?: outputs.datapipeline.PipelineTag[];
}

export function getPipelineOutput(args: GetPipelineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPipelineResult> {
    return pulumi.output(args).apply(a => getPipeline(a, opts))
}

export interface GetPipelineOutputArgs {
    id: pulumi.Input<string>;
}