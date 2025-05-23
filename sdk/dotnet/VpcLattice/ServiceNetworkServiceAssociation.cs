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
    /// Associates a service with a service network.
    /// </summary>
    [AwsNativeResourceType("aws-native:vpclattice:ServiceNetworkServiceAssociation")]
    public partial class ServiceNetworkServiceAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the association between the service network and the service.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ID of the of the association between the service network and the service.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The date and time that the association was created, specified in ISO-8601 format.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The DNS information of the service.
        /// </summary>
        [Output("dnsEntry")]
        public Output<Outputs.ServiceNetworkServiceAssociationDnsEntry?> DnsEntry { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the service.
        /// </summary>
        [Output("serviceArn")]
        public Output<string> ServiceArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the service.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// The ID or ARN of the service.
        /// </summary>
        [Output("serviceIdentifier")]
        public Output<string?> ServiceIdentifier { get; private set; } = null!;

        /// <summary>
        /// The name of the service.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the service network
        /// </summary>
        [Output("serviceNetworkArn")]
        public Output<string> ServiceNetworkArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the service network.
        /// </summary>
        [Output("serviceNetworkId")]
        public Output<string> ServiceNetworkId { get; private set; } = null!;

        /// <summary>
        /// The ID or ARN of the service network. You must use an ARN if the resources are in different accounts.
        /// </summary>
        [Output("serviceNetworkIdentifier")]
        public Output<string?> ServiceNetworkIdentifier { get; private set; } = null!;

        /// <summary>
        /// The name of the service network.
        /// </summary>
        [Output("serviceNetworkName")]
        public Output<string> ServiceNetworkName { get; private set; } = null!;

        /// <summary>
        /// The status of the association between the service network and the service.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.AwsNative.VpcLattice.ServiceNetworkServiceAssociationStatus> Status { get; private set; } = null!;

        /// <summary>
        /// The tags for the association.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceNetworkServiceAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceNetworkServiceAssociation(string name, ServiceNetworkServiceAssociationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:vpclattice:ServiceNetworkServiceAssociation", name, args ?? new ServiceNetworkServiceAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceNetworkServiceAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:vpclattice:ServiceNetworkServiceAssociation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "serviceIdentifier",
                    "serviceNetworkIdentifier",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceNetworkServiceAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceNetworkServiceAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceNetworkServiceAssociation(name, id, options);
        }
    }

    public sealed class ServiceNetworkServiceAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DNS information of the service.
        /// </summary>
        [Input("dnsEntry")]
        public Input<Inputs.ServiceNetworkServiceAssociationDnsEntryArgs>? DnsEntry { get; set; }

        /// <summary>
        /// The ID or ARN of the service.
        /// </summary>
        [Input("serviceIdentifier")]
        public Input<string>? ServiceIdentifier { get; set; }

        /// <summary>
        /// The ID or ARN of the service network. You must use an ARN if the resources are in different accounts.
        /// </summary>
        [Input("serviceNetworkIdentifier")]
        public Input<string>? ServiceNetworkIdentifier { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// The tags for the association.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public ServiceNetworkServiceAssociationArgs()
        {
        }
        public static new ServiceNetworkServiceAssociationArgs Empty => new ServiceNetworkServiceAssociationArgs();
    }
}
