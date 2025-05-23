// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Inputs
{

    /// <summary>
    /// The instance metadata option settings for the infrastructure configuration.
    /// </summary>
    public sealed class InfrastructureConfigurationInstanceMetadataOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limit the number of hops that an instance metadata request can traverse to reach its destination.
        /// </summary>
        [Input("httpPutResponseHopLimit")]
        public Input<int>? HttpPutResponseHopLimit { get; set; }

        /// <summary>
        /// Indicates whether a signed token header is required for instance metadata retrieval requests. The values affect the response as follows: 
        /// </summary>
        [Input("httpTokens")]
        public Input<Pulumi.AwsNative.ImageBuilder.InfrastructureConfigurationInstanceMetadataOptionsHttpTokens>? HttpTokens { get; set; }

        public InfrastructureConfigurationInstanceMetadataOptionsArgs()
        {
        }
        public static new InfrastructureConfigurationInstanceMetadataOptionsArgs Empty => new InfrastructureConfigurationInstanceMetadataOptionsArgs();
    }
}
