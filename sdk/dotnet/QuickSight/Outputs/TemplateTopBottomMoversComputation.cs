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
    public sealed class TemplateTopBottomMoversComputation
    {
        /// <summary>
        /// The category field that is used in a computation.
        /// </summary>
        public readonly Outputs.TemplateDimensionField? Category;
        /// <summary>
        /// The ID for a computation.
        /// </summary>
        public readonly string ComputationId;
        /// <summary>
        /// The mover size setup of the top and bottom movers computation.
        /// </summary>
        public readonly double? MoverSize;
        /// <summary>
        /// The name of a computation.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The sort order setup of the top and bottom movers computation.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateTopBottomSortOrder? SortOrder;
        /// <summary>
        /// The time field that is used in a computation.
        /// </summary>
        public readonly Outputs.TemplateDimensionField? Time;
        /// <summary>
        /// The computation type. Choose from the following options:
        /// 
        /// - TOP: Top movers computation.
        /// - BOTTOM: Bottom movers computation.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateTopBottomComputationType Type;
        /// <summary>
        /// The value field that is used in a computation.
        /// </summary>
        public readonly Outputs.TemplateMeasureField? Value;

        [OutputConstructor]
        private TemplateTopBottomMoversComputation(
            Outputs.TemplateDimensionField? category,

            string computationId,

            double? moverSize,

            string? name,

            Pulumi.AwsNative.QuickSight.TemplateTopBottomSortOrder? sortOrder,

            Outputs.TemplateDimensionField? time,

            Pulumi.AwsNative.QuickSight.TemplateTopBottomComputationType type,

            Outputs.TemplateMeasureField? value)
        {
            Category = category;
            ComputationId = computationId;
            MoverSize = moverSize;
            Name = name;
            SortOrder = sortOrder;
            Time = time;
            Type = type;
            Value = value;
        }
    }
}
