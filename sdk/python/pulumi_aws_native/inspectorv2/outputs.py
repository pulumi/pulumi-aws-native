# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'FilterCriteria',
    'FilterDateFilter',
    'FilterMapFilter',
    'FilterNumberFilter',
    'FilterPackageFilter',
    'FilterPortRangeFilter',
    'FilterStringFilter',
]

@pulumi.output_type
class FilterCriteria(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "awsAccountId":
            suggest = "aws_account_id"
        elif key == "componentId":
            suggest = "component_id"
        elif key == "componentType":
            suggest = "component_type"
        elif key == "ec2InstanceImageId":
            suggest = "ec2_instance_image_id"
        elif key == "ec2InstanceSubnetId":
            suggest = "ec2_instance_subnet_id"
        elif key == "ec2InstanceVpcId":
            suggest = "ec2_instance_vpc_id"
        elif key == "ecrImageArchitecture":
            suggest = "ecr_image_architecture"
        elif key == "ecrImageHash":
            suggest = "ecr_image_hash"
        elif key == "ecrImagePushedAt":
            suggest = "ecr_image_pushed_at"
        elif key == "ecrImageRegistry":
            suggest = "ecr_image_registry"
        elif key == "ecrImageRepositoryName":
            suggest = "ecr_image_repository_name"
        elif key == "ecrImageTags":
            suggest = "ecr_image_tags"
        elif key == "findingArn":
            suggest = "finding_arn"
        elif key == "findingStatus":
            suggest = "finding_status"
        elif key == "findingType":
            suggest = "finding_type"
        elif key == "firstObservedAt":
            suggest = "first_observed_at"
        elif key == "inspectorScore":
            suggest = "inspector_score"
        elif key == "lastObservedAt":
            suggest = "last_observed_at"
        elif key == "networkProtocol":
            suggest = "network_protocol"
        elif key == "portRange":
            suggest = "port_range"
        elif key == "relatedVulnerabilities":
            suggest = "related_vulnerabilities"
        elif key == "resourceId":
            suggest = "resource_id"
        elif key == "resourceTags":
            suggest = "resource_tags"
        elif key == "resourceType":
            suggest = "resource_type"
        elif key == "updatedAt":
            suggest = "updated_at"
        elif key == "vendorSeverity":
            suggest = "vendor_severity"
        elif key == "vulnerabilityId":
            suggest = "vulnerability_id"
        elif key == "vulnerabilitySource":
            suggest = "vulnerability_source"
        elif key == "vulnerablePackages":
            suggest = "vulnerable_packages"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FilterCriteria. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FilterCriteria.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FilterCriteria.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 aws_account_id: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 component_id: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 component_type: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 ec2_instance_image_id: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 ec2_instance_subnet_id: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 ec2_instance_vpc_id: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 ecr_image_architecture: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 ecr_image_hash: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 ecr_image_pushed_at: Optional[Sequence['outputs.FilterDateFilter']] = None,
                 ecr_image_registry: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 ecr_image_repository_name: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 ecr_image_tags: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 finding_arn: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 finding_status: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 finding_type: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 first_observed_at: Optional[Sequence['outputs.FilterDateFilter']] = None,
                 inspector_score: Optional[Sequence['outputs.FilterNumberFilter']] = None,
                 last_observed_at: Optional[Sequence['outputs.FilterDateFilter']] = None,
                 network_protocol: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 port_range: Optional[Sequence['outputs.FilterPortRangeFilter']] = None,
                 related_vulnerabilities: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 resource_id: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 resource_tags: Optional[Sequence['outputs.FilterMapFilter']] = None,
                 resource_type: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 severity: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 title: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 updated_at: Optional[Sequence['outputs.FilterDateFilter']] = None,
                 vendor_severity: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 vulnerability_id: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 vulnerability_source: Optional[Sequence['outputs.FilterStringFilter']] = None,
                 vulnerable_packages: Optional[Sequence['outputs.FilterPackageFilter']] = None):
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if component_id is not None:
            pulumi.set(__self__, "component_id", component_id)
        if component_type is not None:
            pulumi.set(__self__, "component_type", component_type)
        if ec2_instance_image_id is not None:
            pulumi.set(__self__, "ec2_instance_image_id", ec2_instance_image_id)
        if ec2_instance_subnet_id is not None:
            pulumi.set(__self__, "ec2_instance_subnet_id", ec2_instance_subnet_id)
        if ec2_instance_vpc_id is not None:
            pulumi.set(__self__, "ec2_instance_vpc_id", ec2_instance_vpc_id)
        if ecr_image_architecture is not None:
            pulumi.set(__self__, "ecr_image_architecture", ecr_image_architecture)
        if ecr_image_hash is not None:
            pulumi.set(__self__, "ecr_image_hash", ecr_image_hash)
        if ecr_image_pushed_at is not None:
            pulumi.set(__self__, "ecr_image_pushed_at", ecr_image_pushed_at)
        if ecr_image_registry is not None:
            pulumi.set(__self__, "ecr_image_registry", ecr_image_registry)
        if ecr_image_repository_name is not None:
            pulumi.set(__self__, "ecr_image_repository_name", ecr_image_repository_name)
        if ecr_image_tags is not None:
            pulumi.set(__self__, "ecr_image_tags", ecr_image_tags)
        if finding_arn is not None:
            pulumi.set(__self__, "finding_arn", finding_arn)
        if finding_status is not None:
            pulumi.set(__self__, "finding_status", finding_status)
        if finding_type is not None:
            pulumi.set(__self__, "finding_type", finding_type)
        if first_observed_at is not None:
            pulumi.set(__self__, "first_observed_at", first_observed_at)
        if inspector_score is not None:
            pulumi.set(__self__, "inspector_score", inspector_score)
        if last_observed_at is not None:
            pulumi.set(__self__, "last_observed_at", last_observed_at)
        if network_protocol is not None:
            pulumi.set(__self__, "network_protocol", network_protocol)
        if port_range is not None:
            pulumi.set(__self__, "port_range", port_range)
        if related_vulnerabilities is not None:
            pulumi.set(__self__, "related_vulnerabilities", related_vulnerabilities)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_tags is not None:
            pulumi.set(__self__, "resource_tags", resource_tags)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if vendor_severity is not None:
            pulumi.set(__self__, "vendor_severity", vendor_severity)
        if vulnerability_id is not None:
            pulumi.set(__self__, "vulnerability_id", vulnerability_id)
        if vulnerability_source is not None:
            pulumi.set(__self__, "vulnerability_source", vulnerability_source)
        if vulnerable_packages is not None:
            pulumi.set(__self__, "vulnerable_packages", vulnerable_packages)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="componentId")
    def component_id(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "component_id")

    @property
    @pulumi.getter(name="componentType")
    def component_type(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "component_type")

    @property
    @pulumi.getter(name="ec2InstanceImageId")
    def ec2_instance_image_id(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "ec2_instance_image_id")

    @property
    @pulumi.getter(name="ec2InstanceSubnetId")
    def ec2_instance_subnet_id(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "ec2_instance_subnet_id")

    @property
    @pulumi.getter(name="ec2InstanceVpcId")
    def ec2_instance_vpc_id(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "ec2_instance_vpc_id")

    @property
    @pulumi.getter(name="ecrImageArchitecture")
    def ecr_image_architecture(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "ecr_image_architecture")

    @property
    @pulumi.getter(name="ecrImageHash")
    def ecr_image_hash(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "ecr_image_hash")

    @property
    @pulumi.getter(name="ecrImagePushedAt")
    def ecr_image_pushed_at(self) -> Optional[Sequence['outputs.FilterDateFilter']]:
        return pulumi.get(self, "ecr_image_pushed_at")

    @property
    @pulumi.getter(name="ecrImageRegistry")
    def ecr_image_registry(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "ecr_image_registry")

    @property
    @pulumi.getter(name="ecrImageRepositoryName")
    def ecr_image_repository_name(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "ecr_image_repository_name")

    @property
    @pulumi.getter(name="ecrImageTags")
    def ecr_image_tags(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "ecr_image_tags")

    @property
    @pulumi.getter(name="findingArn")
    def finding_arn(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "finding_arn")

    @property
    @pulumi.getter(name="findingStatus")
    def finding_status(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "finding_status")

    @property
    @pulumi.getter(name="findingType")
    def finding_type(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "finding_type")

    @property
    @pulumi.getter(name="firstObservedAt")
    def first_observed_at(self) -> Optional[Sequence['outputs.FilterDateFilter']]:
        return pulumi.get(self, "first_observed_at")

    @property
    @pulumi.getter(name="inspectorScore")
    def inspector_score(self) -> Optional[Sequence['outputs.FilterNumberFilter']]:
        return pulumi.get(self, "inspector_score")

    @property
    @pulumi.getter(name="lastObservedAt")
    def last_observed_at(self) -> Optional[Sequence['outputs.FilterDateFilter']]:
        return pulumi.get(self, "last_observed_at")

    @property
    @pulumi.getter(name="networkProtocol")
    def network_protocol(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "network_protocol")

    @property
    @pulumi.getter(name="portRange")
    def port_range(self) -> Optional[Sequence['outputs.FilterPortRangeFilter']]:
        return pulumi.get(self, "port_range")

    @property
    @pulumi.getter(name="relatedVulnerabilities")
    def related_vulnerabilities(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "related_vulnerabilities")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceTags")
    def resource_tags(self) -> Optional[Sequence['outputs.FilterMapFilter']]:
        return pulumi.get(self, "resource_tags")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def severity(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def title(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[Sequence['outputs.FilterDateFilter']]:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="vendorSeverity")
    def vendor_severity(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "vendor_severity")

    @property
    @pulumi.getter(name="vulnerabilityId")
    def vulnerability_id(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "vulnerability_id")

    @property
    @pulumi.getter(name="vulnerabilitySource")
    def vulnerability_source(self) -> Optional[Sequence['outputs.FilterStringFilter']]:
        return pulumi.get(self, "vulnerability_source")

    @property
    @pulumi.getter(name="vulnerablePackages")
    def vulnerable_packages(self) -> Optional[Sequence['outputs.FilterPackageFilter']]:
        return pulumi.get(self, "vulnerable_packages")


@pulumi.output_type
class FilterDateFilter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "endInclusive":
            suggest = "end_inclusive"
        elif key == "startInclusive":
            suggest = "start_inclusive"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FilterDateFilter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FilterDateFilter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FilterDateFilter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 end_inclusive: Optional[int] = None,
                 start_inclusive: Optional[int] = None):
        if end_inclusive is not None:
            pulumi.set(__self__, "end_inclusive", end_inclusive)
        if start_inclusive is not None:
            pulumi.set(__self__, "start_inclusive", start_inclusive)

    @property
    @pulumi.getter(name="endInclusive")
    def end_inclusive(self) -> Optional[int]:
        return pulumi.get(self, "end_inclusive")

    @property
    @pulumi.getter(name="startInclusive")
    def start_inclusive(self) -> Optional[int]:
        return pulumi.get(self, "start_inclusive")


