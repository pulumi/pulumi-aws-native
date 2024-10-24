// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Cognito::UserPool Resource Type
 */
export class UserPool extends pulumi.CustomResource {
    /**
     * Get an existing UserPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UserPool {
        return new UserPool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cognito:UserPool';

    /**
     * Returns true if the given object is an instance of UserPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPool.__pulumiType;
    }

    /**
     * The available verified method a user can use to recover their password when they call `ForgotPassword` . You can use this setting to define a preferred method when a user has more than one method available. With this setting, SMS doesn't qualify for a valid password recovery mechanism if the user also has SMS multi-factor authentication (MFA) activated. In the absence of this setting, Amazon Cognito uses the legacy behavior to determine the recovery method where SMS is preferred through email.
     */
    public readonly accountRecoverySetting!: pulumi.Output<outputs.cognito.UserPoolAccountRecoverySetting | undefined>;
    /**
     * The settings for administrator creation of users in a user pool. Contains settings for allowing user sign-up, customizing invitation messages to new users, and the amount of time before temporary passwords expire.
     *
     * This data type is a request and response parameter of [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) and [UpdateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UpdateUserPool.html) , and a response parameter of [DescribeUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html) .
     */
    public readonly adminCreateUserConfig!: pulumi.Output<outputs.cognito.UserPoolAdminCreateUserConfig | undefined>;
    /**
     * Attributes supported as an alias for this user pool. Possible values: *phone_number* , *email* , or *preferred_username* .
     */
    public readonly aliasAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the user pool, such as `arn:aws:cognito-idp:us-east-1:123412341234:userpool/us-east-1_123412341` .
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The attributes to be auto-verified. Possible values: *email* , *phone_number* .
     */
    public readonly autoVerifiedAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * When active, `DeletionProtection` prevents accidental deletion of your user
     * pool. Before you can delete a user pool that you have protected against deletion, you
     * must deactivate this feature.
     *
     * When you try to delete a protected user pool in a `DeleteUserPool` API request, Amazon Cognito returns an `InvalidParameterException` error. To delete a protected user pool, send a new `DeleteUserPool` request after you deactivate deletion protection in an `UpdateUserPool` API request.
     */
    public readonly deletionProtection!: pulumi.Output<string | undefined>;
    /**
     * The device-remembering configuration for a user pool. A null value indicates that you have deactivated device remembering in your user pool.
     *
     * > When you provide a value for any `DeviceConfiguration` field, you activate the Amazon Cognito device-remembering feature.
     */
    public readonly deviceConfiguration!: pulumi.Output<outputs.cognito.UserPoolDeviceConfiguration | undefined>;
    public readonly emailAuthenticationMessage!: pulumi.Output<string | undefined>;
    public readonly emailAuthenticationSubject!: pulumi.Output<string | undefined>;
    /**
     * The email configuration of your user pool. The email configuration type sets your preferred sending method, AWS Region, and sender for messages from your user pool.
     */
    public readonly emailConfiguration!: pulumi.Output<outputs.cognito.UserPoolEmailConfiguration | undefined>;
    /**
     * This parameter is no longer used. See [VerificationMessageTemplateType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-verificationmessagetemplate.html) .
     */
    public readonly emailVerificationMessage!: pulumi.Output<string | undefined>;
    /**
     * This parameter is no longer used. See [VerificationMessageTemplateType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-verificationmessagetemplate.html) .
     */
    public readonly emailVerificationSubject!: pulumi.Output<string | undefined>;
    /**
     * Set enabled MFA options on a specified user pool. To disable all MFAs after it has been enabled, set `MfaConfiguration` to `OFF` and remove EnabledMfas. MFAs can only be all disabled if `MfaConfiguration` is `OFF` . After you enable `SMS_MFA` , you can only disable it by setting `MfaConfiguration` to `OFF` . Can be one of the following values:
     *
     * - `SMS_MFA` - Enables MFA with SMS for the user pool. To select this option, you must also provide values for `SmsConfiguration` .
     * - `SOFTWARE_TOKEN_MFA` - Enables software token MFA for the user pool.
     * - `EMAIL_OTP` - Enables MFA with email for the user pool. To select this option, you must provide values for `EmailConfiguration` and within those, set `EmailSendingAccount` to `DEVELOPER` .
     *
     * Allowed values: `SMS_MFA` | `SOFTWARE_TOKEN_MFA` | `EMAIL_OTP`
     */
    public readonly enabledMfas!: pulumi.Output<string[] | undefined>;
    /**
     * A collection of user pool Lambda triggers. Amazon Cognito invokes triggers at several possible stages of authentication operations. Triggers can modify the outcome of the operations that invoked them.
     */
    public readonly lambdaConfig!: pulumi.Output<outputs.cognito.UserPoolLambdaConfig | undefined>;
    /**
     * The multi-factor authentication (MFA) configuration. Valid values include:
     *
     * - `OFF` MFA won't be used for any users.
     * - `ON` MFA is required for all users to sign in.
     * - `OPTIONAL` MFA will be required only for individual users who have an MFA factor activated.
     */
    public readonly mfaConfiguration!: pulumi.Output<string | undefined>;
    /**
     * A list of user pool policies. Contains the policy that sets password-complexity requirements.
     *
     * This data type is a request and response parameter of [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) and [UpdateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UpdateUserPool.html) , and a response parameter of [DescribeUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html) .
     */
    public readonly policies!: pulumi.Output<outputs.cognito.UserPoolPolicies | undefined>;
    /**
     * A friendly name for the IdP.
     */
    public /*out*/ readonly providerName!: pulumi.Output<string>;
    /**
     * The URL of the provider of the Amazon Cognito user pool, specified as a `String` .
     */
    public /*out*/ readonly providerUrl!: pulumi.Output<string>;
    /**
     * An array of schema attributes for the new user pool. These attributes can be standard or custom attributes.
     */
    public readonly schema!: pulumi.Output<outputs.cognito.UserPoolSchemaAttribute[] | undefined>;
    /**
     * The contents of the SMS authentication message.
     */
    public readonly smsAuthenticationMessage!: pulumi.Output<string | undefined>;
    /**
     * The SMS configuration with the settings that your Amazon Cognito user pool must use to send an SMS message from your AWS account through Amazon Simple Notification Service. To send SMS messages with Amazon SNS in the AWS Region that you want, the Amazon Cognito user pool uses an AWS Identity and Access Management (IAM) role in your AWS account .
     */
    public readonly smsConfiguration!: pulumi.Output<outputs.cognito.UserPoolSmsConfiguration | undefined>;
    /**
     * This parameter is no longer used. See [VerificationMessageTemplateType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-verificationmessagetemplate.html) .
     */
    public readonly smsVerificationMessage!: pulumi.Output<string | undefined>;
    /**
     * The settings for updates to user attributes. These settings include the property `AttributesRequireVerificationBeforeUpdate` ,
     * a user-pool setting that tells Amazon Cognito how to handle changes to the value of your users' email address and phone number attributes. For
     * more information, see [Verifying updates to email addresses and phone numbers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html#user-pool-settings-verifications-verify-attribute-updates) .
     */
    public readonly userAttributeUpdateSettings!: pulumi.Output<outputs.cognito.UserPoolUserAttributeUpdateSettings | undefined>;
    /**
     * User pool add-ons. Contains settings for activation of advanced security features. To log user security information but take no action, set to `AUDIT` . To configure automatic security responses to risky traffic to your user pool, set to `ENFORCED` .
     *
     * For more information, see [Adding advanced security to a user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-advanced-security.html) .
     */
    public readonly userPoolAddOns!: pulumi.Output<outputs.cognito.UserPoolAddOns | undefined>;
    /**
     * The ID of the user pool.
     */
    public /*out*/ readonly userPoolId!: pulumi.Output<string>;
    /**
     * A string used to name the user pool.
     */
    public readonly userPoolName!: pulumi.Output<string | undefined>;
    /**
     * The tag keys and values to assign to the user pool. A tag is a label that you can use to categorize and manage user pools in different ways, such as by purpose, owner, environment, or other criteria.
     */
    public readonly userPoolTags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies whether a user can use an email address or phone number as a username when they sign up.
     */
    public readonly usernameAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * Case sensitivity on the username input for the selected sign-in option. When case sensitivity is set to `False` (case insensitive), users can sign in with any combination of capital and lowercase letters. For example, `username` , `USERNAME` , or `UserName` , or for email, `email@example.com` or `EMaiL@eXamplE.Com` . For most use cases, set case sensitivity to `False` (case insensitive) as a best practice. When usernames and email addresses are case insensitive, Amazon Cognito treats any variation in case as the same user, and prevents a case variation from being assigned to the same attribute for a different user.
     *
     * This configuration is immutable after you set it. For more information, see [UsernameConfigurationType](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UsernameConfigurationType.html) .
     */
    public readonly usernameConfiguration!: pulumi.Output<outputs.cognito.UserPoolUsernameConfiguration | undefined>;
    /**
     * The template for the verification message that your user pool delivers to users who set an email address or phone number attribute.
     *
     * Set the email message type that corresponds to your `DefaultEmailOption` selection. For `CONFIRM_WITH_LINK` , specify an `EmailMessageByLink` and leave `EmailMessage` blank. For `CONFIRM_WITH_CODE` , specify an `EmailMessage` and leave `EmailMessageByLink` blank. When you supply both parameters with either choice, Amazon Cognito returns an error.
     */
    public readonly verificationMessageTemplate!: pulumi.Output<outputs.cognito.UserPoolVerificationMessageTemplate | undefined>;

