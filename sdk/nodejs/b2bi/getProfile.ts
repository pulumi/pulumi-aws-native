// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::B2BI::Profile Resource Type
 */
export function getProfile(args: GetProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetProfileResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:b2bi:getProfile", {
        "profileId": args.profileId,
    }, opts);
}

export interface GetProfileArgs {
    profileId: string;
}

export interface GetProfileResult {
    readonly businessName?: string;
    readonly createdAt?: string;
    readonly email?: string;
    readonly logGroupName?: string;
    readonly modifiedAt?: string;
    readonly name?: string;
    readonly phone?: string;
    readonly profileArn?: string;
    readonly profileId?: string;
    readonly tags?: outputs.b2bi.ProfileTag[];
}
/**
 * Definition of AWS::B2BI::Profile Resource Type
 */
export function getProfileOutput(args: GetProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProfileResult> {
    return pulumi.output(args).apply((a: any) => getProfile(a, opts))
}

export interface GetProfileOutputArgs {
    profileId: pulumi.Input<string>;
}