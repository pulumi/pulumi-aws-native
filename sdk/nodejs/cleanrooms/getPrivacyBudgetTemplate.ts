// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a privacy budget within a collaboration
 */
export function getPrivacyBudgetTemplate(args: GetPrivacyBudgetTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivacyBudgetTemplateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cleanrooms:getPrivacyBudgetTemplate", {
        "membershipIdentifier": args.membershipIdentifier,
        "privacyBudgetTemplateIdentifier": args.privacyBudgetTemplateIdentifier,
    }, opts);
}

export interface GetPrivacyBudgetTemplateArgs {
    membershipIdentifier: string;
    privacyBudgetTemplateIdentifier: string;
}

export interface GetPrivacyBudgetTemplateResult {
    readonly arn?: string;
    readonly collaborationArn?: string;
    readonly collaborationIdentifier?: string;
    readonly membershipArn?: string;
    readonly parameters?: outputs.cleanrooms.ParametersProperties;
    readonly privacyBudgetTemplateIdentifier?: string;
    /**
     * An arbitrary set of tags (key-value pairs) for this cleanrooms privacy budget template.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Represents a privacy budget within a collaboration
 */
export function getPrivacyBudgetTemplateOutput(args: GetPrivacyBudgetTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrivacyBudgetTemplateResult> {
    return pulumi.output(args).apply((a: any) => getPrivacyBudgetTemplate(a, opts))
}

export interface GetPrivacyBudgetTemplateOutputArgs {
    membershipIdentifier: pulumi.Input<string>;
    privacyBudgetTemplateIdentifier: pulumi.Input<string>;
}