// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
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
     * A map for Idle Connection Preferences
     */
    readonly idleConnection?: outputs.ssmguiconnect.PreferencesIdleConnectionPreferences[];
}
/**
 * Definition of AWS::SSMGuiConnect::Preferences Resource Type
 */
export function getPreferencesOutput(args: GetPreferencesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPreferencesResult> {
    return pulumi.output(args).apply((a: any) => getPreferences(a, opts))
}

export interface GetPreferencesOutputArgs {
    /**
     * The AWS Account Id that the preference is associated with, used as the unique identifier for this resource.
     */
    accountId: pulumi.Input<string>;
}