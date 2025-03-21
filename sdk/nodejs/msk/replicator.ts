// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::MSK::Replicator
 */
export class Replicator extends pulumi.CustomResource {
    /**
     * Get an existing Replicator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Replicator {
        return new Replicator(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:msk:Replicator';

    /**
     * Returns true if the given object is an instance of Replicator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Replicator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Replicator.__pulumiType;
    }

    /**
     * The current version of the MSK replicator.
     */
    public /*out*/ readonly currentVersion!: pulumi.Output<string>;
    /**
     * A summary description of the replicator.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies a list of Kafka clusters which are targets of the replicator.
     */
    public readonly kafkaClusters!: pulumi.Output<outputs.msk.ReplicatorKafkaCluster[]>;
    /**
     * A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
     */
    public readonly replicationInfoList!: pulumi.Output<outputs.msk.ReplicatorReplicationInfo[]>;
    /**
     * Amazon Resource Name for the created replicator.
     */
    public /*out*/ readonly replicatorArn!: pulumi.Output<string>;
    /**
     * The name of the replicator.
     */
    public readonly replicatorName!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.
     */
    public readonly serviceExecutionRoleArn!: pulumi.Output<string>;
    /**
     * A collection of tags associated with a resource
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Replicator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicatorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.kafkaClusters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kafkaClusters'");
            }
            if ((!args || args.replicationInfoList === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replicationInfoList'");
            }
            if ((!args || args.serviceExecutionRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceExecutionRoleArn'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["kafkaClusters"] = args ? args.kafkaClusters : undefined;
            resourceInputs["replicationInfoList"] = args ? args.replicationInfoList : undefined;
            resourceInputs["replicatorName"] = args ? args.replicatorName : undefined;
            resourceInputs["serviceExecutionRoleArn"] = args ? args.serviceExecutionRoleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["currentVersion"] = undefined /*out*/;
            resourceInputs["replicatorArn"] = undefined /*out*/;
        } else {
            resourceInputs["currentVersion"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["kafkaClusters"] = undefined /*out*/;
            resourceInputs["replicationInfoList"] = undefined /*out*/;
            resourceInputs["replicatorArn"] = undefined /*out*/;
            resourceInputs["replicatorName"] = undefined /*out*/;
            resourceInputs["serviceExecutionRoleArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["description", "kafkaClusters[*]", "replicationInfoList[*].sourceKafkaClusterArn", "replicationInfoList[*].targetCompressionType", "replicationInfoList[*].targetKafkaClusterArn", "replicationInfoList[*].topicReplication.startingPosition", "replicationInfoList[*].topicReplication.topicNameConfiguration", "replicatorName", "serviceExecutionRoleArn"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Replicator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Replicator resource.
 */
export interface ReplicatorArgs {
    /**
     * A summary description of the replicator.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies a list of Kafka clusters which are targets of the replicator.
     */
    kafkaClusters: pulumi.Input<pulumi.Input<inputs.msk.ReplicatorKafkaClusterArgs>[]>;
    /**
     * A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
     */
    replicationInfoList: pulumi.Input<pulumi.Input<inputs.msk.ReplicatorReplicationInfoArgs>[]>;
    /**
     * The name of the replicator.
     */
    replicatorName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.
     */
    serviceExecutionRoleArn: pulumi.Input<string>;
    /**
     * A collection of tags associated with a resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
