// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a stored analysis within a collaboration
 */
export function getAnalysisTemplate(args: GetAnalysisTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetAnalysisTemplateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cleanrooms:getAnalysisTemplate", {
        "analysisTemplateIdentifier": args.analysisTemplateIdentifier,
        "membershipIdentifier": args.membershipIdentifier,
    }, opts);
}

export interface GetAnalysisTemplateArgs {
    analysisTemplateIdentifier: string;
    membershipIdentifier: string;
}

export interface GetAnalysisTemplateResult {
    readonly analysisTemplateIdentifier?: string;
    readonly arn?: string;
    readonly collaborationArn?: string;
    readonly collaborationIdentifier?: string;
    readonly description?: string;
    readonly membershipArn?: string;
    readonly schema?: outputs.cleanrooms.AnalysisTemplateAnalysisSchema;
    /**
     * An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.
     */
    readonly tags?: outputs.cleanrooms.AnalysisTemplateTag[];
}
/**
 * Represents a stored analysis within a collaboration
 */
export function getAnalysisTemplateOutput(args: GetAnalysisTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAnalysisTemplateResult> {
    return pulumi.output(args).apply((a: any) => getAnalysisTemplate(a, opts))
}

export interface GetAnalysisTemplateOutputArgs {
    analysisTemplateIdentifier: pulumi.Input<string>;
    membershipIdentifier: pulumi.Input<string>;
}