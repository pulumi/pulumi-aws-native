// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTSiteWise::Asset
 */
export class Asset extends pulumi.CustomResource {
    /**
     * Get an existing Asset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Asset {
        return new Asset(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iotsitewise:Asset';

    /**
     * Returns true if the given object is an instance of Asset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Asset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Asset.__pulumiType;
    }

    /**
     * The ARN of the asset
     */
    public /*out*/ readonly assetArn!: pulumi.Output<string>;
    /**
     * A description for the asset
     */
    public readonly assetDescription!: pulumi.Output<string | undefined>;
    public readonly assetHierarchies!: pulumi.Output<outputs.iotsitewise.AssetHierarchy[] | undefined>;
    /**
     * The ID of the asset
     */
    public /*out*/ readonly assetId!: pulumi.Output<string>;
    /**
     * The ID of the asset model from which to create the asset.
     */
    public readonly assetModelId!: pulumi.Output<string>;
    /**
     * A unique, friendly name for the asset.
     */
    public readonly assetName!: pulumi.Output<string>;
    public readonly assetProperties!: pulumi.Output<outputs.iotsitewise.AssetProperty[] | undefined>;
    /**
     * A list of key-value pairs that contain metadata for the asset.
     */
    public readonly tags!: pulumi.Output<outputs.iotsitewise.AssetTag[] | undefined>;

    /**
     * Create a Asset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.assetModelId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assetModelId'");
            }
            resourceInputs["assetDescription"] = args ? args.assetDescription : undefined;
            resourceInputs["assetHierarchies"] = args ? args.assetHierarchies : undefined;
            resourceInputs["assetModelId"] = args ? args.assetModelId : undefined;
            resourceInputs["assetName"] = args ? args.assetName : undefined;
            resourceInputs["assetProperties"] = args ? args.assetProperties : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["assetArn"] = undefined /*out*/;
            resourceInputs["assetId"] = undefined /*out*/;
        } else {
            resourceInputs["assetArn"] = undefined /*out*/;
            resourceInputs["assetDescription"] = undefined /*out*/;
            resourceInputs["assetHierarchies"] = undefined /*out*/;
            resourceInputs["assetId"] = undefined /*out*/;
            resourceInputs["assetModelId"] = undefined /*out*/;
            resourceInputs["assetName"] = undefined /*out*/;
            resourceInputs["assetProperties"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Asset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Asset resource.
 */
export interface AssetArgs {
    /**
     * A description for the asset
     */
    assetDescription?: pulumi.Input<string>;
    assetHierarchies?: pulumi.Input<pulumi.Input<inputs.iotsitewise.AssetHierarchyArgs>[]>;
    /**
     * The ID of the asset model from which to create the asset.
     */
    assetModelId: pulumi.Input<string>;
    /**
     * A unique, friendly name for the asset.
     */
    assetName?: pulumi.Input<string>;
    assetProperties?: pulumi.Input<pulumi.Input<inputs.iotsitewise.AssetPropertyArgs>[]>;
    /**
     * A list of key-value pairs that contain metadata for the asset.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.iotsitewise.AssetTagArgs>[]>;
}