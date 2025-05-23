// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Configuration for memory storage
    /// </summary>
    public sealed class AgentMemoryConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabledMemoryTypes")]
        private InputList<Pulumi.AwsNative.Bedrock.AgentMemoryType>? _enabledMemoryTypes;

        /// <summary>
        /// The type of memory that is stored.
        /// </summary>
        public InputList<Pulumi.AwsNative.Bedrock.AgentMemoryType> EnabledMemoryTypes
        {
            get => _enabledMemoryTypes ?? (_enabledMemoryTypes = new InputList<Pulumi.AwsNative.Bedrock.AgentMemoryType>());
            set => _enabledMemoryTypes = value;
        }

        /// <summary>
        /// Contains the configuration for SESSION_SUMMARY memory type enabled for the agent.
        /// </summary>
        [Input("sessionSummaryConfiguration")]
        public Input<Inputs.AgentSessionSummaryConfigurationArgs>? SessionSummaryConfiguration { get; set; }

        /// <summary>
        /// Maximum number of days to store session details
        /// </summary>
        [Input("storageDays")]
        public Input<double>? StorageDays { get; set; }

        public AgentMemoryConfigurationArgs()
        {
        }
        public static new AgentMemoryConfigurationArgs Empty => new AgentMemoryConfigurationArgs();
    }
}
