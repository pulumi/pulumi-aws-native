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
    'AccessPointCreationInfo',
    'AccessPointPosixUser',
    'AccessPointRootDirectory',
    'AccessPointTag',
    'FileSystemBackupPolicy',
    'FileSystemElasticFileSystemTag',
    'FileSystemLifecyclePolicy',
]

@pulumi.output_type
class AccessPointCreationInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ownerGid":
            suggest = "owner_gid"
        elif key == "ownerUid":
            suggest = "owner_uid"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccessPointCreationInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccessPointCreationInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccessPointCreationInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 owner_gid: str,
                 owner_uid: str,
                 permissions: str):
        """
        :param str owner_gid: Specifies the POSIX group ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).
        :param str owner_uid: Specifies the POSIX user ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).
        :param str permissions: Specifies the POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.
        """
        pulumi.set(__self__, "owner_gid", owner_gid)
        pulumi.set(__self__, "owner_uid", owner_uid)
        pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="ownerGid")
    def owner_gid(self) -> str:
        """
        Specifies the POSIX group ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).
        """
        return pulumi.get(self, "owner_gid")

    @property
    @pulumi.getter(name="ownerUid")
    def owner_uid(self) -> str:
        """
        Specifies the POSIX user ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).
        """
        return pulumi.get(self, "owner_uid")

    @property
    @pulumi.getter
    def permissions(self) -> str:
        """
        Specifies the POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.
        """
        return pulumi.get(self, "permissions")


@pulumi.output_type
class AccessPointPosixUser(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "secondaryGids":
            suggest = "secondary_gids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccessPointPosixUser. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccessPointPosixUser.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccessPointPosixUser.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 gid: str,
                 uid: str,
                 secondary_gids: Optional[Sequence[str]] = None):
        """
        :param str gid: The POSIX group ID used for all file system operations using this access point.
        :param str uid: The POSIX user ID used for all file system operations using this access point.
        :param Sequence[str] secondary_gids: Secondary POSIX group IDs used for all file system operations using this access point.
        """
        pulumi.set(__self__, "gid", gid)
        pulumi.set(__self__, "uid", uid)
        if secondary_gids is not None:
            pulumi.set(__self__, "secondary_gids", secondary_gids)

    @property
    @pulumi.getter
    def gid(self) -> str:
        """
        The POSIX group ID used for all file system operations using this access point.
        """
        return pulumi.get(self, "gid")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        The POSIX user ID used for all file system operations using this access point.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="secondaryGids")
    def secondary_gids(self) -> Optional[Sequence[str]]:
        """
        Secondary POSIX group IDs used for all file system operations using this access point.
        """
        return pulumi.get(self, "secondary_gids")


@pulumi.output_type
class AccessPointRootDirectory(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "creationInfo":
            suggest = "creation_info"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccessPointRootDirectory. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccessPointRootDirectory.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccessPointRootDirectory.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 creation_info: Optional['outputs.AccessPointCreationInfo'] = None,
                 path: Optional[str] = None):
        """
        :param 'AccessPointCreationInfo' creation_info: (Optional) Specifies the POSIX IDs and permissions to apply to the access point's RootDirectory. If the RootDirectory>Path specified does not exist, EFS creates the root directory using the CreationInfo settings when a client connects to an access point. When specifying the CreationInfo, you must provide values for all properties.   If you do not provide CreationInfo and the specified RootDirectory>Path does not exist, attempts to mount the file system using the access point will fail. 
        :param str path: Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the CreationInfo.
        """
        if creation_info is not None:
            pulumi.set(__self__, "creation_info", creation_info)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter(name="creationInfo")
    def creation_info(self) -> Optional['outputs.AccessPointCreationInfo']:
        """
        (Optional) Specifies the POSIX IDs and permissions to apply to the access point's RootDirectory. If the RootDirectory>Path specified does not exist, EFS creates the root directory using the CreationInfo settings when a client connects to an access point. When specifying the CreationInfo, you must provide values for all properties.   If you do not provide CreationInfo and the specified RootDirectory>Path does not exist, attempts to mount the file system using the access point will fail. 
        """
        return pulumi.get(self, "creation_info")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        """
        Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the CreationInfo.
        """
        return pulumi.get(self, "path")


@pulumi.output_type
class AccessPointTag(dict):
    def __init__(__self__, *,
                 key: Optional[str] = None,
                 value: Optional[str] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


@pulumi.output_type
class FileSystemBackupPolicy(dict):
    def __init__(__self__, *,
                 status: str):
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")


@pulumi.output_type
class FileSystemElasticFileSystemTag(dict):
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
class FileSystemLifecyclePolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "transitionToIA":
            suggest = "transition_to_ia"
        elif key == "transitionToPrimaryStorageClass":
            suggest = "transition_to_primary_storage_class"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FileSystemLifecyclePolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FileSystemLifecyclePolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FileSystemLifecyclePolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 transition_to_ia: Optional[str] = None,
                 transition_to_primary_storage_class: Optional[str] = None):
        if transition_to_ia is not None:
            pulumi.set(__self__, "transition_to_ia", transition_to_ia)
        if transition_to_primary_storage_class is not None:
            pulumi.set(__self__, "transition_to_primary_storage_class", transition_to_primary_storage_class)

    @property
    @pulumi.getter(name="transitionToIA")
    def transition_to_ia(self) -> Optional[str]:
        return pulumi.get(self, "transition_to_ia")

    @property
    @pulumi.getter(name="transitionToPrimaryStorageClass")
    def transition_to_primary_storage_class(self) -> Optional[str]:
        return pulumi.get(self, "transition_to_primary_storage_class")

