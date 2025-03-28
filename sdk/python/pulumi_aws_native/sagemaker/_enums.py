# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AppResourceSpecInstanceType',
    'AppType',
    'ClusterDeepHealthCheckType',
    'ClusterNodeRecovery',
    'ClusterStatus',
    'DataQualityJobDefinitionBatchTransformInputS3DataDistributionType',
    'DataQualityJobDefinitionBatchTransformInputS3InputMode',
    'DataQualityJobDefinitionEndpointInputS3DataDistributionType',
    'DataQualityJobDefinitionEndpointInputS3InputMode',
    'DataQualityJobDefinitionS3OutputS3UploadMode',
    'DomainAppNetworkAccessType',
    'DomainAppSecurityGroupManagement',
    'DomainAppType',
    'DomainAuthMode',
    'DomainDockerSettingsEnableDockerAccess',
    'DomainLifecycleManagement',
    'DomainMlTools',
    'DomainRStudioServerProAppSettingsAccessStatus',
    'DomainRStudioServerProAppSettingsUserGroup',
    'DomainResourceSpecInstanceType',
    'DomainSettingsExecutionRoleIdentityConfig',
    'DomainSharingSettingsNotebookOutputOption',
    'DomainTagPropagation',
    'DomainUserSettingsStudioWebPortal',
    'FeatureGroupFeatureDefinitionFeatureType',
    'FeatureGroupStorageType',
    'FeatureGroupTableFormat',
    'FeatureGroupThroughputMode',
    'FeatureGroupUnit',
    'ImageVersionJobType',
    'ImageVersionProcessor',
    'ImageVersionVendorGuidance',
    'InferenceComponentCapacitySizeType',
    'InferenceComponentStatus',
    'InferenceExperimentDesiredState',
    'InferenceExperimentEndpointMetadataEndpointStatus',
    'InferenceExperimentModelInfrastructureConfigInfrastructureType',
    'InferenceExperimentStatus',
    'InferenceExperimentType',
    'MlflowTrackingServerTrackingServerSize',
    'ModelBiasJobDefinitionBatchTransformInputS3DataDistributionType',
    'ModelBiasJobDefinitionBatchTransformInputS3InputMode',
    'ModelBiasJobDefinitionEndpointInputS3DataDistributionType',
    'ModelBiasJobDefinitionEndpointInputS3InputMode',
    'ModelBiasJobDefinitionS3OutputS3UploadMode',
    'ModelCardBarChartMetricType',
    'ModelCardLinearGraphMetricType',
    'ModelCardMatrixMetricType',
    'ModelCardModelPackageDetailsModelApprovalStatus',
    'ModelCardModelPackageDetailsModelPackageStatus',
    'ModelCardObjectiveFunctionFunctionPropertiesFunction',
    'ModelCardProcessingStatus',
    'ModelCardRiskRating',
    'ModelCardSimpleMetricType',
    'ModelCardStatus',
    'ModelExplainabilityJobDefinitionBatchTransformInputS3DataDistributionType',
    'ModelExplainabilityJobDefinitionBatchTransformInputS3InputMode',
    'ModelExplainabilityJobDefinitionEndpointInputS3DataDistributionType',
    'ModelExplainabilityJobDefinitionEndpointInputS3InputMode',
    'ModelExplainabilityJobDefinitionS3OutputS3UploadMode',
    'ModelPackageGroupStatus',
    'ModelPackageModelApprovalStatus',
    'ModelPackageModelCardModelCardStatus',
    'ModelPackageS3DataSourceS3DataType',
    'ModelPackageS3ModelDataSourceCompressionType',
    'ModelPackageS3ModelDataSourceS3DataType',
    'ModelPackageSkipModelValidation',
    'ModelPackageStatus',
    'ModelPackageStatusItemStatus',
    'ModelPackageTransformInputCompressionType',
    'ModelPackageTransformInputSplitType',
    'ModelPackageTransformJobDefinitionBatchStrategy',
    'ModelPackageTransformOutputAssembleWith',
    'ModelQualityJobDefinitionBatchTransformInputS3DataDistributionType',
    'ModelQualityJobDefinitionBatchTransformInputS3InputMode',
    'ModelQualityJobDefinitionEndpointInputS3DataDistributionType',
    'ModelQualityJobDefinitionEndpointInputS3InputMode',
    'ModelQualityJobDefinitionProblemType',
    'ModelQualityJobDefinitionS3OutputS3UploadMode',
    'MonitoringScheduleBatchTransformInputS3DataDistributionType',
    'MonitoringScheduleBatchTransformInputS3InputMode',
    'MonitoringScheduleEndpointInputS3DataDistributionType',
    'MonitoringScheduleEndpointInputS3InputMode',
    'MonitoringScheduleMonitoringExecutionSummaryMonitoringExecutionStatus',
    'MonitoringScheduleMonitoringType',
    'MonitoringScheduleS3OutputS3UploadMode',
    'MonitoringScheduleStatus',
    'PartnerAppAuthType',
    'PartnerAppType',
    'ProjectStatus',
    'SpaceAppType',
    'SpaceResourceSpecInstanceType',
    'SpaceSharingSettingsSharingType',
    'StudioLifecycleConfigAppType',
    'UserProfileAppType',
    'UserProfileLifecycleManagement',
    'UserProfileMlTools',
    'UserProfileRStudioServerProAppSettingsAccessStatus',
    'UserProfileRStudioServerProAppSettingsUserGroup',
    'UserProfileResourceSpecInstanceType',
    'UserProfileSharingSettingsNotebookOutputOption',
    'UserProfileUserSettingsStudioWebPortal',
]


