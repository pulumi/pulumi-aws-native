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
    /// Node configuration in a flow
    /// </summary>
    public sealed class FlowNodeConfiguration10PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("collector", required: true)]
        public Input<Inputs.FlowCollectorFlowNodeConfigurationArgs> Collector { get; set; } = null!;

        public FlowNodeConfiguration10PropertiesArgs()
        {
        }
        public static new FlowNodeConfiguration10PropertiesArgs Empty => new FlowNodeConfiguration10PropertiesArgs();
    }
}
