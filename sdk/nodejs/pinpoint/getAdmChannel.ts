// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::ADMChannel
 */
export function getAdmChannel(args: GetAdmChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetAdmChannelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:pinpoint:getAdmChannel", {
        "id": args.id,
    }, opts);
}

export interface GetAdmChannelArgs {
    id: string;
}

export interface GetAdmChannelResult {
    readonly clientId?: string;
    readonly clientSecret?: string;
    readonly enabled?: boolean;
    readonly id?: string;
}
/**
 * Resource Type definition for AWS::Pinpoint::ADMChannel
 */
export function getAdmChannelOutput(args: GetAdmChannelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAdmChannelResult> {
    return pulumi.output(args).apply((a: any) => getAdmChannel(a, opts))
}

export interface GetAdmChannelOutputArgs {
    id: pulumi.Input<string>;
}