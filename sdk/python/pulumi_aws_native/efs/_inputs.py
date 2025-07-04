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
from ._enums import *

__all__ = [
    'AccessPointCreationInfoArgs',
    'AccessPointCreationInfoArgsDict',
    'AccessPointPosixUserArgs',
    'AccessPointPosixUserArgsDict',
    'AccessPointRootDirectoryArgs',
    'AccessPointRootDirectoryArgsDict',
    'FileSystemBackupPolicyArgs',
    'FileSystemBackupPolicyArgsDict',
    'FileSystemLifecyclePolicyArgs',
    'FileSystemLifecyclePolicyArgsDict',
    'FileSystemProtectionArgs',
    'FileSystemProtectionArgsDict',
    'FileSystemReplicationConfigurationArgs',
    'FileSystemReplicationConfigurationArgsDict',
    'FileSystemReplicationDestinationArgs',
    'FileSystemReplicationDestinationArgsDict',
]

MYPY = False

if not MYPY:
    class AccessPointCreationInfoArgsDict(TypedDict):
        """
        Required if the ``RootDirectory`` > ``Path`` specified does not exist. Specifies the POSIX IDs and permissions to apply to the access point's ``RootDirectory`` > ``Path``. If the access point root directory does not exist, EFS creates it with these settings when a client connects to the access point. When specifying ``CreationInfo``, you must include values for all properties. 
         Amazon EFS creates a root directory only if you have provided the CreationInfo: OwnUid, OwnGID, and permissions for the directory. If you do not provide this information, Amazon EFS does not create the root directory. If the root directory does not exist, attempts to mount using the access point will fail.
          If you do not provide ``CreationInfo`` and the specified ``RootDirectory`` does not exist, attempts to mount the file system using the access point will fail.
        """
        owner_gid: pulumi.Input[builtins.str]
        """
        Specifies the POSIX group ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).
        """
        owner_uid: pulumi.Input[builtins.str]
        """
        Specifies the POSIX user ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).
        """
        permissions: pulumi.Input[builtins.str]
        """
        Specifies the POSIX permissions to apply to the ``RootDirectory``, in the format of an octal number representing the file's mode bits.
        """
