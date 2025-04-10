# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ._enums import *

__all__ = [
    'LinkConfiguration',
    'LinkFilter',
]

@pulumi.output_type
class LinkConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "logGroupConfiguration":
            suggest = "log_group_configuration"
        elif key == "metricConfiguration":
            suggest = "metric_configuration"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LinkConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LinkConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LinkConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 log_group_configuration: Optional['outputs.LinkFilter'] = None,
                 metric_configuration: Optional['outputs.LinkFilter'] = None):
        """
        :param 'LinkFilter' log_group_configuration: Use this structure to filter which log groups are to send log events from the source account to the monitoring account.
        :param 'LinkFilter' metric_configuration: Use this structure to filter which metric namespaces are to be shared from the source account to the monitoring account.
        """
        if log_group_configuration is not None:
            pulumi.set(__self__, "log_group_configuration", log_group_configuration)
        if metric_configuration is not None:
            pulumi.set(__self__, "metric_configuration", metric_configuration)

    @property
    @pulumi.getter(name="logGroupConfiguration")
    def log_group_configuration(self) -> Optional['outputs.LinkFilter']:
        """
        Use this structure to filter which log groups are to send log events from the source account to the monitoring account.
        """
        return pulumi.get(self, "log_group_configuration")

    @property
    @pulumi.getter(name="metricConfiguration")
    def metric_configuration(self) -> Optional['outputs.LinkFilter']:
        """
        Use this structure to filter which metric namespaces are to be shared from the source account to the monitoring account.
        """
        return pulumi.get(self, "metric_configuration")


@pulumi.output_type
class LinkFilter(dict):
    def __init__(__self__, *,
                 filter: str):
        pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def filter(self) -> str:
        return pulumi.get(self, "filter")


