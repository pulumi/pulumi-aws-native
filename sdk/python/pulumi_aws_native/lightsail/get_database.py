# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
    'get_database_output',
]

@pulumi.output_type
class GetDatabaseResult:
    def __init__(__self__, backup_retention=None, ca_certificate_identifier=None, database_arn=None, preferred_backup_window=None, preferred_maintenance_window=None, publicly_accessible=None, tags=None):
        if backup_retention and not isinstance(backup_retention, bool):
            raise TypeError("Expected argument 'backup_retention' to be a bool")
        pulumi.set(__self__, "backup_retention", backup_retention)
        if ca_certificate_identifier and not isinstance(ca_certificate_identifier, str):
            raise TypeError("Expected argument 'ca_certificate_identifier' to be a str")
        pulumi.set(__self__, "ca_certificate_identifier", ca_certificate_identifier)
        if database_arn and not isinstance(database_arn, str):
            raise TypeError("Expected argument 'database_arn' to be a str")
        pulumi.set(__self__, "database_arn", database_arn)
        if preferred_backup_window and not isinstance(preferred_backup_window, str):
            raise TypeError("Expected argument 'preferred_backup_window' to be a str")
        pulumi.set(__self__, "preferred_backup_window", preferred_backup_window)
        if preferred_maintenance_window and not isinstance(preferred_maintenance_window, str):
            raise TypeError("Expected argument 'preferred_maintenance_window' to be a str")
        pulumi.set(__self__, "preferred_maintenance_window", preferred_maintenance_window)
        if publicly_accessible and not isinstance(publicly_accessible, bool):
            raise TypeError("Expected argument 'publicly_accessible' to be a bool")
        pulumi.set(__self__, "publicly_accessible", publicly_accessible)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="backupRetention")
    def backup_retention(self) -> Optional[bool]:
        """
        When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.
        """
        return pulumi.get(self, "backup_retention")

    @property
    @pulumi.getter(name="caCertificateIdentifier")
    def ca_certificate_identifier(self) -> Optional[str]:
        """
        Indicates the certificate that needs to be associated with the database.
        """
        return pulumi.get(self, "ca_certificate_identifier")

    @property
    @pulumi.getter(name="databaseArn")
    def database_arn(self) -> Optional[str]:
        return pulumi.get(self, "database_arn")

    @property
    @pulumi.getter(name="preferredBackupWindow")
    def preferred_backup_window(self) -> Optional[str]:
        """
        The daily time range during which automated backups are created for your new database if automated backups are enabled.
        """
        return pulumi.get(self, "preferred_backup_window")

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> Optional[str]:
        """
        The weekly time range during which system maintenance can occur on your new database.
        """
        return pulumi.get(self, "preferred_maintenance_window")

    @property
    @pulumi.getter(name="publiclyAccessible")
    def publicly_accessible(self) -> Optional[bool]:
        """
        Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.
        """
        return pulumi.get(self, "publicly_accessible")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DatabaseTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            backup_retention=self.backup_retention,
            ca_certificate_identifier=self.ca_certificate_identifier,
            database_arn=self.database_arn,
            preferred_backup_window=self.preferred_backup_window,
            preferred_maintenance_window=self.preferred_maintenance_window,
            publicly_accessible=self.publicly_accessible,
            tags=self.tags)


def get_database(relational_database_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Resource Type definition for AWS::Lightsail::Database


    :param str relational_database_name: The name to use for your new Lightsail database resource.
    """
    __args__ = dict()
    __args__['relationalDatabaseName'] = relational_database_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:lightsail:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        backup_retention=__ret__.backup_retention,
        ca_certificate_identifier=__ret__.ca_certificate_identifier,
        database_arn=__ret__.database_arn,
        preferred_backup_window=__ret__.preferred_backup_window,
        preferred_maintenance_window=__ret__.preferred_maintenance_window,
        publicly_accessible=__ret__.publicly_accessible,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_database)
def get_database_output(relational_database_name: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseResult]:
    """
    Resource Type definition for AWS::Lightsail::Database


    :param str relational_database_name: The name to use for your new Lightsail database resource.
    """
    ...