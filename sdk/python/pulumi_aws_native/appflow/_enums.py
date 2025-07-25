# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'ConnectorProfileAuthenticationType',
    'ConnectorProfileConnectionMode',
    'ConnectorProfileConnectorType',
    'ConnectorProfileOAuth2GrantType',
    'FlowAggregationType',
    'FlowAmplitudeConnectorOperator',
    'FlowConnectorType',
    'FlowCustomConnectorSourcePropertiesDataTransferApiPropertiesType',
    'FlowDataTransferApi',
    'FlowDatadogConnectorOperator',
    'FlowDynatraceConnectorOperator',
    'FlowFileType',
    'FlowGoogleAnalyticsConnectorOperator',
    'FlowInforNexusConnectorOperator',
    'FlowMarketoConnectorOperator',
    'FlowOperator',
    'FlowOperatorPropertiesKeys',
    'FlowPardotConnectorOperator',
    'FlowPathPrefix',
    'FlowPrefixFormat',
    'FlowPrefixType',
    'FlowS3ConnectorOperator',
    'FlowS3InputFormatConfigS3InputFileType',
    'FlowSalesforceConnectorOperator',
    'FlowSapoDataConnectorOperator',
    'FlowScheduledTriggerPropertiesDataPullMode',
    'FlowServiceNowConnectorOperator',
    'FlowSingularConnectorOperator',
    'FlowSlackConnectorOperator',
    'FlowStatus',
    'FlowTaskType',
    'FlowTrendmicroConnectorOperator',
    'FlowTriggerType',
    'FlowVeevaConnectorOperator',
    'FlowWriteOperationType',
    'FlowZendeskConnectorOperator',
]


@pulumi.type_token("aws-native:appflow:ConnectorProfileAuthenticationType")
class ConnectorProfileAuthenticationType(builtins.str, Enum):
    OAUTH2 = "OAUTH2"
    APIKEY = "APIKEY"
    BASIC = "BASIC"
    CUSTOM = "CUSTOM"


@pulumi.type_token("aws-native:appflow:ConnectorProfileConnectionMode")
class ConnectorProfileConnectionMode(builtins.str, Enum):
    """
    Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
    """
    PUBLIC = "Public"
    PRIVATE = "Private"


@pulumi.type_token("aws-native:appflow:ConnectorProfileConnectorType")
class ConnectorProfileConnectorType(builtins.str, Enum):
    SALESFORCE = "Salesforce"
    PARDOT = "Pardot"
    SINGULAR = "Singular"
    SLACK = "Slack"
    REDSHIFT = "Redshift"
    MARKETO = "Marketo"
    GOOGLEANALYTICS = "Googleanalytics"
    ZENDESK = "Zendesk"
    SERVICENOW = "Servicenow"
    SAPO_DATA = "SAPOData"
    DATADOG = "Datadog"
    TRENDMICRO = "Trendmicro"
    SNOWFLAKE = "Snowflake"
    DYNATRACE = "Dynatrace"
    INFORNEXUS = "Infornexus"
    AMPLITUDE = "Amplitude"
    VEEVA = "Veeva"
    CUSTOM_CONNECTOR = "CustomConnector"


@pulumi.type_token("aws-native:appflow:ConnectorProfileOAuth2GrantType")
class ConnectorProfileOAuth2GrantType(builtins.str, Enum):
    CLIENT_CREDENTIALS = "CLIENT_CREDENTIALS"
    AUTHORIZATION_CODE = "AUTHORIZATION_CODE"
    JWT_BEARER = "JWT_BEARER"


@pulumi.type_token("aws-native:appflow:FlowAggregationType")
class FlowAggregationType(builtins.str, Enum):
    NONE = "None"
    SINGLE_FILE = "SingleFile"


@pulumi.type_token("aws-native:appflow:FlowAmplitudeConnectorOperator")
class FlowAmplitudeConnectorOperator(builtins.str, Enum):
    BETWEEN = "BETWEEN"


