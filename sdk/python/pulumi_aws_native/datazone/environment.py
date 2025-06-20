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
from ._enums import *
from ._inputs import *

__all__ = ['EnvironmentArgs', 'Environment']

@pulumi.input_type
class EnvironmentArgs:
    def __init__(__self__, *,
                 domain_identifier: pulumi.Input[builtins.str],
                 project_identifier: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 environment_account_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_account_region: Optional[pulumi.Input[builtins.str]] = None,
                 environment_profile_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_role_arn: Optional[pulumi.Input[builtins.str]] = None,
                 glossary_terms: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 user_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentParameterArgs']]]] = None):
        """
        The set of arguments for constructing a Environment resource.
        :param pulumi.Input[builtins.str] domain_identifier: The identifier of the Amazon DataZone domain in which the environment would be created.
        :param pulumi.Input[builtins.str] project_identifier: The ID of the Amazon DataZone project in which the environment would be created.
        :param pulumi.Input[builtins.str] description: The description of the Amazon DataZone environment.
        :param pulumi.Input[builtins.str] environment_account_identifier: The AWS account in which the Amazon DataZone environment is created.
        :param pulumi.Input[builtins.str] environment_account_region: The AWS region in which the Amazon DataZone environment is created.
        :param pulumi.Input[builtins.str] environment_profile_identifier: The ID of the environment profile with which the Amazon DataZone environment would be created.
        :param pulumi.Input[builtins.str] environment_role_arn: Environment role arn for custom aws environment permissions
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] glossary_terms: The glossary terms that can be used in the Amazon DataZone environment.
        :param pulumi.Input[builtins.str] name: The name of the environment.
        :param pulumi.Input[Sequence[pulumi.Input['EnvironmentParameterArgs']]] user_parameters: The user parameters of the Amazon DataZone environment.
        """
        pulumi.set(__self__, "domain_identifier", domain_identifier)
        pulumi.set(__self__, "project_identifier", project_identifier)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environment_account_identifier is not None:
            pulumi.set(__self__, "environment_account_identifier", environment_account_identifier)
        if environment_account_region is not None:
            pulumi.set(__self__, "environment_account_region", environment_account_region)
        if environment_profile_identifier is not None:
            pulumi.set(__self__, "environment_profile_identifier", environment_profile_identifier)
        if environment_role_arn is not None:
            pulumi.set(__self__, "environment_role_arn", environment_role_arn)
        if glossary_terms is not None:
            pulumi.set(__self__, "glossary_terms", glossary_terms)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if user_parameters is not None:
            pulumi.set(__self__, "user_parameters", user_parameters)

    @property
    @pulumi.getter(name="domainIdentifier")
    def domain_identifier(self) -> pulumi.Input[builtins.str]:
        """
        The identifier of the Amazon DataZone domain in which the environment would be created.
        """
        return pulumi.get(self, "domain_identifier")

    @domain_identifier.setter
    def domain_identifier(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "domain_identifier", value)

    @property
    @pulumi.getter(name="projectIdentifier")
    def project_identifier(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the Amazon DataZone project in which the environment would be created.
        """
        return pulumi.get(self, "project_identifier")

    @project_identifier.setter
    def project_identifier(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_identifier", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The description of the Amazon DataZone environment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="environmentAccountIdentifier")
    def environment_account_identifier(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The AWS account in which the Amazon DataZone environment is created.
        """
        return pulumi.get(self, "environment_account_identifier")

    @environment_account_identifier.setter
    def environment_account_identifier(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "environment_account_identifier", value)

    @property
    @pulumi.getter(name="environmentAccountRegion")
    def environment_account_region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The AWS region in which the Amazon DataZone environment is created.
        """
        return pulumi.get(self, "environment_account_region")

    @environment_account_region.setter
    def environment_account_region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "environment_account_region", value)

    @property
    @pulumi.getter(name="environmentProfileIdentifier")
    def environment_profile_identifier(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the environment profile with which the Amazon DataZone environment would be created.
        """
        return pulumi.get(self, "environment_profile_identifier")

    @environment_profile_identifier.setter
    def environment_profile_identifier(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "environment_profile_identifier", value)

    @property
    @pulumi.getter(name="environmentRoleArn")
    def environment_role_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Environment role arn for custom aws environment permissions
        """
        return pulumi.get(self, "environment_role_arn")

    @environment_role_arn.setter
    def environment_role_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "environment_role_arn", value)

    @property
    @pulumi.getter(name="glossaryTerms")
    def glossary_terms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The glossary terms that can be used in the Amazon DataZone environment.
        """
        return pulumi.get(self, "glossary_terms")

    @glossary_terms.setter
    def glossary_terms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "glossary_terms", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="userParameters")
    def user_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentParameterArgs']]]]:
        """
        The user parameters of the Amazon DataZone environment.
        """
        return pulumi.get(self, "user_parameters")

    @user_parameters.setter
    def user_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentParameterArgs']]]]):
        pulumi.set(self, "user_parameters", value)


@pulumi.type_token("aws-native:datazone:Environment")
class Environment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_account_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_account_region: Optional[pulumi.Input[builtins.str]] = None,
                 environment_profile_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_role_arn: Optional[pulumi.Input[builtins.str]] = None,
                 glossary_terms: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 user_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EnvironmentParameterArgs', 'EnvironmentParameterArgsDict']]]]] = None,
                 __props__=None):
        """
        Definition of AWS::DataZone::Environment Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The description of the Amazon DataZone environment.
        :param pulumi.Input[builtins.str] domain_identifier: The identifier of the Amazon DataZone domain in which the environment would be created.
        :param pulumi.Input[builtins.str] environment_account_identifier: The AWS account in which the Amazon DataZone environment is created.
        :param pulumi.Input[builtins.str] environment_account_region: The AWS region in which the Amazon DataZone environment is created.
        :param pulumi.Input[builtins.str] environment_profile_identifier: The ID of the environment profile with which the Amazon DataZone environment would be created.
        :param pulumi.Input[builtins.str] environment_role_arn: Environment role arn for custom aws environment permissions
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] glossary_terms: The glossary terms that can be used in the Amazon DataZone environment.
        :param pulumi.Input[builtins.str] name: The name of the environment.
        :param pulumi.Input[builtins.str] project_identifier: The ID of the Amazon DataZone project in which the environment would be created.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EnvironmentParameterArgs', 'EnvironmentParameterArgsDict']]]] user_parameters: The user parameters of the Amazon DataZone environment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::DataZone::Environment Resource Type

        :param str resource_name: The name of the resource.
        :param EnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_account_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_account_region: Optional[pulumi.Input[builtins.str]] = None,
                 environment_profile_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_role_arn: Optional[pulumi.Input[builtins.str]] = None,
                 glossary_terms: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 user_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EnvironmentParameterArgs', 'EnvironmentParameterArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

            __props__.__dict__["description"] = description
            if domain_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'domain_identifier'")
            __props__.__dict__["domain_identifier"] = domain_identifier
            __props__.__dict__["environment_account_identifier"] = environment_account_identifier
            __props__.__dict__["environment_account_region"] = environment_account_region
            __props__.__dict__["environment_profile_identifier"] = environment_profile_identifier
            __props__.__dict__["environment_role_arn"] = environment_role_arn
            __props__.__dict__["glossary_terms"] = glossary_terms
            __props__.__dict__["name"] = name
            if project_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'project_identifier'")
            __props__.__dict__["project_identifier"] = project_identifier
            __props__.__dict__["user_parameters"] = user_parameters
            __props__.__dict__["aws_account_id"] = None
            __props__.__dict__["aws_account_region"] = None
            __props__.__dict__["aws_id"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["created_by"] = None
            __props__.__dict__["domain_id"] = None
            __props__.__dict__["environment_blueprint_id"] = None
            __props__.__dict__["environment_profile_id"] = None
            __props__.__dict__["project_id"] = None
            __props__.__dict__["provider"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["domainIdentifier", "environmentAccountIdentifier", "environmentAccountRegion", "environmentProfileIdentifier", "projectIdentifier", "userParameters[*]"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Environment, __self__).__init__(
            'aws-native:datazone:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Environment':
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

        __props__.__dict__["aws_account_id"] = None
        __props__.__dict__["aws_account_region"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["created_by"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["domain_id"] = None
        __props__.__dict__["domain_identifier"] = None
        __props__.__dict__["environment_account_identifier"] = None
        __props__.__dict__["environment_account_region"] = None
        __props__.__dict__["environment_blueprint_id"] = None
        __props__.__dict__["environment_profile_id"] = None
        __props__.__dict__["environment_profile_identifier"] = None
        __props__.__dict__["environment_role_arn"] = None
        __props__.__dict__["glossary_terms"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project_id"] = None
        __props__.__dict__["project_identifier"] = None
        __props__.__dict__["provider"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["updated_at"] = None
        __props__.__dict__["user_parameters"] = None
        return Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[builtins.str]:
        """
        The AWS account in which the Amazon DataZone environment is created.
        """
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="awsAccountRegion")
    def aws_account_region(self) -> pulumi.Output[builtins.str]:
        """
        The AWS region in which the Amazon DataZone environment is created.
        """
        return pulumi.get(self, "aws_account_region")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the Amazon DataZone environment.
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The timestamp of when the environment was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon DataZone user who created the environment.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The description of the Amazon DataZone environment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Output[builtins.str]:
        """
        The identifier of the Amazon DataZone domain in which the environment is created.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="domainIdentifier")
    def domain_identifier(self) -> pulumi.Output[builtins.str]:
        """
        The identifier of the Amazon DataZone domain in which the environment would be created.
        """
        return pulumi.get(self, "domain_identifier")

    @property
    @pulumi.getter(name="environmentAccountIdentifier")
    def environment_account_identifier(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The AWS account in which the Amazon DataZone environment is created.
        """
        return pulumi.get(self, "environment_account_identifier")

    @property
    @pulumi.getter(name="environmentAccountRegion")
    def environment_account_region(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The AWS region in which the Amazon DataZone environment is created.
        """
        return pulumi.get(self, "environment_account_region")

    @property
    @pulumi.getter(name="environmentBlueprintId")
    def environment_blueprint_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the blueprint with which the Amazon DataZone environment was created.
        """
        return pulumi.get(self, "environment_blueprint_id")

    @property
    @pulumi.getter(name="environmentProfileId")
    def environment_profile_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the environment profile with which the Amazon DataZone environment was created.
        """
        return pulumi.get(self, "environment_profile_id")

    @property
    @pulumi.getter(name="environmentProfileIdentifier")
    def environment_profile_identifier(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The ID of the environment profile with which the Amazon DataZone environment would be created.
        """
        return pulumi.get(self, "environment_profile_identifier")

    @property
    @pulumi.getter(name="environmentRoleArn")
    def environment_role_arn(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Environment role arn for custom aws environment permissions
        """
        return pulumi.get(self, "environment_role_arn")

    @property
    @pulumi.getter(name="glossaryTerms")
    def glossary_terms(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The glossary terms that can be used in the Amazon DataZone environment.
        """
        return pulumi.get(self, "glossary_terms")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the Amazon DataZone project in which the environment is created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectIdentifier")
    def project_identifier(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the Amazon DataZone project in which the environment would be created.
        """
        return pulumi.get(self, "project_identifier")

    @property
    @pulumi.getter
    def provider(self) -> pulumi.Output[builtins.str]:
        """
        The provider of the Amazon DataZone environment.
        """
        return pulumi.get(self, "provider")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['EnvironmentStatus']:
        """
        The status of the Amazon DataZone environment.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        The timestamp of when the environment was updated.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="userParameters")
    def user_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.EnvironmentParameter']]]:
        """
        The user parameters of the Amazon DataZone environment.
        """
        return pulumi.get(self, "user_parameters")

