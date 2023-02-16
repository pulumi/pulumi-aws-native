// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    /// <summary>
    /// Resource schema for AWS::EC2::NetworkInsightsPath
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:NetworkInsightsPath")]
    public partial class NetworkInsightsPath : global::Pulumi.CustomResource
    {
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        [Output("destination")]
        public Output<string?> Destination { get; private set; } = null!;

        [Output("destinationArn")]
        public Output<string> DestinationArn { get; private set; } = null!;

        [Output("destinationIp")]
        public Output<string?> DestinationIp { get; private set; } = null!;

        [Output("destinationPort")]
        public Output<int?> DestinationPort { get; private set; } = null!;

        [Output("networkInsightsPathArn")]
        public Output<string> NetworkInsightsPathArn { get; private set; } = null!;

        [Output("networkInsightsPathId")]
        public Output<string> NetworkInsightsPathId { get; private set; } = null!;

        [Output("protocol")]
        public Output<Pulumi.AwsNative.EC2.NetworkInsightsPathProtocol> Protocol { get; private set; } = null!;

        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        [Output("sourceArn")]
        public Output<string> SourceArn { get; private set; } = null!;

        [Output("sourceIp")]
        public Output<string?> SourceIp { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.NetworkInsightsPathTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInsightsPath resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInsightsPath(string name, NetworkInsightsPathArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkInsightsPath", name, args ?? new NetworkInsightsPathArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInsightsPath(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkInsightsPath", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInsightsPath resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInsightsPath Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkInsightsPath(name, id, options);
        }
    }

    public sealed class NetworkInsightsPathArgs : global::Pulumi.ResourceArgs
    {
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        [Input("destinationIp")]
        public Input<string>? DestinationIp { get; set; }

        [Input("destinationPort")]
        public Input<int>? DestinationPort { get; set; }

        [Input("protocol", required: true)]
        public Input<Pulumi.AwsNative.EC2.NetworkInsightsPathProtocol> Protocol { get; set; } = null!;

        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        [Input("tags")]
        private InputList<Inputs.NetworkInsightsPathTagArgs>? _tags;
        public InputList<Inputs.NetworkInsightsPathTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.NetworkInsightsPathTagArgs>());
            set => _tags = value;
        }

        public NetworkInsightsPathArgs()
        {
        }
        public static new NetworkInsightsPathArgs Empty => new NetworkInsightsPathArgs();
    }
}