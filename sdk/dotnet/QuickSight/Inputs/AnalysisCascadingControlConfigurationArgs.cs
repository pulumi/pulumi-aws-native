// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisCascadingControlConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceControls")]
        private InputList<Inputs.AnalysisCascadingControlSourceArgs>? _sourceControls;
        public InputList<Inputs.AnalysisCascadingControlSourceArgs> SourceControls
        {
            get => _sourceControls ?? (_sourceControls = new InputList<Inputs.AnalysisCascadingControlSourceArgs>());
            set => _sourceControls = value;
        }

        public AnalysisCascadingControlConfigurationArgs()
        {
        }
        public static new AnalysisCascadingControlConfigurationArgs Empty => new AnalysisCascadingControlConfigurationArgs();
    }
}