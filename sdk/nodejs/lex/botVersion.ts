// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A version is a numbered snapshot of your work that you can publish for use in different parts of your workflow, such as development, beta deployment, and production.
 */
export class BotVersion extends pulumi.CustomResource {
    /**
     * Get an existing BotVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BotVersion {
        return new BotVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lex:BotVersion';

    /**
     * Returns true if the given object is an instance of BotVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BotVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BotVersion.__pulumiType;
    }

    public readonly botId!: pulumi.Output<string>;
    public /*out*/ readonly botVersion!: pulumi.Output<string>;
    public readonly botVersionLocaleSpecification!: pulumi.Output<outputs.lex.BotVersionLocaleSpecification[]>;
    public readonly description!: pulumi.Output<string | undefined>;

    /**
     * Create a BotVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BotVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.botId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'botId'");
            }
            if ((!args || args.botVersionLocaleSpecification === undefined) && !opts.urn) {
                throw new Error("Missing required property 'botVersionLocaleSpecification'");
            }
            resourceInputs["botId"] = args ? args.botId : undefined;
            resourceInputs["botVersionLocaleSpecification"] = args ? args.botVersionLocaleSpecification : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["botVersion"] = undefined /*out*/;
        } else {
            resourceInputs["botId"] = undefined /*out*/;
            resourceInputs["botVersion"] = undefined /*out*/;
            resourceInputs["botVersionLocaleSpecification"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BotVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BotVersion resource.
 */
export interface BotVersionArgs {
    botId: pulumi.Input<string>;
    botVersionLocaleSpecification: pulumi.Input<pulumi.Input<inputs.lex.BotVersionLocaleSpecificationArgs>[]>;
    description?: pulumi.Input<string>;
}