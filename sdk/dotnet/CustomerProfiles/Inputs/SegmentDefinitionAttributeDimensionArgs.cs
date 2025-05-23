// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Inputs
{

    /// <summary>
    /// Specifies attribute based criteria for a segment.
    /// </summary>
    public sealed class SegmentDefinitionAttributeDimensionArgs : global::Pulumi.ResourceArgs
    {
        [Input("dimensionType", required: true)]
        public Input<Pulumi.AwsNative.CustomerProfiles.SegmentDefinitionAttributeDimensionType> DimensionType { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public SegmentDefinitionAttributeDimensionArgs()
        {
        }
        public static new SegmentDefinitionAttributeDimensionArgs Empty => new SegmentDefinitionAttributeDimensionArgs();
    }
}
