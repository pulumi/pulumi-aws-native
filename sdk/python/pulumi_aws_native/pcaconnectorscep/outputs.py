# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'ConnectorMobileDeviceManagement',
    'ConnectorOpenIdConfiguration',
]

@pulumi.output_type
class ConnectorMobileDeviceManagement(dict):
    def __init__(__self__):
        pass


@pulumi.output_type
class ConnectorOpenIdConfiguration(dict):
    def __init__(__self__, *,
                 audience: Optional[str] = None,
                 issuer: Optional[str] = None,
                 subject: Optional[str] = None):
        if audience is not None:
            pulumi.set(__self__, "audience", audience)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)

    @property
    @pulumi.getter
    def audience(self) -> Optional[str]:
        return pulumi.get(self, "audience")

    @property
    @pulumi.getter
    def issuer(self) -> Optional[str]:
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter
    def subject(self) -> Optional[str]:
        return pulumi.get(self, "subject")

