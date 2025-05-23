// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    /// <summary>
    /// A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudfront:Distribution")]
    public partial class Distribution : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The distribution's identifier. For example: `E1U5RQF7T870K0` .
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The distribution's configuration.
        /// </summary>
        [Output("distributionConfig")]
        public Output<Outputs.DistributionConfig> DistributionConfig { get; private set; } = null!;

        /// <summary>
        /// The domain name of the resource, such as `d111111abcdef8.cloudfront.net` .
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// A complex type that contains zero or more ``Tag`` elements.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Distribution resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Distribution(string name, DistributionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudfront:Distribution", name, args ?? new DistributionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Distribution(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudfront:Distribution", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Distribution resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Distribution Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Distribution(name, id, options);
        }
    }

    public sealed class DistributionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The distribution's configuration.
        /// </summary>
        [Input("distributionConfig", required: true)]
        public Input<Inputs.DistributionConfigArgs> DistributionConfig { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// A complex type that contains zero or more ``Tag`` elements.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public DistributionArgs()
        {
        }
        public static new DistributionArgs Empty => new DistributionArgs();
    }
}
