// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::MemoryDB::SubnetGroup resource creates an Amazon MemoryDB Subnet Group.
 */
export class SubnetGroup extends pulumi.CustomResource {
    /**
     * Get an existing SubnetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SubnetGroup {
        return new SubnetGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:memorydb:SubnetGroup';

    /**
     * Returns true if the given object is an instance of SubnetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubnetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubnetGroup.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the subnet group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * An optional description of the subnet group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the subnet group. This value must be unique as it also serves as the subnet group identifier.
     */
    public readonly subnetGroupName!: pulumi.Output<string>;
    /**
     * A list of VPC subnet IDs for the subnet group.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * Supported network types would be a list of network types supported by subnet group and can be either [ipv4] or [ipv4, dual_stack] or [ipv6].
     */
    public /*out*/ readonly supportedNetworkTypes!: pulumi.Output<string[]>;
    /**
     * An array of key-value pairs to apply to this subnet group.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a SubnetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["subnetGroupName"] = args ? args.subnetGroupName : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["supportedNetworkTypes"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["subnetGroupName"] = undefined /*out*/;
            resourceInputs["subnetIds"] = undefined /*out*/;
            resourceInputs["supportedNetworkTypes"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["subnetGroupName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(SubnetGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SubnetGroup resource.
 */
export interface SubnetGroupArgs {
    /**
     * An optional description of the subnet group.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the subnet group. This value must be unique as it also serves as the subnet group identifier.
     */
    subnetGroupName?: pulumi.Input<string>;
    /**
     * A list of VPC subnet IDs for the subnet group.
     */
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of key-value pairs to apply to this subnet group.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
