# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import outputs as _root_outputs

__all__ = [
    'GetPrivacyBudgetTemplateResult',
    'AwaitableGetPrivacyBudgetTemplateResult',
    'get_privacy_budget_template',
    'get_privacy_budget_template_output',
]

@pulumi.output_type
class GetPrivacyBudgetTemplateResult:
    def __init__(__self__, arn=None, collaboration_arn=None, collaboration_identifier=None, membership_arn=None, parameters=None, privacy_budget_template_identifier=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if collaboration_arn and not isinstance(collaboration_arn, str):
            raise TypeError("Expected argument 'collaboration_arn' to be a str")
        pulumi.set(__self__, "collaboration_arn", collaboration_arn)
        if collaboration_identifier and not isinstance(collaboration_identifier, str):
            raise TypeError("Expected argument 'collaboration_identifier' to be a str")
        pulumi.set(__self__, "collaboration_identifier", collaboration_identifier)
        if membership_arn and not isinstance(membership_arn, str):
            raise TypeError("Expected argument 'membership_arn' to be a str")
        pulumi.set(__self__, "membership_arn", membership_arn)
        if parameters and not isinstance(parameters, dict):
            raise TypeError("Expected argument 'parameters' to be a dict")
        pulumi.set(__self__, "parameters", parameters)
        if privacy_budget_template_identifier and not isinstance(privacy_budget_template_identifier, str):
            raise TypeError("Expected argument 'privacy_budget_template_identifier' to be a str")
        pulumi.set(__self__, "privacy_budget_template_identifier", privacy_budget_template_identifier)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The ARN of the privacy budget template.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="collaborationArn")
    def collaboration_arn(self) -> Optional[str]:
        """
        The ARN of the collaboration that contains this privacy budget template.
        """
        return pulumi.get(self, "collaboration_arn")

    @property
    @pulumi.getter(name="collaborationIdentifier")
    def collaboration_identifier(self) -> Optional[str]:
        """
        The unique ID of the collaboration that contains this privacy budget template.
        """
        return pulumi.get(self, "collaboration_identifier")

    @property
    @pulumi.getter(name="membershipArn")
    def membership_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the member who created the privacy budget template.
        """
        return pulumi.get(self, "membership_arn")

    @property
    @pulumi.getter
    def parameters(self) -> Optional['outputs.ParametersProperties']:
        """
        Specifies the epislon and noise parameters for the privacy budget template.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="privacyBudgetTemplateIdentifier")
    def privacy_budget_template_identifier(self) -> Optional[str]:
        """
        A unique identifier for one of your memberships for a collaboration. The privacy budget template is created in the collaboration that this membership belongs to. Accepts a membership ID.
        """
        return pulumi.get(self, "privacy_budget_template_identifier")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An arbitrary set of tags (key-value pairs) for this cleanrooms privacy budget template.
        """
        return pulumi.get(self, "tags")


class AwaitableGetPrivacyBudgetTemplateResult(GetPrivacyBudgetTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivacyBudgetTemplateResult(
            arn=self.arn,
            collaboration_arn=self.collaboration_arn,
            collaboration_identifier=self.collaboration_identifier,
            membership_arn=self.membership_arn,
            parameters=self.parameters,
            privacy_budget_template_identifier=self.privacy_budget_template_identifier,
            tags=self.tags)


def get_privacy_budget_template(membership_identifier: Optional[str] = None,
                                privacy_budget_template_identifier: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivacyBudgetTemplateResult:
    """
    Represents a privacy budget within a collaboration


    :param str membership_identifier: The identifier for a membership resource.
    :param str privacy_budget_template_identifier: A unique identifier for one of your memberships for a collaboration. The privacy budget template is created in the collaboration that this membership belongs to. Accepts a membership ID.
    """
    __args__ = dict()
    __args__['membershipIdentifier'] = membership_identifier
    __args__['privacyBudgetTemplateIdentifier'] = privacy_budget_template_identifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cleanrooms:getPrivacyBudgetTemplate', __args__, opts=opts, typ=GetPrivacyBudgetTemplateResult).value

    return AwaitableGetPrivacyBudgetTemplateResult(
        arn=pulumi.get(__ret__, 'arn'),
        collaboration_arn=pulumi.get(__ret__, 'collaboration_arn'),
        collaboration_identifier=pulumi.get(__ret__, 'collaboration_identifier'),
        membership_arn=pulumi.get(__ret__, 'membership_arn'),
        parameters=pulumi.get(__ret__, 'parameters'),
        privacy_budget_template_identifier=pulumi.get(__ret__, 'privacy_budget_template_identifier'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_privacy_budget_template)
def get_privacy_budget_template_output(membership_identifier: Optional[pulumi.Input[str]] = None,
                                       privacy_budget_template_identifier: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrivacyBudgetTemplateResult]:
    """
    Represents a privacy budget within a collaboration


    :param str membership_identifier: The identifier for a membership resource.
    :param str privacy_budget_template_identifier: A unique identifier for one of your memberships for a collaboration. The privacy budget template is created in the collaboration that this membership belongs to. Accepts a membership ID.
    """
    ...