// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Outputs
{

    /// <summary>
    /// An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.
    /// </summary>
    [OutputType]
    public sealed class ConfigurationSetEventDestinationCloudWatchDestination
    {
        /// <summary>
        /// A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigurationSetEventDestinationDimensionConfiguration> DimensionConfigurations;

        [OutputConstructor]
        private ConfigurationSetEventDestinationCloudWatchDestination(ImmutableArray<Outputs.ConfigurationSetEventDestinationDimensionConfiguration> dimensionConfigurations)
        {
            DimensionConfigurations = dimensionConfigurations;
        }
    }
}
