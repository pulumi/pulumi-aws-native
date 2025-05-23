// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeEnrichmentHttpParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("headerParameters")]
        private InputMap<string>? _headerParameters;

        /// <summary>
        /// The headers that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
        /// </summary>
        public InputMap<string> HeaderParameters
        {
            get => _headerParameters ?? (_headerParameters = new InputMap<string>());
            set => _headerParameters = value;
        }

        [Input("pathParameterValues")]
        private InputList<string>? _pathParameterValues;

        /// <summary>
        /// The path parameter values to be used to populate API Gateway REST API or EventBridge ApiDestination path wildcards ("*").
        /// </summary>
        public InputList<string> PathParameterValues
        {
            get => _pathParameterValues ?? (_pathParameterValues = new InputList<string>());
            set => _pathParameterValues = value;
        }

        [Input("queryStringParameters")]
        private InputMap<string>? _queryStringParameters;

        /// <summary>
        /// The query string keys/values that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
        /// </summary>
        public InputMap<string> QueryStringParameters
        {
            get => _queryStringParameters ?? (_queryStringParameters = new InputMap<string>());
            set => _queryStringParameters = value;
        }

        public PipeEnrichmentHttpParametersArgs()
        {
        }
        public static new PipeEnrichmentHttpParametersArgs Empty => new PipeEnrichmentHttpParametersArgs();
    }
}
