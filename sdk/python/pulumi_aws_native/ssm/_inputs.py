# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'AssociationInstanceAssociationOutputLocationArgs',
    'AssociationS3OutputLocationArgs',
    'AssociationTargetArgs',
    'DocumentAttachmentsSourceArgs',
    'DocumentRequiresArgs',
    'PatchBaselinePatchFilterGroupArgs',
    'PatchBaselinePatchFilterArgs',
    'PatchBaselinePatchSourceArgs',
    'PatchBaselineRuleGroupArgs',
    'PatchBaselineRuleArgs',
    'ResourceDataSyncAwsOrganizationsSourceArgs',
    'ResourceDataSyncS3DestinationArgs',
    'ResourceDataSyncSyncSourceArgs',
]

@pulumi.input_type
class AssociationInstanceAssociationOutputLocationArgs:
    def __init__(__self__, *,
                 s3_location: Optional[pulumi.Input['AssociationS3OutputLocationArgs']] = None):
        if s3_location is not None:
            pulumi.set(__self__, "s3_location", s3_location)

    @property
    @pulumi.getter(name="s3Location")
    def s3_location(self) -> Optional[pulumi.Input['AssociationS3OutputLocationArgs']]:
        return pulumi.get(self, "s3_location")

    @s3_location.setter
    def s3_location(self, value: Optional[pulumi.Input['AssociationS3OutputLocationArgs']]):
        pulumi.set(self, "s3_location", value)


@pulumi.input_type
class AssociationS3OutputLocationArgs:
    def __init__(__self__, *,
                 output_s3_bucket_name: Optional[pulumi.Input[str]] = None,
                 output_s3_key_prefix: Optional[pulumi.Input[str]] = None,
                 output_s3_region: Optional[pulumi.Input[str]] = None):
        if output_s3_bucket_name is not None:
            pulumi.set(__self__, "output_s3_bucket_name", output_s3_bucket_name)
        if output_s3_key_prefix is not None:
            pulumi.set(__self__, "output_s3_key_prefix", output_s3_key_prefix)
        if output_s3_region is not None:
            pulumi.set(__self__, "output_s3_region", output_s3_region)

    @property
    @pulumi.getter(name="outputS3BucketName")
    def output_s3_bucket_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "output_s3_bucket_name")

    @output_s3_bucket_name.setter
    def output_s3_bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_s3_bucket_name", value)

    @property
    @pulumi.getter(name="outputS3KeyPrefix")
    def output_s3_key_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "output_s3_key_prefix")

    @output_s3_key_prefix.setter
    def output_s3_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_s3_key_prefix", value)

    @property
    @pulumi.getter(name="outputS3Region")
    def output_s3_region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "output_s3_region")

    @output_s3_region.setter
    def output_s3_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_s3_region", value)


