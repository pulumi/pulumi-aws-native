// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::NetworkManager::Link type describes a link.
 */
export function getLink(args: GetLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetLinkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:networkmanager:getLink", {
        "globalNetworkId": args.globalNetworkId,
        "linkId": args.linkId,
    }, opts);
}

export interface GetLinkArgs {
    /**
     * The ID of the global network.
     */
    globalNetworkId: string;
    /**
     * The ID of the link.
     */
    linkId: string;
}

export interface GetLinkResult {
    /**
     * The Bandwidth for the link.
     */
    readonly bandwidth?: outputs.networkmanager.LinkBandwidth;
    /**
     * The date and time that the device was created.
     */
    readonly createdAt?: string;
    /**
     * The description of the link.
     */
    readonly description?: string;
    /**
     * The Amazon Resource Name (ARN) of the link.
     */
    readonly linkArn?: string;
    /**
     * The ID of the link.
     */
    readonly linkId?: string;
    /**
     * The provider of the link.
     */
    readonly provider?: string;
    /**
     * The state of the link.
     */
    readonly state?: string;
    /**
     * The tags for the link.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The type of the link.
     */
    readonly type?: string;
}
/**
 * The AWS::NetworkManager::Link type describes a link.
 */
export function getLinkOutput(args: GetLinkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLinkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:networkmanager:getLink", {
        "globalNetworkId": args.globalNetworkId,
        "linkId": args.linkId,
    }, opts);
}

export interface GetLinkOutputArgs {
    /**
     * The ID of the global network.
     */
    globalNetworkId: pulumi.Input<string>;
    /**
     * The ID of the link.
     */
    linkId: pulumi.Input<string>;
}
