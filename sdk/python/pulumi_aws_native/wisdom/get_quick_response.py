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
    'GetQuickResponseResult',
    'AwaitableGetQuickResponseResult',
    'get_quick_response',
    'get_quick_response_output',
]

@pulumi.output_type
class GetQuickResponseResult:
    def __init__(__self__, channels=None, content=None, content_type=None, contents=None, description=None, grouping_configuration=None, is_active=None, language=None, name=None, quick_response_arn=None, quick_response_id=None, shortcut_key=None, status=None, tags=None):
        if channels and not isinstance(channels, list):
            raise TypeError("Expected argument 'channels' to be a list")
        pulumi.set(__self__, "channels", channels)
        if content and not isinstance(content, dict):
            raise TypeError("Expected argument 'content' to be a dict")
        pulumi.set(__self__, "content", content)
        if content_type and not isinstance(content_type, str):
            raise TypeError("Expected argument 'content_type' to be a str")
        pulumi.set(__self__, "content_type", content_type)
        if contents and not isinstance(contents, dict):
            raise TypeError("Expected argument 'contents' to be a dict")
        pulumi.set(__self__, "contents", contents)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if grouping_configuration and not isinstance(grouping_configuration, dict):
            raise TypeError("Expected argument 'grouping_configuration' to be a dict")
        pulumi.set(__self__, "grouping_configuration", grouping_configuration)
        if is_active and not isinstance(is_active, bool):
            raise TypeError("Expected argument 'is_active' to be a bool")
        pulumi.set(__self__, "is_active", is_active)
        if language and not isinstance(language, str):
            raise TypeError("Expected argument 'language' to be a str")
        pulumi.set(__self__, "language", language)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if quick_response_arn and not isinstance(quick_response_arn, str):
            raise TypeError("Expected argument 'quick_response_arn' to be a str")
        pulumi.set(__self__, "quick_response_arn", quick_response_arn)
        if quick_response_id and not isinstance(quick_response_id, str):
            raise TypeError("Expected argument 'quick_response_id' to be a str")
        pulumi.set(__self__, "quick_response_id", quick_response_id)
        if shortcut_key and not isinstance(shortcut_key, str):
            raise TypeError("Expected argument 'shortcut_key' to be a str")
        pulumi.set(__self__, "shortcut_key", shortcut_key)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def channels(self) -> Optional[Sequence['QuickResponseChannelType']]:
        """
        The Amazon Connect contact channels this quick response applies to.
        """
        return pulumi.get(self, "channels")

    @property
    @pulumi.getter
    def content(self) -> Optional['outputs.QuickResponseContentProvider']:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[builtins.str]:
        """
        The media type of the quick response content.
        - Use application/x.quickresponse;format=plain for quick response written in plain text.
        - Use application/x.quickresponse;format=markdown for quick response written in richtext.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def contents(self) -> Optional['outputs.QuickResponseContents']:
        return pulumi.get(self, "contents")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        The description of the quick response.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="groupingConfiguration")
    def grouping_configuration(self) -> Optional['outputs.QuickResponseGroupingConfiguration']:
        return pulumi.get(self, "grouping_configuration")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[builtins.bool]:
        """
        Whether the quick response is active.
        """
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter
    def language(self) -> Optional[builtins.str]:
        """
        The language code value for the language in which the quick response is written. The supported language codes include de_DE, en_US, es_ES, fr_FR, id_ID, it_IT, ja_JP, ko_KR, pt_BR, zh_CN, zh_TW
        """
        return pulumi.get(self, "language")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The name of the quick response.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="quickResponseArn")
    def quick_response_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the quick response.
        """
        return pulumi.get(self, "quick_response_arn")

    @property
    @pulumi.getter(name="quickResponseId")
    def quick_response_id(self) -> Optional[builtins.str]:
        """
        The identifier of the quick response.
        """
        return pulumi.get(self, "quick_response_id")

    @property
    @pulumi.getter(name="shortcutKey")
    def shortcut_key(self) -> Optional[builtins.str]:
        """
        The shortcut key of the quick response. The value should be unique across the knowledge base.
        """
        return pulumi.get(self, "shortcut_key")

    @property
    @pulumi.getter
    def status(self) -> Optional['QuickResponseStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetQuickResponseResult(GetQuickResponseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQuickResponseResult(
            channels=self.channels,
            content=self.content,
            content_type=self.content_type,
            contents=self.contents,
            description=self.description,
            grouping_configuration=self.grouping_configuration,
            is_active=self.is_active,
            language=self.language,
            name=self.name,
            quick_response_arn=self.quick_response_arn,
            quick_response_id=self.quick_response_id,
            shortcut_key=self.shortcut_key,
            status=self.status,
            tags=self.tags)


def get_quick_response(quick_response_arn: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQuickResponseResult:
    """
    Definition of AWS::Wisdom::QuickResponse Resource Type.


    :param builtins.str quick_response_arn: The Amazon Resource Name (ARN) of the quick response.
    """
    __args__ = dict()
    __args__['quickResponseArn'] = quick_response_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:wisdom:getQuickResponse', __args__, opts=opts, typ=GetQuickResponseResult).value

    return AwaitableGetQuickResponseResult(
        channels=pulumi.get(__ret__, 'channels'),
        content=pulumi.get(__ret__, 'content'),
        content_type=pulumi.get(__ret__, 'content_type'),
        contents=pulumi.get(__ret__, 'contents'),
        description=pulumi.get(__ret__, 'description'),
        grouping_configuration=pulumi.get(__ret__, 'grouping_configuration'),
        is_active=pulumi.get(__ret__, 'is_active'),
        language=pulumi.get(__ret__, 'language'),
        name=pulumi.get(__ret__, 'name'),
        quick_response_arn=pulumi.get(__ret__, 'quick_response_arn'),
        quick_response_id=pulumi.get(__ret__, 'quick_response_id'),
        shortcut_key=pulumi.get(__ret__, 'shortcut_key'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'))
def get_quick_response_output(quick_response_arn: Optional[pulumi.Input[builtins.str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetQuickResponseResult]:
    """
    Definition of AWS::Wisdom::QuickResponse Resource Type.


    :param builtins.str quick_response_arn: The Amazon Resource Name (ARN) of the quick response.
    """
    __args__ = dict()
    __args__['quickResponseArn'] = quick_response_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:wisdom:getQuickResponse', __args__, opts=opts, typ=GetQuickResponseResult)
    return __ret__.apply(lambda __response__: GetQuickResponseResult(
        channels=pulumi.get(__response__, 'channels'),
        content=pulumi.get(__response__, 'content'),
        content_type=pulumi.get(__response__, 'content_type'),
        contents=pulumi.get(__response__, 'contents'),
        description=pulumi.get(__response__, 'description'),
        grouping_configuration=pulumi.get(__response__, 'grouping_configuration'),
        is_active=pulumi.get(__response__, 'is_active'),
        language=pulumi.get(__response__, 'language'),
        name=pulumi.get(__response__, 'name'),
        quick_response_arn=pulumi.get(__response__, 'quick_response_arn'),
        quick_response_id=pulumi.get(__response__, 'quick_response_id'),
        shortcut_key=pulumi.get(__response__, 'shortcut_key'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags')))
