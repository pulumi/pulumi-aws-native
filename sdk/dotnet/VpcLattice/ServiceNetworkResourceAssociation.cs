// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice
{
    /// <summary>
    /// VpcLattice ServiceNetworkResourceAssociation CFN resource
    /// </summary>
    [AwsNativeResourceType("aws-native:vpclattice:ServiceNetworkResourceAssociation")]
    public partial class ServiceNetworkResourceAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the association.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ID of the association between the service network and resource configuration.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource configuration associated with the service network.
        /// </summary>
        [Output("resourceConfigurationId")]
        public Output<string?> ResourceConfigurationId { get; private set; } = null!;

        /// <summary>
        /// The ID of the service network associated with the resource configuration.
        /// </summary>
        [Output("serviceNetworkId")]
        public Output<string?> ServiceNetworkId { get; private set; } = null!;

        /// <summary>
        /// A key-value pair to associate with a resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceNetworkResourceAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceNetworkResourceAssociation(string name, ServiceNetworkResourceAssociationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:vpclattice:ServiceNetworkResourceAssociation", name, args ?? new ServiceNetworkResourceAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceNetworkResourceAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:vpclattice:ServiceNetworkResourceAssociation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "resourceConfigurationId",
                    "serviceNetworkId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceNetworkResourceAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceNetworkResourceAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceNetworkResourceAssociation(name, id, options);
        }
    }

    public sealed class ServiceNetworkResourceAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the resource configuration associated with the service network.
        /// </summary>
        [Input("resourceConfigurationId")]
        public Input<string>? ResourceConfigurationId { get; set; }

        /// <summary>
        /// The ID of the service network associated with the resource configuration.
        /// </summary>
        [Input("serviceNetworkId")]
        public Input<string>? ServiceNetworkId { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// A key-value pair to associate with a resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public ServiceNetworkResourceAssociationArgs()
        {
        }
        public static new ServiceNetworkResourceAssociationArgs Empty => new ServiceNetworkResourceAssociationArgs();
    }
}
