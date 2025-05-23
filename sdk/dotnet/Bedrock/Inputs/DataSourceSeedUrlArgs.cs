// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// A seed url object.
    /// </summary>
    public sealed class DataSourceSeedUrlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A web url.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public DataSourceSeedUrlArgs()
        {
        }
        public static new DataSourceSeedUrlArgs Empty => new DataSourceSeedUrlArgs();
    }
}
