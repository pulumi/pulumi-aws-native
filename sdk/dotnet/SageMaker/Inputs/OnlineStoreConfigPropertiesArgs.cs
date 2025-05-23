// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// The configuration of an `OnlineStore` .
    /// </summary>
    public sealed class OnlineStoreConfigPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Turn `OnlineStore` off by specifying `False` for the `EnableOnlineStore` flag. Turn `OnlineStore` on by specifying `True` for the `EnableOnlineStore` flag.
        /// 
        /// The default value is `False` .
        /// </summary>
        [Input("enableOnlineStore")]
        public Input<bool>? EnableOnlineStore { get; set; }

        /// <summary>
        /// Use to specify KMS Key ID ( `KMSKeyId` ) for at-rest encryption of your `OnlineStore` .
        /// </summary>
        [Input("securityConfig")]
        public Input<Inputs.FeatureGroupOnlineStoreSecurityConfigArgs>? SecurityConfig { get; set; }

        /// <summary>
        /// Option for different tiers of low latency storage for real-time data retrieval.
        /// 
        /// - `Standard` : A managed low latency data store for feature groups.
        /// - `InMemory` : A managed data store for feature groups that supports very low latency retrieval.
        /// </summary>
        [Input("storageType")]
        public Input<Pulumi.AwsNative.SageMaker.FeatureGroupStorageType>? StorageType { get; set; }

        /// <summary>
        /// Time to live duration, where the record is hard deleted after the expiration time is reached; `ExpiresAt` = `EventTime` + `TtlDuration` . For information on HardDelete, see the [DeleteRecord](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_feature_store_DeleteRecord.html) API in the Amazon SageMaker API Reference guide.
        /// </summary>
        [Input("ttlDuration")]
        public Input<Inputs.FeatureGroupTtlDurationArgs>? TtlDuration { get; set; }

        public OnlineStoreConfigPropertiesArgs()
        {
        }
        public static new OnlineStoreConfigPropertiesArgs Empty => new OnlineStoreConfigPropertiesArgs();
    }
}
