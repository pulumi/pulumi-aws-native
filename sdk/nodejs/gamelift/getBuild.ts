// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GameLift::Build
 */
export function getBuild(args: GetBuildArgs, opts?: pulumi.InvokeOptions): Promise<GetBuildResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:gamelift:getBuild", {
        "buildId": args.buildId,
    }, opts);
}

export interface GetBuildArgs {
    /**
     * A unique identifier for a build to be deployed on the new fleet. If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created.
     */
    buildId: string;
}

export interface GetBuildResult {
    /**
     * A unique identifier for a build to be deployed on the new fleet. If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created.
     */
    readonly buildId?: string;
    /**
     * A descriptive label that is associated with a build. Build names do not need to be unique.
     */
    readonly name?: string;
    /**
     * Version information that is associated with this build. Version strings do not need to be unique.
     */
    readonly version?: string;
}
/**
 * Resource Type definition for AWS::GameLift::Build
 */
export function getBuildOutput(args: GetBuildOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBuildResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:gamelift:getBuild", {
        "buildId": args.buildId,
    }, opts);
}

export interface GetBuildOutputArgs {
    /**
     * A unique identifier for a build to be deployed on the new fleet. If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created.
     */
    buildId: pulumi.Input<string>;
}
