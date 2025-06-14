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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs

__all__ = ['AutoScalingConfigurationArgs', 'AutoScalingConfiguration']

@pulumi.input_type
class AutoScalingConfigurationArgs:
    def __init__(__self__, *,
                 auto_scaling_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                 max_concurrency: Optional[pulumi.Input[builtins.int]] = None,
                 max_size: Optional[pulumi.Input[builtins.int]] = None,
                 min_size: Optional[pulumi.Input[builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]] = None):
        """
        The set of arguments for constructing a AutoScalingConfiguration resource.
        :param pulumi.Input[builtins.str] auto_scaling_configuration_name: The customer-provided auto scaling configuration name.  When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration.
        :param pulumi.Input[builtins.int] max_concurrency: The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests.
        :param pulumi.Input[builtins.int] max_size: The maximum number of instances that an App Runner service scales up to. At most MaxSize instances actively serve traffic for your service.
        :param pulumi.Input[builtins.int] min_size: The minimum number of instances that App Runner provisions for a service. The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]] tags: A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.
        """
        if auto_scaling_configuration_name is not None:
            pulumi.set(__self__, "auto_scaling_configuration_name", auto_scaling_configuration_name)
        if max_concurrency is not None:
            pulumi.set(__self__, "max_concurrency", max_concurrency)
        if max_size is not None:
            pulumi.set(__self__, "max_size", max_size)
        if min_size is not None:
            pulumi.set(__self__, "min_size", min_size)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="autoScalingConfigurationName")
    def auto_scaling_configuration_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The customer-provided auto scaling configuration name.  When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration.
        """
        return pulumi.get(self, "auto_scaling_configuration_name")

    @auto_scaling_configuration_name.setter
    def auto_scaling_configuration_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auto_scaling_configuration_name", value)

    @property
    @pulumi.getter(name="maxConcurrency")
    def max_concurrency(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests.
        """
        return pulumi.get(self, "max_concurrency")

    @max_concurrency.setter
    def max_concurrency(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_concurrency", value)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum number of instances that an App Runner service scales up to. At most MaxSize instances actively serve traffic for your service.
        """
        return pulumi.get(self, "max_size")

    @max_size.setter
    def max_size(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_size", value)

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The minimum number of instances that App Runner provisions for a service. The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.
        """
        return pulumi.get(self, "min_size")

    @min_size.setter
    def min_size(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "min_size", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]]:
        """
        A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:apprunner:AutoScalingConfiguration")
class AutoScalingConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                 max_concurrency: Optional[pulumi.Input[builtins.int]] = None,
                 max_size: Optional[pulumi.Input[builtins.int]] = None,
                 min_size: Optional[pulumi.Input[builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]]] = None,
                 __props__=None):
        """
        Describes an AWS App Runner automatic configuration resource that enables automatic scaling of instances used to process web requests. You can share an auto scaling configuration across multiple services.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] auto_scaling_configuration_name: The customer-provided auto scaling configuration name.  When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration.
        :param pulumi.Input[builtins.int] max_concurrency: The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests.
        :param pulumi.Input[builtins.int] max_size: The maximum number of instances that an App Runner service scales up to. At most MaxSize instances actively serve traffic for your service.
        :param pulumi.Input[builtins.int] min_size: The minimum number of instances that App Runner provisions for a service. The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]] tags: A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AutoScalingConfigurationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Describes an AWS App Runner automatic configuration resource that enables automatic scaling of instances used to process web requests. You can share an auto scaling configuration across multiple services.

        :param str resource_name: The name of the resource.
        :param AutoScalingConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutoScalingConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                 max_concurrency: Optional[pulumi.Input[builtins.int]] = None,
                 max_size: Optional[pulumi.Input[builtins.int]] = None,
                 min_size: Optional[pulumi.Input[builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AutoScalingConfigurationArgs.__new__(AutoScalingConfigurationArgs)

            __props__.__dict__["auto_scaling_configuration_name"] = auto_scaling_configuration_name
            __props__.__dict__["max_concurrency"] = max_concurrency
            __props__.__dict__["max_size"] = max_size
            __props__.__dict__["min_size"] = min_size
            __props__.__dict__["tags"] = tags
            __props__.__dict__["auto_scaling_configuration_arn"] = None
            __props__.__dict__["auto_scaling_configuration_revision"] = None
            __props__.__dict__["latest"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["autoScalingConfigurationName", "maxConcurrency", "maxSize", "minSize", "tags[*]"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(AutoScalingConfiguration, __self__).__init__(
            'aws-native:apprunner:AutoScalingConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AutoScalingConfiguration':
        """
        Get an existing AutoScalingConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AutoScalingConfigurationArgs.__new__(AutoScalingConfigurationArgs)

        __props__.__dict__["auto_scaling_configuration_arn"] = None
        __props__.__dict__["auto_scaling_configuration_name"] = None
        __props__.__dict__["auto_scaling_configuration_revision"] = None
        __props__.__dict__["latest"] = None
        __props__.__dict__["max_concurrency"] = None
        __props__.__dict__["max_size"] = None
        __props__.__dict__["min_size"] = None
        __props__.__dict__["tags"] = None
        return AutoScalingConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoScalingConfigurationArn")
    def auto_scaling_configuration_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of this auto scaling configuration.
        """
        return pulumi.get(self, "auto_scaling_configuration_arn")

    @property
    @pulumi.getter(name="autoScalingConfigurationName")
    def auto_scaling_configuration_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The customer-provided auto scaling configuration name.  When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration.
        """
        return pulumi.get(self, "auto_scaling_configuration_name")

    @property
    @pulumi.getter(name="autoScalingConfigurationRevision")
    def auto_scaling_configuration_revision(self) -> pulumi.Output[builtins.int]:
        """
        The revision of this auto scaling configuration. It's unique among all the active configurations ("Status": "ACTIVE") that share the same AutoScalingConfigurationName.
        """
        return pulumi.get(self, "auto_scaling_configuration_revision")

    @property
    @pulumi.getter
    def latest(self) -> pulumi.Output[builtins.bool]:
        """
        It's set to true for the configuration with the highest Revision among all configurations that share the same AutoScalingConfigurationName. It's set to false otherwise. App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.
        """
        return pulumi.get(self, "latest")

    @property
    @pulumi.getter(name="maxConcurrency")
    def max_concurrency(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests.
        """
        return pulumi.get(self, "max_concurrency")

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The maximum number of instances that an App Runner service scales up to. At most MaxSize instances actively serve traffic for your service.
        """
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The minimum number of instances that App Runner provisions for a service. The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.
        """
        return pulumi.get(self, "min_size")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.CreateOnlyTag']]]:
        """
        A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.
        """
        return pulumi.get(self, "tags")

