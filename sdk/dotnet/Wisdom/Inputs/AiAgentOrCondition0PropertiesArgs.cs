// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Inputs
{

    public sealed class AiAgentOrCondition0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("andConditions", required: true)]
        private InputList<Inputs.AiAgentTagConditionArgs>? _andConditions;
        public InputList<Inputs.AiAgentTagConditionArgs> AndConditions
        {
            get => _andConditions ?? (_andConditions = new InputList<Inputs.AiAgentTagConditionArgs>());
            set => _andConditions = value;
        }

        public AiAgentOrCondition0PropertiesArgs()
        {
        }
        public static new AiAgentOrCondition0PropertiesArgs Empty => new AiAgentOrCondition0PropertiesArgs();
    }
}
