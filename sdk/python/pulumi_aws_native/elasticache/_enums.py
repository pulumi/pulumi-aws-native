# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'GlobalReplicationGroupMemberRole',
    'UserEngine',
    'UserGroupEngine',
]


class GlobalReplicationGroupMemberRole(str, Enum):
    """
    Indicates the role of the member, primary or secondary.
    """
    PRIMARY = "PRIMARY"
    SECONDARY = "SECONDARY"


class UserEngine(str, Enum):
    """
    Must be redis.
    """
    REDIS = "redis"


class UserGroupEngine(str, Enum):
    """
    Must be redis.
    """
    REDIS = "redis"