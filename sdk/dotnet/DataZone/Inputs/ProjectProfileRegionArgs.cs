// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    public sealed class ProjectProfileRegionArgs : global::Pulumi.ResourceArgs
    {
        [Input("regionName", required: true)]
        public Input<string> RegionName { get; set; } = null!;

        public ProjectProfileRegionArgs()
        {
        }
        public static new ProjectProfileRegionArgs Empty => new ProjectProfileRegionArgs();
    }
}
