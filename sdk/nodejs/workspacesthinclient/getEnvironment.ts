// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource type definition for AWS::WorkSpacesThinClient::Environment.
 */
export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:workspacesthinclient:getEnvironment", {
        "id": args.id,
    }, opts);
}

export interface GetEnvironmentArgs {
    /**
     * Unique identifier of the environment.
     */
    id: string;
}

export interface GetEnvironmentResult {
    /**
     * Activation code for devices associated with environment.
     */
    readonly activationCode?: string;
    /**
     * The environment ARN.
     */
    readonly arn?: string;
    /**
     * The timestamp in unix epoch format when environment was created.
     */
    readonly createdAt?: string;
    /**
     * The ID of the software set to apply.
     */
    readonly desiredSoftwareSetId?: string;
    /**
     * The URL for the identity provider login (only for environments that use AppStream 2.0).
     */
    readonly desktopEndpoint?: string;
    /**
     * The type of VDI.
     */
    readonly desktopType?: enums.workspacesthinclient.EnvironmentDesktopType;
    /**
     * An array of key-value pairs to apply to the newly created devices for this environment.
     */
    readonly deviceCreationTags?: outputs.workspacesthinclient.EnvironmentTag[];
    /**
     * Unique identifier of the environment.
     */
    readonly id?: string;
    /**
     * A specification for a time window to apply software updates.
     */
    readonly maintenanceWindow?: outputs.workspacesthinclient.EnvironmentMaintenanceWindow;
    /**
     * The name of the environment.
     */
    readonly name?: string;
    /**
     * The ID of the software set that is pending to be installed.
     */
    readonly pendingSoftwareSetId?: string;
    /**
     * The version of the software set that is pending to be installed.
     */
    readonly pendingSoftwareSetVersion?: string;
    /**
     * Number of devices registered to the environment.
     */
    readonly registeredDevicesCount?: number;
    /**
     * Describes if the software currently installed on all devices in the environment is a supported version.
     */
    readonly softwareSetComplianceStatus?: enums.workspacesthinclient.EnvironmentSoftwareSetComplianceStatus;
    /**
     * An option to define which software updates to apply.
     */
    readonly softwareSetUpdateMode?: enums.workspacesthinclient.EnvironmentSoftwareSetUpdateMode;
    /**
     * An option to define if software updates should be applied within a maintenance window.
     */
    readonly softwareSetUpdateSchedule?: enums.workspacesthinclient.EnvironmentSoftwareSetUpdateSchedule;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The timestamp in unix epoch format when environment was last updated.
     */
    readonly updatedAt?: string;
}
/**
 * Resource type definition for AWS::WorkSpacesThinClient::Environment.
 */
export function getEnvironmentOutput(args: GetEnvironmentOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEnvironmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:workspacesthinclient:getEnvironment", {
        "id": args.id,
    }, opts);
}

export interface GetEnvironmentOutputArgs {
    /**
     * Unique identifier of the environment.
     */
    id: pulumi.Input<string>;
}
