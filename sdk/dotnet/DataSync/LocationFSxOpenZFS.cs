// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    /// <summary>
    /// Resource schema for AWS::DataSync::LocationFSxOpenZFS.
    /// </summary>
    [AwsNativeResourceType("aws-native:datasync:LocationFSxOpenZFS")]
    public partial class LocationFSxOpenZFS : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the FSx OpenZFS file system.
        /// </summary>
        [Output("fsxFilesystemArn")]
        public Output<string?> FsxFilesystemArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
        /// </summary>
        [Output("locationArn")]
        public Output<string> LocationArn { get; private set; } = null!;

        /// <summary>
        /// The URL of the FSx OpenZFS that was described.
        /// </summary>
        [Output("locationUri")]
        public Output<string> LocationUri { get; private set; } = null!;

        [Output("protocol")]
        public Output<Outputs.LocationFSxOpenZFSProtocol> Protocol { get; private set; } = null!;

        /// <summary>
        /// The ARNs of the security groups that are to use to configure the FSx OpenZFS file system.
        /// </summary>
        [Output("securityGroupArns")]
        public Output<ImmutableArray<string>> SecurityGroupArns { get; private set; } = null!;

        /// <summary>
        /// A subdirectory in the location's path.
        /// </summary>
        [Output("subdirectory")]
        public Output<string?> Subdirectory { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.LocationFSxOpenZFSTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a LocationFSxOpenZFS resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocationFSxOpenZFS(string name, LocationFSxOpenZFSArgs args, CustomResourceOptions? options = null)
            : base("aws-native:datasync:LocationFSxOpenZFS", name, args ?? new LocationFSxOpenZFSArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocationFSxOpenZFS(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:datasync:LocationFSxOpenZFS", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LocationFSxOpenZFS resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocationFSxOpenZFS Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LocationFSxOpenZFS(name, id, options);
        }
    }

    public sealed class LocationFSxOpenZFSArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the FSx OpenZFS file system.
        /// </summary>
        [Input("fsxFilesystemArn")]
        public Input<string>? FsxFilesystemArn { get; set; }

        [Input("protocol", required: true)]
        public Input<Inputs.LocationFSxOpenZFSProtocolArgs> Protocol { get; set; } = null!;

        [Input("securityGroupArns", required: true)]
        private InputList<string>? _securityGroupArns;

        /// <summary>
        /// The ARNs of the security groups that are to use to configure the FSx OpenZFS file system.
        /// </summary>
        public InputList<string> SecurityGroupArns
        {
            get => _securityGroupArns ?? (_securityGroupArns = new InputList<string>());
            set => _securityGroupArns = value;
        }

        /// <summary>
        /// A subdirectory in the location's path.
        /// </summary>
        [Input("subdirectory")]
        public Input<string>? Subdirectory { get; set; }

        [Input("tags")]
        private InputList<Inputs.LocationFSxOpenZFSTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.LocationFSxOpenZFSTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LocationFSxOpenZFSTagArgs>());
            set => _tags = value;
        }

        public LocationFSxOpenZFSArgs()
        {
        }
        public static new LocationFSxOpenZFSArgs Empty => new LocationFSxOpenZFSArgs();
    }
}