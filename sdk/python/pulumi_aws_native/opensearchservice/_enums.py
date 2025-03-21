# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ApplicationAppConfigType',
    'DomainNodeOptionNodeType',
    'DomainRolesKeyIdcType',
    'DomainSubjectKeyIdcType',
]


class ApplicationAppConfigType(str, Enum):
    """
    AppConfig type values.
    """
    OPENSEARCH_DASHBOARDS_DASHBOARD_ADMIN_USERS = "opensearchDashboards.dashboardAdmin.users"
    OPENSEARCH_DASHBOARDS_DASHBOARD_ADMIN_GROUPS = "opensearchDashboards.dashboardAdmin.groups"


class DomainNodeOptionNodeType(str, Enum):
    """
    Container for node type like coordinating.
    """
    COORDINATOR = "coordinator"


class DomainRolesKeyIdcType(str, Enum):
    """
    Roles Key Idc type values.
    """
    GROUP_NAME = "GroupName"
    GROUP_ID = "GroupId"


class DomainSubjectKeyIdcType(str, Enum):
    """
    Subject Key Idc type values.
    """
    USER_NAME = "UserName"
    USER_ID = "UserId"
    EMAIL = "Email"
