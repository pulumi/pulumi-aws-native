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
from ._enums import *

__all__ = [
    'KeyspaceTag',
    'TableBillingMode',
    'TableClusteringKeyColumn',
    'TableColumn',
    'TableEncryptionSpecification',
    'TableProvisionedThroughput',
    'TableTag',
]

@pulumi.output_type
class KeyspaceTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class TableBillingMode(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "provisionedThroughput":
            suggest = "provisioned_throughput"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TableBillingMode. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TableBillingMode.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TableBillingMode.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 mode: 'TableMode',
                 provisioned_throughput: Optional['outputs.TableProvisionedThroughput'] = None):
        pulumi.set(__self__, "mode", mode)
        if provisioned_throughput is not None:
            pulumi.set(__self__, "provisioned_throughput", provisioned_throughput)

    @property
    @pulumi.getter
    def mode(self) -> 'TableMode':
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="provisionedThroughput")
    def provisioned_throughput(self) -> Optional['outputs.TableProvisionedThroughput']:
        return pulumi.get(self, "provisioned_throughput")


@pulumi.output_type
class TableClusteringKeyColumn(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "orderBy":
            suggest = "order_by"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TableClusteringKeyColumn. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TableClusteringKeyColumn.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TableClusteringKeyColumn.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 column: 'outputs.TableColumn',
                 order_by: Optional['TableClusteringKeyColumnOrderBy'] = None):
        pulumi.set(__self__, "column", column)
        if order_by is not None:
            pulumi.set(__self__, "order_by", order_by)

    @property
    @pulumi.getter
    def column(self) -> 'outputs.TableColumn':
        return pulumi.get(self, "column")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional['TableClusteringKeyColumnOrderBy']:
        return pulumi.get(self, "order_by")


@pulumi.output_type
class TableColumn(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "columnName":
            suggest = "column_name"
        elif key == "columnType":
            suggest = "column_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TableColumn. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TableColumn.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TableColumn.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 column_name: str,
                 column_type: str):
        pulumi.set(__self__, "column_name", column_name)
        pulumi.set(__self__, "column_type", column_type)

    @property
    @pulumi.getter(name="columnName")
    def column_name(self) -> str:
        return pulumi.get(self, "column_name")

    @property
    @pulumi.getter(name="columnType")
    def column_type(self) -> str:
        return pulumi.get(self, "column_type")


@pulumi.output_type
class TableEncryptionSpecification(dict):
    """
    Represents the settings used to enable server-side encryption
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionType":
            suggest = "encryption_type"
        elif key == "kmsKeyIdentifier":
            suggest = "kms_key_identifier"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TableEncryptionSpecification. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TableEncryptionSpecification.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TableEncryptionSpecification.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 encryption_type: 'TableEncryptionType',
                 kms_key_identifier: Optional[str] = None):
        """
        Represents the settings used to enable server-side encryption
        """
        pulumi.set(__self__, "encryption_type", encryption_type)
        if kms_key_identifier is not None:
            pulumi.set(__self__, "kms_key_identifier", kms_key_identifier)

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> 'TableEncryptionType':
        return pulumi.get(self, "encryption_type")

    @property
    @pulumi.getter(name="kmsKeyIdentifier")
    def kms_key_identifier(self) -> Optional[str]:
        return pulumi.get(self, "kms_key_identifier")


@pulumi.output_type
class TableProvisionedThroughput(dict):
    """
    Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "readCapacityUnits":
            suggest = "read_capacity_units"
        elif key == "writeCapacityUnits":
            suggest = "write_capacity_units"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TableProvisionedThroughput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TableProvisionedThroughput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TableProvisionedThroughput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 read_capacity_units: int,
                 write_capacity_units: int):
        """
        Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits
        """
        pulumi.set(__self__, "read_capacity_units", read_capacity_units)
        pulumi.set(__self__, "write_capacity_units", write_capacity_units)

    @property
    @pulumi.getter(name="readCapacityUnits")
    def read_capacity_units(self) -> int:
        return pulumi.get(self, "read_capacity_units")

    @property
    @pulumi.getter(name="writeCapacityUnits")
    def write_capacity_units(self) -> int:
        return pulumi.get(self, "write_capacity_units")


@pulumi.output_type
class TableTag(dict):
    """
    A key-value pair to apply to the resource
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A key-value pair to apply to the resource
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

