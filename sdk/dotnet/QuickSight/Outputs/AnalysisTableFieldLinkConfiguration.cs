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
    public sealed class AnalysisTableFieldLinkConfiguration
    {
        /// <summary>
        /// The URL content (text, icon) for the table link configuration.
        /// </summary>
        public readonly Outputs.AnalysisTableFieldLinkContentConfiguration Content;
        /// <summary>
        /// The URL target (new tab, new window, same tab) for the table link configuration.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisUrlTargetConfiguration Target;

        [OutputConstructor]
        private AnalysisTableFieldLinkConfiguration(
            Outputs.AnalysisTableFieldLinkContentConfiguration content,

            Pulumi.AwsNative.QuickSight.AnalysisUrlTargetConfiguration target)
        {
            Content = content;
            Target = target;
        }
    }
}
