// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates a package and storage location in an Amazon S3 access point.
 */
export function getPackage(args: GetPackageArgs, opts?: pulumi.InvokeOptions): Promise<GetPackageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:panorama:getPackage", {
        "packageId": args.packageId,
    }, opts);
}

export interface GetPackageArgs {
    /**
     * The package's ID.
     */
    packageId: string;
}

export interface GetPackageResult {
    /**
     * The package's ARN.
     */
    readonly arn?: string;
    /**
     * When the package was created.
     */
    readonly createdTime?: number;
    /**
     * The package's ID.
     */
    readonly packageId?: string;
    /**
     * A storage location.
     */
    readonly storageLocation?: outputs.panorama.PackageStorageLocation;
    /**
     * Tags for the package.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Creates a package and storage location in an Amazon S3 access point.
 */
export function getPackageOutput(args: GetPackageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPackageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:panorama:getPackage", {
        "packageId": args.packageId,
    }, opts);
}

export interface GetPackageOutputArgs {
    /**
     * The package's ID.
     */
    packageId: pulumi.Input<string>;
}
