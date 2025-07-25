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

__all__ = ['StreamConsumerArgs', 'StreamConsumer']

@pulumi.input_type
class StreamConsumerArgs:
    def __init__(__self__, *,
                 stream_arn: pulumi.Input[builtins.str],
                 consumer_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]] = None):
        """
        The set of arguments for constructing a StreamConsumer resource.
        :param pulumi.Input[builtins.str] stream_arn: The Amazon resource name (ARN) of the Kinesis data stream that you want to register the consumer with.
        :param pulumi.Input[builtins.str] consumer_name: The name of the Kinesis Stream Consumer. For a given Kinesis data stream, each consumer must have a unique name. However, consumer names don't have to be unique across data streams.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]] tags: An arbitrary set of tags (key–value pairs) to associate with the Kinesis consumer.
        """
        pulumi.set(__self__, "stream_arn", stream_arn)
        if consumer_name is not None:
            pulumi.set(__self__, "consumer_name", consumer_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> pulumi.Input[builtins.str]:
        """
        The Amazon resource name (ARN) of the Kinesis data stream that you want to register the consumer with.
        """
        return pulumi.get(self, "stream_arn")

    @stream_arn.setter
    def stream_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "stream_arn", value)

    @property
    @pulumi.getter(name="consumerName")
    def consumer_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Kinesis Stream Consumer. For a given Kinesis data stream, each consumer must have a unique name. However, consumer names don't have to be unique across data streams.
        """
        return pulumi.get(self, "consumer_name")

    @consumer_name.setter
    def consumer_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "consumer_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]]:
        """
        An arbitrary set of tags (key–value pairs) to associate with the Kinesis consumer.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:kinesis:StreamConsumer")
class StreamConsumer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_name: Optional[pulumi.Input[builtins.str]] = None,
                 stream_arn: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Kinesis::StreamConsumer

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] consumer_name: The name of the Kinesis Stream Consumer. For a given Kinesis data stream, each consumer must have a unique name. However, consumer names don't have to be unique across data streams.
        :param pulumi.Input[builtins.str] stream_arn: The Amazon resource name (ARN) of the Kinesis data stream that you want to register the consumer with.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]] tags: An arbitrary set of tags (key–value pairs) to associate with the Kinesis consumer.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StreamConsumerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Kinesis::StreamConsumer

        :param str resource_name: The name of the resource.
        :param StreamConsumerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamConsumerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_name: Optional[pulumi.Input[builtins.str]] = None,
                 stream_arn: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StreamConsumerArgs.__new__(StreamConsumerArgs)

            __props__.__dict__["consumer_name"] = consumer_name
            if stream_arn is None and not opts.urn:
                raise TypeError("Missing required property 'stream_arn'")
            __props__.__dict__["stream_arn"] = stream_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["consumer_arn"] = None
            __props__.__dict__["consumer_creation_timestamp"] = None
            __props__.__dict__["consumer_status"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["consumerName", "streamArn", "tags[*]"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(StreamConsumer, __self__).__init__(
            'aws-native:kinesis:StreamConsumer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'StreamConsumer':
        """
        Get an existing StreamConsumer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = StreamConsumerArgs.__new__(StreamConsumerArgs)

        __props__.__dict__["consumer_arn"] = None
        __props__.__dict__["consumer_creation_timestamp"] = None
        __props__.__dict__["consumer_name"] = None
        __props__.__dict__["consumer_status"] = None
        __props__.__dict__["stream_arn"] = None
        __props__.__dict__["tags"] = None
        return StreamConsumer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="consumerArn")
    def consumer_arn(self) -> pulumi.Output[builtins.str]:
        """
        The ARN returned by Kinesis Data Streams when you registered the consumer. If you don't know the ARN of the consumer that you want to deregister, you can use the ListStreamConsumers operation to get a list of the descriptions of all the consumers that are currently registered with a given data stream. The description of a consumer contains its ARN.
        """
        return pulumi.get(self, "consumer_arn")

    @property
    @pulumi.getter(name="consumerCreationTimestamp")
    def consumer_creation_timestamp(self) -> pulumi.Output[builtins.str]:
        """
        Timestamp when the consumer was created.
        """
        return pulumi.get(self, "consumer_creation_timestamp")

    @property
    @pulumi.getter(name="consumerName")
    def consumer_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the Kinesis Stream Consumer. For a given Kinesis data stream, each consumer must have a unique name. However, consumer names don't have to be unique across data streams.
        """
        return pulumi.get(self, "consumer_name")

    @property
    @pulumi.getter(name="consumerStatus")
    def consumer_status(self) -> pulumi.Output[builtins.str]:
        """
        A consumer can't read data while in the CREATING or DELETING states. Valid Values: CREATING | DELETING | ACTIVE
        """
        return pulumi.get(self, "consumer_status")

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon resource name (ARN) of the Kinesis data stream that you want to register the consumer with.
        """
        return pulumi.get(self, "stream_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.CreateOnlyTag']]]:
        """
        An arbitrary set of tags (key–value pairs) to associate with the Kinesis consumer.
        """
        return pulumi.get(self, "tags")

