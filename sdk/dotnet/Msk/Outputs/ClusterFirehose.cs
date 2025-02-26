// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Outputs
{

    [OutputType]
    public sealed class ClusterFirehose
    {
        public readonly string? DeliveryStream;
        public readonly bool Enabled;

        [OutputConstructor]
        private ClusterFirehose(
            string? deliveryStream,

            bool enabled)
        {
            DeliveryStream = deliveryStream;
            Enabled = enabled;
        }
    }
}
