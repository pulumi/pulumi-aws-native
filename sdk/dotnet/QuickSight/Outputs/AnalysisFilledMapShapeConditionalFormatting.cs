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
    public sealed class AnalysisFilledMapShapeConditionalFormatting
    {
        public readonly string FieldId;
        public readonly Outputs.AnalysisShapeConditionalFormat? Format;

        [OutputConstructor]
        private AnalysisFilledMapShapeConditionalFormatting(
            string fieldId,

            Outputs.AnalysisShapeConditionalFormat? format)
        {
            FieldId = fieldId;
            Format = format;
        }
    }
}