// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceDiscovery
{
    /// <summary>
    /// Resource Type definition for AWS::ServiceDiscovery::PrivateDnsNamespace
    /// </summary>
    [Obsolete(@"PrivateDnsNamespace is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:servicediscovery:PrivateDnsNamespace")]
    public partial class PrivateDnsNamespace : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("hostedZoneId")]
        public Output<string> HostedZoneId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("properties")]
        public Output<Outputs.PrivateDnsNamespaceProperties?> Properties { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.PrivateDnsNamespaceTag>> Tags { get; private set; } = null!;

        [Output("vpc")]
        public Output<string> Vpc { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateDnsNamespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateDnsNamespace(string name, PrivateDnsNamespaceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:servicediscovery:PrivateDnsNamespace", name, args ?? new PrivateDnsNamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateDnsNamespace(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:servicediscovery:PrivateDnsNamespace", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateDnsNamespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateDnsNamespace Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateDnsNamespace(name, id, options);
        }
    }

    public sealed class PrivateDnsNamespaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("properties")]
        public Input<Inputs.PrivateDnsNamespacePropertiesArgs>? Properties { get; set; }

        [Input("tags")]
        private InputList<Inputs.PrivateDnsNamespaceTagArgs>? _tags;
        public InputList<Inputs.PrivateDnsNamespaceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.PrivateDnsNamespaceTagArgs>());
            set => _tags = value;
        }

        [Input("vpc", required: true)]
        public Input<string> Vpc { get; set; } = null!;

        public PrivateDnsNamespaceArgs()
        {
        }
        public static new PrivateDnsNamespaceArgs Empty => new PrivateDnsNamespaceArgs();
    }
}