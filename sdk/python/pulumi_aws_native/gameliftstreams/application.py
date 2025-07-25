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
from ._inputs import *

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 application_source_uri: pulumi.Input[builtins.str],
                 description: pulumi.Input[builtins.str],
                 executable_path: pulumi.Input[builtins.str],
                 runtime_environment: pulumi.Input['ApplicationRuntimeEnvironmentArgs'],
                 application_log_output_uri: Optional[pulumi.Input[builtins.str]] = None,
                 application_log_paths: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Application resource.
        :param pulumi.Input[builtins.str] application_source_uri: The location of the content that you want to stream. Enter an Amazon S3 URI to a bucket that contains your game or other application. The location can have a multi-level prefix structure, but it must include all the files needed to run the content. Amazon GameLift Streams copies everything under the specified location.
               
               This value is immutable. To designate a different content location, create a new application.
               
               > The Amazon S3 bucket and the Amazon GameLift Streams application must be in the same AWS Region.
        :param pulumi.Input[builtins.str] description: A human-readable label for the application. You can update this value later.
        :param pulumi.Input[builtins.str] executable_path: The path and file name of the executable file that launches the content for streaming. Enter a path value that is relative to the location set in `ApplicationSourceUri` .
        :param pulumi.Input['ApplicationRuntimeEnvironmentArgs'] runtime_environment: A set of configuration settings to run the application on a stream group. This configures the operating system, and can include compatibility layers and other drivers.
        :param pulumi.Input[builtins.str] application_log_output_uri: An Amazon S3 URI to a bucket where you would like Amazon GameLift Streams to save application logs. Required if you specify one or more `ApplicationLogPaths` .
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] application_log_paths: Locations of log files that your content generates during a stream session. Enter path values that are relative to the `ApplicationSourceUri` location. You can specify up to 10 log paths. Amazon GameLift Streams uploads designated log files to the Amazon S3 bucket that you specify in `ApplicationLogOutputUri` at the end of a stream session. To retrieve stored log files, call [GetStreamSession](https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamSession.html) and get the `LogFileLocationUri` .
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: A list of labels to assign to the new application resource. Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
        """
        pulumi.set(__self__, "application_source_uri", application_source_uri)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "executable_path", executable_path)
        pulumi.set(__self__, "runtime_environment", runtime_environment)
        if application_log_output_uri is not None:
            pulumi.set(__self__, "application_log_output_uri", application_log_output_uri)
        if application_log_paths is not None:
            pulumi.set(__self__, "application_log_paths", application_log_paths)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="applicationSourceUri")
    def application_source_uri(self) -> pulumi.Input[builtins.str]:
        """
        The location of the content that you want to stream. Enter an Amazon S3 URI to a bucket that contains your game or other application. The location can have a multi-level prefix structure, but it must include all the files needed to run the content. Amazon GameLift Streams copies everything under the specified location.

        This value is immutable. To designate a different content location, create a new application.

        > The Amazon S3 bucket and the Amazon GameLift Streams application must be in the same AWS Region.
        """
        return pulumi.get(self, "application_source_uri")

    @application_source_uri.setter
    def application_source_uri(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "application_source_uri", value)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[builtins.str]:
        """
        A human-readable label for the application. You can update this value later.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="executablePath")
    def executable_path(self) -> pulumi.Input[builtins.str]:
        """
        The path and file name of the executable file that launches the content for streaming. Enter a path value that is relative to the location set in `ApplicationSourceUri` .
        """
        return pulumi.get(self, "executable_path")

    @executable_path.setter
    def executable_path(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "executable_path", value)

    @property
    @pulumi.getter(name="runtimeEnvironment")
    def runtime_environment(self) -> pulumi.Input['ApplicationRuntimeEnvironmentArgs']:
        """
        A set of configuration settings to run the application on a stream group. This configures the operating system, and can include compatibility layers and other drivers.
        """
        return pulumi.get(self, "runtime_environment")

    @runtime_environment.setter
    def runtime_environment(self, value: pulumi.Input['ApplicationRuntimeEnvironmentArgs']):
        pulumi.set(self, "runtime_environment", value)

    @property
    @pulumi.getter(name="applicationLogOutputUri")
    def application_log_output_uri(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An Amazon S3 URI to a bucket where you would like Amazon GameLift Streams to save application logs. Required if you specify one or more `ApplicationLogPaths` .
        """
        return pulumi.get(self, "application_log_output_uri")

    @application_log_output_uri.setter
    def application_log_output_uri(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "application_log_output_uri", value)

    @property
    @pulumi.getter(name="applicationLogPaths")
    def application_log_paths(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Locations of log files that your content generates during a stream session. Enter path values that are relative to the `ApplicationSourceUri` location. You can specify up to 10 log paths. Amazon GameLift Streams uploads designated log files to the Amazon S3 bucket that you specify in `ApplicationLogOutputUri` at the end of a stream session. To retrieve stored log files, call [GetStreamSession](https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamSession.html) and get the `LogFileLocationUri` .
        """
        return pulumi.get(self, "application_log_paths")

    @application_log_paths.setter
    def application_log_paths(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "application_log_paths", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A list of labels to assign to the new application resource. Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:gameliftstreams:Application")
class Application(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_log_output_uri: Optional[pulumi.Input[builtins.str]] = None,
                 application_log_paths: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 application_source_uri: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 executable_path: Optional[pulumi.Input[builtins.str]] = None,
                 runtime_environment: Optional[pulumi.Input[Union['ApplicationRuntimeEnvironmentArgs', 'ApplicationRuntimeEnvironmentArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Definition of AWS::GameLiftStreams::Application Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] application_log_output_uri: An Amazon S3 URI to a bucket where you would like Amazon GameLift Streams to save application logs. Required if you specify one or more `ApplicationLogPaths` .
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] application_log_paths: Locations of log files that your content generates during a stream session. Enter path values that are relative to the `ApplicationSourceUri` location. You can specify up to 10 log paths. Amazon GameLift Streams uploads designated log files to the Amazon S3 bucket that you specify in `ApplicationLogOutputUri` at the end of a stream session. To retrieve stored log files, call [GetStreamSession](https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamSession.html) and get the `LogFileLocationUri` .
        :param pulumi.Input[builtins.str] application_source_uri: The location of the content that you want to stream. Enter an Amazon S3 URI to a bucket that contains your game or other application. The location can have a multi-level prefix structure, but it must include all the files needed to run the content. Amazon GameLift Streams copies everything under the specified location.
               
               This value is immutable. To designate a different content location, create a new application.
               
               > The Amazon S3 bucket and the Amazon GameLift Streams application must be in the same AWS Region.
        :param pulumi.Input[builtins.str] description: A human-readable label for the application. You can update this value later.
        :param pulumi.Input[builtins.str] executable_path: The path and file name of the executable file that launches the content for streaming. Enter a path value that is relative to the location set in `ApplicationSourceUri` .
        :param pulumi.Input[Union['ApplicationRuntimeEnvironmentArgs', 'ApplicationRuntimeEnvironmentArgsDict']] runtime_environment: A set of configuration settings to run the application on a stream group. This configures the operating system, and can include compatibility layers and other drivers.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: A list of labels to assign to the new application resource. Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::GameLiftStreams::Application Resource Type

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_log_output_uri: Optional[pulumi.Input[builtins.str]] = None,
                 application_log_paths: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 application_source_uri: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 executable_path: Optional[pulumi.Input[builtins.str]] = None,
                 runtime_environment: Optional[pulumi.Input[Union['ApplicationRuntimeEnvironmentArgs', 'ApplicationRuntimeEnvironmentArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationArgs.__new__(ApplicationArgs)

            __props__.__dict__["application_log_output_uri"] = application_log_output_uri
            __props__.__dict__["application_log_paths"] = application_log_paths
            if application_source_uri is None and not opts.urn:
                raise TypeError("Missing required property 'application_source_uri'")
            __props__.__dict__["application_source_uri"] = application_source_uri
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if executable_path is None and not opts.urn:
                raise TypeError("Missing required property 'executable_path'")
            __props__.__dict__["executable_path"] = executable_path
            if runtime_environment is None and not opts.urn:
                raise TypeError("Missing required property 'runtime_environment'")
            __props__.__dict__["runtime_environment"] = runtime_environment
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["applicationSourceUri", "executablePath", "runtimeEnvironment"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Application, __self__).__init__(
            'aws-native:gameliftstreams:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ApplicationArgs.__new__(ApplicationArgs)

        __props__.__dict__["application_log_output_uri"] = None
        __props__.__dict__["application_log_paths"] = None
        __props__.__dict__["application_source_uri"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["executable_path"] = None
        __props__.__dict__["runtime_environment"] = None
        __props__.__dict__["tags"] = None
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationLogOutputUri")
    def application_log_output_uri(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        An Amazon S3 URI to a bucket where you would like Amazon GameLift Streams to save application logs. Required if you specify one or more `ApplicationLogPaths` .
        """
        return pulumi.get(self, "application_log_output_uri")

    @property
    @pulumi.getter(name="applicationLogPaths")
    def application_log_paths(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        Locations of log files that your content generates during a stream session. Enter path values that are relative to the `ApplicationSourceUri` location. You can specify up to 10 log paths. Amazon GameLift Streams uploads designated log files to the Amazon S3 bucket that you specify in `ApplicationLogOutputUri` at the end of a stream session. To retrieve stored log files, call [GetStreamSession](https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamSession.html) and get the `LogFileLocationUri` .
        """
        return pulumi.get(self, "application_log_paths")

    @property
    @pulumi.getter(name="applicationSourceUri")
    def application_source_uri(self) -> pulumi.Output[builtins.str]:
        """
        The location of the content that you want to stream. Enter an Amazon S3 URI to a bucket that contains your game or other application. The location can have a multi-level prefix structure, but it must include all the files needed to run the content. Amazon GameLift Streams copies everything under the specified location.

        This value is immutable. To designate a different content location, create a new application.

        > The Amazon S3 bucket and the Amazon GameLift Streams application must be in the same AWS Region.
        """
        return pulumi.get(self, "application_source_uri")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the application resource across all AWS Regions. For example:

        `arn:aws:gameliftstreams:us-west-2:123456789012:application/a-9ZY8X7Wv6` .
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[builtins.str]:
        """
        An ID that uniquely identifies the application resource. For example: `a-9ZY8X7Wv6` .
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[builtins.str]:
        """
        A human-readable label for the application. You can update this value later.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="executablePath")
    def executable_path(self) -> pulumi.Output[builtins.str]:
        """
        The path and file name of the executable file that launches the content for streaming. Enter a path value that is relative to the location set in `ApplicationSourceUri` .
        """
        return pulumi.get(self, "executable_path")

    @property
    @pulumi.getter(name="runtimeEnvironment")
    def runtime_environment(self) -> pulumi.Output['outputs.ApplicationRuntimeEnvironment']:
        """
        A set of configuration settings to run the application on a stream group. This configures the operating system, and can include compatibility layers and other drivers.
        """
        return pulumi.get(self, "runtime_environment")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        A list of labels to assign to the new application resource. Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
        """
        return pulumi.get(self, "tags")

