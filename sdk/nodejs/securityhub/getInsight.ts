// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::SecurityHub::Insight resource represents the AWS Security Hub Insight in your account. An AWS Security Hub insight is a collection of related findings.
 */
export function getInsight(args: GetInsightArgs, opts?: pulumi.InvokeOptions): Promise<GetInsightResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:securityhub:getInsight", {
        "insightArn": args.insightArn,
    }, opts);
}

export interface GetInsightArgs {
    /**
     * The ARN of a Security Hub insight
     */
    insightArn: string;
}

export interface GetInsightResult {
    /**
     * One or more attributes used to filter the findings included in the insight
     */
    readonly filters?: outputs.securityhub.InsightAwsSecurityFindingFilters;
    /**
     * The grouping attribute for the insight's findings
     */
    readonly groupByAttribute?: string;
    /**
     * The ARN of a Security Hub insight
     */
    readonly insightArn?: string;
    /**
     * The name of a Security Hub insight
     */
    readonly name?: string;
}
/**
 * The AWS::SecurityHub::Insight resource represents the AWS Security Hub Insight in your account. An AWS Security Hub insight is a collection of related findings.
 */
export function getInsightOutput(args: GetInsightOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInsightResult> {
    return pulumi.output(args).apply((a: any) => getInsight(a, opts))
}

export interface GetInsightOutputArgs {
    /**
     * The ARN of a Security Hub insight
     */
    insightArn: pulumi.Input<string>;
}