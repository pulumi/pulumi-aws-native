// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for Identity Center (SSO) Application
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sso:getApplication", {
        "applicationArn": args.applicationArn,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * The Application ARN that is returned upon creation of the Identity Center (SSO) Application
     */
    applicationArn: string;
}

export interface GetApplicationResult {
    /**
     * The Application ARN that is returned upon creation of the Identity Center (SSO) Application
     */
    readonly applicationArn?: string;
    /**
     * The description information for the Identity Center (SSO) Application
     */
    readonly description?: string;
    /**
     * The name you want to assign to this Identity Center (SSO) Application
     */
    readonly name?: string;
    /**
     * A structure that describes the options for the portal associated with an application
     */
    readonly portalOptions?: outputs.sso.ApplicationPortalOptionsConfiguration;
    /**
     * Specifies whether the application is enabled or disabled
     */
    readonly status?: enums.sso.ApplicationStatus;
    /**
     * Specifies tags to be attached to the application
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for Identity Center (SSO) Application
 */
export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:sso:getApplication", {
        "applicationArn": args.applicationArn,
    }, opts);
}

export interface GetApplicationOutputArgs {
    /**
     * The Application ARN that is returned upon creation of the Identity Center (SSO) Application
     */
    applicationArn: pulumi.Input<string>;
}
