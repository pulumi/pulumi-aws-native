// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Type of Executors for an Action Group
    /// </summary>
    public sealed class AgentActionGroupExecutor0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of a Lambda.
        /// </summary>
        [Input("lambda", required: true)]
        public Input<string> Lambda { get; set; } = null!;

        public AgentActionGroupExecutor0PropertiesArgs()
        {
        }
        public static new AgentActionGroupExecutor0PropertiesArgs Empty => new AgentActionGroupExecutor0PropertiesArgs();
    }
}