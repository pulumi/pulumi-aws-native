// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    /// <summary>
    /// Document db event source config.
    /// </summary>
    public sealed class EventSourceMappingDocumentDBEventSourceConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The collection name to connect to.
        /// </summary>
        [Input("collectionName")]
        public Input<string>? CollectionName { get; set; }

        /// <summary>
        /// The database name to connect to.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// Include full document in change stream response. The default option will only send the changes made to documents to Lambda. If you want the complete document sent to Lambda, set this to UpdateLookup.
        /// </summary>
        [Input("fullDocument")]
        public Input<Pulumi.AwsNative.Lambda.EventSourceMappingDocumentDBEventSourceConfigFullDocument>? FullDocument { get; set; }

        public EventSourceMappingDocumentDBEventSourceConfigArgs()
        {
        }
        public static new EventSourceMappingDocumentDBEventSourceConfigArgs Empty => new EventSourceMappingDocumentDBEventSourceConfigArgs();
    }
}