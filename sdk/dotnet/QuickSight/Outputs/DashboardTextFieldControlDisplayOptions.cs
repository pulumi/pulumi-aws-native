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
    public sealed class DashboardTextFieldControlDisplayOptions
    {
        public readonly Outputs.DashboardTextControlPlaceholderOptions? PlaceholderOptions;
        public readonly Outputs.DashboardLabelOptions? TitleOptions;

        [OutputConstructor]
        private DashboardTextFieldControlDisplayOptions(
            Outputs.DashboardTextControlPlaceholderOptions? placeholderOptions,

            Outputs.DashboardLabelOptions? titleOptions)
        {
            PlaceholderOptions = placeholderOptions;
            TitleOptions = titleOptions;
        }
    }
}