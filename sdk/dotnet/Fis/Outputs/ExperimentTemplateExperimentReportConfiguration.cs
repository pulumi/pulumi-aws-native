// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fis.Outputs
{

    [OutputType]
    public sealed class ExperimentTemplateExperimentReportConfiguration
    {
        /// <summary>
        /// The data sources for the experiment report.
        /// </summary>
        public readonly Outputs.ExperimentTemplateExperimentReportConfigurationDataSourcesProperties? DataSources;
        /// <summary>
        /// The output destinations of the experiment report.
        /// </summary>
        public readonly Outputs.ExperimentTemplateExperimentReportConfigurationOutputsProperties Outputs;
        /// <summary>
        /// The duration after the experiment end time for the data sources to include in the report.
        /// </summary>
        public readonly string? PostExperimentDuration;
        /// <summary>
        /// The duration before the experiment start time for the data sources to include in the report.
        /// </summary>
        public readonly string? PreExperimentDuration;

        [OutputConstructor]
        private ExperimentTemplateExperimentReportConfiguration(
            Outputs.ExperimentTemplateExperimentReportConfigurationDataSourcesProperties? dataSources,

            Outputs.ExperimentTemplateExperimentReportConfigurationOutputsProperties outputs,

            string? postExperimentDuration,

            string? preExperimentDuration)
        {
            DataSources = dataSources;
            Outputs = outputs;
            PostExperimentDuration = postExperimentDuration;
            PreExperimentDuration = preExperimentDuration;
        }
    }
}
