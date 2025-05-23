// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    /// <summary>
    /// Resource Type definition for AWS::AppSync::DomainName
    /// </summary>
    [AwsNativeResourceType("aws-native:appsync:DomainName")]
    public partial class DomainName : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The domain name provided by AWS AppSync .
        /// </summary>
        [Output("appSyncDomainName")]
        public Output<string> AppSyncDomainName { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the certificate. This will be an AWS Certificate Manager certificate.
        /// </summary>
        [Output("certificateArn")]
        public Output<string> CertificateArn { get; private set; } = null!;

        /// <summary>
        /// The decription for your domain name.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The domain name.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainNameValue { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the Domain Name.
        /// </summary>
        [Output("domainNameArn")]
        public Output<string> DomainNameArn { get; private set; } = null!;

        /// <summary>
        /// The ID of your Amazon Route 53 hosted zone.
        /// </summary>
        [Output("hostedZoneId")]
        public Output<string> HostedZoneId { get; private set; } = null!;

        /// <summary>
        /// A set of tags (key-value pairs) for this domain name.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DomainName resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainName(string name, DomainNameArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appsync:DomainName", name, args ?? new DomainNameArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainName(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appsync:DomainName", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "certificateArn",
                    "domainName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainName resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainName Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DomainName(name, id, options);
        }
    }

    public sealed class DomainNameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the certificate. This will be an AWS Certificate Manager certificate.
        /// </summary>
        [Input("certificateArn", required: true)]
        public Input<string> CertificateArn { get; set; } = null!;

        /// <summary>
        /// The decription for your domain name.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain name.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainNameValue { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// A set of tags (key-value pairs) for this domain name.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public DomainNameArgs()
        {
        }
        public static new DomainNameArgs Empty => new DomainNameArgs();
    }
}
