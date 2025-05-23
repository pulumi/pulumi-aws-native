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
    /// Specifies an Amazon DocumentDB or MongoDB data store to crawl.
    /// </summary>
    public sealed class CrawlerMongoDbTargetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// The path of the Amazon DocumentDB or MongoDB target (database/collection).
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public CrawlerMongoDbTargetArgs()
        {
        }
        public static new CrawlerMongoDbTargetArgs Empty => new CrawlerMongoDbTargetArgs();
    }
}
