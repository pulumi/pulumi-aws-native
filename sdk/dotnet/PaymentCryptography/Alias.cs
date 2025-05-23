// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PaymentCryptography
{
    /// <summary>
    /// Definition of AWS::PaymentCryptography::Alias Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:paymentcryptography:Alias")]
    public partial class Alias : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A friendly name that you can use to refer to a key. The value must begin with `alias/` .
        /// 
        /// &gt; Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in AWS CloudTrail logs and other output.
        /// </summary>
        [Output("aliasName")]
        public Output<string> AliasName { get; private set; } = null!;

        /// <summary>
        /// The `KeyARN` of the key associated with the alias.
        /// </summary>
        [Output("keyArn")]
        public Output<string?> KeyArn { get; private set; } = null!;


        /// <summary>
        /// Create a Alias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Alias(string name, AliasArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:paymentcryptography:Alias", name, args ?? new AliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Alias(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:paymentcryptography:Alias", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "aliasName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Alias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Alias Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Alias(name, id, options);
        }
    }

    public sealed class AliasArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A friendly name that you can use to refer to a key. The value must begin with `alias/` .
        /// 
        /// &gt; Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in AWS CloudTrail logs and other output.
        /// </summary>
        [Input("aliasName")]
        public Input<string>? AliasName { get; set; }

        /// <summary>
        /// The `KeyARN` of the key associated with the alias.
        /// </summary>
        [Input("keyArn")]
        public Input<string>? KeyArn { get; set; }

        public AliasArgs()
        {
        }
        public static new AliasArgs Empty => new AliasArgs();
    }
}
