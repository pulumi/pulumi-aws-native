// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::Entitlement
 */
export function getEntitlement(args: GetEntitlementArgs, opts?: pulumi.InvokeOptions): Promise<GetEntitlementResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appstream:getEntitlement", {
        "name": args.name,
        "stackName": args.stackName,
    }, opts);
}

export interface GetEntitlementArgs {
    name: string;
    stackName: string;
}

export interface GetEntitlementResult {
    readonly appVisibility?: string;
    readonly attributes?: outputs.appstream.EntitlementAttribute[];
    readonly createdTime?: string;
    readonly description?: string;
    readonly lastModifiedTime?: string;
}

export function getEntitlementOutput(args: GetEntitlementOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEntitlementResult> {
    return pulumi.output(args).apply(a => getEntitlement(a, opts))
}

export interface GetEntitlementOutputArgs {
    name: pulumi.Input<string>;
    stackName: pulumi.Input<string>;
}