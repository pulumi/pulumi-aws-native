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

__all__ = [
    'GetConfigurationTemplateResult',
    'AwaitableGetConfigurationTemplateResult',
    'get_configuration_template',
    'get_configuration_template_output',
]

@pulumi.output_type
class GetConfigurationTemplateResult:
    def __init__(__self__, description=None, option_settings=None, template_name=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if option_settings and not isinstance(option_settings, list):
            raise TypeError("Expected argument 'option_settings' to be a list")
        pulumi.set(__self__, "option_settings", option_settings)
        if template_name and not isinstance(template_name, str):
            raise TypeError("Expected argument 'template_name' to be a str")
        pulumi.set(__self__, "template_name", template_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        An optional description for this configuration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="optionSettings")
    def option_settings(self) -> Optional[Sequence['outputs.ConfigurationTemplateConfigurationOptionSetting']]:
        """
        Option values for the Elastic Beanstalk configuration, such as the instance type. If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the AWS Elastic Beanstalk Developer Guide. 
        """
        return pulumi.get(self, "option_settings")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> Optional[str]:
        """
        The name of the configuration template
        """
        return pulumi.get(self, "template_name")


class AwaitableGetConfigurationTemplateResult(GetConfigurationTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfigurationTemplateResult(
            description=self.description,
            option_settings=self.option_settings,
            template_name=self.template_name)


def get_configuration_template(application_name: Optional[str] = None,
                               template_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfigurationTemplateResult:
    """
    Resource Type definition for AWS::ElasticBeanstalk::ConfigurationTemplate


    :param str application_name: The name of the Elastic Beanstalk application to associate with this configuration template. 
    :param str template_name: The name of the configuration template
    """
    __args__ = dict()
    __args__['applicationName'] = application_name
    __args__['templateName'] = template_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:elasticbeanstalk:getConfigurationTemplate', __args__, opts=opts, typ=GetConfigurationTemplateResult).value

    return AwaitableGetConfigurationTemplateResult(
        description=__ret__.description,
        option_settings=__ret__.option_settings,
        template_name=__ret__.template_name)


@_utilities.lift_output_func(get_configuration_template)
def get_configuration_template_output(application_name: Optional[pulumi.Input[str]] = None,
                                      template_name: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConfigurationTemplateResult]:
    """
    Resource Type definition for AWS::ElasticBeanstalk::ConfigurationTemplate


    :param str application_name: The name of the Elastic Beanstalk application to associate with this configuration template. 
    :param str template_name: The name of the configuration template
    """
    ...