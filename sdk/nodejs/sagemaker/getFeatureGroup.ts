// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::FeatureGroup
 */
export function getFeatureGroup(args: GetFeatureGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sagemaker:getFeatureGroup", {
        "featureGroupName": args.featureGroupName,
    }, opts);
}

export interface GetFeatureGroupArgs {
    /**
     * The Name of the FeatureGroup.
     */
    featureGroupName: string;
}

export interface GetFeatureGroupResult {
    /**
     * A timestamp of FeatureGroup creation time.
     */
    readonly creationTime?: string;
    /**
     * An Array of Feature Definition
     */
    readonly featureDefinitions?: outputs.sagemaker.FeatureGroupFeatureDefinition[];
    /**
     * The status of the feature group.
     */
    readonly featureGroupStatus?: string;
    /**
     * The configuration of an `OnlineStore` .
     */
    readonly onlineStoreConfig?: outputs.sagemaker.OnlineStoreConfigProperties;
    /**
     * Used to set feature group throughput configuration. There are two modes: `ON_DEMAND` and `PROVISIONED` . With on-demand mode, you are charged for data reads and writes that your application performs on your feature group. You do not need to specify read and write throughput because Feature Store accommodates your workloads as they ramp up and down. You can switch a feature group to on-demand only once in a 24 hour period. With provisioned throughput mode, you specify the read and write capacity per second that you expect your application to require, and you are billed based on those limits. Exceeding provisioned throughput will result in your requests being throttled.
     *
     * Note: `PROVISIONED` throughput mode is supported only for feature groups that are offline-only, or use the [`Standard`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_OnlineStoreConfig.html#sagemaker-Type-OnlineStoreConfig-StorageType) tier online store.
     */
    readonly throughputConfig?: outputs.sagemaker.FeatureGroupThroughputConfig;
}
/**
 * Resource Type definition for AWS::SageMaker::FeatureGroup
 */
export function getFeatureGroupOutput(args: GetFeatureGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFeatureGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:sagemaker:getFeatureGroup", {
        "featureGroupName": args.featureGroupName,
    }, opts);
}

export interface GetFeatureGroupOutputArgs {
    /**
     * The Name of the FeatureGroup.
     */
    featureGroupName: pulumi.Input<string>;
}
