// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Bedrock::KnowledgeBase Resource Type
 */
export function getKnowledgeBase(args: GetKnowledgeBaseArgs, opts?: pulumi.InvokeOptions): Promise<GetKnowledgeBaseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:bedrock:getKnowledgeBase", {
        "knowledgeBaseId": args.knowledgeBaseId,
    }, opts);
}

export interface GetKnowledgeBaseArgs {
    /**
     * The unique identifier of the knowledge base.
     */
    knowledgeBaseId: string;
}

export interface GetKnowledgeBaseResult {
    /**
     * The time at which the knowledge base was created.
     */
    readonly createdAt?: string;
    /**
     * Description of the Resource.
     */
    readonly description?: string;
    /**
     * A list of reasons that the API operation on the knowledge base failed.
     */
    readonly failureReasons?: string[];
    /**
     * The ARN of the knowledge base.
     */
    readonly knowledgeBaseArn?: string;
    /**
     * The unique identifier of the knowledge base.
     */
    readonly knowledgeBaseId?: string;
    /**
     * The name of the knowledge base.
     */
    readonly name?: string;
    /**
     * The ARN of the IAM role with permissions to invoke API operations on the knowledge base. The ARN must begin with AmazonBedrockExecutionRoleForKnowledgeBase_
     */
    readonly roleArn?: string;
    readonly status?: enums.bedrock.KnowledgeBaseStatus;
    readonly tags?: {[key: string]: string};
    /**
     * The time at which the knowledge base was last updated.
     */
    readonly updatedAt?: string;
}
/**
 * Definition of AWS::Bedrock::KnowledgeBase Resource Type
 */
export function getKnowledgeBaseOutput(args: GetKnowledgeBaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKnowledgeBaseResult> {
    return pulumi.output(args).apply((a: any) => getKnowledgeBase(a, opts))
}

export interface GetKnowledgeBaseOutputArgs {
    /**
     * The unique identifier of the knowledge base.
     */
    knowledgeBaseId: pulumi.Input<string>;
}