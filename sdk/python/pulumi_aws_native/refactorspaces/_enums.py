# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'ApplicationApiGatewayEndpointType',
    'ApplicationProxyType',
    'EnvironmentNetworkFabricType',
    'RouteActivationState',
    'RouteMethod',
    'RouteType',
    'ServiceEndpointType',
]


class ApplicationApiGatewayEndpointType(builtins.str, Enum):
    REGIONAL = "REGIONAL"
    PRIVATE = "PRIVATE"


class ApplicationProxyType(builtins.str, Enum):
    API_GATEWAY = "API_GATEWAY"


class EnvironmentNetworkFabricType(builtins.str, Enum):
    TRANSIT_GATEWAY = "TRANSIT_GATEWAY"
    NONE = "NONE"


class RouteActivationState(builtins.str, Enum):
    INACTIVE = "INACTIVE"
    ACTIVE = "ACTIVE"


class RouteMethod(builtins.str, Enum):
    DELETE = "DELETE"
    GET = "GET"
    HEAD = "HEAD"
    OPTIONS = "OPTIONS"
    PATCH = "PATCH"
    POST = "POST"
    PUT = "PUT"


class RouteType(builtins.str, Enum):
    DEFAULT = "DEFAULT"
    URI_PATH = "URI_PATH"


class ServiceEndpointType(builtins.str, Enum):
    LAMBDA_ = "LAMBDA"
    URL = "URL"
