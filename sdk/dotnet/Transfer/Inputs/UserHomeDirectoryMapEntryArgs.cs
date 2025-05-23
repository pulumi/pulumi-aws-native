// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Transfer.Inputs
{

    public sealed class UserHomeDirectoryMapEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents an entry for `HomeDirectoryMappings` .
        /// </summary>
        [Input("entry", required: true)]
        public Input<string> Entry { get; set; } = null!;

        /// <summary>
        /// Represents the map target that is used in a `HomeDirectoryMapEntry` .
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// Specifies the type of mapping. Set the type to `FILE` if you want the mapping to point to a file, or `DIRECTORY` for the directory to point to a directory.
        /// 
        /// &gt; By default, home directory mappings have a `Type` of `DIRECTORY` when you create a Transfer Family server. You would need to explicitly set `Type` to `FILE` if you want a mapping to have a file target.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AwsNative.Transfer.UserMapType>? Type { get; set; }

        public UserHomeDirectoryMapEntryArgs()
        {
        }
        public static new UserHomeDirectoryMapEntryArgs Empty => new UserHomeDirectoryMapEntryArgs();
    }
}
