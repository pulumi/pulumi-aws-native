# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'PrincipalPermissionsPermission',
    'PrincipalPermissionsResourceType',
]


class PrincipalPermissionsPermission(str, Enum):
    ALL = "ALL"
    SELECT = "SELECT"
    ALTER = "ALTER"
    DROP = "DROP"
    DELETE = "DELETE"
    INSERT = "INSERT"
    DESCRIBE = "DESCRIBE"
    CREATE_DATABASE = "CREATE_DATABASE"
    CREATE_TABLE = "CREATE_TABLE"
    DATA_LOCATION_ACCESS = "DATA_LOCATION_ACCESS"
    CREATE_TAG = "CREATE_TAG"
    ASSOCIATE = "ASSOCIATE"


class PrincipalPermissionsResourceType(str, Enum):
    DATABASE = "DATABASE"
    TABLE = "TABLE"