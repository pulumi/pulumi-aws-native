// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Specifies a listener for an Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.
 */
export class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Listener {
        return new Listener(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:elasticloadbalancingv2:Listener';

    /**
     * Returns true if the given object is an instance of Listener.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Listener {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Listener.__pulumiType;
    }

    /**
     * [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
     */
    public readonly alpnPolicy!: pulumi.Output<string[] | undefined>;
    /**
     * The default SSL server certificate for a secure listener. You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
     *  To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html).
     */
    public readonly certificates!: pulumi.Output<outputs.elasticloadbalancingv2.ListenerCertificate[] | undefined>;
    /**
     * The actions for the default rule. You cannot define a condition for a default rule.
     *  To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html).
     */
    public readonly defaultActions!: pulumi.Output<outputs.elasticloadbalancingv2.ListenerAction[]>;
    /**
     * The Amazon Resource Name (ARN) of the listener.
     */
    public /*out*/ readonly listenerArn!: pulumi.Output<string>;
    /**
     * The listener attributes.
     */
    public readonly listenerAttributes!: pulumi.Output<outputs.elasticloadbalancingv2.ListenerAttribute[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the load balancer.
     */
    public readonly loadBalancerArn!: pulumi.Output<string>;
    /**
     * The mutual authentication configuration information.
     */
    public readonly mutualAuthentication!: pulumi.Output<outputs.elasticloadbalancingv2.ListenerMutualAuthentication | undefined>;
    /**
     * The port on which the load balancer is listening. You can't specify a port for a Gateway Load Balancer.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You can't specify a protocol for a Gateway Load Balancer.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported. For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/describe-ssl-policies.html) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/describe-ssl-policies.html) in the *Network Load Balancers Guide*.
     *  Updating the security policy can result in interruptions if the load balancer is handling a high volume of traffic. To decrease the possibility of an interruption if your load balancer is handling a high volume of traffic, create an additional load balancer or request an LCU reservation.
     */
    public readonly sslPolicy!: pulumi.Output<string | undefined>;

    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.defaultActions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultActions'");
            }
            if ((!args || args.loadBalancerArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerArn'");
            }
            resourceInputs["alpnPolicy"] = args ? args.alpnPolicy : undefined;
            resourceInputs["certificates"] = args ? args.certificates : undefined;
            resourceInputs["defaultActions"] = args ? args.defaultActions : undefined;
            resourceInputs["listenerAttributes"] = args ? args.listenerAttributes : undefined;
            resourceInputs["loadBalancerArn"] = args ? args.loadBalancerArn : undefined;
            resourceInputs["mutualAuthentication"] = args ? args.mutualAuthentication : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            resourceInputs["listenerArn"] = undefined /*out*/;
        } else {
            resourceInputs["alpnPolicy"] = undefined /*out*/;
            resourceInputs["certificates"] = undefined /*out*/;
            resourceInputs["defaultActions"] = undefined /*out*/;
            resourceInputs["listenerArn"] = undefined /*out*/;
            resourceInputs["listenerAttributes"] = undefined /*out*/;
            resourceInputs["loadBalancerArn"] = undefined /*out*/;
            resourceInputs["mutualAuthentication"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
            resourceInputs["sslPolicy"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["loadBalancerArn"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Listener.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    /**
     * [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
     */
    alpnPolicy?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The default SSL server certificate for a secure listener. You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
     *  To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html).
     */
    certificates?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.ListenerCertificateArgs>[]>;
    /**
     * The actions for the default rule. You cannot define a condition for a default rule.
     *  To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html).
     */
    defaultActions: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.ListenerActionArgs>[]>;
    /**
     * The listener attributes.
     */
    listenerAttributes?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.ListenerAttributeArgs>[]>;
    /**
     * The Amazon Resource Name (ARN) of the load balancer.
     */
    loadBalancerArn: pulumi.Input<string>;
    /**
     * The mutual authentication configuration information.
     */
    mutualAuthentication?: pulumi.Input<inputs.elasticloadbalancingv2.ListenerMutualAuthenticationArgs>;
    /**
     * The port on which the load balancer is listening. You can't specify a port for a Gateway Load Balancer.
     */
    port?: pulumi.Input<number>;
    /**
     * The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You can't specify a protocol for a Gateway Load Balancer.
     */
    protocol?: pulumi.Input<string>;
    /**
     * [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported. For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/describe-ssl-policies.html) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/describe-ssl-policies.html) in the *Network Load Balancers Guide*.
     *  Updating the security policy can result in interruptions if the load balancer is handling a high volume of traffic. To decrease the possibility of an interruption if your load balancer is handling a high volume of traffic, create an additional load balancer or request an LCU reservation.
     */
    sslPolicy?: pulumi.Input<string>;
}
