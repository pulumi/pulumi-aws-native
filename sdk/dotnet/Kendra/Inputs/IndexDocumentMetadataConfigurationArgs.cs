// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class IndexDocumentMetadataConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("relevance")]
        public Input<Inputs.IndexRelevanceArgs>? Relevance { get; set; }

        [Input("search")]
        public Input<Inputs.IndexSearchArgs>? Search { get; set; }

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Kendra.IndexDocumentAttributeValueType> Type { get; set; } = null!;

        public IndexDocumentMetadataConfigurationArgs()
        {
        }
        public static new IndexDocumentMetadataConfigurationArgs Empty => new IndexDocumentMetadataConfigurationArgs();
    }
}