@pulumi.type_token("aws-native:appflow:FlowConnectorType")
class FlowConnectorType(builtins.str, Enum):
    SAPO_DATA = "SAPOData"
    SALESFORCE = "Salesforce"
    PARDOT = "Pardot"
    SINGULAR = "Singular"
    SLACK = "Slack"
    REDSHIFT = "Redshift"
    S3 = "S3"
    MARKETO = "Marketo"
    GOOGLEANALYTICS = "Googleanalytics"
    ZENDESK = "Zendesk"
    SERVICENOW = "Servicenow"
    DATADOG = "Datadog"
    TRENDMICRO = "Trendmicro"
    SNOWFLAKE = "Snowflake"
    DYNATRACE = "Dynatrace"
    INFORNEXUS = "Infornexus"
    AMPLITUDE = "Amplitude"
    VEEVA = "Veeva"
    CUSTOM_CONNECTOR = "CustomConnector"
    EVENT_BRIDGE = "EventBridge"
    UPSOLVER = "Upsolver"
    LOOKOUT_METRICS = "LookoutMetrics"


@pulumi.type_token("aws-native:appflow:FlowCustomConnectorSourcePropertiesDataTransferApiPropertiesType")
class FlowCustomConnectorSourcePropertiesDataTransferApiPropertiesType(builtins.str, Enum):
    SYNC = "SYNC"
    ASYNC_ = "ASYNC"
    AUTOMATIC = "AUTOMATIC"


@pulumi.type_token("aws-native:appflow:FlowDataTransferApi")
class FlowDataTransferApi(builtins.str, Enum):
    AUTOMATIC = "AUTOMATIC"
    BULKV2 = "BULKV2"
    REST_SYNC = "REST_SYNC"


@pulumi.type_token("aws-native:appflow:FlowDatadogConnectorOperator")
class FlowDatadogConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    BETWEEN = "BETWEEN"
    EQUAL_TO = "EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowDynatraceConnectorOperator")
class FlowDynatraceConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    BETWEEN = "BETWEEN"
    EQUAL_TO = "EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowFileType")
class FlowFileType(builtins.str, Enum):
    CSV = "CSV"
    JSON = "JSON"
    PARQUET = "PARQUET"


@pulumi.type_token("aws-native:appflow:FlowGoogleAnalyticsConnectorOperator")
class FlowGoogleAnalyticsConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    BETWEEN = "BETWEEN"


@pulumi.type_token("aws-native:appflow:FlowInforNexusConnectorOperator")
class FlowInforNexusConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    BETWEEN = "BETWEEN"
    EQUAL_TO = "EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowMarketoConnectorOperator")
class FlowMarketoConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    LESS_THAN = "LESS_THAN"
    GREATER_THAN = "GREATER_THAN"
    BETWEEN = "BETWEEN"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowOperator")
class FlowOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    LESS_THAN = "LESS_THAN"
    GREATER_THAN = "GREATER_THAN"
    CONTAINS = "CONTAINS"
    BETWEEN = "BETWEEN"
    LESS_THAN_OR_EQUAL_TO = "LESS_THAN_OR_EQUAL_TO"
    GREATER_THAN_OR_EQUAL_TO = "GREATER_THAN_OR_EQUAL_TO"
    EQUAL_TO = "EQUAL_TO"
    NOT_EQUAL_TO = "NOT_EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowOperatorPropertiesKeys")
