// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Connect::PredefinedAttribute
 */
export class PredefinedAttribute extends pulumi.CustomResource {
    /**
     * Get an existing PredefinedAttribute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PredefinedAttribute {
        return new PredefinedAttribute(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:connect:PredefinedAttribute';

    /**
     * Returns true if the given object is an instance of PredefinedAttribute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PredefinedAttribute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PredefinedAttribute.__pulumiType;
    }

    /**
     * The identifier of the Amazon Connect instance.
     */
    public readonly instanceArn!: pulumi.Output<string>;
    /**
     * Last modified region.
     */
    public /*out*/ readonly lastModifiedRegion!: pulumi.Output<string>;
    /**
     * Last modified time.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<number>;
    /**
     * The name of the predefined attribute.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The values of a predefined attribute.
     */
    public readonly values!: pulumi.Output<outputs.connect.ValuesProperties>;

    /**
     * Create a PredefinedAttribute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PredefinedAttributeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceArn'");
            }
            if ((!args || args.values === undefined) && !opts.urn) {
                throw new Error("Missing required property 'values'");
            }
            resourceInputs["instanceArn"] = args ? args.instanceArn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["values"] = args ? args.values : undefined;
            resourceInputs["lastModifiedRegion"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
        } else {
            resourceInputs["instanceArn"] = undefined /*out*/;
            resourceInputs["lastModifiedRegion"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["values"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["instanceArn", "name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(PredefinedAttribute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PredefinedAttribute resource.
 */
export interface PredefinedAttributeArgs {
    /**
     * The identifier of the Amazon Connect instance.
     */
    instanceArn: pulumi.Input<string>;
    /**
     * The name of the predefined attribute.
     */
    name?: pulumi.Input<string>;
    /**
     * The values of a predefined attribute.
     */
    values: pulumi.Input<inputs.connect.ValuesPropertiesArgs>;
}
