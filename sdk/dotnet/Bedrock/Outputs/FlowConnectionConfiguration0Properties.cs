// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// Connection configuration
    /// </summary>
    [OutputType]
    public sealed class FlowConnectionConfiguration0Properties
    {
        public readonly Outputs.FlowDataConnectionConfiguration Data;

        [OutputConstructor]
        private FlowConnectionConfiguration0Properties(Outputs.FlowDataConnectionConfiguration data)
        {
            Data = data;
        }
    }
}