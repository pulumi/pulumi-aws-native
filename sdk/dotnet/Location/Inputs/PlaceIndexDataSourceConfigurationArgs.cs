// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Location.Inputs
{

    public sealed class PlaceIndexDataSourceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("intendedUse")]
        public Input<Pulumi.AwsNative.Location.PlaceIndexIntendedUse>? IntendedUse { get; set; }

        public PlaceIndexDataSourceConfigurationArgs()
        {
        }
        public static new PlaceIndexDataSourceConfigurationArgs Empty => new PlaceIndexDataSourceConfigurationArgs();
    }
}