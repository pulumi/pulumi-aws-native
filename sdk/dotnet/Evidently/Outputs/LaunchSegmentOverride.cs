// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently.Outputs
{

    [OutputType]
    public sealed class LaunchSegmentOverride
    {
        public readonly int EvaluationOrder;
        public readonly string Segment;
        public readonly ImmutableArray<Outputs.LaunchGroupToWeight> Weights;

        [OutputConstructor]
        private LaunchSegmentOverride(
            int evaluationOrder,

            string segment,

            ImmutableArray<Outputs.LaunchGroupToWeight> weights)
        {
            EvaluationOrder = evaluationOrder;
            Segment = segment;
            Weights = weights;
        }
    }
}