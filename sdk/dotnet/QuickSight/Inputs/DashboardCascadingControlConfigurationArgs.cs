// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardCascadingControlConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceControls")]
        private InputList<Inputs.DashboardCascadingControlSourceArgs>? _sourceControls;

        /// <summary>
        /// A list of source controls that determine the values that are used in the current control.
        /// </summary>
        public InputList<Inputs.DashboardCascadingControlSourceArgs> SourceControls
        {
            get => _sourceControls ?? (_sourceControls = new InputList<Inputs.DashboardCascadingControlSourceArgs>());
            set => _sourceControls = value;
        }

        public DashboardCascadingControlConfigurationArgs()
        {
        }
        public static new DashboardCascadingControlConfigurationArgs Empty => new DashboardCascadingControlConfigurationArgs();
    }
}
