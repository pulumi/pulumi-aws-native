// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GuardDuty::Filter
 */
export function getFilter(args: GetFilterArgs, opts?: pulumi.InvokeOptions): Promise<GetFilterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:guardduty:getFilter", {
        "detectorId": args.detectorId,
        "name": args.name,
    }, opts);
}

export interface GetFilterArgs {
    /**
     * The detector ID associated with the GuardDuty account for which you want to create a filter.
     *
     * To find the `detectorId` in the current Region, see the
     * Settings page in the GuardDuty console, or run the [ListDetectors](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListDetectors.html) API.
     */
    detectorId: string;
    /**
     * The name of the filter. Valid characters include period (.), underscore (_), dash (-), and alphanumeric characters. A whitespace is considered to be an invalid character.
     */
    name: string;
}

export interface GetFilterResult {
    /**
     * Specifies the action that is to be applied to the findings that match the filter.
     */
    readonly action?: string;
    /**
     * The description of the filter. Valid characters include alphanumeric characters, and special characters such as hyphen, period, colon, underscore, parentheses ( `{ }` , `[ ]` , and `( )` ), forward slash, horizontal tab, vertical tab, newline, form feed, return, and whitespace.
     */
    readonly description?: string;
    /**
     * Represents the criteria to be used in the filter for querying findings.
     */
    readonly findingCriteria?: outputs.guardduty.FilterFindingCriteria;
    /**
     * Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings. The minimum value for this property is 1 and the maximum is 100.
     *
     * By default, filters may not be created in the same order as they are ranked. To ensure that the filters are created in the expected order, you can use an optional attribute, [DependsOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) , with the following syntax: `"DependsOn":[ "ObjectName" ]` .
     */
    readonly rank?: number;
    /**
     * The tags to be added to a new filter resource. Each tag consists of a key and an optional value, both of which you define.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::GuardDuty::Filter
 */
export function getFilterOutput(args: GetFilterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFilterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:guardduty:getFilter", {
        "detectorId": args.detectorId,
        "name": args.name,
    }, opts);
}

export interface GetFilterOutputArgs {
    /**
     * The detector ID associated with the GuardDuty account for which you want to create a filter.
     *
     * To find the `detectorId` in the current Region, see the
     * Settings page in the GuardDuty console, or run the [ListDetectors](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListDetectors.html) API.
     */
    detectorId: pulumi.Input<string>;
    /**
     * The name of the filter. Valid characters include period (.), underscore (_), dash (-), and alphanumeric characters. A whitespace is considered to be an invalid character.
     */
    name: pulumi.Input<string>;
}
