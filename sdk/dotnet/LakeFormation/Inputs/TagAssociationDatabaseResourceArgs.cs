// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation.Inputs
{

    public sealed class TagAssociationDatabaseResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier for the Data Catalog . By default, it should be the account ID of the caller.
        /// </summary>
        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        /// <summary>
        /// The name of the database resource. Unique to the Data Catalog.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public TagAssociationDatabaseResourceArgs()
        {
        }
        public static new TagAssociationDatabaseResourceArgs Empty => new TagAssociationDatabaseResourceArgs();
    }
}
