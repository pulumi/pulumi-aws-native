// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently.Inputs
{

    public sealed class FeatureVariationObjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("booleanValue")]
        public Input<bool>? BooleanValue { get; set; }

        [Input("doubleValue")]
        public Input<double>? DoubleValue { get; set; }

        [Input("longValue")]
        public Input<double>? LongValue { get; set; }

        [Input("stringValue")]
        public Input<string>? StringValue { get; set; }

        [Input("variationName")]
        public Input<string>? VariationName { get; set; }

        public FeatureVariationObjectArgs()
        {
        }
        public static new FeatureVariationObjectArgs Empty => new FeatureVariationObjectArgs();
    }
}