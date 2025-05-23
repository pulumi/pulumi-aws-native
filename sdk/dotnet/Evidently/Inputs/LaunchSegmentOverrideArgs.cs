// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently.Inputs
{

    public sealed class LaunchSegmentOverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A number indicating the order to use to evaluate segment overrides, if there are more than one. Segment overrides with lower numbers are evaluated first.
        /// </summary>
        [Input("evaluationOrder", required: true)]
        public Input<int> EvaluationOrder { get; set; } = null!;

        /// <summary>
        /// The ARN of the segment to use for this override.
        /// </summary>
        [Input("segment", required: true)]
        public Input<string> Segment { get; set; } = null!;

        [Input("weights", required: true)]
        private InputList<Inputs.LaunchGroupToWeightArgs>? _weights;

        /// <summary>
        /// The traffic allocation percentages among the feature variations to assign to this segment. This is a set of key-value pairs. The keys are variation names. The values represent the amount of traffic to allocate to that variation for this segment. This is expressed in thousandths of a percent, so a weight of 50000 represents 50% of traffic.
        /// </summary>
        public InputList<Inputs.LaunchGroupToWeightArgs> Weights
        {
            get => _weights ?? (_weights = new InputList<Inputs.LaunchGroupToWeightArgs>());
            set => _weights = value;
        }

        public LaunchSegmentOverrideArgs()
        {
        }
        public static new LaunchSegmentOverrideArgs Empty => new LaunchSegmentOverrideArgs();
    }
}
