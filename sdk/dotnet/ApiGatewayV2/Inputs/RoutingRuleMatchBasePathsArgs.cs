// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2.Inputs
{

    public sealed class RoutingRuleMatchBasePathsArgs : global::Pulumi.ResourceArgs
    {
        [Input("anyOf", required: true)]
        private InputList<string>? _anyOf;

        /// <summary>
        /// The string of the case sensitive base path to be matched.
        /// </summary>
        public InputList<string> AnyOf
        {
            get => _anyOf ?? (_anyOf = new InputList<string>());
            set => _anyOf = value;
        }

        public RoutingRuleMatchBasePathsArgs()
        {
        }
        public static new RoutingRuleMatchBasePathsArgs Empty => new RoutingRuleMatchBasePathsArgs();
    }
}
