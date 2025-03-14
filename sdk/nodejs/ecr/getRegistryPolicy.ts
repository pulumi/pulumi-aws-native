// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The ``AWS::ECR::RegistryPolicy`` resource creates or updates the permissions policy for a private registry.
 *  A private registry policy is used to specify permissions for another AWS-account and is used when configuring cross-account replication. For more information, see [Registry permissions](https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html) in the *Amazon Elastic Container Registry User Guide*.
 */
export function getRegistryPolicy(args: GetRegistryPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ecr:getRegistryPolicy", {
        "registryId": args.registryId,
    }, opts);
}

export interface GetRegistryPolicyArgs {
    /**
     * The account ID of the private registry the policy is associated with.
     */
    registryId: string;
}

export interface GetRegistryPolicyResult {
    /**
     * The JSON policy text for your registry.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ECR::RegistryPolicy` for more information about the expected schema for this property.
     */
    readonly policyText?: any;
    /**
     * The account ID of the private registry the policy is associated with.
     */
    readonly registryId?: string;
}
/**
 * The ``AWS::ECR::RegistryPolicy`` resource creates or updates the permissions policy for a private registry.
 *  A private registry policy is used to specify permissions for another AWS-account and is used when configuring cross-account replication. For more information, see [Registry permissions](https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html) in the *Amazon Elastic Container Registry User Guide*.
 */
export function getRegistryPolicyOutput(args: GetRegistryPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRegistryPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ecr:getRegistryPolicy", {
        "registryId": args.registryId,
    }, opts);
}

export interface GetRegistryPolicyOutputArgs {
    /**
     * The account ID of the private registry the policy is associated with.
     */
    registryId: pulumi.Input<string>;
}
