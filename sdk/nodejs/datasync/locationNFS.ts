// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationNFS
 */
export class LocationNFS extends pulumi.CustomResource {
    /**
     * Get an existing LocationNFS resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LocationNFS {
        return new LocationNFS(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:datasync:LocationNFS';

    /**
     * Returns true if the given object is an instance of LocationNFS.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocationNFS {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocationNFS.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the NFS location.
     */
    public /*out*/ readonly locationArn!: pulumi.Output<string>;
    /**
     * The URL of the NFS location that was described.
     */
    public /*out*/ readonly locationUri!: pulumi.Output<string>;
    public readonly mountOptions!: pulumi.Output<outputs.datasync.LocationNFSMountOptions | undefined>;
    public readonly onPremConfig!: pulumi.Output<outputs.datasync.LocationNFSOnPremConfig>;
    /**
     * The name of the NFS server. This value is the IP address or DNS name of the NFS server.
     */
    public readonly serverHostname!: pulumi.Output<string>;
    /**
     * The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.
     */
    public readonly subdirectory!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.datasync.LocationNFSTag[] | undefined>;

    /**
     * Create a LocationNFS resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocationNFSArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.onPremConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'onPremConfig'");
            }
            if ((!args || args.serverHostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverHostname'");
            }
            if ((!args || args.subdirectory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subdirectory'");
            }
            resourceInputs["mountOptions"] = args ? args.mountOptions : undefined;
            resourceInputs["onPremConfig"] = args ? args.onPremConfig : undefined;
            resourceInputs["serverHostname"] = args ? args.serverHostname : undefined;
            resourceInputs["subdirectory"] = args ? args.subdirectory : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["locationArn"] = undefined /*out*/;
            resourceInputs["locationUri"] = undefined /*out*/;
        } else {
            resourceInputs["locationArn"] = undefined /*out*/;
            resourceInputs["locationUri"] = undefined /*out*/;
            resourceInputs["mountOptions"] = undefined /*out*/;
            resourceInputs["onPremConfig"] = undefined /*out*/;
            resourceInputs["serverHostname"] = undefined /*out*/;
            resourceInputs["subdirectory"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LocationNFS.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LocationNFS resource.
 */
export interface LocationNFSArgs {
    mountOptions?: pulumi.Input<inputs.datasync.LocationNFSMountOptionsArgs>;
    onPremConfig: pulumi.Input<inputs.datasync.LocationNFSOnPremConfigArgs>;
    /**
     * The name of the NFS server. This value is the IP address or DNS name of the NFS server.
     */
    serverHostname: pulumi.Input<string>;
    /**
     * The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.
     */
    subdirectory: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.datasync.LocationNFSTagArgs>[]>;
}