// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::IdentityPool
 */
export function getIdentityPool(args: GetIdentityPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityPoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cognito:getIdentityPool", {
        "id": args.id,
    }, opts);
}

export interface GetIdentityPoolArgs {
    id: string;
}

export interface GetIdentityPoolResult {
    /**
     * Enables the Basic (Classic) authentication flow.
     */
    readonly allowClassicFlow?: boolean;
    /**
     * Specifies whether the identity pool supports unauthenticated logins.
     */
    readonly allowUnauthenticatedIdentities?: boolean;
    /**
     * The Amazon Cognito user pools and their client IDs.
     */
    readonly cognitoIdentityProviders?: outputs.cognito.IdentityPoolCognitoIdentityProvider[];
    /**
     * The "domain" Amazon Cognito uses when referencing your users. This name acts as a placeholder that allows your backend and the Amazon Cognito service to communicate about the developer provider. For the `DeveloperProviderName` , you can use letters and periods (.), underscores (_), and dashes (-).
     *
     * *Minimum length* : 1
     *
     * *Maximum length* : 100
     */
    readonly developerProviderName?: string;
    readonly id?: string;
    /**
     * The name of your Amazon Cognito identity pool.
     *
     * *Minimum length* : 1
     *
     * *Maximum length* : 128
     *
     * *Pattern* : `[\w\s+=,.@-]+`
     */
    readonly identityPoolName?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly identityPoolTags?: outputs.Tag[];
    /**
     * The name of the Amazon Cognito identity pool, returned as a string.
     */
    readonly name?: string;
    /**
     * The Amazon Resource Names (ARNs) of the OpenID connect providers.
     */
    readonly openIdConnectProviderArns?: string[];
    /**
     * The Amazon Resource Names (ARNs) of the Security Assertion Markup Language (SAML) providers.
     */
    readonly samlProviderArns?: string[];
    /**
     * Key-value pairs that map provider names to provider app IDs.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Cognito::IdentityPool` for more information about the expected schema for this property.
     */
    readonly supportedLoginProviders?: any;
}
/**
 * Resource Type definition for AWS::Cognito::IdentityPool
 */
export function getIdentityPoolOutput(args: GetIdentityPoolOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIdentityPoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:cognito:getIdentityPool", {
        "id": args.id,
    }, opts);
}

export interface GetIdentityPoolOutputArgs {
    id: pulumi.Input<string>;
}
