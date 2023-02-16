// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElasticLoadBalancingV2::LoadBalancer
 *
 * @deprecated LoadBalancer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        pulumi.log.warn("LoadBalancer is deprecated: LoadBalancer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new LoadBalancer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:elasticloadbalancingv2:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    public /*out*/ readonly canonicalHostedZoneID!: pulumi.Output<string>;
    public /*out*/ readonly dNSName!: pulumi.Output<string>;
    public readonly ipAddressType!: pulumi.Output<string | undefined>;
    public readonly loadBalancerAttributes!: pulumi.Output<outputs.elasticloadbalancingv2.LoadBalancerAttribute[] | undefined>;
    public /*out*/ readonly loadBalancerFullName!: pulumi.Output<string>;
    public /*out*/ readonly loadBalancerName!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly scheme!: pulumi.Output<string | undefined>;
    public readonly securityGroups!: pulumi.Output<string[] | undefined>;
    public readonly subnetMappings!: pulumi.Output<outputs.elasticloadbalancingv2.LoadBalancerSubnetMapping[] | undefined>;
    public readonly subnets!: pulumi.Output<string[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.elasticloadbalancingv2.LoadBalancerTag[] | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated LoadBalancer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LoadBalancer is deprecated: LoadBalancer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            resourceInputs["loadBalancerAttributes"] = args ? args.loadBalancerAttributes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scheme"] = args ? args.scheme : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["subnetMappings"] = args ? args.subnetMappings : undefined;
            resourceInputs["subnets"] = args ? args.subnets : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["canonicalHostedZoneID"] = undefined /*out*/;
            resourceInputs["dNSName"] = undefined /*out*/;
            resourceInputs["loadBalancerFullName"] = undefined /*out*/;
            resourceInputs["loadBalancerName"] = undefined /*out*/;
        } else {
            resourceInputs["canonicalHostedZoneID"] = undefined /*out*/;
            resourceInputs["dNSName"] = undefined /*out*/;
            resourceInputs["ipAddressType"] = undefined /*out*/;
            resourceInputs["loadBalancerAttributes"] = undefined /*out*/;
            resourceInputs["loadBalancerFullName"] = undefined /*out*/;
            resourceInputs["loadBalancerName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["scheme"] = undefined /*out*/;
            resourceInputs["securityGroups"] = undefined /*out*/;
            resourceInputs["subnetMappings"] = undefined /*out*/;
            resourceInputs["subnets"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    ipAddressType?: pulumi.Input<string>;
    loadBalancerAttributes?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.LoadBalancerAttributeArgs>[]>;
    name?: pulumi.Input<string>;
    scheme?: pulumi.Input<string>;
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    subnetMappings?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.LoadBalancerSubnetMappingArgs>[]>;
    subnets?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.LoadBalancerTagArgs>[]>;
    type?: pulumi.Input<string>;
}