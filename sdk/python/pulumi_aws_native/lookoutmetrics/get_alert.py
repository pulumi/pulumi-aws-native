# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAlertResult',
    'AwaitableGetAlertResult',
    'get_alert',
    'get_alert_output',
]

@pulumi.output_type
class GetAlertResult:
    def __init__(__self__, arn=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        ARN assigned to the alert.
        """
        return pulumi.get(self, "arn")


class AwaitableGetAlertResult(GetAlertResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlertResult(
            arn=self.arn)


def get_alert(arn: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAlertResult:
    """
    Resource Type definition for AWS::LookoutMetrics::Alert


    :param str arn: ARN assigned to the alert.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:lookoutmetrics:getAlert', __args__, opts=opts, typ=GetAlertResult).value

    return AwaitableGetAlertResult(
        arn=__ret__.arn)


@_utilities.lift_output_func(get_alert)
def get_alert_output(arn: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAlertResult]:
    """
    Resource Type definition for AWS::LookoutMetrics::Alert


    :param str arn: ARN assigned to the alert.
    """
    ...