    /**
     * Create a UserPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserPoolArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["accountRecoverySetting"] = args ? args.accountRecoverySetting : undefined;
            resourceInputs["adminCreateUserConfig"] = args ? args.adminCreateUserConfig : undefined;
            resourceInputs["aliasAttributes"] = args ? args.aliasAttributes : undefined;
            resourceInputs["autoVerifiedAttributes"] = args ? args.autoVerifiedAttributes : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["deviceConfiguration"] = args ? args.deviceConfiguration : undefined;
            resourceInputs["emailAuthenticationMessage"] = args ? args.emailAuthenticationMessage : undefined;
            resourceInputs["emailAuthenticationSubject"] = args ? args.emailAuthenticationSubject : undefined;
            resourceInputs["emailConfiguration"] = args ? args.emailConfiguration : undefined;
            resourceInputs["emailVerificationMessage"] = args ? args.emailVerificationMessage : undefined;
            resourceInputs["emailVerificationSubject"] = args ? args.emailVerificationSubject : undefined;
            resourceInputs["enabledMfas"] = args ? args.enabledMfas : undefined;
            resourceInputs["lambdaConfig"] = args ? args.lambdaConfig : undefined;
            resourceInputs["mfaConfiguration"] = args ? args.mfaConfiguration : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["schema"] = args ? args.schema : undefined;
            resourceInputs["smsAuthenticationMessage"] = args ? args.smsAuthenticationMessage : undefined;
            resourceInputs["smsConfiguration"] = args ? args.smsConfiguration : undefined;
            resourceInputs["smsVerificationMessage"] = args ? args.smsVerificationMessage : undefined;
            resourceInputs["userAttributeUpdateSettings"] = args ? args.userAttributeUpdateSettings : undefined;
            resourceInputs["userPoolAddOns"] = args ? args.userPoolAddOns : undefined;
            resourceInputs["userPoolName"] = args ? args.userPoolName : undefined;
            resourceInputs["userPoolTags"] = args ? args.userPoolTags : undefined;
            resourceInputs["usernameAttributes"] = args ? args.usernameAttributes : undefined;
            resourceInputs["usernameConfiguration"] = args ? args.usernameConfiguration : undefined;
            resourceInputs["verificationMessageTemplate"] = args ? args.verificationMessageTemplate : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["providerName"] = undefined /*out*/;
            resourceInputs["providerUrl"] = undefined /*out*/;
            resourceInputs["userPoolId"] = undefined /*out*/;
        } else {
            resourceInputs["accountRecoverySetting"] = undefined /*out*/;
            resourceInputs["adminCreateUserConfig"] = undefined /*out*/;
            resourceInputs["aliasAttributes"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["autoVerifiedAttributes"] = undefined /*out*/;
            resourceInputs["deletionProtection"] = undefined /*out*/;
            resourceInputs["deviceConfiguration"] = undefined /*out*/;
            resourceInputs["emailAuthenticationMessage"] = undefined /*out*/;
            resourceInputs["emailAuthenticationSubject"] = undefined /*out*/;
            resourceInputs["emailConfiguration"] = undefined /*out*/;
            resourceInputs["emailVerificationMessage"] = undefined /*out*/;
            resourceInputs["emailVerificationSubject"] = undefined /*out*/;
            resourceInputs["enabledMfas"] = undefined /*out*/;
            resourceInputs["lambdaConfig"] = undefined /*out*/;
            resourceInputs["mfaConfiguration"] = undefined /*out*/;
            resourceInputs["policies"] = undefined /*out*/;
            resourceInputs["providerName"] = undefined /*out*/;
            resourceInputs["providerUrl"] = undefined /*out*/;
            resourceInputs["schema"] = undefined /*out*/;
            resourceInputs["smsAuthenticationMessage"] = undefined /*out*/;
            resourceInputs["smsConfiguration"] = undefined /*out*/;
            resourceInputs["smsVerificationMessage"] = undefined /*out*/;
            resourceInputs["userAttributeUpdateSettings"] = undefined /*out*/;
            resourceInputs["userPoolAddOns"] = undefined /*out*/;
            resourceInputs["userPoolId"] = undefined /*out*/;
            resourceInputs["userPoolName"] = undefined /*out*/;
            resourceInputs["userPoolTags"] = undefined /*out*/;
            resourceInputs["usernameAttributes"] = undefined /*out*/;
            resourceInputs["usernameConfiguration"] = undefined /*out*/;
            resourceInputs["verificationMessageTemplate"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a UserPool resource.
 */
export interface UserPoolArgs {
    /**
     * The available verified method a user can use to recover their password when they call `ForgotPassword` . You can use this setting to define a preferred method when a user has more than one method available. With this setting, SMS doesn't qualify for a valid password recovery mechanism if the user also has SMS multi-factor authentication (MFA) activated. In the absence of this setting, Amazon Cognito uses the legacy behavior to determine the recovery method where SMS is preferred through email.
     */
    accountRecoverySetting?: pulumi.Input<inputs.cognito.UserPoolAccountRecoverySettingArgs>;
    /**
     * The settings for administrator creation of users in a user pool. Contains settings for allowing user sign-up, customizing invitation messages to new users, and the amount of time before temporary passwords expire.
     *
     * This data type is a request and response parameter of [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) and [UpdateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UpdateUserPool.html) , and a response parameter of [DescribeUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html) .
     */
    adminCreateUserConfig?: pulumi.Input<inputs.cognito.UserPoolAdminCreateUserConfigArgs>;
    /**
     * Attributes supported as an alias for this user pool. Possible values: *phone_number* , *email* , or *preferred_username* .
     */
    aliasAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The attributes to be auto-verified. Possible values: *email* , *phone_number* .
     */
    autoVerifiedAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When active, `DeletionProtection` prevents accidental deletion of your user
     * pool. Before you can delete a user pool that you have protected against deletion, you
     * must deactivate this feature.
     *
     * When you try to delete a protected user pool in a `DeleteUserPool` API request, Amazon Cognito returns an `InvalidParameterException` error. To delete a protected user pool, send a new `DeleteUserPool` request after you deactivate deletion protection in an `UpdateUserPool` API request.
     */
    deletionProtection?: pulumi.Input<string>;
    /**
     * The device-remembering configuration for a user pool. A null value indicates that you have deactivated device remembering in your user pool.
     *
     * > When you provide a value for any `DeviceConfiguration` field, you activate the Amazon Cognito device-remembering feature.
     */
    deviceConfiguration?: pulumi.Input<inputs.cognito.UserPoolDeviceConfigurationArgs>;
    emailAuthenticationMessage?: pulumi.Input<string>;
    emailAuthenticationSubject?: pulumi.Input<string>;
    /**
     * The email configuration of your user pool. The email configuration type sets your preferred sending method, AWS Region, and sender for messages from your user pool.
     */
    emailConfiguration?: pulumi.Input<inputs.cognito.UserPoolEmailConfigurationArgs>;
    /**
     * This parameter is no longer used. See [VerificationMessageTemplateType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-verificationmessagetemplate.html) .
     */
    emailVerificationMessage?: pulumi.Input<string>;
    /**
     * This parameter is no longer used. See [VerificationMessageTemplateType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-verificationmessagetemplate.html) .
     */
    emailVerificationSubject?: pulumi.Input<string>;
    /**
     * Set enabled MFA options on a specified user pool. To disable all MFAs after it has been enabled, set `MfaConfiguration` to `OFF` and remove EnabledMfas. MFAs can only be all disabled if `MfaConfiguration` is `OFF` . After you enable `SMS_MFA` , you can only disable it by setting `MfaConfiguration` to `OFF` . Can be one of the following values:
     *
     * - `SMS_MFA` - Enables MFA with SMS for the user pool. To select this option, you must also provide values for `SmsConfiguration` .
     * - `SOFTWARE_TOKEN_MFA` - Enables software token MFA for the user pool.
     * - `EMAIL_OTP` - Enables MFA with email for the user pool. To select this option, you must provide values for `EmailConfiguration` and within those, set `EmailSendingAccount` to `DEVELOPER` .
     *
     * Allowed values: `SMS_MFA` | `SOFTWARE_TOKEN_MFA` | `EMAIL_OTP`
     */
    enabledMfas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A collection of user pool Lambda triggers. Amazon Cognito invokes triggers at several possible stages of authentication operations. Triggers can modify the outcome of the operations that invoked them.
     */
    lambdaConfig?: pulumi.Input<inputs.cognito.UserPoolLambdaConfigArgs>;
    /**
     * The multi-factor authentication (MFA) configuration. Valid values include:
     *
     * - `OFF` MFA won't be used for any users.
     * - `ON` MFA is required for all users to sign in.
     * - `OPTIONAL` MFA will be required only for individual users who have an MFA factor activated.
     */
    mfaConfiguration?: pulumi.Input<string>;
    /**
     * A list of user pool policies. Contains the policy that sets password-complexity requirements.
     *
     * This data type is a request and response parameter of [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) and [UpdateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UpdateUserPool.html) , and a response parameter of [DescribeUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html) .
     */
    policies?: pulumi.Input<inputs.cognito.UserPoolPoliciesArgs>;
    /**
     * An array of schema attributes for the new user pool. These attributes can be standard or custom attributes.
     */
    schema?: pulumi.Input<pulumi.Input<inputs.cognito.UserPoolSchemaAttributeArgs>[]>;
    /**
     * The contents of the SMS authentication message.
     */
    smsAuthenticationMessage?: pulumi.Input<string>;
    /**
     * The SMS configuration with the settings that your Amazon Cognito user pool must use to send an SMS message from your AWS account through Amazon Simple Notification Service. To send SMS messages with Amazon SNS in the AWS Region that you want, the Amazon Cognito user pool uses an AWS Identity and Access Management (IAM) role in your AWS account .
     */
    smsConfiguration?: pulumi.Input<inputs.cognito.UserPoolSmsConfigurationArgs>;
    /**
     * This parameter is no longer used. See [VerificationMessageTemplateType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-verificationmessagetemplate.html) .
     */
    smsVerificationMessage?: pulumi.Input<string>;
    /**
     * The settings for updates to user attributes. These settings include the property `AttributesRequireVerificationBeforeUpdate` ,
     * a user-pool setting that tells Amazon Cognito how to handle changes to the value of your users' email address and phone number attributes. For
     * more information, see [Verifying updates to email addresses and phone numbers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html#user-pool-settings-verifications-verify-attribute-updates) .
     */
    userAttributeUpdateSettings?: pulumi.Input<inputs.cognito.UserPoolUserAttributeUpdateSettingsArgs>;
    /**
     * User pool add-ons. Contains settings for activation of advanced security features. To log user security information but take no action, set to `AUDIT` . To configure automatic security responses to risky traffic to your user pool, set to `ENFORCED` .
     *
     * For more information, see [Adding advanced security to a user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-advanced-security.html) .
     */
    userPoolAddOns?: pulumi.Input<inputs.cognito.UserPoolAddOnsArgs>;
    /**
     * A string used to name the user pool.
     */
    userPoolName?: pulumi.Input<string>;
    /**
     * The tag keys and values to assign to the user pool. A tag is a label that you can use to categorize and manage user pools in different ways, such as by purpose, owner, environment, or other criteria.
     */
    userPoolTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether a user can use an email address or phone number as a username when they sign up.
     */
    usernameAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Case sensitivity on the username input for the selected sign-in option. When case sensitivity is set to `False` (case insensitive), users can sign in with any combination of capital and lowercase letters. For example, `username` , `USERNAME` , or `UserName` , or for email, `email@example.com` or `EMaiL@eXamplE.Com` . For most use cases, set case sensitivity to `False` (case insensitive) as a best practice. When usernames and email addresses are case insensitive, Amazon Cognito treats any variation in case as the same user, and prevents a case variation from being assigned to the same attribute for a different user.
     *
     * This configuration is immutable after you set it. For more information, see [UsernameConfigurationType](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UsernameConfigurationType.html) .
     */
    usernameConfiguration?: pulumi.Input<inputs.cognito.UserPoolUsernameConfigurationArgs>;
    /**
     * The template for the verification message that your user pool delivers to users who set an email address or phone number attribute.
     *
     * Set the email message type that corresponds to your `DefaultEmailOption` selection. For `CONFIRM_WITH_LINK` , specify an `EmailMessageByLink` and leave `EmailMessage` blank. For `CONFIRM_WITH_CODE` , specify an `EmailMessage` and leave `EmailMessageByLink` blank. When you supply both parameters with either choice, Amazon Cognito returns an error.
     */
    verificationMessageTemplate?: pulumi.Input<inputs.cognito.UserPoolVerificationMessageTemplateArgs>;
}
