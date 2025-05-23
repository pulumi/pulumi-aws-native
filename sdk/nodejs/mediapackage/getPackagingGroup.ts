// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaPackage::PackagingGroup
 */
export function getPackagingGroup(args: GetPackagingGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetPackagingGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:mediapackage:getPackagingGroup", {
        "id": args.id,
    }, opts);
}

export interface GetPackagingGroupArgs {
    /**
     * The ID of the PackagingGroup.
     */
    id: string;
}

export interface GetPackagingGroupResult {
    /**
     * The ARN of the PackagingGroup.
     */
    readonly arn?: string;
    /**
     * CDN Authorization
     */
    readonly authorization?: outputs.mediapackage.PackagingGroupAuthorization;
    /**
     * The fully qualified domain name for Assets in the PackagingGroup.
     */
    readonly domainName?: string;
    /**
     * The configuration parameters for egress access logging.
     */
    readonly egressAccessLogs?: outputs.mediapackage.PackagingGroupLogConfiguration;
}
/**
 * Resource schema for AWS::MediaPackage::PackagingGroup
 */
export function getPackagingGroupOutput(args: GetPackagingGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPackagingGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:mediapackage:getPackagingGroup", {
        "id": args.id,
    }, opts);
}

export interface GetPackagingGroupOutputArgs {
    /**
     * The ID of the PackagingGroup.
     */
    id: pulumi.Input<string>;
}
