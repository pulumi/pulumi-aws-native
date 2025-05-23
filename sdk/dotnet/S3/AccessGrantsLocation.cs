// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3
{
    /// <summary>
    /// The AWS::S3::AccessGrantsLocation resource is an Amazon S3 resource type hosted in an access grants instance which can be the target of S3 access grants.
    /// </summary>
    [AwsNativeResourceType("aws-native:s3:AccessGrantsLocation")]
    public partial class AccessGrantsLocation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the specified Access Grants location.
        /// </summary>
        [Output("accessGrantsLocationArn")]
        public Output<string> AccessGrantsLocationArn { get; private set; } = null!;

        /// <summary>
        /// The unique identifier for the specified Access Grants location.
        /// </summary>
        [Output("accessGrantsLocationId")]
        public Output<string> AccessGrantsLocationId { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the access grant location's associated IAM role.
        /// </summary>
        [Output("iamRoleArn")]
        public Output<string?> IamRoleArn { get; private set; } = null!;

        /// <summary>
        /// Descriptor for where the location actually points
        /// </summary>
        [Output("locationScope")]
        public Output<string?> LocationScope { get; private set; } = null!;

        /// <summary>
        /// The AWS resource tags that you are adding to the S3 Access Grants location. Each tag is a label consisting of a user-defined key and value. Tags can help you manage, identify, organize, search for, and filter resources.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.CreateOnlyTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a AccessGrantsLocation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessGrantsLocation(string name, AccessGrantsLocationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:s3:AccessGrantsLocation", name, args ?? new AccessGrantsLocationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessGrantsLocation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:s3:AccessGrantsLocation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "tags[*]",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AccessGrantsLocation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessGrantsLocation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessGrantsLocation(name, id, options);
        }
    }

    public sealed class AccessGrantsLocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the access grant location's associated IAM role.
        /// </summary>
        [Input("iamRoleArn")]
        public Input<string>? IamRoleArn { get; set; }

        /// <summary>
        /// Descriptor for where the location actually points
        /// </summary>
        [Input("locationScope")]
        public Input<string>? LocationScope { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs>? _tags;

        /// <summary>
        /// The AWS resource tags that you are adding to the S3 Access Grants location. Each tag is a label consisting of a user-defined key and value. Tags can help you manage, identify, organize, search for, and filter resources.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs>());
            set => _tags = value;
        }

        public AccessGrantsLocationArgs()
        {
        }
        public static new AccessGrantsLocationArgs Empty => new AccessGrantsLocationArgs();
    }
}
