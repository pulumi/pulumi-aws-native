// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Deadline::LicenseEndpoint Resource Type
 */
export function getLicenseEndpoint(args: GetLicenseEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetLicenseEndpointResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:deadline:getLicenseEndpoint", {
        "arn": args.arn,
    }, opts);
}

export interface GetLicenseEndpointArgs {
    arn: string;
}

export interface GetLicenseEndpointResult {
    readonly arn?: string;
    readonly dnsName?: string;
    readonly licenseEndpointId?: string;
    readonly status?: enums.deadline.LicenseEndpointStatus;
    readonly statusMessage?: string;
}
/**
 * Definition of AWS::Deadline::LicenseEndpoint Resource Type
 */
export function getLicenseEndpointOutput(args: GetLicenseEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLicenseEndpointResult> {
    return pulumi.output(args).apply((a: any) => getLicenseEndpoint(a, opts))
}

export interface GetLicenseEndpointOutputArgs {
    arn: pulumi.Input<string>;
}