// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ImageBuilder::InfrastructureConfiguration
 */
export function getInfrastructureConfiguration(args: GetInfrastructureConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetInfrastructureConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:imagebuilder:getInfrastructureConfiguration", {
        "arn": args.arn,
    }, opts);
}

export interface GetInfrastructureConfigurationArgs {
    /**
     * The Amazon Resource Name (ARN) of the infrastructure configuration.
     */
    arn: string;
}

export interface GetInfrastructureConfigurationResult {
    /**
     * The Amazon Resource Name (ARN) of the infrastructure configuration.
     */
    readonly arn?: string;
    /**
     * The description of the infrastructure configuration.
     */
    readonly description?: string;
    /**
     * The instance metadata option settings for the infrastructure configuration.
     */
    readonly instanceMetadataOptions?: outputs.imagebuilder.InfrastructureConfigurationInstanceMetadataOptions;
    /**
     * The instance profile of the infrastructure configuration.
     */
    readonly instanceProfileName?: string;
    /**
     * The instance types of the infrastructure configuration.
     */
    readonly instanceTypes?: string[];
    /**
     * The EC2 key pair of the infrastructure configuration..
     */
    readonly keyPair?: string;
    /**
     * The logging configuration of the infrastructure configuration.
     */
    readonly logging?: outputs.imagebuilder.InfrastructureConfigurationLogging;
    /**
     * The placement option settings for the infrastructure configuration.
     */
    readonly placement?: outputs.imagebuilder.InfrastructureConfigurationPlacement;
    /**
     * The tags attached to the resource created by Image Builder.
     */
    readonly resourceTags?: {[key: string]: string};
    /**
     * The security group IDs of the infrastructure configuration.
     */
    readonly securityGroupIds?: string[];
    /**
     * The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.
     */
    readonly snsTopicArn?: string;
    /**
     * The subnet ID of the infrastructure configuration.
     */
    readonly subnetId?: string;
    /**
     * The tags associated with the component.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The terminate instance on failure configuration of the infrastructure configuration.
     */
    readonly terminateInstanceOnFailure?: boolean;
}
/**
 * Resource schema for AWS::ImageBuilder::InfrastructureConfiguration
 */
export function getInfrastructureConfigurationOutput(args: GetInfrastructureConfigurationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInfrastructureConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:imagebuilder:getInfrastructureConfiguration", {
        "arn": args.arn,
    }, opts);
}

export interface GetInfrastructureConfigurationOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the infrastructure configuration.
     */
    arn: pulumi.Input<string>;
}
