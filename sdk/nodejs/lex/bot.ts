// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Amazon Lex conversational bot performing automated tasks such as ordering a pizza, booking a hotel, and so on.
 */
export class Bot extends pulumi.CustomResource {
    /**
     * Get an existing Bot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Bot {
        return new Bot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lex:Bot';

    /**
     * Returns true if the given object is an instance of Bot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bot.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specifies whether to build the bot locales after bot creation completes.
     */
    public readonly autoBuildBotLocales!: pulumi.Output<boolean | undefined>;
    public readonly botFileS3Location!: pulumi.Output<outputs.lex.BotS3Location | undefined>;
    /**
     * List of bot locales
     */
    public readonly botLocales!: pulumi.Output<outputs.lex.BotLocale[] | undefined>;
    /**
     * A list of tags to add to the bot, which can only be added at bot creation.
     */
    public readonly botTags!: pulumi.Output<outputs.lex.BotTag[] | undefined>;
    /**
     * Data privacy setting of the Bot.
     */
    public readonly dataPrivacy!: pulumi.Output<outputs.lex.DataPrivacyProperties>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * IdleSessionTTLInSeconds of the resource
     */
    public readonly idleSessionTTLInSeconds!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly roleArn!: pulumi.Output<string>;
    public readonly testBotAliasSettings!: pulumi.Output<outputs.lex.BotTestBotAliasSettings | undefined>;
    /**
     * A list of tags to add to the test alias for a bot, , which can only be added at bot/bot alias creation.
     */
    public readonly testBotAliasTags!: pulumi.Output<outputs.lex.BotTag[] | undefined>;

    /**
     * Create a Bot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BotArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dataPrivacy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataPrivacy'");
            }
            if ((!args || args.idleSessionTTLInSeconds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'idleSessionTTLInSeconds'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["autoBuildBotLocales"] = args ? args.autoBuildBotLocales : undefined;
            resourceInputs["botFileS3Location"] = args ? args.botFileS3Location : undefined;
            resourceInputs["botLocales"] = args ? args.botLocales : undefined;
            resourceInputs["botTags"] = args ? args.botTags : undefined;
            resourceInputs["dataPrivacy"] = args ? args.dataPrivacy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["idleSessionTTLInSeconds"] = args ? args.idleSessionTTLInSeconds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["testBotAliasSettings"] = args ? args.testBotAliasSettings : undefined;
            resourceInputs["testBotAliasTags"] = args ? args.testBotAliasTags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["autoBuildBotLocales"] = undefined /*out*/;
            resourceInputs["botFileS3Location"] = undefined /*out*/;
            resourceInputs["botLocales"] = undefined /*out*/;
            resourceInputs["botTags"] = undefined /*out*/;
            resourceInputs["dataPrivacy"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["idleSessionTTLInSeconds"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["testBotAliasSettings"] = undefined /*out*/;
            resourceInputs["testBotAliasTags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Bot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Bot resource.
 */
export interface BotArgs {
    /**
     * Specifies whether to build the bot locales after bot creation completes.
     */
    autoBuildBotLocales?: pulumi.Input<boolean>;
    botFileS3Location?: pulumi.Input<inputs.lex.BotS3LocationArgs>;
    /**
     * List of bot locales
     */
    botLocales?: pulumi.Input<pulumi.Input<inputs.lex.BotLocaleArgs>[]>;
    /**
     * A list of tags to add to the bot, which can only be added at bot creation.
     */
    botTags?: pulumi.Input<pulumi.Input<inputs.lex.BotTagArgs>[]>;
    /**
     * Data privacy setting of the Bot.
     */
    dataPrivacy: pulumi.Input<inputs.lex.DataPrivacyPropertiesArgs>;
    description?: pulumi.Input<string>;
    /**
     * IdleSessionTTLInSeconds of the resource
     */
    idleSessionTTLInSeconds: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    roleArn: pulumi.Input<string>;
    testBotAliasSettings?: pulumi.Input<inputs.lex.BotTestBotAliasSettingsArgs>;
    /**
     * A list of tags to add to the test alias for a bot, , which can only be added at bot/bot alias creation.
     */
    testBotAliasTags?: pulumi.Input<pulumi.Input<inputs.lex.BotTagArgs>[]>;
}