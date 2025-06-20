# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'AnalysisTemplateAnalysisParameterType',
    'AnalysisTemplateFormat',
    'CollaborationAnalyticsEngine',
    'CollaborationCustomMlMemberAbility',
    'CollaborationJobLogStatus',
    'CollaborationMemberAbility',
    'CollaborationQueryLogStatus',
    'ConfiguredTableAdditionalAnalyses',
    'ConfiguredTableAggregateFunctionName',
    'ConfiguredTableAggregationType',
    'ConfiguredTableAnalysisMethod',
    'ConfiguredTableAnalysisRuleType',
    'ConfiguredTableAssociationAnalysisRuleType',
    'ConfiguredTableJoinOperator',
    'ConfiguredTableJoinRequiredOption',
    'ConfiguredTableScalarFunctions',
    'ConfiguredTableSelectedAnalysisMethod',
    'IdMappingTableInputSourceType',
    'IdNamespaceAssociationInputReferencePropertiesIdNamespaceType',
    'MembershipJobLogStatus',
    'MembershipQueryLogStatus',
    'MembershipResultFormat',
    'PrivacyBudgetTemplateAutoRefresh',
    'PrivacyBudgetTemplatePrivacyBudgetType',
]


@pulumi.type_token("aws-native:cleanrooms:AnalysisTemplateAnalysisParameterType")
class AnalysisTemplateAnalysisParameterType(builtins.str, Enum):
    """
    The type of parameter.
    """
    SMALLINT = "SMALLINT"
    INTEGER = "INTEGER"
    BIGINT = "BIGINT"
    DECIMAL = "DECIMAL"
    REAL = "REAL"
    DOUBLE_PRECISION = "DOUBLE_PRECISION"
    BOOLEAN = "BOOLEAN"
    CHAR = "CHAR"
    VARCHAR = "VARCHAR"
    DATE = "DATE"
    TIMESTAMP = "TIMESTAMP"
    TIMESTAMPTZ = "TIMESTAMPTZ"
    TIME = "TIME"
    TIMETZ = "TIMETZ"
    VARBYTE = "VARBYTE"
    BINARY = "BINARY"
    BYTE = "BYTE"
    CHARACTER = "CHARACTER"
    DOUBLE = "DOUBLE"
    FLOAT = "FLOAT"
    INT = "INT"
    LONG = "LONG"
    NUMERIC = "NUMERIC"
    SHORT = "SHORT"
    STRING = "STRING"
    TIMESTAMP_LTZ = "TIMESTAMP_LTZ"
    TIMESTAMP_NTZ = "TIMESTAMP_NTZ"
    TINYINT = "TINYINT"


@pulumi.type_token("aws-native:cleanrooms:AnalysisTemplateFormat")
class AnalysisTemplateFormat(builtins.str, Enum):
    """
    The format of the analysis template.
    """
    SQL = "SQL"
    PYSPARK10 = "PYSPARK_1_0"


@pulumi.type_token("aws-native:cleanrooms:CollaborationAnalyticsEngine")
class CollaborationAnalyticsEngine(builtins.str, Enum):
    CLEAN_ROOMS_SQL = "CLEAN_ROOMS_SQL"
    SPARK = "SPARK"


@pulumi.type_token("aws-native:cleanrooms:CollaborationCustomMlMemberAbility")
class CollaborationCustomMlMemberAbility(builtins.str, Enum):
    CAN_RECEIVE_MODEL_OUTPUT = "CAN_RECEIVE_MODEL_OUTPUT"
    CAN_RECEIVE_INFERENCE_OUTPUT = "CAN_RECEIVE_INFERENCE_OUTPUT"


