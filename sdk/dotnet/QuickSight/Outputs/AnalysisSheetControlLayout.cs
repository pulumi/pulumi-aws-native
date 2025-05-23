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
    public sealed class AnalysisSheetControlLayout
    {
        /// <summary>
        /// The configuration that determines the elements and canvas size options of sheet control.
        /// </summary>
        public readonly Outputs.AnalysisSheetControlLayoutConfiguration Configuration;

        [OutputConstructor]
        private AnalysisSheetControlLayout(Outputs.AnalysisSheetControlLayoutConfiguration configuration)
        {
            Configuration = configuration;
        }
    }
}
