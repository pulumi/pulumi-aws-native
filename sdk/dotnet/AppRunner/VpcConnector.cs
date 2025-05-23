// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppRunner
{
    /// <summary>
    /// The AWS::AppRunner::VpcConnector resource specifies an App Runner VpcConnector.
    /// </summary>
    [AwsNativeResourceType("aws-native:apprunner:VpcConnector")]
    public partial class VpcConnector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<string>> Subnets { get; private set; } = null!;

        /// <summary>
        /// A list of metadata items that you can associate with your VPC connector resource. A tag is a key-value pair.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.CreateOnlyTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of this VPC connector.
        /// </summary>
        [Output("vpcConnectorArn")]
        public Output<string> VpcConnectorArn { get; private set; } = null!;

        /// <summary>
        /// A name for the VPC connector. If you don't specify a name, AWS CloudFormation generates a name for your VPC connector.
        /// </summary>
        [Output("vpcConnectorName")]
        public Output<string?> VpcConnectorName { get; private set; } = null!;

        /// <summary>
        /// The revision of this VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
        /// </summary>
        [Output("vpcConnectorRevision")]
        public Output<int> VpcConnectorRevision { get; private set; } = null!;


        /// <summary>
        /// Create a VpcConnector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcConnector(string name, VpcConnectorArgs args, CustomResourceOptions? options = null)
            : base("aws-native:apprunner:VpcConnector", name, args ?? new VpcConnectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcConnector(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apprunner:VpcConnector", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "securityGroups[*]",
                    "subnets[*]",
                    "tags[*]",
                    "vpcConnectorName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcConnector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcConnector Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpcConnector(name, id, options);
        }
    }

    public sealed class VpcConnectorArgs : global::Pulumi.ResourceArgs
    {
        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnets", required: true)]
        private InputList<string>? _subnets;

        /// <summary>
        /// A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs>? _tags;

        /// <summary>
        /// A list of metadata items that you can associate with your VPC connector resource. A tag is a key-value pair.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// A name for the VPC connector. If you don't specify a name, AWS CloudFormation generates a name for your VPC connector.
        /// </summary>
        [Input("vpcConnectorName")]
        public Input<string>? VpcConnectorName { get; set; }

        public VpcConnectorArgs()
        {
        }
        public static new VpcConnectorArgs Empty => new VpcConnectorArgs();
    }
}