@pulumi.type_token("aws-native:cleanrooms:CollaborationJobLogStatus")
class CollaborationJobLogStatus(builtins.str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


@pulumi.type_token("aws-native:cleanrooms:CollaborationMemberAbility")
class CollaborationMemberAbility(builtins.str, Enum):
    CAN_QUERY = "CAN_QUERY"
    CAN_RUN_JOB = "CAN_RUN_JOB"
    CAN_RECEIVE_RESULTS = "CAN_RECEIVE_RESULTS"


@pulumi.type_token("aws-native:cleanrooms:CollaborationQueryLogStatus")
class CollaborationQueryLogStatus(builtins.str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


@pulumi.type_token("aws-native:cleanrooms:ConfiguredTableAdditionalAnalyses")
class ConfiguredTableAdditionalAnalyses(builtins.str, Enum):
    ALLOWED = "ALLOWED"
    REQUIRED = "REQUIRED"
    NOT_ALLOWED = "NOT_ALLOWED"


@pulumi.type_token("aws-native:cleanrooms:ConfiguredTableAggregateFunctionName")
class ConfiguredTableAggregateFunctionName(builtins.str, Enum):
    SUM = "SUM"
    SUM_DISTINCT = "SUM_DISTINCT"
    COUNT = "COUNT"
    COUNT_DISTINCT = "COUNT_DISTINCT"
    AVG = "AVG"


@pulumi.type_token("aws-native:cleanrooms:ConfiguredTableAggregationType")
class ConfiguredTableAggregationType(builtins.str, Enum):
    COUNT_DISTINCT = "COUNT_DISTINCT"


@pulumi.type_token("aws-native:cleanrooms:ConfiguredTableAnalysisMethod")
class ConfiguredTableAnalysisMethod(builtins.str, Enum):
    DIRECT_QUERY = "DIRECT_QUERY"
    DIRECT_JOB = "DIRECT_JOB"
    MULTIPLE = "MULTIPLE"


@pulumi.type_token("aws-native:cleanrooms:ConfiguredTableAnalysisRuleType")
class ConfiguredTableAnalysisRuleType(builtins.str, Enum):
    AGGREGATION = "AGGREGATION"
    LIST = "LIST"
    CUSTOM = "CUSTOM"


@pulumi.type_token("aws-native:cleanrooms:ConfiguredTableAssociationAnalysisRuleType")
class ConfiguredTableAssociationAnalysisRuleType(builtins.str, Enum):
    AGGREGATION = "AGGREGATION"
    LIST = "LIST"
    CUSTOM = "CUSTOM"


@pulumi.type_token("aws-native:cleanrooms:ConfiguredTableJoinOperator")
class ConfiguredTableJoinOperator(builtins.str, Enum):
    OR_ = "OR"
    AND_ = "AND"


@pulumi.type_token("aws-native:cleanrooms:ConfiguredTableJoinRequiredOption")
class ConfiguredTableJoinRequiredOption(builtins.str, Enum):
    QUERY_RUNNER = "QUERY_RUNNER"


@pulumi.type_token("aws-native:cleanrooms:ConfiguredTableScalarFunctions")
class ConfiguredTableScalarFunctions(builtins.str, Enum):
    TRUNC = "TRUNC"
    ABS = "ABS"
    CEILING = "CEILING"
    FLOOR = "FLOOR"
    LN = "LN"
    LOG = "LOG"
    ROUND = "ROUND"
    SQRT = "SQRT"
    CAST = "CAST"
    LOWER = "LOWER"
    RTRIM = "RTRIM"
    UPPER = "UPPER"
    COALESCE = "COALESCE"
    CONVERT = "CONVERT"
    CURRENT_DATE = "CURRENT_DATE"
    DATEADD = "DATEADD"
    EXTRACT = "EXTRACT"
    GETDATE = "GETDATE"
    SUBSTRING = "SUBSTRING"
    TO_CHAR = "TO_CHAR"
    TO_DATE = "TO_DATE"
    TO_NUMBER = "TO_NUMBER"
    TO_TIMESTAMP = "TO_TIMESTAMP"
    TRIM = "TRIM"


@pulumi.type_token("aws-native:cleanrooms:ConfiguredTableSelectedAnalysisMethod")
class ConfiguredTableSelectedAnalysisMethod(builtins.str, Enum):
    DIRECT_QUERY = "DIRECT_QUERY"
    DIRECT_JOB = "DIRECT_JOB"


@pulumi.type_token("aws-native:cleanrooms:IdMappingTableInputSourceType")
class IdMappingTableInputSourceType(builtins.str, Enum):
    """
    The type of the input source of the ID mapping table.
    """
    SOURCE = "SOURCE"
    TARGET = "TARGET"


@pulumi.type_token("aws-native:cleanrooms:IdNamespaceAssociationInputReferencePropertiesIdNamespaceType")
class IdNamespaceAssociationInputReferencePropertiesIdNamespaceType(builtins.str, Enum):
    """
    The ID namespace type for this ID namespace association.
    """
    SOURCE = "SOURCE"
    TARGET = "TARGET"


@pulumi.type_token("aws-native:cleanrooms:MembershipJobLogStatus")
class MembershipJobLogStatus(builtins.str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


@pulumi.type_token("aws-native:cleanrooms:MembershipQueryLogStatus")
class MembershipQueryLogStatus(builtins.str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


@pulumi.type_token("aws-native:cleanrooms:MembershipResultFormat")
class MembershipResultFormat(builtins.str, Enum):
    CSV = "CSV"
    PARQUET = "PARQUET"


@pulumi.type_token("aws-native:cleanrooms:PrivacyBudgetTemplateAutoRefresh")
class PrivacyBudgetTemplateAutoRefresh(builtins.str, Enum):
    """
    How often the privacy budget refreshes.

    > If you plan to regularly bring new data into the collaboration, use `CALENDAR_MONTH` to automatically get a new privacy budget for the collaboration every calendar month. Choosing this option allows arbitrary amounts of information to be revealed about rows of the data when repeatedly queried across refreshes. Avoid choosing this if the same rows will be repeatedly queried between privacy budget refreshes.
    """
    CALENDAR_MONTH = "CALENDAR_MONTH"
    NONE = "NONE"


@pulumi.type_token("aws-native:cleanrooms:PrivacyBudgetTemplatePrivacyBudgetType")
class PrivacyBudgetTemplatePrivacyBudgetType(builtins.str, Enum):
    """
    Specifies the type of the privacy budget template.
    """
    DIFFERENTIAL_PRIVACY = "DIFFERENTIAL_PRIVACY"
