# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ClusterStatus',
    'ControlPanelStatus',
    'RoutingControlStatus',
    'SafetyRuleRuleType',
    'SafetyRuleStatus',
]


class ClusterStatus(str, Enum):
    """
    Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
    """
    PENDING = "PENDING"
    DEPLOYED = "DEPLOYED"
    PENDING_DELETION = "PENDING_DELETION"


class ControlPanelStatus(str, Enum):
    """
    The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
    """
    PENDING = "PENDING"
    DEPLOYED = "DEPLOYED"
    PENDING_DELETION = "PENDING_DELETION"


class RoutingControlStatus(str, Enum):
    """
    The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
    """
    PENDING = "PENDING"
    DEPLOYED = "DEPLOYED"
    PENDING_DELETION = "PENDING_DELETION"


class SafetyRuleRuleType(str, Enum):
    """
    A rule can be one of the following: ATLEAST, AND, or OR.
    """
    AND_ = "AND"
    OR_ = "OR"
    ATLEAST = "ATLEAST"


class SafetyRuleStatus(str, Enum):
    """
    The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
    """
    PENDING = "PENDING"
    DEPLOYED = "DEPLOYED"
    PENDING_DELETION = "PENDING_DELETION"