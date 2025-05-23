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
from ._enums import *

__all__ = [
    'GetMailManagerArchiveResult',
    'AwaitableGetMailManagerArchiveResult',
    'get_mail_manager_archive',
    'get_mail_manager_archive_output',
]

@pulumi.output_type
class GetMailManagerArchiveResult:
    def __init__(__self__, archive_arn=None, archive_id=None, archive_name=None, archive_state=None, retention=None, tags=None):
        if archive_arn and not isinstance(archive_arn, str):
            raise TypeError("Expected argument 'archive_arn' to be a str")
        pulumi.set(__self__, "archive_arn", archive_arn)
        if archive_id and not isinstance(archive_id, str):
            raise TypeError("Expected argument 'archive_id' to be a str")
        pulumi.set(__self__, "archive_id", archive_id)
        if archive_name and not isinstance(archive_name, str):
            raise TypeError("Expected argument 'archive_name' to be a str")
        pulumi.set(__self__, "archive_name", archive_name)
        if archive_state and not isinstance(archive_state, str):
            raise TypeError("Expected argument 'archive_state' to be a str")
        pulumi.set(__self__, "archive_state", archive_state)
        if retention and not isinstance(retention, dict):
            raise TypeError("Expected argument 'retention' to be a dict")
        pulumi.set(__self__, "retention", retention)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="archiveArn")
    def archive_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the archive.
        """
        return pulumi.get(self, "archive_arn")

    @property
    @pulumi.getter(name="archiveId")
    def archive_id(self) -> Optional[builtins.str]:
        """
        The unique identifier of the archive.
        """
        return pulumi.get(self, "archive_id")

    @property
    @pulumi.getter(name="archiveName")
    def archive_name(self) -> Optional[builtins.str]:
        """
        A unique name for the new archive.
        """
        return pulumi.get(self, "archive_name")

    @property
    @pulumi.getter(name="archiveState")
    def archive_state(self) -> Optional['MailManagerArchiveArchiveState']:
        """
        The current state of the archive:

        - `ACTIVE` – The archive is ready and available for use.
        - `PENDING_DELETION` – The archive has been marked for deletion and will be permanently deleted in 30 days. No further modifications can be made in this state.
        """
        return pulumi.get(self, "archive_state")

    @property
    @pulumi.getter
    def retention(self) -> Optional['outputs.MailManagerArchiveArchiveRetentionProperties']:
        """
        The period for retaining emails in the archive before automatic deletion.
        """
        return pulumi.get(self, "retention")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
        """
        return pulumi.get(self, "tags")


class AwaitableGetMailManagerArchiveResult(GetMailManagerArchiveResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMailManagerArchiveResult(
            archive_arn=self.archive_arn,
            archive_id=self.archive_id,
            archive_name=self.archive_name,
            archive_state=self.archive_state,
            retention=self.retention,
            tags=self.tags)


def get_mail_manager_archive(archive_id: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMailManagerArchiveResult:
    """
    Definition of AWS::SES::MailManagerArchive Resource Type


    :param builtins.str archive_id: The unique identifier of the archive.
    """
    __args__ = dict()
    __args__['archiveId'] = archive_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ses:getMailManagerArchive', __args__, opts=opts, typ=GetMailManagerArchiveResult).value

    return AwaitableGetMailManagerArchiveResult(
        archive_arn=pulumi.get(__ret__, 'archive_arn'),
        archive_id=pulumi.get(__ret__, 'archive_id'),
        archive_name=pulumi.get(__ret__, 'archive_name'),
        archive_state=pulumi.get(__ret__, 'archive_state'),
        retention=pulumi.get(__ret__, 'retention'),
        tags=pulumi.get(__ret__, 'tags'))
def get_mail_manager_archive_output(archive_id: Optional[pulumi.Input[builtins.str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMailManagerArchiveResult]:
    """
    Definition of AWS::SES::MailManagerArchive Resource Type


    :param builtins.str archive_id: The unique identifier of the archive.
    """
    __args__ = dict()
    __args__['archiveId'] = archive_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:ses:getMailManagerArchive', __args__, opts=opts, typ=GetMailManagerArchiveResult)
    return __ret__.apply(lambda __response__: GetMailManagerArchiveResult(
        archive_arn=pulumi.get(__response__, 'archive_arn'),
        archive_id=pulumi.get(__response__, 'archive_id'),
        archive_name=pulumi.get(__response__, 'archive_name'),
        archive_state=pulumi.get(__response__, 'archive_state'),
        retention=pulumi.get(__response__, 'retention'),
        tags=pulumi.get(__response__, 'tags')))
