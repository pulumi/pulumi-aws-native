// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::ElasticBeanstalk::Application resource specifies an Elastic Beanstalk application.
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:elasticbeanstalk:getApplication", {
        "applicationName": args.applicationName,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
     */
    applicationName: string;
}

export interface GetApplicationResult {
    /**
     * Your description of the application.
     */
    readonly description?: string;
    /**
     * Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
     */
    readonly resourceLifecycleConfig?: outputs.elasticbeanstalk.ApplicationResourceLifecycleConfig;
}

export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationResult> {
    return pulumi.output(args).apply(a => getApplication(a, opts))
}

export interface GetApplicationOutputArgs {
    /**
     * A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
     */
    applicationName: pulumi.Input<string>;
}