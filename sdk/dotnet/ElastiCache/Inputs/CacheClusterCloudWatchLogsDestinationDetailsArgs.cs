// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElastiCache.Inputs
{

    public sealed class CacheClusterCloudWatchLogsDestinationDetailsArgs : global::Pulumi.ResourceArgs
    {
        [Input("logGroup", required: true)]
        public Input<string> LogGroup { get; set; } = null!;

        public CacheClusterCloudWatchLogsDestinationDetailsArgs()
        {
        }
        public static new CacheClusterCloudWatchLogsDestinationDetailsArgs Empty => new CacheClusterCloudWatchLogsDestinationDetailsArgs();
    }
}