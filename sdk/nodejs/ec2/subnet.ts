// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::Subnet
 */
export class Subnet extends pulumi.CustomResource {
    /**
     * Get an existing Subnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Subnet {
        return new Subnet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:Subnet';

    /**
     * Returns true if the given object is an instance of Subnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subnet.__pulumiType;
    }

    public readonly assignIpv6AddressOnCreation!: pulumi.Output<boolean | undefined>;
    public readonly availabilityZone!: pulumi.Output<string | undefined>;
    public readonly availabilityZoneId!: pulumi.Output<string | undefined>;
    public readonly cidrBlock!: pulumi.Output<string | undefined>;
    public readonly enableDns64!: pulumi.Output<boolean | undefined>;
    public readonly ipv6CidrBlock!: pulumi.Output<string | undefined>;
    public /*out*/ readonly ipv6CidrBlocks!: pulumi.Output<string[]>;
    public readonly ipv6Native!: pulumi.Output<boolean | undefined>;
    public readonly mapPublicIpOnLaunch!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly networkAclAssociationId!: pulumi.Output<string>;
    public readonly outpostArn!: pulumi.Output<string | undefined>;
    public readonly privateDnsNameOptionsOnLaunch!: pulumi.Output<outputs.ec2.PrivateDnsNameOptionsOnLaunchProperties | undefined>;
    public /*out*/ readonly subnetId!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.ec2.SubnetTag[] | undefined>;
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Subnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["assignIpv6AddressOnCreation"] = args ? args.assignIpv6AddressOnCreation : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["availabilityZoneId"] = args ? args.availabilityZoneId : undefined;
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["enableDns64"] = args ? args.enableDns64 : undefined;
            resourceInputs["ipv6CidrBlock"] = args ? args.ipv6CidrBlock : undefined;
            resourceInputs["ipv6Native"] = args ? args.ipv6Native : undefined;
            resourceInputs["mapPublicIpOnLaunch"] = args ? args.mapPublicIpOnLaunch : undefined;
            resourceInputs["outpostArn"] = args ? args.outpostArn : undefined;
            resourceInputs["privateDnsNameOptionsOnLaunch"] = args ? args.privateDnsNameOptionsOnLaunch : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["ipv6CidrBlocks"] = undefined /*out*/;
            resourceInputs["networkAclAssociationId"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
        } else {
            resourceInputs["assignIpv6AddressOnCreation"] = undefined /*out*/;
            resourceInputs["availabilityZone"] = undefined /*out*/;
            resourceInputs["availabilityZoneId"] = undefined /*out*/;
            resourceInputs["cidrBlock"] = undefined /*out*/;
            resourceInputs["enableDns64"] = undefined /*out*/;
            resourceInputs["ipv6CidrBlock"] = undefined /*out*/;
            resourceInputs["ipv6CidrBlocks"] = undefined /*out*/;
            resourceInputs["ipv6Native"] = undefined /*out*/;
            resourceInputs["mapPublicIpOnLaunch"] = undefined /*out*/;
            resourceInputs["networkAclAssociationId"] = undefined /*out*/;
            resourceInputs["outpostArn"] = undefined /*out*/;
            resourceInputs["privateDnsNameOptionsOnLaunch"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Subnet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Subnet resource.
 */
export interface SubnetArgs {
    assignIpv6AddressOnCreation?: pulumi.Input<boolean>;
    availabilityZone?: pulumi.Input<string>;
    availabilityZoneId?: pulumi.Input<string>;
    cidrBlock?: pulumi.Input<string>;
    enableDns64?: pulumi.Input<boolean>;
    ipv6CidrBlock?: pulumi.Input<string>;
    ipv6Native?: pulumi.Input<boolean>;
    mapPublicIpOnLaunch?: pulumi.Input<boolean>;
    outpostArn?: pulumi.Input<string>;
    privateDnsNameOptionsOnLaunch?: pulumi.Input<inputs.ec2.PrivateDnsNameOptionsOnLaunchPropertiesArgs>;
    tags?: pulumi.Input<pulumi.Input<inputs.ec2.SubnetTagArgs>[]>;
    vpcId: pulumi.Input<string>;
}