// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateAttributeAggregationFunctionArgs : global::Pulumi.ResourceArgs
    {
        [Input("simpleAttributeAggregation")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateSimpleAttributeAggregationFunction>? SimpleAttributeAggregation { get; set; }

        [Input("valueForMultipleValues")]
        public Input<string>? ValueForMultipleValues { get; set; }

        public TemplateAttributeAggregationFunctionArgs()
        {
        }
        public static new TemplateAttributeAggregationFunctionArgs Empty => new TemplateAttributeAggregationFunctionArgs();
    }
}