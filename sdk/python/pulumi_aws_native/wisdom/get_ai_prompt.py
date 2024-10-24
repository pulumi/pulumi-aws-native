# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'GetAiPromptResult',
    'AwaitableGetAiPromptResult',
    'get_ai_prompt',
    'get_ai_prompt_output',
]

@pulumi.output_type
class GetAiPromptResult:
    def __init__(__self__, ai_prompt_arn=None, ai_prompt_id=None, assistant_arn=None, description=None, template_configuration=None):
        if ai_prompt_arn and not isinstance(ai_prompt_arn, str):
            raise TypeError("Expected argument 'ai_prompt_arn' to be a str")
        pulumi.set(__self__, "ai_prompt_arn", ai_prompt_arn)
        if ai_prompt_id and not isinstance(ai_prompt_id, str):
            raise TypeError("Expected argument 'ai_prompt_id' to be a str")
        pulumi.set(__self__, "ai_prompt_id", ai_prompt_id)
        if assistant_arn and not isinstance(assistant_arn, str):
            raise TypeError("Expected argument 'assistant_arn' to be a str")
        pulumi.set(__self__, "assistant_arn", assistant_arn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if template_configuration and not isinstance(template_configuration, dict):
            raise TypeError("Expected argument 'template_configuration' to be a dict")
        pulumi.set(__self__, "template_configuration", template_configuration)

    @property
    @pulumi.getter(name="aiPromptArn")
    def ai_prompt_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the AI Prompt.
        """
        return pulumi.get(self, "ai_prompt_arn")

    @property
    @pulumi.getter(name="aiPromptId")
    def ai_prompt_id(self) -> Optional[str]:
        """
        The identifier of the Amazon Q in Connect AI prompt.
        """
        return pulumi.get(self, "ai_prompt_id")

    @property
    @pulumi.getter(name="assistantArn")
    def assistant_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the Amazon Q in Connect assistant.
        """
        return pulumi.get(self, "assistant_arn")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the AI Prompt.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="templateConfiguration")
    def template_configuration(self) -> Optional['outputs.AiPromptAiPromptTemplateConfiguration']:
        """
        The configuration of the prompt template for this AI Prompt.
        """
        return pulumi.get(self, "template_configuration")


class AwaitableGetAiPromptResult(GetAiPromptResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAiPromptResult(
            ai_prompt_arn=self.ai_prompt_arn,
            ai_prompt_id=self.ai_prompt_id,
            assistant_arn=self.assistant_arn,
            description=self.description,
            template_configuration=self.template_configuration)


def get_ai_prompt(ai_prompt_id: Optional[str] = None,
                  assistant_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAiPromptResult:
    """
    Definition of AWS::Wisdom::AIPrompt Resource Type


    :param str ai_prompt_id: The identifier of the Amazon Q in Connect AI prompt.
    :param str assistant_id: The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the ARN.
    """
    __args__ = dict()
    __args__['aiPromptId'] = ai_prompt_id
    __args__['assistantId'] = assistant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:wisdom:getAiPrompt', __args__, opts=opts, typ=GetAiPromptResult).value

    return AwaitableGetAiPromptResult(
        ai_prompt_arn=pulumi.get(__ret__, 'ai_prompt_arn'),
        ai_prompt_id=pulumi.get(__ret__, 'ai_prompt_id'),
        assistant_arn=pulumi.get(__ret__, 'assistant_arn'),
        description=pulumi.get(__ret__, 'description'),
        template_configuration=pulumi.get(__ret__, 'template_configuration'))
def get_ai_prompt_output(ai_prompt_id: Optional[pulumi.Input[str]] = None,
                         assistant_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAiPromptResult]:
    """
    Definition of AWS::Wisdom::AIPrompt Resource Type


    :param str ai_prompt_id: The identifier of the Amazon Q in Connect AI prompt.
    :param str assistant_id: The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the ARN.
    """
    __args__ = dict()
    __args__['aiPromptId'] = ai_prompt_id
    __args__['assistantId'] = assistant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:wisdom:getAiPrompt', __args__, opts=opts, typ=GetAiPromptResult)
    return __ret__.apply(lambda __response__: GetAiPromptResult(
        ai_prompt_arn=pulumi.get(__response__, 'ai_prompt_arn'),
        ai_prompt_id=pulumi.get(__response__, 'ai_prompt_id'),
        assistant_arn=pulumi.get(__response__, 'assistant_arn'),
        description=pulumi.get(__response__, 'description'),
        template_configuration=pulumi.get(__response__, 'template_configuration')))
