// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates an Amazon Kinesis Data Analytics application. For information about creating a Kinesis Data Analytics application, see [Creating an Application](https://docs.aws.amazon.com/kinesisanalytics/latest/java/getting-started.html).
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:kinesisanalyticsv2:getApplication", {
        "applicationName": args.applicationName,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * The name of the application.
     */
    applicationName: string;
}

export interface GetApplicationResult {
    /**
     * Use this parameter to configure the application.
     */
    readonly applicationConfiguration?: outputs.kinesisanalyticsv2.ApplicationConfiguration;
    /**
     * The description of the application.
     */
    readonly applicationDescription?: string;
    /**
     * Used to configure start of maintenance window.
     */
    readonly applicationMaintenanceConfiguration?: outputs.kinesisanalyticsv2.ApplicationMaintenanceConfiguration;
    /**
     * Specifies the IAM role that the application uses to access external resources.
     */
    readonly serviceExecutionRole?: string;
    /**
     * A list of one or more tags to assign to the application. A tag is a key-value pair that identifies an application. Note that the maximum number of application tags includes system tags. The maximum number of user-defined application tags is 50.
     */
    readonly tags?: outputs.kinesisanalyticsv2.ApplicationTag[];
}

export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationResult> {
    return pulumi.output(args).apply(a => getApplication(a, opts))
}

export interface GetApplicationOutputArgs {
    /**
     * The name of the application.
     */
    applicationName: pulumi.Input<string>;
}