class FlowOperatorPropertiesKeys(builtins.str, Enum):
    VALUE = "VALUE"
    VALUES = "VALUES"
    DATA_TYPE = "DATA_TYPE"
    UPPER_BOUND = "UPPER_BOUND"
    LOWER_BOUND = "LOWER_BOUND"
    SOURCE_DATA_TYPE = "SOURCE_DATA_TYPE"
    DESTINATION_DATA_TYPE = "DESTINATION_DATA_TYPE"
    VALIDATION_ACTION = "VALIDATION_ACTION"
    MASK_VALUE = "MASK_VALUE"
    MASK_LENGTH = "MASK_LENGTH"
    TRUNCATE_LENGTH = "TRUNCATE_LENGTH"
    MATH_OPERATION_FIELDS_ORDER = "MATH_OPERATION_FIELDS_ORDER"
    CONCAT_FORMAT = "CONCAT_FORMAT"
    SUBFIELD_CATEGORY_MAP = "SUBFIELD_CATEGORY_MAP"
    EXCLUDE_SOURCE_FIELDS_LIST = "EXCLUDE_SOURCE_FIELDS_LIST"
    INCLUDE_NEW_FIELDS = "INCLUDE_NEW_FIELDS"
    ORDERED_PARTITION_KEYS_LIST = "ORDERED_PARTITION_KEYS_LIST"


@pulumi.type_token("aws-native:appflow:FlowPardotConnectorOperator")
class FlowPardotConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    EQUAL_TO = "EQUAL_TO"
    NO_OP = "NO_OP"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"


@pulumi.type_token("aws-native:appflow:FlowPathPrefix")
class FlowPathPrefix(builtins.str, Enum):
    EXECUTION_ID = "EXECUTION_ID"
    SCHEMA_VERSION = "SCHEMA_VERSION"


@pulumi.type_token("aws-native:appflow:FlowPrefixFormat")
class FlowPrefixFormat(builtins.str, Enum):
    YEAR = "YEAR"
    MONTH = "MONTH"
    DAY = "DAY"
    HOUR = "HOUR"
    MINUTE = "MINUTE"


@pulumi.type_token("aws-native:appflow:FlowPrefixType")
class FlowPrefixType(builtins.str, Enum):
    FILENAME = "FILENAME"
    PATH = "PATH"
    PATH_AND_FILENAME = "PATH_AND_FILENAME"


@pulumi.type_token("aws-native:appflow:FlowS3ConnectorOperator")
class FlowS3ConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    LESS_THAN = "LESS_THAN"
    GREATER_THAN = "GREATER_THAN"
    BETWEEN = "BETWEEN"
    LESS_THAN_OR_EQUAL_TO = "LESS_THAN_OR_EQUAL_TO"
    GREATER_THAN_OR_EQUAL_TO = "GREATER_THAN_OR_EQUAL_TO"
    EQUAL_TO = "EQUAL_TO"
    NOT_EQUAL_TO = "NOT_EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowS3InputFormatConfigS3InputFileType")
class FlowS3InputFormatConfigS3InputFileType(builtins.str, Enum):
    """
    The file type that Amazon AppFlow gets from your Amazon S3 bucket.
    """
    CSV = "CSV"
    JSON = "JSON"


@pulumi.type_token("aws-native:appflow:FlowSalesforceConnectorOperator")
class FlowSalesforceConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    LESS_THAN = "LESS_THAN"
    CONTAINS = "CONTAINS"
    GREATER_THAN = "GREATER_THAN"
    BETWEEN = "BETWEEN"
    LESS_THAN_OR_EQUAL_TO = "LESS_THAN_OR_EQUAL_TO"
    GREATER_THAN_OR_EQUAL_TO = "GREATER_THAN_OR_EQUAL_TO"
    EQUAL_TO = "EQUAL_TO"
    NOT_EQUAL_TO = "NOT_EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowSapoDataConnectorOperator")
class FlowSapoDataConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    LESS_THAN = "LESS_THAN"
    CONTAINS = "CONTAINS"
    GREATER_THAN = "GREATER_THAN"
    BETWEEN = "BETWEEN"
    LESS_THAN_OR_EQUAL_TO = "LESS_THAN_OR_EQUAL_TO"
    GREATER_THAN_OR_EQUAL_TO = "GREATER_THAN_OR_EQUAL_TO"
    EQUAL_TO = "EQUAL_TO"
    NOT_EQUAL_TO = "NOT_EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowScheduledTriggerPropertiesDataPullMode")
