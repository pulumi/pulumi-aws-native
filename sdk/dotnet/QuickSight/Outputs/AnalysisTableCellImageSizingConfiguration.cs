// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisTableCellImageSizingConfiguration
    {
        /// <summary>
        /// The cell scaling configuration of the sizing options for the table image configuration.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisTableCellImageScalingConfiguration? TableCellImageScalingConfiguration;

        [OutputConstructor]
        private AnalysisTableCellImageSizingConfiguration(Pulumi.AwsNative.QuickSight.AnalysisTableCellImageScalingConfiguration? tableCellImageScalingConfiguration)
        {
            TableCellImageScalingConfiguration = tableCellImageScalingConfiguration;
        }
    }
}
