// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DMS::InstanceProfile.
 */
export function getInstanceProfile(args: GetInstanceProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceProfileResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:dms:getInstanceProfile", {
        "instanceProfileArn": args.instanceProfileArn,
    }, opts);
}

export interface GetInstanceProfileArgs {
    /**
     * The property describes an ARN of the instance profile.
     */
    instanceProfileArn: string;
}

export interface GetInstanceProfileResult {
    /**
     * The property describes an availability zone of the instance profile.
     */
    readonly availabilityZone?: string;
    /**
     * The optional description of the instance profile.
     */
    readonly description?: string;
    /**
     * The property describes an ARN of the instance profile.
     */
    readonly instanceProfileArn?: string;
    /**
     * The property describes a creating time of the instance profile.
     */
    readonly instanceProfileCreationTime?: string;
    /**
     * The property describes a name for the instance profile.
     */
    readonly instanceProfileName?: string;
    /**
     * The property describes kms key arn for the instance profile.
     */
    readonly kmsKeyArn?: string;
    /**
     * The property describes a network type for the instance profile.
     */
    readonly networkType?: enums.dms.InstanceProfileNetworkType;
    /**
     * The property describes the publicly accessible of the instance profile
     */
    readonly publiclyAccessible?: boolean;
    /**
     * The property describes a subnet group identifier for the instance profile.
     */
    readonly subnetGroupIdentifier?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The property describes vps security groups for the instance profile.
     */
    readonly vpcSecurityGroups?: string[];
}
/**
 * Resource schema for AWS::DMS::InstanceProfile.
 */
export function getInstanceProfileOutput(args: GetInstanceProfileOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInstanceProfileResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:dms:getInstanceProfile", {
        "instanceProfileArn": args.instanceProfileArn,
    }, opts);
}

export interface GetInstanceProfileOutputArgs {
    /**
     * The property describes an ARN of the instance profile.
     */
    instanceProfileArn: pulumi.Input<string>;
}