@pulumi.output_type
class FilterMapFilter(dict):
    def __init__(__self__, *,
                 comparison: 'FilterMapComparison',
                 key: Optional[str] = None,
                 value: Optional[str] = None):
        pulumi.set(__self__, "comparison", comparison)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def comparison(self) -> 'FilterMapComparison':
        return pulumi.get(self, "comparison")

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


@pulumi.output_type
class FilterNumberFilter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "lowerInclusive":
            suggest = "lower_inclusive"
        elif key == "upperInclusive":
            suggest = "upper_inclusive"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FilterNumberFilter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FilterNumberFilter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FilterNumberFilter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 lower_inclusive: Optional[float] = None,
                 upper_inclusive: Optional[float] = None):
        if lower_inclusive is not None:
            pulumi.set(__self__, "lower_inclusive", lower_inclusive)
        if upper_inclusive is not None:
            pulumi.set(__self__, "upper_inclusive", upper_inclusive)

    @property
    @pulumi.getter(name="lowerInclusive")
    def lower_inclusive(self) -> Optional[float]:
        return pulumi.get(self, "lower_inclusive")

    @property
    @pulumi.getter(name="upperInclusive")
    def upper_inclusive(self) -> Optional[float]:
        return pulumi.get(self, "upper_inclusive")


