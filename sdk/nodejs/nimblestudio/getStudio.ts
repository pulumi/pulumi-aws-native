// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::NimbleStudio::Studio
 */
export function getStudio(args: GetStudioArgs, opts?: pulumi.InvokeOptions): Promise<GetStudioResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:nimblestudio:getStudio", {
        "studioId": args.studioId,
    }, opts);
}

export interface GetStudioArgs {
    /**
     * The unique identifier for the studio resource.
     */
    studioId: string;
}

export interface GetStudioResult {
    /**
     * The IAM role that studio admins assume when logging in to the Nimble Studio portal.
     */
    readonly adminRoleArn?: string;
    /**
     * A friendly name for the studio.
     */
    readonly displayName?: string;
    /**
     * The AWS Region where the studio resource is located. For example, `us-west-2` .
     */
    readonly homeRegion?: string;
    /**
     * The IAM Identity Center application client ID that is used to integrate with IAM Identity Center , which enables IAM Identity Center users to log into the  portal.
     */
    readonly ssoClientId?: string;
    /**
     * Configuration of the encryption method that is used for the studio.
     */
    readonly studioEncryptionConfiguration?: outputs.nimblestudio.StudioEncryptionConfiguration;
    /**
     * The unique identifier for the studio resource.
     */
    readonly studioId?: string;
    /**
     * The unique identifier for the studio resource.
     */
    readonly studioUrl?: string;
    /**
     * The IAM role that studio users assume when logging in to the Nimble Studio portal.
     */
    readonly userRoleArn?: string;
}
/**
 * Resource Type definition for AWS::NimbleStudio::Studio
 */
export function getStudioOutput(args: GetStudioOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetStudioResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:nimblestudio:getStudio", {
        "studioId": args.studioId,
    }, opts);
}

export interface GetStudioOutputArgs {
    /**
     * The unique identifier for the studio resource.
     */
    studioId: pulumi.Input<string>;
}
