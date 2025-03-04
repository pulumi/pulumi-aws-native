// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::EC2::TransitGatewayMulticastGroupSource registers and deregisters members and sources (network interfaces) with the transit gateway multicast group
 */
export class TransitGatewayMulticastGroupSource extends pulumi.CustomResource {
    /**
     * Get an existing TransitGatewayMulticastGroupSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TransitGatewayMulticastGroupSource {
        return new TransitGatewayMulticastGroupSource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:TransitGatewayMulticastGroupSource';

    /**
     * Returns true if the given object is an instance of TransitGatewayMulticastGroupSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitGatewayMulticastGroupSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitGatewayMulticastGroupSource.__pulumiType;
    }

    /**
     * The IP address assigned to the transit gateway multicast group.
     */
    public readonly groupIpAddress!: pulumi.Output<string>;
    /**
     * Indicates that the resource is a transit gateway multicast group member.
     */
    public /*out*/ readonly groupMember!: pulumi.Output<boolean>;
    /**
     * Indicates that the resource is a transit gateway multicast group member.
     */
    public /*out*/ readonly groupSource!: pulumi.Output<boolean>;
    /**
     * The ID of the transit gateway attachment.
     */
    public readonly networkInterfaceId!: pulumi.Output<string>;
    /**
     * The ID of the resource.
     */
    public /*out*/ readonly resourceId!: pulumi.Output<string>;
    /**
     * The type of resource, for example a VPC attachment.
     */
    public /*out*/ readonly resourceType!: pulumi.Output<string>;
    /**
     * The source type.
     */
    public /*out*/ readonly sourceType!: pulumi.Output<string>;
    /**
     * The ID of the subnet.
     */
    public /*out*/ readonly subnetId!: pulumi.Output<string>;
    /**
     * The ID of the transit gateway attachment.
     */
    public /*out*/ readonly transitGatewayAttachmentId!: pulumi.Output<string>;
    /**
     * The ID of the transit gateway multicast domain.
     */
    public readonly transitGatewayMulticastDomainId!: pulumi.Output<string>;

    /**
     * Create a TransitGatewayMulticastGroupSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitGatewayMulticastGroupSourceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.groupIpAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupIpAddress'");
            }
            if ((!args || args.networkInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkInterfaceId'");
            }
            if ((!args || args.transitGatewayMulticastDomainId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayMulticastDomainId'");
            }
            resourceInputs["groupIpAddress"] = args ? args.groupIpAddress : undefined;
            resourceInputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            resourceInputs["transitGatewayMulticastDomainId"] = args ? args.transitGatewayMulticastDomainId : undefined;
            resourceInputs["groupMember"] = undefined /*out*/;
            resourceInputs["groupSource"] = undefined /*out*/;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["resourceType"] = undefined /*out*/;
            resourceInputs["sourceType"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
            resourceInputs["transitGatewayAttachmentId"] = undefined /*out*/;
        } else {
            resourceInputs["groupIpAddress"] = undefined /*out*/;
            resourceInputs["groupMember"] = undefined /*out*/;
            resourceInputs["groupSource"] = undefined /*out*/;
            resourceInputs["networkInterfaceId"] = undefined /*out*/;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["resourceType"] = undefined /*out*/;
            resourceInputs["sourceType"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
            resourceInputs["transitGatewayAttachmentId"] = undefined /*out*/;
            resourceInputs["transitGatewayMulticastDomainId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["groupIpAddress", "networkInterfaceId", "transitGatewayMulticastDomainId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TransitGatewayMulticastGroupSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TransitGatewayMulticastGroupSource resource.
 */
export interface TransitGatewayMulticastGroupSourceArgs {
    /**
     * The IP address assigned to the transit gateway multicast group.
     */
    groupIpAddress: pulumi.Input<string>;
    /**
     * The ID of the transit gateway attachment.
     */
    networkInterfaceId: pulumi.Input<string>;
    /**
     * The ID of the transit gateway multicast domain.
     */
    transitGatewayMulticastDomainId: pulumi.Input<string>;
}
