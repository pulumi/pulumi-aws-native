// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MWAA
{
    public static class GetEnvironment
    {
        /// <summary>
        /// Resource schema for AWS::MWAA::Environment
        /// </summary>
        public static Task<GetEnvironmentResult> InvokeAsync(GetEnvironmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentResult>("aws-native:mwaa:getEnvironment", args ?? new GetEnvironmentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::MWAA::Environment
        /// </summary>
        public static Output<GetEnvironmentResult> Invoke(GetEnvironmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvironmentResult>("aws-native:mwaa:getEnvironment", args ?? new GetEnvironmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetEnvironmentArgs()
        {
        }
        public static new GetEnvironmentArgs Empty => new GetEnvironmentArgs();
    }

    public sealed class GetEnvironmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetEnvironmentInvokeArgs()
        {
        }
        public static new GetEnvironmentInvokeArgs Empty => new GetEnvironmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnvironmentResult
    {
        /// <summary>
        /// Key/value pairs representing Airflow configuration variables.
        ///     Keys are prefixed by their section:
        /// 
        ///     [core]
        ///     dags_folder={AIRFLOW_HOME}/dags
        /// 
        ///     Would be represented as
        /// 
        ///     "core.dags_folder": "{AIRFLOW_HOME}/dags"
        /// </summary>
        public readonly object? AirflowConfigurationOptions;
        public readonly string? AirflowVersion;
        public readonly string? Arn;
        public readonly string? DagS3Path;
        public readonly string? EnvironmentClass;
        public readonly string? ExecutionRoleArn;
        public readonly Outputs.EnvironmentLoggingConfiguration? LoggingConfiguration;
        public readonly int? MaxWorkers;
        public readonly int? MinWorkers;
        public readonly Outputs.EnvironmentNetworkConfiguration? NetworkConfiguration;
        public readonly string? PluginsS3ObjectVersion;
        public readonly string? PluginsS3Path;
        public readonly string? RequirementsS3ObjectVersion;
        public readonly string? RequirementsS3Path;
        public readonly int? Schedulers;
        public readonly string? SourceBucketArn;
        public readonly string? StartupScriptS3ObjectVersion;
        public readonly string? StartupScriptS3Path;
        /// <summary>
        /// A map of tags for the environment.
        /// </summary>
        public readonly object? Tags;
        public readonly Pulumi.AwsNative.MWAA.EnvironmentWebserverAccessMode? WebserverAccessMode;
        public readonly string? WebserverUrl;
        public readonly string? WeeklyMaintenanceWindowStart;

        [OutputConstructor]
        private GetEnvironmentResult(
            object? airflowConfigurationOptions,

            string? airflowVersion,

            string? arn,

            string? dagS3Path,

            string? environmentClass,

            string? executionRoleArn,

            Outputs.EnvironmentLoggingConfiguration? loggingConfiguration,

            int? maxWorkers,

            int? minWorkers,

            Outputs.EnvironmentNetworkConfiguration? networkConfiguration,

            string? pluginsS3ObjectVersion,

            string? pluginsS3Path,

            string? requirementsS3ObjectVersion,

            string? requirementsS3Path,

            int? schedulers,

            string? sourceBucketArn,

            string? startupScriptS3ObjectVersion,

            string? startupScriptS3Path,

            object? tags,

            Pulumi.AwsNative.MWAA.EnvironmentWebserverAccessMode? webserverAccessMode,

            string? webserverUrl,

            string? weeklyMaintenanceWindowStart)
        {
            AirflowConfigurationOptions = airflowConfigurationOptions;
            AirflowVersion = airflowVersion;
            Arn = arn;
            DagS3Path = dagS3Path;
            EnvironmentClass = environmentClass;
            ExecutionRoleArn = executionRoleArn;
            LoggingConfiguration = loggingConfiguration;
            MaxWorkers = maxWorkers;
            MinWorkers = minWorkers;
            NetworkConfiguration = networkConfiguration;
            PluginsS3ObjectVersion = pluginsS3ObjectVersion;
            PluginsS3Path = pluginsS3Path;
            RequirementsS3ObjectVersion = requirementsS3ObjectVersion;
            RequirementsS3Path = requirementsS3Path;
            Schedulers = schedulers;
            SourceBucketArn = sourceBucketArn;
            StartupScriptS3ObjectVersion = startupScriptS3ObjectVersion;
            StartupScriptS3Path = startupScriptS3Path;
            Tags = tags;
            WebserverAccessMode = webserverAccessMode;
            WebserverUrl = webserverUrl;
            WeeklyMaintenanceWindowStart = weeklyMaintenanceWindowStart;
        }
    }
}