// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Shield
{
    /// <summary>
    /// Enables AWS Shield Advanced for a specific AWS resource. The resource can be an Amazon CloudFront distribution, Amazon Route 53 hosted zone, AWS Global Accelerator standard accelerator, Elastic IP Address, Application Load Balancer, or a Classic Load Balancer. You can protect Amazon EC2 instances and Network Load Balancers by association with protected Amazon EC2 Elastic IP addresses.
    /// </summary>
    [Obsolete(@"Protection is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:shield:Protection")]
    public partial class Protection : global::Pulumi.CustomResource
    {
        [Output("applicationLayerAutomaticResponseConfiguration")]
        public Output<Outputs.ProtectionApplicationLayerAutomaticResponseConfiguration?> ApplicationLayerAutomaticResponseConfiguration { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Names (ARNs) of the health check to associate with the protection.
        /// </summary>
        [Output("healthCheckArns")]
        public Output<ImmutableArray<string>> HealthCheckArns { get; private set; } = null!;

        /// <summary>
        /// Friendly name for the Protection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN (Amazon Resource Name) of the protection.
        /// </summary>
        [Output("protectionArn")]
        public Output<string> ProtectionArn { get; private set; } = null!;

        /// <summary>
        /// The unique identifier (ID) of the protection.
        /// </summary>
        [Output("protectionId")]
        public Output<string> ProtectionId { get; private set; } = null!;

        /// <summary>
        /// The ARN (Amazon Resource Name) of the resource to be protected.
        /// </summary>
        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;

        /// <summary>
        /// One or more tag key-value pairs for the Protection object.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ProtectionTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Protection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Protection(string name, ProtectionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:shield:Protection", name, args ?? new ProtectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Protection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:shield:Protection", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Protection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Protection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Protection(name, id, options);
        }
    }

    public sealed class ProtectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicationLayerAutomaticResponseConfiguration")]
        public Input<Inputs.ProtectionApplicationLayerAutomaticResponseConfigurationArgs>? ApplicationLayerAutomaticResponseConfiguration { get; set; }

        [Input("healthCheckArns")]
        private InputList<string>? _healthCheckArns;

        /// <summary>
        /// The Amazon Resource Names (ARNs) of the health check to associate with the protection.
        /// </summary>
        public InputList<string> HealthCheckArns
        {
            get => _healthCheckArns ?? (_healthCheckArns = new InputList<string>());
            set => _healthCheckArns = value;
        }

        /// <summary>
        /// Friendly name for the Protection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN (Amazon Resource Name) of the resource to be protected.
        /// </summary>
        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.ProtectionTagArgs>? _tags;

        /// <summary>
        /// One or more tag key-value pairs for the Protection object.
        /// </summary>
        public InputList<Inputs.ProtectionTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ProtectionTagArgs>());
            set => _tags = value;
        }

        public ProtectionArgs()
        {
        }
        public static new ProtectionArgs Empty => new ProtectionArgs();
    }
}