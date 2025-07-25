# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'CloudAutonomousVmClusterComputeModel',
    'CloudAutonomousVmClusterLicenseModel',
    'CloudAutonomousVmClusterMaintenanceWindowDaysOfWeekItem',
    'CloudAutonomousVmClusterMaintenanceWindowMonthsItem',
    'CloudAutonomousVmClusterMaintenanceWindowPreference',
    'CloudVmClusterLicenseModel',
]


@pulumi.type_token("aws-native:odb:CloudAutonomousVmClusterComputeModel")
class CloudAutonomousVmClusterComputeModel(builtins.str, Enum):
    """
    The compute model of the Autonomous VM cluster: ECPU or OCPU.
    """
    ECPU = "ECPU"
    OCPU = "OCPU"


@pulumi.type_token("aws-native:odb:CloudAutonomousVmClusterLicenseModel")
class CloudAutonomousVmClusterLicenseModel(builtins.str, Enum):
    """
    The Oracle license model that applies to the Autonomous VM cluster. Valid values are LICENSE_INCLUDED or BRING_YOUR_OWN_LICENSE.
    """
    BRING_YOUR_OWN_LICENSE = "BRING_YOUR_OWN_LICENSE"
    LICENSE_INCLUDED = "LICENSE_INCLUDED"


@pulumi.type_token("aws-native:odb:CloudAutonomousVmClusterMaintenanceWindowDaysOfWeekItem")
class CloudAutonomousVmClusterMaintenanceWindowDaysOfWeekItem(builtins.str, Enum):
    MONDAY = "MONDAY"
    TUESDAY = "TUESDAY"
    WEDNESDAY = "WEDNESDAY"
    THURSDAY = "THURSDAY"
    FRIDAY = "FRIDAY"
    SATURDAY = "SATURDAY"
    SUNDAY = "SUNDAY"


@pulumi.type_token("aws-native:odb:CloudAutonomousVmClusterMaintenanceWindowMonthsItem")
class CloudAutonomousVmClusterMaintenanceWindowMonthsItem(builtins.str, Enum):
    JANUARY = "JANUARY"
    FEBRUARY = "FEBRUARY"
    MARCH = "MARCH"
    APRIL = "APRIL"
    MAY = "MAY"
    JUNE = "JUNE"
    JULY = "JULY"
    AUGUST = "AUGUST"
    SEPTEMBER = "SEPTEMBER"
    OCTOBER = "OCTOBER"
    NOVEMBER = "NOVEMBER"
    DECEMBER = "DECEMBER"


@pulumi.type_token("aws-native:odb:CloudAutonomousVmClusterMaintenanceWindowPreference")
class CloudAutonomousVmClusterMaintenanceWindowPreference(builtins.str, Enum):
    """
    The preference for the maintenance window scheduling.
    """
    NO_PREFERENCE = "NO_PREFERENCE"
    CUSTOM_PREFERENCE = "CUSTOM_PREFERENCE"


@pulumi.type_token("aws-native:odb:CloudVmClusterLicenseModel")
class CloudVmClusterLicenseModel(builtins.str, Enum):
    """
    The Oracle license model applied to the VM cluster.
    """
    BRING_YOUR_OWN_LICENSE = "BRING_YOUR_OWN_LICENSE"
    LICENSE_INCLUDED = "LICENSE_INCLUDED"
