// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Inputs
{

    /// <summary>
    /// Contains information about an asset model hierarchy.
    /// </summary>
    public sealed class AssetModelHierarchyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the asset model. All assets in this hierarchy must be instances of the child AssetModelId asset model.
        /// </summary>
        [Input("childAssetModelId", required: true)]
        public Input<string> ChildAssetModelId { get; set; } = null!;

        /// <summary>
        /// Customer provided ID for hierarchy.
        /// </summary>
        [Input("logicalId", required: true)]
        public Input<string> LogicalId { get; set; } = null!;

        /// <summary>
        /// The name of the asset model hierarchy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public AssetModelHierarchyArgs()
        {
        }
        public static new AssetModelHierarchyArgs Empty => new AssetModelHierarchyArgs();
    }
}