class FlowScheduledTriggerPropertiesDataPullMode(builtins.str, Enum):
    """
    Specifies whether a scheduled flow has an incremental data transfer or a complete data transfer for each flow run.
    """
    INCREMENTAL = "Incremental"
    COMPLETE = "Complete"


@pulumi.type_token("aws-native:appflow:FlowServiceNowConnectorOperator")
class FlowServiceNowConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    LESS_THAN = "LESS_THAN"
    CONTAINS = "CONTAINS"
    GREATER_THAN = "GREATER_THAN"
    BETWEEN = "BETWEEN"
    LESS_THAN_OR_EQUAL_TO = "LESS_THAN_OR_EQUAL_TO"
    GREATER_THAN_OR_EQUAL_TO = "GREATER_THAN_OR_EQUAL_TO"
    EQUAL_TO = "EQUAL_TO"
    NOT_EQUAL_TO = "NOT_EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowSingularConnectorOperator")
class FlowSingularConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    EQUAL_TO = "EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowSlackConnectorOperator")
class FlowSlackConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    BETWEEN = "BETWEEN"
    EQUAL_TO = "EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowStatus")
class FlowStatus(builtins.str, Enum):
    """
    Flow activation status for Scheduled- and Event-triggered flows
    """
    ACTIVE = "Active"
    SUSPENDED = "Suspended"
    DRAFT = "Draft"


@pulumi.type_token("aws-native:appflow:FlowTaskType")
class FlowTaskType(builtins.str, Enum):
    ARITHMETIC = "Arithmetic"
    FILTER = "Filter"
    MAP = "Map"
    MAP_ALL = "Map_all"
    MASK = "Mask"
    MERGE = "Merge"
    PASSTHROUGH = "Passthrough"
    TRUNCATE = "Truncate"
    VALIDATE = "Validate"
    PARTITION = "Partition"


@pulumi.type_token("aws-native:appflow:FlowTrendmicroConnectorOperator")
class FlowTrendmicroConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    EQUAL_TO = "EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowTriggerType")
class FlowTriggerType(builtins.str, Enum):
    SCHEDULED = "Scheduled"
    EVENT = "Event"
    ON_DEMAND = "OnDemand"


@pulumi.type_token("aws-native:appflow:FlowVeevaConnectorOperator")
class FlowVeevaConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    LESS_THAN = "LESS_THAN"
    GREATER_THAN = "GREATER_THAN"
    BETWEEN = "BETWEEN"
    LESS_THAN_OR_EQUAL_TO = "LESS_THAN_OR_EQUAL_TO"
    GREATER_THAN_OR_EQUAL_TO = "GREATER_THAN_OR_EQUAL_TO"
    EQUAL_TO = "EQUAL_TO"
    NOT_EQUAL_TO = "NOT_EQUAL_TO"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"


@pulumi.type_token("aws-native:appflow:FlowWriteOperationType")
class FlowWriteOperationType(builtins.str, Enum):
    INSERT = "INSERT"
    UPSERT = "UPSERT"
    UPDATE = "UPDATE"
    DELETE = "DELETE"


@pulumi.type_token("aws-native:appflow:FlowZendeskConnectorOperator")
class FlowZendeskConnectorOperator(builtins.str, Enum):
    PROJECTION = "PROJECTION"
    GREATER_THAN = "GREATER_THAN"
    ADDITION = "ADDITION"
    MULTIPLICATION = "MULTIPLICATION"
    DIVISION = "DIVISION"
    SUBTRACTION = "SUBTRACTION"
    MASK_ALL = "MASK_ALL"
    MASK_FIRST_N = "MASK_FIRST_N"
    MASK_LAST_N = "MASK_LAST_N"
    VALIDATE_NON_NULL = "VALIDATE_NON_NULL"
    VALIDATE_NON_ZERO = "VALIDATE_NON_ZERO"
    VALIDATE_NON_NEGATIVE = "VALIDATE_NON_NEGATIVE"
    VALIDATE_NUMERIC = "VALIDATE_NUMERIC"
    NO_OP = "NO_OP"
