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
    'GetWorkspaceResult',
    'AwaitableGetWorkspaceResult',
    'get_workspace',
    'get_workspace_output',
]

@pulumi.output_type
class GetWorkspaceResult:
    def __init__(__self__, account_access_type=None, authentication_providers=None, creation_timestamp=None, data_sources=None, description=None, endpoint=None, grafana_version=None, id=None, modification_timestamp=None, name=None, network_access_control=None, notification_destinations=None, organization_role_name=None, organizational_units=None, permission_type=None, role_arn=None, saml_configuration=None, saml_configuration_status=None, sso_client_id=None, stack_set_name=None, status=None, vpc_configuration=None):
        if account_access_type and not isinstance(account_access_type, str):
            raise TypeError("Expected argument 'account_access_type' to be a str")
        pulumi.set(__self__, "account_access_type", account_access_type)
        if authentication_providers and not isinstance(authentication_providers, list):
            raise TypeError("Expected argument 'authentication_providers' to be a list")
        pulumi.set(__self__, "authentication_providers", authentication_providers)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if data_sources and not isinstance(data_sources, list):
            raise TypeError("Expected argument 'data_sources' to be a list")
        pulumi.set(__self__, "data_sources", data_sources)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if grafana_version and not isinstance(grafana_version, str):
            raise TypeError("Expected argument 'grafana_version' to be a str")
        pulumi.set(__self__, "grafana_version", grafana_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if modification_timestamp and not isinstance(modification_timestamp, str):
            raise TypeError("Expected argument 'modification_timestamp' to be a str")
        pulumi.set(__self__, "modification_timestamp", modification_timestamp)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_access_control and not isinstance(network_access_control, dict):
            raise TypeError("Expected argument 'network_access_control' to be a dict")
        pulumi.set(__self__, "network_access_control", network_access_control)
        if notification_destinations and not isinstance(notification_destinations, list):
            raise TypeError("Expected argument 'notification_destinations' to be a list")
        pulumi.set(__self__, "notification_destinations", notification_destinations)
        if organization_role_name and not isinstance(organization_role_name, str):
            raise TypeError("Expected argument 'organization_role_name' to be a str")
        pulumi.set(__self__, "organization_role_name", organization_role_name)
        if organizational_units and not isinstance(organizational_units, list):
            raise TypeError("Expected argument 'organizational_units' to be a list")
        pulumi.set(__self__, "organizational_units", organizational_units)
        if permission_type and not isinstance(permission_type, str):
            raise TypeError("Expected argument 'permission_type' to be a str")
        pulumi.set(__self__, "permission_type", permission_type)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if saml_configuration and not isinstance(saml_configuration, dict):
            raise TypeError("Expected argument 'saml_configuration' to be a dict")
        pulumi.set(__self__, "saml_configuration", saml_configuration)
        if saml_configuration_status and not isinstance(saml_configuration_status, str):
            raise TypeError("Expected argument 'saml_configuration_status' to be a str")
        pulumi.set(__self__, "saml_configuration_status", saml_configuration_status)
        if sso_client_id and not isinstance(sso_client_id, str):
            raise TypeError("Expected argument 'sso_client_id' to be a str")
        pulumi.set(__self__, "sso_client_id", sso_client_id)
        if stack_set_name and not isinstance(stack_set_name, str):
            raise TypeError("Expected argument 'stack_set_name' to be a str")
        pulumi.set(__self__, "stack_set_name", stack_set_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vpc_configuration and not isinstance(vpc_configuration, dict):
            raise TypeError("Expected argument 'vpc_configuration' to be a dict")
        pulumi.set(__self__, "vpc_configuration", vpc_configuration)

    @property
    @pulumi.getter(name="accountAccessType")
    def account_access_type(self) -> Optional['WorkspaceAccountAccessType']:
        return pulumi.get(self, "account_access_type")

    @property
    @pulumi.getter(name="authenticationProviders")
    def authentication_providers(self) -> Optional[Sequence['WorkspaceAuthenticationProviderTypes']]:
        """
        List of authentication providers to enable.
        """
        return pulumi.get(self, "authentication_providers")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[str]:
        """
        Timestamp when the workspace was created.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="dataSources")
    def data_sources(self) -> Optional[Sequence['WorkspaceDataSourceType']]:
        """
        List of data sources on the service managed IAM role.
        """
        return pulumi.get(self, "data_sources")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of a workspace.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[str]:
        """
        Endpoint for the Grafana workspace.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="grafanaVersion")
    def grafana_version(self) -> Optional[str]:
        """
        Version of Grafana the workspace is currently using.
        """
        return pulumi.get(self, "grafana_version")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The id that uniquely identifies a Grafana workspace.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="modificationTimestamp")
    def modification_timestamp(self) -> Optional[str]:
        """
        Timestamp when the workspace was last modified
        """
        return pulumi.get(self, "modification_timestamp")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The user friendly name of a workspace.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkAccessControl")
    def network_access_control(self) -> Optional['outputs.WorkspaceNetworkAccessControl']:
        return pulumi.get(self, "network_access_control")

    @property
    @pulumi.getter(name="notificationDestinations")
    def notification_destinations(self) -> Optional[Sequence['WorkspaceNotificationDestinationType']]:
        """
        List of notification destinations on the customers service managed IAM role that the Grafana workspace can query.
        """
        return pulumi.get(self, "notification_destinations")

    @property
    @pulumi.getter(name="organizationRoleName")
    def organization_role_name(self) -> Optional[str]:
        """
        The name of an IAM role that already exists to use with AWS Organizations to access AWS data sources and notification channels in other accounts in an organization.
        """
        return pulumi.get(self, "organization_role_name")

    @property
    @pulumi.getter(name="organizationalUnits")
    def organizational_units(self) -> Optional[Sequence[str]]:
        """
        List of Organizational Units containing AWS accounts the Grafana workspace can pull data from.
        """
        return pulumi.get(self, "organizational_units")

    @property
    @pulumi.getter(name="permissionType")
    def permission_type(self) -> Optional['WorkspacePermissionType']:
        return pulumi.get(self, "permission_type")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        """
        IAM Role that will be used to grant the Grafana workspace access to a customers AWS resources.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="samlConfiguration")
    def saml_configuration(self) -> Optional['outputs.WorkspaceSamlConfiguration']:
        return pulumi.get(self, "saml_configuration")

    @property
    @pulumi.getter(name="samlConfigurationStatus")
    def saml_configuration_status(self) -> Optional['WorkspaceSamlConfigurationStatus']:
        return pulumi.get(self, "saml_configuration_status")

    @property
    @pulumi.getter(name="ssoClientId")
    def sso_client_id(self) -> Optional[str]:
        """
        The client ID of the AWS SSO Managed Application.
        """
        return pulumi.get(self, "sso_client_id")

    @property
    @pulumi.getter(name="stackSetName")
    def stack_set_name(self) -> Optional[str]:
        """
        The name of the AWS CloudFormation stack set to use to generate IAM roles to be used for this workspace.
        """
        return pulumi.get(self, "stack_set_name")

    @property
    @pulumi.getter
    def status(self) -> Optional['WorkspaceStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcConfiguration")
    def vpc_configuration(self) -> Optional['outputs.WorkspaceVpcConfiguration']:
        return pulumi.get(self, "vpc_configuration")


class AwaitableGetWorkspaceResult(GetWorkspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkspaceResult(
            account_access_type=self.account_access_type,
            authentication_providers=self.authentication_providers,
            creation_timestamp=self.creation_timestamp,
            data_sources=self.data_sources,
            description=self.description,
            endpoint=self.endpoint,
            grafana_version=self.grafana_version,
            id=self.id,
            modification_timestamp=self.modification_timestamp,
            name=self.name,
            network_access_control=self.network_access_control,
            notification_destinations=self.notification_destinations,
            organization_role_name=self.organization_role_name,
            organizational_units=self.organizational_units,
            permission_type=self.permission_type,
            role_arn=self.role_arn,
            saml_configuration=self.saml_configuration,
            saml_configuration_status=self.saml_configuration_status,
            sso_client_id=self.sso_client_id,
            stack_set_name=self.stack_set_name,
            status=self.status,
            vpc_configuration=self.vpc_configuration)


def get_workspace(id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkspaceResult:
    """
    Definition of AWS::Grafana::Workspace Resource Type


    :param str id: The id that uniquely identifies a Grafana workspace.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:grafana:getWorkspace', __args__, opts=opts, typ=GetWorkspaceResult).value

    return AwaitableGetWorkspaceResult(
        account_access_type=__ret__.account_access_type,
        authentication_providers=__ret__.authentication_providers,
        creation_timestamp=__ret__.creation_timestamp,
        data_sources=__ret__.data_sources,
        description=__ret__.description,
        endpoint=__ret__.endpoint,
        grafana_version=__ret__.grafana_version,
        id=__ret__.id,
        modification_timestamp=__ret__.modification_timestamp,
        name=__ret__.name,
        network_access_control=__ret__.network_access_control,
        notification_destinations=__ret__.notification_destinations,
        organization_role_name=__ret__.organization_role_name,
        organizational_units=__ret__.organizational_units,
        permission_type=__ret__.permission_type,
        role_arn=__ret__.role_arn,
        saml_configuration=__ret__.saml_configuration,
        saml_configuration_status=__ret__.saml_configuration_status,
        sso_client_id=__ret__.sso_client_id,
        stack_set_name=__ret__.stack_set_name,
        status=__ret__.status,
        vpc_configuration=__ret__.vpc_configuration)


@_utilities.lift_output_func(get_workspace)
def get_workspace_output(id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkspaceResult]:
    """
    Definition of AWS::Grafana::Workspace Resource Type


    :param str id: The id that uniquely identifies a Grafana workspace.
    """
    ...