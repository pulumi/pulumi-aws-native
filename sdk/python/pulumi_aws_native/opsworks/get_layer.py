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

__all__ = [
    'GetLayerResult',
    'AwaitableGetLayerResult',
    'get_layer',
    'get_layer_output',
]

@pulumi.output_type
class GetLayerResult:
    def __init__(__self__, attributes=None, auto_assign_elastic_ips=None, auto_assign_public_ips=None, custom_instance_profile_arn=None, custom_json=None, custom_recipes=None, custom_security_group_ids=None, enable_auto_healing=None, id=None, install_updates_on_boot=None, lifecycle_event_configuration=None, load_based_auto_scaling=None, name=None, packages=None, shortname=None, tags=None, use_ebs_optimized_instances=None, volume_configurations=None):
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if auto_assign_elastic_ips and not isinstance(auto_assign_elastic_ips, bool):
            raise TypeError("Expected argument 'auto_assign_elastic_ips' to be a bool")
        pulumi.set(__self__, "auto_assign_elastic_ips", auto_assign_elastic_ips)
        if auto_assign_public_ips and not isinstance(auto_assign_public_ips, bool):
            raise TypeError("Expected argument 'auto_assign_public_ips' to be a bool")
        pulumi.set(__self__, "auto_assign_public_ips", auto_assign_public_ips)
        if custom_instance_profile_arn and not isinstance(custom_instance_profile_arn, str):
            raise TypeError("Expected argument 'custom_instance_profile_arn' to be a str")
        pulumi.set(__self__, "custom_instance_profile_arn", custom_instance_profile_arn)
        if custom_json and not isinstance(custom_json, dict):
            raise TypeError("Expected argument 'custom_json' to be a dict")
        pulumi.set(__self__, "custom_json", custom_json)
        if custom_recipes and not isinstance(custom_recipes, dict):
            raise TypeError("Expected argument 'custom_recipes' to be a dict")
        pulumi.set(__self__, "custom_recipes", custom_recipes)
        if custom_security_group_ids and not isinstance(custom_security_group_ids, list):
            raise TypeError("Expected argument 'custom_security_group_ids' to be a list")
        pulumi.set(__self__, "custom_security_group_ids", custom_security_group_ids)
        if enable_auto_healing and not isinstance(enable_auto_healing, bool):
            raise TypeError("Expected argument 'enable_auto_healing' to be a bool")
        pulumi.set(__self__, "enable_auto_healing", enable_auto_healing)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if install_updates_on_boot and not isinstance(install_updates_on_boot, bool):
            raise TypeError("Expected argument 'install_updates_on_boot' to be a bool")
        pulumi.set(__self__, "install_updates_on_boot", install_updates_on_boot)
        if lifecycle_event_configuration and not isinstance(lifecycle_event_configuration, dict):
            raise TypeError("Expected argument 'lifecycle_event_configuration' to be a dict")
        pulumi.set(__self__, "lifecycle_event_configuration", lifecycle_event_configuration)
        if load_based_auto_scaling and not isinstance(load_based_auto_scaling, dict):
            raise TypeError("Expected argument 'load_based_auto_scaling' to be a dict")
        pulumi.set(__self__, "load_based_auto_scaling", load_based_auto_scaling)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if packages and not isinstance(packages, list):
            raise TypeError("Expected argument 'packages' to be a list")
        pulumi.set(__self__, "packages", packages)
        if shortname and not isinstance(shortname, str):
            raise TypeError("Expected argument 'shortname' to be a str")
        pulumi.set(__self__, "shortname", shortname)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if use_ebs_optimized_instances and not isinstance(use_ebs_optimized_instances, bool):
            raise TypeError("Expected argument 'use_ebs_optimized_instances' to be a bool")
        pulumi.set(__self__, "use_ebs_optimized_instances", use_ebs_optimized_instances)
        if volume_configurations and not isinstance(volume_configurations, list):
            raise TypeError("Expected argument 'volume_configurations' to be a list")
        pulumi.set(__self__, "volume_configurations", volume_configurations)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Any]:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="autoAssignElasticIps")
    def auto_assign_elastic_ips(self) -> Optional[bool]:
        return pulumi.get(self, "auto_assign_elastic_ips")

    @property
    @pulumi.getter(name="autoAssignPublicIps")
    def auto_assign_public_ips(self) -> Optional[bool]:
        return pulumi.get(self, "auto_assign_public_ips")

    @property
    @pulumi.getter(name="customInstanceProfileArn")
    def custom_instance_profile_arn(self) -> Optional[str]:
        return pulumi.get(self, "custom_instance_profile_arn")

    @property
    @pulumi.getter(name="customJson")
    def custom_json(self) -> Optional[Any]:
        return pulumi.get(self, "custom_json")

    @property
    @pulumi.getter(name="customRecipes")
    def custom_recipes(self) -> Optional['outputs.LayerRecipes']:
        return pulumi.get(self, "custom_recipes")

    @property
    @pulumi.getter(name="customSecurityGroupIds")
    def custom_security_group_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "custom_security_group_ids")

    @property
    @pulumi.getter(name="enableAutoHealing")
    def enable_auto_healing(self) -> Optional[bool]:
        return pulumi.get(self, "enable_auto_healing")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="installUpdatesOnBoot")
    def install_updates_on_boot(self) -> Optional[bool]:
        return pulumi.get(self, "install_updates_on_boot")

    @property
    @pulumi.getter(name="lifecycleEventConfiguration")
    def lifecycle_event_configuration(self) -> Optional['outputs.LayerLifecycleEventConfiguration']:
        return pulumi.get(self, "lifecycle_event_configuration")

    @property
    @pulumi.getter(name="loadBasedAutoScaling")
    def load_based_auto_scaling(self) -> Optional['outputs.LayerLoadBasedAutoScaling']:
        return pulumi.get(self, "load_based_auto_scaling")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def packages(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "packages")

    @property
    @pulumi.getter
    def shortname(self) -> Optional[str]:
        return pulumi.get(self, "shortname")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.LayerTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="useEbsOptimizedInstances")
    def use_ebs_optimized_instances(self) -> Optional[bool]:
        return pulumi.get(self, "use_ebs_optimized_instances")

    @property
    @pulumi.getter(name="volumeConfigurations")
    def volume_configurations(self) -> Optional[Sequence['outputs.LayerVolumeConfiguration']]:
        return pulumi.get(self, "volume_configurations")


