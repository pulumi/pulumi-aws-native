# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'AgentActionGroupSignature',
    'AgentActionGroupState',
    'AgentAliasStatus',
    'AgentCollaboration',
    'AgentCreationMode',
    'AgentCustomControlMethod',
    'AgentKnowledgeBaseState',
    'AgentMemoryType',
    'AgentOrchestrationType',
    'AgentPromptState',
    'AgentPromptType',
    'AgentRelayConversationHistory',
    'AgentRequireConfirmation',
    'AgentStatus',
    'AgentType',
    'ApplicationInferenceProfileInferenceProfileStatus',
    'ApplicationInferenceProfileInferenceProfileType',
    'BlueprintStage',
    'BlueprintType',
    'DataAutomationProjectAudioExtractionCategoryType',
    'DataAutomationProjectAudioStandardGenerativeFieldType',
    'DataAutomationProjectBlueprintStage',
    'DataAutomationProjectDesiredModality',
    'DataAutomationProjectDocumentExtractionGranularityType',
    'DataAutomationProjectDocumentOutputTextFormatType',
    'DataAutomationProjectImageExtractionCategoryType',
    'DataAutomationProjectImageStandardGenerativeFieldType',
    'DataAutomationProjectStage',
    'DataAutomationProjectState',
    'DataAutomationProjectStatus',
    'DataAutomationProjectVideoExtractionCategoryType',
    'DataAutomationProjectVideoStandardGenerativeFieldType',
    'DataSourceChunkingStrategy',
    'DataSourceConfluenceSourceConfigurationAuthType',
    'DataSourceConfluenceSourceConfigurationHostType',
    'DataSourceContextEnrichmentType',
    'DataSourceCrawlFilterConfigurationType',
    'DataSourceDataDeletionPolicy',
    'DataSourceEnrichmentStrategyMethod',
    'DataSourceParsingModality',
    'DataSourceParsingStrategy',
    'DataSourceSalesforceSourceConfigurationAuthType',
    'DataSourceSharePointSourceConfigurationAuthType',
    'DataSourceSharePointSourceConfigurationHostType',
    'DataSourceStatus',
    'DataSourceTransformationStepToApply',
    'DataSourceType',
    'DataSourceWebScopeType',
    'FlowAliasConcurrencyType',
    'FlowConnectionType',
    'FlowNodeIoDataType',
    'FlowNodeType',
    'FlowPromptTemplateType',
    'FlowStatus',
    'FlowSupportedLanguages',
    'FlowVersionFlowConnectionType',
    'FlowVersionFlowNodeIoDataType',
    'FlowVersionFlowNodeType',
    'FlowVersionFlowStatus',
    'FlowVersionPromptTemplateType',
    'FlowVersionSupportedLanguages',
    'GuardrailContentFilterAction',
    'GuardrailContentFilterType',
    'GuardrailContextualGroundingAction',
    'GuardrailContextualGroundingFilterType',
    'GuardrailFilterStrength',
    'GuardrailManagedWordsType',
    'GuardrailModality',
    'GuardrailPiiEntityType',
    'GuardrailSensitiveInformationAction',
    'GuardrailStatus',
    'GuardrailTopicAction',
    'GuardrailTopicType',
    'GuardrailWordAction',
    'IntelligentPromptRouterPromptRouterStatus',
    'IntelligentPromptRouterPromptRouterType',
    'KnowledgeBaseBedrockEmbeddingModelConfigurationEmbeddingDataType',
    'KnowledgeBaseInclusionType',
    'KnowledgeBaseQueryEngineType',
    'KnowledgeBaseRedshiftProvisionedAuthType',
    'KnowledgeBaseRedshiftQueryEngineStorageType',
    'KnowledgeBaseRedshiftQueryEngineType',
    'KnowledgeBaseRedshiftServerlessAuthType',
    'KnowledgeBaseStatus',
    'KnowledgeBaseStorageType',
    'KnowledgeBaseSupplementalDataStorageLocationType',
    'KnowledgeBaseType',
    'PromptCachePointType',
    'PromptConversationRole',
    'PromptTemplateType',
    'PromptVersionCachePointType',
    'PromptVersionConversationRole',
    'PromptVersionPromptTemplateType',
]


