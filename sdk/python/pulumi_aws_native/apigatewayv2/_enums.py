# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'DomainNameRoutingMode',
]


class DomainNameRoutingMode(builtins.str, Enum):
    API_MAPPING_ONLY = "API_MAPPING_ONLY"
    ROUTING_RULE_THEN_API_MAPPING = "ROUTING_RULE_THEN_API_MAPPING"
    ROUTING_RULE_ONLY = "ROUTING_RULE_ONLY"
