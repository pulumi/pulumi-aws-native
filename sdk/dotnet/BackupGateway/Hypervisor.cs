// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.BackupGateway
{
    /// <summary>
    /// Definition of AWS::BackupGateway::Hypervisor Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:backupgateway:Hypervisor")]
    public partial class Hypervisor : global::Pulumi.CustomResource
    {
        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        [Output("hypervisorArn")]
        public Output<string> HypervisorArn { get; private set; } = null!;

        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;

        [Output("logGroupArn")]
        public Output<string?> LogGroupArn { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.HypervisorTag>> Tags { get; private set; } = null!;

        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Hypervisor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Hypervisor(string name, HypervisorArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:backupgateway:Hypervisor", name, args ?? new HypervisorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Hypervisor(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:backupgateway:Hypervisor", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Hypervisor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Hypervisor Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Hypervisor(name, id, options);
        }
    }

    public sealed class HypervisorArgs : global::Pulumi.ResourceArgs
    {
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        [Input("logGroupArn")]
        public Input<string>? LogGroupArn { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("tags")]
        private InputList<Inputs.HypervisorTagArgs>? _tags;
        public InputList<Inputs.HypervisorTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.HypervisorTagArgs>());
            set => _tags = value;
        }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public HypervisorArgs()
        {
        }
        public static new HypervisorArgs Empty => new HypervisorArgs();
    }
}