@pulumi.input_type
class AssociationTargetArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 values: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class DocumentAttachmentsSourceArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input['DocumentAttachmentsSourceKey']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input['DocumentAttachmentsSourceKey'] key: The key of a key-value pair that identifies the location of an attachment to a document.
        :param pulumi.Input[str] name: The name of the document attachment file.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: The value of a key-value pair that identifies the location of an attachment to a document. The format for Value depends on the type of key you specify.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input['DocumentAttachmentsSourceKey']]:
        """
        The key of a key-value pair that identifies the location of an attachment to a document.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input['DocumentAttachmentsSourceKey']]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the document attachment file.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The value of a key-value pair that identifies the location of an attachment to a document. The format for Value depends on the type of key you specify.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class DocumentRequiresArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of the required SSM document. The name can be an Amazon Resource Name (ARN).
        :param pulumi.Input[str] version: The document version required by the current document.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the required SSM document. The name can be an Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The document version required by the current document.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class PatchBaselinePatchFilterGroupArgs:
    def __init__(__self__, *,
                 patch_filters: Optional[pulumi.Input[Sequence[pulumi.Input['PatchBaselinePatchFilterArgs']]]] = None):
        """
        The patch filter group that defines the criteria for the rule.
        """
        if patch_filters is not None:
            pulumi.set(__self__, "patch_filters", patch_filters)

    @property
    @pulumi.getter(name="patchFilters")
    def patch_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PatchBaselinePatchFilterArgs']]]]:
        return pulumi.get(self, "patch_filters")

    @patch_filters.setter
    def patch_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PatchBaselinePatchFilterArgs']]]]):
        pulumi.set(self, "patch_filters", value)


@pulumi.input_type
class PatchBaselinePatchFilterArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input['PatchBaselinePatchFilterKey']] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Defines which patches should be included in a patch baseline.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input['PatchBaselinePatchFilterKey']]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input['PatchBaselinePatchFilterKey']]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class PatchBaselinePatchSourceArgs:
    def __init__(__self__, *,
                 configuration: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 products: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Information about the patches to use to update the instances, including target operating systems and source repository. Applies to Linux instances only.
        """
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if products is not None:
            pulumi.set(__self__, "products", products)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def products(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "products")

    @products.setter
    def products(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "products", value)


@pulumi.input_type
class PatchBaselineRuleGroupArgs:
    def __init__(__self__, *,
                 patch_rules: Optional[pulumi.Input[Sequence[pulumi.Input['PatchBaselineRuleArgs']]]] = None):
        """
        A set of rules defining the approval rules for a patch baseline.
        """
        if patch_rules is not None:
            pulumi.set(__self__, "patch_rules", patch_rules)

    @property
    @pulumi.getter(name="patchRules")
    def patch_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PatchBaselineRuleArgs']]]]:
        return pulumi.get(self, "patch_rules")

    @patch_rules.setter
    def patch_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PatchBaselineRuleArgs']]]]):
        pulumi.set(self, "patch_rules", value)


