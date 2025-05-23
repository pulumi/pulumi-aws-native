// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift
{
    /// <summary>
    /// The AWS::GameLift::Location resource creates an Amazon GameLift (GameLift) custom location.
    /// </summary>
    [AwsNativeResourceType("aws-native:gamelift:Location")]
    public partial class Location : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A unique identifier for the custom location. For example, `arn:aws:gamelift:[region]::location/location-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912` .
        /// </summary>
        [Output("locationArn")]
        public Output<string> LocationArn { get; private set; } = null!;

        /// <summary>
        /// A descriptive name for the custom location.
        /// </summary>
        [Output("locationName")]
        public Output<string> LocationName { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Location resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Location(string name, LocationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:gamelift:Location", name, args ?? new LocationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Location(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:gamelift:Location", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "locationName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Location resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Location Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Location(name, id, options);
        }
    }

    public sealed class LocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A descriptive name for the custom location.
        /// </summary>
        [Input("locationName")]
        public Input<string>? LocationName { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public LocationArgs()
        {
        }
        public static new LocationArgs Empty => new LocationArgs();
    }
}
