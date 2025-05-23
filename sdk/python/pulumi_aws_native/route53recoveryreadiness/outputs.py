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

__all__ = [
    'ResourceSetDnsTargetResource',
    'ResourceSetNlbResource',
    'ResourceSetR53ResourceRecord',
    'ResourceSetResource',
    'ResourceSetTargetResource',
]

@pulumi.output_type
class ResourceSetDnsTargetResource(dict):
    """
    A component for DNS/routing control readiness checks.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "domainName":
            suggest = "domain_name"
        elif key == "hostedZoneArn":
            suggest = "hosted_zone_arn"
        elif key == "recordSetId":
            suggest = "record_set_id"
        elif key == "recordType":
            suggest = "record_type"
        elif key == "targetResource":
            suggest = "target_resource"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceSetDnsTargetResource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceSetDnsTargetResource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceSetDnsTargetResource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 domain_name: Optional[builtins.str] = None,
                 hosted_zone_arn: Optional[builtins.str] = None,
                 record_set_id: Optional[builtins.str] = None,
                 record_type: Optional[builtins.str] = None,
                 target_resource: Optional['outputs.ResourceSetTargetResource'] = None):
        """
        A component for DNS/routing control readiness checks.
        :param builtins.str domain_name: The domain name that acts as an ingress point to a portion of the customer application.
        :param builtins.str hosted_zone_arn: The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.
        :param builtins.str record_set_id: The Route 53 record set ID that will uniquely identify a DNS record, given a name and a type.
        :param builtins.str record_type: The type of DNS record of the target resource.
        :param 'ResourceSetTargetResource' target_resource: The target resource that the Route 53 record points to.
        """
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if hosted_zone_arn is not None:
            pulumi.set(__self__, "hosted_zone_arn", hosted_zone_arn)
        if record_set_id is not None:
            pulumi.set(__self__, "record_set_id", record_set_id)
        if record_type is not None:
            pulumi.set(__self__, "record_type", record_type)
        if target_resource is not None:
            pulumi.set(__self__, "target_resource", target_resource)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[builtins.str]:
        """
        The domain name that acts as an ingress point to a portion of the customer application.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="hostedZoneArn")
    def hosted_zone_arn(self) -> Optional[builtins.str]:
        """
        The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.
        """
        return pulumi.get(self, "hosted_zone_arn")

    @property
    @pulumi.getter(name="recordSetId")
    def record_set_id(self) -> Optional[builtins.str]:
        """
        The Route 53 record set ID that will uniquely identify a DNS record, given a name and a type.
        """
        return pulumi.get(self, "record_set_id")

    @property
    @pulumi.getter(name="recordType")
    def record_type(self) -> Optional[builtins.str]:
        """
        The type of DNS record of the target resource.
        """
        return pulumi.get(self, "record_type")

    @property
    @pulumi.getter(name="targetResource")
    def target_resource(self) -> Optional['outputs.ResourceSetTargetResource']:
        """
        The target resource that the Route 53 record points to.
        """
        return pulumi.get(self, "target_resource")


