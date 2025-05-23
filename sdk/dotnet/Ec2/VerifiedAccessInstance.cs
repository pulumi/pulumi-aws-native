// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    /// <summary>
    /// The AWS::EC2::VerifiedAccessInstance resource creates an AWS EC2 Verified Access Instance.
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:VerifiedAccessInstance")]
    public partial class VerifiedAccessInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Introduce CidrEndpointsCustomSubDomain property to represent the domain (say, ava.my-company.com)
        /// </summary>
        [Output("cidrEndpointsCustomSubDomain")]
        public Output<string?> CidrEndpointsCustomSubDomain { get; private set; } = null!;

        /// <summary>
        /// Property to represent the name servers assoicated with the domain that AVA manages (say, ['ns1.amazonaws.com', 'ns2.amazonaws.com', 'ns3.amazonaws.com', 'ns4.amazonaws.com']).
        /// </summary>
        [Output("cidrEndpointsCustomSubDomainNameServers")]
        public Output<ImmutableArray<string>> CidrEndpointsCustomSubDomainNameServers { get; private set; } = null!;

        /// <summary>
        /// Time this Verified Access Instance was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// A description for the AWS Verified Access instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates whether FIPS is enabled
        /// </summary>
        [Output("fipsEnabled")]
        public Output<bool?> FipsEnabled { get; private set; } = null!;

        /// <summary>
        /// Time this Verified Access Instance was last updated.
        /// </summary>
        [Output("lastUpdatedTime")]
        public Output<string> LastUpdatedTime { get; private set; } = null!;

        /// <summary>
        /// The configuration options for AWS Verified Access instances.
        /// </summary>
        [Output("loggingConfigurations")]
        public Output<Outputs.VerifiedAccessInstanceVerifiedAccessLogs?> LoggingConfigurations { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS Verified Access instance.
        /// </summary>
        [Output("verifiedAccessInstanceId")]
        public Output<string> VerifiedAccessInstanceId { get; private set; } = null!;

        /// <summary>
        /// The IDs of the AWS Verified Access trust providers.
        /// </summary>
        [Output("verifiedAccessTrustProviderIds")]
        public Output<ImmutableArray<string>> VerifiedAccessTrustProviderIds { get; private set; } = null!;

        /// <summary>
        /// AWS Verified Access trust providers.
        /// </summary>
        [Output("verifiedAccessTrustProviders")]
        public Output<ImmutableArray<Outputs.VerifiedAccessInstanceVerifiedAccessTrustProvider>> VerifiedAccessTrustProviders { get; private set; } = null!;


        /// <summary>
        /// Create a VerifiedAccessInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VerifiedAccessInstance(string name, VerifiedAccessInstanceArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ec2:VerifiedAccessInstance", name, args ?? new VerifiedAccessInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VerifiedAccessInstance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:VerifiedAccessInstance", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VerifiedAccessInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VerifiedAccessInstance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VerifiedAccessInstance(name, id, options);
        }
    }

    public sealed class VerifiedAccessInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Introduce CidrEndpointsCustomSubDomain property to represent the domain (say, ava.my-company.com)
        /// </summary>
        [Input("cidrEndpointsCustomSubDomain")]
        public Input<string>? CidrEndpointsCustomSubDomain { get; set; }

        /// <summary>
        /// A description for the AWS Verified Access instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether FIPS is enabled
        /// </summary>
        [Input("fipsEnabled")]
        public Input<bool>? FipsEnabled { get; set; }

        /// <summary>
        /// The configuration options for AWS Verified Access instances.
        /// </summary>
        [Input("loggingConfigurations")]
        public Input<Inputs.VerifiedAccessInstanceVerifiedAccessLogsArgs>? LoggingConfigurations { get; set; }

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

        [Input("verifiedAccessTrustProviderIds")]
        private InputList<string>? _verifiedAccessTrustProviderIds;

        /// <summary>
        /// The IDs of the AWS Verified Access trust providers.
        /// </summary>
        public InputList<string> VerifiedAccessTrustProviderIds
        {
            get => _verifiedAccessTrustProviderIds ?? (_verifiedAccessTrustProviderIds = new InputList<string>());
            set => _verifiedAccessTrustProviderIds = value;
        }

        [Input("verifiedAccessTrustProviders")]
        private InputList<Inputs.VerifiedAccessInstanceVerifiedAccessTrustProviderArgs>? _verifiedAccessTrustProviders;

        /// <summary>
        /// AWS Verified Access trust providers.
        /// </summary>
        public InputList<Inputs.VerifiedAccessInstanceVerifiedAccessTrustProviderArgs> VerifiedAccessTrustProviders
        {
            get => _verifiedAccessTrustProviders ?? (_verifiedAccessTrustProviders = new InputList<Inputs.VerifiedAccessInstanceVerifiedAccessTrustProviderArgs>());
            set => _verifiedAccessTrustProviders = value;
        }

        public VerifiedAccessInstanceArgs()
        {
        }
        public static new VerifiedAccessInstanceArgs Empty => new VerifiedAccessInstanceArgs();
    }
}
