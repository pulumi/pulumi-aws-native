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
    public sealed class IndexDocumentAttributeConfiguration
    {
        /// <summary>
        /// The name of the document attribute.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Information about whether the document attribute can be used by an end user to search for information on their web experience.
        /// </summary>
        public readonly Pulumi.AwsNative.QBusiness.QBusinessIndexStatus? Search;
        /// <summary>
        /// The type of document attribute.
        /// </summary>
        public readonly Pulumi.AwsNative.QBusiness.IndexAttributeType? Type;

        [OutputConstructor]
        private IndexDocumentAttributeConfiguration(
            string? name,

            Pulumi.AwsNative.QBusiness.QBusinessIndexStatus? search,

            Pulumi.AwsNative.QBusiness.IndexAttributeType? type)
        {
            Name = name;
            Search = search;
            Type = type;
        }
    }
}