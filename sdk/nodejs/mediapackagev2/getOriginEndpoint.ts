// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * <p>Represents an origin endpoint that is associated with a channel, offering a dynamically repackaged version of its content through various streaming media protocols. The content can be efficiently disseminated to end-users via a Content Delivery Network (CDN), like Amazon CloudFront.</p>
 */
export function getOriginEndpoint(args: GetOriginEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetOriginEndpointResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:mediapackagev2:getOriginEndpoint", {
        "arn": args.arn,
    }, opts);
}

export interface GetOriginEndpointArgs {
    /**
     * <p>The Amazon Resource Name (ARN) associated with the resource.</p>
     */
    arn: string;
}

export interface GetOriginEndpointResult {
    /**
     * <p>The Amazon Resource Name (ARN) associated with the resource.</p>
     */
    readonly arn?: string;
    readonly containerType?: enums.mediapackagev2.OriginEndpointContainerType;
    /**
     * <p>The date and time the origin endpoint was created.</p>
     */
    readonly createdAt?: string;
    /**
     * <p>Enter any descriptive text that helps you to identify the origin endpoint.</p>
     */
    readonly description?: string;
    /**
     * <p>An HTTP live streaming (HLS) manifest configuration.</p>
     */
    readonly hlsManifests?: outputs.mediapackagev2.OriginEndpointHlsManifestConfiguration[];
    /**
     * <p>A low-latency HLS manifest configuration.</p>
     */
    readonly lowLatencyHlsManifests?: outputs.mediapackagev2.OriginEndpointLowLatencyHlsManifestConfiguration[];
    /**
     * <p>The date and time the origin endpoint was modified.</p>
     */
    readonly modifiedAt?: string;
    readonly segment?: outputs.mediapackagev2.OriginEndpointSegment;
    /**
     * <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p>
     */
    readonly startoverWindowSeconds?: number;
    readonly tags?: outputs.mediapackagev2.OriginEndpointTag[];
}
/**
 * <p>Represents an origin endpoint that is associated with a channel, offering a dynamically repackaged version of its content through various streaming media protocols. The content can be efficiently disseminated to end-users via a Content Delivery Network (CDN), like Amazon CloudFront.</p>
 */
export function getOriginEndpointOutput(args: GetOriginEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOriginEndpointResult> {
    return pulumi.output(args).apply((a: any) => getOriginEndpoint(a, opts))
}

export interface GetOriginEndpointOutputArgs {
    /**
     * <p>The Amazon Resource Name (ARN) associated with the resource.</p>
     */
    arn: pulumi.Input<string>;
}