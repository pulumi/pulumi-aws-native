// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An object representing an Amazon EKS IdentityProviderConfig.
 */
export function getIdentityProviderConfig(args: GetIdentityProviderConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityProviderConfigResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:eks:getIdentityProviderConfig", {
        "clusterName": args.clusterName,
        "identityProviderConfigName": args.identityProviderConfigName,
        "type": args.type,
    }, opts);
}

export interface GetIdentityProviderConfigArgs {
    /**
     * The name of the identity provider configuration.
     */
    clusterName: string;
    /**
     * The name of the OIDC provider configuration.
     */
    identityProviderConfigName: string;
    /**
     * The type of the identity provider configuration.
     */
    type: enums.eks.IdentityProviderConfigType;
}

export interface GetIdentityProviderConfigResult {
    /**
     * The ARN of the configuration.
     */
    readonly identityProviderConfigArn?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * An object representing an Amazon EKS IdentityProviderConfig.
 */
export function getIdentityProviderConfigOutput(args: GetIdentityProviderConfigOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIdentityProviderConfigResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:eks:getIdentityProviderConfig", {
        "clusterName": args.clusterName,
        "identityProviderConfigName": args.identityProviderConfigName,
        "type": args.type,
    }, opts);
}

export interface GetIdentityProviderConfigOutputArgs {
    /**
     * The name of the identity provider configuration.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The name of the OIDC provider configuration.
     */
    identityProviderConfigName: pulumi.Input<string>;
    /**
     * The type of the identity provider configuration.
     */
    type: pulumi.Input<enums.eks.IdentityProviderConfigType>;
}
