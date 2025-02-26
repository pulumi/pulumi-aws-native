// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Definition of AWS::MediaLive::EventBridgeRuleTemplateGroup Resource Type
 */
export function getEventBridgeRuleTemplateGroup(args: GetEventBridgeRuleTemplateGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetEventBridgeRuleTemplateGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:medialive:getEventBridgeRuleTemplateGroup", {
        "identifier": args.identifier,
    }, opts);
}

export interface GetEventBridgeRuleTemplateGroupArgs {
    identifier: string;
}

export interface GetEventBridgeRuleTemplateGroupResult {
    /**
     * An eventbridge rule template group's ARN (Amazon Resource Name)
     */
    readonly arn?: string;
    /**
     * The date and time of resource creation.
     */
    readonly createdAt?: string;
    /**
     * A resource's optional description.
     */
    readonly description?: string;
    /**
     * An eventbridge rule template group's id. AWS provided template groups have ids that start with `aws-`
     */
    readonly id?: string;
    readonly identifier?: string;
    /**
     * The date and time of latest resource modification.
     */
    readonly modifiedAt?: string;
}
/**
 * Definition of AWS::MediaLive::EventBridgeRuleTemplateGroup Resource Type
 */
export function getEventBridgeRuleTemplateGroupOutput(args: GetEventBridgeRuleTemplateGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEventBridgeRuleTemplateGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:medialive:getEventBridgeRuleTemplateGroup", {
        "identifier": args.identifier,
    }, opts);
}

export interface GetEventBridgeRuleTemplateGroupOutputArgs {
    identifier: pulumi.Input<string>;
}
