// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Inputs
{

    /// <summary>
    /// &lt;p&gt;Slate VOD source configuration.&lt;/p&gt;
    /// </summary>
    public sealed class ChannelSlateSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The name of the source location where the slate VOD source is stored.&lt;/p&gt;
        /// </summary>
        [Input("sourceLocationName")]
        public Input<string>? SourceLocationName { get; set; }

        /// <summary>
        /// &lt;p&gt;The slate VOD source name. The VOD source must already exist in a source location before it can be used for slate.&lt;/p&gt;
        /// </summary>
        [Input("vodSourceName")]
        public Input<string>? VodSourceName { get; set; }

        public ChannelSlateSourceArgs()
        {
        }
        public static new ChannelSlateSourceArgs Empty => new ChannelSlateSourceArgs();
    }
}