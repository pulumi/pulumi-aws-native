# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ExperimentMetricGoalObjectDesiredChange',
    'FeatureEvaluationStrategy',
]


class ExperimentMetricGoalObjectDesiredChange(str, Enum):
    INCREASE = "INCREASE"
    DECREASE = "DECREASE"


class FeatureEvaluationStrategy(str, Enum):
    ALL_RULES = "ALL_RULES"
    DEFAULT_VARIATION = "DEFAULT_VARIATION"