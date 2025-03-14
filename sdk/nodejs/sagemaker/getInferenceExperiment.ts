// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::InferenceExperiment
 */
export function getInferenceExperiment(args: GetInferenceExperimentArgs, opts?: pulumi.InvokeOptions): Promise<GetInferenceExperimentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sagemaker:getInferenceExperiment", {
        "name": args.name,
    }, opts);
}

export interface GetInferenceExperimentArgs {
    /**
     * The name for the inference experiment.
     */
    name: string;
}

export interface GetInferenceExperimentResult {
    /**
     * The Amazon Resource Name (ARN) of the inference experiment.
     */
    readonly arn?: string;
    /**
     * The timestamp at which you created the inference experiment.
     */
    readonly creationTime?: string;
    /**
     * The Amazon S3 location and configuration for storing inference request and response data.
     */
    readonly dataStorageConfig?: outputs.sagemaker.InferenceExperimentDataStorageConfig;
    /**
     * The description of the inference experiment.
     */
    readonly description?: string;
    /**
     * The desired state of the experiment after starting or stopping operation.
     */
    readonly desiredState?: enums.sagemaker.InferenceExperimentDesiredState;
    readonly endpointMetadata?: outputs.sagemaker.InferenceExperimentEndpointMetadata;
    /**
     * The timestamp at which you last modified the inference experiment.
     */
    readonly lastModifiedTime?: string;
    /**
     * An array of ModelVariantConfig objects. Each ModelVariantConfig object in the array describes the infrastructure configuration for the corresponding variant.
     */
    readonly modelVariants?: outputs.sagemaker.InferenceExperimentModelVariantConfig[];
    /**
     * The duration for which the inference experiment ran or will run.
     *
     * The maximum duration that you can set for an inference experiment is 30 days.
     */
    readonly schedule?: outputs.sagemaker.InferenceExperimentSchedule;
    /**
     * The configuration of `ShadowMode` inference experiment type, which shows the production variant that takes all the inference requests, and the shadow variant to which Amazon SageMaker replicates a percentage of the inference requests. For the shadow variant it also shows the percentage of requests that Amazon SageMaker replicates.
     */
    readonly shadowModeConfig?: outputs.sagemaker.InferenceExperimentShadowModeConfig;
    /**
     * The status of the inference experiment.
     */
    readonly status?: enums.sagemaker.InferenceExperimentStatus;
    /**
     * The error message or client-specified reason from the StopInferenceExperiment API, that explains the status of the inference experiment.
     */
    readonly statusReason?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::SageMaker::InferenceExperiment
 */
export function getInferenceExperimentOutput(args: GetInferenceExperimentOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInferenceExperimentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:sagemaker:getInferenceExperiment", {
        "name": args.name,
    }, opts);
}

export interface GetInferenceExperimentOutputArgs {
    /**
     * The name for the inference experiment.
     */
    name: pulumi.Input<string>;
}