class AppResourceSpecInstanceType(str, Enum):
    """
    The instance type that the image version runs on.
    """
    SYSTEM = "system"
    ML_T3_MICRO = "ml.t3.micro"
    ML_T3_SMALL = "ml.t3.small"
    ML_T3_MEDIUM = "ml.t3.medium"
    ML_T3_LARGE = "ml.t3.large"
    ML_T3_XLARGE = "ml.t3.xlarge"
    ML_T32XLARGE = "ml.t3.2xlarge"
    ML_M5_LARGE = "ml.m5.large"
    ML_M5_XLARGE = "ml.m5.xlarge"
    ML_M52XLARGE = "ml.m5.2xlarge"
    ML_M54XLARGE = "ml.m5.4xlarge"
    ML_M58XLARGE = "ml.m5.8xlarge"
    ML_M512XLARGE = "ml.m5.12xlarge"
    ML_M516XLARGE = "ml.m5.16xlarge"
    ML_M524XLARGE = "ml.m5.24xlarge"
    ML_C5_LARGE = "ml.c5.large"
    ML_C5_XLARGE = "ml.c5.xlarge"
    ML_C52XLARGE = "ml.c5.2xlarge"
    ML_C54XLARGE = "ml.c5.4xlarge"
    ML_C59XLARGE = "ml.c5.9xlarge"
    ML_C512XLARGE = "ml.c5.12xlarge"
    ML_C518XLARGE = "ml.c5.18xlarge"
    ML_C524XLARGE = "ml.c5.24xlarge"
    ML_P32XLARGE = "ml.p3.2xlarge"
    ML_P38XLARGE = "ml.p3.8xlarge"
    ML_P316XLARGE = "ml.p3.16xlarge"
    ML_G4DN_XLARGE = "ml.g4dn.xlarge"
    ML_G4DN2XLARGE = "ml.g4dn.2xlarge"
    ML_G4DN4XLARGE = "ml.g4dn.4xlarge"
    ML_G4DN8XLARGE = "ml.g4dn.8xlarge"
    ML_G4DN12XLARGE = "ml.g4dn.12xlarge"
    ML_G4DN16XLARGE = "ml.g4dn.16xlarge"
    ML_R5_LARGE = "ml.r5.large"
    ML_R5_XLARGE = "ml.r5.xlarge"
    ML_R52XLARGE = "ml.r5.2xlarge"
    ML_R54XLARGE = "ml.r5.4xlarge"
    ML_R58XLARGE = "ml.r5.8xlarge"
    ML_R512XLARGE = "ml.r5.12xlarge"
    ML_R516XLARGE = "ml.r5.16xlarge"
    ML_R524XLARGE = "ml.r5.24xlarge"
    ML_P3DN24XLARGE = "ml.p3dn.24xlarge"
    ML_M5D_LARGE = "ml.m5d.large"
    ML_M5D_XLARGE = "ml.m5d.xlarge"
    ML_M5D2XLARGE = "ml.m5d.2xlarge"
    ML_M5D4XLARGE = "ml.m5d.4xlarge"
    ML_M5D8XLARGE = "ml.m5d.8xlarge"
    ML_M5D12XLARGE = "ml.m5d.12xlarge"
    ML_M5D16XLARGE = "ml.m5d.16xlarge"
    ML_M5D24XLARGE = "ml.m5d.24xlarge"
    ML_G5_XLARGE = "ml.g5.xlarge"
    ML_G52XLARGE = "ml.g5.2xlarge"
    ML_G54XLARGE = "ml.g5.4xlarge"
    ML_G58XLARGE = "ml.g5.8xlarge"
    ML_G512XLARGE = "ml.g5.12xlarge"
    ML_G516XLARGE = "ml.g5.16xlarge"
    ML_G524XLARGE = "ml.g5.24xlarge"
    ML_G548XLARGE = "ml.g5.48xlarge"
    ML_P4D24XLARGE = "ml.p4d.24xlarge"
    ML_P4DE24XLARGE = "ml.p4de.24xlarge"
    ML_GEOSPATIAL_INTERACTIVE = "ml.geospatial.interactive"
    ML_TRN12XLARGE = "ml.trn1.2xlarge"
    ML_TRN132XLARGE = "ml.trn1.32xlarge"
    ML_TRN1N32XLARGE = "ml.trn1n.32xlarge"


class AppType(str, Enum):
    """
    The type of app.
    """
    JUPYTER_SERVER = "JupyterServer"
    KERNEL_GATEWAY = "KernelGateway"
    R_STUDIO_SERVER_PRO = "RStudioServerPro"
    R_SESSION_GATEWAY = "RSessionGateway"
    CANVAS = "Canvas"


class ClusterDeepHealthCheckType(str, Enum):
    """
    The type of deep health check(s) to be performed on the instances in the SageMaker HyperPod cluster instance group.
    """
    INSTANCE_STRESS = "InstanceStress"
    INSTANCE_CONNECTIVITY = "InstanceConnectivity"


class ClusterNodeRecovery(str, Enum):
    """
    If node auto-recovery is set to true, faulty nodes will be replaced or rebooted when a failure is detected. If set to false, nodes will be labelled when a fault is detected.
    """
    AUTOMATIC = "Automatic"
    NONE = "None"


class ClusterStatus(str, Enum):
    """
    The status of the HyperPod Cluster.
    """
    CREATING = "Creating"
    DELETING = "Deleting"
    FAILED = "Failed"
    IN_SERVICE = "InService"
    ROLLING_BACK = "RollingBack"
    SYSTEM_UPDATING = "SystemUpdating"
    UPDATING = "Updating"


class DataQualityJobDefinitionBatchTransformInputS3DataDistributionType(str, Enum):
    """
    Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
    """
    FULLY_REPLICATED = "FullyReplicated"
    SHARDED_BY_S3_KEY = "ShardedByS3Key"


class DataQualityJobDefinitionBatchTransformInputS3InputMode(str, Enum):
    """
    Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
    """
    PIPE = "Pipe"
    FILE = "File"


class DataQualityJobDefinitionEndpointInputS3DataDistributionType(str, Enum):
    """
    Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
    """
    FULLY_REPLICATED = "FullyReplicated"
    SHARDED_BY_S3_KEY = "ShardedByS3Key"


class DataQualityJobDefinitionEndpointInputS3InputMode(str, Enum):
    """
    Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
    """
    PIPE = "Pipe"
    FILE = "File"


