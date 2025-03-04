// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::Endpoint
 */
export function getEndpoint(args: GetEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sagemaker:getEndpoint", {
        "endpointArn": args.endpointArn,
    }, opts);
}

export interface GetEndpointArgs {
    /**
     * The Amazon Resource Name (ARN) of the endpoint.
     */
    endpointArn: string;
}

export interface GetEndpointResult {
    /**
     * Specifies deployment configuration for updating the SageMaker endpoint. Includes rollback and update policies.
     */
    readonly deploymentConfig?: outputs.sagemaker.EndpointDeploymentConfig;
    /**
     * The Amazon Resource Name (ARN) of the endpoint.
     */
    readonly endpointArn?: string;
    /**
     * The name of the endpoint configuration for the SageMaker endpoint. This is a required property.
     */
    readonly endpointConfigName?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::SageMaker::Endpoint
 */
export function getEndpointOutput(args: GetEndpointOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEndpointResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:sagemaker:getEndpoint", {
        "endpointArn": args.endpointArn,
    }, opts);
}

export interface GetEndpointOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the endpoint.
     */
    endpointArn: pulumi.Input<string>;
}
