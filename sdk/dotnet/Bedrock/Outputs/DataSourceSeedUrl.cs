// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// A seed url object.
    /// </summary>
    [OutputType]
    public sealed class DataSourceSeedUrl
    {
        /// <summary>
        /// A web url.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private DataSourceSeedUrl(string url)
        {
            Url = url;
        }
    }
}