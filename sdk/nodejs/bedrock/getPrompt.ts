// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Bedrock::Prompt Resource Type
 */
export function getPrompt(args: GetPromptArgs, opts?: pulumi.InvokeOptions): Promise<GetPromptResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:bedrock:getPrompt", {
        "arn": args.arn,
    }, opts);
}

export interface GetPromptArgs {
    /**
     * ARN of a prompt resource possibly with a version
     */
    arn: string;
}

export interface GetPromptResult {
    /**
     * ARN of a prompt resource possibly with a version
     */
    readonly arn?: string;
    /**
     * Time Stamp.
     */
    readonly createdAt?: string;
    /**
     * A KMS key ARN
     */
    readonly customerEncryptionKeyArn?: string;
    /**
     * Name for a variant.
     */
    readonly defaultVariant?: string;
    /**
     * Name for a prompt resource.
     */
    readonly description?: string;
    /**
     * Identifier for a Prompt
     */
    readonly id?: string;
    /**
     * Name for a prompt resource.
     */
    readonly name?: string;
    /**
     * Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
     *
     * - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
     * - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
     */
    readonly tags?: {[key: string]: string};
    /**
     * Time Stamp.
     */
    readonly updatedAt?: string;
    /**
     * List of prompt variants
     */
    readonly variants?: outputs.bedrock.PromptVariant[];
    /**
     * Draft Version.
     */
    readonly version?: string;
}
/**
 * Definition of AWS::Bedrock::Prompt Resource Type
 */
export function getPromptOutput(args: GetPromptOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPromptResult> {
    return pulumi.output(args).apply((a: any) => getPrompt(a, opts))
}

export interface GetPromptOutputArgs {
    /**
     * ARN of a prompt resource possibly with a version
     */
    arn: pulumi.Input<string>;
}