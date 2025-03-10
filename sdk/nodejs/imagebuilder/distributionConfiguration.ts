// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ImageBuilder::DistributionConfiguration
 */
export class DistributionConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing DistributionConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DistributionConfiguration {
        return new DistributionConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:imagebuilder:DistributionConfiguration';

    /**
     * Returns true if the given object is an instance of DistributionConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DistributionConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DistributionConfiguration.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the distribution configuration.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of the distribution configuration.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The distributions of the distribution configuration.
     */
    public readonly distributions!: pulumi.Output<outputs.imagebuilder.DistributionConfigurationDistribution[]>;
    /**
     * The name of the distribution configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The tags associated with the component.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a DistributionConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DistributionConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.distributions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'distributions'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["distributions"] = args ? args.distributions : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["distributions"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DistributionConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DistributionConfiguration resource.
 */
export interface DistributionConfigurationArgs {
    /**
     * The description of the distribution configuration.
     */
    description?: pulumi.Input<string>;
    /**
     * The distributions of the distribution configuration.
     */
    distributions: pulumi.Input<pulumi.Input<inputs.imagebuilder.DistributionConfigurationDistributionArgs>[]>;
    /**
     * The name of the distribution configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * The tags associated with the component.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
