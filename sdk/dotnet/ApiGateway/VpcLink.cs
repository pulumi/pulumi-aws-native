// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    /// <summary>
    /// The ``AWS::ApiGateway::VpcLink`` resource creates an API Gateway VPC link for a REST API to access resources in an Amazon Virtual Private Cloud (VPC). For more information, see [vpclink:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateVpcLink.html) in the ``Amazon API Gateway REST API Reference``.
    /// </summary>
    [AwsNativeResourceType("aws-native:apigateway:VpcLink")]
    public partial class VpcLink : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the VPC link.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name used to label and identify the VPC link.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An array of arbitrary tags (key-value pairs) to associate with the VPC link.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ARN of the network load balancer of the VPC targeted by the VPC link. The network load balancer must be owned by the same AWS account of the API owner.
        /// </summary>
        [Output("targetArns")]
        public Output<ImmutableArray<string>> TargetArns { get; private set; } = null!;

        /// <summary>
        /// The ID for the VPC link. For example: `abc123` .
        /// </summary>
        [Output("vpcLinkId")]
        public Output<string> VpcLinkId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcLink(string name, VpcLinkArgs args, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:VpcLink", name, args ?? new VpcLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcLink(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:VpcLink", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "targetArns[*]",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcLink Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpcLink(name, id, options);
        }
    }

    public sealed class VpcLinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the VPC link.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name used to label and identify the VPC link.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of arbitrary tags (key-value pairs) to associate with the VPC link.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("targetArns", required: true)]
        private InputList<string>? _targetArns;

        /// <summary>
        /// The ARN of the network load balancer of the VPC targeted by the VPC link. The network load balancer must be owned by the same AWS account of the API owner.
        /// </summary>
        public InputList<string> TargetArns
        {
            get => _targetArns ?? (_targetArns = new InputList<string>());
            set => _targetArns = value;
        }

        public VpcLinkArgs()
        {
        }
        public static new VpcLinkArgs Empty => new VpcLinkArgs();
    }
}
