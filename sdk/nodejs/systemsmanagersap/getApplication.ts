// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::SystemsManagerSAP::Application
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:systemsmanagersap:getApplication", {
        "arn": args.arn,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * The ARN of the SSM-SAP application
     */
    arn: string;
}

export interface GetApplicationResult {
    /**
     * The ID of the application.
     */
    readonly applicationId?: string;
    /**
     * The type of the application.
     */
    readonly applicationType?: enums.systemsmanagersap.ApplicationType;
    /**
     * The ARN of the SSM-SAP application
     */
    readonly arn?: string;
    /**
     * The tags of a SystemsManagerSAP application.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource schema for AWS::SystemsManagerSAP::Application
 */
export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:systemsmanagersap:getApplication", {
        "arn": args.arn,
    }, opts);
}

export interface GetApplicationOutputArgs {
    /**
     * The ARN of the SSM-SAP application
     */
    arn: pulumi.Input<string>;
}
