// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    /// <summary>
    /// Resource Type definition for AWS::Connect::ApprovedOrigin
    /// </summary>
    [AwsNativeResourceType("aws-native:connect:ApprovedOrigin")]
    public partial class ApprovedOrigin : global::Pulumi.CustomResource
    {
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        [Output("origin")]
        public Output<string> Origin { get; private set; } = null!;


        /// <summary>
        /// Create a ApprovedOrigin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApprovedOrigin(string name, ApprovedOriginArgs args, CustomResourceOptions? options = null)
            : base("aws-native:connect:ApprovedOrigin", name, args ?? new ApprovedOriginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApprovedOrigin(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:connect:ApprovedOrigin", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ApprovedOrigin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApprovedOrigin Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApprovedOrigin(name, id, options);
        }
    }

    public sealed class ApprovedOriginArgs : global::Pulumi.ResourceArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("origin", required: true)]
        public Input<string> Origin { get; set; } = null!;

        public ApprovedOriginArgs()
        {
        }
        public static new ApprovedOriginArgs Empty => new ApprovedOriginArgs();
    }
}