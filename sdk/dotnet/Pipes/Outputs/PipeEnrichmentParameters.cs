// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeEnrichmentParameters
    {
        /// <summary>
        /// Contains the HTTP parameters to use when the target is a API Gateway REST endpoint or EventBridge ApiDestination.
        /// 
        /// If you specify an API Gateway REST API or EventBridge ApiDestination as a target, you can use this parameter to specify headers, path parameters, and query string keys/values as part of your target invoking request. If you're using ApiDestinations, the corresponding Connection can also have these values configured. In case of any conflicting keys, values from the Connection take precedence.
        /// </summary>
        public readonly Outputs.PipeEnrichmentHttpParameters? HttpParameters;
        /// <summary>
        /// Valid JSON text passed to the enrichment. In this case, nothing from the event itself is passed to the enrichment. For more information, see [The JavaScript Object Notation (JSON) Data Interchange Format](https://docs.aws.amazon.com/http://www.rfc-editor.org/rfc/rfc7159.txt) .
        /// 
        /// To remove an input template, specify an empty string.
        /// </summary>
        public readonly string? InputTemplate;

        [OutputConstructor]
        private PipeEnrichmentParameters(
            Outputs.PipeEnrichmentHttpParameters? httpParameters,

            string? inputTemplate)
        {
            HttpParameters = httpParameters;
            InputTemplate = inputTemplate;
        }
    }
}
