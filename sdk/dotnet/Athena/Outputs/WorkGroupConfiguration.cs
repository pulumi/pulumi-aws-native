// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Athena.Outputs
{

    [OutputType]
    public sealed class WorkGroupConfiguration
    {
        public readonly int? BytesScannedCutoffPerQuery;
        public readonly bool? EnforceWorkGroupConfiguration;
        public readonly Outputs.WorkGroupEngineVersion? EngineVersion;
        public readonly bool? PublishCloudWatchMetricsEnabled;
        public readonly bool? RequesterPaysEnabled;
        public readonly Outputs.WorkGroupResultConfiguration? ResultConfiguration;

        [OutputConstructor]
        private WorkGroupConfiguration(
            int? bytesScannedCutoffPerQuery,

            bool? enforceWorkGroupConfiguration,

            Outputs.WorkGroupEngineVersion? engineVersion,

            bool? publishCloudWatchMetricsEnabled,

            bool? requesterPaysEnabled,

            Outputs.WorkGroupResultConfiguration? resultConfiguration)
        {
            BytesScannedCutoffPerQuery = bytesScannedCutoffPerQuery;
            EnforceWorkGroupConfiguration = enforceWorkGroupConfiguration;
            EngineVersion = engineVersion;
            PublishCloudWatchMetricsEnabled = publishCloudWatchMetricsEnabled;
            RequesterPaysEnabled = requesterPaysEnabled;
            ResultConfiguration = resultConfiguration;
        }
    }
}