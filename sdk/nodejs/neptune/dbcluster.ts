// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::Neptune::DBCluster resource creates an Amazon Neptune DB cluster.
 */
export class DBCluster extends pulumi.CustomResource {
    /**
     * Get an existing DBCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DBCluster {
        return new DBCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:neptune:DBCluster';

    /**
     * Returns true if the given object is an instance of DBCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBCluster.__pulumiType;
    }

    /**
     * Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf.
     */
    public readonly associatedRoles!: pulumi.Output<outputs.neptune.DBClusterRole[] | undefined>;
    /**
     * Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in.
     */
    public readonly availabilityZones!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the number of days for which automatic DB snapshots are retained.
     */
    public readonly backupRetentionPeriod!: pulumi.Output<number | undefined>;
    /**
     * The resource id for the DB cluster. For example: `cluster-ABCD1234EFGH5678IJKL90MNOP`. The cluster ID uniquely identifies the cluster and is used in things like IAM authentication policies.
     */
    public /*out*/ readonly clusterResourceId!: pulumi.Output<string>;
    /**
     * The DB cluster identifier. Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string.
     */
    public readonly dBClusterIdentifier!: pulumi.Output<string | undefined>;
    /**
     * Provides the name of the DB cluster parameter group.
     */
    public readonly dBClusterParameterGroupName!: pulumi.Output<string | undefined>;
    /**
     * Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group.
     */
    public readonly dBSubnetGroupName!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether or not the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies a list of log types that are enabled for export to CloudWatch Logs.
     */
    public readonly enableCloudwatchLogsExports!: pulumi.Output<string[] | undefined>;
    /**
     * The connection endpoint for the DB cluster. For example: mystack-mydbcluster-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Indicates the database engine version.
     */
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    /**
     * True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.
     */
    public readonly iamAuthEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * If `StorageEncrypted` is true, the Amazon KMS key identifier for the encrypted DB cluster.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the port that the database engine is listening on.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;
    /**
     * Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.
     */
    public readonly preferredBackupWindow!: pulumi.Output<string | undefined>;
    /**
     * Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
     */
    public readonly preferredMaintenanceWindow!: pulumi.Output<string | undefined>;
    /**
     * The reader endpoint for the DB cluster. For example: mystack-mydbcluster-ro-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com
     */
    public /*out*/ readonly readEndpoint!: pulumi.Output<string>;
    /**
     * Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
     *
     * If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
     *
     * If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
     */
    public readonly restoreToTime!: pulumi.Output<string | undefined>;
    /**
     * Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
     *
     * If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
     *
     * If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
     */
    public readonly restoreType!: pulumi.Output<string | undefined>;
    /**
     * Specifies the identifier for a DB cluster snapshot. Must match the identifier of an existing snapshot.
     *
     * After you restore a DB cluster using a SnapshotIdentifier, you must specify the same SnapshotIdentifier for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed.
     *
     * However, if you don't specify the SnapshotIdentifier, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, the DB cluster is restored from the snapshot specified by the SnapshotIdentifier, and the original DB cluster is deleted.
     */
    public readonly snapshotIdentifier!: pulumi.Output<string | undefined>;
    /**
     * Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
     *
     * If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
     *
     * If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
     */
    public readonly sourceDBClusterIdentifier!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the DB cluster is encrypted.
     *
     * If you specify the `DBClusterIdentifier`, `DBSnapshotIdentifier`, or `SourceDBInstanceIdentifier` property, don't specify this property. The value is inherited from the cluster, snapshot, or source DB instance. If you specify the KmsKeyId property, you must enable encryption.
     *
     * If you specify the KmsKeyId, you must enable encryption by setting StorageEncrypted to true.
     */
    public readonly storageEncrypted!: pulumi.Output<boolean | undefined>;
    /**
     * The tags assigned to this cluster.
     */
    public readonly tags!: pulumi.Output<outputs.neptune.DBClusterTag[] | undefined>;
    /**
     * Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
     *
     * If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
     *
     * If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
     */
    public readonly useLatestRestorableTime!: pulumi.Output<boolean | undefined>;
    /**
     * Provides a list of VPC security groups that the DB cluster belongs to.
     */
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a DBCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DBClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["associatedRoles"] = args ? args.associatedRoles : undefined;
            resourceInputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            resourceInputs["backupRetentionPeriod"] = args ? args.backupRetentionPeriod : undefined;
            resourceInputs["dBClusterIdentifier"] = args ? args.dBClusterIdentifier : undefined;
            resourceInputs["dBClusterParameterGroupName"] = args ? args.dBClusterParameterGroupName : undefined;
            resourceInputs["dBSubnetGroupName"] = args ? args.dBSubnetGroupName : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["enableCloudwatchLogsExports"] = args ? args.enableCloudwatchLogsExports : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["iamAuthEnabled"] = args ? args.iamAuthEnabled : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["preferredBackupWindow"] = args ? args.preferredBackupWindow : undefined;
            resourceInputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            resourceInputs["restoreToTime"] = args ? args.restoreToTime : undefined;
            resourceInputs["restoreType"] = args ? args.restoreType : undefined;
            resourceInputs["snapshotIdentifier"] = args ? args.snapshotIdentifier : undefined;
            resourceInputs["sourceDBClusterIdentifier"] = args ? args.sourceDBClusterIdentifier : undefined;
            resourceInputs["storageEncrypted"] = args ? args.storageEncrypted : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["useLatestRestorableTime"] = args ? args.useLatestRestorableTime : undefined;
            resourceInputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            resourceInputs["clusterResourceId"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["readEndpoint"] = undefined /*out*/;
        } else {
            resourceInputs["associatedRoles"] = undefined /*out*/;
            resourceInputs["availabilityZones"] = undefined /*out*/;
            resourceInputs["backupRetentionPeriod"] = undefined /*out*/;
            resourceInputs["clusterResourceId"] = undefined /*out*/;
            resourceInputs["dBClusterIdentifier"] = undefined /*out*/;
            resourceInputs["dBClusterParameterGroupName"] = undefined /*out*/;
            resourceInputs["dBSubnetGroupName"] = undefined /*out*/;
            resourceInputs["deletionProtection"] = undefined /*out*/;
            resourceInputs["enableCloudwatchLogsExports"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["engineVersion"] = undefined /*out*/;
            resourceInputs["iamAuthEnabled"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["preferredBackupWindow"] = undefined /*out*/;
            resourceInputs["preferredMaintenanceWindow"] = undefined /*out*/;
            resourceInputs["readEndpoint"] = undefined /*out*/;
            resourceInputs["restoreToTime"] = undefined /*out*/;
            resourceInputs["restoreType"] = undefined /*out*/;
            resourceInputs["snapshotIdentifier"] = undefined /*out*/;
            resourceInputs["sourceDBClusterIdentifier"] = undefined /*out*/;
            resourceInputs["storageEncrypted"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["useLatestRestorableTime"] = undefined /*out*/;
            resourceInputs["vpcSecurityGroupIds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DBCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBCluster resource.
 */
export interface DBClusterArgs {
    /**
     * Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf.
     */
    associatedRoles?: pulumi.Input<pulumi.Input<inputs.neptune.DBClusterRoleArgs>[]>;
    /**
     * Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in.
     */
    availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the number of days for which automatic DB snapshots are retained.
     */
    backupRetentionPeriod?: pulumi.Input<number>;
    /**
     * The DB cluster identifier. Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string.
     */
    dBClusterIdentifier?: pulumi.Input<string>;
    /**
     * Provides the name of the DB cluster parameter group.
     */
    dBClusterParameterGroupName?: pulumi.Input<string>;
    /**
     * Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group.
     */
    dBSubnetGroupName?: pulumi.Input<string>;
    /**
     * Indicates whether or not the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Specifies a list of log types that are enabled for export to CloudWatch Logs.
     */
    enableCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates the database engine version.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.
     */
    iamAuthEnabled?: pulumi.Input<boolean>;
    /**
     * If `StorageEncrypted` is true, the Amazon KMS key identifier for the encrypted DB cluster.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.
     */
    preferredBackupWindow?: pulumi.Input<string>;
    /**
     * Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
     */
    preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
     *
     * If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
     *
     * If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
     */
    restoreToTime?: pulumi.Input<string>;
    /**
     * Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
     *
     * If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
     *
     * If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
     */
    restoreType?: pulumi.Input<string>;
    /**
     * Specifies the identifier for a DB cluster snapshot. Must match the identifier of an existing snapshot.
     *
     * After you restore a DB cluster using a SnapshotIdentifier, you must specify the same SnapshotIdentifier for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed.
     *
     * However, if you don't specify the SnapshotIdentifier, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, the DB cluster is restored from the snapshot specified by the SnapshotIdentifier, and the original DB cluster is deleted.
     */
    snapshotIdentifier?: pulumi.Input<string>;
    /**
     * Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
     *
     * If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
     *
     * If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
     */
    sourceDBClusterIdentifier?: pulumi.Input<string>;
    /**
     * Indicates whether the DB cluster is encrypted.
     *
     * If you specify the `DBClusterIdentifier`, `DBSnapshotIdentifier`, or `SourceDBInstanceIdentifier` property, don't specify this property. The value is inherited from the cluster, snapshot, or source DB instance. If you specify the KmsKeyId property, you must enable encryption.
     *
     * If you specify the KmsKeyId, you must enable encryption by setting StorageEncrypted to true.
     */
    storageEncrypted?: pulumi.Input<boolean>;
    /**
     * The tags assigned to this cluster.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.neptune.DBClusterTagArgs>[]>;
    /**
     * Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
     *
     * If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
     *
     * If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
     */
    useLatestRestorableTime?: pulumi.Input<boolean>;
    /**
     * Provides a list of VPC security groups that the DB cluster belongs to.
     */
    vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}