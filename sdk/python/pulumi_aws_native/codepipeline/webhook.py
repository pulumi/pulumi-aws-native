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
from ._inputs import *

__all__ = ['WebhookArgs', 'Webhook']

@pulumi.input_type
class WebhookArgs:
    def __init__(__self__, *,
                 authentication: pulumi.Input[str],
                 authentication_configuration: pulumi.Input['WebhookAuthConfigurationArgs'],
                 filters: pulumi.Input[Sequence[pulumi.Input['WebhookFilterRuleArgs']]],
                 target_action: pulumi.Input[str],
                 target_pipeline: pulumi.Input[str],
                 target_pipeline_version: pulumi.Input[int],
                 name: Optional[pulumi.Input[str]] = None,
                 register_with_third_party: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Webhook resource.
        """
        pulumi.set(__self__, "authentication", authentication)
        pulumi.set(__self__, "authentication_configuration", authentication_configuration)
        pulumi.set(__self__, "filters", filters)
        pulumi.set(__self__, "target_action", target_action)
        pulumi.set(__self__, "target_pipeline", target_pipeline)
        pulumi.set(__self__, "target_pipeline_version", target_pipeline_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if register_with_third_party is not None:
            pulumi.set(__self__, "register_with_third_party", register_with_third_party)

    @property
    @pulumi.getter
    def authentication(self) -> pulumi.Input[str]:
        return pulumi.get(self, "authentication")

    @authentication.setter
    def authentication(self, value: pulumi.Input[str]):
        pulumi.set(self, "authentication", value)

    @property
    @pulumi.getter(name="authenticationConfiguration")
    def authentication_configuration(self) -> pulumi.Input['WebhookAuthConfigurationArgs']:
        return pulumi.get(self, "authentication_configuration")

    @authentication_configuration.setter
    def authentication_configuration(self, value: pulumi.Input['WebhookAuthConfigurationArgs']):
        pulumi.set(self, "authentication_configuration", value)

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Input[Sequence[pulumi.Input['WebhookFilterRuleArgs']]]:
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: pulumi.Input[Sequence[pulumi.Input['WebhookFilterRuleArgs']]]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter(name="targetAction")
    def target_action(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target_action")

    @target_action.setter
    def target_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_action", value)

    @property
    @pulumi.getter(name="targetPipeline")
    def target_pipeline(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target_pipeline")

    @target_pipeline.setter
    def target_pipeline(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_pipeline", value)

    @property
    @pulumi.getter(name="targetPipelineVersion")
    def target_pipeline_version(self) -> pulumi.Input[int]:
        return pulumi.get(self, "target_pipeline_version")

    @target_pipeline_version.setter
    def target_pipeline_version(self, value: pulumi.Input[int]):
        pulumi.set(self, "target_pipeline_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="registerWithThirdParty")
    def register_with_third_party(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "register_with_third_party")

    @register_with_third_party.setter
    def register_with_third_party(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "register_with_third_party", value)


warnings.warn("""Webhook is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Webhook(pulumi.CustomResource):
    warnings.warn("""Webhook is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication: Optional[pulumi.Input[str]] = None,
                 authentication_configuration: Optional[pulumi.Input[pulumi.InputType['WebhookAuthConfigurationArgs']]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookFilterRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 register_with_third_party: Optional[pulumi.Input[bool]] = None,
                 target_action: Optional[pulumi.Input[str]] = None,
                 target_pipeline: Optional[pulumi.Input[str]] = None,
                 target_pipeline_version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::CodePipeline::Webhook

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::CodePipeline::Webhook

        :param str resource_name: The name of the resource.
        :param WebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication: Optional[pulumi.Input[str]] = None,
                 authentication_configuration: Optional[pulumi.Input[pulumi.InputType['WebhookAuthConfigurationArgs']]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookFilterRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 register_with_third_party: Optional[pulumi.Input[bool]] = None,
                 target_action: Optional[pulumi.Input[str]] = None,
                 target_pipeline: Optional[pulumi.Input[str]] = None,
                 target_pipeline_version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        pulumi.log.warn("""Webhook is deprecated: Webhook is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebhookArgs.__new__(WebhookArgs)

            if authentication is None and not opts.urn:
                raise TypeError("Missing required property 'authentication'")
            __props__.__dict__["authentication"] = authentication
            if authentication_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'authentication_configuration'")
            __props__.__dict__["authentication_configuration"] = authentication_configuration
            if filters is None and not opts.urn:
                raise TypeError("Missing required property 'filters'")
            __props__.__dict__["filters"] = filters
            __props__.__dict__["name"] = name
            __props__.__dict__["register_with_third_party"] = register_with_third_party
            if target_action is None and not opts.urn:
                raise TypeError("Missing required property 'target_action'")
            __props__.__dict__["target_action"] = target_action
            if target_pipeline is None and not opts.urn:
                raise TypeError("Missing required property 'target_pipeline'")
            __props__.__dict__["target_pipeline"] = target_pipeline
            if target_pipeline_version is None and not opts.urn:
                raise TypeError("Missing required property 'target_pipeline_version'")
            __props__.__dict__["target_pipeline_version"] = target_pipeline_version
            __props__.__dict__["url"] = None
        super(Webhook, __self__).__init__(
            'aws-native:codepipeline:Webhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Webhook':
        """
        Get an existing Webhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WebhookArgs.__new__(WebhookArgs)

        __props__.__dict__["authentication"] = None
        __props__.__dict__["authentication_configuration"] = None
        __props__.__dict__["filters"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["register_with_third_party"] = None
        __props__.__dict__["target_action"] = None
        __props__.__dict__["target_pipeline"] = None
        __props__.__dict__["target_pipeline_version"] = None
        __props__.__dict__["url"] = None
        return Webhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authentication(self) -> pulumi.Output[str]:
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter(name="authenticationConfiguration")
    def authentication_configuration(self) -> pulumi.Output['outputs.WebhookAuthConfiguration']:
        return pulumi.get(self, "authentication_configuration")

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Sequence['outputs.WebhookFilterRule']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="registerWithThirdParty")
    def register_with_third_party(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "register_with_third_party")

    @property
    @pulumi.getter(name="targetAction")
    def target_action(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target_action")

    @property
    @pulumi.getter(name="targetPipeline")
    def target_pipeline(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target_pipeline")

    @property
    @pulumi.getter(name="targetPipelineVersion")
    def target_pipeline_version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "target_pipeline_version")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")
