// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    /// <summary>
    /// Resource Type definition for AWS::CloudFront::OriginAccessControl
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudfront:OriginAccessControl")]
    public partial class OriginAccessControl : global::Pulumi.CustomResource
    {
        [Output("originAccessControlConfig")]
        public Output<Outputs.OriginAccessControlConfig> OriginAccessControlConfig { get; private set; } = null!;


        /// <summary>
        /// Create a OriginAccessControl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OriginAccessControl(string name, OriginAccessControlArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudfront:OriginAccessControl", name, args ?? new OriginAccessControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OriginAccessControl(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudfront:OriginAccessControl", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OriginAccessControl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OriginAccessControl Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OriginAccessControl(name, id, options);
        }
    }

    public sealed class OriginAccessControlArgs : global::Pulumi.ResourceArgs
    {
        [Input("originAccessControlConfig", required: true)]
        public Input<Inputs.OriginAccessControlConfigArgs> OriginAccessControlConfig { get; set; } = null!;

        public OriginAccessControlArgs()
        {
        }
        public static new OriginAccessControlArgs Empty => new OriginAccessControlArgs();
    }
}