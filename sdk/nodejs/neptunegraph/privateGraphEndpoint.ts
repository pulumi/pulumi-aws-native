// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::NeptuneGraph::PrivateGraphEndpoint resource creates an Amazon NeptuneGraph PrivateGraphEndpoint.
 */
export class PrivateGraphEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing PrivateGraphEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateGraphEndpoint {
        return new PrivateGraphEndpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:neptunegraph:PrivateGraphEndpoint';

    /**
     * Returns true if the given object is an instance of PrivateGraphEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateGraphEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateGraphEndpoint.__pulumiType;
    }

    /**
     * The auto-generated Graph Id assigned by the service.
     */
    public readonly graphIdentifier!: pulumi.Output<string>;
    /**
     * PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.
     *
     *  For example, if GraphIdentifier is `g-12a3bcdef4` and VpcId is `vpc-0a12bc34567de8f90`, the generated PrivateGraphEndpointIdentifier will be `g-12a3bcdef4_vpc-0a12bc34567de8f90`
     */
    public /*out*/ readonly privateGraphEndpointIdentifier!: pulumi.Output<string>;
    /**
     * The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
     */
    public readonly subnetIds!: pulumi.Output<string[] | undefined>;
    /**
     * VPC endpoint that provides a private connection between the Graph and specified VPC.
     */
    public /*out*/ readonly vpcEndpointId!: pulumi.Output<string>;
    /**
     * The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a PrivateGraphEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateGraphEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.graphIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'graphIdentifier'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["graphIdentifier"] = args ? args.graphIdentifier : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["privateGraphEndpointIdentifier"] = undefined /*out*/;
            resourceInputs["vpcEndpointId"] = undefined /*out*/;
        } else {
            resourceInputs["graphIdentifier"] = undefined /*out*/;
            resourceInputs["privateGraphEndpointIdentifier"] = undefined /*out*/;
            resourceInputs["securityGroupIds"] = undefined /*out*/;
            resourceInputs["subnetIds"] = undefined /*out*/;
            resourceInputs["vpcEndpointId"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["graphIdentifier", "securityGroupIds[*]", "subnetIds[*]", "vpcId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(PrivateGraphEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateGraphEndpoint resource.
 */
export interface PrivateGraphEndpointArgs {
    /**
     * The auto-generated Graph Id assigned by the service.
     */
    graphIdentifier: pulumi.Input<string>;
    /**
     * The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
     */
    vpcId: pulumi.Input<string>;
}