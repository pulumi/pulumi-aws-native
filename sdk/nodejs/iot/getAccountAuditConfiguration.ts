// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Configures the Device Defender audit settings for this account. Settings include how audit notifications are sent and which audit checks are enabled or disabled.
 */
export function getAccountAuditConfiguration(args: GetAccountAuditConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountAuditConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iot:getAccountAuditConfiguration", {
        "accountId": args.accountId,
    }, opts);
}

export interface GetAccountAuditConfigurationArgs {
    /**
     * Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).
     */
    accountId: string;
}

export interface GetAccountAuditConfigurationResult {
    /**
     * Specifies which audit checks are enabled and disabled for this account.
     *
     * Some data collection might start immediately when certain checks are enabled. When a check is disabled, any data collected so far in relation to the check is deleted. To disable a check, set the value of the `Enabled:` key to `false` .
     *
     * If an enabled check is removed from the template, it will also be disabled.
     *
     * You can't disable a check if it's used by any scheduled audit. You must delete the check from the scheduled audit or delete the scheduled audit itself to disable the check.
     *
     * For more information on available audit checks see [AWS::IoT::AccountAuditConfiguration AuditCheckConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html)
     */
    readonly auditCheckConfigurations?: outputs.iot.AccountAuditConfigurationAuditCheckConfigurations;
    /**
     * Information about the targets to which audit notifications are sent.
     */
    readonly auditNotificationTargetConfigurations?: outputs.iot.AccountAuditConfigurationAuditNotificationTargetConfigurations;
    /**
     * The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.
     */
    readonly roleArn?: string;
}
/**
 * Configures the Device Defender audit settings for this account. Settings include how audit notifications are sent and which audit checks are enabled or disabled.
 */
export function getAccountAuditConfigurationOutput(args: GetAccountAuditConfigurationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAccountAuditConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:iot:getAccountAuditConfiguration", {
        "accountId": args.accountId,
    }, opts);
}

export interface GetAccountAuditConfigurationOutputArgs {
    /**
     * Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).
     */
    accountId: pulumi.Input<string>;
}
