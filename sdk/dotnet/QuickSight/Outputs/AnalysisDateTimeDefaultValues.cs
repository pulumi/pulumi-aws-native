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
    public sealed class AnalysisDateTimeDefaultValues
    {
        /// <summary>
        /// The dynamic value of the `DataTimeDefaultValues` . Different defaults are displayed according to users, groups, and values mapping.
        /// </summary>
        public readonly Outputs.AnalysisDynamicDefaultValue? DynamicValue;
        /// <summary>
        /// The rolling date of the `DataTimeDefaultValues` . The date is determined from the dataset based on input expression.
        /// </summary>
        public readonly Outputs.AnalysisRollingDateConfiguration? RollingDate;
        /// <summary>
        /// The static values of the `DataTimeDefaultValues` .
        /// </summary>
        public readonly ImmutableArray<string> StaticValues;

        [OutputConstructor]
        private AnalysisDateTimeDefaultValues(
            Outputs.AnalysisDynamicDefaultValue? dynamicValue,

            Outputs.AnalysisRollingDateConfiguration? rollingDate,

            ImmutableArray<string> staticValues)
        {
            DynamicValue = dynamicValue;
            RollingDate = rollingDate;
            StaticValues = staticValues;
        }
    }
}
