// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Outputs
{

    [OutputType]
    public sealed class VolumeSnaplockRetentionPeriod
    {
        public readonly Outputs.VolumeRetentionPeriod DefaultRetention;
        public readonly Outputs.VolumeRetentionPeriod MaximumRetention;
        public readonly Outputs.VolumeRetentionPeriod MinimumRetention;

        [OutputConstructor]
        private VolumeSnaplockRetentionPeriod(
            Outputs.VolumeRetentionPeriod defaultRetention,

            Outputs.VolumeRetentionPeriod maximumRetention,

            Outputs.VolumeRetentionPeriod minimumRetention)
        {
            DefaultRetention = defaultRetention;
            MaximumRetention = maximumRetention;
            MinimumRetention = minimumRetention;
        }
    }
}