// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline.Outputs
{

    [OutputType]
    public sealed class FleetVCpuCountRange
    {
        public readonly int? Max;
        public readonly int Min;

        [OutputConstructor]
        private FleetVCpuCountRange(
            int? max,

            int min)
        {
            Max = max;
            Min = min;
        }
    }
}