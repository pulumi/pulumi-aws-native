// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Outputs
{

    [OutputType]
    public sealed class IndexTextDocumentStatistics
    {
        /// <summary>
        /// The total size, in bytes, of the indexed documents.
        /// </summary>
        public readonly double? IndexedTextBytes;
        /// <summary>
        /// The number of text documents indexed.
        /// </summary>
        public readonly double? IndexedTextDocumentCount;

        [OutputConstructor]
        private IndexTextDocumentStatistics(
            double? indexedTextBytes,

            double? indexedTextDocumentCount)
        {
            IndexedTextBytes = indexedTextBytes;
            IndexedTextDocumentCount = indexedTextDocumentCount;
        }
    }
}