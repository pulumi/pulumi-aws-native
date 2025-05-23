// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Omics.Inputs
{

    /// <summary>
    /// Server-side encryption (SSE) settings for a store.
    /// </summary>
    public sealed class SequenceStoreSseConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An encryption key ARN.
        /// </summary>
        [Input("keyArn")]
        public Input<string>? KeyArn { get; set; }

        /// <summary>
        /// The encryption type.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Omics.SequenceStoreEncryptionType> Type { get; set; } = null!;

        public SequenceStoreSseConfigArgs()
        {
        }
        public static new SequenceStoreSseConfigArgs Empty => new SequenceStoreSseConfigArgs();
    }
}
