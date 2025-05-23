// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2.Inputs
{

    public sealed class RoutingRuleActionInvokeApiArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        [Input("stage", required: true)]
        public Input<string> Stage { get; set; } = null!;

        [Input("stripBasePath")]
        public Input<bool>? StripBasePath { get; set; }

        public RoutingRuleActionInvokeApiArgs()
        {
        }
        public static new RoutingRuleActionInvokeApiArgs Empty => new RoutingRuleActionInvokeApiArgs();
    }
}
