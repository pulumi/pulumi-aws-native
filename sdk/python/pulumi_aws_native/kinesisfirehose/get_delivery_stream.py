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
    'GetDeliveryStreamResult',
    'AwaitableGetDeliveryStreamResult',
    'get_delivery_stream',
    'get_delivery_stream_output',
]

@pulumi.output_type
class GetDeliveryStreamResult:
    def __init__(__self__, amazonopensearchservice_destination_configuration=None, arn=None, delivery_stream_encryption_configuration_input=None, elasticsearch_destination_configuration=None, extended_s3_destination_configuration=None, http_endpoint_destination_configuration=None, redshift_destination_configuration=None, s3_destination_configuration=None, splunk_destination_configuration=None, tags=None):
        if amazonopensearchservice_destination_configuration and not isinstance(amazonopensearchservice_destination_configuration, dict):
            raise TypeError("Expected argument 'amazonopensearchservice_destination_configuration' to be a dict")
        pulumi.set(__self__, "amazonopensearchservice_destination_configuration", amazonopensearchservice_destination_configuration)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if delivery_stream_encryption_configuration_input and not isinstance(delivery_stream_encryption_configuration_input, dict):
            raise TypeError("Expected argument 'delivery_stream_encryption_configuration_input' to be a dict")
        pulumi.set(__self__, "delivery_stream_encryption_configuration_input", delivery_stream_encryption_configuration_input)
        if elasticsearch_destination_configuration and not isinstance(elasticsearch_destination_configuration, dict):
            raise TypeError("Expected argument 'elasticsearch_destination_configuration' to be a dict")
        pulumi.set(__self__, "elasticsearch_destination_configuration", elasticsearch_destination_configuration)
        if extended_s3_destination_configuration and not isinstance(extended_s3_destination_configuration, dict):
            raise TypeError("Expected argument 'extended_s3_destination_configuration' to be a dict")
        pulumi.set(__self__, "extended_s3_destination_configuration", extended_s3_destination_configuration)
        if http_endpoint_destination_configuration and not isinstance(http_endpoint_destination_configuration, dict):
            raise TypeError("Expected argument 'http_endpoint_destination_configuration' to be a dict")
        pulumi.set(__self__, "http_endpoint_destination_configuration", http_endpoint_destination_configuration)
        if redshift_destination_configuration and not isinstance(redshift_destination_configuration, dict):
            raise TypeError("Expected argument 'redshift_destination_configuration' to be a dict")
        pulumi.set(__self__, "redshift_destination_configuration", redshift_destination_configuration)
        if s3_destination_configuration and not isinstance(s3_destination_configuration, dict):
            raise TypeError("Expected argument 's3_destination_configuration' to be a dict")
        pulumi.set(__self__, "s3_destination_configuration", s3_destination_configuration)
        if splunk_destination_configuration and not isinstance(splunk_destination_configuration, dict):
            raise TypeError("Expected argument 'splunk_destination_configuration' to be a dict")
        pulumi.set(__self__, "splunk_destination_configuration", splunk_destination_configuration)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="amazonopensearchserviceDestinationConfiguration")
    def amazonopensearchservice_destination_configuration(self) -> Optional['outputs.DeliveryStreamAmazonopensearchserviceDestinationConfiguration']:
        return pulumi.get(self, "amazonopensearchservice_destination_configuration")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="deliveryStreamEncryptionConfigurationInput")
    def delivery_stream_encryption_configuration_input(self) -> Optional['outputs.DeliveryStreamEncryptionConfigurationInput']:
        return pulumi.get(self, "delivery_stream_encryption_configuration_input")

    @property
    @pulumi.getter(name="elasticsearchDestinationConfiguration")
    def elasticsearch_destination_configuration(self) -> Optional['outputs.DeliveryStreamElasticsearchDestinationConfiguration']:
        return pulumi.get(self, "elasticsearch_destination_configuration")

    @property
    @pulumi.getter(name="extendedS3DestinationConfiguration")
    def extended_s3_destination_configuration(self) -> Optional['outputs.DeliveryStreamExtendedS3DestinationConfiguration']:
        return pulumi.get(self, "extended_s3_destination_configuration")

    @property
    @pulumi.getter(name="httpEndpointDestinationConfiguration")
    def http_endpoint_destination_configuration(self) -> Optional['outputs.DeliveryStreamHttpEndpointDestinationConfiguration']:
        return pulumi.get(self, "http_endpoint_destination_configuration")

    @property
    @pulumi.getter(name="redshiftDestinationConfiguration")
    def redshift_destination_configuration(self) -> Optional['outputs.DeliveryStreamRedshiftDestinationConfiguration']:
        return pulumi.get(self, "redshift_destination_configuration")

    @property
    @pulumi.getter(name="s3DestinationConfiguration")
    def s3_destination_configuration(self) -> Optional['outputs.DeliveryStreamS3DestinationConfiguration']:
        return pulumi.get(self, "s3_destination_configuration")

    @property
    @pulumi.getter(name="splunkDestinationConfiguration")
    def splunk_destination_configuration(self) -> Optional['outputs.DeliveryStreamSplunkDestinationConfiguration']:
        return pulumi.get(self, "splunk_destination_configuration")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DeliveryStreamTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetDeliveryStreamResult(GetDeliveryStreamResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeliveryStreamResult(
            amazonopensearchservice_destination_configuration=self.amazonopensearchservice_destination_configuration,
            arn=self.arn,
            delivery_stream_encryption_configuration_input=self.delivery_stream_encryption_configuration_input,
            elasticsearch_destination_configuration=self.elasticsearch_destination_configuration,
            extended_s3_destination_configuration=self.extended_s3_destination_configuration,
            http_endpoint_destination_configuration=self.http_endpoint_destination_configuration,
            redshift_destination_configuration=self.redshift_destination_configuration,
            s3_destination_configuration=self.s3_destination_configuration,
            splunk_destination_configuration=self.splunk_destination_configuration,
            tags=self.tags)


def get_delivery_stream(delivery_stream_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeliveryStreamResult:
    """
    Resource Type definition for AWS::KinesisFirehose::DeliveryStream
    """
    __args__ = dict()
    __args__['deliveryStreamName'] = delivery_stream_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:kinesisfirehose:getDeliveryStream', __args__, opts=opts, typ=GetDeliveryStreamResult).value

    return AwaitableGetDeliveryStreamResult(
        amazonopensearchservice_destination_configuration=__ret__.amazonopensearchservice_destination_configuration,
        arn=__ret__.arn,
        delivery_stream_encryption_configuration_input=__ret__.delivery_stream_encryption_configuration_input,
        elasticsearch_destination_configuration=__ret__.elasticsearch_destination_configuration,
        extended_s3_destination_configuration=__ret__.extended_s3_destination_configuration,
        http_endpoint_destination_configuration=__ret__.http_endpoint_destination_configuration,
        redshift_destination_configuration=__ret__.redshift_destination_configuration,
        s3_destination_configuration=__ret__.s3_destination_configuration,
        splunk_destination_configuration=__ret__.splunk_destination_configuration,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_delivery_stream)
def get_delivery_stream_output(delivery_stream_name: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeliveryStreamResult]:
    """
    Resource Type definition for AWS::KinesisFirehose::DeliveryStream
    """
    ...