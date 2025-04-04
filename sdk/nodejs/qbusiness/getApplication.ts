// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::QBusiness::Application Resource Type
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:qbusiness:getApplication", {
        "applicationId": args.applicationId,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * The identifier for the Amazon Q Business application.
     */
    applicationId: string;
}

export interface GetApplicationResult {
    /**
     * The Amazon Resource Name (ARN) of the Amazon Q Business application.
     */
    readonly applicationArn?: string;
    /**
     * The identifier for the Amazon Q Business application.
     */
    readonly applicationId?: string;
    /**
     * Configuration information for the file upload during chat feature.
     */
    readonly attachmentsConfiguration?: outputs.qbusiness.ApplicationAttachmentsConfiguration;
    /**
     * Subscription configuration information for an Amazon Q Business application using IAM identity federation for user management.
     */
    readonly autoSubscriptionConfiguration?: outputs.qbusiness.ApplicationAutoSubscriptionConfiguration;
    /**
     * The Unix timestamp when the Amazon Q Business application was created.
     */
    readonly createdAt?: string;
    /**
     * A description for the Amazon Q Business application.
     */
    readonly description?: string;
    /**
     * The name of the Amazon Q Business application.
     */
    readonly displayName?: string;
    /**
     * The Amazon Resource Name (ARN) of the AWS IAM Identity Center instance attached to your Amazon Q Business application.
     */
    readonly identityCenterApplicationArn?: string;
    /**
     * Configuration information about chat response personalization. For more information, see [Personalizing chat responses](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/personalizing-chat-responses.html) .
     */
    readonly personalizationConfiguration?: outputs.qbusiness.ApplicationPersonalizationConfiguration;
    /**
     * Configuration information about Amazon Q Apps.
     */
    readonly qAppsConfiguration?: outputs.qbusiness.ApplicationQAppsConfiguration;
    /**
     * The Amazon Resource Name (ARN) of an IAM role with permissions to access your Amazon CloudWatch logs and metrics. If this property is not specified, Amazon Q Business will create a [service linked role (SLR)](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles.html#slr-permissions) and use it as the application's role.
     */
    readonly roleArn?: string;
    /**
     * The status of the Amazon Q Business application. The application is ready to use when the status is `ACTIVE` .
     */
    readonly status?: enums.qbusiness.ApplicationStatus;
    /**
     * A list of key-value pairs that identify or categorize your Amazon Q Business application. You can also use tags to help control access to the application. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The Unix timestamp when the Amazon Q Business application was last updated.
     */
    readonly updatedAt?: string;
}
/**
 * Definition of AWS::QBusiness::Application Resource Type
 */
export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:qbusiness:getApplication", {
        "applicationId": args.applicationId,
    }, opts);
}

export interface GetApplicationOutputArgs {
    /**
     * The identifier for the Amazon Q Business application.
     */
    applicationId: pulumi.Input<string>;
}
