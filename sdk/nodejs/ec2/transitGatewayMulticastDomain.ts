// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::EC2::TransitGatewayMulticastDomain type
 */
export class TransitGatewayMulticastDomain extends pulumi.CustomResource {
    /**
     * Get an existing TransitGatewayMulticastDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TransitGatewayMulticastDomain {
        return new TransitGatewayMulticastDomain(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:TransitGatewayMulticastDomain';

    /**
     * Returns true if the given object is an instance of TransitGatewayMulticastDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitGatewayMulticastDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitGatewayMulticastDomain.__pulumiType;
    }

    /**
     * The time the transit gateway multicast domain was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The options for the transit gateway multicast domain.
     */
    public readonly options!: pulumi.Output<outputs.ec2.OptionsProperties | undefined>;
    /**
     * The state of the transit gateway multicast domain.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The tags for the transit gateway multicast domain.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The ID of the transit gateway.
     */
    public readonly transitGatewayId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the transit gateway multicast domain.
     */
    public /*out*/ readonly transitGatewayMulticastDomainArn!: pulumi.Output<string>;
    /**
     * The ID of the transit gateway multicast domain.
     */
    public /*out*/ readonly transitGatewayMulticastDomainId!: pulumi.Output<string>;

    /**
     * Create a TransitGatewayMulticastDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitGatewayMulticastDomainArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.transitGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayId'");
            }
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["transitGatewayMulticastDomainArn"] = undefined /*out*/;
            resourceInputs["transitGatewayMulticastDomainId"] = undefined /*out*/;
        } else {
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["options"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["transitGatewayId"] = undefined /*out*/;
            resourceInputs["transitGatewayMulticastDomainArn"] = undefined /*out*/;
            resourceInputs["transitGatewayMulticastDomainId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["transitGatewayId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TransitGatewayMulticastDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TransitGatewayMulticastDomain resource.
 */
export interface TransitGatewayMulticastDomainArgs {
    /**
     * The options for the transit gateway multicast domain.
     */
    options?: pulumi.Input<inputs.ec2.OptionsPropertiesArgs>;
    /**
     * The tags for the transit gateway multicast domain.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The ID of the transit gateway.
     */
    transitGatewayId: pulumi.Input<string>;
}
