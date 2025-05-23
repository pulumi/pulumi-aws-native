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
    public sealed class TemplateNegativeValueConfiguration
    {
        /// <summary>
        /// Determines the display mode of the negative value configuration.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateNegativeValueDisplayMode DisplayMode;

        [OutputConstructor]
        private TemplateNegativeValueConfiguration(Pulumi.AwsNative.QuickSight.TemplateNegativeValueDisplayMode displayMode)
        {
            DisplayMode = displayMode;
        }
    }
}
