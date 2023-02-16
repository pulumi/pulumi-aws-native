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
    'GetStackSetResult',
    'AwaitableGetStackSetResult',
    'get_stack_set',
    'get_stack_set_output',
]

@pulumi.output_type
class GetStackSetResult:
    def __init__(__self__, administration_role_arn=None, auto_deployment=None, capabilities=None, description=None, execution_role_name=None, managed_execution=None, parameters=None, stack_instances_group=None, stack_set_id=None, tags=None, template_body=None):
        if administration_role_arn and not isinstance(administration_role_arn, str):
            raise TypeError("Expected argument 'administration_role_arn' to be a str")
        pulumi.set(__self__, "administration_role_arn", administration_role_arn)
        if auto_deployment and not isinstance(auto_deployment, dict):
            raise TypeError("Expected argument 'auto_deployment' to be a dict")
        pulumi.set(__self__, "auto_deployment", auto_deployment)
        if capabilities and not isinstance(capabilities, list):
            raise TypeError("Expected argument 'capabilities' to be a list")
        pulumi.set(__self__, "capabilities", capabilities)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if execution_role_name and not isinstance(execution_role_name, str):
            raise TypeError("Expected argument 'execution_role_name' to be a str")
        pulumi.set(__self__, "execution_role_name", execution_role_name)
        if managed_execution and not isinstance(managed_execution, dict):
            raise TypeError("Expected argument 'managed_execution' to be a dict")
        pulumi.set(__self__, "managed_execution", managed_execution)
        if parameters and not isinstance(parameters, list):
            raise TypeError("Expected argument 'parameters' to be a list")
        pulumi.set(__self__, "parameters", parameters)
        if stack_instances_group and not isinstance(stack_instances_group, list):
            raise TypeError("Expected argument 'stack_instances_group' to be a list")
        pulumi.set(__self__, "stack_instances_group", stack_instances_group)
        if stack_set_id and not isinstance(stack_set_id, str):
            raise TypeError("Expected argument 'stack_set_id' to be a str")
        pulumi.set(__self__, "stack_set_id", stack_set_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if template_body and not isinstance(template_body, str):
            raise TypeError("Expected argument 'template_body' to be a str")
        pulumi.set(__self__, "template_body", template_body)

    @property
    @pulumi.getter(name="administrationRoleARN")
    def administration_role_arn(self) -> Optional[str]:
        """
        The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.
        """
        return pulumi.get(self, "administration_role_arn")

    @property
    @pulumi.getter(name="autoDeployment")
    def auto_deployment(self) -> Optional['outputs.StackSetAutoDeployment']:
        """
        Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED.
        """
        return pulumi.get(self, "auto_deployment")

    @property
    @pulumi.getter
    def capabilities(self) -> Optional[Sequence['StackSetCapability']]:
        """
        In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.
        """
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description of the stack set. You can use the description to identify the stack set's purpose or other important information.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="executionRoleName")
    def execution_role_name(self) -> Optional[str]:
        """
        The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.
        """
        return pulumi.get(self, "execution_role_name")

    @property
    @pulumi.getter(name="managedExecution")
    def managed_execution(self) -> Optional['outputs.ManagedExecutionProperties']:
        """
        Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
        """
        return pulumi.get(self, "managed_execution")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Sequence['outputs.StackSetParameter']]:
        """
        The input parameters for the stack set template.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="stackInstancesGroup")
    def stack_instances_group(self) -> Optional[Sequence['outputs.StackSetStackInstances']]:
        """
        A group of stack instances with parameters in some specific accounts and regions.
        """
        return pulumi.get(self, "stack_instances_group")

    @property
    @pulumi.getter(name="stackSetId")
    def stack_set_id(self) -> Optional[str]:
        """
        The ID of the stack set that you're creating.
        """
        return pulumi.get(self, "stack_set_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.StackSetTag']]:
        """
        The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateBody")
    def template_body(self) -> Optional[str]:
        """
        The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
        """
        return pulumi.get(self, "template_body")


class AwaitableGetStackSetResult(GetStackSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStackSetResult(
            administration_role_arn=self.administration_role_arn,
            auto_deployment=self.auto_deployment,
            capabilities=self.capabilities,
            description=self.description,
            execution_role_name=self.execution_role_name,
            managed_execution=self.managed_execution,
            parameters=self.parameters,
            stack_instances_group=self.stack_instances_group,
            stack_set_id=self.stack_set_id,
            tags=self.tags,
            template_body=self.template_body)


def get_stack_set(stack_set_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStackSetResult:
    """
    StackSet as a resource provides one-click experience for provisioning a StackSet and StackInstances


    :param str stack_set_id: The ID of the stack set that you're creating.
    """
    __args__ = dict()
    __args__['stackSetId'] = stack_set_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cloudformation:getStackSet', __args__, opts=opts, typ=GetStackSetResult).value

    return AwaitableGetStackSetResult(
        administration_role_arn=__ret__.administration_role_arn,
        auto_deployment=__ret__.auto_deployment,
        capabilities=__ret__.capabilities,
        description=__ret__.description,
        execution_role_name=__ret__.execution_role_name,
        managed_execution=__ret__.managed_execution,
        parameters=__ret__.parameters,
        stack_instances_group=__ret__.stack_instances_group,
        stack_set_id=__ret__.stack_set_id,
        tags=__ret__.tags,
        template_body=__ret__.template_body)


@_utilities.lift_output_func(get_stack_set)
def get_stack_set_output(stack_set_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStackSetResult]:
    """
    StackSet as a resource provides one-click experience for provisioning a StackSet and StackInstances


    :param str stack_set_id: The ID of the stack set that you're creating.
    """
    ...