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

__all__ = [
    'GetComputeNodeGroupResult',
    'AwaitableGetComputeNodeGroupResult',
    'get_compute_node_group',
    'get_compute_node_group_output',
]

@pulumi.output_type
class GetComputeNodeGroupResult:
    def __init__(__self__, ami_id=None, arn=None, custom_launch_template=None, error_info=None, iam_instance_profile_arn=None, id=None, purchase_option=None, scaling_configuration=None, slurm_configuration=None, spot_options=None, status=None, subnet_ids=None, tags=None):
        if ami_id and not isinstance(ami_id, str):
            raise TypeError("Expected argument 'ami_id' to be a str")
        pulumi.set(__self__, "ami_id", ami_id)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if custom_launch_template and not isinstance(custom_launch_template, dict):
            raise TypeError("Expected argument 'custom_launch_template' to be a dict")
        pulumi.set(__self__, "custom_launch_template", custom_launch_template)
        if error_info and not isinstance(error_info, list):
            raise TypeError("Expected argument 'error_info' to be a list")
        pulumi.set(__self__, "error_info", error_info)
        if iam_instance_profile_arn and not isinstance(iam_instance_profile_arn, str):
            raise TypeError("Expected argument 'iam_instance_profile_arn' to be a str")
        pulumi.set(__self__, "iam_instance_profile_arn", iam_instance_profile_arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if purchase_option and not isinstance(purchase_option, str):
            raise TypeError("Expected argument 'purchase_option' to be a str")
        pulumi.set(__self__, "purchase_option", purchase_option)
        if scaling_configuration and not isinstance(scaling_configuration, dict):
            raise TypeError("Expected argument 'scaling_configuration' to be a dict")
        pulumi.set(__self__, "scaling_configuration", scaling_configuration)
        if slurm_configuration and not isinstance(slurm_configuration, dict):
            raise TypeError("Expected argument 'slurm_configuration' to be a dict")
        pulumi.set(__self__, "slurm_configuration", slurm_configuration)
        if spot_options and not isinstance(spot_options, dict):
            raise TypeError("Expected argument 'spot_options' to be a dict")
        pulumi.set(__self__, "spot_options", spot_options)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="amiId")
    def ami_id(self) -> Optional[builtins.str]:
        """
        The ID of the Amazon Machine Image (AMI) that AWS PCS uses to launch instances. If not provided, AWS PCS uses the AMI ID specified in the custom launch template.
        """
        return pulumi.get(self, "ami_id")

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        The unique Amazon Resource Name (ARN) of the compute node group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="customLaunchTemplate")
    def custom_launch_template(self) -> Optional['outputs.CustomLaunchTemplateProperties']:
        """
        An Amazon EC2 launch template AWS PCS uses to launch compute nodes.
        """
        return pulumi.get(self, "custom_launch_template")

    @property
    @pulumi.getter(name="errorInfo")
    def error_info(self) -> Optional[Sequence['outputs.ComputeNodeGroupErrorInfo']]:
        """
        The list of errors that occurred during compute node group provisioning.
        """
        return pulumi.get(self, "error_info")

    @property
    @pulumi.getter(name="iamInstanceProfileArn")
    def iam_instance_profile_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when launching EC2 instances. The role contained in your instance profile must have pcs:RegisterComputeNodeGroupInstance permissions attached to provision instances correctly.
        """
        return pulumi.get(self, "iam_instance_profile_arn")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        The generated unique ID of the compute node group.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="purchaseOption")
    def purchase_option(self) -> Optional['ComputeNodeGroupPurchaseOption']:
        """
        Specifies how EC2 instances are purchased on your behalf. AWS PCS supports On-Demand and Spot instances. For more information, see Instance purchasing options in the Amazon Elastic Compute Cloud User Guide. If you don't provide this option, it defaults to On-Demand.
        """
        return pulumi.get(self, "purchase_option")

    @property
    @pulumi.getter(name="scalingConfiguration")
    def scaling_configuration(self) -> Optional['outputs.ScalingConfigurationProperties']:
        """
        Specifies the boundaries of the compute node group auto scaling.
        """
        return pulumi.get(self, "scaling_configuration")

    @property
    @pulumi.getter(name="slurmConfiguration")
    def slurm_configuration(self) -> Optional['outputs.SlurmConfigurationProperties']:
        """
        Additional options related to the Slurm scheduler.
        """
        return pulumi.get(self, "slurm_configuration")

    @property
    @pulumi.getter(name="spotOptions")
    def spot_options(self) -> Optional['outputs.SpotOptionsProperties']:
        """
        Additional configuration when you specify SPOT as the purchase option.
        """
        return pulumi.get(self, "spot_options")

    @property
    @pulumi.getter
    def status(self) -> Optional['ComputeNodeGroupStatus']:
        """
        The provisioning status of the compute node group. The provisioning status doesn't indicate the overall health of the compute node group.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[Sequence[builtins.str]]:
        """
        The list of subnet IDs where instances are provisioned by the compute node group. The subnets must be in the same VPC as the cluster.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.
        """
        return pulumi.get(self, "tags")


