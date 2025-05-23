// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Contains the configurations to use Neptune Analytics as Vector Store.
    /// </summary>
    public sealed class KnowledgeBaseNeptuneAnalyticsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains the names of the fields to which to map information about the vector store.
        /// </summary>
        [Input("fieldMapping", required: true)]
        public Input<Inputs.KnowledgeBaseNeptuneAnalyticsFieldMappingArgs> FieldMapping { get; set; } = null!;

        /// <summary>
        /// ARN for Neptune Analytics graph database.
        /// </summary>
        [Input("graphArn", required: true)]
        public Input<string> GraphArn { get; set; } = null!;

        public KnowledgeBaseNeptuneAnalyticsConfigurationArgs()
        {
        }
        public static new KnowledgeBaseNeptuneAnalyticsConfigurationArgs Empty => new KnowledgeBaseNeptuneAnalyticsConfigurationArgs();
    }
}
