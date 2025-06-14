// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::FIS::ExperimentTemplate
 */
export function getExperimentTemplate(args: GetExperimentTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetExperimentTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:fis:getExperimentTemplate", {
        "id": args.id,
    }, opts);
}

export interface GetExperimentTemplateArgs {
    /**
     * The ID of the experiment template.
     */
    id: string;
}

export interface GetExperimentTemplateResult {
    /**
     * The actions for the experiment.
     */
    readonly actions?: {[key: string]: outputs.fis.ExperimentTemplateAction};
    /**
     * The description for the experiment template.
     */
    readonly description?: string;
    /**
     * The experiment options for an experiment template.
     */
    readonly experimentOptions?: outputs.fis.ExperimentTemplateExperimentOptions;
    /**
     * Describes the report configuration for the experiment template.
     */
    readonly experimentReportConfiguration?: outputs.fis.ExperimentTemplateExperimentReportConfiguration;
    /**
     * The ID of the experiment template.
     */
    readonly id?: string;
    /**
     * The configuration for experiment logging.
     */
    readonly logConfiguration?: outputs.fis.ExperimentTemplateLogConfiguration;
    /**
     * The Amazon Resource Name (ARN) of an IAM role.
     */
    readonly roleArn?: string;
    /**
     * The stop conditions for the experiment.
     */
    readonly stopConditions?: outputs.fis.ExperimentTemplateStopCondition[];
    /**
     * The tags for the experiment template.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The targets for the experiment.
     */
    readonly targets?: {[key: string]: outputs.fis.ExperimentTemplateTarget};
}
/**
 * Resource schema for AWS::FIS::ExperimentTemplate
 */
export function getExperimentTemplateOutput(args: GetExperimentTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetExperimentTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:fis:getExperimentTemplate", {
        "id": args.id,
    }, opts);
}

export interface GetExperimentTemplateOutputArgs {
    /**
     * The ID of the experiment template.
     */
    id: pulumi.Input<string>;
}
