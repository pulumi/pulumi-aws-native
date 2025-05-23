// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Inputs
{

    /// <summary>
    /// Specify checkpoint version for update. This is only required to update the Compatibility.
    /// </summary>
    public sealed class SchemaVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the latest version needs to be updated.
        /// </summary>
        [Input("isLatest")]
        public Input<bool>? IsLatest { get; set; }

        /// <summary>
        /// Indicates the version number in the schema to update.
        /// </summary>
        [Input("versionNumber")]
        public Input<int>? VersionNumber { get; set; }

        public SchemaVersionArgs()
        {
        }
        public static new SchemaVersionArgs Empty => new SchemaVersionArgs();
    }
}
