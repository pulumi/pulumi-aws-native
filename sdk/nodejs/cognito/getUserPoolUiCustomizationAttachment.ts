// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment
 */
export function getUserPoolUiCustomizationAttachment(args: GetUserPoolUiCustomizationAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetUserPoolUiCustomizationAttachmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cognito:getUserPoolUiCustomizationAttachment", {
        "clientId": args.clientId,
        "userPoolId": args.userPoolId,
    }, opts);
}

export interface GetUserPoolUiCustomizationAttachmentArgs {
    /**
     * The app client ID for your UI customization. When this value isn't present, the customization applies to all user pool app clients that don't have client-level settings..
     */
    clientId: string;
    /**
     * The ID of the user pool where you want to apply branding to the classic hosted UI.
     */
    userPoolId: string;
}

export interface GetUserPoolUiCustomizationAttachmentResult {
    /**
     * A plaintext CSS file that contains the custom fields that you want to apply to your user pool or app client. To download a template, go to the Amazon Cognito console. Navigate to your user pool *App clients* tab, select *Login pages* , edit *Hosted UI (classic) style* , and select the link to `CSS template.css` .
     */
    readonly css?: string;
}
/**
 * Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment
 */
export function getUserPoolUiCustomizationAttachmentOutput(args: GetUserPoolUiCustomizationAttachmentOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUserPoolUiCustomizationAttachmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:cognito:getUserPoolUiCustomizationAttachment", {
        "clientId": args.clientId,
        "userPoolId": args.userPoolId,
    }, opts);
}

export interface GetUserPoolUiCustomizationAttachmentOutputArgs {
    /**
     * The app client ID for your UI customization. When this value isn't present, the customization applies to all user pool app clients that don't have client-level settings..
     */
    clientId: pulumi.Input<string>;
    /**
     * The ID of the user pool where you want to apply branding to the classic hosted UI.
     */
    userPoolId: pulumi.Input<string>;
}
