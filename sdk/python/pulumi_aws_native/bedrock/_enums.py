# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AgentActionGroupSignature',
    'AgentActionGroupState',
    'AgentAliasStatus',
    'AgentCreationMode',
    'AgentCustomControlMethod',
    'AgentKnowledgeBaseState',
    'AgentPromptState',
    'AgentPromptType',
    'AgentStatus',
    'AgentType',
    'DataSourceChunkingStrategy',
    'DataSourceConfluenceSourceConfigurationAuthType',
    'DataSourceConfluenceSourceConfigurationHostType',
    'DataSourceCrawlFilterConfigurationType',
    'DataSourceDataDeletionPolicy',
    'DataSourceParsingStrategy',
    'DataSourceSalesforceSourceConfigurationAuthType',
    'DataSourceSharePointSourceConfigurationAuthType',
    'DataSourceSharePointSourceConfigurationHostType',
    'DataSourceStatus',
    'DataSourceTransformationStepToApply',
    'DataSourceType',
    'DataSourceWebScopeType',
    'FlowConnectionType',
    'FlowNodeIoDataType',
    'FlowNodeType',
    'FlowPromptTemplateType',
    'FlowStatus',
    'FlowVersionFlowConnectionType',
    'FlowVersionFlowNodeIoDataType',
    'FlowVersionFlowNodeType',
    'FlowVersionFlowStatus',
    'FlowVersionPromptTemplateType',
    'GuardrailContentFilterType',
    'GuardrailContextualGroundingFilterType',
    'GuardrailFilterStrength',
    'GuardrailManagedWordsType',
    'GuardrailPiiEntityType',
    'GuardrailSensitiveInformationAction',
    'GuardrailStatus',
    'GuardrailTopicType',
    'KnowledgeBaseStatus',
    'KnowledgeBaseStorageType',
    'KnowledgeBaseType',
    'PromptTemplateType',
    'PromptVersionPromptTemplateType',
]


class AgentActionGroupSignature(str, Enum):
    """
    Action Group Signature for a BuiltIn Action
    """
    AMAZON_USER_INPUT = "AMAZON.UserInput"


