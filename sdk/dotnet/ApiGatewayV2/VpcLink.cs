// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2
{
    /// <summary>
    /// The ``AWS::ApiGatewayV2::VpcLink`` resource creates a VPC link. Supported only for HTTP APIs. The VPC link status must transition from ``PENDING`` to ``AVAILABLE`` to successfully create a VPC link, which can take up to 10 minutes. To learn more, see [Working with VPC Links for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-vpc-links.html) in the *API Gateway Developer Guide*.
    /// </summary>
    [AwsNativeResourceType("aws-native:apigatewayv2:VpcLink")]
    public partial class VpcLink : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the VPC link.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of security group IDs for the VPC link.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// A list of subnet IDs to include in the VPC link.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// The collection of tags. Each tag element is associated with a given resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The VPC link ID.
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
            : base("aws-native:apigatewayv2:VpcLink", name, args ?? new VpcLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcLink(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apigatewayv2:VpcLink", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "securityGroupIds[*]",
                    "subnetIds[*]",
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
        /// The name of the VPC link.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs for the VPC link.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of subnet IDs to include in the VPC link.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The collection of tags. Each tag element is associated with a given resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public VpcLinkArgs()
        {
        }
        public static new VpcLinkArgs Empty => new VpcLinkArgs();
    }
}
