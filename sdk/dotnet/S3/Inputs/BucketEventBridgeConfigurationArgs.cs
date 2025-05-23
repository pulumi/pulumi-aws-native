// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Amazon S3 can send events to Amazon EventBridge whenever certain events happen in your bucket, see [Using EventBridge](https://docs.aws.amazon.com/AmazonS3/latest/userguide/EventBridge.html) in the *Amazon S3 User Guide*.
    ///  Unlike other destinations, delivery of events to EventBridge can be either enabled or disabled for a bucket. If enabled, all events will be sent to EventBridge and you can use EventBridge rules to route events to additional targets. For more information, see [What Is Amazon EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-what-is.html) in the *Amazon EventBridge User Guide*
    /// </summary>
    public sealed class BucketEventBridgeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables delivery of events to Amazon EventBridge.
        /// </summary>
        [Input("eventBridgeEnabled", required: true)]
        public Input<bool> EventBridgeEnabled { get; set; } = null!;

        public BucketEventBridgeConfigurationArgs()
        {
        }
        public static new BucketEventBridgeConfigurationArgs Empty => new BucketEventBridgeConfigurationArgs();
    }
}
