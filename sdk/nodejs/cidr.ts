// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function cidr(args: CidrArgs, opts?: pulumi.InvokeOptions): Promise<CidrResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:index:cidr", {
        "cidrBits": args.cidrBits,
        "count": args.count,
        "ipBlock": args.ipBlock,
    }, opts);
}

export interface CidrArgs {
    cidrBits: number;
    count: number;
    ipBlock: string;
}

export interface CidrResult {
    readonly subnets: string[];
}

export function cidrOutput(args: CidrOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CidrResult> {
    return pulumi.output(args).apply(a => cidr(a, opts))
}

export interface CidrOutputArgs {
    cidrBits: pulumi.Input<number>;
    count: pulumi.Input<number>;
    ipBlock: pulumi.Input<string>;
}