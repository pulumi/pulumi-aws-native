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
    'GetChannelResult',
    'AwaitableGetChannelResult',
    'get_channel',
    'get_channel_output',
]

@pulumi.output_type
class GetChannelResult:
    def __init__(__self__, arn=None, filler_slate=None, log_configuration=None, playback_mode=None, tags=None, time_shift_configuration=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if filler_slate and not isinstance(filler_slate, dict):
            raise TypeError("Expected argument 'filler_slate' to be a dict")
        pulumi.set(__self__, "filler_slate", filler_slate)
        if log_configuration and not isinstance(log_configuration, dict):
            raise TypeError("Expected argument 'log_configuration' to be a dict")
        pulumi.set(__self__, "log_configuration", log_configuration)
        if playback_mode and not isinstance(playback_mode, str):
            raise TypeError("Expected argument 'playback_mode' to be a str")
        pulumi.set(__self__, "playback_mode", playback_mode)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if time_shift_configuration and not isinstance(time_shift_configuration, dict):
            raise TypeError("Expected argument 'time_shift_configuration' to be a dict")
        pulumi.set(__self__, "time_shift_configuration", time_shift_configuration)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        <p>The ARN of the channel.</p>
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="fillerSlate")
    def filler_slate(self) -> Optional['outputs.ChannelSlateSource']:
        return pulumi.get(self, "filler_slate")

    @property
    @pulumi.getter(name="logConfiguration")
    def log_configuration(self) -> Optional['outputs.ChannelLogConfigurationForChannel']:
        return pulumi.get(self, "log_configuration")

    @property
    @pulumi.getter(name="playbackMode")
    def playback_mode(self) -> Optional['ChannelPlaybackMode']:
        return pulumi.get(self, "playback_mode")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ChannelTag']]:
        """
        The tags to assign to the channel.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeShiftConfiguration")
    def time_shift_configuration(self) -> Optional['outputs.ChannelTimeShiftConfiguration']:
        return pulumi.get(self, "time_shift_configuration")


class AwaitableGetChannelResult(GetChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetChannelResult(
            arn=self.arn,
            filler_slate=self.filler_slate,
            log_configuration=self.log_configuration,
            playback_mode=self.playback_mode,
            tags=self.tags,
            time_shift_configuration=self.time_shift_configuration)


def get_channel(channel_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetChannelResult:
    """
    Definition of AWS::MediaTailor::Channel Resource Type
    """
    __args__ = dict()
    __args__['channelName'] = channel_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:mediatailor:getChannel', __args__, opts=opts, typ=GetChannelResult).value

    return AwaitableGetChannelResult(
        arn=pulumi.get(__ret__, 'arn'),
        filler_slate=pulumi.get(__ret__, 'filler_slate'),
        log_configuration=pulumi.get(__ret__, 'log_configuration'),
        playback_mode=pulumi.get(__ret__, 'playback_mode'),
        tags=pulumi.get(__ret__, 'tags'),
        time_shift_configuration=pulumi.get(__ret__, 'time_shift_configuration'))


@_utilities.lift_output_func(get_channel)
def get_channel_output(channel_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetChannelResult]:
    """
    Definition of AWS::MediaTailor::Channel Resource Type
    """
    ...