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
    'GetFleetResult',
    'AwaitableGetFleetResult',
    'get_fleet',
    'get_fleet_output',
]

@pulumi.output_type
class GetFleetResult:
    def __init__(__self__, arn=None, capabilities=None, configuration=None, description=None, display_name=None, fleet_id=None, host_configuration=None, max_worker_count=None, min_worker_count=None, role_arn=None, status=None, tags=None, worker_count=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if capabilities and not isinstance(capabilities, dict):
            raise TypeError("Expected argument 'capabilities' to be a dict")
        pulumi.set(__self__, "capabilities", capabilities)
        if configuration and not isinstance(configuration, dict):
            raise TypeError("Expected argument 'configuration' to be a dict")
        pulumi.set(__self__, "configuration", configuration)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if fleet_id and not isinstance(fleet_id, str):
            raise TypeError("Expected argument 'fleet_id' to be a str")
        pulumi.set(__self__, "fleet_id", fleet_id)
        if host_configuration and not isinstance(host_configuration, dict):
            raise TypeError("Expected argument 'host_configuration' to be a dict")
        pulumi.set(__self__, "host_configuration", host_configuration)
        if max_worker_count and not isinstance(max_worker_count, int):
            raise TypeError("Expected argument 'max_worker_count' to be a int")
        pulumi.set(__self__, "max_worker_count", max_worker_count)
        if min_worker_count and not isinstance(min_worker_count, int):
            raise TypeError("Expected argument 'min_worker_count' to be a int")
        pulumi.set(__self__, "min_worker_count", min_worker_count)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if worker_count and not isinstance(worker_count, int):
            raise TypeError("Expected argument 'worker_count' to be a int")
        pulumi.set(__self__, "worker_count", worker_count)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) assigned to the fleet.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def capabilities(self) -> Optional['outputs.FleetCapabilities']:
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter
    def configuration(self) -> Optional[Any]:
        """
        The configuration details for the fleet.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        A description that helps identify what the fleet is used for.

        > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[builtins.str]:
        """
        The display name of the fleet summary to update.

        > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="fleetId")
    def fleet_id(self) -> Optional[builtins.str]:
        """
        The fleet ID.
        """
        return pulumi.get(self, "fleet_id")

    @property
    @pulumi.getter(name="hostConfiguration")
    def host_configuration(self) -> Optional['outputs.FleetHostConfiguration']:
        """
        Provides a script that runs as a worker is starting up that you can use to provide additional configuration for workers in your fleet.

        To remove a script from a fleet, use the [UpdateFleet](https://docs.aws.amazon.com/deadline-cloud/latest/APIReference/API_UpdateFleet.html) operation with the `hostConfiguration` `scriptBody` parameter set to an empty string ("").
        """
        return pulumi.get(self, "host_configuration")

    @property
    @pulumi.getter(name="maxWorkerCount")
    def max_worker_count(self) -> Optional[builtins.int]:
        """
        The maximum number of workers specified in the fleet.
        """
        return pulumi.get(self, "max_worker_count")

    @property
    @pulumi.getter(name="minWorkerCount")
    def min_worker_count(self) -> Optional[builtins.int]:
        """
        The minimum number of workers in the fleet.
        """
        return pulumi.get(self, "min_worker_count")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[builtins.str]:
        """
        The IAM role that workers in the fleet use when processing jobs.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def status(self) -> Optional['FleetStatus']:
        """
        The status of the fleet.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="workerCount")
    def worker_count(self) -> Optional[builtins.int]:
        """
        The number of workers in the fleet summary.
        """
        return pulumi.get(self, "worker_count")


class AwaitableGetFleetResult(GetFleetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFleetResult(
            arn=self.arn,
            capabilities=self.capabilities,
            configuration=self.configuration,
            description=self.description,
            display_name=self.display_name,
            fleet_id=self.fleet_id,
            host_configuration=self.host_configuration,
            max_worker_count=self.max_worker_count,
            min_worker_count=self.min_worker_count,
            role_arn=self.role_arn,
            status=self.status,
            tags=self.tags,
            worker_count=self.worker_count)


def get_fleet(arn: Optional[builtins.str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFleetResult:
    """
    Definition of AWS::Deadline::Fleet Resource Type


    :param builtins.str arn: The Amazon Resource Name (ARN) assigned to the fleet.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:deadline:getFleet', __args__, opts=opts, typ=GetFleetResult).value

    return AwaitableGetFleetResult(
        arn=pulumi.get(__ret__, 'arn'),
        capabilities=pulumi.get(__ret__, 'capabilities'),
        configuration=pulumi.get(__ret__, 'configuration'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        fleet_id=pulumi.get(__ret__, 'fleet_id'),
        host_configuration=pulumi.get(__ret__, 'host_configuration'),
        max_worker_count=pulumi.get(__ret__, 'max_worker_count'),
        min_worker_count=pulumi.get(__ret__, 'min_worker_count'),
        role_arn=pulumi.get(__ret__, 'role_arn'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        worker_count=pulumi.get(__ret__, 'worker_count'))
def get_fleet_output(arn: Optional[pulumi.Input[builtins.str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFleetResult]:
    """
    Definition of AWS::Deadline::Fleet Resource Type


    :param builtins.str arn: The Amazon Resource Name (ARN) assigned to the fleet.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:deadline:getFleet', __args__, opts=opts, typ=GetFleetResult)
    return __ret__.apply(lambda __response__: GetFleetResult(
        arn=pulumi.get(__response__, 'arn'),
        capabilities=pulumi.get(__response__, 'capabilities'),
        configuration=pulumi.get(__response__, 'configuration'),
        description=pulumi.get(__response__, 'description'),
        display_name=pulumi.get(__response__, 'display_name'),
        fleet_id=pulumi.get(__response__, 'fleet_id'),
        host_configuration=pulumi.get(__response__, 'host_configuration'),
        max_worker_count=pulumi.get(__response__, 'max_worker_count'),
        min_worker_count=pulumi.get(__response__, 'min_worker_count'),
        role_arn=pulumi.get(__response__, 'role_arn'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags'),
        worker_count=pulumi.get(__response__, 'worker_count')))
