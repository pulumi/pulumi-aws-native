// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionImagePullSecretArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Provides a unique identifier for the `ImagePullSecret` . This object is required when `EksPodProperties$imagePullSecrets` is used.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public JobDefinitionImagePullSecretArgs()
        {
        }
        public static new JobDefinitionImagePullSecretArgs Empty => new JobDefinitionImagePullSecretArgs();
    }
}