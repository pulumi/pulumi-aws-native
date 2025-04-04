// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Config::ConfigurationAggregator
 */
export function getConfigurationAggregator(args: GetConfigurationAggregatorArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationAggregatorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:configuration:getConfigurationAggregator", {
        "configurationAggregatorName": args.configurationAggregatorName,
    }, opts);
}

export interface GetConfigurationAggregatorArgs {
    /**
     * The name of the aggregator.
     */
    configurationAggregatorName: string;
}

export interface GetConfigurationAggregatorResult {
    /**
     * Provides a list of source accounts and regions to be aggregated.
     */
    readonly accountAggregationSources?: outputs.configuration.ConfigurationAggregatorAccountAggregationSource[];
    /**
     * The Amazon Resource Name (ARN) of the aggregator.
     */
    readonly configurationAggregatorArn?: string;
    /**
     * Provides an organization and list of regions to be aggregated.
     */
    readonly organizationAggregationSource?: outputs.configuration.ConfigurationAggregatorOrganizationAggregationSource;
    /**
     * The tags for the configuration aggregator.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::Config::ConfigurationAggregator
 */
export function getConfigurationAggregatorOutput(args: GetConfigurationAggregatorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetConfigurationAggregatorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:configuration:getConfigurationAggregator", {
        "configurationAggregatorName": args.configurationAggregatorName,
    }, opts);
}

export interface GetConfigurationAggregatorOutputArgs {
    /**
     * The name of the aggregator.
     */
    configurationAggregatorName: pulumi.Input<string>;
}
