// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Properties related to the Amazon Elastic Block Store volume. Must be provided if storage type is Amazon EBS and must not be provided if storage type is not Amazon EBS
    /// </summary>
    public sealed class DomainDefaultEbsStorageSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default size of the Amazon EBS volume in Gb
        /// </summary>
        [Input("defaultEbsVolumeSizeInGb", required: true)]
        public Input<int> DefaultEbsVolumeSizeInGb { get; set; } = null!;

        /// <summary>
        /// Maximum size of the Amazon EBS volume in Gb. Must be greater than or equal to the DefaultEbsVolumeSizeInGb.
        /// </summary>
        [Input("maximumEbsVolumeSizeInGb", required: true)]
        public Input<int> MaximumEbsVolumeSizeInGb { get; set; } = null!;

        public DomainDefaultEbsStorageSettingsArgs()
        {
        }
        public static new DomainDefaultEbsStorageSettingsArgs Empty => new DomainDefaultEbsStorageSettingsArgs();
    }
}