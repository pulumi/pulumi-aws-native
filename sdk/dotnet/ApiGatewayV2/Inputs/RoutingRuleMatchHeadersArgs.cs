// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2.Inputs
{

    public sealed class RoutingRuleMatchHeadersArgs : global::Pulumi.ResourceArgs
    {
        [Input("anyOf", required: true)]
        private InputList<Inputs.RoutingRuleMatchHeaderValueArgs>? _anyOf;
        public InputList<Inputs.RoutingRuleMatchHeaderValueArgs> AnyOf
        {
            get => _anyOf ?? (_anyOf = new InputList<Inputs.RoutingRuleMatchHeaderValueArgs>());
            set => _anyOf = value;
        }

        public RoutingRuleMatchHeadersArgs()
        {
        }
        public static new RoutingRuleMatchHeadersArgs Empty => new RoutingRuleMatchHeadersArgs();
    }
}
