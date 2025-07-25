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
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetStackResult',
    'AwaitableGetStackResult',
    'get_stack',
    'get_stack_output',
]

@pulumi.output_type
class GetStackResult:
    def __init__(__self__, capabilities=None, change_set_id=None, creation_time=None, description=None, disable_rollback=None, enable_termination_protection=None, last_update_time=None, notification_arns=None, outputs=None, parameters=None, parent_id=None, role_arn=None, root_id=None, stack_id=None, stack_policy_body=None, stack_status=None, stack_status_reason=None, tags=None, template_body=None, timeout_in_minutes=None):
        if capabilities and not isinstance(capabilities, list):
            raise TypeError("Expected argument 'capabilities' to be a list")
        pulumi.set(__self__, "capabilities", capabilities)
        if change_set_id and not isinstance(change_set_id, str):
            raise TypeError("Expected argument 'change_set_id' to be a str")
        pulumi.set(__self__, "change_set_id", change_set_id)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disable_rollback and not isinstance(disable_rollback, bool):
            raise TypeError("Expected argument 'disable_rollback' to be a bool")
        pulumi.set(__self__, "disable_rollback", disable_rollback)
        if enable_termination_protection and not isinstance(enable_termination_protection, bool):
            raise TypeError("Expected argument 'enable_termination_protection' to be a bool")
        pulumi.set(__self__, "enable_termination_protection", enable_termination_protection)
        if last_update_time and not isinstance(last_update_time, str):
            raise TypeError("Expected argument 'last_update_time' to be a str")
        pulumi.set(__self__, "last_update_time", last_update_time)
        if notification_arns and not isinstance(notification_arns, list):
            raise TypeError("Expected argument 'notification_arns' to be a list")
        pulumi.set(__self__, "notification_arns", notification_arns)
        if outputs and not isinstance(outputs, list):
            raise TypeError("Expected argument 'outputs' to be a list")
        pulumi.set(__self__, "outputs", outputs)
        if parameters and not isinstance(parameters, dict):
            raise TypeError("Expected argument 'parameters' to be a dict")
        pulumi.set(__self__, "parameters", parameters)
        if parent_id and not isinstance(parent_id, str):
            raise TypeError("Expected argument 'parent_id' to be a str")
        pulumi.set(__self__, "parent_id", parent_id)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if root_id and not isinstance(root_id, str):
            raise TypeError("Expected argument 'root_id' to be a str")
        pulumi.set(__self__, "root_id", root_id)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if stack_policy_body and not isinstance(stack_policy_body, dict):
            raise TypeError("Expected argument 'stack_policy_body' to be a dict")
        pulumi.set(__self__, "stack_policy_body", stack_policy_body)
        if stack_status and not isinstance(stack_status, str):
            raise TypeError("Expected argument 'stack_status' to be a str")
        pulumi.set(__self__, "stack_status", stack_status)
        if stack_status_reason and not isinstance(stack_status_reason, str):
            raise TypeError("Expected argument 'stack_status_reason' to be a str")
        pulumi.set(__self__, "stack_status_reason", stack_status_reason)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if template_body and not isinstance(template_body, dict):
            raise TypeError("Expected argument 'template_body' to be a dict")
        pulumi.set(__self__, "template_body", template_body)
        if timeout_in_minutes and not isinstance(timeout_in_minutes, int):
            raise TypeError("Expected argument 'timeout_in_minutes' to be a int")
        pulumi.set(__self__, "timeout_in_minutes", timeout_in_minutes)

    @property
    @pulumi.getter
    def capabilities(self) -> Optional[Sequence['StackCapabilitiesItem']]:
        """
        In some cases, you must explicitly acknowledge that your stack template contains certain capabilities in order for CloudFormation to create the stack.

        - `CAPABILITY_IAM` and `CAPABILITY_NAMED_IAM`

        Some stack templates might include resources that can affect permissions in your AWS account ; for example, by creating new AWS Identity and Access Management (IAM) users. For those stacks, you must explicitly acknowledge this by specifying one of these capabilities.

        The following IAM resources require you to specify either the `CAPABILITY_IAM` or `CAPABILITY_NAMED_IAM` capability.

        - If you have IAM resources, you can specify either capability.
        - If you have IAM resources with custom names, you *must* specify `CAPABILITY_NAMED_IAM` .
        - If you don't specify either of these capabilities, CloudFormation returns an `InsufficientCapabilities` error.

        If your stack template contains these resources, we recommend that you review all permissions associated with them and edit their permissions if necessary.

        - [AWS::IAM::AccessKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-accesskey.html)
        - [AWS::IAM::Group](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-group.html)
        - [AWS::IAM::InstanceProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-instanceprofile.html)
        - [AWS::IAM::Policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-policy.html)
        - [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-role.html)
        - [AWS::IAM::User](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-user.html)
        - [AWS::IAM::UserToGroupAddition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-usertogroupaddition.html)

        For more information, see [Acknowledging IAM resources in CloudFormation templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/control-access-with-iam.html#using-iam-capabilities) in the *AWS CloudFormation User Guide* .
        - `CAPABILITY_AUTO_EXPAND`

        Some template contain macros. Macros perform custom processing on templates; this can include simple actions like find-and-replace operations, all the way to extensive transformations of entire templates. Because of this, users typically create a change set from the processed template, so that they can review the changes resulting from the macros before actually creating the stack. If your stack template contains one or more macros, and you choose to create a stack directly from the processed template, without first reviewing the resulting changes in a change set, you must acknowledge this capability. This includes the [AWS::Include](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-include.html) and [AWS::Serverless](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-serverless.html) transforms, which are macros hosted by CloudFormation .

        If you want to create a stack from a stack template that contains macros *and* nested stacks, you must create the stack directly from the template using this capability.

        > You should only create stacks directly from a stack template that contains macros if you know what processing the macro performs.
        > 
        > Each macro relies on an underlying Lambda service function for processing stack templates. Be aware that the Lambda function owner can update the function operation without CloudFormation being notified. 

        For more information, see [Perform custom processing on CloudFormation templates with template macros](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html) in the *AWS CloudFormation User Guide* .
        """
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter(name="changeSetId")
    def change_set_id(self) -> Optional[builtins.str]:
        """
        The unique ID of the change set.
        """
        return pulumi.get(self, "change_set_id")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[builtins.str]:
        """
        The time at which the stack was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        A user-defined description associated with the stack.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableRollback")
    def disable_rollback(self) -> Optional[builtins.bool]:
        """
        Set to `true` to disable rollback of the stack if stack creation failed. You can specify either `DisableRollback` or `OnFailure` , but not both.

        Default: `false`
        """
        return pulumi.get(self, "disable_rollback")

    @property
    @pulumi.getter(name="enableTerminationProtection")
    def enable_termination_protection(self) -> Optional[builtins.bool]:
        """
        Whether to enable termination protection on the specified stack. If a user attempts to delete a stack with termination protection enabled, the operation fails and the stack remains unchanged. For more information, see [Protect CloudFormation stacks from being deleted](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-protect-stacks.html) in the *AWS CloudFormation User Guide* . Termination protection is deactivated on stacks by default.

        For nested stacks, termination protection is set on the root stack and can't be changed directly on the nested stack.
        """
        return pulumi.get(self, "enable_termination_protection")

    @property
    @pulumi.getter(name="lastUpdateTime")
    def last_update_time(self) -> Optional[builtins.str]:
        """
        The time the stack was last updated. This field will only be returned if the stack has been updated at least once.
        """
        return pulumi.get(self, "last_update_time")

    @property
    @pulumi.getter(name="notificationArns")
    def notification_arns(self) -> Optional[Sequence[builtins.str]]:
        """
        The Amazon SNS topic ARNs to publish stack related events. You can find your Amazon SNS topic ARNs using the Amazon SNS console or your Command Line Interface (CLI).
        """
        return pulumi.get(self, "notification_arns")

    @property
    @pulumi.getter
    def outputs(self) -> Optional[Sequence['outputs.StackOutput']]:
        """
        A list of output structures.
        """
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, builtins.str]]:
        """
        The set value pairs that represent the parameters passed to CloudFormation when this nested stack is created. Each parameter has a name corresponding to a parameter defined in the embedded template and a value representing the value that you want to set for the parameter.

        > If you use the `Ref` function to pass a parameter value to a nested stack, comma-delimited list parameters must be of type `String` . In other words, you can't pass values that are of type `CommaDelimitedList` to nested stacks. 

        Required if the nested stack requires input parameters.

        Whether an update causes interruptions depends on the resources that are being updated. An update never causes a nested stack to be replaced.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[builtins.str]:
        """
        For nested stacks, the stack ID of the direct parent of this stack. For the first level of nested stacks, the root stack is also the parent stack.

        For more information, see [Nested stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html) in the *AWS CloudFormation User Guide* .
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes to create the stack. CloudFormation uses the role's credentials to make calls on your behalf. CloudFormation always uses this role for all future operations on the stack. Provided that users have permission to operate on the stack, CloudFormation uses this role even if the users don't have permission to pass it. Ensure that the role grants least privilege.

        If you don't specify a value, CloudFormation uses the role that was previously associated with the stack. If no role is available, CloudFormation uses a temporary session that's generated from your user credentials.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="rootId")
    def root_id(self) -> Optional[builtins.str]:
        """
        For nested stacks, the stack ID of the top-level stack to which the nested stack ultimately belongs.

        For more information, see [Nested stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html) in the *AWS CloudFormation User Guide* .
        """
        return pulumi.get(self, "root_id")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> Optional[builtins.str]:
        """
        Unique identifier of the stack.
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="stackPolicyBody")
    def stack_policy_body(self) -> Optional[Any]:
        """
        Structure that contains the stack policy body. For more information, see [Prevent updates to stack resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html) in the *AWS CloudFormation User Guide* . You can specify either the `StackPolicyBody` or the `StackPolicyURL` parameter, but not both.

        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CloudFormation::Stack` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "stack_policy_body")

    @property
    @pulumi.getter(name="stackStatus")
    def stack_status(self) -> Optional['StackStatus']:
        """
        Current status of the stack.
        """
        return pulumi.get(self, "stack_status")

    @property
    @pulumi.getter(name="stackStatusReason")
    def stack_status_reason(self) -> Optional[builtins.str]:
        """
        Success/failure message associated with the stack status.
        """
        return pulumi.get(self, "stack_status_reason")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Key-value pairs to associate with this stack. CloudFormation also propagates these tags to the resources created in the stack. A maximum number of 50 tags can be specified.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateBody")
    def template_body(self) -> Optional[Any]:
        """
        Structure that contains the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.

        Conditional: You must specify either the `TemplateBody` or the `TemplateURL` parameter, but not both.

        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CloudFormation::Stack` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "template_body")

    @property
    @pulumi.getter(name="timeoutInMinutes")
    def timeout_in_minutes(self) -> Optional[builtins.int]:
        """
        The length of time, in minutes, that CloudFormation waits for the nested stack to reach the `CREATE_COMPLETE` state. The default is no timeout. When CloudFormation detects that the nested stack has reached the `CREATE_COMPLETE` state, it marks the nested stack resource as `CREATE_COMPLETE` in the parent stack and resumes creating the parent stack. If the timeout period expires before the nested stack reaches `CREATE_COMPLETE` , CloudFormation marks the nested stack as failed and rolls back both the nested stack and parent stack.

        Updates aren't supported.
        """
        return pulumi.get(self, "timeout_in_minutes")


class AwaitableGetStackResult(GetStackResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStackResult(
            capabilities=self.capabilities,
            change_set_id=self.change_set_id,
            creation_time=self.creation_time,
            description=self.description,
            disable_rollback=self.disable_rollback,
            enable_termination_protection=self.enable_termination_protection,
            last_update_time=self.last_update_time,
            notification_arns=self.notification_arns,
            outputs=self.outputs,
            parameters=self.parameters,
            parent_id=self.parent_id,
            role_arn=self.role_arn,
            root_id=self.root_id,
            stack_id=self.stack_id,
            stack_policy_body=self.stack_policy_body,
            stack_status=self.stack_status,
            stack_status_reason=self.stack_status_reason,
            tags=self.tags,
            template_body=self.template_body,
            timeout_in_minutes=self.timeout_in_minutes)


def get_stack(stack_id: Optional[builtins.str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStackResult:
    """
    The AWS::CloudFormation::Stack resource nests a stack as a resource in a top-level template.


    :param builtins.str stack_id: Unique identifier of the stack.
    """
    __args__ = dict()
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cloudformation:getStack', __args__, opts=opts, typ=GetStackResult).value

    return AwaitableGetStackResult(
        capabilities=pulumi.get(__ret__, 'capabilities'),
        change_set_id=pulumi.get(__ret__, 'change_set_id'),
        creation_time=pulumi.get(__ret__, 'creation_time'),
        description=pulumi.get(__ret__, 'description'),
        disable_rollback=pulumi.get(__ret__, 'disable_rollback'),
        enable_termination_protection=pulumi.get(__ret__, 'enable_termination_protection'),
        last_update_time=pulumi.get(__ret__, 'last_update_time'),
        notification_arns=pulumi.get(__ret__, 'notification_arns'),
        outputs=pulumi.get(__ret__, 'outputs'),
        parameters=pulumi.get(__ret__, 'parameters'),
        parent_id=pulumi.get(__ret__, 'parent_id'),
        role_arn=pulumi.get(__ret__, 'role_arn'),
        root_id=pulumi.get(__ret__, 'root_id'),
        stack_id=pulumi.get(__ret__, 'stack_id'),
        stack_policy_body=pulumi.get(__ret__, 'stack_policy_body'),
        stack_status=pulumi.get(__ret__, 'stack_status'),
        stack_status_reason=pulumi.get(__ret__, 'stack_status_reason'),
        tags=pulumi.get(__ret__, 'tags'),
        template_body=pulumi.get(__ret__, 'template_body'),
        timeout_in_minutes=pulumi.get(__ret__, 'timeout_in_minutes'))
def get_stack_output(stack_id: Optional[pulumi.Input[builtins.str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStackResult]:
    """
    The AWS::CloudFormation::Stack resource nests a stack as a resource in a top-level template.


    :param builtins.str stack_id: Unique identifier of the stack.
    """
    __args__ = dict()
    __args__['stackId'] = stack_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:cloudformation:getStack', __args__, opts=opts, typ=GetStackResult)
    return __ret__.apply(lambda __response__: GetStackResult(
        capabilities=pulumi.get(__response__, 'capabilities'),
        change_set_id=pulumi.get(__response__, 'change_set_id'),
        creation_time=pulumi.get(__response__, 'creation_time'),
        description=pulumi.get(__response__, 'description'),
        disable_rollback=pulumi.get(__response__, 'disable_rollback'),
        enable_termination_protection=pulumi.get(__response__, 'enable_termination_protection'),
        last_update_time=pulumi.get(__response__, 'last_update_time'),
        notification_arns=pulumi.get(__response__, 'notification_arns'),
        outputs=pulumi.get(__response__, 'outputs'),
        parameters=pulumi.get(__response__, 'parameters'),
        parent_id=pulumi.get(__response__, 'parent_id'),
        role_arn=pulumi.get(__response__, 'role_arn'),
        root_id=pulumi.get(__response__, 'root_id'),
        stack_id=pulumi.get(__response__, 'stack_id'),
        stack_policy_body=pulumi.get(__response__, 'stack_policy_body'),
        stack_status=pulumi.get(__response__, 'stack_status'),
        stack_status_reason=pulumi.get(__response__, 'stack_status_reason'),
        tags=pulumi.get(__response__, 'tags'),
        template_body=pulumi.get(__response__, 'template_body'),
        timeout_in_minutes=pulumi.get(__response__, 'timeout_in_minutes')))
