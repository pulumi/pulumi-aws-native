// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The ``AWS::ApiGateway::Deployment`` resource deploys an API Gateway ``RestApi`` resource to a stage so that clients can call the API over the internet. The stage acts as an environment.
 */
export function getDeployment(args: GetDeploymentArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:apigateway:getDeployment", {
        "deploymentId": args.deploymentId,
        "restApiId": args.restApiId,
    }, opts);
}

export interface GetDeploymentArgs {
    /**
     * The ID for the deployment. For example: `abc123` .
     */
    deploymentId: string;
    /**
     * The string identifier of the associated RestApi.
     */
    restApiId: string;
}

export interface GetDeploymentResult {
    /**
     * The ID for the deployment. For example: `abc123` .
     */
    readonly deploymentId?: string;
    /**
     * The description for the Deployment resource to create.
     */
    readonly description?: string;
}
/**
 * The ``AWS::ApiGateway::Deployment`` resource deploys an API Gateway ``RestApi`` resource to a stage so that clients can call the API over the internet. The stage acts as an environment.
 */
export function getDeploymentOutput(args: GetDeploymentOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDeploymentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:apigateway:getDeployment", {
        "deploymentId": args.deploymentId,
        "restApiId": args.restApiId,
    }, opts);
}

export interface GetDeploymentOutputArgs {
    /**
     * The ID for the deployment. For example: `abc123` .
     */
    deploymentId: pulumi.Input<string>;
    /**
     * The string identifier of the associated RestApi.
     */
    restApiId: pulumi.Input<string>;
}
