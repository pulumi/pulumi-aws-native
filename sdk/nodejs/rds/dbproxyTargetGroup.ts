// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::RDS::DBProxyTargetGroup
 */
export class DBProxyTargetGroup extends pulumi.CustomResource {
    /**
     * Get an existing DBProxyTargetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DBProxyTargetGroup {
        return new DBProxyTargetGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:rds:DBProxyTargetGroup';

    /**
     * Returns true if the given object is an instance of DBProxyTargetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBProxyTargetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBProxyTargetGroup.__pulumiType;
    }

    public readonly connectionPoolConfigurationInfo!: pulumi.Output<outputs.rds.DBProxyTargetGroupConnectionPoolConfigurationInfoFormat | undefined>;
    public readonly dBClusterIdentifiers!: pulumi.Output<string[] | undefined>;
    public readonly dBInstanceIdentifiers!: pulumi.Output<string[] | undefined>;
    /**
     * The identifier for the proxy.
     */
    public readonly dBProxyName!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) representing the target group.
     */
    public /*out*/ readonly targetGroupArn!: pulumi.Output<string>;
    /**
     * The identifier for the DBProxyTargetGroup
     */
    public readonly targetGroupName!: pulumi.Output<enums.rds.DBProxyTargetGroupTargetGroupName>;

    /**
     * Create a DBProxyTargetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DBProxyTargetGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dBProxyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dBProxyName'");
            }
            if ((!args || args.targetGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetGroupName'");
            }
            resourceInputs["connectionPoolConfigurationInfo"] = args ? args.connectionPoolConfigurationInfo : undefined;
            resourceInputs["dBClusterIdentifiers"] = args ? args.dBClusterIdentifiers : undefined;
            resourceInputs["dBInstanceIdentifiers"] = args ? args.dBInstanceIdentifiers : undefined;
            resourceInputs["dBProxyName"] = args ? args.dBProxyName : undefined;
            resourceInputs["targetGroupName"] = args ? args.targetGroupName : undefined;
            resourceInputs["targetGroupArn"] = undefined /*out*/;
        } else {
            resourceInputs["connectionPoolConfigurationInfo"] = undefined /*out*/;
            resourceInputs["dBClusterIdentifiers"] = undefined /*out*/;
            resourceInputs["dBInstanceIdentifiers"] = undefined /*out*/;
            resourceInputs["dBProxyName"] = undefined /*out*/;
            resourceInputs["targetGroupArn"] = undefined /*out*/;
            resourceInputs["targetGroupName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DBProxyTargetGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBProxyTargetGroup resource.
 */
export interface DBProxyTargetGroupArgs {
    connectionPoolConfigurationInfo?: pulumi.Input<inputs.rds.DBProxyTargetGroupConnectionPoolConfigurationInfoFormatArgs>;
    dBClusterIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    dBInstanceIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The identifier for the proxy.
     */
    dBProxyName: pulumi.Input<string>;
    /**
     * The identifier for the DBProxyTargetGroup
     */
    targetGroupName: pulumi.Input<enums.rds.DBProxyTargetGroupTargetGroupName>;
}