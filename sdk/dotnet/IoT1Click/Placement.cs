// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT1Click
{
    /// <summary>
    /// Resource Type definition for AWS::IoT1Click::Placement
    /// </summary>
    [Obsolete(@"Placement is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:iot1click:Placement")]
    public partial class Placement : global::Pulumi.CustomResource
    {
        [Output("associatedDevices")]
        public Output<object?> AssociatedDevices { get; private set; } = null!;

        [Output("attributes")]
        public Output<object?> Attributes { get; private set; } = null!;

        [Output("placementName")]
        public Output<string?> PlacementName { get; private set; } = null!;

        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;


        /// <summary>
        /// Create a Placement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Placement(string name, PlacementArgs args, CustomResourceOptions? options = null)
            : base("aws-native:iot1click:Placement", name, args ?? new PlacementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Placement(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iot1click:Placement", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Placement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Placement Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Placement(name, id, options);
        }
    }

    public sealed class PlacementArgs : global::Pulumi.ResourceArgs
    {
        [Input("associatedDevices")]
        public Input<object>? AssociatedDevices { get; set; }

        [Input("attributes")]
        public Input<object>? Attributes { get; set; }

        [Input("placementName")]
        public Input<string>? PlacementName { get; set; }

        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        public PlacementArgs()
        {
        }
        public static new PlacementArgs Empty => new PlacementArgs();
    }
}