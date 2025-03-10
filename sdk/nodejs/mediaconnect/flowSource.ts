// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaConnect::FlowSource
 */
export class FlowSource extends pulumi.CustomResource {
    /**
     * Get an existing FlowSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FlowSource {
        return new FlowSource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:mediaconnect:FlowSource';

    /**
     * Returns true if the given object is an instance of FlowSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlowSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlowSource.__pulumiType;
    }

    /**
     * The type of encryption that is used on the content ingested from this source.
     */
    public readonly decryption!: pulumi.Output<outputs.mediaconnect.FlowSourceEncryption | undefined>;
    /**
     * A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
     */
    public readonly entitlementArn!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the flow.
     */
    public readonly flowArn!: pulumi.Output<string | undefined>;
    /**
     * The source configuration for cloud flows receiving a stream from a bridge.
     */
    public readonly gatewayBridgeSource!: pulumi.Output<outputs.mediaconnect.FlowSourceGatewayBridgeSource | undefined>;
    /**
     * The IP address that the flow will be listening on for incoming content.
     */
    public /*out*/ readonly ingestIp!: pulumi.Output<string>;
    /**
     * The port that the flow will be listening on for incoming content.
     */
    public readonly ingestPort!: pulumi.Output<number | undefined>;
    /**
     * The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
     */
    public readonly maxBitrate!: pulumi.Output<number | undefined>;
    /**
     * The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
     */
    public readonly maxLatency!: pulumi.Output<number | undefined>;
    /**
     * The minimum latency in milliseconds.
     */
    public readonly minLatency!: pulumi.Output<number | undefined>;
    /**
     * The name of the source.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The protocol that is used by the source.
     */
    public readonly protocol!: pulumi.Output<enums.mediaconnect.FlowSourceProtocol | undefined>;
    /**
     * The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
     */
    public readonly senderControlPort!: pulumi.Output<number | undefined>;
    /**
     * The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
     */
    public readonly senderIpAddress!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the source.
     */
    public /*out*/ readonly sourceArn!: pulumi.Output<string>;
    /**
     * The port that the flow will be listening on for incoming content.(ReadOnly)
     */
    public /*out*/ readonly sourceIngestPort!: pulumi.Output<string>;
    /**
     * Source IP or domain name for SRT-caller protocol.
     */
    public readonly sourceListenerAddress!: pulumi.Output<string | undefined>;
    /**
     * Source port for SRT-caller protocol.
     */
    public readonly sourceListenerPort!: pulumi.Output<number | undefined>;
    /**
     * The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
     */
    public readonly streamId!: pulumi.Output<string | undefined>;
    /**
     * The name of the VPC Interface this Source is configured with.
     */
    public readonly vpcInterfaceName!: pulumi.Output<string | undefined>;
    /**
     * The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
     */
    public readonly whitelistCidr!: pulumi.Output<string | undefined>;

    /**
     * Create a FlowSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlowSourceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            resourceInputs["decryption"] = args ? args.decryption : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["entitlementArn"] = args ? args.entitlementArn : undefined;
            resourceInputs["flowArn"] = args ? args.flowArn : undefined;
            resourceInputs["gatewayBridgeSource"] = args ? args.gatewayBridgeSource : undefined;
            resourceInputs["ingestPort"] = args ? args.ingestPort : undefined;
            resourceInputs["maxBitrate"] = args ? args.maxBitrate : undefined;
            resourceInputs["maxLatency"] = args ? args.maxLatency : undefined;
            resourceInputs["minLatency"] = args ? args.minLatency : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["senderControlPort"] = args ? args.senderControlPort : undefined;
            resourceInputs["senderIpAddress"] = args ? args.senderIpAddress : undefined;
            resourceInputs["sourceListenerAddress"] = args ? args.sourceListenerAddress : undefined;
            resourceInputs["sourceListenerPort"] = args ? args.sourceListenerPort : undefined;
            resourceInputs["streamId"] = args ? args.streamId : undefined;
            resourceInputs["vpcInterfaceName"] = args ? args.vpcInterfaceName : undefined;
            resourceInputs["whitelistCidr"] = args ? args.whitelistCidr : undefined;
            resourceInputs["ingestIp"] = undefined /*out*/;
            resourceInputs["sourceArn"] = undefined /*out*/;
            resourceInputs["sourceIngestPort"] = undefined /*out*/;
        } else {
            resourceInputs["decryption"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["entitlementArn"] = undefined /*out*/;
            resourceInputs["flowArn"] = undefined /*out*/;
            resourceInputs["gatewayBridgeSource"] = undefined /*out*/;
            resourceInputs["ingestIp"] = undefined /*out*/;
            resourceInputs["ingestPort"] = undefined /*out*/;
            resourceInputs["maxBitrate"] = undefined /*out*/;
            resourceInputs["maxLatency"] = undefined /*out*/;
            resourceInputs["minLatency"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
            resourceInputs["senderControlPort"] = undefined /*out*/;
            resourceInputs["senderIpAddress"] = undefined /*out*/;
            resourceInputs["sourceArn"] = undefined /*out*/;
            resourceInputs["sourceIngestPort"] = undefined /*out*/;
            resourceInputs["sourceListenerAddress"] = undefined /*out*/;
            resourceInputs["sourceListenerPort"] = undefined /*out*/;
            resourceInputs["streamId"] = undefined /*out*/;
            resourceInputs["vpcInterfaceName"] = undefined /*out*/;
            resourceInputs["whitelistCidr"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(FlowSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FlowSource resource.
 */
export interface FlowSourceArgs {
    /**
     * The type of encryption that is used on the content ingested from this source.
     */
    decryption?: pulumi.Input<inputs.mediaconnect.FlowSourceEncryptionArgs>;
    /**
     * A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.
     */
    description: pulumi.Input<string>;
    /**
     * The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
     */
    entitlementArn?: pulumi.Input<string>;
    /**
     * The ARN of the flow.
     */
    flowArn?: pulumi.Input<string>;
    /**
     * The source configuration for cloud flows receiving a stream from a bridge.
     */
    gatewayBridgeSource?: pulumi.Input<inputs.mediaconnect.FlowSourceGatewayBridgeSourceArgs>;
    /**
     * The port that the flow will be listening on for incoming content.
     */
    ingestPort?: pulumi.Input<number>;
    /**
     * The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
     */
    maxBitrate?: pulumi.Input<number>;
    /**
     * The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
     */
    maxLatency?: pulumi.Input<number>;
    /**
     * The minimum latency in milliseconds.
     */
    minLatency?: pulumi.Input<number>;
    /**
     * The name of the source.
     */
    name?: pulumi.Input<string>;
    /**
     * The protocol that is used by the source.
     */
    protocol?: pulumi.Input<enums.mediaconnect.FlowSourceProtocol>;
    /**
     * The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
     */
    senderControlPort?: pulumi.Input<number>;
    /**
     * The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
     */
    senderIpAddress?: pulumi.Input<string>;
    /**
     * Source IP or domain name for SRT-caller protocol.
     */
    sourceListenerAddress?: pulumi.Input<string>;
    /**
     * Source port for SRT-caller protocol.
     */
    sourceListenerPort?: pulumi.Input<number>;
    /**
     * The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
     */
    streamId?: pulumi.Input<string>;
    /**
     * The name of the VPC Interface this Source is configured with.
     */
    vpcInterfaceName?: pulumi.Input<string>;
    /**
     * The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
     */
    whitelistCidr?: pulumi.Input<string>;
}
