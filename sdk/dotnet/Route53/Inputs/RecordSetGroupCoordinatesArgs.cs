// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53.Inputs
{

    public sealed class RecordSetGroupCoordinatesArgs : global::Pulumi.ResourceArgs
    {
        [Input("latitude", required: true)]
        public Input<string> Latitude { get; set; } = null!;

        [Input("longitude", required: true)]
        public Input<string> Longitude { get; set; } = null!;

        public RecordSetGroupCoordinatesArgs()
        {
        }
        public static new RecordSetGroupCoordinatesArgs Empty => new RecordSetGroupCoordinatesArgs();
    }
}