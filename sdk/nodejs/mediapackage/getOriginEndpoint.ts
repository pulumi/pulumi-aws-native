// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaPackage::OriginEndpoint
 */
export function getOriginEndpoint(args: GetOriginEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetOriginEndpointResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:mediapackage:getOriginEndpoint", {
        "id": args.id,
    }, opts);
}

export interface GetOriginEndpointArgs {
    /**
     * The ID of the OriginEndpoint.
     */
    id: string;
}

export interface GetOriginEndpointResult {
    /**
     * The Amazon Resource Name (ARN) assigned to the OriginEndpoint.
     */
    readonly arn?: string;
    readonly authorization?: outputs.mediapackage.OriginEndpointAuthorization;
    /**
     * The ID of the Channel the OriginEndpoint is associated with.
     */
    readonly channelId?: string;
    readonly cmafPackage?: outputs.mediapackage.OriginEndpointCmafPackage;
    readonly dashPackage?: outputs.mediapackage.OriginEndpointDashPackage;
    /**
     * A short text description of the OriginEndpoint.
     */
    readonly description?: string;
    readonly hlsPackage?: outputs.mediapackage.OriginEndpointHlsPackage;
    /**
     * A short string appended to the end of the OriginEndpoint URL.
     */
    readonly manifestName?: string;
    readonly mssPackage?: outputs.mediapackage.OriginEndpointMssPackage;
    /**
     * Control whether origination of video is allowed for this OriginEndpoint. If set to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of access control. If set to DENY, the OriginEndpoint may not be requested. This can be helpful for Live to VOD harvesting, or for temporarily disabling origination
     */
    readonly origination?: enums.mediapackage.OriginEndpointOrigination;
    /**
     * Maximum duration (seconds) of content to retain for startover playback. If not specified, startover playback will be disabled for the OriginEndpoint.
     */
    readonly startoverWindowSeconds?: number;
    /**
     * A collection of tags associated with a resource
     */
    readonly tags?: outputs.mediapackage.OriginEndpointTag[];
    /**
     * Amount of delay (seconds) to enforce on the playback of live content. If not specified, there will be no time delay in effect for the OriginEndpoint.
     */
    readonly timeDelaySeconds?: number;
    /**
     * The URL of the packaged OriginEndpoint for consumption.
     */
    readonly url?: string;
    /**
     * A list of source IP CIDR blocks that will be allowed to access the OriginEndpoint.
     */
    readonly whitelist?: string[];
}

export function getOriginEndpointOutput(args: GetOriginEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOriginEndpointResult> {
    return pulumi.output(args).apply(a => getOriginEndpoint(a, opts))
}

export interface GetOriginEndpointOutputArgs {
    /**
     * The ID of the OriginEndpoint.
     */
    id: pulumi.Input<string>;
}