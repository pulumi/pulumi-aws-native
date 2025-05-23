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
    /// An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.
    /// </summary>
    [OutputType]
    public sealed class ConfigurationSetReputationOptions
    {
        /// <summary>
        /// If true , tracking of reputation metrics is enabled for the configuration set. If false , tracking of reputation metrics is disabled for the configuration set.
        /// </summary>
        public readonly bool? ReputationMetricsEnabled;

        [OutputConstructor]
        private ConfigurationSetReputationOptions(bool? reputationMetricsEnabled)
        {
            ReputationMetricsEnabled = reputationMetricsEnabled;
        }
    }
}
