# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'PreferencesIdleConnectionAlertArgs',
    'PreferencesIdleConnectionPreferencesArgs',
    'PreferencesIdleConnectionTimeoutArgs',
]

@pulumi.input_type
class PreferencesIdleConnectionAlertArgs:
    def __init__(__self__, *,
                 value: pulumi.Input[int],
                 type: Optional[pulumi.Input['PreferencesIdleConnectionAlertType']] = None):
        pulumi.set(__self__, "value", value)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[int]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[int]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['PreferencesIdleConnectionAlertType']]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['PreferencesIdleConnectionAlertType']]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class PreferencesIdleConnectionPreferencesArgs:
    def __init__(__self__, *,
                 alert: Optional[pulumi.Input['PreferencesIdleConnectionAlertArgs']] = None,
                 timeout: Optional[pulumi.Input['PreferencesIdleConnectionTimeoutArgs']] = None):
        """
        Idle Connection Preferences
        """
        if alert is not None:
            pulumi.set(__self__, "alert", alert)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def alert(self) -> Optional[pulumi.Input['PreferencesIdleConnectionAlertArgs']]:
        return pulumi.get(self, "alert")

    @alert.setter
    def alert(self, value: Optional[pulumi.Input['PreferencesIdleConnectionAlertArgs']]):
        pulumi.set(self, "alert", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input['PreferencesIdleConnectionTimeoutArgs']]:
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input['PreferencesIdleConnectionTimeoutArgs']]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class PreferencesIdleConnectionTimeoutArgs:
    def __init__(__self__, *,
                 value: pulumi.Input[int],
                 type: Optional[pulumi.Input['PreferencesIdleConnectionTimeoutType']] = None):
        pulumi.set(__self__, "value", value)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[int]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[int]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['PreferencesIdleConnectionTimeoutType']]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['PreferencesIdleConnectionTimeoutType']]):
        pulumi.set(self, "type", value)

