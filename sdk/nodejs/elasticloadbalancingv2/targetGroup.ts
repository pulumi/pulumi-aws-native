// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup
 *
 * @deprecated TargetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class TargetGroup extends pulumi.CustomResource {
    /**
     * Get an existing TargetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TargetGroup {
        pulumi.log.warn("TargetGroup is deprecated: TargetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new TargetGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:elasticloadbalancingv2:TargetGroup';

    /**
     * Returns true if the given object is an instance of TargetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetGroup.__pulumiType;
    }

    public readonly healthCheckEnabled!: pulumi.Output<boolean | undefined>;
    public readonly healthCheckIntervalSeconds!: pulumi.Output<number | undefined>;
    public readonly healthCheckPath!: pulumi.Output<string | undefined>;
    public readonly healthCheckPort!: pulumi.Output<string | undefined>;
    public readonly healthCheckProtocol!: pulumi.Output<string | undefined>;
    public readonly healthCheckTimeoutSeconds!: pulumi.Output<number | undefined>;
    public readonly healthyThresholdCount!: pulumi.Output<number | undefined>;
    public readonly ipAddressType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly loadBalancerArns!: pulumi.Output<string[]>;
    public readonly matcher!: pulumi.Output<outputs.elasticloadbalancingv2.TargetGroupMatcher | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly port!: pulumi.Output<number | undefined>;
    public readonly protocol!: pulumi.Output<string | undefined>;
    public readonly protocolVersion!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.elasticloadbalancingv2.TargetGroupTag[] | undefined>;
    public readonly targetGroupAttributes!: pulumi.Output<outputs.elasticloadbalancingv2.TargetGroupAttribute[] | undefined>;
    public /*out*/ readonly targetGroupFullName!: pulumi.Output<string>;
    public /*out*/ readonly targetGroupName!: pulumi.Output<string>;
    public readonly targetType!: pulumi.Output<string | undefined>;
    public readonly targets!: pulumi.Output<outputs.elasticloadbalancingv2.TargetGroupTargetDescription[] | undefined>;
    public readonly unhealthyThresholdCount!: pulumi.Output<number | undefined>;
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a TargetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated TargetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: TargetGroupArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("TargetGroup is deprecated: TargetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["healthCheckEnabled"] = args ? args.healthCheckEnabled : undefined;
            resourceInputs["healthCheckIntervalSeconds"] = args ? args.healthCheckIntervalSeconds : undefined;
            resourceInputs["healthCheckPath"] = args ? args.healthCheckPath : undefined;
            resourceInputs["healthCheckPort"] = args ? args.healthCheckPort : undefined;
            resourceInputs["healthCheckProtocol"] = args ? args.healthCheckProtocol : undefined;
            resourceInputs["healthCheckTimeoutSeconds"] = args ? args.healthCheckTimeoutSeconds : undefined;
            resourceInputs["healthyThresholdCount"] = args ? args.healthyThresholdCount : undefined;
            resourceInputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            resourceInputs["matcher"] = args ? args.matcher : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["protocolVersion"] = args ? args.protocolVersion : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetGroupAttributes"] = args ? args.targetGroupAttributes : undefined;
            resourceInputs["targetType"] = args ? args.targetType : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["unhealthyThresholdCount"] = args ? args.unhealthyThresholdCount : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["loadBalancerArns"] = undefined /*out*/;
            resourceInputs["targetGroupFullName"] = undefined /*out*/;
            resourceInputs["targetGroupName"] = undefined /*out*/;
        } else {
            resourceInputs["healthCheckEnabled"] = undefined /*out*/;
            resourceInputs["healthCheckIntervalSeconds"] = undefined /*out*/;
            resourceInputs["healthCheckPath"] = undefined /*out*/;
            resourceInputs["healthCheckPort"] = undefined /*out*/;
            resourceInputs["healthCheckProtocol"] = undefined /*out*/;
            resourceInputs["healthCheckTimeoutSeconds"] = undefined /*out*/;
            resourceInputs["healthyThresholdCount"] = undefined /*out*/;
            resourceInputs["ipAddressType"] = undefined /*out*/;
            resourceInputs["loadBalancerArns"] = undefined /*out*/;
            resourceInputs["matcher"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
            resourceInputs["protocolVersion"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["targetGroupAttributes"] = undefined /*out*/;
            resourceInputs["targetGroupFullName"] = undefined /*out*/;
            resourceInputs["targetGroupName"] = undefined /*out*/;
            resourceInputs["targetType"] = undefined /*out*/;
            resourceInputs["targets"] = undefined /*out*/;
            resourceInputs["unhealthyThresholdCount"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TargetGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TargetGroup resource.
 */
export interface TargetGroupArgs {
    healthCheckEnabled?: pulumi.Input<boolean>;
    healthCheckIntervalSeconds?: pulumi.Input<number>;
    healthCheckPath?: pulumi.Input<string>;
    healthCheckPort?: pulumi.Input<string>;
    healthCheckProtocol?: pulumi.Input<string>;
    healthCheckTimeoutSeconds?: pulumi.Input<number>;
    healthyThresholdCount?: pulumi.Input<number>;
    ipAddressType?: pulumi.Input<string>;
    matcher?: pulumi.Input<inputs.elasticloadbalancingv2.TargetGroupMatcherArgs>;
    name?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    protocol?: pulumi.Input<string>;
    protocolVersion?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.TargetGroupTagArgs>[]>;
    targetGroupAttributes?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.TargetGroupAttributeArgs>[]>;
    targetType?: pulumi.Input<string>;
    targets?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.TargetGroupTargetDescriptionArgs>[]>;
    unhealthyThresholdCount?: pulumi.Input<number>;
    vpcId?: pulumi.Input<string>;
}