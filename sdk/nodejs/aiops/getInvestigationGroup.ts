// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::AIOps::InvestigationGroup Resource Type
 */
export function getInvestigationGroup(args: GetInvestigationGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetInvestigationGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:aiops:getInvestigationGroup", {
        "arn": args.arn,
    }, opts);
}

export interface GetInvestigationGroupArgs {
    /**
     * The Amazon Resource Name (ARN) of the investigation group.
     */
    arn: string;
}

export interface GetInvestigationGroupResult {
    /**
     * The Amazon Resource Name (ARN) of the investigation group.
     */
    readonly arn?: string;
    /**
     * An array of key-value pairs of notification channels to apply to this resource.
     */
    readonly chatbotNotificationChannels?: outputs.aiops.InvestigationGroupChatbotNotificationChannel[];
    /**
     * The date and time that the investigation group was created.
     */
    readonly createdAt?: string;
    /**
     * The name of the user who created the investigation group.
     */
    readonly createdBy?: string;
    /**
     * An array of cross account configurations.
     */
    readonly crossAccountConfigurations?: outputs.aiops.InvestigationGroupCrossAccountConfiguration[];
    /**
     * Specifies the customer managed AWS KMS key that the investigation group uses to encrypt data, if there is one. If not, the investigation group uses an AWS key to encrypt the data.
     */
    readonly encryptionConfig?: outputs.aiops.InvestigationGroupEncryptionConfigMap;
    /**
     * Investigation Group policy
     */
    readonly investigationGroupPolicy?: string;
    /**
     * Flag to enable cloud trail history
     */
    readonly isCloudTrailEventHistoryEnabled?: boolean;
    /**
     * The date and time that the investigation group was most recently modified.
     */
    readonly lastModifiedAt?: string;
    /**
     * The name of the user who created the investigation group.
     */
    readonly lastModifiedBy?: string;
    /**
     * The ARN of the IAM role that the investigation group uses for permissions to gather data.
     */
    readonly roleArn?: string;
    /**
     * Displays the custom tag keys for custom applications in your system that you have specified in the investigation group. Resource tags help CloudWatch investigations narrow the search space when it is unable to discover definite relationships between resources.
     */
    readonly tagKeyBoundaries?: string[];
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Definition of AWS::AIOps::InvestigationGroup Resource Type
 */
export function getInvestigationGroupOutput(args: GetInvestigationGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInvestigationGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:aiops:getInvestigationGroup", {
        "arn": args.arn,
    }, opts);
}

export interface GetInvestigationGroupOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the investigation group.
     */
    arn: pulumi.Input<string>;
}
