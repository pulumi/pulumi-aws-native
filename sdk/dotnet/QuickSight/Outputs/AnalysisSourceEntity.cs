// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;The source entity of an analysis.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class AnalysisSourceEntity
    {
        /// <summary>
        /// The source template for the source entity of the analysis.
        /// </summary>
        public readonly Outputs.AnalysisSourceTemplate? SourceTemplate;

        [OutputConstructor]
        private AnalysisSourceEntity(Outputs.AnalysisSourceTemplate? sourceTemplate)
        {
            SourceTemplate = sourceTemplate;
        }
    }
}
