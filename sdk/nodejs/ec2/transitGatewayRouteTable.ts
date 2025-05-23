// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TransitGatewayRouteTable
 */
export class TransitGatewayRouteTable extends pulumi.CustomResource {
    /**
     * Get an existing TransitGatewayRouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TransitGatewayRouteTable {
        return new TransitGatewayRouteTable(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:TransitGatewayRouteTable';

    /**
     * Returns true if the given object is an instance of TransitGatewayRouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitGatewayRouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitGatewayRouteTable.__pulumiType;
    }

    /**
     * Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The ID of the transit gateway.
     */
    public readonly transitGatewayId!: pulumi.Output<string>;
    /**
     * Transit Gateway Route Table primary identifier
     */
    public /*out*/ readonly transitGatewayRouteTableId!: pulumi.Output<string>;

    /**
     * Create a TransitGatewayRouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitGatewayRouteTableArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.transitGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayId'");
            }
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
            resourceInputs["transitGatewayRouteTableId"] = undefined /*out*/;
        } else {
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["transitGatewayId"] = undefined /*out*/;
            resourceInputs["transitGatewayRouteTableId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["transitGatewayId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TransitGatewayRouteTable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TransitGatewayRouteTable resource.
 */
export interface TransitGatewayRouteTableArgs {
    /**
     * Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The ID of the transit gateway.
     */
    transitGatewayId: pulumi.Input<string>;
}
