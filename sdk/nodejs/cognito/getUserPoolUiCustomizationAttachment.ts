// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
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
    clientId: string;
    userPoolId: string;
}

export interface GetUserPoolUiCustomizationAttachmentResult {
    readonly css?: string;
}
/**
 * Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment
 */
export function getUserPoolUiCustomizationAttachmentOutput(args: GetUserPoolUiCustomizationAttachmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserPoolUiCustomizationAttachmentResult> {
    return pulumi.output(args).apply((a: any) => getUserPoolUiCustomizationAttachment(a, opts))
}

export interface GetUserPoolUiCustomizationAttachmentOutputArgs {
    clientId: pulumi.Input<string>;
    userPoolId: pulumi.Input<string>;
}