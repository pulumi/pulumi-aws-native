// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Connect::QuickConnect
 */
export function getQuickConnect(args: GetQuickConnectArgs, opts?: pulumi.InvokeOptions): Promise<GetQuickConnectResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:connect:getQuickConnect", {
        "quickConnectArn": args.quickConnectArn,
    }, opts);
}

export interface GetQuickConnectArgs {
    /**
     * The Amazon Resource Name (ARN) for the quick connect.
     */
    quickConnectArn: string;
}

export interface GetQuickConnectResult {
    /**
     * The description of the quick connect.
     */
    readonly description?: string;
    /**
     * The identifier of the Amazon Connect instance.
     */
    readonly instanceArn?: string;
    /**
     * The name of the quick connect.
     */
    readonly name?: string;
    /**
     * The Amazon Resource Name (ARN) for the quick connect.
     */
    readonly quickConnectArn?: string;
    /**
     * Configuration settings for the quick connect.
     */
    readonly quickConnectConfig?: outputs.connect.QuickConnectConfig;
    /**
     * The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).
     */
    readonly quickConnectType?: enums.connect.QuickConnectType;
    /**
     * One or more tags.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::Connect::QuickConnect
 */
export function getQuickConnectOutput(args: GetQuickConnectOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetQuickConnectResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:connect:getQuickConnect", {
        "quickConnectArn": args.quickConnectArn,
    }, opts);
}

export interface GetQuickConnectOutputArgs {
    /**
     * The Amazon Resource Name (ARN) for the quick connect.
     */
    quickConnectArn: pulumi.Input<string>;
}
