// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource type definition for AWS::NetworkFirewall::RuleGroup
 */
export class RuleGroup extends pulumi.CustomResource {
    /**
     * Get an existing RuleGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RuleGroup {
        return new RuleGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:networkfirewall:RuleGroup';

    /**
     * Returns true if the given object is an instance of RuleGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleGroup.__pulumiType;
    }

    public readonly capacity!: pulumi.Output<number>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly ruleGroup!: pulumi.Output<outputs.networkfirewall.RuleGroup | undefined>;
    public /*out*/ readonly ruleGroupArn!: pulumi.Output<string>;
    public /*out*/ readonly ruleGroupId!: pulumi.Output<string>;
    public readonly ruleGroupName!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.networkfirewall.RuleGroupTag[] | undefined>;
    public readonly type!: pulumi.Output<enums.networkfirewall.RuleGroupTypeEnum>;

    /**
     * Create a RuleGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.capacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'capacity'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["capacity"] = args ? args.capacity : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ruleGroup"] = args ? args.ruleGroup : undefined;
            resourceInputs["ruleGroupName"] = args ? args.ruleGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["ruleGroupArn"] = undefined /*out*/;
            resourceInputs["ruleGroupId"] = undefined /*out*/;
        } else {
            resourceInputs["capacity"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["ruleGroup"] = undefined /*out*/;
            resourceInputs["ruleGroupArn"] = undefined /*out*/;
            resourceInputs["ruleGroupId"] = undefined /*out*/;
            resourceInputs["ruleGroupName"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RuleGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RuleGroup resource.
 */
export interface RuleGroupArgs {
    capacity: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    ruleGroup?: pulumi.Input<inputs.networkfirewall.RuleGroupArgs>;
    ruleGroupName?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.networkfirewall.RuleGroupTagArgs>[]>;
    type: pulumi.Input<enums.networkfirewall.RuleGroupTypeEnum>;
}