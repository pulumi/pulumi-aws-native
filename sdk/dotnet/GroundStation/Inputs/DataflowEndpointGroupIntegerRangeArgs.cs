// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Inputs
{

    /// <summary>
    /// An integer range that has a minimum and maximum value.
    /// </summary>
    public sealed class DataflowEndpointGroupIntegerRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A maximum value.
        /// </summary>
        [Input("maximum")]
        public Input<int>? Maximum { get; set; }

        /// <summary>
        /// A minimum value.
        /// </summary>
        [Input("minimum")]
        public Input<int>? Minimum { get; set; }

        public DataflowEndpointGroupIntegerRangeArgs()
        {
        }
        public static new DataflowEndpointGroupIntegerRangeArgs Empty => new DataflowEndpointGroupIntegerRangeArgs();
    }
}
