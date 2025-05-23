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
    /// Data connection configuration
    /// </summary>
    public sealed class FlowDataConnectionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of a node output in a flow
        /// </summary>
        [Input("sourceOutput", required: true)]
        public Input<string> SourceOutput { get; set; } = null!;

        /// <summary>
        /// Name of a node input in a flow
        /// </summary>
        [Input("targetInput", required: true)]
        public Input<string> TargetInput { get; set; } = null!;

        public FlowDataConnectionConfigurationArgs()
        {
        }
        public static new FlowDataConnectionConfigurationArgs Empty => new FlowDataConnectionConfigurationArgs();
    }
}
