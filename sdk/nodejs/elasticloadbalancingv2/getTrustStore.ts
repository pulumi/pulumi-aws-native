// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElasticLoadBalancingV2::TrustStore
 */
export function getTrustStore(args: GetTrustStoreArgs, opts?: pulumi.InvokeOptions): Promise<GetTrustStoreResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:elasticloadbalancingv2:getTrustStore", {
        "trustStoreArn": args.trustStoreArn,
    }, opts);
}

export interface GetTrustStoreArgs {
    /**
     * The Amazon Resource Name (ARN) of the trust store.
     */
    trustStoreArn: string;
}

export interface GetTrustStoreResult {
    /**
     * The number of certificates associated with the trust store.
     */
    readonly numberOfCaCertificates?: number;
    /**
     * The status of the trust store, could be either of ACTIVE or CREATING.
     */
    readonly status?: string;
    /**
     * The tags to assign to the trust store.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The Amazon Resource Name (ARN) of the trust store.
     */
    readonly trustStoreArn?: string;
}
/**
 * Resource Type definition for AWS::ElasticLoadBalancingV2::TrustStore
 */
export function getTrustStoreOutput(args: GetTrustStoreOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTrustStoreResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:elasticloadbalancingv2:getTrustStore", {
        "trustStoreArn": args.trustStoreArn,
    }, opts);
}

export interface GetTrustStoreOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the trust store.
     */
    trustStoreArn: pulumi.Input<string>;
}
