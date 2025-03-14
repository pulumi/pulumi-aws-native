// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::B2BI::Capability Resource Type
 */
export class Capability extends pulumi.CustomResource {
    /**
     * Get an existing Capability resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Capability {
        return new Capability(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:b2bi:Capability';

    /**
     * Returns true if the given object is an instance of Capability.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Capability {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Capability.__pulumiType;
    }

    /**
     * Returns an Amazon Resource Name (ARN) for a specific AWS resource, such as a capability, partnership, profile, or transformer.
     */
    public /*out*/ readonly capabilityArn!: pulumi.Output<string>;
    /**
     * Returns a system-assigned unique identifier for the capability.
     */
    public /*out*/ readonly capabilityId!: pulumi.Output<string>;
    /**
     * Specifies a structure that contains the details for a capability.
     */
    public readonly configuration!: pulumi.Output<outputs.b2bi.CapabilityConfigurationProperties>;
    /**
     * Returns a timestamp for creation date and time of the capability.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Specifies one or more locations in Amazon S3, each specifying an EDI document that can be used with this capability. Each item contains the name of the bucket and the key, to identify the document's location.
     */
    public readonly instructionsDocuments!: pulumi.Output<outputs.b2bi.CapabilityS3Location[] | undefined>;
    /**
     * Returns a timestamp that identifies the most recent date and time that the capability was modified.
     */
    public /*out*/ readonly modifiedAt!: pulumi.Output<string>;
    /**
     * The display name of the capability.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type. You can attach this metadata to resources (capabilities, partnerships, and so on) for any purpose.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * Returns the type of the capability. Currently, only `edi` is supported.
     */
    public readonly type!: pulumi.Output<enums.b2bi.CapabilityType>;

    /**
     * Create a Capability resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CapabilityArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["instructionsDocuments"] = args ? args.instructionsDocuments : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["capabilityArn"] = undefined /*out*/;
            resourceInputs["capabilityId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["modifiedAt"] = undefined /*out*/;
        } else {
            resourceInputs["capabilityArn"] = undefined /*out*/;
            resourceInputs["capabilityId"] = undefined /*out*/;
            resourceInputs["configuration"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["instructionsDocuments"] = undefined /*out*/;
            resourceInputs["modifiedAt"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["type"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Capability.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Capability resource.
 */
export interface CapabilityArgs {
    /**
     * Specifies a structure that contains the details for a capability.
     */
    configuration: pulumi.Input<inputs.b2bi.CapabilityConfigurationPropertiesArgs>;
    /**
     * Specifies one or more locations in Amazon S3, each specifying an EDI document that can be used with this capability. Each item contains the name of the bucket and the key, to identify the document's location.
     */
    instructionsDocuments?: pulumi.Input<pulumi.Input<inputs.b2bi.CapabilityS3LocationArgs>[]>;
    /**
     * The display name of the capability.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type. You can attach this metadata to resources (capabilities, partnerships, and so on) for any purpose.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * Returns the type of the capability. Currently, only `edi` is supported.
     */
    type: pulumi.Input<enums.b2bi.CapabilityType>;
}
