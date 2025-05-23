// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// Model for configuration of a Topic
    /// </summary>
    [OutputType]
    public sealed class TopicConfigOptions
    {
        /// <summary>
        /// Enables Amazon Q Business Insights for a `Topic` .
        /// </summary>
        public readonly bool? QBusinessInsightsEnabled;

        [OutputConstructor]
        private TopicConfigOptions(bool? qBusinessInsightsEnabled)
        {
            QBusinessInsightsEnabled = qBusinessInsightsEnabled;
        }
    }
}