class AgentActionGroupSignature(builtins.str, Enum):
    """
    Action Group Signature for a BuiltIn Action
    """
    AMAZON_USER_INPUT = "AMAZON.UserInput"
    AMAZON_CODE_INTERPRETER = "AMAZON.CodeInterpreter"


class AgentActionGroupState(builtins.str, Enum):
    """
    State of the action group
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class AgentAliasStatus(builtins.str, Enum):
    """
    The statuses an Agent Alias can be in.
    """
    CREATING = "CREATING"
    PREPARED = "PREPARED"
    FAILED = "FAILED"
    UPDATING = "UPDATING"
    DELETING = "DELETING"


class AgentCollaboration(builtins.str, Enum):
    """
    Agent collaboration state
    """
    DISABLED = "DISABLED"
    SUPERVISOR = "SUPERVISOR"
    SUPERVISOR_ROUTER = "SUPERVISOR_ROUTER"


class AgentCreationMode(builtins.str, Enum):
    """
    Creation Mode for Prompt Configuration.
    """
    DEFAULT = "DEFAULT"
    OVERRIDDEN = "OVERRIDDEN"


class AgentCustomControlMethod(builtins.str, Enum):
    """
    Custom control of action execution
    """
    RETURN_CONTROL = "RETURN_CONTROL"


class AgentKnowledgeBaseState(builtins.str, Enum):
    """
    State of the knowledge base; whether it is enabled or disabled
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class AgentMemoryType(builtins.str, Enum):
    """
    Memory type
    """
    SESSION_SUMMARY = "SESSION_SUMMARY"


class AgentOrchestrationType(builtins.str, Enum):
    """
    Types of orchestration strategy for agents
    """
    DEFAULT = "DEFAULT"
    CUSTOM_ORCHESTRATION = "CUSTOM_ORCHESTRATION"


class AgentPromptState(builtins.str, Enum):
    """
    Prompt State.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class AgentPromptType(builtins.str, Enum):
    """
    Prompt Type.
    """
    PRE_PROCESSING = "PRE_PROCESSING"
    ORCHESTRATION = "ORCHESTRATION"
    POST_PROCESSING = "POST_PROCESSING"
    ROUTING_CLASSIFIER = "ROUTING_CLASSIFIER"
    MEMORY_SUMMARIZATION = "MEMORY_SUMMARIZATION"
    KNOWLEDGE_BASE_RESPONSE_GENERATION = "KNOWLEDGE_BASE_RESPONSE_GENERATION"


class AgentRelayConversationHistory(builtins.str, Enum):
    """
    Relay conversation history state
    """
    TO_COLLABORATOR = "TO_COLLABORATOR"
    DISABLED = "DISABLED"


class AgentRequireConfirmation(builtins.str, Enum):
    """
    ENUM to check if action requires user confirmation
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class AgentStatus(builtins.str, Enum):
    """
    Schema Type for Action APIs.
    """
    CREATING = "CREATING"
    PREPARING = "PREPARING"
    PREPARED = "PREPARED"
    NOT_PREPARED = "NOT_PREPARED"
    DELETING = "DELETING"
    FAILED = "FAILED"
    VERSIONING = "VERSIONING"
    UPDATING = "UPDATING"


class AgentType(builtins.str, Enum):
    """
    Parameter Type
    """
    STRING = "string"
    NUMBER = "number"
    INTEGER = "integer"
    BOOLEAN = "boolean"
    ARRAY = "array"


class ApplicationInferenceProfileInferenceProfileStatus(builtins.str, Enum):
    """
    Status of the Inference Profile
    """
    ACTIVE = "ACTIVE"


class ApplicationInferenceProfileInferenceProfileType(builtins.str, Enum):
    """
    Type of the Inference Profile
    """
    APPLICATION = "APPLICATION"
    SYSTEM_DEFINED = "SYSTEM_DEFINED"