@pulumi.output_type
class FilterPackageFilter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sourceLayerHash":
            suggest = "source_layer_hash"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FilterPackageFilter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FilterPackageFilter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FilterPackageFilter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 architecture: Optional['outputs.FilterStringFilter'] = None,
                 epoch: Optional['outputs.FilterNumberFilter'] = None,
                 name: Optional['outputs.FilterStringFilter'] = None,
                 release: Optional['outputs.FilterStringFilter'] = None,
                 source_layer_hash: Optional['outputs.FilterStringFilter'] = None,
                 version: Optional['outputs.FilterStringFilter'] = None):
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)
        if epoch is not None:
            pulumi.set(__self__, "epoch", epoch)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if release is not None:
            pulumi.set(__self__, "release", release)
        if source_layer_hash is not None:
            pulumi.set(__self__, "source_layer_hash", source_layer_hash)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def architecture(self) -> Optional['outputs.FilterStringFilter']:
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter
    def epoch(self) -> Optional['outputs.FilterNumberFilter']:
        return pulumi.get(self, "epoch")

    @property
    @pulumi.getter
    def name(self) -> Optional['outputs.FilterStringFilter']:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def release(self) -> Optional['outputs.FilterStringFilter']:
        return pulumi.get(self, "release")

    @property
    @pulumi.getter(name="sourceLayerHash")
    def source_layer_hash(self) -> Optional['outputs.FilterStringFilter']:
        return pulumi.get(self, "source_layer_hash")

    @property
    @pulumi.getter
    def version(self) -> Optional['outputs.FilterStringFilter']:
        return pulumi.get(self, "version")


@pulumi.output_type
class FilterPortRangeFilter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "beginInclusive":
            suggest = "begin_inclusive"
        elif key == "endInclusive":
            suggest = "end_inclusive"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FilterPortRangeFilter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FilterPortRangeFilter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FilterPortRangeFilter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 begin_inclusive: Optional[int] = None,
                 end_inclusive: Optional[int] = None):
        if begin_inclusive is not None:
            pulumi.set(__self__, "begin_inclusive", begin_inclusive)
        if end_inclusive is not None:
            pulumi.set(__self__, "end_inclusive", end_inclusive)

    @property
    @pulumi.getter(name="beginInclusive")
    def begin_inclusive(self) -> Optional[int]:
        return pulumi.get(self, "begin_inclusive")

    @property
    @pulumi.getter(name="endInclusive")
    def end_inclusive(self) -> Optional[int]:
        return pulumi.get(self, "end_inclusive")


@pulumi.output_type
class FilterStringFilter(dict):
    def __init__(__self__, *,
                 comparison: 'FilterStringComparison',
                 value: str):
        pulumi.set(__self__, "comparison", comparison)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def comparison(self) -> 'FilterStringComparison':
        return pulumi.get(self, "comparison")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

