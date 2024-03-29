# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'MigrationProjectDataProviderDescriptorArgs',
    'ReplicationConfigComputeConfigArgs',
    'SchemaConversionApplicationAttributesPropertiesArgs',
    'Settings0PropertiesPostgreSqlSettingsPropertiesArgs',
    'Settings0PropertiesArgs',
    'Settings1PropertiesMySqlSettingsPropertiesArgs',
    'Settings1PropertiesArgs',
    'Settings2PropertiesOracleSettingsPropertiesArgs',
    'Settings2PropertiesArgs',
    'Settings3PropertiesMicrosoftSqlServerSettingsPropertiesArgs',
    'Settings3PropertiesArgs',
]

@pulumi.input_type
class MigrationProjectDataProviderDescriptorArgs:
    def __init__(__self__, *,
                 data_provider_arn: Optional[pulumi.Input[str]] = None,
                 data_provider_identifier: Optional[pulumi.Input[str]] = None,
                 data_provider_name: Optional[pulumi.Input[str]] = None,
                 secrets_manager_access_role_arn: Optional[pulumi.Input[str]] = None,
                 secrets_manager_secret_id: Optional[pulumi.Input[str]] = None):
        """
        It is an object that describes Source and Target DataProviders and credentials for connecting to databases that are used in MigrationProject
        """
        if data_provider_arn is not None:
            pulumi.set(__self__, "data_provider_arn", data_provider_arn)
        if data_provider_identifier is not None:
            pulumi.set(__self__, "data_provider_identifier", data_provider_identifier)
        if data_provider_name is not None:
            pulumi.set(__self__, "data_provider_name", data_provider_name)
        if secrets_manager_access_role_arn is not None:
            pulumi.set(__self__, "secrets_manager_access_role_arn", secrets_manager_access_role_arn)
        if secrets_manager_secret_id is not None:
            pulumi.set(__self__, "secrets_manager_secret_id", secrets_manager_secret_id)

    @property
    @pulumi.getter(name="dataProviderArn")
    def data_provider_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data_provider_arn")

    @data_provider_arn.setter
    def data_provider_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_provider_arn", value)

    @property
    @pulumi.getter(name="dataProviderIdentifier")
    def data_provider_identifier(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data_provider_identifier")

    @data_provider_identifier.setter
    def data_provider_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_provider_identifier", value)

    @property
    @pulumi.getter(name="dataProviderName")
    def data_provider_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data_provider_name")

    @data_provider_name.setter
    def data_provider_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_provider_name", value)

    @property
    @pulumi.getter(name="secretsManagerAccessRoleArn")
    def secrets_manager_access_role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secrets_manager_access_role_arn")

    @secrets_manager_access_role_arn.setter
    def secrets_manager_access_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secrets_manager_access_role_arn", value)

    @property
    @pulumi.getter(name="secretsManagerSecretId")
    def secrets_manager_secret_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secrets_manager_secret_id")

    @secrets_manager_secret_id.setter
    def secrets_manager_secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secrets_manager_secret_id", value)


@pulumi.input_type
class ReplicationConfigComputeConfigArgs:
    def __init__(__self__, *,
                 max_capacity_units: pulumi.Input[int],
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 dns_name_servers: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 min_capacity_units: Optional[pulumi.Input[int]] = None,
                 multi_az: Optional[pulumi.Input[bool]] = None,
                 preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
                 replication_subnet_group_id: Optional[pulumi.Input[str]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Configuration parameters for provisioning a AWS DMS Serverless replication
        """
        pulumi.set(__self__, "max_capacity_units", max_capacity_units)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if dns_name_servers is not None:
            pulumi.set(__self__, "dns_name_servers", dns_name_servers)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if min_capacity_units is not None:
            pulumi.set(__self__, "min_capacity_units", min_capacity_units)
        if multi_az is not None:
            pulumi.set(__self__, "multi_az", multi_az)
        if preferred_maintenance_window is not None:
            pulumi.set(__self__, "preferred_maintenance_window", preferred_maintenance_window)
        if replication_subnet_group_id is not None:
            pulumi.set(__self__, "replication_subnet_group_id", replication_subnet_group_id)
        if vpc_security_group_ids is not None:
            pulumi.set(__self__, "vpc_security_group_ids", vpc_security_group_ids)

    @property
    @pulumi.getter(name="maxCapacityUnits")
    def max_capacity_units(self) -> pulumi.Input[int]:
        return pulumi.get(self, "max_capacity_units")

    @max_capacity_units.setter
    def max_capacity_units(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_capacity_units", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="dnsNameServers")
    def dns_name_servers(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dns_name_servers")

    @dns_name_servers.setter
    def dns_name_servers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_name_servers", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="minCapacityUnits")
    def min_capacity_units(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "min_capacity_units")

    @min_capacity_units.setter
    def min_capacity_units(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_capacity_units", value)

    @property
    @pulumi.getter(name="multiAz")
    def multi_az(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "multi_az")

    @multi_az.setter
    def multi_az(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_az", value)

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "preferred_maintenance_window")

    @preferred_maintenance_window.setter
    def preferred_maintenance_window(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "preferred_maintenance_window", value)

    @property
    @pulumi.getter(name="replicationSubnetGroupId")
    def replication_subnet_group_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "replication_subnet_group_id")

    @replication_subnet_group_id.setter
    def replication_subnet_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_subnet_group_id", value)

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "vpc_security_group_ids")

    @vpc_security_group_ids.setter
    def vpc_security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_security_group_ids", value)


@pulumi.input_type
class SchemaConversionApplicationAttributesPropertiesArgs:
    def __init__(__self__, *,
                 s3_bucket_path: Optional[pulumi.Input[str]] = None,
                 s3_bucket_role_arn: Optional[pulumi.Input[str]] = None):
        """
        The property describes schema conversion application attributes for the migration project.
        """
        if s3_bucket_path is not None:
            pulumi.set(__self__, "s3_bucket_path", s3_bucket_path)
        if s3_bucket_role_arn is not None:
            pulumi.set(__self__, "s3_bucket_role_arn", s3_bucket_role_arn)

    @property
    @pulumi.getter(name="s3BucketPath")
    def s3_bucket_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "s3_bucket_path")

    @s3_bucket_path.setter
    def s3_bucket_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_bucket_path", value)

    @property
    @pulumi.getter(name="s3BucketRoleArn")
    def s3_bucket_role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "s3_bucket_role_arn")

    @s3_bucket_role_arn.setter
    def s3_bucket_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_bucket_role_arn", value)


@pulumi.input_type
class Settings0PropertiesPostgreSqlSettingsPropertiesArgs:
    def __init__(__self__, *,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input['DataProviderDmsSslModeValue']] = None):
        if certificate_arn is not None:
            pulumi.set(__self__, "certificate_arn", certificate_arn)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if ssl_mode is not None:
            pulumi.set(__self__, "ssl_mode", ssl_mode)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate_arn")

    @certificate_arn.setter
    def certificate_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_arn", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="sslMode")
    def ssl_mode(self) -> Optional[pulumi.Input['DataProviderDmsSslModeValue']]:
        return pulumi.get(self, "ssl_mode")

    @ssl_mode.setter
    def ssl_mode(self, value: Optional[pulumi.Input['DataProviderDmsSslModeValue']]):
        pulumi.set(self, "ssl_mode", value)


@pulumi.input_type
class Settings0PropertiesArgs:
    def __init__(__self__, *,
                 postgre_sql_settings: Optional[pulumi.Input['Settings0PropertiesPostgreSqlSettingsPropertiesArgs']] = None):
        """
        PostgreSqlSettings property identifier.
        """
        if postgre_sql_settings is not None:
            pulumi.set(__self__, "postgre_sql_settings", postgre_sql_settings)

    @property
    @pulumi.getter(name="postgreSqlSettings")
    def postgre_sql_settings(self) -> Optional[pulumi.Input['Settings0PropertiesPostgreSqlSettingsPropertiesArgs']]:
        return pulumi.get(self, "postgre_sql_settings")

    @postgre_sql_settings.setter
    def postgre_sql_settings(self, value: Optional[pulumi.Input['Settings0PropertiesPostgreSqlSettingsPropertiesArgs']]):
        pulumi.set(self, "postgre_sql_settings", value)


@pulumi.input_type
class Settings1PropertiesMySqlSettingsPropertiesArgs:
    def __init__(__self__, *,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input['DataProviderDmsSslModeValue']] = None):
        if certificate_arn is not None:
            pulumi.set(__self__, "certificate_arn", certificate_arn)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if ssl_mode is not None:
            pulumi.set(__self__, "ssl_mode", ssl_mode)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate_arn")

    @certificate_arn.setter
    def certificate_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_arn", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="sslMode")
    def ssl_mode(self) -> Optional[pulumi.Input['DataProviderDmsSslModeValue']]:
        return pulumi.get(self, "ssl_mode")

    @ssl_mode.setter
    def ssl_mode(self, value: Optional[pulumi.Input['DataProviderDmsSslModeValue']]):
        pulumi.set(self, "ssl_mode", value)


@pulumi.input_type
class Settings1PropertiesArgs:
    def __init__(__self__, *,
                 my_sql_settings: Optional[pulumi.Input['Settings1PropertiesMySqlSettingsPropertiesArgs']] = None):
        """
        MySqlSettings property identifier.
        """
        if my_sql_settings is not None:
            pulumi.set(__self__, "my_sql_settings", my_sql_settings)

    @property
    @pulumi.getter(name="mySqlSettings")
    def my_sql_settings(self) -> Optional[pulumi.Input['Settings1PropertiesMySqlSettingsPropertiesArgs']]:
        return pulumi.get(self, "my_sql_settings")

    @my_sql_settings.setter
    def my_sql_settings(self, value: Optional[pulumi.Input['Settings1PropertiesMySqlSettingsPropertiesArgs']]):
        pulumi.set(self, "my_sql_settings", value)


@pulumi.input_type
class Settings2PropertiesOracleSettingsPropertiesArgs:
    def __init__(__self__, *,
                 asm_server: Optional[pulumi.Input[str]] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secrets_manager_oracle_asm_access_role_arn: Optional[pulumi.Input[str]] = None,
                 secrets_manager_oracle_asm_secret_id: Optional[pulumi.Input[str]] = None,
                 secrets_manager_security_db_encryption_access_role_arn: Optional[pulumi.Input[str]] = None,
                 secrets_manager_security_db_encryption_secret_id: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input['DataProviderDmsSslModeValue']] = None):
        if asm_server is not None:
            pulumi.set(__self__, "asm_server", asm_server)
        if certificate_arn is not None:
            pulumi.set(__self__, "certificate_arn", certificate_arn)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if secrets_manager_oracle_asm_access_role_arn is not None:
            pulumi.set(__self__, "secrets_manager_oracle_asm_access_role_arn", secrets_manager_oracle_asm_access_role_arn)
        if secrets_manager_oracle_asm_secret_id is not None:
            pulumi.set(__self__, "secrets_manager_oracle_asm_secret_id", secrets_manager_oracle_asm_secret_id)
        if secrets_manager_security_db_encryption_access_role_arn is not None:
            pulumi.set(__self__, "secrets_manager_security_db_encryption_access_role_arn", secrets_manager_security_db_encryption_access_role_arn)
        if secrets_manager_security_db_encryption_secret_id is not None:
            pulumi.set(__self__, "secrets_manager_security_db_encryption_secret_id", secrets_manager_security_db_encryption_secret_id)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if ssl_mode is not None:
            pulumi.set(__self__, "ssl_mode", ssl_mode)

    @property
    @pulumi.getter(name="asmServer")
    def asm_server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "asm_server")

    @asm_server.setter
    def asm_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asm_server", value)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate_arn")

    @certificate_arn.setter
    def certificate_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_arn", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="secretsManagerOracleAsmAccessRoleArn")
    def secrets_manager_oracle_asm_access_role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secrets_manager_oracle_asm_access_role_arn")

    @secrets_manager_oracle_asm_access_role_arn.setter
    def secrets_manager_oracle_asm_access_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secrets_manager_oracle_asm_access_role_arn", value)

    @property
    @pulumi.getter(name="secretsManagerOracleAsmSecretId")
    def secrets_manager_oracle_asm_secret_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secrets_manager_oracle_asm_secret_id")

    @secrets_manager_oracle_asm_secret_id.setter
    def secrets_manager_oracle_asm_secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secrets_manager_oracle_asm_secret_id", value)

    @property
    @pulumi.getter(name="secretsManagerSecurityDbEncryptionAccessRoleArn")
    def secrets_manager_security_db_encryption_access_role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secrets_manager_security_db_encryption_access_role_arn")

    @secrets_manager_security_db_encryption_access_role_arn.setter
    def secrets_manager_security_db_encryption_access_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secrets_manager_security_db_encryption_access_role_arn", value)

    @property
    @pulumi.getter(name="secretsManagerSecurityDbEncryptionSecretId")
    def secrets_manager_security_db_encryption_secret_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secrets_manager_security_db_encryption_secret_id")

    @secrets_manager_security_db_encryption_secret_id.setter
    def secrets_manager_security_db_encryption_secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secrets_manager_security_db_encryption_secret_id", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="sslMode")
    def ssl_mode(self) -> Optional[pulumi.Input['DataProviderDmsSslModeValue']]:
        return pulumi.get(self, "ssl_mode")

    @ssl_mode.setter
    def ssl_mode(self, value: Optional[pulumi.Input['DataProviderDmsSslModeValue']]):
        pulumi.set(self, "ssl_mode", value)


@pulumi.input_type
class Settings2PropertiesArgs:
    def __init__(__self__, *,
                 oracle_settings: Optional[pulumi.Input['Settings2PropertiesOracleSettingsPropertiesArgs']] = None):
        """
        OracleSettings property identifier.
        """
        if oracle_settings is not None:
            pulumi.set(__self__, "oracle_settings", oracle_settings)

    @property
    @pulumi.getter(name="oracleSettings")
    def oracle_settings(self) -> Optional[pulumi.Input['Settings2PropertiesOracleSettingsPropertiesArgs']]:
        return pulumi.get(self, "oracle_settings")

    @oracle_settings.setter
    def oracle_settings(self, value: Optional[pulumi.Input['Settings2PropertiesOracleSettingsPropertiesArgs']]):
        pulumi.set(self, "oracle_settings", value)


@pulumi.input_type
class Settings3PropertiesMicrosoftSqlServerSettingsPropertiesArgs:
    def __init__(__self__, *,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input['DataProviderDmsSslModeValue']] = None):
        if certificate_arn is not None:
            pulumi.set(__self__, "certificate_arn", certificate_arn)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if ssl_mode is not None:
            pulumi.set(__self__, "ssl_mode", ssl_mode)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate_arn")

    @certificate_arn.setter
    def certificate_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_arn", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="sslMode")
    def ssl_mode(self) -> Optional[pulumi.Input['DataProviderDmsSslModeValue']]:
        return pulumi.get(self, "ssl_mode")

    @ssl_mode.setter
    def ssl_mode(self, value: Optional[pulumi.Input['DataProviderDmsSslModeValue']]):
        pulumi.set(self, "ssl_mode", value)


@pulumi.input_type
class Settings3PropertiesArgs:
    def __init__(__self__, *,
                 microsoft_sql_server_settings: Optional[pulumi.Input['Settings3PropertiesMicrosoftSqlServerSettingsPropertiesArgs']] = None):
        """
        MicrosoftSqlServerSettings property identifier.
        """
        if microsoft_sql_server_settings is not None:
            pulumi.set(__self__, "microsoft_sql_server_settings", microsoft_sql_server_settings)

    @property
    @pulumi.getter(name="microsoftSqlServerSettings")
    def microsoft_sql_server_settings(self) -> Optional[pulumi.Input['Settings3PropertiesMicrosoftSqlServerSettingsPropertiesArgs']]:
        return pulumi.get(self, "microsoft_sql_server_settings")

    @microsoft_sql_server_settings.setter
    def microsoft_sql_server_settings(self, value: Optional[pulumi.Input['Settings3PropertiesMicrosoftSqlServerSettingsPropertiesArgs']]):
        pulumi.set(self, "microsoft_sql_server_settings", value)


