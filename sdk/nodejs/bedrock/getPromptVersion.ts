// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Bedrock::PromptVersion Resource Type
 */
export function getPromptVersion(args: GetPromptVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetPromptVersionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:bedrock:getPromptVersion", {
        "arn": args.arn,
    }, opts);
}

export interface GetPromptVersionArgs {
    /**
     * ARN of a prompt version resource
     */
    arn: string;
}

export interface GetPromptVersionResult {
    /**
     * ARN of a prompt version resource
     */
    readonly arn?: string;
    /**
     * Time Stamp.
     */
    readonly createdAt?: string;
    /**
     * Name for a variant.
     */
    readonly defaultVariant?: string;
    /**
     * Name for a prompt resource.
     */
    readonly name?: string;
    /**
     * Identifier for a Prompt
     */
    readonly promptId?: string;
    /**
     * Time Stamp.
     */
    readonly updatedAt?: string;
    /**
     * List of prompt variants
     */
    readonly variants?: outputs.bedrock.PromptVersionPromptVariant[];
    /**
     * Version.
     */
    readonly version?: string;
}
/**
 * Definition of AWS::Bedrock::PromptVersion Resource Type
 */
export function getPromptVersionOutput(args: GetPromptVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPromptVersionResult> {
    return pulumi.output(args).apply((a: any) => getPromptVersion(a, opts))
}

export interface GetPromptVersionOutputArgs {
    /**
     * ARN of a prompt version resource
     */
    arn: pulumi.Input<string>;
}