// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::CodeSigningConfig.
 */
export function getCodeSigningConfig(args: GetCodeSigningConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetCodeSigningConfigResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:lambda:getCodeSigningConfig", {
        "codeSigningConfigArn": args.codeSigningConfigArn,
    }, opts);
}

export interface GetCodeSigningConfigArgs {
    /**
     * A unique Arn for CodeSigningConfig resource
     */
    codeSigningConfigArn: string;
}

export interface GetCodeSigningConfigResult {
    /**
     * When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list
     */
    readonly allowedPublishers?: outputs.lambda.CodeSigningConfigAllowedPublishers;
    /**
     * A unique Arn for CodeSigningConfig resource
     */
    readonly codeSigningConfigArn?: string;
    /**
     * A unique identifier for CodeSigningConfig resource
     */
    readonly codeSigningConfigId?: string;
    /**
     * Policies to control how to act if a signature is invalid
     */
    readonly codeSigningPolicies?: outputs.lambda.CodeSigningConfigCodeSigningPolicies;
    /**
     * A description of the CodeSigningConfig
     */
    readonly description?: string;
    /**
     * A list of tags to apply to CodeSigningConfig resource
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::Lambda::CodeSigningConfig.
 */
export function getCodeSigningConfigOutput(args: GetCodeSigningConfigOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCodeSigningConfigResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:lambda:getCodeSigningConfig", {
        "codeSigningConfigArn": args.codeSigningConfigArn,
    }, opts);
}

export interface GetCodeSigningConfigOutputArgs {
    /**
     * A unique Arn for CodeSigningConfig resource
     */
    codeSigningConfigArn: pulumi.Input<string>;
}