@pulumi.input_type
class PatchBaselineRuleArgs:
    def __init__(__self__, *,
                 approve_after_days: Optional[pulumi.Input[int]] = None,
                 approve_until_date: Optional[pulumi.Input[str]] = None,
                 compliance_level: Optional[pulumi.Input['PatchBaselineRuleComplianceLevel']] = None,
                 enable_non_security: Optional[pulumi.Input[bool]] = None,
                 patch_filter_group: Optional[pulumi.Input['PatchBaselinePatchFilterGroupArgs']] = None):
        """
        Defines an approval rule for a patch baseline.
        """
        if approve_after_days is not None:
            pulumi.set(__self__, "approve_after_days", approve_after_days)
        if approve_until_date is not None:
            pulumi.set(__self__, "approve_until_date", approve_until_date)
        if compliance_level is not None:
            pulumi.set(__self__, "compliance_level", compliance_level)
        if enable_non_security is not None:
            pulumi.set(__self__, "enable_non_security", enable_non_security)
        if patch_filter_group is not None:
            pulumi.set(__self__, "patch_filter_group", patch_filter_group)

    @property
    @pulumi.getter(name="approveAfterDays")
    def approve_after_days(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "approve_after_days")

    @approve_after_days.setter
    def approve_after_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "approve_after_days", value)

    @property
    @pulumi.getter(name="approveUntilDate")
    def approve_until_date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "approve_until_date")

    @approve_until_date.setter
    def approve_until_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "approve_until_date", value)

    @property
    @pulumi.getter(name="complianceLevel")
    def compliance_level(self) -> Optional[pulumi.Input['PatchBaselineRuleComplianceLevel']]:
        return pulumi.get(self, "compliance_level")

    @compliance_level.setter
    def compliance_level(self, value: Optional[pulumi.Input['PatchBaselineRuleComplianceLevel']]):
        pulumi.set(self, "compliance_level", value)

    @property
    @pulumi.getter(name="enableNonSecurity")
    def enable_non_security(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_non_security")

    @enable_non_security.setter
    def enable_non_security(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_non_security", value)

    @property
    @pulumi.getter(name="patchFilterGroup")
    def patch_filter_group(self) -> Optional[pulumi.Input['PatchBaselinePatchFilterGroupArgs']]:
        return pulumi.get(self, "patch_filter_group")

    @patch_filter_group.setter
    def patch_filter_group(self, value: Optional[pulumi.Input['PatchBaselinePatchFilterGroupArgs']]):
        pulumi.set(self, "patch_filter_group", value)


@pulumi.input_type
class ResourceDataSyncAwsOrganizationsSourceArgs:
    def __init__(__self__, *,
                 organization_source_type: pulumi.Input[str],
                 organizational_units: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "organization_source_type", organization_source_type)
        if organizational_units is not None:
            pulumi.set(__self__, "organizational_units", organizational_units)

    @property
    @pulumi.getter(name="organizationSourceType")
    def organization_source_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_source_type")

    @organization_source_type.setter
    def organization_source_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_source_type", value)

    @property
    @pulumi.getter(name="organizationalUnits")
    def organizational_units(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "organizational_units")

    @organizational_units.setter
    def organizational_units(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "organizational_units", value)


@pulumi.input_type
class ResourceDataSyncS3DestinationArgs:
    def __init__(__self__, *,
                 bucket_name: pulumi.Input[str],
                 bucket_region: pulumi.Input[str],
                 sync_format: pulumi.Input[str],
                 bucket_prefix: Optional[pulumi.Input[str]] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "bucket_name", bucket_name)
        pulumi.set(__self__, "bucket_region", bucket_region)
        pulumi.set(__self__, "sync_format", sync_format)
        if bucket_prefix is not None:
            pulumi.set(__self__, "bucket_prefix", bucket_prefix)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="bucketRegion")
    def bucket_region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bucket_region")

    @bucket_region.setter
    def bucket_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_region", value)

    @property
    @pulumi.getter(name="syncFormat")
    def sync_format(self) -> pulumi.Input[str]:
        return pulumi.get(self, "sync_format")

    @sync_format.setter
    def sync_format(self, value: pulumi.Input[str]):
        pulumi.set(self, "sync_format", value)

    @property
    @pulumi.getter(name="bucketPrefix")
    def bucket_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bucket_prefix")

    @bucket_prefix.setter
    def bucket_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_prefix", value)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_arn", value)


@pulumi.input_type
class ResourceDataSyncSyncSourceArgs:
    def __init__(__self__, *,
                 source_regions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 source_type: pulumi.Input[str],
                 aws_organizations_source: Optional[pulumi.Input['ResourceDataSyncAwsOrganizationsSourceArgs']] = None,
                 include_future_regions: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "source_regions", source_regions)
        pulumi.set(__self__, "source_type", source_type)
        if aws_organizations_source is not None:
            pulumi.set(__self__, "aws_organizations_source", aws_organizations_source)
        if include_future_regions is not None:
            pulumi.set(__self__, "include_future_regions", include_future_regions)

    @property
    @pulumi.getter(name="sourceRegions")
    def source_regions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "source_regions")

    @source_regions.setter
    def source_regions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "source_regions", value)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source_type")

    @source_type.setter
    def source_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_type", value)

    @property
    @pulumi.getter(name="awsOrganizationsSource")
    def aws_organizations_source(self) -> Optional[pulumi.Input['ResourceDataSyncAwsOrganizationsSourceArgs']]:
        return pulumi.get(self, "aws_organizations_source")

    @aws_organizations_source.setter
    def aws_organizations_source(self, value: Optional[pulumi.Input['ResourceDataSyncAwsOrganizationsSourceArgs']]):
        pulumi.set(self, "aws_organizations_source", value)

    @property
    @pulumi.getter(name="includeFutureRegions")
    def include_future_regions(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_future_regions")

    @include_future_regions.setter
    def include_future_regions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_future_regions", value)


