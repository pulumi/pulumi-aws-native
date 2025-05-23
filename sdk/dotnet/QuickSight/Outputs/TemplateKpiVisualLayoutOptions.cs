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
    public sealed class TemplateKpiVisualLayoutOptions
    {
        /// <summary>
        /// The standard layout of the KPI visual.
        /// </summary>
        public readonly Outputs.TemplateKpiVisualStandardLayout? StandardLayout;

        [OutputConstructor]
        private TemplateKpiVisualLayoutOptions(Outputs.TemplateKpiVisualStandardLayout? standardLayout)
        {
            StandardLayout = standardLayout;
        }
    }
}
