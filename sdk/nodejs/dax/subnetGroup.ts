// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DAX::SubnetGroup
 *
 * @deprecated SubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
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
        pulumi.log.warn("SubnetGroup is deprecated: SubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new SubnetGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:dax:SubnetGroup';

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

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly subnetGroupName!: pulumi.Output<string | undefined>;
    public readonly subnetIds!: pulumi.Output<string[]>;

    /**
     * Create a SubnetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated SubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: SubnetGroupArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("SubnetGroup is deprecated: SubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["subnetGroupName"] = args ? args.subnetGroupName : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["subnetGroupName"] = undefined /*out*/;
            resourceInputs["subnetIds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SubnetGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SubnetGroup resource.
 */
export interface SubnetGroupArgs {
    description?: pulumi.Input<string>;
    subnetGroupName?: pulumi.Input<string>;
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
}