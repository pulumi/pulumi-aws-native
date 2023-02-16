// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless
{
    /// <summary>
    /// Create and manage NetworkAnalyzerConfiguration resource.
    /// </summary>
    [AwsNativeResourceType("aws-native:iotwireless:NetworkAnalyzerConfiguration")]
    public partial class NetworkAnalyzerConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Arn for network analyzer configuration, Returned upon successful create.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the new resource
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the network analyzer configuration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.NetworkAnalyzerConfigurationTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Trace content for your wireless gateway and wireless device resources
        /// </summary>
        [Output("traceContent")]
        public Output<Outputs.TraceContentProperties?> TraceContent { get; private set; } = null!;

        /// <summary>
        /// List of wireless gateway resources that have been added to the network analyzer configuration
        /// </summary>
        [Output("wirelessDevices")]
        public Output<ImmutableArray<string>> WirelessDevices { get; private set; } = null!;

        /// <summary>
        /// List of wireless gateway resources that have been added to the network analyzer configuration
        /// </summary>
        [Output("wirelessGateways")]
        public Output<ImmutableArray<string>> WirelessGateways { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkAnalyzerConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkAnalyzerConfiguration(string name, NetworkAnalyzerConfigurationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:iotwireless:NetworkAnalyzerConfiguration", name, args ?? new NetworkAnalyzerConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkAnalyzerConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iotwireless:NetworkAnalyzerConfiguration", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkAnalyzerConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkAnalyzerConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkAnalyzerConfiguration(name, id, options);
        }
    }

    public sealed class NetworkAnalyzerConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the new resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the network analyzer configuration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.NetworkAnalyzerConfigurationTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.NetworkAnalyzerConfigurationTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.NetworkAnalyzerConfigurationTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Trace content for your wireless gateway and wireless device resources
        /// </summary>
        [Input("traceContent")]
        public Input<Inputs.TraceContentPropertiesArgs>? TraceContent { get; set; }

        [Input("wirelessDevices")]
        private InputList<string>? _wirelessDevices;

        /// <summary>
        /// List of wireless gateway resources that have been added to the network analyzer configuration
        /// </summary>
        public InputList<string> WirelessDevices
        {
            get => _wirelessDevices ?? (_wirelessDevices = new InputList<string>());
            set => _wirelessDevices = value;
        }

        [Input("wirelessGateways")]
        private InputList<string>? _wirelessGateways;

        /// <summary>
        /// List of wireless gateway resources that have been added to the network analyzer configuration
        /// </summary>
        public InputList<string> WirelessGateways
        {
            get => _wirelessGateways ?? (_wirelessGateways = new InputList<string>());
            set => _wirelessGateways = value;
        }

        public NetworkAnalyzerConfigurationArgs()
        {
        }
        public static new NetworkAnalyzerConfigurationArgs Empty => new NetworkAnalyzerConfigurationArgs();
    }
}