// *** WARNING: this file was generated by pulumi-language-nodejs. ***
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
    /**
     * Returns the identifier for the analysis template.
     *
     * Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`
     */
    analysisTemplateIdentifier: string;
    /**
     * The identifier for a membership resource.
     */
    membershipIdentifier: string;
}

export interface GetAnalysisTemplateResult {
    /**
     * Returns the identifier for the analysis template.
     *
     * Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`
     */
    readonly analysisTemplateIdentifier?: string;
    /**
     * Returns the Amazon Resource Name (ARN) of the analysis template.
     *
     * Example: `arn:aws:cleanrooms:us-east-1:111122223333:membership/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111/analysistemplates/a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`
     */
    readonly arn?: string;
    /**
     * Returns the unique ARN for the analysis template’s associated collaboration.
     *
     * Example: `arn:aws:cleanrooms:us-east-1:111122223333:collaboration/a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`
     */
    readonly collaborationArn?: string;
    /**
     * Returns the unique ID for the associated collaboration of the analysis template.
     *
     * Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`
     */
    readonly collaborationIdentifier?: string;
    /**
     * The description of the analysis template.
     */
    readonly description?: string;
    /**
     * Returns the Amazon Resource Name (ARN) of the member who created the analysis template.
     *
     * Example: `arn:aws:cleanrooms:us-east-1:111122223333:membership/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
     */
    readonly membershipArn?: string;
    /**
     * The source metadata for the analysis template.
     */
    readonly sourceMetadata?: outputs.cleanrooms.AnalysisTemplateAnalysisSourceMetadataProperties;
    /**
     * An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Represents a stored analysis within a collaboration
 */
export function getAnalysisTemplateOutput(args: GetAnalysisTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAnalysisTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:cleanrooms:getAnalysisTemplate", {
        "analysisTemplateIdentifier": args.analysisTemplateIdentifier,
        "membershipIdentifier": args.membershipIdentifier,
    }, opts);
}

export interface GetAnalysisTemplateOutputArgs {
    /**
     * Returns the identifier for the analysis template.
     *
     * Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`
     */
    analysisTemplateIdentifier: pulumi.Input<string>;
    /**
     * The identifier for a membership resource.
     */
    membershipIdentifier: pulumi.Input<string>;
}
