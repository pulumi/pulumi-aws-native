// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::DocDBElastic::Cluster Amazon DocumentDB (with MongoDB compatibility) Elastic Scale resource describes a Cluster
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:docdbelastic:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    public readonly adminUserName!: pulumi.Output<string>;
    public readonly adminUserPassword!: pulumi.Output<string | undefined>;
    public readonly authType!: pulumi.Output<string>;
    public /*out*/ readonly clusterArn!: pulumi.Output<string>;
    public /*out*/ readonly clusterEndpoint!: pulumi.Output<string>;
    public readonly clusterName!: pulumi.Output<string>;
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    public readonly preferredMaintenanceWindow!: pulumi.Output<string | undefined>;
    public readonly shardCapacity!: pulumi.Output<number>;
    public readonly shardCount!: pulumi.Output<number>;
    public readonly subnetIds!: pulumi.Output<string[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.docdbelastic.ClusterTag[] | undefined>;
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.adminUserName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminUserName'");
            }
            if ((!args || args.authType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authType'");
            }
            if ((!args || args.shardCapacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shardCapacity'");
            }
            if ((!args || args.shardCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shardCount'");
            }
            resourceInputs["adminUserName"] = args ? args.adminUserName : undefined;
            resourceInputs["adminUserPassword"] = args ? args.adminUserPassword : undefined;
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            resourceInputs["shardCapacity"] = args ? args.shardCapacity : undefined;
            resourceInputs["shardCount"] = args ? args.shardCount : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            resourceInputs["clusterArn"] = undefined /*out*/;
            resourceInputs["clusterEndpoint"] = undefined /*out*/;
        } else {
            resourceInputs["adminUserName"] = undefined /*out*/;
            resourceInputs["adminUserPassword"] = undefined /*out*/;
            resourceInputs["authType"] = undefined /*out*/;
            resourceInputs["clusterArn"] = undefined /*out*/;
            resourceInputs["clusterEndpoint"] = undefined /*out*/;
            resourceInputs["clusterName"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["preferredMaintenanceWindow"] = undefined /*out*/;
            resourceInputs["shardCapacity"] = undefined /*out*/;
            resourceInputs["shardCount"] = undefined /*out*/;
            resourceInputs["subnetIds"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vpcSecurityGroupIds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    adminUserName: pulumi.Input<string>;
    adminUserPassword?: pulumi.Input<string>;
    authType: pulumi.Input<string>;
    clusterName?: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
    preferredMaintenanceWindow?: pulumi.Input<string>;
    shardCapacity: pulumi.Input<number>;
    shardCount: pulumi.Input<number>;
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.docdbelastic.ClusterTagArgs>[]>;
    vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}