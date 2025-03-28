// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Wisdom::KnowledgeBase Resource Type
 */
export function getKnowledgeBase(args: GetKnowledgeBaseArgs, opts?: pulumi.InvokeOptions): Promise<GetKnowledgeBaseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:wisdom:getKnowledgeBase", {
        "knowledgeBaseId": args.knowledgeBaseId,
    }, opts);
}

export interface GetKnowledgeBaseArgs {
    /**
     * The ID of the knowledge base.
     */
    knowledgeBaseId: string;
}

export interface GetKnowledgeBaseResult {
    /**
     * The Amazon Resource Name (ARN) of the knowledge base.
     */
    readonly knowledgeBaseArn?: string;
    /**
     * The ID of the knowledge base.
     */
    readonly knowledgeBaseId?: string;
    /**
     * Information about how to render the content.
     */
    readonly renderingConfiguration?: outputs.wisdom.KnowledgeBaseRenderingConfiguration;
    /**
     * Contains details about how to ingest the documents in a data source.
     */
    readonly vectorIngestionConfiguration?: outputs.wisdom.KnowledgeBaseVectorIngestionConfiguration;
}
/**
 * Definition of AWS::Wisdom::KnowledgeBase Resource Type
 */
export function getKnowledgeBaseOutput(args: GetKnowledgeBaseOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetKnowledgeBaseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:wisdom:getKnowledgeBase", {
        "knowledgeBaseId": args.knowledgeBaseId,
    }, opts);
}

export interface GetKnowledgeBaseOutputArgs {
    /**
     * The ID of the knowledge base.
     */
    knowledgeBaseId: pulumi.Input<string>;
}
