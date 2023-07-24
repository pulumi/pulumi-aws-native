# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'DocumentClassifierAugmentedManifestsListItemArgs',
    'DocumentClassifierDocumentReaderConfigArgs',
    'DocumentClassifierDocumentsArgs',
    'DocumentClassifierInputDataConfigArgs',
    'DocumentClassifierOutputDataConfigArgs',
    'DocumentClassifierTagArgs',
    'DocumentClassifierVpcConfigArgs',
    'FlywheelDataSecurityConfigArgs',
    'FlywheelDocumentClassificationConfigArgs',
    'FlywheelEntityRecognitionConfigArgs',
    'FlywheelEntityTypesListItemArgs',
    'FlywheelTagArgs',
    'FlywheelTaskConfigArgs',
    'FlywheelVpcConfigArgs',
]

@pulumi.input_type
class DocumentClassifierAugmentedManifestsListItemArgs:
    def __init__(__self__, *,
                 attribute_names: pulumi.Input[Sequence[pulumi.Input[str]]],
                 s3_uri: pulumi.Input[str],
                 split: Optional[pulumi.Input['DocumentClassifierAugmentedManifestsListItemSplit']] = None):
        pulumi.set(__self__, "attribute_names", attribute_names)
        pulumi.set(__self__, "s3_uri", s3_uri)
        if split is not None:
            pulumi.set(__self__, "split", split)

    @property
    @pulumi.getter(name="attributeNames")
    def attribute_names(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "attribute_names")

    @attribute_names.setter
    def attribute_names(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "attribute_names", value)

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> pulumi.Input[str]:
        return pulumi.get(self, "s3_uri")

    @s3_uri.setter
    def s3_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_uri", value)

    @property
    @pulumi.getter
    def split(self) -> Optional[pulumi.Input['DocumentClassifierAugmentedManifestsListItemSplit']]:
        return pulumi.get(self, "split")

    @split.setter
    def split(self, value: Optional[pulumi.Input['DocumentClassifierAugmentedManifestsListItemSplit']]):
        pulumi.set(self, "split", value)


@pulumi.input_type
class DocumentClassifierDocumentReaderConfigArgs:
    def __init__(__self__, *,
                 document_read_action: pulumi.Input['DocumentClassifierDocumentReaderConfigDocumentReadAction'],
                 document_read_mode: Optional[pulumi.Input['DocumentClassifierDocumentReaderConfigDocumentReadMode']] = None,
                 feature_types: Optional[pulumi.Input[Sequence[pulumi.Input['DocumentClassifierDocumentReaderConfigFeatureTypesItem']]]] = None):
        pulumi.set(__self__, "document_read_action", document_read_action)
        if document_read_mode is not None:
            pulumi.set(__self__, "document_read_mode", document_read_mode)
        if feature_types is not None:
            pulumi.set(__self__, "feature_types", feature_types)

    @property
    @pulumi.getter(name="documentReadAction")
    def document_read_action(self) -> pulumi.Input['DocumentClassifierDocumentReaderConfigDocumentReadAction']:
        return pulumi.get(self, "document_read_action")

    @document_read_action.setter
    def document_read_action(self, value: pulumi.Input['DocumentClassifierDocumentReaderConfigDocumentReadAction']):
        pulumi.set(self, "document_read_action", value)

    @property
    @pulumi.getter(name="documentReadMode")
    def document_read_mode(self) -> Optional[pulumi.Input['DocumentClassifierDocumentReaderConfigDocumentReadMode']]:
        return pulumi.get(self, "document_read_mode")

    @document_read_mode.setter
    def document_read_mode(self, value: Optional[pulumi.Input['DocumentClassifierDocumentReaderConfigDocumentReadMode']]):
        pulumi.set(self, "document_read_mode", value)

    @property
    @pulumi.getter(name="featureTypes")
    def feature_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DocumentClassifierDocumentReaderConfigFeatureTypesItem']]]]:
        return pulumi.get(self, "feature_types")

    @feature_types.setter
    def feature_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DocumentClassifierDocumentReaderConfigFeatureTypesItem']]]]):
        pulumi.set(self, "feature_types", value)


@pulumi.input_type
class DocumentClassifierDocumentsArgs:
    def __init__(__self__, *,
                 s3_uri: pulumi.Input[str],
                 test_s3_uri: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "s3_uri", s3_uri)
        if test_s3_uri is not None:
            pulumi.set(__self__, "test_s3_uri", test_s3_uri)

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> pulumi.Input[str]:
        return pulumi.get(self, "s3_uri")

    @s3_uri.setter
    def s3_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_uri", value)

    @property
    @pulumi.getter(name="testS3Uri")
    def test_s3_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "test_s3_uri")

    @test_s3_uri.setter
    def test_s3_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "test_s3_uri", value)