elif False:
    AccessPointCreationInfoArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AccessPointCreationInfoArgs:
    def __init__(__self__, *,
                 owner_gid: pulumi.Input[builtins.str],
                 owner_uid: pulumi.Input[builtins.str],
                 permissions: pulumi.Input[builtins.str]):
        """
        Required if the ``RootDirectory`` > ``Path`` specified does not exist. Specifies the POSIX IDs and permissions to apply to the access point's ``RootDirectory`` > ``Path``. If the access point root directory does not exist, EFS creates it with these settings when a client connects to the access point. When specifying ``CreationInfo``, you must include values for all properties. 
         Amazon EFS creates a root directory only if you have provided the CreationInfo: OwnUid, OwnGID, and permissions for the directory. If you do not provide this information, Amazon EFS does not create the root directory. If the root directory does not exist, attempts to mount using the access point will fail.
          If you do not provide ``CreationInfo`` and the specified ``RootDirectory`` does not exist, attempts to mount the file system using the access point will fail.
        :param pulumi.Input[builtins.str] owner_gid: Specifies the POSIX group ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).
        :param pulumi.Input[builtins.str] owner_uid: Specifies the POSIX user ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).
        :param pulumi.Input[builtins.str] permissions: Specifies the POSIX permissions to apply to the ``RootDirectory``, in the format of an octal number representing the file's mode bits.
        """
        pulumi.set(__self__, "owner_gid", owner_gid)
        pulumi.set(__self__, "owner_uid", owner_uid)
        pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="ownerGid")
    def owner_gid(self) -> pulumi.Input[builtins.str]:
        """
        Specifies the POSIX group ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).
        """
        return pulumi.get(self, "owner_gid")

    @owner_gid.setter
    def owner_gid(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "owner_gid", value)

    @property
    @pulumi.getter(name="ownerUid")
    def owner_uid(self) -> pulumi.Input[builtins.str]:
        """
        Specifies the POSIX user ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).
        """
        return pulumi.get(self, "owner_uid")

    @owner_uid.setter
    def owner_uid(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "owner_uid", value)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[builtins.str]:
        """
        Specifies the POSIX permissions to apply to the ``RootDirectory``, in the format of an octal number representing the file's mode bits.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "permissions", value)


if not MYPY:
    class AccessPointPosixUserArgsDict(TypedDict):
        """
        The full POSIX identity, including the user ID, group ID, and any secondary group IDs, on the access point that is used for all file system operations performed by NFS clients using the access point.
        """
        gid: pulumi.Input[builtins.str]
        """
        The POSIX group ID used for all file system operations using this access point.
        """
        uid: pulumi.Input[builtins.str]
        """
        The POSIX user ID used for all file system operations using this access point.
        """
        secondary_gids: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        """
        Secondary POSIX group IDs used for all file system operations using this access point.
        """
elif False:
    AccessPointPosixUserArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AccessPointPosixUserArgs:
    def __init__(__self__, *,
                 gid: pulumi.Input[builtins.str],
                 uid: pulumi.Input[builtins.str],
                 secondary_gids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The full POSIX identity, including the user ID, group ID, and any secondary group IDs, on the access point that is used for all file system operations performed by NFS clients using the access point.
        :param pulumi.Input[builtins.str] gid: The POSIX group ID used for all file system operations using this access point.
        :param pulumi.Input[builtins.str] uid: The POSIX user ID used for all file system operations using this access point.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] secondary_gids: Secondary POSIX group IDs used for all file system operations using this access point.
        """
        pulumi.set(__self__, "gid", gid)
        pulumi.set(__self__, "uid", uid)
        if secondary_gids is not None:
            pulumi.set(__self__, "secondary_gids", secondary_gids)

    @property
    @pulumi.getter
    def gid(self) -> pulumi.Input[builtins.str]:
        """
        The POSIX group ID used for all file system operations using this access point.
        """
        return pulumi.get(self, "gid")

    @gid.setter
    def gid(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "gid", value)

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Input[builtins.str]:
        """
        The POSIX user ID used for all file system operations using this access point.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "uid", value)

    @property
    @pulumi.getter(name="secondaryGids")
    def secondary_gids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Secondary POSIX group IDs used for all file system operations using this access point.
        """
        return pulumi.get(self, "secondary_gids")

    @secondary_gids.setter
    def secondary_gids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "secondary_gids", value)


if not MYPY:
    class AccessPointRootDirectoryArgsDict(TypedDict):
        """
        Specifies the directory on the Amazon EFS file system that the access point provides access to. The access point exposes the specified file system path as the root directory of your file system to applications using the access point. NFS clients using the access point can only access data in the access point's ``RootDirectory`` and its subdirectories.
        """
        creation_info: NotRequired[pulumi.Input['AccessPointCreationInfoArgsDict']]
        """
        (Optional) Specifies the POSIX IDs and permissions to apply to the access point's ``RootDirectory``. If the ``RootDirectory`` > ``Path`` specified does not exist, EFS creates the root directory using the ``CreationInfo`` settings when a client connects to an access point. When specifying the ``CreationInfo``, you must provide values for all properties. 
          If you do not provide ``CreationInfo`` and the specified ``RootDirectory`` > ``Path`` does not exist, attempts to mount the file system using the access point will fail.
        """
        path: NotRequired[pulumi.Input[builtins.str]]
        """
        Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the ``CreationInfo``.
        """
