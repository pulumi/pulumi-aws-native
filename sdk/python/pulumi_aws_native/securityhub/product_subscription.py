# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProductSubscriptionArgs', 'ProductSubscription']

@pulumi.input_type
class ProductSubscriptionArgs:
    def __init__(__self__, *,
                 product_arn: pulumi.Input[str]):
        """
        The set of arguments for constructing a ProductSubscription resource.
        :param pulumi.Input[str] product_arn: The generic ARN of the product being subscribed to
        """
        pulumi.set(__self__, "product_arn", product_arn)

    @property
    @pulumi.getter(name="productArn")
    def product_arn(self) -> pulumi.Input[str]:
        """
        The generic ARN of the product being subscribed to
        """
        return pulumi.get(self, "product_arn")

    @product_arn.setter
    def product_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "product_arn", value)


class ProductSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 product_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The AWS::SecurityHub::ProductSubscription resource represents a subscription to a service that is allowed to generate findings for your Security Hub account. One product subscription resource is created for each product enabled.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] product_arn: The generic ARN of the product being subscribed to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProductSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::SecurityHub::ProductSubscription resource represents a subscription to a service that is allowed to generate findings for your Security Hub account. One product subscription resource is created for each product enabled.

        :param str resource_name: The name of the resource.
        :param ProductSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProductSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 product_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProductSubscriptionArgs.__new__(ProductSubscriptionArgs)

            if product_arn is None and not opts.urn:
                raise TypeError("Missing required property 'product_arn'")
            __props__.__dict__["product_arn"] = product_arn
            __props__.__dict__["product_subscription_arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["productArn"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ProductSubscription, __self__).__init__(
            'aws-native:securityhub:ProductSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ProductSubscription':
        """
        Get an existing ProductSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProductSubscriptionArgs.__new__(ProductSubscriptionArgs)

        __props__.__dict__["product_arn"] = None
        __props__.__dict__["product_subscription_arn"] = None
        return ProductSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="productArn")
    def product_arn(self) -> pulumi.Output[str]:
        """
        The generic ARN of the product being subscribed to
        """
        return pulumi.get(self, "product_arn")

    @property
    @pulumi.getter(name="productSubscriptionArn")
    def product_subscription_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the product subscription for the account
        """
        return pulumi.get(self, "product_subscription_arn")