class DataQualityJobDefinitionS3OutputS3UploadMode(str, Enum):
    """
    Whether to upload the results of the monitoring job continuously or after the job completes.
    """
    CONTINUOUS = "Continuous"
    END_OF_JOB = "EndOfJob"


class DomainAppNetworkAccessType(str, Enum):
    """
    Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.
    """
    PUBLIC_INTERNET_ONLY = "PublicInternetOnly"
    VPC_ONLY = "VpcOnly"


class DomainAppSecurityGroupManagement(str, Enum):
    """
    The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.
    """
    SERVICE = "Service"
    CUSTOMER = "Customer"


class DomainAppType(str, Enum):
    JUPYTER_SERVER = "JupyterServer"
    TENSOR_BOARD = "TensorBoard"
    R_STUDIO_SERVER_PRO = "RStudioServerPro"
    JUPYTER_LAB = "JupyterLab"
    CODE_EDITOR = "CodeEditor"
    DETAILED_PROFILER = "DetailedProfiler"
    CANVAS = "Canvas"


class DomainAuthMode(str, Enum):
    """
    The mode of authentication that members use to access the domain.
    """
    SSO = "SSO"
    IAM = "IAM"


class DomainDockerSettingsEnableDockerAccess(str, Enum):
    """
    The flag to enable/disable docker-proxy server
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class DomainLifecycleManagement(str, Enum):
    """
    A flag to enable/disable AppLifecycleManagement settings
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class DomainMlTools(str, Enum):
    DATA_WRANGLER = "DataWrangler"
    FEATURE_STORE = "FeatureStore"
    EMR_CLUSTERS = "EmrClusters"
    AUTO_ML = "AutoMl"
    EXPERIMENTS = "Experiments"
    TRAINING = "Training"
    MODEL_EVALUATION = "ModelEvaluation"
    PIPELINES = "Pipelines"
    MODELS = "Models"
    JUMP_START = "JumpStart"
    INFERENCE_RECOMMENDER = "InferenceRecommender"
    ENDPOINTS = "Endpoints"
    PROJECTS = "Projects"
    INFERENCE_OPTIMIZATION = "InferenceOptimization"
    HYPER_POD_CLUSTERS = "HyperPodClusters"
    COMET = "Comet"
    DEEPCHECKS_LLM_EVALUATION = "DeepchecksLLMEvaluation"
    FIDDLER = "Fiddler"
    LAKERA_GUARD = "LakeraGuard"
    PERFORMANCE_EVALUATION = "PerformanceEvaluation"


