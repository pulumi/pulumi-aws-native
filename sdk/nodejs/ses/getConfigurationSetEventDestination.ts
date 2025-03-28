// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SES::ConfigurationSetEventDestination
 */
export function getConfigurationSetEventDestination(args: GetConfigurationSetEventDestinationArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationSetEventDestinationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ses:getConfigurationSetEventDestination", {
        "id": args.id,
    }, opts);
}

export interface GetConfigurationSetEventDestinationArgs {
    id: string;
}

export interface GetConfigurationSetEventDestinationResult {
    /**
     * The event destination object.
     */
    readonly eventDestination?: outputs.ses.ConfigurationSetEventDestinationEventDestination;
    readonly id?: string;
}
/**
 * Resource Type definition for AWS::SES::ConfigurationSetEventDestination
 */
export function getConfigurationSetEventDestinationOutput(args: GetConfigurationSetEventDestinationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetConfigurationSetEventDestinationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ses:getConfigurationSetEventDestination", {
        "id": args.id,
    }, opts);
}

export interface GetConfigurationSetEventDestinationOutputArgs {
    id: pulumi.Input<string>;
}
