# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AppAutoBranchCreationConfigStage',
    'BranchStage',
]


class AppAutoBranchCreationConfigStage(str, Enum):
    EXPERIMENTAL = "EXPERIMENTAL"
    BETA = "BETA"
    PULL_REQUEST = "PULL_REQUEST"
    PRODUCTION = "PRODUCTION"
    DEVELOPMENT = "DEVELOPMENT"


class BranchStage(str, Enum):
    EXPERIMENTAL = "EXPERIMENTAL"
    BETA = "BETA"
    PULL_REQUEST = "PULL_REQUEST"
    PRODUCTION = "PRODUCTION"
    DEVELOPMENT = "DEVELOPMENT"