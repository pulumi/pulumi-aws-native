// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
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
        public readonly Outputs.AnalysisCalculatedMeasureField? CalculatedMeasureField;
        public readonly Outputs.AnalysisCategoricalMeasureField? CategoricalMeasureField;
        public readonly Outputs.AnalysisDateMeasureField? DateMeasureField;
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