// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SES::Template
 */
export function getTemplate(args: GetTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ses:getTemplate", {
        "id": args.id,
    }, opts);
}

export interface GetTemplateArgs {
    id: string;
}

export interface GetTemplateResult {
    readonly id?: string;
    /**
     * The content of the email, composed of a subject line and either an HTML part or a text-only part.
     */
    readonly template?: outputs.ses.Template;
}
/**
 * Resource Type definition for AWS::SES::Template
 */
export function getTemplateOutput(args: GetTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ses:getTemplate", {
        "id": args.id,
    }, opts);
}

export interface GetTemplateOutputArgs {
    id: pulumi.Input<string>;
}
