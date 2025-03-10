// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Connect::ContactFlowModule.
 */
export function getContactFlowModule(args: GetContactFlowModuleArgs, opts?: pulumi.InvokeOptions): Promise<GetContactFlowModuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:connect:getContactFlowModule", {
        "contactFlowModuleArn": args.contactFlowModuleArn,
    }, opts);
}

export interface GetContactFlowModuleArgs {
    /**
     * The identifier of the contact flow module (ARN).
     */
    contactFlowModuleArn: string;
}

export interface GetContactFlowModuleResult {
    /**
     * The identifier of the contact flow module (ARN).
     */
    readonly contactFlowModuleArn?: string;
    /**
     * The content of the contact flow module in JSON format.
     */
    readonly content?: string;
    /**
     * The description of the contact flow module.
     */
    readonly description?: string;
    /**
     * The identifier of the Amazon Connect instance (ARN).
     */
    readonly instanceArn?: string;
    /**
     * The name of the contact flow module.
     */
    readonly name?: string;
    /**
     * The state of the contact flow module.
     */
    readonly state?: string;
    /**
     * The status of the contact flow module.
     */
    readonly status?: string;
    /**
     * One or more tags.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::Connect::ContactFlowModule.
 */
export function getContactFlowModuleOutput(args: GetContactFlowModuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetContactFlowModuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:connect:getContactFlowModule", {
        "contactFlowModuleArn": args.contactFlowModuleArn,
    }, opts);
}

export interface GetContactFlowModuleOutputArgs {
    /**
     * The identifier of the contact flow module (ARN).
     */
    contactFlowModuleArn: pulumi.Input<string>;
}