@pulumi.input_type
class DocumentClassifierInputDataConfigArgs:
    def __init__(__self__, *,
                 augmented_manifests: Optional[pulumi.Input[Sequence[pulumi.Input['DocumentClassifierAugmentedManifestsListItemArgs']]]] = None,
                 data_format: Optional[pulumi.Input['DocumentClassifierInputDataConfigDataFormat']] = None,
                 document_reader_config: Optional[pulumi.Input['DocumentClassifierDocumentReaderConfigArgs']] = None,
                 document_type: Optional[pulumi.Input['DocumentClassifierInputDataConfigDocumentType']] = None,
                 documents: Optional[pulumi.Input['DocumentClassifierDocumentsArgs']] = None,
                 label_delimiter: Optional[pulumi.Input[str]] = None,
                 s3_uri: Optional[pulumi.Input[str]] = None,
                 test_s3_uri: Optional[pulumi.Input[str]] = None):
        if augmented_manifests is not None:
            pulumi.set(__self__, "augmented_manifests", augmented_manifests)
        if data_format is not None:
            pulumi.set(__self__, "data_format", data_format)
        if document_reader_config is not None:
            pulumi.set(__self__, "document_reader_config", document_reader_config)
        if document_type is not None:
            pulumi.set(__self__, "document_type", document_type)
        if documents is not None:
            pulumi.set(__self__, "documents", documents)
        if label_delimiter is not None:
            pulumi.set(__self__, "label_delimiter", label_delimiter)
        if s3_uri is not None:
            pulumi.set(__self__, "s3_uri", s3_uri)
        if test_s3_uri is not None:
            pulumi.set(__self__, "test_s3_uri", test_s3_uri)

    @property
    @pulumi.getter(name="augmentedManifests")
    def augmented_manifests(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DocumentClassifierAugmentedManifestsListItemArgs']]]]:
        return pulumi.get(self, "augmented_manifests")

    @augmented_manifests.setter
    def augmented_manifests(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DocumentClassifierAugmentedManifestsListItemArgs']]]]):
        pulumi.set(self, "augmented_manifests", value)

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> Optional[pulumi.Input['DocumentClassifierInputDataConfigDataFormat']]:
        return pulumi.get(self, "data_format")

    @data_format.setter
    def data_format(self, value: Optional[pulumi.Input['DocumentClassifierInputDataConfigDataFormat']]):
        pulumi.set(self, "data_format", value)

    @property
    @pulumi.getter(name="documentReaderConfig")
    def document_reader_config(self) -> Optional[pulumi.Input['DocumentClassifierDocumentReaderConfigArgs']]:
        return pulumi.get(self, "document_reader_config")

    @document_reader_config.setter
    def document_reader_config(self, value: Optional[pulumi.Input['DocumentClassifierDocumentReaderConfigArgs']]):
        pulumi.set(self, "document_reader_config", value)

    @property
    @pulumi.getter(name="documentType")
    def document_type(self) -> Optional[pulumi.Input['DocumentClassifierInputDataConfigDocumentType']]:
        return pulumi.get(self, "document_type")

    @document_type.setter
    def document_type(self, value: Optional[pulumi.Input['DocumentClassifierInputDataConfigDocumentType']]):
        pulumi.set(self, "document_type", value)

    @property
    @pulumi.getter
    def documents(self) -> Optional[pulumi.Input['DocumentClassifierDocumentsArgs']]:
        return pulumi.get(self, "documents")

    @documents.setter
    def documents(self, value: Optional[pulumi.Input['DocumentClassifierDocumentsArgs']]):
        pulumi.set(self, "documents", value)

    @property
    @pulumi.getter(name="labelDelimiter")
    def label_delimiter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "label_delimiter")

    @label_delimiter.setter
    def label_delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label_delimiter", value)

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "s3_uri")

    @s3_uri.setter
    def s3_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_uri", value)

    @property
    @pulumi.getter(name="testS3Uri")
    def test_s3_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "test_s3_uri")

    @test_s3_uri.setter
    def test_s3_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "test_s3_uri", value)


@pulumi.input_type
class DocumentClassifierOutputDataConfigArgs:
    def __init__(__self__, *,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 s3_uri: Optional[pulumi.Input[str]] = None):
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if s3_uri is not None:
            pulumi.set(__self__, "s3_uri", s3_uri)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "s3_uri")

    @s3_uri.setter
    def s3_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_uri", value)


