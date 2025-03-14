// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.
 */
export class MicrosoftTeamsChannelConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing MicrosoftTeamsChannelConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MicrosoftTeamsChannelConfiguration {
        return new MicrosoftTeamsChannelConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:chatbot:MicrosoftTeamsChannelConfiguration';

    /**
     * Returns true if the given object is an instance of MicrosoftTeamsChannelConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MicrosoftTeamsChannelConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MicrosoftTeamsChannelConfiguration.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the configuration
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the configuration
     */
    public readonly configurationName!: pulumi.Output<string>;
    /**
     * ARNs of Custom Actions to associate with notifications in the provided chat channel.
     */
    public readonly customizationResourceArns!: pulumi.Output<string[] | undefined>;
    /**
     * The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
     */
    public readonly guardrailPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * The ARN of the IAM role that defines the permissions for AWS Chatbot
     */
    public readonly iamRoleArn!: pulumi.Output<string>;
    /**
     * Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs
     */
    public readonly loggingLevel!: pulumi.Output<string | undefined>;
    /**
     * ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.
     */
    public readonly snsTopicArns!: pulumi.Output<string[] | undefined>;
    /**
     * The tags to add to the configuration
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The id of the Microsoft Teams team
     */
    public readonly teamId!: pulumi.Output<string>;
    /**
     * The id of the Microsoft Teams channel
     */
    public readonly teamsChannelId!: pulumi.Output<string>;
    /**
     * The name of the Microsoft Teams channel
     */
    public readonly teamsChannelName!: pulumi.Output<string | undefined>;
    /**
     * The id of the Microsoft Teams tenant
     */
    public readonly teamsTenantId!: pulumi.Output<string>;
    /**
     * Enables use of a user role requirement in your chat configuration
     */
    public readonly userRoleRequired!: pulumi.Output<boolean | undefined>;

    /**
     * Create a MicrosoftTeamsChannelConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MicrosoftTeamsChannelConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.iamRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'iamRoleArn'");
            }
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            if ((!args || args.teamsChannelId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamsChannelId'");
            }
            if ((!args || args.teamsTenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamsTenantId'");
            }
            resourceInputs["configurationName"] = args ? args.configurationName : undefined;
            resourceInputs["customizationResourceArns"] = args ? args.customizationResourceArns : undefined;
            resourceInputs["guardrailPolicies"] = args ? args.guardrailPolicies : undefined;
            resourceInputs["iamRoleArn"] = args ? args.iamRoleArn : undefined;
            resourceInputs["loggingLevel"] = args ? args.loggingLevel : undefined;
            resourceInputs["snsTopicArns"] = args ? args.snsTopicArns : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["teamsChannelId"] = args ? args.teamsChannelId : undefined;
            resourceInputs["teamsChannelName"] = args ? args.teamsChannelName : undefined;
            resourceInputs["teamsTenantId"] = args ? args.teamsTenantId : undefined;
            resourceInputs["userRoleRequired"] = args ? args.userRoleRequired : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["configurationName"] = undefined /*out*/;
            resourceInputs["customizationResourceArns"] = undefined /*out*/;
            resourceInputs["guardrailPolicies"] = undefined /*out*/;
            resourceInputs["iamRoleArn"] = undefined /*out*/;
            resourceInputs["loggingLevel"] = undefined /*out*/;
            resourceInputs["snsTopicArns"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["teamId"] = undefined /*out*/;
            resourceInputs["teamsChannelId"] = undefined /*out*/;
            resourceInputs["teamsChannelName"] = undefined /*out*/;
            resourceInputs["teamsTenantId"] = undefined /*out*/;
            resourceInputs["userRoleRequired"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["configurationName", "teamId", "teamsTenantId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(MicrosoftTeamsChannelConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MicrosoftTeamsChannelConfiguration resource.
 */
export interface MicrosoftTeamsChannelConfigurationArgs {
    /**
     * The name of the configuration
     */
    configurationName?: pulumi.Input<string>;
    /**
     * ARNs of Custom Actions to associate with notifications in the provided chat channel.
     */
    customizationResourceArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
     */
    guardrailPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ARN of the IAM role that defines the permissions for AWS Chatbot
     */
    iamRoleArn: pulumi.Input<string>;
    /**
     * Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs
     */
    loggingLevel?: pulumi.Input<string>;
    /**
     * ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.
     */
    snsTopicArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tags to add to the configuration
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The id of the Microsoft Teams team
     */
    teamId: pulumi.Input<string>;
    /**
     * The id of the Microsoft Teams channel
     */
    teamsChannelId: pulumi.Input<string>;
    /**
     * The name of the Microsoft Teams channel
     */
    teamsChannelName?: pulumi.Input<string>;
    /**
     * The id of the Microsoft Teams tenant
     */
    teamsTenantId: pulumi.Input<string>;
    /**
     * Enables use of a user role requirement in your chat configuration
     */
    userRoleRequired?: pulumi.Input<boolean>;
}
