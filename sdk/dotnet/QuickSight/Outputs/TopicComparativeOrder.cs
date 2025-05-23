// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TopicComparativeOrder
    {
        /// <summary>
        /// The list of columns to be used in the ordering.
        /// </summary>
        public readonly ImmutableArray<string> SpecifedOrder;
        /// <summary>
        /// The treat of undefined specified values. Valid values for this structure are `LEAST` and `MOST` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TopicUndefinedSpecifiedValueType? TreatUndefinedSpecifiedValues;
        /// <summary>
        /// The ordering type for a column. Valid values for this structure are `GREATER_IS_BETTER` , `LESSER_IS_BETTER` and `SPECIFIED` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TopicColumnOrderingType? UseOrdering;

        [OutputConstructor]
        private TopicComparativeOrder(
            ImmutableArray<string> specifedOrder,

            Pulumi.AwsNative.QuickSight.TopicUndefinedSpecifiedValueType? treatUndefinedSpecifiedValues,

            Pulumi.AwsNative.QuickSight.TopicColumnOrderingType? useOrdering)
        {
            SpecifedOrder = specifedOrder;
            TreatUndefinedSpecifiedValues = treatUndefinedSpecifiedValues;
            UseOrdering = useOrdering;
        }
    }
}
