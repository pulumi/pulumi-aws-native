// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::SecurityHub::ProductSubscription resource represents a subscription to a service that is allowed to generate findings for your Security Hub account. One product subscription resource is created for each product enabled.
 */
export function getProductSubscription(args: GetProductSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetProductSubscriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:securityhub:getProductSubscription", {
        "productSubscriptionArn": args.productSubscriptionArn,
    }, opts);
}

export interface GetProductSubscriptionArgs {
    /**
     * The ARN of the product subscription for the account
     */
    productSubscriptionArn: string;
}

export interface GetProductSubscriptionResult {
    /**
     * The ARN of the product subscription for the account
     */
    readonly productSubscriptionArn?: string;
}
/**
 * The AWS::SecurityHub::ProductSubscription resource represents a subscription to a service that is allowed to generate findings for your Security Hub account. One product subscription resource is created for each product enabled.
 */
export function getProductSubscriptionOutput(args: GetProductSubscriptionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetProductSubscriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:securityhub:getProductSubscription", {
        "productSubscriptionArn": args.productSubscriptionArn,
    }, opts);
}

export interface GetProductSubscriptionOutputArgs {
    /**
     * The ARN of the product subscription for the account
     */
    productSubscriptionArn: pulumi.Input<string>;
}
