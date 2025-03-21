// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SecurityLake::AwsLogSource
 */
export function getAwsLogSource(args: GetAwsLogSourceArgs, opts?: pulumi.InvokeOptions): Promise<GetAwsLogSourceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:securitylake:getAwsLogSource", {
        "sourceName": args.sourceName,
        "sourceVersion": args.sourceVersion,
    }, opts);
}

export interface GetAwsLogSourceArgs {
    /**
     * The name for a AWS source. This must be a Regionally unique value.
     */
    sourceName: string;
    /**
     * The version for a AWS source. This must be a Regionally unique value.
     */
    sourceVersion: string;
}

export interface GetAwsLogSourceResult {
    /**
     * AWS account where you want to collect logs from.
     */
    readonly accounts?: string[];
}
/**
 * Resource Type definition for AWS::SecurityLake::AwsLogSource
 */
export function getAwsLogSourceOutput(args: GetAwsLogSourceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAwsLogSourceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:securitylake:getAwsLogSource", {
        "sourceName": args.sourceName,
        "sourceVersion": args.sourceVersion,
    }, opts);
}

export interface GetAwsLogSourceOutputArgs {
    /**
     * The name for a AWS source. This must be a Regionally unique value.
     */
    sourceName: pulumi.Input<string>;
    /**
     * The version for a AWS source. This must be a Regionally unique value.
     */
    sourceVersion: pulumi.Input<string>;
}
