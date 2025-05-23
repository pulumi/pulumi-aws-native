// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dms
{
    /// <summary>
    /// Resource schema for AWS::DMS::InstanceProfile.
    /// </summary>
    [AwsNativeResourceType("aws-native:dms:InstanceProfile")]
    public partial class InstanceProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The property describes an availability zone of the instance profile.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string?> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The optional description of the instance profile.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The property describes an ARN of the instance profile.
        /// </summary>
        [Output("instanceProfileArn")]
        public Output<string> InstanceProfileArn { get; private set; } = null!;

        /// <summary>
        /// The property describes a creating time of the instance profile.
        /// </summary>
        [Output("instanceProfileCreationTime")]
        public Output<string> InstanceProfileCreationTime { get; private set; } = null!;

        /// <summary>
        /// The property describes an identifier for the instance profile. It is used for describing/deleting/modifying. Can be name/arn
        /// </summary>
        [Output("instanceProfileIdentifier")]
        public Output<string?> InstanceProfileIdentifier { get; private set; } = null!;

        /// <summary>
        /// The property describes a name for the instance profile.
        /// </summary>
        [Output("instanceProfileName")]
        public Output<string?> InstanceProfileName { get; private set; } = null!;

        /// <summary>
        /// The property describes kms key arn for the instance profile.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// The property describes a network type for the instance profile.
        /// </summary>
        [Output("networkType")]
        public Output<Pulumi.AwsNative.Dms.InstanceProfileNetworkType?> NetworkType { get; private set; } = null!;

        /// <summary>
        /// The property describes the publicly accessible of the instance profile
        /// </summary>
        [Output("publiclyAccessible")]
        public Output<bool?> PubliclyAccessible { get; private set; } = null!;

        /// <summary>
        /// The property describes a subnet group identifier for the instance profile.
        /// </summary>
        [Output("subnetGroupIdentifier")]
        public Output<string?> SubnetGroupIdentifier { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The property describes vps security groups for the instance profile.
        /// </summary>
        [Output("vpcSecurityGroups")]
        public Output<ImmutableArray<string>> VpcSecurityGroups { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceProfile(string name, InstanceProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:dms:InstanceProfile", name, args ?? new InstanceProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:dms:InstanceProfile", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InstanceProfile(name, id, options);
        }
    }

    public sealed class InstanceProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The property describes an availability zone of the instance profile.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The optional description of the instance profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The property describes an identifier for the instance profile. It is used for describing/deleting/modifying. Can be name/arn
        /// </summary>
        [Input("instanceProfileIdentifier")]
        public Input<string>? InstanceProfileIdentifier { get; set; }

        /// <summary>
        /// The property describes a name for the instance profile.
        /// </summary>
        [Input("instanceProfileName")]
        public Input<string>? InstanceProfileName { get; set; }

        /// <summary>
        /// The property describes kms key arn for the instance profile.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The property describes a network type for the instance profile.
        /// </summary>
        [Input("networkType")]
        public Input<Pulumi.AwsNative.Dms.InstanceProfileNetworkType>? NetworkType { get; set; }

        /// <summary>
        /// The property describes the publicly accessible of the instance profile
        /// </summary>
        [Input("publiclyAccessible")]
        public Input<bool>? PubliclyAccessible { get; set; }

        /// <summary>
        /// The property describes a subnet group identifier for the instance profile.
        /// </summary>
        [Input("subnetGroupIdentifier")]
        public Input<string>? SubnetGroupIdentifier { get; set; }

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

        [Input("vpcSecurityGroups")]
        private InputList<string>? _vpcSecurityGroups;

        /// <summary>
        /// The property describes vps security groups for the instance profile.
        /// </summary>
        public InputList<string> VpcSecurityGroups
        {
            get => _vpcSecurityGroups ?? (_vpcSecurityGroups = new InputList<string>());
            set => _vpcSecurityGroups = value;
        }

        public InstanceProfileArgs()
        {
        }
        public static new InstanceProfileArgs Empty => new InstanceProfileArgs();
    }
}
