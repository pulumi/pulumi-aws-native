// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects. When you set this value to &lt;code&gt;LENIENT&lt;/code&gt;, validation is skipped for specific errors.&lt;/p&gt;
    /// </summary>
    public sealed class AnalysisValidationStrategyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The mode of validation for the asset to be created or updated. When you set this value to `STRICT` , strict validation for every error is enforced. When you set this value to `LENIENT` , validation is skipped for specific UI errors.
        /// </summary>
        [Input("mode", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisValidationStrategyMode> Mode { get; set; } = null!;

        public AnalysisValidationStrategyArgs()
        {
        }
        public static new AnalysisValidationStrategyArgs Empty => new AnalysisValidationStrategyArgs();
    }
}
