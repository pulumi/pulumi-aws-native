// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::SSMGuiConnect::Preferences Resource Type
 */
export function getPreferences(args: GetPreferencesArgs, opts?: pulumi.InvokeOptions): Promise<GetPreferencesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ssmguiconnect:getPreferences", {
        "accountId": args.accountId,
    }, opts);
}

export interface GetPreferencesArgs {
    /**
     * The AWS Account Id that the preference is associated with, used as the unique identifier for this resource.
     */
    accountId: string;
}

export interface GetPreferencesResult {
    /**
     * The AWS Account Id that the preference is associated with, used as the unique identifier for this resource.
     */
    readonly accountId?: string;
    /**
     * The set of preferences used for recording RDP connections in the requesting AWS account and AWS Region. This includes details such as which S3 bucket recordings are stored in.
     */
    readonly connectionRecordingPreferences?: outputs.ssmguiconnect.ConnectionRecordingPreferencesProperties;
}
/**
 * Definition of AWS::SSMGuiConnect::Preferences Resource Type
 */
export function getPreferencesOutput(args: GetPreferencesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPreferencesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ssmguiconnect:getPreferences", {
        "accountId": args.accountId,
    }, opts);
}

export interface GetPreferencesOutputArgs {
    /**
     * The AWS Account Id that the preference is associated with, used as the unique identifier for this resource.
     */
    accountId: pulumi.Input<string>;
}
