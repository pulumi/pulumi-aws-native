// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    /// <summary>
    /// An object that defines the open and click tracking options for emails that you send using the configuration set.
    /// </summary>
    public sealed class ConfigurationSetTrackingOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to use for tracking open and click events.
        /// </summary>
        [Input("customRedirectDomain")]
        public Input<string>? CustomRedirectDomain { get; set; }

        /// <summary>
        /// The https policy to use for tracking open and click events.
        /// </summary>
        [Input("httpsPolicy")]
        public Input<string>? HttpsPolicy { get; set; }

        public ConfigurationSetTrackingOptionsArgs()
        {
        }
        public static new ConfigurationSetTrackingOptionsArgs Empty => new ConfigurationSetTrackingOptionsArgs();
    }
}
