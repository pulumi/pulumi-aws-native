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
from ._enums import *
from ._inputs import *

__all__ = ['AnomalySubscriptionArgs', 'AnomalySubscription']

@pulumi.input_type
class AnomalySubscriptionArgs:
    def __init__(__self__, *,
                 frequency: pulumi.Input['AnomalySubscriptionFrequency'],
                 monitor_arn_list: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subscribers: pulumi.Input[Sequence[pulumi.Input['AnomalySubscriptionSubscriberArgs']]],
                 subscription_name: pulumi.Input[str],
                 threshold: pulumi.Input[float],
                 resource_tags: Optional[pulumi.Input[Sequence[pulumi.Input['AnomalySubscriptionResourceTagArgs']]]] = None):
        """
        The set of arguments for constructing a AnomalySubscription resource.
        :param pulumi.Input['AnomalySubscriptionFrequency'] frequency: The frequency at which anomaly reports are sent over email. 
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitor_arn_list: A list of cost anomaly monitors.
        :param pulumi.Input[Sequence[pulumi.Input['AnomalySubscriptionSubscriberArgs']]] subscribers: A list of subscriber
        :param pulumi.Input[str] subscription_name: The name of the subscription.
        :param pulumi.Input[float] threshold: The dollar value that triggers a notification if the threshold is exceeded. 
        :param pulumi.Input[Sequence[pulumi.Input['AnomalySubscriptionResourceTagArgs']]] resource_tags: Tags to assign to subscription.
        """
        pulumi.set(__self__, "frequency", frequency)
        pulumi.set(__self__, "monitor_arn_list", monitor_arn_list)
        pulumi.set(__self__, "subscribers", subscribers)
        pulumi.set(__self__, "subscription_name", subscription_name)
        pulumi.set(__self__, "threshold", threshold)
        if resource_tags is not None:
            pulumi.set(__self__, "resource_tags", resource_tags)

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Input['AnomalySubscriptionFrequency']:
        """
        The frequency at which anomaly reports are sent over email. 
        """
        return pulumi.get(self, "frequency")

    @frequency.setter
    def frequency(self, value: pulumi.Input['AnomalySubscriptionFrequency']):
        pulumi.set(self, "frequency", value)

    @property
    @pulumi.getter(name="monitorArnList")
    def monitor_arn_list(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of cost anomaly monitors.
        """
        return pulumi.get(self, "monitor_arn_list")

    @monitor_arn_list.setter
    def monitor_arn_list(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "monitor_arn_list", value)

    @property
    @pulumi.getter
    def subscribers(self) -> pulumi.Input[Sequence[pulumi.Input['AnomalySubscriptionSubscriberArgs']]]:
        """
        A list of subscriber
        """
        return pulumi.get(self, "subscribers")

    @subscribers.setter
    def subscribers(self, value: pulumi.Input[Sequence[pulumi.Input['AnomalySubscriptionSubscriberArgs']]]):
        pulumi.set(self, "subscribers", value)

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> pulumi.Input[str]:
        """
        The name of the subscription.
        """
        return pulumi.get(self, "subscription_name")

    @subscription_name.setter
    def subscription_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "subscription_name", value)

    @property
    @pulumi.getter
    def threshold(self) -> pulumi.Input[float]:
        """
        The dollar value that triggers a notification if the threshold is exceeded. 
        """
        return pulumi.get(self, "threshold")

    @threshold.setter
    def threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "threshold", value)

    @property
    @pulumi.getter(name="resourceTags")
    def resource_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AnomalySubscriptionResourceTagArgs']]]]:
        """
        Tags to assign to subscription.
        """
        return pulumi.get(self, "resource_tags")

    @resource_tags.setter
    def resource_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AnomalySubscriptionResourceTagArgs']]]]):
        pulumi.set(self, "resource_tags", value)


warnings.warn("""AnomalySubscription is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class AnomalySubscription(pulumi.CustomResource):
    warnings.warn("""AnomalySubscription is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 frequency: Optional[pulumi.Input['AnomalySubscriptionFrequency']] = None,
                 monitor_arn_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnomalySubscriptionResourceTagArgs']]]]] = None,
                 subscribers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnomalySubscriptionSubscriberArgs']]]]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 threshold: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        """
        AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. Create subscription to be notified

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['AnomalySubscriptionFrequency'] frequency: The frequency at which anomaly reports are sent over email. 
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitor_arn_list: A list of cost anomaly monitors.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnomalySubscriptionResourceTagArgs']]]] resource_tags: Tags to assign to subscription.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnomalySubscriptionSubscriberArgs']]]] subscribers: A list of subscriber
        :param pulumi.Input[str] subscription_name: The name of the subscription.
        :param pulumi.Input[float] threshold: The dollar value that triggers a notification if the threshold is exceeded. 
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AnomalySubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. Create subscription to be notified

        :param str resource_name: The name of the resource.
        :param AnomalySubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AnomalySubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 frequency: Optional[pulumi.Input['AnomalySubscriptionFrequency']] = None,
                 monitor_arn_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnomalySubscriptionResourceTagArgs']]]]] = None,
                 subscribers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnomalySubscriptionSubscriberArgs']]]]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 threshold: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        pulumi.log.warn("""AnomalySubscription is deprecated: AnomalySubscription is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AnomalySubscriptionArgs.__new__(AnomalySubscriptionArgs)

            if frequency is None and not opts.urn:
                raise TypeError("Missing required property 'frequency'")
            __props__.__dict__["frequency"] = frequency
            if monitor_arn_list is None and not opts.urn:
                raise TypeError("Missing required property 'monitor_arn_list'")
            __props__.__dict__["monitor_arn_list"] = monitor_arn_list
            __props__.__dict__["resource_tags"] = resource_tags
            if subscribers is None and not opts.urn:
                raise TypeError("Missing required property 'subscribers'")
            __props__.__dict__["subscribers"] = subscribers
            if subscription_name is None and not opts.urn:
                raise TypeError("Missing required property 'subscription_name'")
            __props__.__dict__["subscription_name"] = subscription_name
            if threshold is None and not opts.urn:
                raise TypeError("Missing required property 'threshold'")
            __props__.__dict__["threshold"] = threshold
            __props__.__dict__["account_id"] = None
            __props__.__dict__["subscription_arn"] = None
        super(AnomalySubscription, __self__).__init__(
            'aws-native:ce:AnomalySubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AnomalySubscription':
        """
        Get an existing AnomalySubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AnomalySubscriptionArgs.__new__(AnomalySubscriptionArgs)

        __props__.__dict__["account_id"] = None
        __props__.__dict__["frequency"] = None
        __props__.__dict__["monitor_arn_list"] = None
        __props__.__dict__["resource_tags"] = None
        __props__.__dict__["subscribers"] = None
        __props__.__dict__["subscription_arn"] = None
        __props__.__dict__["subscription_name"] = None
        __props__.__dict__["threshold"] = None
        return AnomalySubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The accountId
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Output['AnomalySubscriptionFrequency']:
        """
        The frequency at which anomaly reports are sent over email. 
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter(name="monitorArnList")
    def monitor_arn_list(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of cost anomaly monitors.
        """
        return pulumi.get(self, "monitor_arn_list")

    @property
    @pulumi.getter(name="resourceTags")
    def resource_tags(self) -> pulumi.Output[Optional[Sequence['outputs.AnomalySubscriptionResourceTag']]]:
        """
        Tags to assign to subscription.
        """
        return pulumi.get(self, "resource_tags")

    @property
    @pulumi.getter
    def subscribers(self) -> pulumi.Output[Sequence['outputs.AnomalySubscriptionSubscriber']]:
        """
        A list of subscriber
        """
        return pulumi.get(self, "subscribers")

    @property
    @pulumi.getter(name="subscriptionArn")
    def subscription_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "subscription_arn")

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> pulumi.Output[str]:
        """
        The name of the subscription.
        """
        return pulumi.get(self, "subscription_name")

    @property
    @pulumi.getter
    def threshold(self) -> pulumi.Output[float]:
        """
        The dollar value that triggers a notification if the threshold is exceeded. 
        """
        return pulumi.get(self, "threshold")
