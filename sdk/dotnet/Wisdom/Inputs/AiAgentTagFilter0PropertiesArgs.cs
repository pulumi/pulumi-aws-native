// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Inputs
{

    public sealed class AiAgentTagFilter0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("tagCondition", required: true)]
        public Input<Inputs.AiAgentTagConditionArgs> TagCondition { get; set; } = null!;

        public AiAgentTagFilter0PropertiesArgs()
        {
        }
        public static new AiAgentTagFilter0PropertiesArgs Empty => new AiAgentTagFilter0PropertiesArgs();
    }
}
