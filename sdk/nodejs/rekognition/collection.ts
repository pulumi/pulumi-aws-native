// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::Rekognition::Collection type creates an Amazon Rekognition Collection. A collection is a logical grouping of information about detected faces which can later be referenced for searches on the group
 */
export class Collection extends pulumi.CustomResource {
    /**
     * Get an existing Collection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Collection {
        return new Collection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:rekognition:Collection';

    /**
     * Returns true if the given object is an instance of Collection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Collection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Collection.__pulumiType;
    }

    /**
     * Returns the Amazon Resource Name of the collection.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ID for the collection that you are creating.
     */
    public readonly collectionId!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Collection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CollectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.collectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'collectionId'");
            }
            resourceInputs["collectionId"] = args ? args.collectionId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["collectionId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["collectionId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Collection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Collection resource.
 */
export interface CollectionArgs {
    /**
     * ID for the collection that you are creating.
     */
    collectionId: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
