// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::ModelBiasJobDefinition
 */
export function getModelBiasJobDefinition(args: GetModelBiasJobDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetModelBiasJobDefinitionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
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
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    readonly jobDefinitionArn?: string;
}
/**
 * Resource Type definition for AWS::SageMaker::ModelBiasJobDefinition
 */
export function getModelBiasJobDefinitionOutput(args: GetModelBiasJobDefinitionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetModelBiasJobDefinitionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:sagemaker:getModelBiasJobDefinition", {
        "jobDefinitionArn": args.jobDefinitionArn,
    }, opts);
}

export interface GetModelBiasJobDefinitionOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    jobDefinitionArn: pulumi.Input<string>;
}
