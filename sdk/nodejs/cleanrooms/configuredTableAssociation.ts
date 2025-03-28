// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a table that can be queried within a collaboration
 */
export class ConfiguredTableAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ConfiguredTableAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConfiguredTableAssociation {
        return new ConfiguredTableAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cleanrooms:ConfiguredTableAssociation';

    /**
     * Returns true if the given object is an instance of ConfiguredTableAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfiguredTableAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfiguredTableAssociation.__pulumiType;
    }

    /**
     * Returns the Amazon Resource Name (ARN) of the specified configured table association.
     *
     * Example: `arn:aws:cleanrooms:us-east-1:111122223333:configuredtable/a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * An analysis rule for a configured table association. This analysis rule specifies how data from the table can be used within its associated collaboration. In the console, the `ConfiguredTableAssociationAnalysisRule` is referred to as the *collaboration analysis rule* .
     */
    public readonly configuredTableAssociationAnalysisRules!: pulumi.Output<outputs.cleanrooms.ConfiguredTableAssociationAnalysisRule[] | undefined>;
    /**
     * Returns the unique identifier of the specified configured table association.
     *
     * Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`
     */
    public /*out*/ readonly configuredTableAssociationIdentifier!: pulumi.Output<string>;
    /**
     * A unique identifier for the configured table to be associated to. Currently accepts a configured table ID.
     */
    public readonly configuredTableIdentifier!: pulumi.Output<string>;
    /**
     * A description of the configured table association.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The unique ID for the membership this configured table association belongs to.
     */
    public readonly membershipIdentifier!: pulumi.Output<string>;
    /**
     * The name of the configured table association, in lowercase. The table is identified by this name when running protected queries against the underlying data.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The service will assume this role to access catalog metadata and query the table.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a ConfiguredTableAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfiguredTableAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.configuredTableIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuredTableIdentifier'");
            }
            if ((!args || args.membershipIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'membershipIdentifier'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["configuredTableAssociationAnalysisRules"] = args ? args.configuredTableAssociationAnalysisRules : undefined;
            resourceInputs["configuredTableIdentifier"] = args ? args.configuredTableIdentifier : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["membershipIdentifier"] = args ? args.membershipIdentifier : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["configuredTableAssociationIdentifier"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["configuredTableAssociationAnalysisRules"] = undefined /*out*/;
            resourceInputs["configuredTableAssociationIdentifier"] = undefined /*out*/;
            resourceInputs["configuredTableIdentifier"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["membershipIdentifier"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["configuredTableIdentifier", "membershipIdentifier", "name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ConfiguredTableAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConfiguredTableAssociation resource.
 */
export interface ConfiguredTableAssociationArgs {
    /**
     * An analysis rule for a configured table association. This analysis rule specifies how data from the table can be used within its associated collaboration. In the console, the `ConfiguredTableAssociationAnalysisRule` is referred to as the *collaboration analysis rule* .
     */
    configuredTableAssociationAnalysisRules?: pulumi.Input<pulumi.Input<inputs.cleanrooms.ConfiguredTableAssociationAnalysisRuleArgs>[]>;
    /**
     * A unique identifier for the configured table to be associated to. Currently accepts a configured table ID.
     */
    configuredTableIdentifier: pulumi.Input<string>;
    /**
     * A description of the configured table association.
     */
    description?: pulumi.Input<string>;
    /**
     * The unique ID for the membership this configured table association belongs to.
     */
    membershipIdentifier: pulumi.Input<string>;
    /**
     * The name of the configured table association, in lowercase. The table is identified by this name when running protected queries against the underlying data.
     */
    name?: pulumi.Input<string>;
    /**
     * The service will assume this role to access catalog metadata and query the table.
     */
    roleArn: pulumi.Input<string>;
    /**
     * An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
