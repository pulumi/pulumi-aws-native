# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from .. import outputs as _root_outputs

__all__ = [
    'GetDbClusterResult',
    'AwaitableGetDbClusterResult',
    'get_db_cluster',
    'get_db_cluster_output',
]

@pulumi.output_type
class GetDbClusterResult:
    def __init__(__self__, associated_roles=None, backup_retention_period=None, cluster_resource_id=None, copy_tags_to_snapshot=None, db_cluster_parameter_group_name=None, db_port=None, deletion_protection=None, enable_cloudwatch_logs_exports=None, endpoint=None, engine_version=None, iam_auth_enabled=None, port=None, preferred_backup_window=None, preferred_maintenance_window=None, read_endpoint=None, serverless_scaling_configuration=None, tags=None, vpc_security_group_ids=None):
        if associated_roles and not isinstance(associated_roles, list):
            raise TypeError("Expected argument 'associated_roles' to be a list")
        pulumi.set(__self__, "associated_roles", associated_roles)
        if backup_retention_period and not isinstance(backup_retention_period, int):
            raise TypeError("Expected argument 'backup_retention_period' to be a int")
        pulumi.set(__self__, "backup_retention_period", backup_retention_period)
        if cluster_resource_id and not isinstance(cluster_resource_id, str):
            raise TypeError("Expected argument 'cluster_resource_id' to be a str")
        pulumi.set(__self__, "cluster_resource_id", cluster_resource_id)
        if copy_tags_to_snapshot and not isinstance(copy_tags_to_snapshot, bool):
            raise TypeError("Expected argument 'copy_tags_to_snapshot' to be a bool")
        pulumi.set(__self__, "copy_tags_to_snapshot", copy_tags_to_snapshot)
        if db_cluster_parameter_group_name and not isinstance(db_cluster_parameter_group_name, str):
            raise TypeError("Expected argument 'db_cluster_parameter_group_name' to be a str")
        pulumi.set(__self__, "db_cluster_parameter_group_name", db_cluster_parameter_group_name)
        if db_port and not isinstance(db_port, int):
            raise TypeError("Expected argument 'db_port' to be a int")
        pulumi.set(__self__, "db_port", db_port)
        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError("Expected argument 'deletion_protection' to be a bool")
        pulumi.set(__self__, "deletion_protection", deletion_protection)
        if enable_cloudwatch_logs_exports and not isinstance(enable_cloudwatch_logs_exports, list):
            raise TypeError("Expected argument 'enable_cloudwatch_logs_exports' to be a list")
        pulumi.set(__self__, "enable_cloudwatch_logs_exports", enable_cloudwatch_logs_exports)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        pulumi.set(__self__, "engine_version", engine_version)
        if iam_auth_enabled and not isinstance(iam_auth_enabled, bool):
            raise TypeError("Expected argument 'iam_auth_enabled' to be a bool")
        pulumi.set(__self__, "iam_auth_enabled", iam_auth_enabled)
        if port and not isinstance(port, str):
            raise TypeError("Expected argument 'port' to be a str")
        pulumi.set(__self__, "port", port)
        if preferred_backup_window and not isinstance(preferred_backup_window, str):
            raise TypeError("Expected argument 'preferred_backup_window' to be a str")
        pulumi.set(__self__, "preferred_backup_window", preferred_backup_window)
        if preferred_maintenance_window and not isinstance(preferred_maintenance_window, str):
            raise TypeError("Expected argument 'preferred_maintenance_window' to be a str")
        pulumi.set(__self__, "preferred_maintenance_window", preferred_maintenance_window)
        if read_endpoint and not isinstance(read_endpoint, str):
            raise TypeError("Expected argument 'read_endpoint' to be a str")
        pulumi.set(__self__, "read_endpoint", read_endpoint)
        if serverless_scaling_configuration and not isinstance(serverless_scaling_configuration, dict):
            raise TypeError("Expected argument 'serverless_scaling_configuration' to be a dict")
        pulumi.set(__self__, "serverless_scaling_configuration", serverless_scaling_configuration)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if vpc_security_group_ids and not isinstance(vpc_security_group_ids, list):
            raise TypeError("Expected argument 'vpc_security_group_ids' to be a list")
        pulumi.set(__self__, "vpc_security_group_ids", vpc_security_group_ids)

    @property
    @pulumi.getter(name="associatedRoles")
    def associated_roles(self) -> Optional[Sequence['outputs.DbClusterDbClusterRole']]:
        """
        Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf.
        """
        return pulumi.get(self, "associated_roles")

    @property
    @pulumi.getter(name="backupRetentionPeriod")
    def backup_retention_period(self) -> Optional[builtins.int]:
        """
        Specifies the number of days for which automatic DB snapshots are retained.
        """
        return pulumi.get(self, "backup_retention_period")

    @property
    @pulumi.getter(name="clusterResourceId")
    def cluster_resource_id(self) -> Optional[builtins.str]:
        """
        The resource id for the DB cluster. For example: `cluster-ABCD1234EFGH5678IJKL90MNOP`. The cluster ID uniquely identifies the cluster and is used in things like IAM authentication policies.
        """
        return pulumi.get(self, "cluster_resource_id")

    @property
    @pulumi.getter(name="copyTagsToSnapshot")
    def copy_tags_to_snapshot(self) -> Optional[builtins.bool]:
        """
        A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default behaviour is not to copy them.
        """
        return pulumi.get(self, "copy_tags_to_snapshot")

    @property
    @pulumi.getter(name="dbClusterParameterGroupName")
    def db_cluster_parameter_group_name(self) -> Optional[builtins.str]:
        """
        Provides the name of the DB cluster parameter group.
        """
        return pulumi.get(self, "db_cluster_parameter_group_name")

    @property
    @pulumi.getter(name="dbPort")
    def db_port(self) -> Optional[builtins.int]:
        """
        The port number on which the DB instances in the DB cluster accept connections. 

        If not specified, the default port used is `8182`. 

        Note: `Port` property will soon be deprecated from this resource. Please update existing templates to rename it with new property `DBPort` having same functionalities.
        """
        return pulumi.get(self, "db_port")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> Optional[builtins.bool]:
        """
        Indicates whether or not the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter(name="enableCloudwatchLogsExports")
    def enable_cloudwatch_logs_exports(self) -> Optional[Sequence[builtins.str]]:
        """
        Specifies a list of log types that are enabled for export to CloudWatch Logs.
        """
        return pulumi.get(self, "enable_cloudwatch_logs_exports")

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[builtins.str]:
        """
        The connection endpoint for the DB cluster. For example: `mystack-mydbcluster-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[builtins.str]:
        """
        Indicates the database engine version.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="iamAuthEnabled")
    def iam_auth_enabled(self) -> Optional[builtins.bool]:
        """
        True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.
        """
        return pulumi.get(self, "iam_auth_enabled")

    @property
    @pulumi.getter
    def port(self) -> Optional[builtins.str]:
        """
        The port number on which the DB cluster accepts connections. For example: `8182`.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="preferredBackupWindow")
    def preferred_backup_window(self) -> Optional[builtins.str]:
        """
        Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.
        """
        return pulumi.get(self, "preferred_backup_window")

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> Optional[builtins.str]:
        """
        Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
        """
        return pulumi.get(self, "preferred_maintenance_window")

    @property
    @pulumi.getter(name="readEndpoint")
    def read_endpoint(self) -> Optional[builtins.str]:
        """
        The reader endpoint for the DB cluster. For example: `mystack-mydbcluster-ro-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`
        """
        return pulumi.get(self, "read_endpoint")

    @property
    @pulumi.getter(name="serverlessScalingConfiguration")
    def serverless_scaling_configuration(self) -> Optional['outputs.DbClusterServerlessScalingConfiguration']:
        """
        Contains the scaling configuration used by the Neptune Serverless Instances within this DB cluster.
        """
        return pulumi.get(self, "serverless_scaling_configuration")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        The tags assigned to this cluster.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> Optional[Sequence[builtins.str]]:
        """
        Provides a list of VPC security groups that the DB cluster belongs to.
        """
        return pulumi.get(self, "vpc_security_group_ids")


class AwaitableGetDbClusterResult(GetDbClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDbClusterResult(
            associated_roles=self.associated_roles,
            backup_retention_period=self.backup_retention_period,
            cluster_resource_id=self.cluster_resource_id,
            copy_tags_to_snapshot=self.copy_tags_to_snapshot,
            db_cluster_parameter_group_name=self.db_cluster_parameter_group_name,
            db_port=self.db_port,
            deletion_protection=self.deletion_protection,
            enable_cloudwatch_logs_exports=self.enable_cloudwatch_logs_exports,
            endpoint=self.endpoint,
            engine_version=self.engine_version,
            iam_auth_enabled=self.iam_auth_enabled,
            port=self.port,
            preferred_backup_window=self.preferred_backup_window,
            preferred_maintenance_window=self.preferred_maintenance_window,
            read_endpoint=self.read_endpoint,
            serverless_scaling_configuration=self.serverless_scaling_configuration,
            tags=self.tags,
            vpc_security_group_ids=self.vpc_security_group_ids)


def get_db_cluster(db_cluster_identifier: Optional[builtins.str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDbClusterResult:
    """
    The AWS::Neptune::DBCluster resource creates an Amazon Neptune DB cluster.


    :param builtins.str db_cluster_identifier: The DB cluster identifier. Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string.
    """
    __args__ = dict()
    __args__['dbClusterIdentifier'] = db_cluster_identifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:neptune:getDbCluster', __args__, opts=opts, typ=GetDbClusterResult).value

    return AwaitableGetDbClusterResult(
        associated_roles=pulumi.get(__ret__, 'associated_roles'),
        backup_retention_period=pulumi.get(__ret__, 'backup_retention_period'),
        cluster_resource_id=pulumi.get(__ret__, 'cluster_resource_id'),
        copy_tags_to_snapshot=pulumi.get(__ret__, 'copy_tags_to_snapshot'),
        db_cluster_parameter_group_name=pulumi.get(__ret__, 'db_cluster_parameter_group_name'),
        db_port=pulumi.get(__ret__, 'db_port'),
        deletion_protection=pulumi.get(__ret__, 'deletion_protection'),
        enable_cloudwatch_logs_exports=pulumi.get(__ret__, 'enable_cloudwatch_logs_exports'),
        endpoint=pulumi.get(__ret__, 'endpoint'),
        engine_version=pulumi.get(__ret__, 'engine_version'),
        iam_auth_enabled=pulumi.get(__ret__, 'iam_auth_enabled'),
        port=pulumi.get(__ret__, 'port'),
        preferred_backup_window=pulumi.get(__ret__, 'preferred_backup_window'),
        preferred_maintenance_window=pulumi.get(__ret__, 'preferred_maintenance_window'),
        read_endpoint=pulumi.get(__ret__, 'read_endpoint'),
        serverless_scaling_configuration=pulumi.get(__ret__, 'serverless_scaling_configuration'),
        tags=pulumi.get(__ret__, 'tags'),
        vpc_security_group_ids=pulumi.get(__ret__, 'vpc_security_group_ids'))
def get_db_cluster_output(db_cluster_identifier: Optional[pulumi.Input[builtins.str]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDbClusterResult]:
    """
    The AWS::Neptune::DBCluster resource creates an Amazon Neptune DB cluster.


    :param builtins.str db_cluster_identifier: The DB cluster identifier. Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string.
    """
    __args__ = dict()
    __args__['dbClusterIdentifier'] = db_cluster_identifier
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:neptune:getDbCluster', __args__, opts=opts, typ=GetDbClusterResult)
    return __ret__.apply(lambda __response__: GetDbClusterResult(
        associated_roles=pulumi.get(__response__, 'associated_roles'),
        backup_retention_period=pulumi.get(__response__, 'backup_retention_period'),
        cluster_resource_id=pulumi.get(__response__, 'cluster_resource_id'),
        copy_tags_to_snapshot=pulumi.get(__response__, 'copy_tags_to_snapshot'),
        db_cluster_parameter_group_name=pulumi.get(__response__, 'db_cluster_parameter_group_name'),
        db_port=pulumi.get(__response__, 'db_port'),
        deletion_protection=pulumi.get(__response__, 'deletion_protection'),
        enable_cloudwatch_logs_exports=pulumi.get(__response__, 'enable_cloudwatch_logs_exports'),
        endpoint=pulumi.get(__response__, 'endpoint'),
        engine_version=pulumi.get(__response__, 'engine_version'),
        iam_auth_enabled=pulumi.get(__response__, 'iam_auth_enabled'),
        port=pulumi.get(__response__, 'port'),
        preferred_backup_window=pulumi.get(__response__, 'preferred_backup_window'),
        preferred_maintenance_window=pulumi.get(__response__, 'preferred_maintenance_window'),
        read_endpoint=pulumi.get(__response__, 'read_endpoint'),
        serverless_scaling_configuration=pulumi.get(__response__, 'serverless_scaling_configuration'),
        tags=pulumi.get(__response__, 'tags'),
        vpc_security_group_ids=pulumi.get(__response__, 'vpc_security_group_ids')))
