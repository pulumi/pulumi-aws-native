// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::Host
 */
export class Host extends pulumi.CustomResource {
    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Host {
        return new Host(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:Host';

    /**
     * Returns true if the given object is an instance of Host.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Host {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Host.__pulumiType;
    }

    /**
     * The ID of the Outpost hardware asset.
     */
    public readonly assetId!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
     */
    public readonly autoPlacement!: pulumi.Output<string | undefined>;
    /**
     * The Availability Zone in which to allocate the Dedicated Host.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * ID of the host created.
     */
    public /*out*/ readonly hostId!: pulumi.Output<string>;
    /**
     * Automatically allocates a new dedicated host and moves your instances on to it if a degradation is detected on your current host.
     */
    public readonly hostMaintenance!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
     */
    public readonly hostRecovery!: pulumi.Output<string | undefined>;
    /**
     * Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.
     */
    public readonly instanceFamily!: pulumi.Output<string | undefined>;
    /**
     * Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.
     */
    public readonly instanceType!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.
     */
    public readonly outpostArn!: pulumi.Output<string | undefined>;
    /**
     * Any tags assigned to the Host.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Host resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            resourceInputs["assetId"] = args ? args.assetId : undefined;
            resourceInputs["autoPlacement"] = args ? args.autoPlacement : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["hostMaintenance"] = args ? args.hostMaintenance : undefined;
            resourceInputs["hostRecovery"] = args ? args.hostRecovery : undefined;
            resourceInputs["instanceFamily"] = args ? args.instanceFamily : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["outpostArn"] = args ? args.outpostArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["hostId"] = undefined /*out*/;
        } else {
            resourceInputs["assetId"] = undefined /*out*/;
            resourceInputs["autoPlacement"] = undefined /*out*/;
            resourceInputs["availabilityZone"] = undefined /*out*/;
            resourceInputs["hostId"] = undefined /*out*/;
            resourceInputs["hostMaintenance"] = undefined /*out*/;
            resourceInputs["hostRecovery"] = undefined /*out*/;
            resourceInputs["instanceFamily"] = undefined /*out*/;
            resourceInputs["instanceType"] = undefined /*out*/;
            resourceInputs["outpostArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["assetId", "availabilityZone", "instanceFamily", "instanceType", "outpostArn"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Host.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Host resource.
 */
export interface HostArgs {
    /**
     * The ID of the Outpost hardware asset.
     */
    assetId?: pulumi.Input<string>;
    /**
     * Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
     */
    autoPlacement?: pulumi.Input<string>;
    /**
     * The Availability Zone in which to allocate the Dedicated Host.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * Automatically allocates a new dedicated host and moves your instances on to it if a degradation is detected on your current host.
     */
    hostMaintenance?: pulumi.Input<string>;
    /**
     * Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
     */
    hostRecovery?: pulumi.Input<string>;
    /**
     * Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.
     */
    instanceFamily?: pulumi.Input<string>;
    /**
     * Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.
     */
    outpostArn?: pulumi.Input<string>;
    /**
     * Any tags assigned to the Host.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
