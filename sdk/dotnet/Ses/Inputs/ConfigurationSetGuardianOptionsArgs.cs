// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    /// <summary>
    /// Preferences regarding the Guardian feature.
    /// </summary>
    public sealed class ConfigurationSetGuardianOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether emails sent with this configuration set have optimized delivery algorithm enabled.
        /// </summary>
        [Input("optimizedSharedDelivery", required: true)]
        public Input<string> OptimizedSharedDelivery { get; set; } = null!;

        public ConfigurationSetGuardianOptionsArgs()
        {
        }
        public static new ConfigurationSetGuardianOptionsArgs Empty => new ConfigurationSetGuardianOptionsArgs();
    }
}