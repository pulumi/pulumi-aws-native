# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'ApplicationApiGatewayProxyInput',
    'ApplicationTag',
    'EnvironmentTag',
    'RouteDefaultRouteInput',
    'RouteTag',
    'RouteUriPathRouteInput',
    'ServiceLambdaEndpointInput',
    'ServiceTag',
    'ServiceUrlEndpointInput',
]

@pulumi.output_type
class ApplicationApiGatewayProxyInput(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "endpointType":
            suggest = "endpoint_type"
        elif key == "stageName":
            suggest = "stage_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationApiGatewayProxyInput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationApiGatewayProxyInput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationApiGatewayProxyInput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 endpoint_type: Optional['ApplicationApiGatewayEndpointType'] = None,
                 stage_name: Optional[str] = None):
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)
        if stage_name is not None:
            pulumi.set(__self__, "stage_name", stage_name)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional['ApplicationApiGatewayEndpointType']:
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> Optional[str]:
        return pulumi.get(self, "stage_name")


@pulumi.output_type
class ApplicationTag(dict):
    """
    A label for tagging Environment resource
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A label for tagging Environment resource
        :param str key: A string used to identify this tag
        :param str value: A string containing the value for the tag
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        A string used to identify this tag
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        A string containing the value for the tag
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class EnvironmentTag(dict):
    """
    A label for tagging Environment resource
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A label for tagging Environment resource
        :param str key: A string used to identify this tag
        :param str value: A string containing the value for the tag
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        A string used to identify this tag
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        A string containing the value for the tag
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class RouteDefaultRouteInput(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "activationState":
            suggest = "activation_state"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RouteDefaultRouteInput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RouteDefaultRouteInput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RouteDefaultRouteInput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 activation_state: 'RouteActivationState'):
        pulumi.set(__self__, "activation_state", activation_state)

    @property
    @pulumi.getter(name="activationState")
    def activation_state(self) -> 'RouteActivationState':
        return pulumi.get(self, "activation_state")


@pulumi.output_type
class RouteTag(dict):
    """
    A label for tagging Environment resource
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A label for tagging Environment resource
        :param str key: A string used to identify this tag
        :param str value: A string containing the value for the tag
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        A string used to identify this tag
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        A string containing the value for the tag
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class RouteUriPathRouteInput(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "activationState":
            suggest = "activation_state"
        elif key == "includeChildPaths":
            suggest = "include_child_paths"
        elif key == "sourcePath":
            suggest = "source_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RouteUriPathRouteInput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RouteUriPathRouteInput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RouteUriPathRouteInput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 activation_state: 'RouteActivationState',
                 include_child_paths: Optional[bool] = None,
                 methods: Optional[Sequence['RouteMethod']] = None,
                 source_path: Optional[str] = None):
        pulumi.set(__self__, "activation_state", activation_state)
        if include_child_paths is not None:
            pulumi.set(__self__, "include_child_paths", include_child_paths)
        if methods is not None:
            pulumi.set(__self__, "methods", methods)
        if source_path is not None:
            pulumi.set(__self__, "source_path", source_path)

    @property
    @pulumi.getter(name="activationState")
    def activation_state(self) -> 'RouteActivationState':
        return pulumi.get(self, "activation_state")

    @property
    @pulumi.getter(name="includeChildPaths")
    def include_child_paths(self) -> Optional[bool]:
        return pulumi.get(self, "include_child_paths")

    @property
    @pulumi.getter
    def methods(self) -> Optional[Sequence['RouteMethod']]:
        return pulumi.get(self, "methods")

    @property
    @pulumi.getter(name="sourcePath")
    def source_path(self) -> Optional[str]:
        return pulumi.get(self, "source_path")


@pulumi.output_type
class ServiceLambdaEndpointInput(dict):
    def __init__(__self__, *,
                 arn: str):
        pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")


@pulumi.output_type
class ServiceTag(dict):
    """
    A label for tagging Environment resource
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A label for tagging Environment resource
        :param str key: A string used to identify this tag
        :param str value: A string containing the value for the tag
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        A string used to identify this tag
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        A string containing the value for the tag
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class ServiceUrlEndpointInput(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "healthUrl":
            suggest = "health_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceUrlEndpointInput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceUrlEndpointInput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceUrlEndpointInput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 url: str,
                 health_url: Optional[str] = None):
        pulumi.set(__self__, "url", url)
        if health_url is not None:
            pulumi.set(__self__, "health_url", health_url)

    @property
    @pulumi.getter
    def url(self) -> str:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="healthUrl")
    def health_url(self) -> Optional[str]:
        return pulumi.get(self, "health_url")

