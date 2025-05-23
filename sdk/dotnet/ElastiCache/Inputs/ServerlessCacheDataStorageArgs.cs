// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElastiCache.Inputs
{

    /// <summary>
    /// The cached data capacity of the Serverless Cache.
    /// </summary>
    public sealed class ServerlessCacheDataStorageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum cached data capacity of the Serverless Cache.
        /// </summary>
        [Input("maximum")]
        public Input<int>? Maximum { get; set; }

        /// <summary>
        /// The minimum cached data capacity of the Serverless Cache.
        /// </summary>
        [Input("minimum")]
        public Input<int>? Minimum { get; set; }

        /// <summary>
        /// The unit of cached data capacity of the Serverless Cache.
        /// </summary>
        [Input("unit", required: true)]
        public Input<Pulumi.AwsNative.ElastiCache.ServerlessCacheDataStorageUnit> Unit { get; set; } = null!;

        public ServerlessCacheDataStorageArgs()
        {
        }
        public static new ServerlessCacheDataStorageArgs Empty => new ServerlessCacheDataStorageArgs();
    }
}
