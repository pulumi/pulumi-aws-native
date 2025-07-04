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
from ._enums import *
from ._inputs import *

__all__ = ['AppArgs', 'App']

@pulumi.input_type
class AppArgs:
    def __init__(__self__, *,
                 app_template_body: pulumi.Input[builtins.str],
                 resource_mappings: pulumi.Input[Sequence[pulumi.Input['AppResourceMappingArgs']]],
                 app_assessment_schedule: Optional[pulumi.Input['AppAssessmentSchedule']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 event_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input['AppEventSubscriptionArgs']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 permission_model: Optional[pulumi.Input['AppPermissionModelArgs']] = None,
                 resiliency_policy_arn: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a App resource.
        :param pulumi.Input[builtins.str] app_template_body: A string containing full ResilienceHub app template body.
        :param pulumi.Input[Sequence[pulumi.Input['AppResourceMappingArgs']]] resource_mappings: An array of ResourceMapping objects.
        :param pulumi.Input['AppAssessmentSchedule'] app_assessment_schedule: Assessment execution schedule.
        :param pulumi.Input[builtins.str] description: App description.
        :param pulumi.Input[Sequence[pulumi.Input['AppEventSubscriptionArgs']]] event_subscriptions: The list of events you would like to subscribe and get notification for.
        :param pulumi.Input[builtins.str] name: Name of the app.
        :param pulumi.Input['AppPermissionModelArgs'] permission_model: Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment.
        :param pulumi.Input[builtins.str] resiliency_policy_arn: Amazon Resource Name (ARN) of the Resiliency Policy.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Tags assigned to the resource. A tag is a label that you assign to an AWS resource. Each tag consists of a key/value pair.
        """
        pulumi.set(__self__, "app_template_body", app_template_body)
        pulumi.set(__self__, "resource_mappings", resource_mappings)
        if app_assessment_schedule is not None:
            pulumi.set(__self__, "app_assessment_schedule", app_assessment_schedule)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if event_subscriptions is not None:
            pulumi.set(__self__, "event_subscriptions", event_subscriptions)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if permission_model is not None:
            pulumi.set(__self__, "permission_model", permission_model)
        if resiliency_policy_arn is not None:
            pulumi.set(__self__, "resiliency_policy_arn", resiliency_policy_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="appTemplateBody")
    def app_template_body(self) -> pulumi.Input[builtins.str]:
        """
        A string containing full ResilienceHub app template body.
        """
        return pulumi.get(self, "app_template_body")

    @app_template_body.setter
    def app_template_body(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "app_template_body", value)

    @property
    @pulumi.getter(name="resourceMappings")
    def resource_mappings(self) -> pulumi.Input[Sequence[pulumi.Input['AppResourceMappingArgs']]]:
        """
        An array of ResourceMapping objects.
        """
        return pulumi.get(self, "resource_mappings")

    @resource_mappings.setter
    def resource_mappings(self, value: pulumi.Input[Sequence[pulumi.Input['AppResourceMappingArgs']]]):
        pulumi.set(self, "resource_mappings", value)

    @property
    @pulumi.getter(name="appAssessmentSchedule")
    def app_assessment_schedule(self) -> Optional[pulumi.Input['AppAssessmentSchedule']]:
        """
        Assessment execution schedule.
        """
        return pulumi.get(self, "app_assessment_schedule")

    @app_assessment_schedule.setter
    def app_assessment_schedule(self, value: Optional[pulumi.Input['AppAssessmentSchedule']]):
        pulumi.set(self, "app_assessment_schedule", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        App description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="eventSubscriptions")
    def event_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AppEventSubscriptionArgs']]]]:
        """
        The list of events you would like to subscribe and get notification for.
        """
        return pulumi.get(self, "event_subscriptions")

    @event_subscriptions.setter
    def event_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AppEventSubscriptionArgs']]]]):
        pulumi.set(self, "event_subscriptions", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the app.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="permissionModel")
    def permission_model(self) -> Optional[pulumi.Input['AppPermissionModelArgs']]:
        """
        Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment.
        """
        return pulumi.get(self, "permission_model")

    @permission_model.setter
    def permission_model(self, value: Optional[pulumi.Input['AppPermissionModelArgs']]):
        pulumi.set(self, "permission_model", value)

    @property
    @pulumi.getter(name="resiliencyPolicyArn")
    def resiliency_policy_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Amazon Resource Name (ARN) of the Resiliency Policy.
        """
        return pulumi.get(self, "resiliency_policy_arn")

    @resiliency_policy_arn.setter
    def resiliency_policy_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resiliency_policy_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Tags assigned to the resource. A tag is a label that you assign to an AWS resource. Each tag consists of a key/value pair.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:resiliencehub:App")
class App(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_assessment_schedule: Optional[pulumi.Input['AppAssessmentSchedule']] = None,
                 app_template_body: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 event_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AppEventSubscriptionArgs', 'AppEventSubscriptionArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 permission_model: Optional[pulumi.Input[Union['AppPermissionModelArgs', 'AppPermissionModelArgsDict']]] = None,
                 resiliency_policy_arn: Optional[pulumi.Input[builtins.str]] = None,
                 resource_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AppResourceMappingArgs', 'AppResourceMappingArgsDict']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Resource Type Definition for AWS::ResilienceHub::App.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['AppAssessmentSchedule'] app_assessment_schedule: Assessment execution schedule.
        :param pulumi.Input[builtins.str] app_template_body: A string containing full ResilienceHub app template body.
        :param pulumi.Input[builtins.str] description: App description.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AppEventSubscriptionArgs', 'AppEventSubscriptionArgsDict']]]] event_subscriptions: The list of events you would like to subscribe and get notification for.
        :param pulumi.Input[builtins.str] name: Name of the app.
        :param pulumi.Input[Union['AppPermissionModelArgs', 'AppPermissionModelArgsDict']] permission_model: Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment.
        :param pulumi.Input[builtins.str] resiliency_policy_arn: Amazon Resource Name (ARN) of the Resiliency Policy.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AppResourceMappingArgs', 'AppResourceMappingArgsDict']]]] resource_mappings: An array of ResourceMapping objects.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Tags assigned to the resource. A tag is a label that you assign to an AWS resource. Each tag consists of a key/value pair.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type Definition for AWS::ResilienceHub::App.

        :param str resource_name: The name of the resource.
        :param AppArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_assessment_schedule: Optional[pulumi.Input['AppAssessmentSchedule']] = None,
                 app_template_body: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 event_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AppEventSubscriptionArgs', 'AppEventSubscriptionArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 permission_model: Optional[pulumi.Input[Union['AppPermissionModelArgs', 'AppPermissionModelArgsDict']]] = None,
                 resiliency_policy_arn: Optional[pulumi.Input[builtins.str]] = None,
                 resource_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AppResourceMappingArgs', 'AppResourceMappingArgsDict']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppArgs.__new__(AppArgs)

            __props__.__dict__["app_assessment_schedule"] = app_assessment_schedule
            if app_template_body is None and not opts.urn:
                raise TypeError("Missing required property 'app_template_body'")
            __props__.__dict__["app_template_body"] = app_template_body
            __props__.__dict__["description"] = description
            __props__.__dict__["event_subscriptions"] = event_subscriptions
            __props__.__dict__["name"] = name
            __props__.__dict__["permission_model"] = permission_model
            __props__.__dict__["resiliency_policy_arn"] = resiliency_policy_arn
            if resource_mappings is None and not opts.urn:
                raise TypeError("Missing required property 'resource_mappings'")
            __props__.__dict__["resource_mappings"] = resource_mappings
            __props__.__dict__["tags"] = tags
            __props__.__dict__["app_arn"] = None
            __props__.__dict__["drift_status"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(App, __self__).__init__(
            'aws-native:resiliencehub:App',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'App':
        """
        Get an existing App resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AppArgs.__new__(AppArgs)

        __props__.__dict__["app_arn"] = None
        __props__.__dict__["app_assessment_schedule"] = None
        __props__.__dict__["app_template_body"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["drift_status"] = None
        __props__.__dict__["event_subscriptions"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["permission_model"] = None
        __props__.__dict__["resiliency_policy_arn"] = None
        __props__.__dict__["resource_mappings"] = None
        __props__.__dict__["tags"] = None
        return App(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appArn")
    def app_arn(self) -> pulumi.Output[builtins.str]:
        """
        Amazon Resource Name (ARN) of the App.
        """
        return pulumi.get(self, "app_arn")

    @property
    @pulumi.getter(name="appAssessmentSchedule")
    def app_assessment_schedule(self) -> pulumi.Output[Optional['AppAssessmentSchedule']]:
        """
        Assessment execution schedule.
        """
        return pulumi.get(self, "app_assessment_schedule")

    @property
    @pulumi.getter(name="appTemplateBody")
    def app_template_body(self) -> pulumi.Output[builtins.str]:
        """
        A string containing full ResilienceHub app template body.
        """
        return pulumi.get(self, "app_template_body")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        App description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="driftStatus")
    def drift_status(self) -> pulumi.Output['AppDriftStatus']:
        """
        Indicates if compliance drifts (deviations) were detected while running an assessment for your application.
        """
        return pulumi.get(self, "drift_status")

    @property
    @pulumi.getter(name="eventSubscriptions")
    def event_subscriptions(self) -> pulumi.Output[Optional[Sequence['outputs.AppEventSubscription']]]:
        """
        The list of events you would like to subscribe and get notification for.
        """
        return pulumi.get(self, "event_subscriptions")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the app.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="permissionModel")
    def permission_model(self) -> pulumi.Output[Optional['outputs.AppPermissionModel']]:
        """
        Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment.
        """
        return pulumi.get(self, "permission_model")

    @property
    @pulumi.getter(name="resiliencyPolicyArn")
    def resiliency_policy_arn(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Amazon Resource Name (ARN) of the Resiliency Policy.
        """
        return pulumi.get(self, "resiliency_policy_arn")

    @property
    @pulumi.getter(name="resourceMappings")
    def resource_mappings(self) -> pulumi.Output[Sequence['outputs.AppResourceMapping']]:
        """
        An array of ResourceMapping objects.
        """
        return pulumi.get(self, "resource_mappings")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Tags assigned to the resource. A tag is a label that you assign to an AWS resource. Each tag consists of a key/value pair.
        """
        return pulumi.get(self, "tags")