class AwaitableGetComputeNodeGroupResult(GetComputeNodeGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetComputeNodeGroupResult(
            ami_id=self.ami_id,
            arn=self.arn,
            custom_launch_template=self.custom_launch_template,
            error_info=self.error_info,
            iam_instance_profile_arn=self.iam_instance_profile_arn,
            id=self.id,
            purchase_option=self.purchase_option,
            scaling_configuration=self.scaling_configuration,
            slurm_configuration=self.slurm_configuration,
            spot_options=self.spot_options,
            status=self.status,
            subnet_ids=self.subnet_ids,
            tags=self.tags)


def get_compute_node_group(arn: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetComputeNodeGroupResult:
    """
    AWS::PCS::ComputeNodeGroup resource creates an AWS PCS compute node group.


    :param builtins.str arn: The unique Amazon Resource Name (ARN) of the compute node group.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:pcs:getComputeNodeGroup', __args__, opts=opts, typ=GetComputeNodeGroupResult).value

    return AwaitableGetComputeNodeGroupResult(
        ami_id=pulumi.get(__ret__, 'ami_id'),
        arn=pulumi.get(__ret__, 'arn'),
        custom_launch_template=pulumi.get(__ret__, 'custom_launch_template'),
        error_info=pulumi.get(__ret__, 'error_info'),
        iam_instance_profile_arn=pulumi.get(__ret__, 'iam_instance_profile_arn'),
        id=pulumi.get(__ret__, 'id'),
        purchase_option=pulumi.get(__ret__, 'purchase_option'),
        scaling_configuration=pulumi.get(__ret__, 'scaling_configuration'),
        slurm_configuration=pulumi.get(__ret__, 'slurm_configuration'),
        spot_options=pulumi.get(__ret__, 'spot_options'),
        status=pulumi.get(__ret__, 'status'),
        subnet_ids=pulumi.get(__ret__, 'subnet_ids'),
        tags=pulumi.get(__ret__, 'tags'))
def get_compute_node_group_output(arn: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetComputeNodeGroupResult]:
    """
    AWS::PCS::ComputeNodeGroup resource creates an AWS PCS compute node group.


    :param builtins.str arn: The unique Amazon Resource Name (ARN) of the compute node group.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:pcs:getComputeNodeGroup', __args__, opts=opts, typ=GetComputeNodeGroupResult)
    return __ret__.apply(lambda __response__: GetComputeNodeGroupResult(
        ami_id=pulumi.get(__response__, 'ami_id'),
        arn=pulumi.get(__response__, 'arn'),
        custom_launch_template=pulumi.get(__response__, 'custom_launch_template'),
        error_info=pulumi.get(__response__, 'error_info'),
        iam_instance_profile_arn=pulumi.get(__response__, 'iam_instance_profile_arn'),
        id=pulumi.get(__response__, 'id'),
        purchase_option=pulumi.get(__response__, 'purchase_option'),
        scaling_configuration=pulumi.get(__response__, 'scaling_configuration'),
        slurm_configuration=pulumi.get(__response__, 'slurm_configuration'),
        spot_options=pulumi.get(__response__, 'spot_options'),
        status=pulumi.get(__response__, 'status'),
        subnet_ids=pulumi.get(__response__, 'subnet_ids'),
        tags=pulumi.get(__response__, 'tags')))
