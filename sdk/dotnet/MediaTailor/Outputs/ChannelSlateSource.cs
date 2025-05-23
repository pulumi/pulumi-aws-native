// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Outputs
{

    /// <summary>
    /// &lt;p&gt;Slate VOD source configuration.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class ChannelSlateSource
    {
        /// <summary>
        /// &lt;p&gt;The name of the source location where the slate VOD source is stored.&lt;/p&gt;
        /// </summary>
        public readonly string? SourceLocationName;
        /// <summary>
        /// &lt;p&gt;The slate VOD source name. The VOD source must already exist in a source location before it can be used for slate.&lt;/p&gt;
        /// </summary>
        public readonly string? VodSourceName;

        [OutputConstructor]
        private ChannelSlateSource(
            string? sourceLocationName,

            string? vodSourceName)
        {
            SourceLocationName = sourceLocationName;
            VodSourceName = vodSourceName;
        }
    }
}
