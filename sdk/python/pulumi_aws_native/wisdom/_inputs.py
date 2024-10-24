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
from ._enums import *

__all__ = [
    'AiPromptAiPromptTemplateConfigurationArgs',
    'AiPromptAiPromptTemplateConfigurationArgsDict',
    'AssistantAssociationAssociationDataArgs',
    'AssistantAssociationAssociationDataArgsDict',
    'AssistantServerSideEncryptionConfigurationArgs',
    'AssistantServerSideEncryptionConfigurationArgsDict',
    'KnowledgeBaseAppIntegrationsConfigurationArgs',
    'KnowledgeBaseAppIntegrationsConfigurationArgsDict',
    'KnowledgeBaseRenderingConfigurationArgs',
    'KnowledgeBaseRenderingConfigurationArgsDict',
    'KnowledgeBaseServerSideEncryptionConfigurationArgs',
    'KnowledgeBaseServerSideEncryptionConfigurationArgsDict',
    'KnowledgeBaseSourceConfigurationArgs',
    'KnowledgeBaseSourceConfigurationArgsDict',
]

MYPY = False

if not MYPY:
    class AiPromptAiPromptTemplateConfigurationArgsDict(TypedDict):
        pass
elif False:
    AiPromptAiPromptTemplateConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AiPromptAiPromptTemplateConfigurationArgs:
    def __init__(__self__):
        pass


if not MYPY:
    class AssistantAssociationAssociationDataArgsDict(TypedDict):
        knowledge_base_id: pulumi.Input[str]
        """
        The identifier of the knowledge base.
        """
elif False:
    AssistantAssociationAssociationDataArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AssistantAssociationAssociationDataArgs:
    def __init__(__self__, *,
                 knowledge_base_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] knowledge_base_id: The identifier of the knowledge base.
        """
        pulumi.set(__self__, "knowledge_base_id", knowledge_base_id)

    @property
    @pulumi.getter(name="knowledgeBaseId")
    def knowledge_base_id(self) -> pulumi.Input[str]:
        """
        The identifier of the knowledge base.
        """
        return pulumi.get(self, "knowledge_base_id")

    @knowledge_base_id.setter
    def knowledge_base_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "knowledge_base_id", value)


if not MYPY:
    class AssistantServerSideEncryptionConfigurationArgsDict(TypedDict):
        kms_key_id: NotRequired[pulumi.Input[str]]
        """
        The customer managed key used for encryption. The customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom. To use Wisdom with chat, the key policy must also allow `kms:Decrypt` , `kms:GenerateDataKey*` , and `kms:DescribeKey` permissions to the `connect.amazonaws.com` service principal. For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) . For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) in the *AWS Key Management Service Developer Guide* .
        """
elif False:
    AssistantServerSideEncryptionConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AssistantServerSideEncryptionConfigurationArgs:
    def __init__(__self__, *,
                 kms_key_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] kms_key_id: The customer managed key used for encryption. The customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom. To use Wisdom with chat, the key policy must also allow `kms:Decrypt` , `kms:GenerateDataKey*` , and `kms:DescribeKey` permissions to the `connect.amazonaws.com` service principal. For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) . For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) in the *AWS Key Management Service Developer Guide* .
        """
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The customer managed key used for encryption. The customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom. To use Wisdom with chat, the key policy must also allow `kms:Decrypt` , `kms:GenerateDataKey*` , and `kms:DescribeKey` permissions to the `connect.amazonaws.com` service principal. For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) . For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) in the *AWS Key Management Service Developer Guide* .
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)


