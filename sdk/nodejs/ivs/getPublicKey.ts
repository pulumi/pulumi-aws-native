// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IVS::PublicKey
 */
export function getPublicKey(args: GetPublicKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicKeyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ivs:getPublicKey", {
        "arn": args.arn,
    }, opts);
}

export interface GetPublicKeyArgs {
    /**
     * Key-pair identifier.
     */
    arn: string;
}

export interface GetPublicKeyResult {
    /**
     * Key-pair identifier.
     */
    readonly arn?: string;
    /**
     * Key-pair identifier.
     */
    readonly fingerprint?: string;
    /**
     * A list of key-value pairs that contain metadata for the asset model.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::IVS::PublicKey
 */
export function getPublicKeyOutput(args: GetPublicKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublicKeyResult> {
    return pulumi.output(args).apply((a: any) => getPublicKey(a, opts))
}

export interface GetPublicKeyOutputArgs {
    /**
     * Key-pair identifier.
     */
    arn: pulumi.Input<string>;
}