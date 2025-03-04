// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SecretsManager::ResourcePolicy
 */
export function getResourcePolicy(args: GetResourcePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcePolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:secretsmanager:getResourcePolicy", {
        "id": args.id,
    }, opts);
}

export interface GetResourcePolicyArgs {
    /**
     * The Arn of the secret.
     */
    id: string;
}

export interface GetResourcePolicyResult {
    /**
     * The Arn of the secret.
     */
    readonly id?: string;
    /**
     * A JSON-formatted string for an AWS resource-based policy.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SecretsManager::ResourcePolicy` for more information about the expected schema for this property.
     */
    readonly resourcePolicy?: any;
}
/**
 * Resource Type definition for AWS::SecretsManager::ResourcePolicy
 */
export function getResourcePolicyOutput(args: GetResourcePolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetResourcePolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:secretsmanager:getResourcePolicy", {
        "id": args.id,
    }, opts);
}

export interface GetResourcePolicyOutputArgs {
    /**
     * The Arn of the secret.
     */
    id: pulumi.Input<string>;
}