class BlueprintStage(builtins.str, Enum):
    """
    Stage of the Blueprint
    """
    DEVELOPMENT = "DEVELOPMENT"
    LIVE = "LIVE"


class BlueprintType(builtins.str, Enum):
    """
    Modality Type
    """
    DOCUMENT = "DOCUMENT"
    IMAGE = "IMAGE"
    AUDIO = "AUDIO"
    VIDEO = "VIDEO"


class DataAutomationProjectAudioExtractionCategoryType(builtins.str, Enum):
    AUDIO_CONTENT_MODERATION = "AUDIO_CONTENT_MODERATION"
    TRANSCRIPT = "TRANSCRIPT"
    TOPIC_CONTENT_MODERATION = "TOPIC_CONTENT_MODERATION"


class DataAutomationProjectAudioStandardGenerativeFieldType(builtins.str, Enum):
    AUDIO_SUMMARY = "AUDIO_SUMMARY"
    IAB = "IAB"
    TOPIC_SUMMARY = "TOPIC_SUMMARY"


class DataAutomationProjectBlueprintStage(builtins.str, Enum):
    """
    Stage of the Blueprint
    """
    DEVELOPMENT = "DEVELOPMENT"
    LIVE = "LIVE"


class DataAutomationProjectDesiredModality(builtins.str, Enum):
    DOCUMENT = "DOCUMENT"
    IMAGE = "IMAGE"
    VIDEO = "VIDEO"
    AUDIO = "AUDIO"


class DataAutomationProjectDocumentExtractionGranularityType(builtins.str, Enum):
    DOCUMENT = "DOCUMENT"
    PAGE = "PAGE"
    ELEMENT = "ELEMENT"
    WORD = "WORD"
    LINE = "LINE"


class DataAutomationProjectDocumentOutputTextFormatType(builtins.str, Enum):
    PLAIN_TEXT = "PLAIN_TEXT"
    MARKDOWN = "MARKDOWN"
    HTML = "HTML"
    CSV = "CSV"


class DataAutomationProjectImageExtractionCategoryType(builtins.str, Enum):
    CONTENT_MODERATION = "CONTENT_MODERATION"
    TEXT_DETECTION = "TEXT_DETECTION"
    LOGOS = "LOGOS"


class DataAutomationProjectImageStandardGenerativeFieldType(builtins.str, Enum):
    IMAGE_SUMMARY = "IMAGE_SUMMARY"
    IAB = "IAB"


class DataAutomationProjectStage(builtins.str, Enum):
    """
    Stage of the Project
    """
    DEVELOPMENT = "DEVELOPMENT"
    LIVE = "LIVE"


