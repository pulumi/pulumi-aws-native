// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A service network is a logical boundary for a collection of services. You can associate services and VPCs with a service network.
 */
export function getServiceNetwork(args: GetServiceNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceNetworkResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:vpclattice:getServiceNetwork", {
        "arn": args.arn,
    }, opts);
}

export interface GetServiceNetworkArgs {
    arn: string;
}

export interface GetServiceNetworkResult {
    readonly arn?: string;
    readonly authType?: enums.vpclattice.ServiceNetworkAuthType;
    readonly createdAt?: string;
    readonly id?: string;
    readonly lastUpdatedAt?: string;
    readonly tags?: outputs.vpclattice.ServiceNetworkTag[];
}
/**
 * A service network is a logical boundary for a collection of services. You can associate services and VPCs with a service network.
 */
export function getServiceNetworkOutput(args: GetServiceNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceNetworkResult> {
    return pulumi.output(args).apply((a: any) => getServiceNetwork(a, opts))
}

export interface GetServiceNetworkOutputArgs {
    arn: pulumi.Input<string>;
}