class AgentActionGroupState(str, Enum):
    """
    State of the action group
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class AgentAliasStatus(str, Enum):
    """
    The statuses an Agent Alias can be in.
    """
    CREATING = "CREATING"
    PREPARED = "PREPARED"
    FAILED = "FAILED"
    UPDATING = "UPDATING"
    DELETING = "DELETING"


class AgentCreationMode(str, Enum):
    """
    Creation Mode for Prompt Configuration.
    """
    DEFAULT = "DEFAULT"
    OVERRIDDEN = "OVERRIDDEN"


class AgentCustomControlMethod(str, Enum):
    """
    Custom control of action execution
    """
    RETURN_CONTROL = "RETURN_CONTROL"


class AgentKnowledgeBaseState(str, Enum):
    """
    State of the knowledge base; whether it is enabled or disabled
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class AgentPromptState(str, Enum):
    """
    Prompt State.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class AgentPromptType(str, Enum):
    """
    Prompt Type.
    """
    PRE_PROCESSING = "PRE_PROCESSING"
    ORCHESTRATION = "ORCHESTRATION"
    POST_PROCESSING = "POST_PROCESSING"
    KNOWLEDGE_BASE_RESPONSE_GENERATION = "KNOWLEDGE_BASE_RESPONSE_GENERATION"


class AgentStatus(str, Enum):
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


class AgentType(str, Enum):
    """
    Parameter Type
    """
    STRING = "string"
    NUMBER = "number"
    INTEGER = "integer"
    BOOLEAN = "boolean"
    ARRAY = "array"


class DataSourceChunkingStrategy(str, Enum):
    """
    Knowledge base can split your source data into chunks. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. You have the following options for chunking your data. If you opt for NONE, then you may want to pre-process your files by splitting them up such that each file corresponds to a chunk.
    """
    FIXED_SIZE = "FIXED_SIZE"
    NONE = "NONE"
    HIERARCHICAL = "HIERARCHICAL"
    SEMANTIC = "SEMANTIC"


class DataSourceConfluenceSourceConfigurationAuthType(str, Enum):
    """
    The supported authentication type to authenticate and connect to your Confluence instance.
    """
    BASIC = "BASIC"
    OAUTH2_CLIENT_CREDENTIALS = "OAUTH2_CLIENT_CREDENTIALS"


class DataSourceConfluenceSourceConfigurationHostType(str, Enum):
    """
    The supported host type, whether online/cloud or server/on-premises.
    """
    SAAS = "SAAS"


class DataSourceCrawlFilterConfigurationType(str, Enum):
    """
    The crawl filter type.
    """
    PATTERN = "PATTERN"


class DataSourceDataDeletionPolicy(str, Enum):
    """
    The deletion policy for the data source.
    """
    RETAIN = "RETAIN"
    DELETE = "DELETE"


class DataSourceParsingStrategy(str, Enum):
    """
    The parsing strategy for the data source.
    """
    BEDROCK_FOUNDATION_MODEL = "BEDROCK_FOUNDATION_MODEL"


class DataSourceSalesforceSourceConfigurationAuthType(str, Enum):
    """
    The supported authentication type to authenticate and connect to your Salesforce instance.
    """
    OAUTH2_CLIENT_CREDENTIALS = "OAUTH2_CLIENT_CREDENTIALS"


class DataSourceSharePointSourceConfigurationAuthType(str, Enum):
    """
    The supported authentication type to authenticate and connect to your SharePoint site/sites.
    """
    OAUTH2_CLIENT_CREDENTIALS = "OAUTH2_CLIENT_CREDENTIALS"


class DataSourceSharePointSourceConfigurationHostType(str, Enum):
    """
    The supported host type, whether online/cloud or server/on-premises.
    """
    ONLINE = "ONLINE"


class DataSourceStatus(str, Enum):
    """
    The status of a data source.
    """
    AVAILABLE = "AVAILABLE"
    DELETING = "DELETING"
    DELETE_UNSUCCESSFUL = "DELETE_UNSUCCESSFUL"


class DataSourceTransformationStepToApply(str, Enum):
    """
    When the service applies the transformation.
    """
    POST_CHUNKING = "POST_CHUNKING"


class DataSourceType(str, Enum):
    """
    The type of the data source location.
    """
    S3 = "S3"
    CONFLUENCE = "CONFLUENCE"
    SALESFORCE = "SALESFORCE"
    SHAREPOINT = "SHAREPOINT"
    WEB = "WEB"


class DataSourceWebScopeType(str, Enum):
    """
    The scope that a web crawl job will be restricted to.
    """
    HOST_ONLY = "HOST_ONLY"
    SUBDOMAINS = "SUBDOMAINS"


class FlowConnectionType(str, Enum):
    """
    Connection type
    """
    DATA = "Data"
    CONDITIONAL = "Conditional"


class FlowNodeIoDataType(str, Enum):
    """
    Type of input/output for a node in a flow
    """
    STRING = "String"
    NUMBER = "Number"
    BOOLEAN = "Boolean"
    OBJECT = "Object"
    ARRAY = "Array"


class FlowNodeType(str, Enum):
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


class FlowPromptTemplateType(str, Enum):
    """
    Prompt template type
    """
    TEXT = "TEXT"


class FlowStatus(str, Enum):
    """
    Schema Type for Flow APIs
    """
    FAILED = "Failed"
    PREPARED = "Prepared"
    PREPARING = "Preparing"
    NOT_PREPARED = "NotPrepared"


class FlowVersionFlowConnectionType(str, Enum):
    """
    Connection type
    """
    DATA = "Data"
    CONDITIONAL = "Conditional"


class FlowVersionFlowNodeIoDataType(str, Enum):
    """
    Type of input/output for a node in a flow
    """
    STRING = "String"
    NUMBER = "Number"
    BOOLEAN = "Boolean"
    OBJECT = "Object"
    ARRAY = "Array"


class FlowVersionFlowNodeType(str, Enum):
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


class FlowVersionFlowStatus(str, Enum):
    """
    Schema Type for Flow APIs
    """
    FAILED = "Failed"
    PREPARED = "Prepared"
    PREPARING = "Preparing"
    NOT_PREPARED = "NotPrepared"


class FlowVersionPromptTemplateType(str, Enum):
    """
    Prompt template type
    """
    TEXT = "TEXT"


class GuardrailContentFilterType(str, Enum):
    """
    Type of filter in content policy
    """
    SEXUAL = "SEXUAL"
    VIOLENCE = "VIOLENCE"
    HATE = "HATE"
    INSULTS = "INSULTS"
    MISCONDUCT = "MISCONDUCT"
    PROMPT_ATTACK = "PROMPT_ATTACK"


class GuardrailContextualGroundingFilterType(str, Enum):
    """
    Type of contextual grounding filter
    """
    GROUNDING = "GROUNDING"
    RELEVANCE = "RELEVANCE"


class GuardrailFilterStrength(str, Enum):
    """
    Strength for filters
    """
    NONE = "NONE"
    LOW = "LOW"
    MEDIUM = "MEDIUM"
    HIGH = "HIGH"


class GuardrailManagedWordsType(str, Enum):
    """
    Options for managed words.
    """
    PROFANITY = "PROFANITY"


class GuardrailPiiEntityType(str, Enum):
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


class GuardrailSensitiveInformationAction(str, Enum):
    """
    Options for sensitive information action.
    """
    BLOCK = "BLOCK"
    ANONYMIZE = "ANONYMIZE"


class GuardrailStatus(str, Enum):
    """
    Status of the guardrail
    """
    CREATING = "CREATING"
    UPDATING = "UPDATING"
    VERSIONING = "VERSIONING"
    READY = "READY"
    FAILED = "FAILED"
    DELETING = "DELETING"


class GuardrailTopicType(str, Enum):
    """
    Type of topic in a policy
    """
    DENY = "DENY"


class KnowledgeBaseStatus(str, Enum):
    """
    The status of a knowledge base.
    """
    CREATING = "CREATING"
    ACTIVE = "ACTIVE"
    DELETING = "DELETING"
    UPDATING = "UPDATING"
    FAILED = "FAILED"
    DELETE_UNSUCCESSFUL = "DELETE_UNSUCCESSFUL"


class KnowledgeBaseStorageType(str, Enum):
    """
    The storage type of a knowledge base.
    """
    OPENSEARCH_SERVERLESS = "OPENSEARCH_SERVERLESS"
    PINECONE = "PINECONE"
    RDS = "RDS"
    MONGO_DB_ATLAS = "MONGO_DB_ATLAS"


class KnowledgeBaseType(str, Enum):
    """
    The type of a knowledge base.
    """
    VECTOR = "VECTOR"


class PromptTemplateType(str, Enum):
    """
    Prompt template type
    """
    TEXT = "TEXT"


class PromptVersionPromptTemplateType(str, Enum):
    """
    Prompt template type
    """
    TEXT = "TEXT"
