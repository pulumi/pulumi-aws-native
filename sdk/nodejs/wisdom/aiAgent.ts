// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Wisdom::AIAgent Resource Type
 */
export class AiAgent extends pulumi.CustomResource {
    /**
     * Get an existing AiAgent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AiAgent {
        return new AiAgent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wisdom:AiAgent';

    /**
     * Returns true if the given object is an instance of AiAgent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AiAgent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AiAgent.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the AI agent.
     */
    public /*out*/ readonly aiAgentArn!: pulumi.Output<string>;
    /**
     * The identifier of the AI Agent.
     */
    public /*out*/ readonly aiAgentId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the Amazon Q in Connect assistant.
     */
    public /*out*/ readonly assistantArn!: pulumi.Output<string>;
    /**
     * The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the ARN.
     */
    public readonly assistantId!: pulumi.Output<string>;
    /**
     * Configuration for the AI Agent.
     */
    public readonly configuration!: pulumi.Output<outputs.wisdom.AiAgentAiAgentConfiguration0Properties | outputs.wisdom.AiAgentAiAgentConfiguration1Properties | outputs.wisdom.AiAgentAiAgentConfiguration2Properties>;
    /**
     * The description of the AI Agent.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly modifiedTimeSeconds!: pulumi.Output<number>;
    /**
     * The name of the AI Agent.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The tags used to organize, track, or control access for this resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the AI Agent.
     */
    public readonly type!: pulumi.Output<enums.wisdom.AiAgentAiAgentType>;

    /**
     * Create a AiAgent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AiAgentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.assistantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assistantId'");
            }
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["assistantId"] = args ? args.assistantId : undefined;
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["aiAgentArn"] = undefined /*out*/;
            resourceInputs["aiAgentId"] = undefined /*out*/;
            resourceInputs["assistantArn"] = undefined /*out*/;
            resourceInputs["modifiedTimeSeconds"] = undefined /*out*/;
        } else {
            resourceInputs["aiAgentArn"] = undefined /*out*/;
            resourceInputs["aiAgentId"] = undefined /*out*/;
            resourceInputs["assistantArn"] = undefined /*out*/;
            resourceInputs["assistantId"] = undefined /*out*/;
            resourceInputs["configuration"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["modifiedTimeSeconds"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["assistantId", "name", "tags.*", "type"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(AiAgent.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AiAgent resource.
 */
export interface AiAgentArgs {
    /**
     * The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the ARN.
     */
    assistantId: pulumi.Input<string>;
    /**
     * Configuration for the AI Agent.
     */
    configuration: pulumi.Input<inputs.wisdom.AiAgentAiAgentConfiguration0PropertiesArgs | inputs.wisdom.AiAgentAiAgentConfiguration1PropertiesArgs | inputs.wisdom.AiAgentAiAgentConfiguration2PropertiesArgs>;
    /**
     * The description of the AI Agent.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the AI Agent.
     */
    name?: pulumi.Input<string>;
    /**
     * The tags used to organize, track, or control access for this resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the AI Agent.
     */
    type: pulumi.Input<enums.wisdom.AiAgentAiAgentType>;
}
