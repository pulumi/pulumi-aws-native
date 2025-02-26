// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * resource definition
 */
export function getSoftwarePackage(args: GetSoftwarePackageArgs, opts?: pulumi.InvokeOptions): Promise<GetSoftwarePackageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iot:getSoftwarePackage", {
        "packageName": args.packageName,
    }, opts);
}

export interface GetSoftwarePackageArgs {
    packageName: string;
}

export interface GetSoftwarePackageResult {
    readonly description?: string;
    /**
     * The Amazon Resource Name (ARN) for the package.
     */
    readonly packageArn?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * resource definition
 */
export function getSoftwarePackageOutput(args: GetSoftwarePackageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSoftwarePackageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:iot:getSoftwarePackage", {
        "packageName": args.packageName,
    }, opts);
}

export interface GetSoftwarePackageOutputArgs {
    packageName: pulumi.Input<string>;
}
