# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ConnectorKafkaClusterClientAuthenticationType',
    'ConnectorKafkaClusterEncryptionInTransitType',
    'CustomPluginContentType',
]


class ConnectorKafkaClusterClientAuthenticationType(str, Enum):
    """
    The type of client authentication used to connect to the Kafka cluster. Value NONE means that no client authentication is used.
    """
    NONE = "NONE"
    IAM = "IAM"


class ConnectorKafkaClusterEncryptionInTransitType(str, Enum):
    """
    The type of encryption in transit to the Kafka cluster.
    """
    PLAINTEXT = "PLAINTEXT"
    TLS = "TLS"


class CustomPluginContentType(str, Enum):
    """
    The type of the plugin file.
    """
    JAR = "JAR"
    ZIP = "ZIP"
