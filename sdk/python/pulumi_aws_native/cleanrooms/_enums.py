# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CollaborationMemberAbility',
    'CollaborationQueryLogStatus',
    'ConfiguredTableAnalysisMethod',
    'ConfiguredTableAnalysisRuleType',
    'MembershipQueryLogStatus',
]


class CollaborationMemberAbility(str, Enum):
    CAN_QUERY = "CAN_QUERY"
    CAN_RECEIVE_RESULTS = "CAN_RECEIVE_RESULTS"


class CollaborationQueryLogStatus(str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class ConfiguredTableAnalysisMethod(str, Enum):
    DIRECT_QUERY = "DIRECT_QUERY"


class ConfiguredTableAnalysisRuleType(str, Enum):
    AGGREGATION = "AGGREGATION"
    LIST = "LIST"


class MembershipQueryLogStatus(str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"