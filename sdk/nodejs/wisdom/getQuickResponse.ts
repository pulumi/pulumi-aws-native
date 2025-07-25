// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Wisdom::QuickResponse Resource Type.
 */
export function getQuickResponse(args: GetQuickResponseArgs, opts?: pulumi.InvokeOptions): Promise<GetQuickResponseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:wisdom:getQuickResponse", {
        "quickResponseArn": args.quickResponseArn,
    }, opts);
}

export interface GetQuickResponseArgs {
    /**
     * The Amazon Resource Name (ARN) of the quick response.
     */
    quickResponseArn: string;
}

export interface GetQuickResponseResult {
    /**
     * The Amazon Connect contact channels this quick response applies to.
     */
    readonly channels?: enums.wisdom.QuickResponseChannelType[];
    readonly content?: outputs.wisdom.QuickResponseContentProvider;
    /**
     * The media type of the quick response content.
     * - Use application/x.quickresponse;format=plain for quick response written in plain text.
     * - Use application/x.quickresponse;format=markdown for quick response written in richtext.
     */
    readonly contentType?: string;
    readonly contents?: outputs.wisdom.QuickResponseContents;
    /**
     * The description of the quick response.
     */
    readonly description?: string;
    readonly groupingConfiguration?: outputs.wisdom.QuickResponseGroupingConfiguration;
    /**
     * Whether the quick response is active.
     */
    readonly isActive?: boolean;
    /**
     * The language code value for the language in which the quick response is written. The supported language codes include de_DE, en_US, es_ES, fr_FR, id_ID, it_IT, ja_JP, ko_KR, pt_BR, zh_CN, zh_TW
     */
    readonly language?: string;
    /**
     * The name of the quick response.
     */
    readonly name?: string;
    /**
     * The Amazon Resource Name (ARN) of the quick response.
     */
    readonly quickResponseArn?: string;
    /**
     * The identifier of the quick response.
     */
    readonly quickResponseId?: string;
    /**
     * The shortcut key of the quick response. The value should be unique across the knowledge base.
     */
    readonly shortcutKey?: string;
    readonly status?: enums.wisdom.QuickResponseStatus;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Definition of AWS::Wisdom::QuickResponse Resource Type.
 */
export function getQuickResponseOutput(args: GetQuickResponseOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetQuickResponseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:wisdom:getQuickResponse", {
        "quickResponseArn": args.quickResponseArn,
    }, opts);
}

export interface GetQuickResponseOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the quick response.
     */
    quickResponseArn: pulumi.Input<string>;
}
