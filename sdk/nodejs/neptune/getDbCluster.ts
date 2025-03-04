// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::Neptune::DBCluster resource creates an Amazon Neptune DB cluster.
 */
export function getDbCluster(args: GetDbClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetDbClusterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:neptune:getDbCluster", {
        "dbClusterIdentifier": args.dbClusterIdentifier,
    }, opts);
}

export interface GetDbClusterArgs {
    /**
     * The DB cluster identifier. Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string.
     */
    dbClusterIdentifier: string;
}

export interface GetDbClusterResult {
    /**
     * Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf.
     */
    readonly associatedRoles?: outputs.neptune.DbClusterDbClusterRole[];
    /**
     * Specifies the number of days for which automatic DB snapshots are retained.
     */
    readonly backupRetentionPeriod?: number;
    /**
     * The resource id for the DB cluster. For example: `cluster-ABCD1234EFGH5678IJKL90MNOP`. The cluster ID uniquely identifies the cluster and is used in things like IAM authentication policies.
     */
    readonly clusterResourceId?: string;
    /**
     * A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default behaviour is not to copy them.
     */
    readonly copyTagsToSnapshot?: boolean;
    /**
     * Provides the name of the DB cluster parameter group.
     */
    readonly dbClusterParameterGroupName?: string;
    /**
     * The port number on which the DB instances in the DB cluster accept connections. 
     *
     * If not specified, the default port used is `8182`. 
     *
     * Note: `Port` property will soon be deprecated from this resource. Please update existing templates to rename it with new property `DBPort` having same functionalities.
     */
    readonly dbPort?: number;
    /**
     * Indicates whether or not the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled.
     */
    readonly deletionProtection?: boolean;
    /**
     * Specifies a list of log types that are enabled for export to CloudWatch Logs.
     */
    readonly enableCloudwatchLogsExports?: string[];
    /**
     * The connection endpoint for the DB cluster. For example: `mystack-mydbcluster-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`
     */
    readonly endpoint?: string;
    /**
     * Indicates the database engine version.
     */
    readonly engineVersion?: string;
    /**
     * True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.
     */
    readonly iamAuthEnabled?: boolean;
    /**
     * The port number on which the DB cluster accepts connections. For example: `8182`.
     */
    readonly port?: string;
    /**
     * Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.
     */
    readonly preferredBackupWindow?: string;
    /**
     * Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
     */
    readonly preferredMaintenanceWindow?: string;
    /**
     * The reader endpoint for the DB cluster. For example: `mystack-mydbcluster-ro-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`
     */
    readonly readEndpoint?: string;
    /**
     * Contains the scaling configuration used by the Neptune Serverless Instances within this DB cluster.
     */
    readonly serverlessScalingConfiguration?: outputs.neptune.DbClusterServerlessScalingConfiguration;
    /**
     * The tags assigned to this cluster.
     */
    readonly tags?: outputs.Tag[];
    /**
     * Provides a list of VPC security groups that the DB cluster belongs to.
     */
    readonly vpcSecurityGroupIds?: string[];
}
/**
 * The AWS::Neptune::DBCluster resource creates an Amazon Neptune DB cluster.
 */
export function getDbClusterOutput(args: GetDbClusterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDbClusterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:neptune:getDbCluster", {
        "dbClusterIdentifier": args.dbClusterIdentifier,
    }, opts);
}

export interface GetDbClusterOutputArgs {
    /**
     * The DB cluster identifier. Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string.
     */
    dbClusterIdentifier: pulumi.Input<string>;
}
