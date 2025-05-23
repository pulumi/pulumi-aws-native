// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Inputs
{

    /// <summary>
    /// Contains a property type, which can be one of attribute, measurement, metric, or transform.
    /// </summary>
    public sealed class AssetModelPropertyTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies an asset attribute property. An attribute generally contains static information, such as the serial number of an [IIoT](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Internet_of_things#Industrial_applications) wind turbine.
        /// </summary>
        [Input("attribute")]
        public Input<Inputs.AssetModelAttributeArgs>? Attribute { get; set; }

        /// <summary>
        /// Specifies an asset metric property. A metric contains a mathematical expression that uses aggregate functions to process all input data points over a time interval and output a single data point, such as to calculate the average hourly temperature.
        /// </summary>
        [Input("metric")]
        public Input<Inputs.AssetModelMetricArgs>? Metric { get; set; }

        /// <summary>
        /// Specifies an asset transform property. A transform contains a mathematical expression that maps a property's data points from one form to another, such as a unit conversion from Celsius to Fahrenheit.
        /// </summary>
        [Input("transform")]
        public Input<Inputs.AssetModelTransformArgs>? Transform { get; set; }

        /// <summary>
        /// The type of property type, which can be one of `Attribute` , `Measurement` , `Metric` , or `Transform` .
        /// </summary>
        [Input("typeName", required: true)]
        public Input<Pulumi.AwsNative.IoTSiteWise.AssetModelTypeName> TypeName { get; set; } = null!;

        public AssetModelPropertyTypeArgs()
        {
        }
        public static new AssetModelPropertyTypeArgs Empty => new AssetModelPropertyTypeArgs();
    }
}
