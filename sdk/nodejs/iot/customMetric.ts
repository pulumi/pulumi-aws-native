// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A custom metric published by your devices to Device Defender.
 */
export class CustomMetric extends pulumi.CustomResource {
    /**
     * Get an existing CustomMetric resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomMetric {
        return new CustomMetric(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iot:CustomMetric';

    /**
     * Returns true if the given object is an instance of CustomMetric.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomMetric {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomMetric.__pulumiType;
    }

    /**
     * Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Number (ARN) of the custom metric.
     */
    public /*out*/ readonly metricArn!: pulumi.Output<string>;
    /**
     * The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.
     */
    public readonly metricName!: pulumi.Output<string | undefined>;
    /**
     * The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
     */
    public readonly metricType!: pulumi.Output<enums.iot.CustomMetricMetricType>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.iot.CustomMetricTag[] | undefined>;

    /**
     * Create a CustomMetric resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomMetricArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.metricType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricType'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["metricName"] = args ? args.metricName : undefined;
            resourceInputs["metricType"] = args ? args.metricType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["metricArn"] = undefined /*out*/;
        } else {
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["metricArn"] = undefined /*out*/;
            resourceInputs["metricName"] = undefined /*out*/;
            resourceInputs["metricType"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomMetric.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomMetric resource.
 */
export interface CustomMetricArgs {
    /**
     * Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.
     */
    metricName?: pulumi.Input<string>;
    /**
     * The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
     */
    metricType: pulumi.Input<enums.iot.CustomMetricMetricType>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.iot.CustomMetricTagArgs>[]>;
}