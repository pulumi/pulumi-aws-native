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

__all__ = [
    'ExecutionPlanCapacityUnitsConfigurationArgs',
    'ExecutionPlanCapacityUnitsConfigurationArgsDict',
]

MYPY = False

if not MYPY:
    class ExecutionPlanCapacityUnitsConfigurationArgsDict(TypedDict):
        rescore_capacity_units: pulumi.Input[builtins.int]
        """
        The amount of extra capacity for your rescore execution plan.

        A single extra capacity unit for a rescore execution plan provides 0.01 rescore requests per second. You can add up to 1000 extra capacity units.
        """
elif False:
    ExecutionPlanCapacityUnitsConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ExecutionPlanCapacityUnitsConfigurationArgs:
    def __init__(__self__, *,
                 rescore_capacity_units: pulumi.Input[builtins.int]):
        """
        :param pulumi.Input[builtins.int] rescore_capacity_units: The amount of extra capacity for your rescore execution plan.
               
               A single extra capacity unit for a rescore execution plan provides 0.01 rescore requests per second. You can add up to 1000 extra capacity units.
        """
        pulumi.set(__self__, "rescore_capacity_units", rescore_capacity_units)

    @property
    @pulumi.getter(name="rescoreCapacityUnits")
    def rescore_capacity_units(self) -> pulumi.Input[builtins.int]:
        """
        The amount of extra capacity for your rescore execution plan.

        A single extra capacity unit for a rescore execution plan provides 0.01 rescore requests per second. You can add up to 1000 extra capacity units.
        """
        return pulumi.get(self, "rescore_capacity_units")

    @rescore_capacity_units.setter
    def rescore_capacity_units(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "rescore_capacity_units", value)