if not MYPY:
    class KnowledgeBaseAppIntegrationsConfigurationArgsDict(TypedDict):
        app_integration_arn: pulumi.Input[str]
        """
        The Amazon Resource Name (ARN) of the AppIntegrations DataIntegration to use for ingesting content.

        - For [Salesforce](https://docs.aws.amazon.com/https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm) , your AppIntegrations DataIntegration must have an ObjectConfiguration if objectFields is not provided, including at least `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , and `IsDeleted` as source fields.
        - For [ServiceNow](https://docs.aws.amazon.com/https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api) , your AppIntegrations DataIntegration must have an ObjectConfiguration if objectFields is not provided, including at least `number` , `short_description` , `sys_mod_count` , `workflow_state` , and `active` as source fields.
        - For [Zendesk](https://docs.aws.amazon.com/https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/) , your AppIntegrations DataIntegration must have an ObjectConfiguration if `objectFields` is not provided, including at least `id` , `title` , `updated_at` , and `draft` as source fields.
        - For [SharePoint](https://docs.aws.amazon.com/https://learn.microsoft.com/en-us/sharepoint/dev/sp-add-ins/sharepoint-net-server-csom-jsom-and-rest-api-index) , your AppIntegrations DataIntegration must have a FileConfiguration, including only file extensions that are among `docx` , `pdf` , `html` , `htm` , and `txt` .
        - For [Amazon S3](https://docs.aws.amazon.com/s3/) , the ObjectConfiguration and FileConfiguration of your AppIntegrations DataIntegration must be null. The `SourceURI` of your DataIntegration must use the following format: `s3://your_s3_bucket_name` .

        > The bucket policy of the corresponding S3 bucket must allow the AWS principal `app-integrations.amazonaws.com` to perform `s3:ListBucket` , `s3:GetObject` , and `s3:GetBucketLocation` against the bucket.
        """
        object_fields: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        The fields from the source that are made available to your agents in Amazon Q in Connect. Optional if ObjectConfiguration is included in the provided DataIntegration.

        - For [Salesforce](https://docs.aws.amazon.com/https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm) , you must include at least `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , and `IsDeleted` .
        - For [ServiceNow](https://docs.aws.amazon.com/https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api) , you must include at least `number` , `short_description` , `sys_mod_count` , `workflow_state` , and `active` .
        - For [Zendesk](https://docs.aws.amazon.com/https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/) , you must include at least `id` , `title` , `updated_at` , and `draft` .

        Make sure to include additional fields. These fields are indexed and used to source recommendations.
        """
elif False:
    KnowledgeBaseAppIntegrationsConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class KnowledgeBaseAppIntegrationsConfigurationArgs:
    def __init__(__self__, *,
                 app_integration_arn: pulumi.Input[str],
                 object_fields: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] app_integration_arn: The Amazon Resource Name (ARN) of the AppIntegrations DataIntegration to use for ingesting content.
               
               - For [Salesforce](https://docs.aws.amazon.com/https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm) , your AppIntegrations DataIntegration must have an ObjectConfiguration if objectFields is not provided, including at least `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , and `IsDeleted` as source fields.
               - For [ServiceNow](https://docs.aws.amazon.com/https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api) , your AppIntegrations DataIntegration must have an ObjectConfiguration if objectFields is not provided, including at least `number` , `short_description` , `sys_mod_count` , `workflow_state` , and `active` as source fields.
               - For [Zendesk](https://docs.aws.amazon.com/https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/) , your AppIntegrations DataIntegration must have an ObjectConfiguration if `objectFields` is not provided, including at least `id` , `title` , `updated_at` , and `draft` as source fields.
               - For [SharePoint](https://docs.aws.amazon.com/https://learn.microsoft.com/en-us/sharepoint/dev/sp-add-ins/sharepoint-net-server-csom-jsom-and-rest-api-index) , your AppIntegrations DataIntegration must have a FileConfiguration, including only file extensions that are among `docx` , `pdf` , `html` , `htm` , and `txt` .
               - For [Amazon S3](https://docs.aws.amazon.com/s3/) , the ObjectConfiguration and FileConfiguration of your AppIntegrations DataIntegration must be null. The `SourceURI` of your DataIntegration must use the following format: `s3://your_s3_bucket_name` .
               
               > The bucket policy of the corresponding S3 bucket must allow the AWS principal `app-integrations.amazonaws.com` to perform `s3:ListBucket` , `s3:GetObject` , and `s3:GetBucketLocation` against the bucket.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] object_fields: The fields from the source that are made available to your agents in Amazon Q in Connect. Optional if ObjectConfiguration is included in the provided DataIntegration.
               
               - For [Salesforce](https://docs.aws.amazon.com/https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm) , you must include at least `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , and `IsDeleted` .
               - For [ServiceNow](https://docs.aws.amazon.com/https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api) , you must include at least `number` , `short_description` , `sys_mod_count` , `workflow_state` , and `active` .
               - For [Zendesk](https://docs.aws.amazon.com/https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/) , you must include at least `id` , `title` , `updated_at` , and `draft` .
               
               Make sure to include additional fields. These fields are indexed and used to source recommendations.
        """
        pulumi.set(__self__, "app_integration_arn", app_integration_arn)
        if object_fields is not None:
            pulumi.set(__self__, "object_fields", object_fields)

    @property
    @pulumi.getter(name="appIntegrationArn")
    def app_integration_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the AppIntegrations DataIntegration to use for ingesting content.

        - For [Salesforce](https://docs.aws.amazon.com/https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm) , your AppIntegrations DataIntegration must have an ObjectConfiguration if objectFields is not provided, including at least `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , and `IsDeleted` as source fields.
        - For [ServiceNow](https://docs.aws.amazon.com/https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api) , your AppIntegrations DataIntegration must have an ObjectConfiguration if objectFields is not provided, including at least `number` , `short_description` , `sys_mod_count` , `workflow_state` , and `active` as source fields.
        - For [Zendesk](https://docs.aws.amazon.com/https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/) , your AppIntegrations DataIntegration must have an ObjectConfiguration if `objectFields` is not provided, including at least `id` , `title` , `updated_at` , and `draft` as source fields.
        - For [SharePoint](https://docs.aws.amazon.com/https://learn.microsoft.com/en-us/sharepoint/dev/sp-add-ins/sharepoint-net-server-csom-jsom-and-rest-api-index) , your AppIntegrations DataIntegration must have a FileConfiguration, including only file extensions that are among `docx` , `pdf` , `html` , `htm` , and `txt` .
        - For [Amazon S3](https://docs.aws.amazon.com/s3/) , the ObjectConfiguration and FileConfiguration of your AppIntegrations DataIntegration must be null. The `SourceURI` of your DataIntegration must use the following format: `s3://your_s3_bucket_name` .

        > The bucket policy of the corresponding S3 bucket must allow the AWS principal `app-integrations.amazonaws.com` to perform `s3:ListBucket` , `s3:GetObject` , and `s3:GetBucketLocation` against the bucket.
        """
        return pulumi.get(self, "app_integration_arn")

    @app_integration_arn.setter
    def app_integration_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_integration_arn", value)

    @property
    @pulumi.getter(name="objectFields")
    def object_fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The fields from the source that are made available to your agents in Amazon Q in Connect. Optional if ObjectConfiguration is included in the provided DataIntegration.

        - For [Salesforce](https://docs.aws.amazon.com/https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm) , you must include at least `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , and `IsDeleted` .
        - For [ServiceNow](https://docs.aws.amazon.com/https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api) , you must include at least `number` , `short_description` , `sys_mod_count` , `workflow_state` , and `active` .
        - For [Zendesk](https://docs.aws.amazon.com/https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/) , you must include at least `id` , `title` , `updated_at` , and `draft` .

        Make sure to include additional fields. These fields are indexed and used to source recommendations.
        """
        return pulumi.get(self, "object_fields")

    @object_fields.setter
    def object_fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "object_fields", value)


if not MYPY:
    class KnowledgeBaseRenderingConfigurationArgsDict(TypedDict):
        template_uri: NotRequired[pulumi.Input[str]]
        """
        A URI template containing exactly one variable in `${variableName}` format. This can only be set for `EXTERNAL` knowledge bases. For Salesforce, ServiceNow, and Zendesk, the variable must be one of the following:

        - Salesforce: `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , or `IsDeleted`
        - ServiceNow: `number` , `short_description` , `sys_mod_count` , `workflow_state` , or `active`
        - Zendesk: `id` , `title` , `updated_at` , or `draft`

        The variable is replaced with the actual value for a piece of content when calling [GetContent](https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_GetContent.html) .
        """
elif False:
    KnowledgeBaseRenderingConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class KnowledgeBaseRenderingConfigurationArgs:
    def __init__(__self__, *,
                 template_uri: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] template_uri: A URI template containing exactly one variable in `${variableName}` format. This can only be set for `EXTERNAL` knowledge bases. For Salesforce, ServiceNow, and Zendesk, the variable must be one of the following:
               
               - Salesforce: `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , or `IsDeleted`
               - ServiceNow: `number` , `short_description` , `sys_mod_count` , `workflow_state` , or `active`
               - Zendesk: `id` , `title` , `updated_at` , or `draft`
               
               The variable is replaced with the actual value for a piece of content when calling [GetContent](https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_GetContent.html) .
        """
        if template_uri is not None:
            pulumi.set(__self__, "template_uri", template_uri)

    @property
    @pulumi.getter(name="templateUri")
    def template_uri(self) -> Optional[pulumi.Input[str]]:
        """
        A URI template containing exactly one variable in `${variableName}` format. This can only be set for `EXTERNAL` knowledge bases. For Salesforce, ServiceNow, and Zendesk, the variable must be one of the following:

        - Salesforce: `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , or `IsDeleted`
        - ServiceNow: `number` , `short_description` , `sys_mod_count` , `workflow_state` , or `active`
        - Zendesk: `id` , `title` , `updated_at` , or `draft`

        The variable is replaced with the actual value for a piece of content when calling [GetContent](https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_GetContent.html) .
        """
        return pulumi.get(self, "template_uri")

    @template_uri.setter
    def template_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_uri", value)


if not MYPY:
    class KnowledgeBaseServerSideEncryptionConfigurationArgsDict(TypedDict):
        kms_key_id: NotRequired[pulumi.Input[str]]
        """
        The customer managed key used for encryption.

        This customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom.

        For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) . For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) .
        """
elif False:
    KnowledgeBaseServerSideEncryptionConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class KnowledgeBaseServerSideEncryptionConfigurationArgs:
    def __init__(__self__, *,
                 kms_key_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] kms_key_id: The customer managed key used for encryption.
               
               This customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom.
               
               For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) . For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) .
        """
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The customer managed key used for encryption.

        This customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom.

        For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) . For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) .
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)


if not MYPY:
    class KnowledgeBaseSourceConfigurationArgsDict(TypedDict):
        app_integrations: NotRequired[pulumi.Input['KnowledgeBaseAppIntegrationsConfigurationArgsDict']]
        """
        Configuration information for Amazon AppIntegrations to automatically ingest content.
        """
elif False:
    KnowledgeBaseSourceConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class KnowledgeBaseSourceConfigurationArgs:
    def __init__(__self__, *,
                 app_integrations: Optional[pulumi.Input['KnowledgeBaseAppIntegrationsConfigurationArgs']] = None):
        """
        :param pulumi.Input['KnowledgeBaseAppIntegrationsConfigurationArgs'] app_integrations: Configuration information for Amazon AppIntegrations to automatically ingest content.
        """
        if app_integrations is not None:
            pulumi.set(__self__, "app_integrations", app_integrations)

    @property
    @pulumi.getter(name="appIntegrations")
    def app_integrations(self) -> Optional[pulumi.Input['KnowledgeBaseAppIntegrationsConfigurationArgs']]:
        """
        Configuration information for Amazon AppIntegrations to automatically ingest content.
        """
        return pulumi.get(self, "app_integrations")

    @app_integrations.setter
    def app_integrations(self, value: Optional[pulumi.Input['KnowledgeBaseAppIntegrationsConfigurationArgs']]):
        pulumi.set(self, "app_integrations", value)