class DataAutomationProjectState(builtins.str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class DataAutomationProjectStatus(builtins.str, Enum):
    COMPLETED = "COMPLETED"
    IN_PROGRESS = "IN_PROGRESS"
    FAILED = "FAILED"


class DataAutomationProjectVideoExtractionCategoryType(builtins.str, Enum):
    CONTENT_MODERATION = "CONTENT_MODERATION"
    TEXT_DETECTION = "TEXT_DETECTION"
    TRANSCRIPT = "TRANSCRIPT"
    LOGOS = "LOGOS"


class DataAutomationProjectVideoStandardGenerativeFieldType(builtins.str, Enum):
    VIDEO_SUMMARY = "VIDEO_SUMMARY"
    IAB = "IAB"
    CHAPTER_SUMMARY = "CHAPTER_SUMMARY"


class DataSourceChunkingStrategy(builtins.str, Enum):
    """
    Knowledge base can split your source data into chunks. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. You have the following options for chunking your data. If you opt for NONE, then you may want to pre-process your files by splitting them up such that each file corresponds to a chunk.
    """
    FIXED_SIZE = "FIXED_SIZE"
    NONE = "NONE"
    HIERARCHICAL = "HIERARCHICAL"
    SEMANTIC = "SEMANTIC"


class DataSourceConfluenceSourceConfigurationAuthType(builtins.str, Enum):
    """
    The supported authentication type to authenticate and connect to your Confluence instance.
    """
    BASIC = "BASIC"
    OAUTH2_CLIENT_CREDENTIALS = "OAUTH2_CLIENT_CREDENTIALS"


class DataSourceConfluenceSourceConfigurationHostType(builtins.str, Enum):
    """
    The supported host type, whether online/cloud or server/on-premises.
    """
    SAAS = "SAAS"


class DataSourceContextEnrichmentType(builtins.str, Enum):
    """
    Enrichment type to be used for the vector database.
    """
    BEDROCK_FOUNDATION_MODEL = "BEDROCK_FOUNDATION_MODEL"


class DataSourceCrawlFilterConfigurationType(builtins.str, Enum):
    """
    The crawl filter type.
    """
    PATTERN = "PATTERN"


class DataSourceDataDeletionPolicy(builtins.str, Enum):
    """
    The deletion policy for the data source.
    """
    RETAIN = "RETAIN"
    DELETE = "DELETE"


class DataSourceEnrichmentStrategyMethod(builtins.str, Enum):
    """
    Enrichment Strategy method.
    """
    CHUNK_ENTITY_EXTRACTION = "CHUNK_ENTITY_EXTRACTION"


class DataSourceParsingModality(builtins.str, Enum):
    """
    Determine how will parsed content be stored.
    """
    MULTIMODAL = "MULTIMODAL"


class DataSourceParsingStrategy(builtins.str, Enum):
    """
    The parsing strategy for the data source.
    """
    BEDROCK_FOUNDATION_MODEL = "BEDROCK_FOUNDATION_MODEL"
    BEDROCK_DATA_AUTOMATION = "BEDROCK_DATA_AUTOMATION"


class DataSourceSalesforceSourceConfigurationAuthType(builtins.str, Enum):
    """
    The supported authentication type to authenticate and connect to your Salesforce instance.
    """
    OAUTH2_CLIENT_CREDENTIALS = "OAUTH2_CLIENT_CREDENTIALS"


class DataSourceSharePointSourceConfigurationAuthType(builtins.str, Enum):
    """
    The supported authentication type to authenticate and connect to your SharePoint site/sites.
    """
    OAUTH2_CLIENT_CREDENTIALS = "OAUTH2_CLIENT_CREDENTIALS"
    OAUTH2_SHAREPOINT_APP_ONLY_CLIENT_CREDENTIALS = "OAUTH2_SHAREPOINT_APP_ONLY_CLIENT_CREDENTIALS"


class DataSourceSharePointSourceConfigurationHostType(builtins.str, Enum):
    """
    The supported host type, whether online/cloud or server/on-premises.
    """
    ONLINE = "ONLINE"


class DataSourceStatus(builtins.str, Enum):
    """
    The status of a data source.
    """
    AVAILABLE = "AVAILABLE"
    DELETING = "DELETING"
    DELETE_UNSUCCESSFUL = "DELETE_UNSUCCESSFUL"


class DataSourceTransformationStepToApply(builtins.str, Enum):
    """
    When the service applies the transformation.
    """
    POST_CHUNKING = "POST_CHUNKING"


class DataSourceType(builtins.str, Enum):
    """
    The type of the data source location.
    """
    S3 = "S3"
    CONFLUENCE = "CONFLUENCE"
    SALESFORCE = "SALESFORCE"
    SHAREPOINT = "SHAREPOINT"
    WEB = "WEB"
    CUSTOM = "CUSTOM"
    REDSHIFT_METADATA = "REDSHIFT_METADATA"


class DataSourceWebScopeType(builtins.str, Enum):
    """
    The scope that a web crawl job will be restricted to.
    """
    HOST_ONLY = "HOST_ONLY"
    SUBDOMAINS = "SUBDOMAINS"


class FlowAliasConcurrencyType(builtins.str, Enum):
    AUTOMATIC = "Automatic"
    MANUAL = "Manual"


class FlowConnectionType(builtins.str, Enum):
    """
    Connection type
    """
    DATA = "Data"
    CONDITIONAL = "Conditional"


class FlowNodeIoDataType(builtins.str, Enum):
    """
    Type of input/output for a node in a flow
    """
    STRING = "String"
    NUMBER = "Number"
    BOOLEAN = "Boolean"
    OBJECT = "Object"
    ARRAY = "Array"


class FlowNodeType(builtins.str, Enum):
    """
    Flow node types
    """
    INPUT_TYPE = "Input"
    OUTPUT_TYPE = "Output"
    KNOWLEDGE_BASE = "KnowledgeBase"
    CONDITION = "Condition"
    LEX = "Lex"
    PROMPT = "Prompt"
    LAMBDA_FUNCTION = "LambdaFunction"
    AGENT = "Agent"
    STORAGE = "Storage"
    RETRIEVAL = "Retrieval"
    ITERATOR = "Iterator"
    COLLECTOR = "Collector"
    INLINE_CODE = "InlineCode"


class FlowPromptTemplateType(builtins.str, Enum):
    """
    Prompt template type
    """
    TEXT = "TEXT"


class FlowStatus(builtins.str, Enum):
    """
    Schema Type for Flow APIs
    """
    FAILED = "Failed"
    PREPARED = "Prepared"
    PREPARING = "Preparing"
    NOT_PREPARED = "NotPrepared"


class FlowSupportedLanguages(builtins.str, Enum):
    """
    Enum encodes the supported language type
    """
    PYTHON3 = "Python_3"


class FlowVersionFlowConnectionType(builtins.str, Enum):
    """
    Connection type
    """
    DATA = "Data"
    CONDITIONAL = "Conditional"


class FlowVersionFlowNodeIoDataType(builtins.str, Enum):
    """
    Type of input/output for a node in a flow
    """
    STRING = "String"
    NUMBER = "Number"
    BOOLEAN = "Boolean"
    OBJECT = "Object"
    ARRAY = "Array"


class FlowVersionFlowNodeType(builtins.str, Enum):
    """
    Flow node types
    """
    INPUT_TYPE = "Input"
    OUTPUT_TYPE = "Output"
    KNOWLEDGE_BASE = "KnowledgeBase"
    CONDITION = "Condition"
    LEX = "Lex"
    PROMPT = "Prompt"
    LAMBDA_FUNCTION = "LambdaFunction"
    AGENT = "Agent"
    ITERATOR = "Iterator"
    COLLECTOR = "Collector"
    STORAGE = "Storage"
    RETRIEVAL = "Retrieval"
    INLINE_CODE = "InlineCode"


class FlowVersionFlowStatus(builtins.str, Enum):
    """
    Schema Type for Flow APIs
    """
    FAILED = "Failed"
    PREPARED = "Prepared"
    PREPARING = "Preparing"
    NOT_PREPARED = "NotPrepared"


class FlowVersionPromptTemplateType(builtins.str, Enum):
    """
    Prompt template type
    """
    TEXT = "TEXT"


class FlowVersionSupportedLanguages(builtins.str, Enum):
    """
    Enum encodes the supported language type
    """
    PYTHON3 = "Python_3"


class GuardrailContentFilterAction(builtins.str, Enum):
    BLOCK = "BLOCK"
    NONE = "NONE"


class GuardrailContentFilterType(builtins.str, Enum):
    """
    Type of filter in content policy
    """
    SEXUAL = "SEXUAL"
    VIOLENCE = "VIOLENCE"
    HATE = "HATE"
    INSULTS = "INSULTS"
    MISCONDUCT = "MISCONDUCT"
    PROMPT_ATTACK = "PROMPT_ATTACK"


class GuardrailContextualGroundingAction(builtins.str, Enum):
    BLOCK = "BLOCK"
    NONE = "NONE"


class GuardrailContextualGroundingFilterType(builtins.str, Enum):
    """
    Type of contextual grounding filter
    """
    GROUNDING = "GROUNDING"
    RELEVANCE = "RELEVANCE"


class GuardrailFilterStrength(builtins.str, Enum):
    """
    Strength for filters
    """
    NONE = "NONE"
    LOW = "LOW"
    MEDIUM = "MEDIUM"
    HIGH = "HIGH"


class GuardrailManagedWordsType(builtins.str, Enum):
    """
    Options for managed words.
    """
    PROFANITY = "PROFANITY"


class GuardrailModality(builtins.str, Enum):
    """
    Modality for filters
    """
    TEXT = "TEXT"
    IMAGE = "IMAGE"


class GuardrailPiiEntityType(builtins.str, Enum):
    """
    The currently supported PII entities
    """
    ADDRESS = "ADDRESS"
    AGE = "AGE"
    AWS_ACCESS_KEY = "AWS_ACCESS_KEY"
    AWS_SECRET_KEY = "AWS_SECRET_KEY"
    CA_HEALTH_NUMBER = "CA_HEALTH_NUMBER"
    CA_SOCIAL_INSURANCE_NUMBER = "CA_SOCIAL_INSURANCE_NUMBER"
    CREDIT_DEBIT_CARD_CVV = "CREDIT_DEBIT_CARD_CVV"
    CREDIT_DEBIT_CARD_EXPIRY = "CREDIT_DEBIT_CARD_EXPIRY"
    CREDIT_DEBIT_CARD_NUMBER = "CREDIT_DEBIT_CARD_NUMBER"
    DRIVER_ID = "DRIVER_ID"
    EMAIL = "EMAIL"
    INTERNATIONAL_BANK_ACCOUNT_NUMBER = "INTERNATIONAL_BANK_ACCOUNT_NUMBER"
    IP_ADDRESS = "IP_ADDRESS"
    LICENSE_PLATE = "LICENSE_PLATE"
    MAC_ADDRESS = "MAC_ADDRESS"
    NAME = "NAME"
    PASSWORD = "PASSWORD"
    PHONE = "PHONE"
    PIN = "PIN"
    SWIFT_CODE = "SWIFT_CODE"
    UK_NATIONAL_HEALTH_SERVICE_NUMBER = "UK_NATIONAL_HEALTH_SERVICE_NUMBER"
    UK_NATIONAL_INSURANCE_NUMBER = "UK_NATIONAL_INSURANCE_NUMBER"
    UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER = "UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER"
    URL = "URL"
    USERNAME = "USERNAME"
    US_BANK_ACCOUNT_NUMBER = "US_BANK_ACCOUNT_NUMBER"
    US_BANK_ROUTING_NUMBER = "US_BANK_ROUTING_NUMBER"
    US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER = "US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER"
    US_PASSPORT_NUMBER = "US_PASSPORT_NUMBER"
    US_SOCIAL_SECURITY_NUMBER = "US_SOCIAL_SECURITY_NUMBER"
    VEHICLE_IDENTIFICATION_NUMBER = "VEHICLE_IDENTIFICATION_NUMBER"


class GuardrailSensitiveInformationAction(builtins.str, Enum):
    """
    Options for sensitive information action.
    """
    BLOCK = "BLOCK"
    ANONYMIZE = "ANONYMIZE"
    NONE = "NONE"


class GuardrailStatus(builtins.str, Enum):
    """
    Status of the guardrail
    """
    CREATING = "CREATING"
    UPDATING = "UPDATING"
    VERSIONING = "VERSIONING"
    READY = "READY"
    FAILED = "FAILED"
    DELETING = "DELETING"


class GuardrailTopicAction(builtins.str, Enum):
    BLOCK = "BLOCK"
    NONE = "NONE"


class GuardrailTopicType(builtins.str, Enum):
    """
    Type of topic in a policy
    """
    DENY = "DENY"


class GuardrailWordAction(builtins.str, Enum):
    BLOCK = "BLOCK"
    NONE = "NONE"


class IntelligentPromptRouterPromptRouterStatus(builtins.str, Enum):
    """
    Status of a PromptRouter
    """
    AVAILABLE = "AVAILABLE"


class IntelligentPromptRouterPromptRouterType(builtins.str, Enum):
    """
    Type of a Prompt Router
    """
    CUSTOM = "custom"
    DEFAULT = "default"


class KnowledgeBaseBedrockEmbeddingModelConfigurationEmbeddingDataType(builtins.str, Enum):
    """
    The data type for the vectors when using a model to convert text into vector embeddings.
    """
    FLOAT32 = "FLOAT32"
    BINARY = "BINARY"


class KnowledgeBaseInclusionType(builtins.str, Enum):
    """
    Include or Exclude status for an entity
    """
    INCLUDE = "INCLUDE"
    EXCLUDE = "EXCLUDE"


class KnowledgeBaseQueryEngineType(builtins.str, Enum):
    """
    SQL query engine type
    """
    REDSHIFT = "REDSHIFT"


class KnowledgeBaseRedshiftProvisionedAuthType(builtins.str, Enum):
    """
    Provisioned Redshift auth type
    """
    IAM = "IAM"
    USERNAME_PASSWORD = "USERNAME_PASSWORD"
    USERNAME = "USERNAME"


class KnowledgeBaseRedshiftQueryEngineStorageType(builtins.str, Enum):
    """
    Redshift query engine storage type
    """
    REDSHIFT = "REDSHIFT"
    AWS_DATA_CATALOG = "AWS_DATA_CATALOG"


class KnowledgeBaseRedshiftQueryEngineType(builtins.str, Enum):
    """
    Redshift query engine type
    """
    SERVERLESS = "SERVERLESS"
    PROVISIONED = "PROVISIONED"


class KnowledgeBaseRedshiftServerlessAuthType(builtins.str, Enum):
    """
    Serverless Redshift auth type
    """
    IAM = "IAM"
    USERNAME_PASSWORD = "USERNAME_PASSWORD"


class KnowledgeBaseStatus(builtins.str, Enum):
    """
    The status of a knowledge base.
    """
    CREATING = "CREATING"
    ACTIVE = "ACTIVE"
    DELETING = "DELETING"
    UPDATING = "UPDATING"
    FAILED = "FAILED"
    DELETE_UNSUCCESSFUL = "DELETE_UNSUCCESSFUL"


class KnowledgeBaseStorageType(builtins.str, Enum):
    """
    The storage type of a knowledge base.
    """
    OPENSEARCH_SERVERLESS = "OPENSEARCH_SERVERLESS"
    PINECONE = "PINECONE"
    RDS = "RDS"
    MONGO_DB_ATLAS = "MONGO_DB_ATLAS"
    NEPTUNE_ANALYTICS = "NEPTUNE_ANALYTICS"
    OPENSEARCH_MANAGED_CLUSTER = "OPENSEARCH_MANAGED_CLUSTER"


class KnowledgeBaseSupplementalDataStorageLocationType(builtins.str, Enum):
    """
    Supplemental data storage location type.
    """
    S3 = "S3"


class KnowledgeBaseType(builtins.str, Enum):
    """
    The type of a knowledge base.
    """
    VECTOR = "VECTOR"
    KENDRA = "KENDRA"
    SQL = "SQL"


class PromptCachePointType(builtins.str, Enum):
    """
    CachePoint types for CachePointBlock
    """
    DEFAULT = "default"


class PromptConversationRole(builtins.str, Enum):
    """
    Conversation roles for the chat prompt
    """
    USER = "user"
    ASSISTANT = "assistant"


class PromptTemplateType(builtins.str, Enum):
    """
    Prompt template type
    """
    TEXT = "TEXT"
    CHAT = "CHAT"


class PromptVersionCachePointType(builtins.str, Enum):
    """
    CachePoint types for CachePointBlock
    """
    DEFAULT = "default"


class PromptVersionConversationRole(builtins.str, Enum):
    """
    Conversation roles for the chat prompt
    """
    USER = "user"
    ASSISTANT = "assistant"


class PromptVersionPromptTemplateType(builtins.str, Enum):
    """
    Prompt template type
    """
    TEXT = "TEXT"
    CHAT = "CHAT"
