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
    public sealed class AnalysisMeasureField
    {
        /// <summary>
        /// The calculated measure field only used in pivot tables.
        /// </summary>
        public readonly Outputs.AnalysisCalculatedMeasureField? CalculatedMeasureField;
        /// <summary>
        /// The measure type field with categorical type columns.
        /// </summary>
        public readonly Outputs.AnalysisCategoricalMeasureField? CategoricalMeasureField;
        /// <summary>
        /// The measure type field with date type columns.
        /// </summary>
        public readonly Outputs.AnalysisDateMeasureField? DateMeasureField;
        /// <summary>
        /// The measure type field with numerical type columns.
        /// </summary>
        public readonly Outputs.AnalysisNumericalMeasureField? NumericalMeasureField;

        [OutputConstructor]
        private AnalysisMeasureField(
            Outputs.AnalysisCalculatedMeasureField? calculatedMeasureField,

            Outputs.AnalysisCategoricalMeasureField? categoricalMeasureField,

            Outputs.AnalysisDateMeasureField? dateMeasureField,

            Outputs.AnalysisNumericalMeasureField? numericalMeasureField)
        {
            CalculatedMeasureField = calculatedMeasureField;
            CategoricalMeasureField = categoricalMeasureField;
            DateMeasureField = dateMeasureField;
            NumericalMeasureField = numericalMeasureField;
        }
    }
}
