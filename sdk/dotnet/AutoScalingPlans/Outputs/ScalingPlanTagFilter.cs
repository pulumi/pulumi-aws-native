// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScalingPlans.Outputs
{

    [OutputType]
    public sealed class ScalingPlanTagFilter
    {
        public readonly string Key;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private ScalingPlanTagFilter(
            string key,

            ImmutableArray<string> values)
        {
            Key = key;
            Values = values;
        }
    }
}