elif False:
    AccessPointRootDirectoryArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AccessPointRootDirectoryArgs:
    def __init__(__self__, *,
                 creation_info: Optional[pulumi.Input['AccessPointCreationInfoArgs']] = None,
                 path: Optional[pulumi.Input[builtins.str]] = None):
        """
        Specifies the directory on the Amazon EFS file system that the access point provides access to. The access point exposes the specified file system path as the root directory of your file system to applications using the access point. NFS clients using the access point can only access data in the access point's ``RootDirectory`` and its subdirectories.
        :param pulumi.Input['AccessPointCreationInfoArgs'] creation_info: (Optional) Specifies the POSIX IDs and permissions to apply to the access point's ``RootDirectory``. If the ``RootDirectory`` > ``Path`` specified does not exist, EFS creates the root directory using the ``CreationInfo`` settings when a client connects to an access point. When specifying the ``CreationInfo``, you must provide values for all properties. 
                 If you do not provide ``CreationInfo`` and the specified ``RootDirectory`` > ``Path`` does not exist, attempts to mount the file system using the access point will fail.
        :param pulumi.Input[builtins.str] path: Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the ``CreationInfo``.
        """
        if creation_info is not None:
            pulumi.set(__self__, "creation_info", creation_info)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter(name="creationInfo")
    def creation_info(self) -> Optional[pulumi.Input['AccessPointCreationInfoArgs']]:
        """
        (Optional) Specifies the POSIX IDs and permissions to apply to the access point's ``RootDirectory``. If the ``RootDirectory`` > ``Path`` specified does not exist, EFS creates the root directory using the ``CreationInfo`` settings when a client connects to an access point. When specifying the ``CreationInfo``, you must provide values for all properties. 
          If you do not provide ``CreationInfo`` and the specified ``RootDirectory`` > ``Path`` does not exist, attempts to mount the file system using the access point will fail.
        """
        return pulumi.get(self, "creation_info")

    @creation_info.setter
    def creation_info(self, value: Optional[pulumi.Input['AccessPointCreationInfoArgs']]):
        pulumi.set(self, "creation_info", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the ``CreationInfo``.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "path", value)


if not MYPY:
    class FileSystemBackupPolicyArgsDict(TypedDict):
        """
        The backup policy turns automatic backups for the file system on or off.
        """
        status: pulumi.Input['FileSystemBackupPolicyStatus']
        """
        Set the backup policy status for the file system.
          +  *ENABLED* - Turns automatic backups on for the file system. 
          +  *DISABLED* - Turns automatic backups off for the file system.
        """
