// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * VPC Route Server Endpoint
 */
export class RouteServerEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing RouteServerEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RouteServerEndpoint {
        return new RouteServerEndpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:RouteServerEndpoint';

    /**
     * Returns true if the given object is an instance of RouteServerEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteServerEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteServerEndpoint.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the Route Server Endpoint.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID of the Route Server Endpoint.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Elastic Network Interface IP address owned by the Route Server Endpoint
     */
    public /*out*/ readonly eniAddress!: pulumi.Output<string>;
    /**
     * Elastic Network Interface ID owned by the Route Server Endpoint
     */
    public /*out*/ readonly eniId!: pulumi.Output<string>;
    /**
     * Route Server ID
     */
    public readonly routeServerId!: pulumi.Output<string>;
    /**
     * Subnet ID
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * VPC ID
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a RouteServerEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteServerEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.routeServerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeServerId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["routeServerId"] = args ? args.routeServerId : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["eniAddress"] = undefined /*out*/;
            resourceInputs["eniId"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["eniAddress"] = undefined /*out*/;
            resourceInputs["eniId"] = undefined /*out*/;
            resourceInputs["routeServerId"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["routeServerId", "subnetId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(RouteServerEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RouteServerEndpoint resource.
 */
export interface RouteServerEndpointArgs {
    /**
     * Route Server ID
     */
    routeServerId: pulumi.Input<string>;
    /**
     * Subnet ID
     */
    subnetId: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
