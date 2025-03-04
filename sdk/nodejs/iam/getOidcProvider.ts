// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::OIDCProvider
 */
export function getOidcProvider(args: GetOidcProviderArgs, opts?: pulumi.InvokeOptions): Promise<GetOidcProviderResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iam:getOidcProvider", {
        "arn": args.arn,
    }, opts);
}

export interface GetOidcProviderArgs {
    /**
     * Amazon Resource Name (ARN) of the OIDC provider
     */
    arn: string;
}

export interface GetOidcProviderResult {
    /**
     * Amazon Resource Name (ARN) of the OIDC provider
     */
    readonly arn?: string;
    /**
     * A list of client IDs (also known as audiences) that are associated with the specified IAM OIDC provider resource object. For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
     */
    readonly clientIdList?: string[];
    /**
     * A list of tags that are attached to the specified IAM OIDC provider. The returned list of tags is sorted by tag key. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide* .
     */
    readonly tags?: outputs.Tag[];
    /**
     * A list of certificate thumbprints that are associated with the specified IAM OIDC provider resource object. For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
     *
     * This property is optional. If it is not included, IAM will retrieve and use the top intermediate certificate authority (CA) thumbprint of the OpenID Connect identity provider server certificate.
     */
    readonly thumbprintList?: string[];
}
/**
 * Resource Type definition for AWS::IAM::OIDCProvider
 */
export function getOidcProviderOutput(args: GetOidcProviderOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOidcProviderResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:iam:getOidcProvider", {
        "arn": args.arn,
    }, opts);
}

export interface GetOidcProviderOutputArgs {
    /**
     * Amazon Resource Name (ARN) of the OIDC provider
     */
    arn: pulumi.Input<string>;
}
