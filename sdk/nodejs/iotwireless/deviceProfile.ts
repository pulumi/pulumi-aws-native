// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Device Profile's resource schema demonstrating some basic constructs and validation rules.
 */
export class DeviceProfile extends pulumi.CustomResource {
    /**
     * Get an existing DeviceProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DeviceProfile {
        return new DeviceProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iotwireless:DeviceProfile';

    /**
     * Returns true if the given object is an instance of DeviceProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceProfile.__pulumiType;
    }

    /**
     * Service profile Arn. Returned after successful create.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation
     */
    public readonly loRaWAN!: pulumi.Output<outputs.iotwireless.DeviceProfileLoRaWANDeviceProfile | undefined>;
    /**
     * Name of service profile
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * A list of key-value pairs that contain metadata for the device profile.
     */
    public readonly tags!: pulumi.Output<outputs.iotwireless.DeviceProfileTag[] | undefined>;

    /**
     * Create a DeviceProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DeviceProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["loRaWAN"] = args ? args.loRaWAN : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["loRaWAN"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeviceProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DeviceProfile resource.
 */
export interface DeviceProfileArgs {
    /**
     * LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation
     */
    loRaWAN?: pulumi.Input<inputs.iotwireless.DeviceProfileLoRaWANDeviceProfileArgs>;
    /**
     * Name of service profile
     */
    name?: pulumi.Input<string>;
    /**
     * A list of key-value pairs that contain metadata for the device profile.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.iotwireless.DeviceProfileTagArgs>[]>;
}