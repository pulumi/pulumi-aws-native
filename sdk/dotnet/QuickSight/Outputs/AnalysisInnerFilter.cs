// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisInnerFilter
    {
        public readonly Outputs.AnalysisCategoryInnerFilter? CategoryInnerFilter;

        [OutputConstructor]
        private AnalysisInnerFilter(Outputs.AnalysisCategoryInnerFilter? categoryInnerFilter)
        {
            CategoryInnerFilter = categoryInnerFilter;
        }
    }
}