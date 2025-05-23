// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Outputs
{

    [OutputType]
    public sealed class RuleHttpParameters
    {
        /// <summary>
        /// The headers that need to be sent as part of request invoking the API Gateway API or EventBridge ApiDestination.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? HeaderParameters;
        /// <summary>
        /// The path parameter values to be used to populate API Gateway API or EventBridge ApiDestination path wildcards ("*").
        /// </summary>
        public readonly ImmutableArray<string> PathParameterValues;
        /// <summary>
        /// The query string keys/values that need to be sent as part of request invoking the API Gateway API or EventBridge ApiDestination.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? QueryStringParameters;

        [OutputConstructor]
        private RuleHttpParameters(
            ImmutableDictionary<string, string>? headerParameters,

            ImmutableArray<string> pathParameterValues,

            ImmutableDictionary<string, string>? queryStringParameters)
        {
            HeaderParameters = headerParameters;
            PathParameterValues = pathParameterValues;
            QueryStringParameters = queryStringParameters;
        }
    }
}
