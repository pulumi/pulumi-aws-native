// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::ServiceLinkedRole
 */
export function getServiceLinkedRole(args: GetServiceLinkedRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceLinkedRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iam:getServiceLinkedRole", {
        "roleName": args.roleName,
    }, opts);
}

export interface GetServiceLinkedRoleArgs {
    /**
     * The name of the role.
     */
    roleName: string;
}

export interface GetServiceLinkedRoleResult {
    /**
     * The description of the role.
     */
    readonly description?: string;
    /**
     * The name of the role.
     */
    readonly roleName?: string;
}
/**
 * Resource Type definition for AWS::IAM::ServiceLinkedRole
 */
export function getServiceLinkedRoleOutput(args: GetServiceLinkedRoleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceLinkedRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:iam:getServiceLinkedRole", {
        "roleName": args.roleName,
    }, opts);
}

export interface GetServiceLinkedRoleOutputArgs {
    /**
     * The name of the role.
     */
    roleName: pulumi.Input<string>;
}
