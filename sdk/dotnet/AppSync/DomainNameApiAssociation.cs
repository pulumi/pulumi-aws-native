// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    /// <summary>
    /// Resource Type definition for AWS::AppSync::DomainNameApiAssociation
    /// </summary>
    [AwsNativeResourceType("aws-native:appsync:DomainNameApiAssociation")]
    public partial class DomainNameApiAssociation : global::Pulumi.CustomResource
    {
        [Output("apiAssociationIdentifier")]
        public Output<string> ApiAssociationIdentifier { get; private set; } = null!;

        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;


        /// <summary>
        /// Create a DomainNameApiAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainNameApiAssociation(string name, DomainNameApiAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appsync:DomainNameApiAssociation", name, args ?? new DomainNameApiAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainNameApiAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appsync:DomainNameApiAssociation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DomainNameApiAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainNameApiAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DomainNameApiAssociation(name, id, options);
        }
    }

    public sealed class DomainNameApiAssociationArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        public DomainNameApiAssociationArgs()
        {
        }
        public static new DomainNameApiAssociationArgs Empty => new DomainNameApiAssociationArgs();
    }
}