@pulumi.input_type
class DocumentClassifierTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DocumentClassifierVpcConfigArgs:
    def __init__(__self__, *,
                 security_group_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subnets: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnets", subnets)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter
    def subnets(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "subnets")

    @subnets.setter
    def subnets(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subnets", value)


@pulumi.input_type
class FlywheelDataSecurityConfigArgs:
    def __init__(__self__, *,
                 data_lake_kms_key_id: Optional[pulumi.Input[str]] = None,
                 model_kms_key_id: Optional[pulumi.Input[str]] = None,
                 volume_kms_key_id: Optional[pulumi.Input[str]] = None,
                 vpc_config: Optional[pulumi.Input['FlywheelVpcConfigArgs']] = None):
        if data_lake_kms_key_id is not None:
            pulumi.set(__self__, "data_lake_kms_key_id", data_lake_kms_key_id)
        if model_kms_key_id is not None:
            pulumi.set(__self__, "model_kms_key_id", model_kms_key_id)
        if volume_kms_key_id is not None:
            pulumi.set(__self__, "volume_kms_key_id", volume_kms_key_id)
        if vpc_config is not None:
            pulumi.set(__self__, "vpc_config", vpc_config)

    @property
    @pulumi.getter(name="dataLakeKmsKeyId")
    def data_lake_kms_key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data_lake_kms_key_id")

    @data_lake_kms_key_id.setter
    def data_lake_kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_lake_kms_key_id", value)

    @property
    @pulumi.getter(name="modelKmsKeyId")
    def model_kms_key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "model_kms_key_id")

    @model_kms_key_id.setter
    def model_kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "model_kms_key_id", value)

    @property
    @pulumi.getter(name="volumeKmsKeyId")
    def volume_kms_key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "volume_kms_key_id")

    @volume_kms_key_id.setter
    def volume_kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_kms_key_id", value)

    @property
    @pulumi.getter(name="vpcConfig")
    def vpc_config(self) -> Optional[pulumi.Input['FlywheelVpcConfigArgs']]:
        return pulumi.get(self, "vpc_config")

    @vpc_config.setter
    def vpc_config(self, value: Optional[pulumi.Input['FlywheelVpcConfigArgs']]):
        pulumi.set(self, "vpc_config", value)


@pulumi.input_type
class FlywheelDocumentClassificationConfigArgs:
    def __init__(__self__, *,
                 mode: pulumi.Input['FlywheelDocumentClassificationConfigMode'],
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "mode", mode)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input['FlywheelDocumentClassificationConfigMode']:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input['FlywheelDocumentClassificationConfigMode']):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)


@pulumi.input_type
class FlywheelEntityRecognitionConfigArgs:
    def __init__(__self__, *,
                 entity_types: Optional[pulumi.Input[Sequence[pulumi.Input['FlywheelEntityTypesListItemArgs']]]] = None):
        if entity_types is not None:
            pulumi.set(__self__, "entity_types", entity_types)

    @property
    @pulumi.getter(name="entityTypes")
    def entity_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FlywheelEntityTypesListItemArgs']]]]:
        return pulumi.get(self, "entity_types")

    @entity_types.setter
    def entity_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FlywheelEntityTypesListItemArgs']]]]):
        pulumi.set(self, "entity_types", value)


@pulumi.input_type
class FlywheelEntityTypesListItemArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class FlywheelTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class FlywheelTaskConfigArgs:
    def __init__(__self__, *,
                 language_code: pulumi.Input['FlywheelTaskConfigLanguageCode'],
                 document_classification_config: Optional[pulumi.Input['FlywheelDocumentClassificationConfigArgs']] = None,
                 entity_recognition_config: Optional[pulumi.Input['FlywheelEntityRecognitionConfigArgs']] = None):
        pulumi.set(__self__, "language_code", language_code)
        if document_classification_config is not None:
            pulumi.set(__self__, "document_classification_config", document_classification_config)
        if entity_recognition_config is not None:
            pulumi.set(__self__, "entity_recognition_config", entity_recognition_config)

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> pulumi.Input['FlywheelTaskConfigLanguageCode']:
        return pulumi.get(self, "language_code")

    @language_code.setter
    def language_code(self, value: pulumi.Input['FlywheelTaskConfigLanguageCode']):
        pulumi.set(self, "language_code", value)

    @property
    @pulumi.getter(name="documentClassificationConfig")
    def document_classification_config(self) -> Optional[pulumi.Input['FlywheelDocumentClassificationConfigArgs']]:
        return pulumi.get(self, "document_classification_config")

    @document_classification_config.setter
    def document_classification_config(self, value: Optional[pulumi.Input['FlywheelDocumentClassificationConfigArgs']]):
        pulumi.set(self, "document_classification_config", value)

    @property
    @pulumi.getter(name="entityRecognitionConfig")
    def entity_recognition_config(self) -> Optional[pulumi.Input['FlywheelEntityRecognitionConfigArgs']]:
        return pulumi.get(self, "entity_recognition_config")

    @entity_recognition_config.setter
    def entity_recognition_config(self, value: Optional[pulumi.Input['FlywheelEntityRecognitionConfigArgs']]):
        pulumi.set(self, "entity_recognition_config", value)


@pulumi.input_type
class FlywheelVpcConfigArgs:
    def __init__(__self__, *,
                 security_group_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subnets: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnets", subnets)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter
    def subnets(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "subnets")

    @subnets.setter
    def subnets(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subnets", value)