class DomainRStudioServerProAppSettingsAccessStatus(str, Enum):
    """
    Indicates whether the current user has access to the RStudioServerPro app.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class DomainRStudioServerProAppSettingsUserGroup(str, Enum):
    """
    The level of permissions that the user has within the RStudioServerPro app. This value defaults to User. The Admin value allows the user access to the RStudio Administrative Dashboard.
    """
    R_STUDIO_ADMIN = "R_STUDIO_ADMIN"
    R_STUDIO_USER = "R_STUDIO_USER"


class DomainResourceSpecInstanceType(str, Enum):
    """
    The instance type that the image version runs on.
    """
    SYSTEM = "system"
    ML_T3_MICRO = "ml.t3.micro"
    ML_T3_SMALL = "ml.t3.small"
    ML_T3_MEDIUM = "ml.t3.medium"
    ML_T3_LARGE = "ml.t3.large"
    ML_T3_XLARGE = "ml.t3.xlarge"
    ML_T32XLARGE = "ml.t3.2xlarge"
    ML_M5_LARGE = "ml.m5.large"
    ML_M5_XLARGE = "ml.m5.xlarge"
    ML_M52XLARGE = "ml.m5.2xlarge"
    ML_M54XLARGE = "ml.m5.4xlarge"
    ML_M58XLARGE = "ml.m5.8xlarge"
    ML_M512XLARGE = "ml.m5.12xlarge"
    ML_M516XLARGE = "ml.m5.16xlarge"
    ML_M524XLARGE = "ml.m5.24xlarge"
    ML_C5_LARGE = "ml.c5.large"
    ML_C5_XLARGE = "ml.c5.xlarge"
    ML_C52XLARGE = "ml.c5.2xlarge"
    ML_C54XLARGE = "ml.c5.4xlarge"
    ML_C59XLARGE = "ml.c5.9xlarge"
    ML_C512XLARGE = "ml.c5.12xlarge"
    ML_C518XLARGE = "ml.c5.18xlarge"
    ML_C524XLARGE = "ml.c5.24xlarge"
    ML_P32XLARGE = "ml.p3.2xlarge"
    ML_P38XLARGE = "ml.p3.8xlarge"
    ML_P316XLARGE = "ml.p3.16xlarge"
    ML_G4DN_XLARGE = "ml.g4dn.xlarge"
    ML_G4DN2XLARGE = "ml.g4dn.2xlarge"
    ML_G4DN4XLARGE = "ml.g4dn.4xlarge"
    ML_G4DN8XLARGE = "ml.g4dn.8xlarge"
    ML_G4DN12XLARGE = "ml.g4dn.12xlarge"
    ML_G4DN16XLARGE = "ml.g4dn.16xlarge"
    ML_R5_LARGE = "ml.r5.large"
    ML_R5_XLARGE = "ml.r5.xlarge"
    ML_R52XLARGE = "ml.r5.2xlarge"
    ML_R54XLARGE = "ml.r5.4xlarge"
    ML_R58XLARGE = "ml.r5.8xlarge"
    ML_R512XLARGE = "ml.r5.12xlarge"
    ML_R516XLARGE = "ml.r5.16xlarge"
    ML_R524XLARGE = "ml.r5.24xlarge"
    ML_P3DN24XLARGE = "ml.p3dn.24xlarge"
    ML_M5D_LARGE = "ml.m5d.large"
    ML_M5D_XLARGE = "ml.m5d.xlarge"
    ML_M5D2XLARGE = "ml.m5d.2xlarge"
    ML_M5D4XLARGE = "ml.m5d.4xlarge"
    ML_M5D8XLARGE = "ml.m5d.8xlarge"
    ML_M5D12XLARGE = "ml.m5d.12xlarge"
    ML_M5D16XLARGE = "ml.m5d.16xlarge"
    ML_M5D24XLARGE = "ml.m5d.24xlarge"
    ML_G5_XLARGE = "ml.g5.xlarge"
    ML_G52XLARGE = "ml.g5.2xlarge"
    ML_G54XLARGE = "ml.g5.4xlarge"
    ML_G58XLARGE = "ml.g5.8xlarge"
    ML_G512XLARGE = "ml.g5.12xlarge"
    ML_G516XLARGE = "ml.g5.16xlarge"
    ML_G524XLARGE = "ml.g5.24xlarge"
    ML_G548XLARGE = "ml.g5.48xlarge"
    ML_P4D24XLARGE = "ml.p4d.24xlarge"
    ML_P4DE24XLARGE = "ml.p4de.24xlarge"
    ML_GEOSPATIAL_INTERACTIVE = "ml.geospatial.interactive"
    ML_TRN12XLARGE = "ml.trn1.2xlarge"
    ML_TRN132XLARGE = "ml.trn1.32xlarge"
    ML_TRN1N32XLARGE = "ml.trn1n.32xlarge"


class DomainSettingsExecutionRoleIdentityConfig(str, Enum):
    """
    The configuration for attaching a SageMaker user profile name to the execution role as a sts:SourceIdentity key.
    """
    USER_PROFILE_NAME = "USER_PROFILE_NAME"
    DISABLED = "DISABLED"


class DomainSharingSettingsNotebookOutputOption(str, Enum):
    """
    Whether to include the notebook cell output when sharing the notebook. The default is Disabled.
    """
    ALLOWED = "Allowed"
    DISABLED = "Disabled"


class DomainTagPropagation(str, Enum):
    """
    Indicates whether the tags added to Domain, User Profile and Space entity is propagated to all SageMaker resources.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class DomainUserSettingsStudioWebPortal(str, Enum):
    """
    Indicates whether the Studio experience is available to users. If not, users cannot access Studio.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class FeatureGroupFeatureDefinitionFeatureType(str, Enum):
    """
    The value type of a feature. Valid values are Integral, Fractional, or String.
    """
    INTEGRAL = "Integral"
    FRACTIONAL = "Fractional"
    STRING = "String"


class FeatureGroupStorageType(str, Enum):
    STANDARD = "Standard"
    IN_MEMORY = "InMemory"


class FeatureGroupTableFormat(str, Enum):
    """
    Format for the offline store feature group. Iceberg is the optimal format for feature groups shared between offline and online stores.
    """
    ICEBERG = "Iceberg"
    GLUE = "Glue"


class FeatureGroupThroughputMode(str, Enum):
    """
    Throughput mode configuration of the feature group
    """
    ON_DEMAND = "OnDemand"
    PROVISIONED = "Provisioned"


class FeatureGroupUnit(str, Enum):
    """
    Unit of ttl configuration
    """
    SECONDS = "Seconds"
    MINUTES = "Minutes"
    HOURS = "Hours"
    DAYS = "Days"
    WEEKS = "Weeks"


class ImageVersionJobType(str, Enum):
    """
    Indicates SageMaker job type compatibility.
    """
    TRAINING = "TRAINING"
    INFERENCE = "INFERENCE"
    NOTEBOOK_KERNEL = "NOTEBOOK_KERNEL"


class ImageVersionProcessor(str, Enum):
    """
    Indicates CPU or GPU compatibility.
    """
    CPU = "CPU"
    GPU = "GPU"


class ImageVersionVendorGuidance(str, Enum):
    """
    The availability of the image version specified by the maintainer.
    """
    NOT_PROVIDED = "NOT_PROVIDED"
    STABLE = "STABLE"
    TO_BE_ARCHIVED = "TO_BE_ARCHIVED"
    ARCHIVED = "ARCHIVED"


class InferenceComponentCapacitySizeType(str, Enum):
    COPY_COUNT = "COPY_COUNT"
    CAPACITY_PERCENT = "CAPACITY_PERCENT"


class InferenceComponentStatus(str, Enum):
    IN_SERVICE = "InService"
    CREATING = "Creating"
    UPDATING = "Updating"
    FAILED = "Failed"
    DELETING = "Deleting"


class InferenceExperimentDesiredState(str, Enum):
    """
    The desired state of the experiment after starting or stopping operation.
    """
    RUNNING = "Running"
    COMPLETED = "Completed"
    CANCELLED = "Cancelled"


class InferenceExperimentEndpointMetadataEndpointStatus(str, Enum):
    """
    The status of the endpoint. For possible values of the status of an endpoint.
    """
    CREATING = "Creating"
    UPDATING = "Updating"
    SYSTEM_UPDATING = "SystemUpdating"
    ROLLING_BACK = "RollingBack"
    IN_SERVICE = "InService"
    OUT_OF_SERVICE = "OutOfService"
    DELETING = "Deleting"
    FAILED = "Failed"


class InferenceExperimentModelInfrastructureConfigInfrastructureType(str, Enum):
    """
    The type of the inference experiment that you want to run.
    """
    REAL_TIME_INFERENCE = "RealTimeInference"


class InferenceExperimentStatus(str, Enum):
    """
    The status of the inference experiment.
    """
    CREATING = "Creating"
    CREATED = "Created"
    UPDATING = "Updating"
    STARTING = "Starting"
    STOPPING = "Stopping"
    RUNNING = "Running"
    COMPLETED = "Completed"
    CANCELLED = "Cancelled"


class InferenceExperimentType(str, Enum):
    """
    The type of the inference experiment that you want to run.
    """
    SHADOW_MODE = "ShadowMode"


class MlflowTrackingServerTrackingServerSize(str, Enum):
    """
    The size of the MLFlow Tracking Server.
    """
    SMALL = "Small"
    MEDIUM = "Medium"
    LARGE = "Large"


class ModelBiasJobDefinitionBatchTransformInputS3DataDistributionType(str, Enum):
    """
    Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
    """
    FULLY_REPLICATED = "FullyReplicated"
    SHARDED_BY_S3_KEY = "ShardedByS3Key"


class ModelBiasJobDefinitionBatchTransformInputS3InputMode(str, Enum):
    """
    Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
    """
    PIPE = "Pipe"
    FILE = "File"


class ModelBiasJobDefinitionEndpointInputS3DataDistributionType(str, Enum):
    """
    Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
    """
    FULLY_REPLICATED = "FullyReplicated"
    SHARDED_BY_S3_KEY = "ShardedByS3Key"


class ModelBiasJobDefinitionEndpointInputS3InputMode(str, Enum):
    """
    Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
    """
    PIPE = "Pipe"
    FILE = "File"


class ModelBiasJobDefinitionS3OutputS3UploadMode(str, Enum):
    """
    Whether to upload the results of the monitoring job continuously or after the job completes.
    """
    CONTINUOUS = "Continuous"
    END_OF_JOB = "EndOfJob"


class ModelCardBarChartMetricType(str, Enum):
    BAR_CHART = "bar_chart"


class ModelCardLinearGraphMetricType(str, Enum):
    LINEAR_GRAPH = "linear_graph"


class ModelCardMatrixMetricType(str, Enum):
    MATRIX = "matrix"


class ModelCardModelPackageDetailsModelApprovalStatus(str, Enum):
    """
    Current approval status of model package
    """
    APPROVED = "Approved"
    REJECTED = "Rejected"
    PENDING_MANUAL_APPROVAL = "PendingManualApproval"


class ModelCardModelPackageDetailsModelPackageStatus(str, Enum):
    """
    Current status of model package
    """
    PENDING = "Pending"
    IN_PROGRESS = "InProgress"
    COMPLETED = "Completed"
    FAILED = "Failed"
    DELETING = "Deleting"


class ModelCardObjectiveFunctionFunctionPropertiesFunction(str, Enum):
    MAXIMIZE = "Maximize"
    MINIMIZE = "Minimize"


class ModelCardProcessingStatus(str, Enum):
    """
    The processing status of model card deletion. The ModelCardProcessingStatus updates throughout the different deletion steps.
    """
    UNSET_VALUE = "UnsetValue"
    DELETE_IN_PROGRESS = "DeleteInProgress"
    DELETE_PENDING = "DeletePending"
    CONTENT_DELETED = "ContentDeleted"
    EXPORT_JOBS_DELETED = "ExportJobsDeleted"
    DELETE_COMPLETED = "DeleteCompleted"
    DELETE_FAILED = "DeleteFailed"


class ModelCardRiskRating(str, Enum):
    """
    Risk rating of model.
    """
    HIGH = "High"
    MEDIUM = "Medium"
    LOW = "Low"
    UNKNOWN = "Unknown"


class ModelCardSimpleMetricType(str, Enum):
    NUMBER = "number"
    STRING = "string"
    BOOLEAN = "boolean"


class ModelCardStatus(str, Enum):
    """
    The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.
    """
    DRAFT = "Draft"
    PENDING_REVIEW = "PendingReview"
    APPROVED = "Approved"
    ARCHIVED = "Archived"


class ModelExplainabilityJobDefinitionBatchTransformInputS3DataDistributionType(str, Enum):
    """
    Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
    """
    FULLY_REPLICATED = "FullyReplicated"
    SHARDED_BY_S3_KEY = "ShardedByS3Key"


class ModelExplainabilityJobDefinitionBatchTransformInputS3InputMode(str, Enum):
    """
    Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
    """
    PIPE = "Pipe"
    FILE = "File"


class ModelExplainabilityJobDefinitionEndpointInputS3DataDistributionType(str, Enum):
    """
    Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
    """
    FULLY_REPLICATED = "FullyReplicated"
    SHARDED_BY_S3_KEY = "ShardedByS3Key"


class ModelExplainabilityJobDefinitionEndpointInputS3InputMode(str, Enum):
    """
    Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
    """
    PIPE = "Pipe"
    FILE = "File"


class ModelExplainabilityJobDefinitionS3OutputS3UploadMode(str, Enum):
    """
    Whether to upload the results of the monitoring job continuously or after the job completes.
    """
    CONTINUOUS = "Continuous"
    END_OF_JOB = "EndOfJob"


class ModelPackageGroupStatus(str, Enum):
    """
    The status of a modelpackage group job.
    """
    PENDING = "Pending"
    IN_PROGRESS = "InProgress"
    COMPLETED = "Completed"
    FAILED = "Failed"
    DELETING = "Deleting"
    DELETE_FAILED = "DeleteFailed"


class ModelPackageModelApprovalStatus(str, Enum):
    """
    The approval status of the model package.
    """
    APPROVED = "Approved"
    REJECTED = "Rejected"
    PENDING_MANUAL_APPROVAL = "PendingManualApproval"


class ModelPackageModelCardModelCardStatus(str, Enum):
    """
    The approval status of the model card within your organization.
    """
    DRAFT = "Draft"
    PENDING_REVIEW = "PendingReview"
    APPROVED = "Approved"
    ARCHIVED = "Archived"


class ModelPackageS3DataSourceS3DataType(str, Enum):
    """
    The S3 Data Source Type
    """
    MANIFEST_FILE = "ManifestFile"
    S3_PREFIX = "S3Prefix"
    AUGMENTED_MANIFEST_FILE = "AugmentedManifestFile"


class ModelPackageS3ModelDataSourceCompressionType(str, Enum):
    """
    Specifies how the ML model data is prepared.
    """
    NONE = "None"
    GZIP = "Gzip"


class ModelPackageS3ModelDataSourceS3DataType(str, Enum):
    """
    Specifies the type of ML model data to deploy.
    """
    S3_PREFIX = "S3Prefix"
    S3_OBJECT = "S3Object"


class ModelPackageSkipModelValidation(str, Enum):
    """
    Indicates if you want to skip model validation.
    """
    NONE = "None"
    ALL = "All"


class ModelPackageStatus(str, Enum):
    """
    The current status of the model package.
    """
    PENDING = "Pending"
    DELETING = "Deleting"
    IN_PROGRESS = "InProgress"
    COMPLETED = "Completed"
    FAILED = "Failed"


class ModelPackageStatusItemStatus(str, Enum):
    """
    The current status.
    """
    NOT_STARTED = "NotStarted"
    FAILED = "Failed"
    IN_PROGRESS = "InProgress"
    COMPLETED = "Completed"


class ModelPackageTransformInputCompressionType(str, Enum):
    """
    If your transform data is compressed, specify the compression type. Amazon SageMaker automatically decompresses the data for the transform job accordingly. The default value is None.
    """
    NONE = "None"
    GZIP = "Gzip"


class ModelPackageTransformInputSplitType(str, Enum):
    """
    The method to use to split the transform job's data files into smaller batches. 
    """
    NONE = "None"
    TF_RECORD = "TFRecord"
    LINE = "Line"
    RECORD_IO = "RecordIO"


class ModelPackageTransformJobDefinitionBatchStrategy(str, Enum):
    """
    A string that determines the number of records included in a single mini-batch.
    """
    MULTI_RECORD = "MultiRecord"
    SINGLE_RECORD = "SingleRecord"


class ModelPackageTransformOutputAssembleWith(str, Enum):
    """
    Defines how to assemble the results of the transform job as a single S3 object.
    """
    NONE = "None"
    LINE = "Line"


class ModelQualityJobDefinitionBatchTransformInputS3DataDistributionType(str, Enum):
    """
    Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
    """
    FULLY_REPLICATED = "FullyReplicated"
    SHARDED_BY_S3_KEY = "ShardedByS3Key"


class ModelQualityJobDefinitionBatchTransformInputS3InputMode(str, Enum):
    """
    Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
    """
    PIPE = "Pipe"
    FILE = "File"


class ModelQualityJobDefinitionEndpointInputS3DataDistributionType(str, Enum):
    """
    Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
    """
    FULLY_REPLICATED = "FullyReplicated"
    SHARDED_BY_S3_KEY = "ShardedByS3Key"


class ModelQualityJobDefinitionEndpointInputS3InputMode(str, Enum):
    """
    Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
    """
    PIPE = "Pipe"
    FILE = "File"


class ModelQualityJobDefinitionProblemType(str, Enum):
    """
    The status of the monitoring job.
    """
    BINARY_CLASSIFICATION = "BinaryClassification"
    MULTICLASS_CLASSIFICATION = "MulticlassClassification"
    REGRESSION = "Regression"


class ModelQualityJobDefinitionS3OutputS3UploadMode(str, Enum):
    """
    Whether to upload the results of the monitoring job continuously or after the job completes.
    """
    CONTINUOUS = "Continuous"
    END_OF_JOB = "EndOfJob"


class MonitoringScheduleBatchTransformInputS3DataDistributionType(str, Enum):
    """
    Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
    """
    FULLY_REPLICATED = "FullyReplicated"
    SHARDED_BY_S3_KEY = "ShardedByS3Key"


class MonitoringScheduleBatchTransformInputS3InputMode(str, Enum):
    """
    Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
    """
    PIPE = "Pipe"
    FILE = "File"


class MonitoringScheduleEndpointInputS3DataDistributionType(str, Enum):
    """
    Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
    """
    FULLY_REPLICATED = "FullyReplicated"
    SHARDED_BY_S3_KEY = "ShardedByS3Key"


class MonitoringScheduleEndpointInputS3InputMode(str, Enum):
    """
    Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
    """
    PIPE = "Pipe"
    FILE = "File"


class MonitoringScheduleMonitoringExecutionSummaryMonitoringExecutionStatus(str, Enum):
    """
    The status of the monitoring job.
    """
    PENDING = "Pending"
    COMPLETED = "Completed"
    COMPLETED_WITH_VIOLATIONS = "CompletedWithViolations"
    IN_PROGRESS = "InProgress"
    FAILED = "Failed"
    STOPPING = "Stopping"
    STOPPED = "Stopped"


class MonitoringScheduleMonitoringType(str, Enum):
    """
    The type of monitoring job.
    """
    DATA_QUALITY = "DataQuality"
    MODEL_QUALITY = "ModelQuality"
    MODEL_BIAS = "ModelBias"
    MODEL_EXPLAINABILITY = "ModelExplainability"


class MonitoringScheduleS3OutputS3UploadMode(str, Enum):
    """
    Whether to upload the results of the monitoring job continuously or after the job completes.
    """
    CONTINUOUS = "Continuous"
    END_OF_JOB = "EndOfJob"


class MonitoringScheduleStatus(str, Enum):
    """
    The status of a schedule job.
    """
    PENDING = "Pending"
    FAILED = "Failed"
    SCHEDULED = "Scheduled"
    STOPPED = "Stopped"


class PartnerAppAuthType(str, Enum):
    """
    The Auth type of PartnerApp.
    """
    IAM = "IAM"


class PartnerAppType(str, Enum):
    """
    The type of PartnerApp.
    """
    LAKERA_GUARD = "lakera-guard"
    COMET = "comet"
    DEEPCHECKS_LLM_EVALUATION = "deepchecks-llm-evaluation"
    FIDDLER = "fiddler"


class ProjectStatus(str, Enum):
    """
    The status of a project.
    """
    PENDING = "Pending"
    CREATE_IN_PROGRESS = "CreateInProgress"
    CREATE_COMPLETED = "CreateCompleted"
    CREATE_FAILED = "CreateFailed"
    DELETE_IN_PROGRESS = "DeleteInProgress"
    DELETE_FAILED = "DeleteFailed"
    DELETE_COMPLETED = "DeleteCompleted"


class SpaceAppType(str, Enum):
    JUPYTER_SERVER = "JupyterServer"
    KERNEL_GATEWAY = "KernelGateway"
    TENSOR_BOARD = "TensorBoard"
    R_STUDIO_SERVER_PRO = "RStudioServerPro"
    R_SESSION_GATEWAY = "RSessionGateway"
    JUPYTER_LAB = "JupyterLab"
    CODE_EDITOR = "CodeEditor"


class SpaceResourceSpecInstanceType(str, Enum):
    """
    The instance type that the image version runs on.
    """
    SYSTEM = "system"
    ML_T3_MICRO = "ml.t3.micro"
    ML_T3_SMALL = "ml.t3.small"
    ML_T3_MEDIUM = "ml.t3.medium"
    ML_T3_LARGE = "ml.t3.large"
    ML_T3_XLARGE = "ml.t3.xlarge"
    ML_T32XLARGE = "ml.t3.2xlarge"
    ML_M5_LARGE = "ml.m5.large"
    ML_M5_XLARGE = "ml.m5.xlarge"
    ML_M52XLARGE = "ml.m5.2xlarge"
    ML_M54XLARGE = "ml.m5.4xlarge"
    ML_M58XLARGE = "ml.m5.8xlarge"
    ML_M512XLARGE = "ml.m5.12xlarge"
    ML_M516XLARGE = "ml.m5.16xlarge"
    ML_M524XLARGE = "ml.m5.24xlarge"
    ML_C5_LARGE = "ml.c5.large"
    ML_C5_XLARGE = "ml.c5.xlarge"
    ML_C52XLARGE = "ml.c5.2xlarge"
    ML_C54XLARGE = "ml.c5.4xlarge"
    ML_C59XLARGE = "ml.c5.9xlarge"
    ML_C512XLARGE = "ml.c5.12xlarge"
    ML_C518XLARGE = "ml.c5.18xlarge"
    ML_C524XLARGE = "ml.c5.24xlarge"
    ML_P32XLARGE = "ml.p3.2xlarge"
    ML_P38XLARGE = "ml.p3.8xlarge"
    ML_P316XLARGE = "ml.p3.16xlarge"
    ML_G4DN_XLARGE = "ml.g4dn.xlarge"
    ML_G4DN2XLARGE = "ml.g4dn.2xlarge"
    ML_G4DN4XLARGE = "ml.g4dn.4xlarge"
    ML_G4DN8XLARGE = "ml.g4dn.8xlarge"
    ML_G4DN12XLARGE = "ml.g4dn.12xlarge"
    ML_G4DN16XLARGE = "ml.g4dn.16xlarge"
    ML_R5_LARGE = "ml.r5.large"
    ML_R5_XLARGE = "ml.r5.xlarge"
    ML_R52XLARGE = "ml.r5.2xlarge"
    ML_R54XLARGE = "ml.r5.4xlarge"
    ML_R58XLARGE = "ml.r5.8xlarge"
    ML_R512XLARGE = "ml.r5.12xlarge"
    ML_R516XLARGE = "ml.r5.16xlarge"
    ML_R524XLARGE = "ml.r5.24xlarge"
    ML_P3DN24XLARGE = "ml.p3dn.24xlarge"
    ML_M5D_LARGE = "ml.m5d.large"
    ML_M5D_XLARGE = "ml.m5d.xlarge"
    ML_M5D2XLARGE = "ml.m5d.2xlarge"
    ML_M5D4XLARGE = "ml.m5d.4xlarge"
    ML_M5D8XLARGE = "ml.m5d.8xlarge"
    ML_M5D12XLARGE = "ml.m5d.12xlarge"
    ML_M5D16XLARGE = "ml.m5d.16xlarge"
    ML_M5D24XLARGE = "ml.m5d.24xlarge"
    ML_G5_XLARGE = "ml.g5.xlarge"
    ML_G52XLARGE = "ml.g5.2xlarge"
    ML_G54XLARGE = "ml.g5.4xlarge"
    ML_G58XLARGE = "ml.g5.8xlarge"
    ML_G512XLARGE = "ml.g5.12xlarge"
    ML_G516XLARGE = "ml.g5.16xlarge"
    ML_G524XLARGE = "ml.g5.24xlarge"
    ML_G548XLARGE = "ml.g5.48xlarge"
    ML_P4D24XLARGE = "ml.p4d.24xlarge"
    ML_P4DE24XLARGE = "ml.p4de.24xlarge"
    ML_GEOSPATIAL_INTERACTIVE = "ml.geospatial.interactive"
    ML_TRN12XLARGE = "ml.trn1.2xlarge"
    ML_TRN132XLARGE = "ml.trn1.32xlarge"
    ML_TRN1N32XLARGE = "ml.trn1n.32xlarge"


class SpaceSharingSettingsSharingType(str, Enum):
    """
    Specifies the sharing type of the space.
    """
    PRIVATE = "Private"
    SHARED = "Shared"


class StudioLifecycleConfigAppType(str, Enum):
    """
    The App type that the Lifecycle Configuration is attached to.
    """
    JUPYTER_SERVER = "JupyterServer"
    KERNEL_GATEWAY = "KernelGateway"
    CODE_EDITOR = "CodeEditor"
    JUPYTER_LAB = "JupyterLab"


class UserProfileAppType(str, Enum):
    JUPYTER_SERVER = "JupyterServer"
    TENSOR_BOARD = "TensorBoard"
    R_STUDIO_SERVER_PRO = "RStudioServerPro"
    JUPYTER_LAB = "JupyterLab"
    CODE_EDITOR = "CodeEditor"
    DETAILED_PROFILER = "DetailedProfiler"
    CANVAS = "Canvas"


class UserProfileLifecycleManagement(str, Enum):
    """
    A flag to enable/disable AppLifecycleManagement settings
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class UserProfileMlTools(str, Enum):
    DATA_WRANGLER = "DataWrangler"
    FEATURE_STORE = "FeatureStore"
    EMR_CLUSTERS = "EmrClusters"
    AUTO_ML = "AutoMl"
    EXPERIMENTS = "Experiments"
    TRAINING = "Training"
    MODEL_EVALUATION = "ModelEvaluation"
    PIPELINES = "Pipelines"
    MODELS = "Models"
    JUMP_START = "JumpStart"
    INFERENCE_RECOMMENDER = "InferenceRecommender"
    ENDPOINTS = "Endpoints"
    PROJECTS = "Projects"
    INFERENCE_OPTIMIZATION = "InferenceOptimization"
    HYPER_POD_CLUSTERS = "HyperPodClusters"
    COMET = "Comet"
    DEEPCHECKS_LLM_EVALUATION = "DeepchecksLLMEvaluation"
    FIDDLER = "Fiddler"
    LAKERA_GUARD = "LakeraGuard"
    PERFORMANCE_EVALUATION = "PerformanceEvaluation"


