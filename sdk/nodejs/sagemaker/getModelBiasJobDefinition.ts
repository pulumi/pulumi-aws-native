// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::ModelBiasJobDefinition
 */
export function getModelBiasJobDefinition(args: GetModelBiasJobDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetModelBiasJobDefinitionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sagemaker:getModelBiasJobDefinition", {
        "jobDefinitionArn": args.jobDefinitionArn,
    }, opts);
}

export interface GetModelBiasJobDefinitionArgs {
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    jobDefinitionArn: string;
}

export interface GetModelBiasJobDefinitionResult {
    /**
     * The time at which the job definition was created.
     */
    readonly creationTime?: string;
    readonly endpointName?: string;
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    readonly jobDefinitionArn?: string;
}

export function getModelBiasJobDefinitionOutput(args: GetModelBiasJobDefinitionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModelBiasJobDefinitionResult> {
    return pulumi.output(args).apply(a => getModelBiasJobDefinition(a, opts))
}

export interface GetModelBiasJobDefinitionOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    jobDefinitionArn: pulumi.Input<string>;
}