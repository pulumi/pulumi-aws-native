// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WorkspacesInstances::VolumeAssociation
 */
export class VolumeAssociation extends pulumi.CustomResource {
    /**
     * Get an existing VolumeAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VolumeAssociation {
        return new VolumeAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:workspacesinstances:VolumeAssociation';

    /**
     * Returns true if the given object is an instance of VolumeAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VolumeAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VolumeAssociation.__pulumiType;
    }

    /**
     * The device name for the volume attachment
     */
    public readonly device!: pulumi.Output<string>;
    /**
     * Mode to use when disassociating the volume
     */
    public readonly disassociateMode!: pulumi.Output<enums.workspacesinstances.VolumeAssociationDisassociateMode | undefined>;
    /**
     * ID of the volume to attach to the workspace instance
     */
    public readonly volumeId!: pulumi.Output<string>;
    /**
     * ID of the workspace instance to associate with the volume
     */
    public readonly workspaceInstanceId!: pulumi.Output<string>;

    /**
     * Create a VolumeAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.device === undefined) && !opts.urn) {
                throw new Error("Missing required property 'device'");
            }
            if ((!args || args.volumeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeId'");
            }
            if ((!args || args.workspaceInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceInstanceId'");
            }
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["disassociateMode"] = args ? args.disassociateMode : undefined;
            resourceInputs["volumeId"] = args ? args.volumeId : undefined;
            resourceInputs["workspaceInstanceId"] = args ? args.workspaceInstanceId : undefined;
        } else {
            resourceInputs["device"] = undefined /*out*/;
            resourceInputs["disassociateMode"] = undefined /*out*/;
            resourceInputs["volumeId"] = undefined /*out*/;
            resourceInputs["workspaceInstanceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["device", "volumeId", "workspaceInstanceId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(VolumeAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VolumeAssociation resource.
 */
export interface VolumeAssociationArgs {
    /**
     * The device name for the volume attachment
     */
    device: pulumi.Input<string>;
    /**
     * Mode to use when disassociating the volume
     */
    disassociateMode?: pulumi.Input<enums.workspacesinstances.VolumeAssociationDisassociateMode>;
    /**
     * ID of the volume to attach to the workspace instance
     */
    volumeId: pulumi.Input<string>;
    /**
     * ID of the workspace instance to associate with the volume
     */
    workspaceInstanceId: pulumi.Input<string>;
}
