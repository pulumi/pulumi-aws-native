// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElasticBeanstalk::ConfigurationTemplate
 */
export function getConfigurationTemplate(args: GetConfigurationTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationTemplateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:elasticbeanstalk:getConfigurationTemplate", {
        "applicationName": args.applicationName,
        "templateName": args.templateName,
    }, opts);
}

export interface GetConfigurationTemplateArgs {
    /**
     * The name of the Elastic Beanstalk application to associate with this configuration template. 
     */
    applicationName: string;
    /**
     * The name of the configuration template
     */
    templateName: string;
}

export interface GetConfigurationTemplateResult {
    /**
     * An optional description for this configuration.
     */
    readonly description?: string;
    /**
     * Option values for the Elastic Beanstalk configuration, such as the instance type. If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the AWS Elastic Beanstalk Developer Guide. 
     */
    readonly optionSettings?: outputs.elasticbeanstalk.ConfigurationTemplateConfigurationOptionSetting[];
    /**
     * The name of the configuration template
     */
    readonly templateName?: string;
}

export function getConfigurationTemplateOutput(args: GetConfigurationTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConfigurationTemplateResult> {
    return pulumi.output(args).apply(a => getConfigurationTemplate(a, opts))
}

export interface GetConfigurationTemplateOutputArgs {
    /**
     * The name of the Elastic Beanstalk application to associate with this configuration template. 
     */
    applicationName: pulumi.Input<string>;
    /**
     * The name of the configuration template
     */
    templateName: pulumi.Input<string>;
}