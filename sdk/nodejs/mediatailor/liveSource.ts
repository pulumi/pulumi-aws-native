// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::MediaTailor::LiveSource Resource Type
 */
export class LiveSource extends pulumi.CustomResource {
    /**
     * Get an existing LiveSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LiveSource {
        return new LiveSource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:mediatailor:LiveSource';

    /**
     * Returns true if the given object is an instance of LiveSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LiveSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LiveSource.__pulumiType;
    }

    /**
     * <p>The ARN of the live source.</p>
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * <p>A list of HTTP package configuration parameters for this live source.</p>
     */
    public readonly httpPackageConfigurations!: pulumi.Output<outputs.mediatailor.LiveSourceHttpPackageConfiguration[]>;
    public readonly liveSourceName!: pulumi.Output<string>;
    public readonly sourceLocationName!: pulumi.Output<string>;
    /**
     * The tags to assign to the live source.
     */
    public readonly tags!: pulumi.Output<outputs.mediatailor.LiveSourceTag[] | undefined>;

    /**
     * Create a LiveSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LiveSourceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.httpPackageConfigurations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'httpPackageConfigurations'");
            }
            if ((!args || args.sourceLocationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceLocationName'");
            }
            resourceInputs["httpPackageConfigurations"] = args ? args.httpPackageConfigurations : undefined;
            resourceInputs["liveSourceName"] = args ? args.liveSourceName : undefined;
            resourceInputs["sourceLocationName"] = args ? args.sourceLocationName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["httpPackageConfigurations"] = undefined /*out*/;
            resourceInputs["liveSourceName"] = undefined /*out*/;
            resourceInputs["sourceLocationName"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["liveSourceName", "sourceLocationName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(LiveSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LiveSource resource.
 */
export interface LiveSourceArgs {
    /**
     * <p>A list of HTTP package configuration parameters for this live source.</p>
     */
    httpPackageConfigurations: pulumi.Input<pulumi.Input<inputs.mediatailor.LiveSourceHttpPackageConfigurationArgs>[]>;
    liveSourceName?: pulumi.Input<string>;
    sourceLocationName: pulumi.Input<string>;
    /**
     * The tags to assign to the live source.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.mediatailor.LiveSourceTagArgs>[]>;
}