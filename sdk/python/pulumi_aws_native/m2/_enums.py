# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'ApplicationEngineType',
    'EnvironmentEngineType',
    'EnvironmentNetworkType',
]


class ApplicationEngineType(builtins.str, Enum):
    MICROFOCUS = "microfocus"
    BLUAGE = "bluage"


class EnvironmentEngineType(builtins.str, Enum):
    """
    The target platform for the environment.
    """
    MICROFOCUS = "microfocus"
    BLUAGE = "bluage"


class EnvironmentNetworkType(builtins.str, Enum):
    IPV4 = "ipv4"
    DUAL = "dual"