class AwaitableGetLayerResult(GetLayerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLayerResult(
            attributes=self.attributes,
            auto_assign_elastic_ips=self.auto_assign_elastic_ips,
            auto_assign_public_ips=self.auto_assign_public_ips,
            custom_instance_profile_arn=self.custom_instance_profile_arn,
            custom_json=self.custom_json,
            custom_recipes=self.custom_recipes,
            custom_security_group_ids=self.custom_security_group_ids,
            enable_auto_healing=self.enable_auto_healing,
            id=self.id,
            install_updates_on_boot=self.install_updates_on_boot,
            lifecycle_event_configuration=self.lifecycle_event_configuration,
            load_based_auto_scaling=self.load_based_auto_scaling,
            name=self.name,
            packages=self.packages,
            shortname=self.shortname,
            tags=self.tags,
            use_ebs_optimized_instances=self.use_ebs_optimized_instances,
            volume_configurations=self.volume_configurations)


def get_layer(id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLayerResult:
    """
    Resource Type definition for AWS::OpsWorks::Layer
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:opsworks:getLayer', __args__, opts=opts, typ=GetLayerResult).value

    return AwaitableGetLayerResult(
        attributes=__ret__.attributes,
        auto_assign_elastic_ips=__ret__.auto_assign_elastic_ips,
        auto_assign_public_ips=__ret__.auto_assign_public_ips,
        custom_instance_profile_arn=__ret__.custom_instance_profile_arn,
        custom_json=__ret__.custom_json,
        custom_recipes=__ret__.custom_recipes,
        custom_security_group_ids=__ret__.custom_security_group_ids,
        enable_auto_healing=__ret__.enable_auto_healing,
        id=__ret__.id,
        install_updates_on_boot=__ret__.install_updates_on_boot,
        lifecycle_event_configuration=__ret__.lifecycle_event_configuration,
        load_based_auto_scaling=__ret__.load_based_auto_scaling,
        name=__ret__.name,
        packages=__ret__.packages,
        shortname=__ret__.shortname,
        tags=__ret__.tags,
        use_ebs_optimized_instances=__ret__.use_ebs_optimized_instances,
        volume_configurations=__ret__.volume_configurations)


@_utilities.lift_output_func(get_layer)
def get_layer_output(id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLayerResult]:
    """
    Resource Type definition for AWS::OpsWorks::Layer
    """
    ...