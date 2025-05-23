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
    'ApplicationRuntimeEnvironment',
    'StreamGroupDefaultApplication',
    'StreamGroupLocationConfiguration',
]

@pulumi.output_type
class ApplicationRuntimeEnvironment(dict):
    def __init__(__self__, *,
                 type: builtins.str,
                 version: builtins.str):
        """
        :param builtins.str type: The operating system and other drivers. For Proton, this also includes the Proton compatibility layer.
        :param builtins.str version: Versioned container environment for the application operating system.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The operating system and other drivers. For Proton, this also includes the Proton compatibility layer.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> builtins.str:
        """
        Versioned container environment for the application operating system.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class StreamGroupDefaultApplication(dict):
    def __init__(__self__, *,
                 arn: Optional[builtins.str] = None,
                 id: Optional[builtins.str] = None):
        """
        :param builtins.str arn: An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the application resource. Example ARN: `arn:aws:gameliftstreams:us-west-2:111122223333:application/a-9ZY8X7Wv6` .
        :param builtins.str id: An ID that uniquely identifies the application resource. Example ID: `a-9ZY8X7Wv6` .
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the application resource. Example ARN: `arn:aws:gameliftstreams:us-west-2:111122223333:application/a-9ZY8X7Wv6` .
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        An ID that uniquely identifies the application resource. Example ID: `a-9ZY8X7Wv6` .
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class StreamGroupLocationConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "locationName":
            suggest = "location_name"
        elif key == "alwaysOnCapacity":
            suggest = "always_on_capacity"
        elif key == "onDemandCapacity":
            suggest = "on_demand_capacity"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StreamGroupLocationConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StreamGroupLocationConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StreamGroupLocationConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 location_name: builtins.str,
                 always_on_capacity: Optional[builtins.int] = None,
                 on_demand_capacity: Optional[builtins.int] = None):
        """
        :param builtins.str location_name: A location's name. For example, `us-east-1` . For a complete list of locations that Amazon GameLift Streams supports, refer to [Regions, quotas, and limitations](https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/regions-quotas.html) in the *Amazon GameLift Streams Developer Guide* .
        :param builtins.int always_on_capacity: The streaming capacity that is allocated and ready to handle stream requests without delay. You pay for this capacity whether it's in use or not. Best for quickest time from streaming request to streaming session.
        :param builtins.int on_demand_capacity: The streaming capacity that Amazon GameLift Streams can allocate in response to stream requests, and then de-allocate when the session has terminated. This offers a cost control measure at the expense of a greater startup time (typically under 5 minutes).
        """
        pulumi.set(__self__, "location_name", location_name)
        if always_on_capacity is not None:
            pulumi.set(__self__, "always_on_capacity", always_on_capacity)
        if on_demand_capacity is not None:
            pulumi.set(__self__, "on_demand_capacity", on_demand_capacity)

    @property
    @pulumi.getter(name="locationName")
    def location_name(self) -> builtins.str:
        """
        A location's name. For example, `us-east-1` . For a complete list of locations that Amazon GameLift Streams supports, refer to [Regions, quotas, and limitations](https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/regions-quotas.html) in the *Amazon GameLift Streams Developer Guide* .
        """
        return pulumi.get(self, "location_name")

    @property
    @pulumi.getter(name="alwaysOnCapacity")
    def always_on_capacity(self) -> Optional[builtins.int]:
        """
        The streaming capacity that is allocated and ready to handle stream requests without delay. You pay for this capacity whether it's in use or not. Best for quickest time from streaming request to streaming session.
        """
        return pulumi.get(self, "always_on_capacity")

    @property
    @pulumi.getter(name="onDemandCapacity")
    def on_demand_capacity(self) -> Optional[builtins.int]:
        """
        The streaming capacity that Amazon GameLift Streams can allocate in response to stream requests, and then de-allocate when the session has terminated. This offers a cost control measure at the expense of a greater startup time (typically under 5 minutes).
        """
        return pulumi.get(self, "on_demand_capacity")