@pulumi.output_type
class ResourceSetNlbResource(dict):
    """
    The Network Load Balancer resource that a DNS target resource points to.
    """
    def __init__(__self__, *,
                 arn: Optional[builtins.str] = None):
        """
        The Network Load Balancer resource that a DNS target resource points to.
        :param builtins.str arn: A Network Load Balancer resource Amazon Resource Name (ARN).
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        A Network Load Balancer resource Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "arn")


@pulumi.output_type
class ResourceSetR53ResourceRecord(dict):
    """
    The Route 53 resource that a DNS target resource record points to.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "domainName":
            suggest = "domain_name"
        elif key == "recordSetId":
            suggest = "record_set_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceSetR53ResourceRecord. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceSetR53ResourceRecord.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceSetR53ResourceRecord.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 domain_name: Optional[builtins.str] = None,
                 record_set_id: Optional[builtins.str] = None):
        """
        The Route 53 resource that a DNS target resource record points to.
        :param builtins.str domain_name: The DNS target domain name.
        :param builtins.str record_set_id: The Resource Record set id.
        """
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if record_set_id is not None:
            pulumi.set(__self__, "record_set_id", record_set_id)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[builtins.str]:
        """
        The DNS target domain name.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="recordSetId")
    def record_set_id(self) -> Optional[builtins.str]:
        """
        The Resource Record set id.
        """
        return pulumi.get(self, "record_set_id")


@pulumi.output_type
class ResourceSetResource(dict):
    """
    The resource element of a ResourceSet
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "componentId":
            suggest = "component_id"
        elif key == "dnsTargetResource":
            suggest = "dns_target_resource"
        elif key == "readinessScopes":
            suggest = "readiness_scopes"
        elif key == "resourceArn":
            suggest = "resource_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceSetResource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceSetResource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceSetResource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 component_id: Optional[builtins.str] = None,
                 dns_target_resource: Optional['outputs.ResourceSetDnsTargetResource'] = None,
                 readiness_scopes: Optional[Sequence[builtins.str]] = None,
                 resource_arn: Optional[builtins.str] = None):
        """
        The resource element of a ResourceSet
        :param builtins.str component_id: The component identifier of the resource, generated when DNS target resource is used.
        :param 'ResourceSetDnsTargetResource' dns_target_resource: A component for DNS/routing control readiness checks. This is a required setting when `ResourceSet` `ResourceSetType` is set to `AWS::Route53RecoveryReadiness::DNSTargetResource` . Do not set it for any other `ResourceSetType` setting.
        :param Sequence[builtins.str] readiness_scopes: A list of recovery group Amazon Resource Names (ARNs) and cell ARNs that this resource is contained within.
        :param builtins.str resource_arn: The Amazon Resource Name (ARN) of the AWS resource.
        """
        if component_id is not None:
            pulumi.set(__self__, "component_id", component_id)
        if dns_target_resource is not None:
            pulumi.set(__self__, "dns_target_resource", dns_target_resource)
        if readiness_scopes is not None:
            pulumi.set(__self__, "readiness_scopes", readiness_scopes)
        if resource_arn is not None:
            pulumi.set(__self__, "resource_arn", resource_arn)

    @property
    @pulumi.getter(name="componentId")
    def component_id(self) -> Optional[builtins.str]:
        """
        The component identifier of the resource, generated when DNS target resource is used.
        """
        return pulumi.get(self, "component_id")

    @property
    @pulumi.getter(name="dnsTargetResource")
    def dns_target_resource(self) -> Optional['outputs.ResourceSetDnsTargetResource']:
        """
        A component for DNS/routing control readiness checks. This is a required setting when `ResourceSet` `ResourceSetType` is set to `AWS::Route53RecoveryReadiness::DNSTargetResource` . Do not set it for any other `ResourceSetType` setting.
        """
        return pulumi.get(self, "dns_target_resource")

    @property
    @pulumi.getter(name="readinessScopes")
    def readiness_scopes(self) -> Optional[Sequence[builtins.str]]:
        """
        A list of recovery group Amazon Resource Names (ARNs) and cell ARNs that this resource is contained within.
        """
        return pulumi.get(self, "readiness_scopes")

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the AWS resource.
        """
        return pulumi.get(self, "resource_arn")


@pulumi.output_type
class ResourceSetTargetResource(dict):
    """
    The target resource that the Route 53 record points to.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "nlbResource":
            suggest = "nlb_resource"
        elif key == "r53Resource":
            suggest = "r53_resource"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceSetTargetResource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceSetTargetResource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceSetTargetResource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 nlb_resource: Optional['outputs.ResourceSetNlbResource'] = None,
                 r53_resource: Optional['outputs.ResourceSetR53ResourceRecord'] = None):
        """
        The target resource that the Route 53 record points to.
        :param 'ResourceSetNlbResource' nlb_resource: The Network Load Balancer resource that a DNS target resource points to.
        :param 'ResourceSetR53ResourceRecord' r53_resource: The Route 53 resource that a DNS target resource record points to.
        """
        if nlb_resource is not None:
            pulumi.set(__self__, "nlb_resource", nlb_resource)
        if r53_resource is not None:
            pulumi.set(__self__, "r53_resource", r53_resource)

    @property
    @pulumi.getter(name="nlbResource")
    def nlb_resource(self) -> Optional['outputs.ResourceSetNlbResource']:
        """
        The Network Load Balancer resource that a DNS target resource points to.
        """
        return pulumi.get(self, "nlb_resource")

    @property
    @pulumi.getter(name="r53Resource")
    def r53_resource(self) -> Optional['outputs.ResourceSetR53ResourceRecord']:
        """
        The Route 53 resource that a DNS target resource record points to.
        """
        return pulumi.get(self, "r53_resource")


