// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Bedrock::AgentAlias Resource Type
 */
export class AgentAlias extends pulumi.CustomResource {
    /**
     * Get an existing AgentAlias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AgentAlias {
        return new AgentAlias(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:bedrock:AgentAlias';

    /**
     * Returns true if the given object is an instance of AgentAlias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentAlias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentAlias.__pulumiType;
    }

    /**
     * Arn representation of the Agent Alias.
     */
    public /*out*/ readonly agentAliasArn!: pulumi.Output<string>;
    /**
     * The list of history events for an alias for an Agent.
     */
    public /*out*/ readonly agentAliasHistoryEvents!: pulumi.Output<outputs.bedrock.AgentAliasHistoryEvent[]>;
    /**
     * Id for an Agent Alias generated at the server side.
     */
    public /*out*/ readonly agentAliasId!: pulumi.Output<string>;
    /**
     * Name for a resource.
     */
    public readonly agentAliasName!: pulumi.Output<string>;
    /**
     * The status of the alias of the agent and whether it is ready for use. The following statuses are possible:
     *
     * - CREATING – The agent alias is being created.
     * - PREPARED – The agent alias is finished being created or updated and is ready to be invoked.
     * - FAILED – The agent alias API operation failed.
     * - UPDATING – The agent alias is being updated.
     * - DELETING – The agent alias is being deleted.
     * - DISSOCIATED - The agent alias has no version associated with it.
     */
    public /*out*/ readonly agentAliasStatus!: pulumi.Output<enums.bedrock.AgentAliasStatus>;
    /**
     * Identifier for a resource.
     */
    public readonly agentId!: pulumi.Output<string>;
    /**
     * Time Stamp.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Description of the Resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Routing configuration for an Agent alias.
     */
    public readonly routingConfiguration!: pulumi.Output<outputs.bedrock.AgentAliasRoutingConfigurationListItem[] | undefined>;
    /**
     * Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
     *
     * - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
     * - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Time Stamp.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a AgentAlias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentAliasArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            resourceInputs["agentAliasName"] = args ? args.agentAliasName : undefined;
            resourceInputs["agentId"] = args ? args.agentId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["routingConfiguration"] = args ? args.routingConfiguration : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["agentAliasArn"] = undefined /*out*/;
            resourceInputs["agentAliasHistoryEvents"] = undefined /*out*/;
            resourceInputs["agentAliasId"] = undefined /*out*/;
            resourceInputs["agentAliasStatus"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["agentAliasArn"] = undefined /*out*/;
            resourceInputs["agentAliasHistoryEvents"] = undefined /*out*/;
            resourceInputs["agentAliasId"] = undefined /*out*/;
            resourceInputs["agentAliasName"] = undefined /*out*/;
            resourceInputs["agentAliasStatus"] = undefined /*out*/;
            resourceInputs["agentId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["routingConfiguration"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["agentId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(AgentAlias.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AgentAlias resource.
 */
export interface AgentAliasArgs {
    /**
     * Name for a resource.
     */
    agentAliasName?: pulumi.Input<string>;
    /**
     * Identifier for a resource.
     */
    agentId: pulumi.Input<string>;
    /**
     * Description of the Resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Routing configuration for an Agent alias.
     */
    routingConfiguration?: pulumi.Input<pulumi.Input<inputs.bedrock.AgentAliasRoutingConfigurationListItemArgs>[]>;
    /**
     * Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
     *
     * - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
     * - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
