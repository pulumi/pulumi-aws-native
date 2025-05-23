// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kms
{
    /// <summary>
    /// The ``AWS::KMS::Alias`` resource specifies a display name for a [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys). You can use an alias to identify a KMS key in the KMS console, in the [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html) operation, and in [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations), such as [Decrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html) and [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey.html).
    ///   Adding, deleting, or updating an alias can allow or deny permission to the KMS key. For details, see [ABAC for](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *Developer Guide*.
    ///   Using an alias to refer to a KMS key can help you simplify key management. For example, an alias in your code can be associated with different KMS keys in different AWS-Regions. For more information, see [Using aliases](https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html) in the *Developer Guide*.
    ///  When specifying an alias, observe the following rules.
    ///   +  Each alias is associated with one KMS key, but multiple aliases can be associated with the same KMS key.
    ///   +  The alias and its associated KMS key must be in the same AWS-account and Region.
    ///   +  The alias name must be unique in the AWS-account and Region. However, you can create aliases with the same name in different AWS-Regions. For example, you can have an ``alias/projectKey`` in multiple Regions, each of which is associated with a KMS key in its Region.
    ///   +  Each alias name must begin with ``alias/`` followed by a name, such as ``alias/exampleKey``. The alias name can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). Alias names cannot begin with ``alias/aws/``. That alias name prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
    /// 
    ///   *Regions*
    ///  KMS CloudFormation resources are available in all AWS-Regions in which KMS and CFN are supported.
    /// </summary>
    [AwsNativeResourceType("aws-native:kms:Alias")]
    public partial class Alias : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the alias name. This value must begin with ``alias/`` followed by a name, such as ``alias/ExampleAlias``. 
        ///   If you change the value of the ``AliasName`` property, the existing alias is deleted and a new alias is created for the specified KMS key. This change can disrupt applications that use the alias. It can also allow or deny access to a KMS key affected by attribute-based access control (ABAC).
        ///   The alias must be string of 1-256 characters. It can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). The alias name cannot begin with ``alias/aws/``. The ``alias/aws/`` prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
        /// </summary>
        [Output("aliasName")]
        public Output<string> AliasName { get; private set; } = null!;

        /// <summary>
        /// Associates the alias with the specified [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk). The KMS key must be in the same AWS-account and Region.
        ///  A valid key ID is required. If you supply a null or empty string value, this operation returns an error.
        ///  For help finding the key ID and ARN, see [Finding the key ID and ARN](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html#find-cmk-id-arn) in the *Developer Guide*.
        ///  Specify the key ID or the key ARN of the KMS key.
        ///  For example:
        ///   +  Key ID: ``1234abcd-12ab-34cd-56ef-1234567890ab``
        ///   +  Key ARN: ``arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab``
        ///   
        ///  To get the key ID and key ARN for a KMS key, use [ListKeys](https://docs.aws.amazon.com/kms/latest/APIReference/API_ListKeys.html) or [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html).
        /// </summary>
        [Output("targetKeyId")]
        public Output<string> TargetKeyId { get; private set; } = null!;


        /// <summary>
        /// Create a Alias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Alias(string name, AliasArgs args, CustomResourceOptions? options = null)
            : base("aws-native:kms:Alias", name, args ?? new AliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Alias(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:kms:Alias", name, null, MakeResourceOptions(options, id))
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
        /// Specifies the alias name. This value must begin with ``alias/`` followed by a name, such as ``alias/ExampleAlias``. 
        ///   If you change the value of the ``AliasName`` property, the existing alias is deleted and a new alias is created for the specified KMS key. This change can disrupt applications that use the alias. It can also allow or deny access to a KMS key affected by attribute-based access control (ABAC).
        ///   The alias must be string of 1-256 characters. It can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). The alias name cannot begin with ``alias/aws/``. The ``alias/aws/`` prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
        /// </summary>
        [Input("aliasName")]
        public Input<string>? AliasName { get; set; }

        /// <summary>
        /// Associates the alias with the specified [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk). The KMS key must be in the same AWS-account and Region.
        ///  A valid key ID is required. If you supply a null or empty string value, this operation returns an error.
        ///  For help finding the key ID and ARN, see [Finding the key ID and ARN](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html#find-cmk-id-arn) in the *Developer Guide*.
        ///  Specify the key ID or the key ARN of the KMS key.
        ///  For example:
        ///   +  Key ID: ``1234abcd-12ab-34cd-56ef-1234567890ab``
        ///   +  Key ARN: ``arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab``
        ///   
        ///  To get the key ID and key ARN for a KMS key, use [ListKeys](https://docs.aws.amazon.com/kms/latest/APIReference/API_ListKeys.html) or [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html).
        /// </summary>
        [Input("targetKeyId", required: true)]
        public Input<string> TargetKeyId { get; set; } = null!;

        public AliasArgs()
        {
        }
        public static new AliasArgs Empty => new AliasArgs();
    }
}
