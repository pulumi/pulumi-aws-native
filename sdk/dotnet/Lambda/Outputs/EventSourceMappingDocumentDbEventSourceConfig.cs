// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    /// <summary>
    /// Document db event source config.
    /// </summary>
    [OutputType]
    public sealed class EventSourceMappingDocumentDbEventSourceConfig
    {
        /// <summary>
        /// The collection name to connect to.
        /// </summary>
        public readonly string? CollectionName;
        /// <summary>
        /// The database name to connect to.
        /// </summary>
        public readonly string? DatabaseName;
        /// <summary>
        /// Include full document in change stream response. The default option will only send the changes made to documents to Lambda. If you want the complete document sent to Lambda, set this to UpdateLookup.
        /// </summary>
        public readonly Pulumi.AwsNative.Lambda.EventSourceMappingDocumentDbEventSourceConfigFullDocument? FullDocument;

        [OutputConstructor]
        private EventSourceMappingDocumentDbEventSourceConfig(
            string? collectionName,

            string? databaseName,

            Pulumi.AwsNative.Lambda.EventSourceMappingDocumentDbEventSourceConfigFullDocument? fullDocument)
        {
            CollectionName = collectionName;
            DatabaseName = databaseName;
            FullDocument = fullDocument;
        }
    }
}