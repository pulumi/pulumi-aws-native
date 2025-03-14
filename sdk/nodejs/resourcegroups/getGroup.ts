// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Schema for ResourceGroups::Group
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:resourcegroups:getGroup", {
        "name": args.name,
    }, opts);
}

export interface GetGroupArgs {
    /**
     * The name of the resource group
     */
    name: string;
}

export interface GetGroupResult {
    /**
     * The Resource Group ARN.
     */
    readonly arn?: string;
    /**
     * The service configuration currently associated with the resource group and in effect for the members of the resource group. A `Configuration` consists of one or more `ConfigurationItem` entries. For information about service configurations for resource groups and how to construct them, see [Service configurations for resource groups](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html) in the *AWS Resource Groups User Guide* .
     *
     * > You can include either a `Configuration` or a `ResourceQuery` , but not both.
     */
    readonly configuration?: outputs.resourcegroups.GroupConfigurationItem[];
    /**
     * The description of the resource group
     */
    readonly description?: string;
    /**
     * The resource query structure that is used to dynamically determine which AWS resources are members of the associated resource group. For more information about queries and how to construct them, see [Build queries and groups in AWS Resource Groups](https://docs.aws.amazon.com//ARG/latest/userguide/gettingstarted-query.html) in the *AWS Resource Groups User Guide*
     *
     * > - You can include either a `ResourceQuery` or a `Configuration` , but not both.
     * > - You can specify the group's membership either by using a `ResourceQuery` or by using a list of `Resources` , but not both.
     */
    readonly resourceQuery?: outputs.resourcegroups.GroupResourceQuery;
    /**
     * A list of the Amazon Resource Names (ARNs) of AWS resources that you want to add to the specified group.
     *
     * > - You can specify the group membership either by using a list of `Resources` or by using a `ResourceQuery` , but not both.
     * > - You can include a `Resources` property only if you also specify a `Configuration` property.
     */
    readonly resources?: string[];
    /**
     * The tag key and value pairs that are attached to the resource group.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Schema for ResourceGroups::Group
 */
export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:resourcegroups:getGroup", {
        "name": args.name,
    }, opts);
}

export interface GetGroupOutputArgs {
    /**
     * The name of the resource group
     */
    name: pulumi.Input<string>;
}
