// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::CleanRoomsML::TrainingDataset Resource Type
 */
export class TrainingDataset extends pulumi.CustomResource {
    /**
     * Get an existing TrainingDataset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TrainingDataset {
        return new TrainingDataset(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cleanroomsml:TrainingDataset';

    /**
     * Returns true if the given object is an instance of TrainingDataset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrainingDataset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrainingDataset.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly roleArn!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<enums.cleanroomsml.TrainingDatasetStatus>;
    /**
     * An arbitrary set of tags (key-value pairs) for this cleanrooms-ml training dataset.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    public readonly trainingData!: pulumi.Output<outputs.cleanroomsml.TrainingDatasetDataset[]>;
    public /*out*/ readonly trainingDatasetArn!: pulumi.Output<string>;

    /**
     * Create a TrainingDataset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TrainingDatasetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            if ((!args || args.trainingData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trainingData'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trainingData"] = args ? args.trainingData : undefined;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["trainingDatasetArn"] = undefined /*out*/;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["trainingData"] = undefined /*out*/;
            resourceInputs["trainingDatasetArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["description", "name", "roleArn", "trainingData[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TrainingDataset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TrainingDataset resource.
 */
export interface TrainingDatasetArgs {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    roleArn: pulumi.Input<string>;
    /**
     * An arbitrary set of tags (key-value pairs) for this cleanrooms-ml training dataset.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    trainingData: pulumi.Input<pulumi.Input<inputs.cleanroomsml.TrainingDatasetDatasetArgs>[]>;
}