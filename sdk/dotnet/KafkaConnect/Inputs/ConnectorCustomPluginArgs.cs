// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KafkaConnect.Inputs
{

    /// <summary>
    /// Details about a custom plugin.
    /// </summary>
    public sealed class ConnectorCustomPluginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the custom plugin to use.
        /// </summary>
        [Input("customPluginArn", required: true)]
        public Input<string> CustomPluginArn { get; set; } = null!;

        /// <summary>
        /// The revision of the custom plugin to use.
        /// </summary>
        [Input("revision", required: true)]
        public Input<int> Revision { get; set; } = null!;

        public ConnectorCustomPluginArgs()
        {
        }
        public static new ConnectorCustomPluginArgs Empty => new ConnectorCustomPluginArgs();
    }
}