class UserProfileRStudioServerProAppSettingsAccessStatus(str, Enum):
    """
    Indicates whether the current user has access to the RStudioServerPro app.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class UserProfileRStudioServerProAppSettingsUserGroup(str, Enum):
    """
    The level of permissions that the user has within the RStudioServerPro app. This value defaults to User. The Admin value allows the user access to the RStudio Administrative Dashboard.
    """
    R_STUDIO_ADMIN = "R_STUDIO_ADMIN"
    R_STUDIO_USER = "R_STUDIO_USER"


class UserProfileResourceSpecInstanceType(str, Enum):
    """
    The instance type that the image version runs on.
    """
    SYSTEM = "system"
    ML_T3_MICRO = "ml.t3.micro"
    ML_T3_SMALL = "ml.t3.small"
    ML_T3_MEDIUM = "ml.t3.medium"
    ML_T3_LARGE = "ml.t3.large"
    ML_T3_XLARGE = "ml.t3.xlarge"
    ML_T32XLARGE = "ml.t3.2xlarge"
    ML_M5_LARGE = "ml.m5.large"
    ML_M5_XLARGE = "ml.m5.xlarge"
    ML_M52XLARGE = "ml.m5.2xlarge"
    ML_M54XLARGE = "ml.m5.4xlarge"
    ML_M58XLARGE = "ml.m5.8xlarge"
    ML_M512XLARGE = "ml.m5.12xlarge"
    ML_M516XLARGE = "ml.m5.16xlarge"
    ML_M524XLARGE = "ml.m5.24xlarge"
    ML_C5_LARGE = "ml.c5.large"
    ML_C5_XLARGE = "ml.c5.xlarge"
    ML_C52XLARGE = "ml.c5.2xlarge"
    ML_C54XLARGE = "ml.c5.4xlarge"
    ML_C59XLARGE = "ml.c5.9xlarge"
    ML_C512XLARGE = "ml.c5.12xlarge"
    ML_C518XLARGE = "ml.c5.18xlarge"
    ML_C524XLARGE = "ml.c5.24xlarge"
    ML_P32XLARGE = "ml.p3.2xlarge"
    ML_P38XLARGE = "ml.p3.8xlarge"
    ML_P316XLARGE = "ml.p3.16xlarge"
    ML_G4DN_XLARGE = "ml.g4dn.xlarge"
    ML_G4DN2XLARGE = "ml.g4dn.2xlarge"
    ML_G4DN4XLARGE = "ml.g4dn.4xlarge"
    ML_G4DN8XLARGE = "ml.g4dn.8xlarge"
    ML_G4DN12XLARGE = "ml.g4dn.12xlarge"
    ML_G4DN16XLARGE = "ml.g4dn.16xlarge"
    ML_R5_LARGE = "ml.r5.large"
    ML_R5_XLARGE = "ml.r5.xlarge"
    ML_R52XLARGE = "ml.r5.2xlarge"
    ML_R54XLARGE = "ml.r5.4xlarge"
    ML_R58XLARGE = "ml.r5.8xlarge"
    ML_R512XLARGE = "ml.r5.12xlarge"
    ML_R516XLARGE = "ml.r5.16xlarge"
    ML_R524XLARGE = "ml.r5.24xlarge"
    ML_P3DN24XLARGE = "ml.p3dn.24xlarge"
    ML_M5D_LARGE = "ml.m5d.large"
    ML_M5D_XLARGE = "ml.m5d.xlarge"
    ML_M5D2XLARGE = "ml.m5d.2xlarge"
    ML_M5D4XLARGE = "ml.m5d.4xlarge"
    ML_M5D8XLARGE = "ml.m5d.8xlarge"
    ML_M5D12XLARGE = "ml.m5d.12xlarge"
    ML_M5D16XLARGE = "ml.m5d.16xlarge"
    ML_M5D24XLARGE = "ml.m5d.24xlarge"
    ML_G5_XLARGE = "ml.g5.xlarge"
    ML_G52XLARGE = "ml.g5.2xlarge"
    ML_G54XLARGE = "ml.g5.4xlarge"
    ML_G58XLARGE = "ml.g5.8xlarge"
    ML_G512XLARGE = "ml.g5.12xlarge"
    ML_G516XLARGE = "ml.g5.16xlarge"
    ML_G524XLARGE = "ml.g5.24xlarge"
    ML_G548XLARGE = "ml.g5.48xlarge"
    ML_P4D24XLARGE = "ml.p4d.24xlarge"
    ML_P4DE24XLARGE = "ml.p4de.24xlarge"
    ML_GEOSPATIAL_INTERACTIVE = "ml.geospatial.interactive"
    ML_TRN12XLARGE = "ml.trn1.2xlarge"
    ML_TRN132XLARGE = "ml.trn1.32xlarge"
    ML_TRN1N32XLARGE = "ml.trn1n.32xlarge"


class UserProfileSharingSettingsNotebookOutputOption(str, Enum):
    """
    Whether to include the notebook cell output when sharing the notebook. The default is Disabled.
    """
    ALLOWED = "Allowed"
    DISABLED = "Disabled"


class UserProfileUserSettingsStudioWebPortal(str, Enum):
    """
    Indicates whether the Studio experience is available to users. If not, users cannot access Studio.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"