elif False:
    FileSystemBackupPolicyArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FileSystemBackupPolicyArgs:
    def __init__(__self__, *,
                 status: pulumi.Input['FileSystemBackupPolicyStatus']):
        """
        The backup policy turns automatic backups for the file system on or off.
        :param pulumi.Input['FileSystemBackupPolicyStatus'] status: Set the backup policy status for the file system.
                 +  *ENABLED* - Turns automatic backups on for the file system. 
                 +  *DISABLED* - Turns automatic backups off for the file system.
        """
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input['FileSystemBackupPolicyStatus']:
        """
        Set the backup policy status for the file system.
          +  *ENABLED* - Turns automatic backups on for the file system. 
          +  *DISABLED* - Turns automatic backups off for the file system.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input['FileSystemBackupPolicyStatus']):
        pulumi.set(self, "status", value)


if not MYPY:
    class FileSystemLifecyclePolicyArgsDict(TypedDict):
        """
        Describes a policy used by Lifecycle management that specifies when to transition files into and out of the EFS storage classes. For more information, see [Managing file system storage](https://docs.aws.amazon.com/efs/latest/ug/lifecycle-management-efs.html).
           +  Each ``LifecyclePolicy`` object can have only a single transition. This means that in a request body, ``LifecyclePolicies`` must be structured as an array of ``LifecyclePolicy`` objects, one object for each transition, ``TransitionToIA``, ``TransitionToArchive``, ``TransitionToPrimaryStorageClass``.
          +  See the AWS::EFS::FileSystem examples for the correct ``LifecyclePolicy`` structure. Do not use the syntax shown on this page.
        """
        transition_to_archive: NotRequired[pulumi.Input[builtins.str]]
        """
        The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.
        """
        transition_to_ia: NotRequired[pulumi.Input[builtins.str]]
        """
        The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Infrequent Access (IA) storage. Metadata operations such as listing the contents of a directory don't count as file access events.
        """
        transition_to_primary_storage_class: NotRequired[pulumi.Input[builtins.str]]
        """
        Whether to move files back to primary (Standard) storage after they are accessed in IA or Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.
        """
elif False:
    FileSystemLifecyclePolicyArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FileSystemLifecyclePolicyArgs:
    def __init__(__self__, *,
                 transition_to_archive: Optional[pulumi.Input[builtins.str]] = None,
                 transition_to_ia: Optional[pulumi.Input[builtins.str]] = None,
                 transition_to_primary_storage_class: Optional[pulumi.Input[builtins.str]] = None):
        """
        Describes a policy used by Lifecycle management that specifies when to transition files into and out of the EFS storage classes. For more information, see [Managing file system storage](https://docs.aws.amazon.com/efs/latest/ug/lifecycle-management-efs.html).
           +  Each ``LifecyclePolicy`` object can have only a single transition. This means that in a request body, ``LifecyclePolicies`` must be structured as an array of ``LifecyclePolicy`` objects, one object for each transition, ``TransitionToIA``, ``TransitionToArchive``, ``TransitionToPrimaryStorageClass``.
          +  See the AWS::EFS::FileSystem examples for the correct ``LifecyclePolicy`` structure. Do not use the syntax shown on this page.
        :param pulumi.Input[builtins.str] transition_to_archive: The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.
        :param pulumi.Input[builtins.str] transition_to_ia: The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Infrequent Access (IA) storage. Metadata operations such as listing the contents of a directory don't count as file access events.
        :param pulumi.Input[builtins.str] transition_to_primary_storage_class: Whether to move files back to primary (Standard) storage after they are accessed in IA or Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.
        """
        if transition_to_archive is not None:
            pulumi.set(__self__, "transition_to_archive", transition_to_archive)
        if transition_to_ia is not None:
            pulumi.set(__self__, "transition_to_ia", transition_to_ia)
        if transition_to_primary_storage_class is not None:
            pulumi.set(__self__, "transition_to_primary_storage_class", transition_to_primary_storage_class)

    @property
    @pulumi.getter(name="transitionToArchive")
    def transition_to_archive(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.
        """
        return pulumi.get(self, "transition_to_archive")

    @transition_to_archive.setter
    def transition_to_archive(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "transition_to_archive", value)

    @property
    @pulumi.getter(name="transitionToIa")
    def transition_to_ia(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Infrequent Access (IA) storage. Metadata operations such as listing the contents of a directory don't count as file access events.
        """
        return pulumi.get(self, "transition_to_ia")

    @transition_to_ia.setter
    def transition_to_ia(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "transition_to_ia", value)

    @property
    @pulumi.getter(name="transitionToPrimaryStorageClass")
    def transition_to_primary_storage_class(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Whether to move files back to primary (Standard) storage after they are accessed in IA or Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.
        """
        return pulumi.get(self, "transition_to_primary_storage_class")

    @transition_to_primary_storage_class.setter
    def transition_to_primary_storage_class(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "transition_to_primary_storage_class", value)


if not MYPY:
    class FileSystemProtectionArgsDict(TypedDict):
        """
        Describes the protection on the file system.
        """
        replication_overwrite_protection: NotRequired[pulumi.Input['FileSystemProtectionReplicationOverwriteProtection']]
        """
        The status of the file system's replication overwrite protection.
          +  ``ENABLED`` – The file system cannot be used as the destination file system in a replication configuration. The file system is writeable. Replication overwrite protection is ``ENABLED`` by default. 
          +  ``DISABLED`` – The file system can be used as the destination file system in a replication configuration. The file system is read-only and can only be modified by EFS replication.
          +  ``REPLICATING`` – The file system is being used as the destination file system in a replication configuration. The file system is read-only and is modified only by EFS replication.
          
         If the replication configuration is deleted, the file system's replication overwrite protection is re-enabled, the file system becomes writeable.
        """
elif False:
    FileSystemProtectionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FileSystemProtectionArgs:
    def __init__(__self__, *,
                 replication_overwrite_protection: Optional[pulumi.Input['FileSystemProtectionReplicationOverwriteProtection']] = None):
        """
        Describes the protection on the file system.
        :param pulumi.Input['FileSystemProtectionReplicationOverwriteProtection'] replication_overwrite_protection: The status of the file system's replication overwrite protection.
                 +  ``ENABLED`` – The file system cannot be used as the destination file system in a replication configuration. The file system is writeable. Replication overwrite protection is ``ENABLED`` by default. 
                 +  ``DISABLED`` – The file system can be used as the destination file system in a replication configuration. The file system is read-only and can only be modified by EFS replication.
                 +  ``REPLICATING`` – The file system is being used as the destination file system in a replication configuration. The file system is read-only and is modified only by EFS replication.
                 
                If the replication configuration is deleted, the file system's replication overwrite protection is re-enabled, the file system becomes writeable.
        """
        if replication_overwrite_protection is not None:
            pulumi.set(__self__, "replication_overwrite_protection", replication_overwrite_protection)

    @property
    @pulumi.getter(name="replicationOverwriteProtection")
    def replication_overwrite_protection(self) -> Optional[pulumi.Input['FileSystemProtectionReplicationOverwriteProtection']]:
        """
        The status of the file system's replication overwrite protection.
          +  ``ENABLED`` – The file system cannot be used as the destination file system in a replication configuration. The file system is writeable. Replication overwrite protection is ``ENABLED`` by default. 
          +  ``DISABLED`` – The file system can be used as the destination file system in a replication configuration. The file system is read-only and can only be modified by EFS replication.
          +  ``REPLICATING`` – The file system is being used as the destination file system in a replication configuration. The file system is read-only and is modified only by EFS replication.
          
         If the replication configuration is deleted, the file system's replication overwrite protection is re-enabled, the file system becomes writeable.
        """
        return pulumi.get(self, "replication_overwrite_protection")

    @replication_overwrite_protection.setter
    def replication_overwrite_protection(self, value: Optional[pulumi.Input['FileSystemProtectionReplicationOverwriteProtection']]):
        pulumi.set(self, "replication_overwrite_protection", value)


if not MYPY:
    class FileSystemReplicationConfigurationArgsDict(TypedDict):
        """
        Describes the replication configuration for a specific file system.
        """
        destinations: NotRequired[pulumi.Input[Sequence[pulumi.Input['FileSystemReplicationDestinationArgsDict']]]]
        """
        An array of destination objects. Only one destination object is supported.
        """
elif False:
    FileSystemReplicationConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FileSystemReplicationConfigurationArgs:
    def __init__(__self__, *,
                 destinations: Optional[pulumi.Input[Sequence[pulumi.Input['FileSystemReplicationDestinationArgs']]]] = None):
        """
        Describes the replication configuration for a specific file system.
        :param pulumi.Input[Sequence[pulumi.Input['FileSystemReplicationDestinationArgs']]] destinations: An array of destination objects. Only one destination object is supported.
        """
        if destinations is not None:
            pulumi.set(__self__, "destinations", destinations)

    @property
    @pulumi.getter
    def destinations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FileSystemReplicationDestinationArgs']]]]:
        """
        An array of destination objects. Only one destination object is supported.
        """
        return pulumi.get(self, "destinations")

    @destinations.setter
    def destinations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FileSystemReplicationDestinationArgs']]]]):
        pulumi.set(self, "destinations", value)


if not MYPY:
    class FileSystemReplicationDestinationArgsDict(TypedDict):
        """
        Describes the destination file system in the replication configuration.
        """
        availability_zone_name: NotRequired[pulumi.Input[builtins.str]]
        """
        For One Zone file systems, the replication configuration must specify the Availability Zone in which the destination file system is located. 
         Use the format ``us-east-1a`` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the *Amazon EFS User Guide*.
          One Zone file system type is not available in all Availability Zones in AWS-Regions where Amazon EFS is available.
        """
        file_system_id: NotRequired[pulumi.Input[builtins.str]]
        """
        The ID of the destination Amazon EFS file system.
        """
        kms_key_id: NotRequired[pulumi.Input[builtins.str]]
        """
        The ID of an kms-key-long used to protect the encrypted file system.
        """
        region: NotRequired[pulumi.Input[builtins.str]]
        """
        The AWS-Region in which the destination file system is located.
          For One Zone file systems, the replication configuration must specify the AWS-Region in which the destination file system is located.
        """
        role_arn: NotRequired[pulumi.Input[builtins.str]]
        """
        The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
        """
        status: NotRequired[pulumi.Input[builtins.str]]
        """
        Describes the status of the replication configuration. For more information about replication status, see [Viewing replication details](https://docs.aws.amazon.com//efs/latest/ug/awsbackup.html#restoring-backup-efsmonitoring-replication-status.html) in the *Amazon EFS User Guide*.
        """
        status_message: NotRequired[pulumi.Input[builtins.str]]
        """
        Message that provides details about the ``PAUSED`` or ``ERRROR`` state of the replication destination configuration. For more information about replication status messages, see [Viewing replication details](https://docs.aws.amazon.com//efs/latest/ug/awsbackup.html#restoring-backup-efsmonitoring-replication-status.html) in the *Amazon EFS User Guide*.
        """
elif False:
    FileSystemReplicationDestinationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FileSystemReplicationDestinationArgs:
    def __init__(__self__, *,
                 availability_zone_name: Optional[pulumi.Input[builtins.str]] = None,
                 file_system_id: Optional[pulumi.Input[builtins.str]] = None,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 role_arn: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 status_message: Optional[pulumi.Input[builtins.str]] = None):
        """
        Describes the destination file system in the replication configuration.
        :param pulumi.Input[builtins.str] availability_zone_name: For One Zone file systems, the replication configuration must specify the Availability Zone in which the destination file system is located. 
                Use the format ``us-east-1a`` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the *Amazon EFS User Guide*.
                 One Zone file system type is not available in all Availability Zones in AWS-Regions where Amazon EFS is available.
        :param pulumi.Input[builtins.str] file_system_id: The ID of the destination Amazon EFS file system.
        :param pulumi.Input[builtins.str] kms_key_id: The ID of an kms-key-long used to protect the encrypted file system.
        :param pulumi.Input[builtins.str] region: The AWS-Region in which the destination file system is located.
                 For One Zone file systems, the replication configuration must specify the AWS-Region in which the destination file system is located.
        :param pulumi.Input[builtins.str] role_arn: The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
        :param pulumi.Input[builtins.str] status: Describes the status of the replication configuration. For more information about replication status, see [Viewing replication details](https://docs.aws.amazon.com//efs/latest/ug/awsbackup.html#restoring-backup-efsmonitoring-replication-status.html) in the *Amazon EFS User Guide*.
        :param pulumi.Input[builtins.str] status_message: Message that provides details about the ``PAUSED`` or ``ERRROR`` state of the replication destination configuration. For more information about replication status messages, see [Viewing replication details](https://docs.aws.amazon.com//efs/latest/ug/awsbackup.html#restoring-backup-efsmonitoring-replication-status.html) in the *Amazon EFS User Guide*.
        """
        if availability_zone_name is not None:
            pulumi.set(__self__, "availability_zone_name", availability_zone_name)
        if file_system_id is not None:
            pulumi.set(__self__, "file_system_id", file_system_id)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if status_message is not None:
            pulumi.set(__self__, "status_message", status_message)

    @property
    @pulumi.getter(name="availabilityZoneName")
    def availability_zone_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        For One Zone file systems, the replication configuration must specify the Availability Zone in which the destination file system is located. 
         Use the format ``us-east-1a`` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the *Amazon EFS User Guide*.
          One Zone file system type is not available in all Availability Zones in AWS-Regions where Amazon EFS is available.
        """
        return pulumi.get(self, "availability_zone_name")

    @availability_zone_name.setter
    def availability_zone_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "availability_zone_name", value)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the destination Amazon EFS file system.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of an kms-key-long used to protect the encrypted file system.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The AWS-Region in which the destination file system is located.
          For One Zone file systems, the replication configuration must specify the AWS-Region in which the destination file system is located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Describes the status of the replication configuration. For more information about replication status, see [Viewing replication details](https://docs.aws.amazon.com//efs/latest/ug/awsbackup.html#restoring-backup-efsmonitoring-replication-status.html) in the *Amazon EFS User Guide*.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Message that provides details about the ``PAUSED`` or ``ERRROR`` state of the replication destination configuration. For more information about replication status messages, see [Viewing replication details](https://docs.aws.amazon.com//efs/latest/ug/awsbackup.html#restoring-backup-efsmonitoring-replication-status.html) in the *Amazon EFS User Guide*.
        """
        return pulumi.get(self, "status_message")

    @status_message.setter
    def status_message(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status_message", value)


