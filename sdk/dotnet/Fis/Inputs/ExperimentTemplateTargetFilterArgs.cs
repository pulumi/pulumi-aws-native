// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fis.Inputs
{

    /// <summary>
    /// Describes a filter used for the target resource input in an experiment template.
    /// </summary>
    public sealed class ExperimentTemplateTargetFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ExperimentTemplateTargetFilterArgs()
        {
        }
        public static new ExperimentTemplateTargetFilterArgs Empty => new ExperimentTemplateTargetFilterArgs();
    }
}
