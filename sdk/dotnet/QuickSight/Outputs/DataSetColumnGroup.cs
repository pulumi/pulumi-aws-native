// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;Groupings of columns that work together in certain Amazon QuickSight features. This is
    ///             a variant type structure. For this structure to be valid, only one of the attributes can
    ///             be non-null.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetColumnGroup
    {
        public readonly Outputs.DataSetGeoSpatialColumnGroup? GeoSpatialColumnGroup;

        [OutputConstructor]
        private DataSetColumnGroup(Outputs.DataSetGeoSpatialColumnGroup? geoSpatialColumnGroup)
        {
            GeoSpatialColumnGroup = geoSpatialColumnGroup;
        }
    }
}