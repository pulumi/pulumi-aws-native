// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PaymentCryptography
{
    /// <summary>
    /// Definition of AWS::PaymentCryptography::Key Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:paymentcryptography:Key")]
    public partial class Key : global::Pulumi.CustomResource
    {
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("exportable")]
        public Output<bool> Exportable { get; private set; } = null!;

        [Output("keyAttributes")]
        public Output<Outputs.KeyAttributes> KeyAttributes { get; private set; } = null!;

        [Output("keyCheckValueAlgorithm")]
        public Output<Pulumi.AwsNative.PaymentCryptography.KeyCheckValueAlgorithm?> KeyCheckValueAlgorithm { get; private set; } = null!;

        [Output("keyIdentifier")]
        public Output<string> KeyIdentifier { get; private set; } = null!;

        [Output("keyOrigin")]
        public Output<Pulumi.AwsNative.PaymentCryptography.KeyOrigin> KeyOrigin { get; private set; } = null!;

        [Output("keyState")]
        public Output<Pulumi.AwsNative.PaymentCryptography.KeyState> KeyState { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Key resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Key(string name, KeyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:paymentcryptography:Key", name, args ?? new KeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Key(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:paymentcryptography:Key", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Key resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Key Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Key(name, id, options);
        }
    }

    public sealed class KeyArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("exportable", required: true)]
        public Input<bool> Exportable { get; set; } = null!;

        [Input("keyAttributes", required: true)]
        public Input<Inputs.KeyAttributesArgs> KeyAttributes { get; set; } = null!;

        [Input("keyCheckValueAlgorithm")]
        public Input<Pulumi.AwsNative.PaymentCryptography.KeyCheckValueAlgorithm>? KeyCheckValueAlgorithm { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public KeyArgs()
        {
        }
        public static new KeyArgs Empty => new KeyArgs();
    }
}