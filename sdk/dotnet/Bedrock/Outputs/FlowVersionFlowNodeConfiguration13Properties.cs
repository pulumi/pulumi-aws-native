// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// Node configuration in a flow
    /// </summary>
    [OutputType]
    public sealed class FlowVersionFlowNodeConfiguration13Properties
    {
        public readonly Outputs.FlowVersionLoopFlowNodeConfiguration Loop;

        [OutputConstructor]
        private FlowVersionFlowNodeConfiguration13Properties(Outputs.FlowVersionLoopFlowNodeConfiguration loop)
        {
            Loop = loop;
        }
    }
}
