// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution.Inputs
{

    public sealed class IdNamespaceRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("matchingKeys", required: true)]
        private InputList<string>? _matchingKeys;
        public InputList<string> MatchingKeys
        {
            get => _matchingKeys ?? (_matchingKeys = new InputList<string>());
            set => _matchingKeys = value;
        }

        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        public IdNamespaceRuleArgs()
        {
        }
        public static new IdNamespaceRuleArgs Empty => new IdNamespaceRuleArgs();
    }
}