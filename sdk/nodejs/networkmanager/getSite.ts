// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::NetworkManager::Site type describes a site.
 */
export function getSite(args: GetSiteArgs, opts?: pulumi.InvokeOptions): Promise<GetSiteResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:networkmanager:getSite", {
        "globalNetworkId": args.globalNetworkId,
        "siteId": args.siteId,
    }, opts);
}

export interface GetSiteArgs {
    /**
     * The ID of the global network.
     */
    globalNetworkId: string;
    /**
     * The ID of the site.
     */
    siteId: string;
}

export interface GetSiteResult {
    /**
     * The date and time that the device was created.
     */
    readonly createdAt?: string;
    /**
     * The description of the site.
     */
    readonly description?: string;
    /**
     * The location of the site.
     */
    readonly location?: outputs.networkmanager.SiteLocation;
    /**
     * The Amazon Resource Name (ARN) of the site.
     */
    readonly siteArn?: string;
    /**
     * The ID of the site.
     */
    readonly siteId?: string;
    /**
     * The state of the site.
     */
    readonly state?: string;
    /**
     * The tags for the site.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * The AWS::NetworkManager::Site type describes a site.
 */
export function getSiteOutput(args: GetSiteOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSiteResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:networkmanager:getSite", {
        "globalNetworkId": args.globalNetworkId,
        "siteId": args.siteId,
    }, opts);
}

export interface GetSiteOutputArgs {
    /**
     * The ID of the global network.
     */
    globalNetworkId: pulumi.Input<string>;
    /**
     * The ID of the site.
     */
    siteId: pulumi.Input<string>;
}
